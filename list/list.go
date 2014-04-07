package list

import (
	"encoding/json"
	"github.com/steveoc64/tiedot/db"
)

type MessageFormat struct {
	Action string
	Entity string
	Data   interface{}
}

// Make the LIST and LOOKUP cache a global object
var ListCache map[string]interface{}
var LookupCache map[string]map[string]map[string]interface{}

// This is not as bad as it looks !!    LookupCache is a map of existing cached objects
// Each object is a collection of records of a particular table, keyed by the Selected Unique Key field
// Each row pointed to in that collection is formatted as a map[string], keyed on the field name
// Hence the awful looking 3D map structure below. Dont worry - it works well, and saves lots of casting when using the
// returned lookup object.

func Init() {
	ListCache = make(map[string]interface{})
	LookupCache = make(map[string]map[string]map[string]interface{})
}

func Clear(theEntity string) {
	delete(ListCache, theEntity)
	delete(LookupCache, theEntity)
}

// For a given entity, return a slice of bytes, being a JSON representation of that list
func Get(col *db.Col, theEntity string) (MessageFormat, bool) {
	var myData map[string]interface{}

	if records, ok := ListCache[theEntity]; ok {
		return records.(MessageFormat), true
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
	records := MessageFormat{"List", theEntity, results}
	ListCache[theEntity] = records
	return records, false
}

// Get a list of the given entity, mapped to a map[string]interface
func Lookup(col *db.Col, theEntity string, theKey string) (map[string]map[string]interface{}, bool) {
	var myData map[string]interface{}

	if lookupRecord, ok := LookupCache[theEntity]; ok {
		return lookupRecord, true
	}

	// Not cached, so Build a new result set using tiedot embedded query processor
	queryStr := `{"eq": "` + theEntity + `", "in": ["Type"]}`
	var query interface{}
	json.Unmarshal([]byte(queryStr), &query)
	queryResult := make(map[uint64]struct{}) // query result (document IDs) goes into map keys

	if err := db.EvalQuery(query, col, &queryResult); err != nil {
		panic(err)
	}

	results := make(map[string]map[string]interface{}, 0)

	for id := range queryResult {
		col.Read(id, &myData)

		theID := myData["@id"].(string)
		theData := myData["Data"].(map[string]interface{})
		theData["@id"] = theID

		key := theData[theKey]
		results[key.(string)] = theData
	}
	LookupCache[theEntity] = results
	return results, false
}
