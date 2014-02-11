package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/appio/websocket"
	"github.com/codegangsta/martini"
	"github.com/steveoc64/ActionFront/gamedatadb"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Command line flags
var (
	port = flag.Int("port", 8080, "port to access the unitEditor")
)

// Convert a GO structure to a map[string]interface{}
func toMap(thing interface{}) map[string]interface{} {
	var jsonThing, err = json.Marshal(thing)
	if err != nil {
		panic(err)
	}
	var retval = map[string]interface{}{}
	json.Unmarshal(jsonThing, &retval)
	return retval
}

// init the DB, and return a ref to the GameData collection
func initDB() *db.Col {
	rand.Seed(time.Now().UTC().UnixNano())

	// Create and open database
	os.RemoveAll("database")
	dir := "database"
	os.MkdirAll(dir, os.ModePerm)

	myDB, err := db.OpenDB(dir)
	if err != nil {
		panic(err)
	}

	if err := myDB.Create("GameData", 1); err == nil {
		// This is a fresh DB so insert some default unit types
		gameData := myDB.Use("GameData")

		gamedatadb.CreateGameData(gameData)
	}
	myDB.Scrub("GameData")
	return myDB.Use("GameData")
}

// Pool of websocket connections
var connections map[*websocket.Conn]bool

type messageFormat struct {
	Action string
	Entity string
	Data   interface{}
}

// Send message to all known connections
func sendAll(msg []byte) {
	for conn := range connections {
		sendMsg(conn, msg)
	}
}

// Send a message to all connections except this one
func sendOthers(fromConn *websocket.Conn, msg []byte) {
	for conn := range connections {
		if conn != fromConn {
			sendMsg(conn, msg)
		}
	}
}

// Send a message to a specified connection
func sendMsg(conn *websocket.Conn, msg []byte) {
	if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		delete(connections, conn)
		return
	}
}

// Get all GameData records into a slice of bytes
func getsdfsdfAllUnitTypes(col *db.Col) map[uint64]struct{} {
	queryResult := make(map[uint64]struct{}) // query result (document IDs) goes into map keys

	if err := db.EvalAllIDs(col, &queryResult); err != nil {
		panic(err)
	}
	return queryResult
}

// Martini handler for incoming socket request - runs forever until socket connection is closed
//
// A DataSocket handler defines a standard protocol, which automates the realtime updates
// for multiple clients collaborating on editing a number of documents.
//
// TODO - need to document the wire protocol.
//   General idea is to wrap JSON document records in a standard envelope with an Action and an Entity
//
//   The Action opcode allows for CRUD requests from the client and realtime updates from the server
//   The Entity opcode allows for multiplexing several doc types through the same NoSQL collection
//
// On the backend, a single handler will server multiple clients for a single collection.
// Fire up extra dataSocketHandlers for different collections
//
// On the frontend, a single call to DataSocket.connect($scope) will bind the current scope of an ng-grid
// for a single Entity type to the global DataSocket. This includes kicking off the first LIST request.

func dataSocketHandler(w http.ResponseWriter, r *http.Request, gameData *db.Col) {

	// Perform handshake and upgrade connection
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println(err)
		return
	}

	// record new connection in our map
	connections[conn] = true
	log.Println("New GameData Websocket connection ", connections)
	defer conn.Close()

	var RxMsg map[string]interface{}
	var myGameData map[string]interface{}

	// loop forever
	for {
		_, msg, err := conn.ReadMessage()
		// Connection died - remove it from the list of connections
		if err != nil {
			delete(connections, conn)
			log.Println("Removed connection ", connections)
			return
		}

		// Got a message on the socket !!!
		log.Printf("<- %s", msg)
		json.Unmarshal(msg, &RxMsg)
		theEntity := RxMsg["Entity"].(string)

		switch RxMsg["Action"] {
		case "List":
			// List all records for the given entity
			log.Println("LIST request:", theEntity)
			resultIDs := make(map[uint64]struct{}) // query result (document IDs) goes into map keys
			if err := db.EvalAllIDs(gameData, &resultIDs); err != nil {
				panic(err)
			}
			results := make([]interface{}, 0)

			for id := range resultIDs {
				gameData.Read(id, &myGameData)

				if myGameData["Type"].(string) == theEntity {
					theID := myGameData["@id"].(string)
					theData := myGameData["Data"].(map[string]interface{})
					theData["@id"] = theID
					results = append(results, theData)
				}
			}
			msg, _ := json.Marshal(messageFormat{"List", theEntity, results})
			sendMsg(conn, msg)

		case "Update":
			log.Println("UPDATE request:", theEntity, RxMsg["Data"])
			myGameData = RxMsg["Data"].(map[string]interface{})
			docID := myGameData["@id"]
			delete(myGameData, "@id") // strip the ID out of this record
			myDocID, _ := strconv.ParseUint(docID.(string), 0, 64)
			switch myDocID {
			case 0:
				// Insert as new record
				log.Println("Add NEW Record", myGameData)
				if myDocID, err = gameData.Insert(gamedatadb.DataMap(theEntity, myGameData)); err != nil {
					panic(err)
				}
				log.Printf("Inserted as ID %d", myDocID)
				gameData.Read(myDocID, &myGameData)
				msg, _ := json.Marshal(messageFormat{"Update", theEntity, myGameData})
				sendAll(msg)

			default:
				// Write to existing record
				log.Println("Write Record ID", myDocID, myGameData)
				if err := gameData.Update(myDocID, gamedatadb.DataMap(theEntity, myGameData)); err != nil {
					panic(err)
				}
				//gameData.Read(myDocID, &myGameData)
				myGameData["@id"] = docID
				msg, _ := json.Marshal(messageFormat{"Update", theEntity, myGameData})
				sendOthers(conn, msg)
			}

		case "Delete":
			log.Println("DELETE request:", RxMsg["Entity"])

		case "Get":
			log.Println("GET request:", RxMsg["Entity"])

		default:
			log.Println("WTF ?", RxMsg)
		}
	}
}

// Main loop
func main() {

	flag.Parse()

	connections = make(map[*websocket.Conn]bool)

	// Classic defaults for webserver - serve up files from public dir
	m := martini.Classic()
	m.Map(initDB())
	m.Get("/GameData", dataSocketHandler)

	// Run the actual webserver
	addr := fmt.Sprintf(":%d", *port)
	log.Println("ActionFront GameData Editor starting on port ", addr)

	http.ListenAndServe(addr, m)
}
