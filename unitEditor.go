package main

import (
	"bytes"
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

type InfantryType struct {
	Nation    string
	From      uint16
	To        uint16
	Name      string
	Rating    string
	DrillBook string
	Layout    string
	Fire      int8
	Elite     int8
	Equip     string
	Skirmish  string
	Street    string
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

		// Range of French line infantry types for various years
		ut.Insert(toMap(InfantryType{"France", 1805, 1807, "Elite Ligne", "Elite", "French", "5L 1S", 0, 2, "Musket", "Excellent", "Excellent"}))
		ut.Insert(toMap(InfantryType{"France", 1805, 1807, "Crack Ligne", "CrackLine", "French", "5L 1S", 0, 2, "Musket", "Excellent", "Excellent"}))
		ut.Insert(toMap(InfantryType{"France", 1805, 1807, "Veteran Ligne", "Veteran", "French", "5L 1S", 0, 2, "Musket", "Average", "Good"}))

		ut.Insert(toMap(InfantryType{"France", 1808, 1812, "Elite Ligne", "Elite", "French", "3L 1E", 0, 2, "Musket", "Excellent", "Good"}))
		ut.Insert(toMap(InfantryType{"France", 1808, 1812, "Crack Ligne", "CrackLine", "French", "3L 1E", 0, 2, "Musket", "Excellent", "Good"}))
		ut.Insert(toMap(InfantryType{"France", 1808, 1812, "Veteran Ligne", "Veteran", "French", "3L 1E", 0, 2, "Musket", "Good", "Good"}))
		ut.Insert(toMap(InfantryType{"France", 1808, 1812, "Regular Ligne", "Veteran", "French", "3L 1E", 0, 2, "Musket", "Average", "Good"}))
		ut.Insert(toMap(InfantryType{"France", 1808, 1812, "Conscript Ligne", "Veteran", "Conscript", "4L", 0, 2, "Musket", "Poor", "Good"}))

		ut.Insert(toMap(InfantryType{"France", 1813, 1814, "Veteran Ligne", "Veteran", "French", "2L 1E", 0, 2, "Musket", "Average", "Good"}))
		ut.Insert(toMap(InfantryType{"France", 1813, 1814, "Conscript Ligne", "Conscript", "Conscript", "3L", 0, 2, "Musket", "Poor", "Poor"}))
		ut.Insert(toMap(InfantryType{"France", 1813, 1814, "Provisional Ligne", "Veteran", "French", "2L", 0, 2, "Musket", "Poor", "Poor"}))

		ut.Insert(toMap(InfantryType{"France", 1815, 1815, "Elites", "Elite", "French", "2L 1E", 0, 2, "Musket", "Excellent", "Good"}))
		ut.Insert(toMap(InfantryType{"France", 1815, 1815, "Crack Ligne", "CrackLine", "French", "2L 1E", 0, 2, "Musket", "Excellent", "Good"}))
		ut.Insert(toMap(InfantryType{"France", 1815, 1815, "Veteran Ligne", "Veteran", "French", "2L 1E", 0, 2, "Musket", "Good", "Good"}))

		ut.Insert(toMap(InfantryType{"France", 1805, 1807, "Elite Legere", "Elite", "French", "5E", 0, 2, "Musket", "Excellent", "Excellent"}))
		ut.Insert(toMap(InfantryType{"France", 1805, 1807, "Crack Legere", "CrackLine", "French", "5E", 0, 2, "Musket", "Excellent", "Excellent"}))

		ut.Insert(toMap(InfantryType{"France", 1808, 1812, "Elite Legere", "Elite", "French", "4E", 0, 2, "Musket", "Excellent", "Good"}))
		ut.Insert(toMap(InfantryType{"France", 1808, 1812, "Crack Legere", "CrackLine", "French", "4E", 0, 2, "Musket", "Excellent", "Good"}))
		ut.Insert(toMap(InfantryType{"France", 1808, 1812, "Veteran Legere", "Elite", "French", "3E", 0, 2, "Musket", "Good", "Good"}))
		ut.Insert(toMap(InfantryType{"France", 1808, 1812, "Regular Legere", "CrackLine", "French", "3E", 0, 2, "Musket", "Average", "Average"}))

		ut.Insert(toMap(InfantryType{"France", 1813, 1814, "Crack Legere", "CrackLine", "French", "3E", 0, 2, "Musket", "Excellent", "Good"}))
		ut.Insert(toMap(InfantryType{"France", 1813, 1814, "Veteran Legere", "Veteran", "French", "3E", 0, 2, "Musket", "Good", "Good"}))
		ut.Insert(toMap(InfantryType{"France", 1813, 1814, "Conscript Legere", "Conscript", "French", "3E", 0, 2, "Musket", "Poor", "Poor"}))

		ut.Insert(toMap(InfantryType{"France", 1815, 1815, "Elite Legere", "Elite", "French", "3E", 0, 2, "Musket", "Excellent", "Good"}))
		ut.Insert(toMap(InfantryType{"France", 1815, 1815, "Veteran Legere", "Veteran", "French", "3E", 0, 2, "Musket", "Excellent", "Good"}))
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
		if bytes.Equal(msg, []byte("init")) {
			log.Println("Received INIT message - send all records")
			allUnitTypeIds := getAllUnitTypes(unitTypes)

			allUnits = make([]interface{}, 0)
			for id := range allUnitTypeIds {
				unitTypes.Read(id, &utMap)
				allUnits = append(allUnits, utMap)
			}
			allUnitsMsg, _ := json.Marshal(allUnits)
			sendMsg(conn, allUnitsMsg)
		} else {
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
}

// Main loop
func main() {

	flag.Parse()

	connections = make(map[*websocket.Conn]bool)

	// Classic defaults for webserver - serve up files from public dir
	m := martini.Classic()
	m.Map(initUnitTypesDB())
	m.Get("/Socket", wsHandler)

	// Run the actual webserver
	addr := fmt.Sprintf(":%d", *port)
	log.Println("ActionFront Unit Editor starting on port ", addr)

	http.ListenAndServe(addr, m)
}
