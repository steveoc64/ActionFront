package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/appio/websocket"
	"github.com/codegangsta/martini"
	"github.com/steveoc64/ActionFront/gamedatadb"
	"github.com/steveoc64/tiedot/db"
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

// Make the LIST cache a global object
var ListCache map[string]interface{}

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
	ListCache = make(map[string]interface{})

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

// For a given entity, return a slice of bytes, being a JSON representation of that list
func getList(col *db.Col, theEntity string) (messageFormat, bool) {
	var myData map[string]interface{}

	if records, ok := ListCache[theEntity]; ok {
		return records.(messageFormat), true
	}

	// Not cached, so Build a new result set using tiedot embedded query processor
	queryStr := `{"eq": "` + theEntity + `", "in": ["Type"]}`
	var query interface{}
	json.Unmarshal([]byte(queryStr), &query)
	queryResult := make(map[uint64]struct{}) // query result (document IDs) goes into map keys

	if err := db.EvalQuery(query, col, &queryResult); err != nil {
		panic(err)
	}

	results := make([]interface{}, 0)

	for id := range queryResult {
		col.Read(id, &myData)

		theID := myData["@id"].(string)
		theData := myData["Data"].(map[string]interface{})
		theData["@id"] = theID
		results = append(results, theData)
	}
	records := messageFormat{"List", theEntity, results}
	ListCache[theEntity] = records
	return records, false
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

func tilde(c bool) string {
	if c {
		return "~"
	} else {
		return ""
	}
}

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
		//log.Printf("<- %s", msg)
		json.Unmarshal(msg, &RxMsg)

		switch RxMsg["Action"] {
		case "MList":
			// Return an array of lists
			startTime := time.Now()
			Entities := RxMsg["Entities"].([]interface{})
			EntityNames := ""
			var mmsg []messageFormat

			for _, Entity := range Entities {
				theEntity := Entity.(string)
				EntityNames = EntityNames + theEntity + " "
				msg, cached := getList(gameData, theEntity)
				log.Printf("MLIST request: %s (%s)%s", theEntity, time.Since(startTime), tilde(cached))
				startTime = time.Now()

				mmsg = append(mmsg, msg)
			}
			msg, _ = json.Marshal(messageFormat{"MList", EntityNames, mmsg})
			sendMsg(conn, msg)

		case "List":
			// List all records for the given entity
			startTime := time.Now()
			theEntity := RxMsg["Entity"].(string)

			records, cached := getList(gameData, theEntity)
			log.Printf("LIST request: %s (%s)%s", theEntity, time.Since(startTime), tilde(cached))
			msg, _ = json.Marshal(records)
			sendMsg(conn, msg)

		case "Add":
			// Invalidate the LIST cache for this entity before we do any updates
			theEntity := RxMsg["Entity"].(string)

			delete(ListCache, theEntity)
			log.Println("ADD request:", theEntity, RxMsg["Data"])

			myGameData = RxMsg["Data"].(map[string]interface{})
			myDocID, err := gameData.Insert(gamedatadb.DataMap(theEntity, myGameData))
			if err != nil {
				panic(err)
			}
			delete(ListCache, theEntity)
			log.Printf("Inserted as ID %d", myDocID)
			myGameData["@id"] = strconv.FormatUint(myDocID, 10)
			msg, _ := json.Marshal(messageFormat{"Add", theEntity, myGameData})
			sendAll(msg)

		case "Update":
			// Invalidate the LIST cache for this entity before we do any updates
			theEntity := RxMsg["Entity"].(string)

			delete(ListCache, theEntity)

			log.Println("UPDATE request:", theEntity, RxMsg["Data"])

			myGameData = RxMsg["Data"].(map[string]interface{})
			docID := myGameData["@id"]
			delete(myGameData, "@id") // strip the ID out of this record
			myDocID, _ := strconv.ParseUint(docID.(string), 0, 64)
			if myDocID > 0 {
				// Write to existing record
				log.Println("Write Record ID", myDocID, myGameData)
				if err := gameData.Update(myDocID, gamedatadb.DataMap(theEntity, myGameData)); err != nil {
					panic(err)
				}
				delete(ListCache, theEntity)
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
