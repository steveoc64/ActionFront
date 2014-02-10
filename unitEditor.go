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

type Infantry struct {
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

type Cavalry struct {
	Nation    string
	From      uint16
	To        uint16
	Name      string
	Rating    string
	Shock     uint16
	Squadrons uint8
	Move      string
	Skirmish  string
}

type Artillery struct {
	Nation   string
	From     uint16
	To       uint16
	Name     string
	Rating   string
	Class    uint8
	Guns     string
	HW       string
	Sections uint8
	Horse    bool
}

// Names of fields here shortened to help make the JSON daatbase more sensible
type Drill struct {
	EF uint8 // Efficiency. Range 1-10. value 1 = 10%, value 10 = 100%
	FR uint8 // Max frontage of this unit in line
	SS uint8 // How many Semi skirmish elements allowed
	SK uint8 // How many full skirmish elements allowed
}

type DrillBook struct {
	Name    string
	Entries map[string]Drill
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

func UnitTypeMap(mainType string, thing interface{}) map[string]interface{} {
	var jsonThing, err = json.Marshal(thing)
	if err != nil {
		panic(err)
	}
	var thingMap = map[string]interface{}{}
	json.Unmarshal(jsonThing, &thingMap)

	var retval = map[string]interface{}{}
	retval["Type"] = mainType
	retval["Data"] = thingMap
	return retval
}

// init the DB, and return a ref to the UnitTypes collection
func initUnitTypesDB() *db.Col {
	rand.Seed(time.Now().UTC().UnixNano())

	// Create and open database
	os.RemoveAll("database")
	dir := "database"
	os.MkdirAll(dir, os.ModePerm)

	myDB, err := db.OpenDB(dir)
	if err != nil {
		panic(err)
	}

	if err := myDB.Create("UnitTypes", 1); err == nil {
		// This is a fresh DB so insert some default unit types
		ut := myDB.Use("UnitTypes")

		// Create some DrillBooks
		ut.Insert(UnitTypeMap("Drill", DrillBook{"Conscript", map[string]Drill{
			"Line":           Drill{5, 2, 0, 0},
			"MarchColumn":    Drill{8, 1, 0, 0},
			"AttackColumn":   Drill{7, 1, 1, 0},
			"ClosedColumn":   Drill{6, 1, 0, 0},
			"ScreenedColumn": Drill{5, 1, 1, 0}}}))

		ut.Insert(UnitTypeMap("Drill", DrillBook{"French", map[string]Drill{
			"Line":         Drill{8, 3, 0, 0},
			"MarchColumn":  Drill{10, 1, 0, 0},
			"AttackColumn": Drill{9, 1, 1, 1},
			"ClosedColumn": Drill{8, 1, 0, 1},
			"Square":       Drill{7, 1, 0, 1}}}))

		ut.Insert(UnitTypeMap("Drill", DrillBook{"LightInfantry", map[string]Drill{
			"FullSK":       Drill{7, 12, 0, 0},
			"HalfSK":       Drill{8, 6, 0, 0},
			"Screen":       Drill{8, 6, 0, 0},
			"Line":         Drill{7, 3, 0, 0},
			"MarchColumn":  Drill{10, 1, 0, 0},
			"AttackColumn": Drill{9, 1, 1, 1},
			"ClosedColumn": Drill{8, 1, 0, 1},
			"Square":       Drill{6, 1, 0, 1}}}))

		ut.Insert(UnitTypeMap("Drill", DrillBook{"Prussian", map[string]Drill{
			"Line":         Drill{7, 4, 0, 1},
			"Oblique":      Drill{6, 4, 0, 1},
			"ScreenedLine": Drill{7, 3, 1, 1},
			"MarchColumn":  Drill{9, 1, 0, 0},
			"AttackColumn": Drill{8, 2, 1, 1},
			"ClosedColumn": Drill{7, 2, 0, 0},
			"Square":       Drill{6, 1, 0, 0}}}))

		ut.Insert(UnitTypeMap("Drill", DrillBook{"Russian", map[string]Drill{
			"Line":         Drill{7, 2, 0, 0},
			"MarchColumn":  Drill{9, 1, 0, 0},
			"AttackColumn": Drill{8, 1, 0, 0},
			"Square":       Drill{6, 1, 0, 0}}}))

		ut.Insert(UnitTypeMap("Drill", DrillBook{"Austrian", map[string]Drill{
			"Line":         Drill{7, 4, 0, 1},
			"ScreenedLine": Drill{7, 4, 1, 1},
			"MarchColumn":  Drill{9, 1, 0, 0},
			"AttackColumn": Drill{8, 2, 1, 1},
			"ClosedColumn": Drill{7, 2, 0, 0},
			"Square":       Drill{6, 1, 0, 0}}}))

		// Range of French line infantry types for various years
		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1805, 1807, "Elite Ligne", "Elite", "French", "5L 1S", 0, 2, "Musket", "Excellent", "Excellent"}))
		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1805, 1807, "Crack Ligne", "CrackLine", "French", "5L 1S", 0, 2, "Musket", "Excellent", "Excellent"}))
		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1805, 1807, "Veteran Ligne", "Veteran", "French", "5L 1S", 0, 2, "Musket", "Average", "Good"}))

		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1808, 1812, "Elite Ligne", "Elite", "French", "3L 1E", 0, 2, "Musket", "Excellent", "Good"}))
		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1808, 1812, "Crack Ligne", "CrackLine", "French", "3L 1E", 0, 2, "Musket", "Excellent", "Good"}))
		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1808, 1812, "Veteran Ligne", "Veteran", "French", "3L 1E", 0, 2, "Musket", "Good", "Good"}))
		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1808, 1812, "Regular Ligne", "Veteran", "French", "3L 1E", 0, 2, "Musket", "Average", "Good"}))
		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1808, 1812, "Conscript Ligne", "Veteran", "Conscript", "4L", 0, 2, "Musket", "Poor", "Good"}))

		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1813, 1814, "Veteran Ligne", "Veteran", "French", "2L 1E", 0, 2, "Musket", "Average", "Good"}))
		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1813, 1814, "Conscript Ligne", "Conscript", "Conscript", "3L", 0, 2, "Musket", "Poor", "Poor"}))
		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1813, 1814, "Provisional Ligne", "Veteran", "French", "2L", 0, 2, "Musket", "Poor", "Poor"}))

		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1815, 1815, "Elites", "Elite", "French", "2L 1E", 0, 2, "Musket", "Excellent", "Good"}))
		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1815, 1815, "Crack Ligne", "CrackLine", "French", "2L 1E", 0, 2, "Musket", "Excellent", "Good"}))
		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1815, 1815, "Veteran Ligne", "Veteran", "French", "2L 1E", 0, 2, "Musket", "Good", "Good"}))

		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1805, 1807, "Elite Legere", "Elite", "French", "5E", 0, 2, "Musket", "Excellent", "Excellent"}))
		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1805, 1807, "Crack Legere", "CrackLine", "French", "5E", 0, 2, "Musket", "Excellent", "Excellent"}))

		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1808, 1812, "Elite Legere", "Elite", "French", "4E", 0, 2, "Musket", "Excellent", "Good"}))
		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1808, 1812, "Crack Legere", "CrackLine", "French", "4E", 0, 2, "Musket", "Excellent", "Good"}))
		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1808, 1812, "Veteran Legere", "Elite", "French", "3E", 0, 2, "Musket", "Good", "Good"}))
		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1808, 1812, "Regular Legere", "CrackLine", "French", "3E", 0, 2, "Musket", "Average", "Average"}))

		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1813, 1814, "Crack Legere", "CrackLine", "French", "3E", 0, 2, "Musket", "Excellent", "Good"}))
		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1813, 1814, "Veteran Legere", "Veteran", "French", "3E", 0, 2, "Musket", "Good", "Good"}))
		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1813, 1814, "Conscript Legere", "Conscript", "French", "3E", 0, 2, "Musket", "Poor", "Poor"}))

		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1815, 1815, "Elite Legere", "Elite", "French", "3E", 0, 2, "Musket", "Excellent", "Good"}))
		ut.Insert(UnitTypeMap("Infantry", Infantry{"France", 1815, 1815, "Veteran Legere", "Veteran", "French", "3E", 0, 2, "Musket", "Excellent", "Good"}))

		/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		// Add some Cavalry
		// French
		ut.Insert(UnitTypeMap("Cavalry", Cavalry{"France", 1804, 1812, "Hussar", "Elite", 18, 3, "Light", "Good"}))
		ut.Insert(UnitTypeMap("Cavalry", Cavalry{"France", 1813, 1814, "Hussar", "CrackLine", 16, 4, "Light", "Good"}))
		ut.Insert(UnitTypeMap("Cavalry", Cavalry{"France", 1815, 1815, "Hussar", "Elite", 18, 3, "Light", "Good"}))

		/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		// Add some Artillery
		// French Line
		ut.Insert(UnitTypeMap("Artillery", Artillery{"France", 1812, 1815, "Young Guard", "Grenadier", 1, "6pdr", "5.5\"", 4, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"France", 1792, 1815, "Line Reserve", "CrackLine", 1, "12pdr", "6\"", 4, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"France", 1791, 1806, "Line", "CrackLine", 1, "8pdr", "5.5\"", 3, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"France", 1807, 1815, "Line", "CrackLine", 1, "6pdr", "5.5\"", 4, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"France", 1812, 1812, "Regimental", "Veteran", 1, "4pdr", "", 1, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"France", 1813, 1814, "Regimental", "Veteran", 1, "6pdr", "", 1, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"France", 1791, 1809, "Horse", "Elite", 1, "8pdr", "5.5\"", 3, true}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"France", 1810, 1815, "Horse", "Elite", 1, "6pdr", "5.5\"", 3, true}))

		// French Guard
		ut.Insert(UnitTypeMap("Artillery", Artillery{"France", 1804, 1805, "Guard Horse", "OldGuard", 0, "8pdr", "5.5\"", 4, true}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"France", 1806, 1806, "Volante", "OldGuard", 0, "8pdr", "5.5\"", 3, true}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"France", 1807, 1815, "Volante", "OldGuard", 0, "6pdr", "5.5\"", 4, true}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"France", 1808, 1815, "Guard Reserve", "OldGuard", 0, "12pdr", "6\"", 4, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"France", 1810, 1815, "Guard Divisional", "OldGuard", 0, "6pdr", "5.5\"", 4, true}))

		// British
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Britain", 1792, 1809, "Royal Foot", "Grenadier", 1, "6pdr", "5.5\"", 3, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Britain", 1810, 1815, "Royal Foot", "Grenadier", 1, "9pdr", "5.5\"", 3, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Britain", 1792, 1809, "Royal Horse", "Grenadier", 1, "6pdr", "5.5\"", 3, true}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Britain", 1810, 1815, "Royal Horse", "Grenadier", 1, "9pdr", "5.5\"", 3, true}))

		// Russian
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Russia", 1792, 1810, "Guard", "Guard", 1, "12pdr", "18pdr L", 5, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Russia", 1811, 1815, "Guard", "Guard", 1, "12pdr", "18pdr L", 6, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Russia", 1792, 1810, "Guard Horse", "Grenadier", 1, "6pdr", "9pdr L", 5, true}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Russia", 1811, 1815, "Guard Horse", "Grenadier", 1, "6pdr", "9pdr L", 4, true}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Russia", 1792, 1815, "Line Heavy", "Elite", 2, "12pdr", "18pdr L", 6, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Russia", 1792, 1815, "Line Light", "Elite", 2, "6pdr", "9pdr L", 6, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Russia", 1792, 1809, "Battalion Guns", "Veteran", 2, "6pdr", "", 1, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Russia", 1792, 1815, "Line", "CrackLine", 2, "6pdr", "9pdr L", 6, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Russia", 1792, 1815, "Flying Cossack", "Conscript", 3, "2pdr", "", 5, true}))

		// Prussian
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Preussen", 1792, 1815, "Guard Horse", "Grenadier", 1, "6pdr", "7pdr", 4, true}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Preussen", 1792, 1815, "Line", "Veteran", 2, "6pdr", "10pdr", 4, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Preussen", 1792, 1815, "Reserve", "Veteran", 2, "12pdr", "10pdr", 4, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Preussen", 1792, 1807, "Battalion Guns", "CrackLine", 3, "3pdr", "", 1, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Preussen", 1792, 1815, "Horse", "CrackLine", 2, "6pdr", "7pdr", 3, true}))

		// Austrian
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Austria", 1792, 1815, "Line", "CrackLine", 2, "6pdr", "7pdr", 3, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Austria", 1792, 1815, "Reserve", "CrackLine", 2, "12pdr", "10pdr", 3, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Austria", 1792, 1815, "Brigade", "CrackLine", 2, "6pdr", "", 4, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Austria", 1792, 1800, "Battalion Guns", "CrackLine", 2, "6pdr", "", 1, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Austria", 1792, 1800, "Grenz Bn Guns", "CrackLine", 2, "3pdr", "", 1, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Austria", 1792, 1815, "Kavallarie", "CrackLine", 2, "6pdr", "7pdr", 3, true}))

		//Minor Powers
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Sweden", 1792, 1815, "Line", "CrackLine", 2, "6pdr", "5.5\"", 3, false}))
		ut.Insert(UnitTypeMap("Artillery", Artillery{"Sweden", 1792, 1815, "Reserve", "CrackLine", 2, "12pdr", "", 3, false}))

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
	var utMap map[string]interface{}
	var allUnits []interface{}
	var theID string
	var theData map[string]interface{}

	for id := range allUnitTypeIds {
		unitTypes.Read(id, &utMap)
		//log.Printf("%+v", utMap)
		if utMap["Type"].(string) == "Infantry" {
			theID = utMap["@id"].(string)
			theData = utMap["Data"].(map[string]interface{})
			theData["@id"] = theID
			allUnits = append(allUnits, theData)
		}
	}
	log.Printf("All units %+v", allUnits)
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
				if utMap["Type"].(string) == "Infantry" {
					theID = utMap["@id"].(string)
					theData = utMap["Data"].(map[string]interface{})
					theData["@id"] = theID
					allUnits = append(allUnits, theData)
				}
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
