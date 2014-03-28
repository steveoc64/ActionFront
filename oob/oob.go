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
}

type Infantry struct {
	Nation string
	Name   string
}

type Cavalry struct {
	Nation string
	Name   string
}

type Artillery struct {
	Nation string
	Name   string
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
	return i.Nation
}
func (c Cavalry) nation() string {
	return c.Nation
}
func (a Artillery) nation() string {
	return a.Nation
}
func (i Infantry) name() string {
	return i.Name
}
func (c Cavalry) name() string {
	return c.Name
}
func (a Artillery) name() string {
	return a.Name
}

type Group struct {
	Name  string
	Units map[string]Unit
}

type ME struct {
	Name   string
	Year   uint16
	Groups map[string]Group
}

func CreateOOB(oobData *db.Col) {

	a := Group{"Bde Ragencourt", map[string]Unit{
		"1st Grenz":    Infantry{"Austria", "Grenz"},
		"2nd Grenz":    Infantry{"Austria", "Grenz"},
		"12th Dragoon": Cavalry{"Austria", "Dragoon"},
		"4th Hussar":   Cavalry{"Austria", "Hussar"},
		"Fld Bty":      Artillery{"Austria", "Line"},
	}}

	log.Println("a = ", a)
	/*
		// Create some DrillBooks
		oobData.Insert(DataMap("ME", ME{"Light Division", 1813, map[string]Group{
			"Bde Ragencourt": map[string]Unit{
				"1st Grenz":    Infantry{"Austria", "Grenz"},
				"2nd Grenz":    Infantry{"Austria", "Grenz"},
				"12th Dragoon": Cavalry{"Austria", "Dragoon"},
				"4th Hussar":   Cavalry{"Austria", "Hussar"},
				"Fld Bty":      Artillery{"Austria", "Line"},
			},
		}}))
	*/
	// Now create some indexes
	log.Println("Creating Index on Type")
	if err := oobData.Index([]string{"Type", "Name"}); err != nil {
		panic(err)
	}
}
