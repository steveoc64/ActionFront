package oob

import (
	"encoding/json"
	"github.com/steveoc64/tiedot/db"
	"log"
)

// Create a DataMap envelope with type name and a JSON representation of the thing
func DataMap(typeName string, thing interface{}) map[string]interface{} {
	var jsonThing, err = json.Marshal(thing)
	if err != nil {
		panic(err)
	}
	var thingMap = map[string]interface{}{}
	json.Unmarshal(jsonThing, &thingMap)

	var retval = map[string]interface{}{}
	retval["Type"] = typeName
	retval["Data"] = thingMap
	return retval
}

// Create a Fresh Database of GameData from scratch

type Unit interface {
	getType() string
	nation() string
	name() string
	descr() string
}

type Infantry struct {
	N string
	T string
	D string
}

type Cavalry struct {
	N string
	T string
	D string
}

type Artillery struct {
	N string
	T string
	D string
}

func (i Infantry) getType() string {
	return "I"
}
func (c Cavalry) getType() string {
	return "C"
}
func (a Artillery) getType() string {
	return "A"
}
func (i Infantry) nation() string {
	return i.N
}
func (c Cavalry) nation() string {
	return c.N
}
func (a Artillery) nation() string {
	return a.N
}
func (i Infantry) name() string {
	return i.T
}
func (c Cavalry) name() string {
	return c.T
}
func (a Artillery) name() string {
	return a.T
}
func (i Infantry) descr() string {
	return i.D
}
func (c Cavalry) descr() string {
	return c.D
}
func (a Artillery) descr() string {
	return a.D
}

// Group name only has relevance within the ME that it belongs to
type Group struct {
	Units map[string]Unit
}

type ME struct {
	Nation string
	Year   uint16
	Name   string
	Groups map[string]Group
}

type Corps struct {
	Nation string
	Year   uint16
	Name   string
	MEs    []string
}

func CreateOOB(oobData *db.Col) {

	// Create a range of ME level units
	oobData.Insert(DataMap("ME", ME{"Austria", 1813, "Light Division Hardegg", map[string]Group{
		"Bde Regencourt": Group{map[string]Unit{
			"1st Grenz":    Infantry{"Austria", "Grenz", "Lt Blue"},
			"2nd Grenz":    Infantry{"Austria", "Grenz", "Lt Yellow"},
			"12th Dragoon": Cavalry{"Austria", "Dragoon", "Regt Graf Reisch, Lt Blue"},
			"4th Hussar":   Cavalry{"Austria", "Hussar", "Green, red breeches, lt blue shako, white lace"},
			"Fld Bty":      Artillery{"Austria", "Line", "6lb"},
		}},
	}}))

	oobData.Insert(DataMap("ME", ME{"Austria", 1813, "I Corps Artillery Reserve", map[string]Group{
		"Artillery Reserve": Group{map[string]Unit{
			"12lb Bty": Artillery{"Austria", "Reserve", "12lb"},
			"6lb Bty":  Artillery{"Austria", "Line", "6lb"},
		}},
	}}))

	oobData.Insert(DataMap("ME", ME{"Austria", 1813, "Line Division Wimpffen", map[string]Group{
		"Bde Giffing": Group{map[string]Unit{
			"1/54":    Infantry{"Austria", "Line", "Lt Green"},
			"2/54":    Infantry{"Austria", "Line", "Park Green"},
			"3/54":    Infantry{"Austria", "Line", "Park Green"},
			"1/25":    Infantry{"Austria", "Line", "Grey Green"},
			"2/25":    Infantry{"Austria", "Line", "Sea Green"},
			"3/25":    Infantry{"Austria", "Line", "Sea Green"},
			"Fld Bty": Artillery{"Austria", "Line", "6lb"},
		}},
		"Bde Chervenka": Group{map[string]Unit{
			"1/35":    Infantry{"Austria", "Line", "Amarinth Red"},
			"2/35":    Infantry{"Austria", "Line", "Amarinth Red"},
			"3/35":    Infantry{"Austria", "Line", "Amarinth Red"},
			"1/42":    Infantry{"Austria", "Line", "Orange"},
			"2/42":    Infantry{"Austria", "Line", "Orange"},
			"3/42":    Infantry{"Austria", "Line", "Orange"},
			"Fld Bty": Artillery{"Austria", "Line", "6lb"},
		}},
	}}))

	oobData.Insert(DataMap("ME", ME{"Austria", 1813, "Line Division Greth", map[string]Group{
		"Bde Mulheim": Group{map[string]Unit{
			"1/9":     Infantry{"Austria", "New Conscript", "Apple Green"},
			"2/9":     Infantry{"Austria", "New Conscript", "Apple Green"},
			"1/30":    Infantry{"Austria", "New Conscript", "Pike Grey"},
			"2/30":    Infantry{"Austria", "New Conscript", "Pike Grey"},
			"3/30":    Infantry{"Austria", "New Conscript", "6lb"},
			"Fld Bty": Artillery{"Austria", "Line", "6lb"},
		}},
		"Bde Quasdanovich": Group{map[string]Unit{
			"1/21":    Infantry{"Austria", "Line", "Sea Green"},
			"2/21":    Infantry{"Austria", "Line", "Sea Green"},
			"3/21":    Infantry{"Austria", "Line", "Sea Green"},
			"1/17":    Infantry{"Austria", "Line", "Lt Brown"},
			"2/17":    Infantry{"Austria", "Line", "Lt Brown"},
			"3/17":    Infantry{"Austria", "Line", "Lt Brown"},
			"Fld Bty": Artillery{"Austria", "Line", "6lb"},
		}},
	}}))

	oobData.Insert(DataMap("Corps", Corps{"Austria", 1813, "I Division Colloredo", []string{
		"Light Division Hardegg",
		"Line Division Wimpffen",
		"Line Division Greth",
		"I Corps Artillery Reserve",
	}}))

	// Now create some indexes
	log.Println("Creating Index on Type")
	if err := oobData.Index([]string{"Type", "Name"}); err != nil {
		panic(err)
	}
}
