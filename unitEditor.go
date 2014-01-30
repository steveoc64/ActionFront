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
	"time"
)

// Command line flags
var (
	port = flag.Int("port", 8080, "port to access the unitEditor")
)

// init the DB, and return a ref to the UnitTypes collection
func initUnitTypesDB() *db.Col {
	rand.Seed(time.Now().UTC().UnixNano())

	// Create and open database
	dir := "database/types"
	os.MkdirAll(dir, os.ModePerm)

	myDB, err := db.OpenDB(dir)
	if err != nil {
		panic(err)
	}

	if err := myDB.Create("UnitTypes", 1); err == nil {
		// This is a fresh DB so insert some default unit types
		ut := myDB.Use("UnitTypes")
		ut.Insert(map[string]interface{}{"name": "French Ligne", "rating": "Veteran", "size": "2L 1E"})
		ut.Insert(map[string]interface{}{"name": "French Legere", "rating": "Elite", "size": "3E"})
		ut.Insert(map[string]interface{}{"name": "French Conscript", "rating": "Conscript", "size": "3L"})
		ut.Insert(map[string]interface{}{"name": "Prussian Line", "rating": "CrackLine", "size": "4L"})
		ut.Insert(map[string]interface{}{"name": "Prussian Fusilier", "rating": "CrackLine", "size": "1L 3E"})
		ut.Insert(map[string]interface{}{"name": "Prussian Reserve", "rating": "Regular", "size": "4L"})
		ut.Insert(map[string]interface{}{"name": "Prussian Landwehr", "rating": "Landwehr", "size": "4L"})
	}
	myDB.Scrub("UnitTypes")
	return myDB.Use("UnitTypes")
}

// Pool of connections
var connections map[*websocket.Conn]bool

// Send message to all known connections
func sendAll(msg []byte) {
	for conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			delete(connections, conn)
			return
		}
	}
}

// Get all unit types into a slice of bytes
func getAllUnitTypes(col *db.Col) []byte {
	var query interface{}
	json.Unmarshal([]byte("[all]"), &query)

	queryResult := make(map[uint64]struct{}) // query result (document IDs) goes into map keys

	if err := db.EvalQuery(query, col, &queryResult); err != nil {
		panic(err)
	}
	log.Println(queryResult)

	// Query results are physical document IDs
	var readBack interface{}
	for id := range queryResult {
		log.Printf("Query returned document ID %d\n", id)
		col.Read(id, &readBack)
		log.Println(readBack)
	}
	return []byte("all")
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
	log.Println("New Websocket connection ", connections)
	defer conn.Close()

	getAllUnitTypes(unitTypes)

	// send the unitTypes to the new socket
	if err := conn.WriteMessage(websocket.TextMessage, []byte("all unit types")); err != nil {
		log.Println(err)
		delete(connections, conn)
		log.Println("Removed connection ", connections)
		return
	}

	// loop forever
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			delete(connections, conn)
			log.Println("Removed connection ", connections)
			return
		}
		log.Println(string(msg))
		sendAll(msg)
	}
}

// Main loop
func main() {

	flag.Parse()

	connections = make(map[*websocket.Conn]bool)

	// Classic defaults for webserver - serve up files from public dir
	m := martini.Classic()
	m.Map(initUnitTypesDB())
	m.Get("/socket", wsHandler)

	// Run the actual webserver
	addr := fmt.Sprintf(":%d", *port)
	log.Println("ActionFront Unit Editor starting on port ", addr)

	http.ListenAndServe(addr, m)
}
