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

// Make the LIST cache a global object
var ListCache map[string]interface{}

func Init() {
	ListCache = make(map[string]interface{})
}

func Clear(theEntity string) {
	delete(ListCache, theEntity)
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
