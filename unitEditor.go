package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/appio/websocket"
	"github.com/codegangsta/martini"
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

type UnitType struct {
	Name      string
	Rating    string
	Men       uint16
	Size      string
	Firepower int8
	DrillBook string
}

func (ut UnitType) toJSON() []byte {
	var retval, err = json.Marshal(ut)
	if err != nil {
		panic(err)
	}
	return retval
}

func toMap(thing interface{}) map[string]interface{} {
	var jsonThing, err = json.Marshal(thing)
	if err != nil {
		panic(err)
	}
	var retval = map[string]interface{}{}
	json.Unmarshal(jsonThing, &retval)
	return retval
}

// init the DB, and return a ref to the UnitTypes collection
func initUnitTypesDB() *db.Col {
	rand.Seed(time.Now().UTC().UnixNano())

	// Create and open database
	os.RemoveAll("database")
	dir := "database/types"
	os.MkdirAll(dir, os.ModePerm)

	myDB, err := db.OpenDB(dir)
	if err != nil {
		panic(err)
	}

	if err := myDB.Create("UnitTypes", 1); err == nil {
		// This is a fresh DB so insert some default unit types
		ut := myDB.Use("UnitTypes")

		ut.Insert(toMap(UnitType{"French Ligne", "Veteran", 720, "2L 1E", 10, "French Veteran"}))
		ut.Insert(toMap(UnitType{"French Legere", "Elite", 720, "3E", 12, "French Veteran"}))
		ut.Insert(toMap(UnitType{"French Provisional", "Regular", 720, "3L", 10, "French Conscript"}))
		ut.Insert(toMap(UnitType{"French Conscript", "Conscript", 720, "3L", 8, "French Conscript"}))
		ut.Insert(toMap(UnitType{"Prussian Line", "CrackLine", 960, "4L 1S", 10, "Prussian"}))
		ut.Insert(toMap(UnitType{"Prussian Fusilier", "CrackLine", 960, "2L 2E", 12, "Prussian"}))
		ut.Insert(toMap(UnitType{"Prussian Reserve", "Regular", 960, "4L 1S", 8, "Prussian"}))
		ut.Insert(toMap(UnitType{"Prussian Landwehr", "Landwehr", 960, "4L", 6, "Landwehr"}))
	}
	myDB.Scrub("UnitTypes")
	return myDB.Use("UnitTypes")
}

// Pool of connections
var connections map[*websocket.Conn]bool

// Send message to all known connections
func sendAll(msg []byte) {
	for conn := range connections {
		sendMsg(conn, msg)
	}
}

func sendOthers(fromConn *websocket.Conn, msg []byte) {
	for conn := range connections {
		if conn != fromConn {
			sendMsg(conn, msg)
		}
	}
}

func sendMsg(conn *websocket.Conn, msg []byte) {
	if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		delete(connections, conn)
		return
	}
}

// Get all unit types into a slice of bytes
func getAllUnitTypes(col *db.Col) map[uint64]struct{} {
	queryResult := make(map[uint64]struct{}) // query result (document IDs) goes into map keys

	if err := db.EvalAllIDs(col, &queryResult); err != nil {
		panic(err)
	}
	return queryResult
}

// Martini handler for incoming socket request - runs forever until socket connection is closed
func wsHandler(w http.ResponseWriter, r *http.Request, unitTypes *db.Col) {

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
	log.Println("New UnitTypes Websocket connection ", connections)
	defer conn.Close()

	// kick the new connection off by sending a list of unit types in a single message
	allUnitTypeIds := getAllUnitTypes(unitTypes)
	var utMap interface{}
	var allUnits []interface{}

	for id := range allUnitTypeIds {
		unitTypes.Read(id, &utMap)
		allUnits = append(allUnits, utMap)
	}
	//log.Printf("All units %+v", allUnits)
	allUnitsMsg, _ := json.Marshal(allUnits)
	sendMsg(conn, allUnitsMsg)

	var myUnitData map[string]interface{}

	// loop forever
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			delete(connections, conn)
			log.Println("Removed connection ", connections)
			return
		}
		log.Printf("Received Message %s", msg)
		json.Unmarshal(msg, &myUnitData)
		//log.Printf("myUnitData %+v", myUnitData)
		docID := myUnitData["@id"]
		delete(myUnitData, "@id")
		//log.Printf("myUnitData truncated %+v", myUnitData)
		myDocID, _ := strconv.ParseUint(docID.(string), 0, 64)
		//log.Printf("myDoc ID as uint64 = %d", myDocID)

		// Insert or Update or Delete ?
		switch myDocID {
		case 0:
			log.Println("Insert New Record")
			if myDocID, err = unitTypes.Insert(myUnitData); err != nil {
				panic(err)
			}
			log.Printf("Inserted as ID %d", myDocID)
			unitTypes.Read(myDocID, &utMap)
			msg, _ := json.Marshal(utMap)
			sendAll(msg)
		default:
			switch myUnitData["Name"] {
			case "":
				log.Println("Deleting Record", myDocID)
				unitTypes.Delete(myDocID)
				sendAll([]byte(fmt.Sprintf("%d", myDocID)))
			default:
				log.Println("Update Record", myDocID)
				if err := unitTypes.Update(myDocID, myUnitData); err != nil {
					panic(err)
				}
				// Tell other connected clients about the updated UnitType
				sendOthers(conn, msg)
			}
		}
	}
}

// Main loop
func main() {

	flag.Parse()

	connections = make(map[*websocket.Conn]bool)

	// Classic defaults for webserver - serve up files from public dir
	m := martini.Classic()
	m.Map(initUnitTypesDB())
	m.Get("/UnitTypeSocket", wsHandler)

	// Run the actual webserver
	addr := fmt.Sprintf(":%d", *port)
	log.Println("ActionFront Unit Editor starting on port ", addr)

	http.ListenAndServe(addr, m)
}
