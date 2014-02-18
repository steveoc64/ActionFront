package gamedatadb

import (
	"encoding/json"
	"github.com/HouzuoGuo/tiedot/db"
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
	Shock     bool
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

// Command and Control tables

type InitTable struct {
	Factor string
	Value  int8
}

type CorpsOrder struct {
	Order       string
	MEOrders    []string
	Stipulation string
}

type MEOrder struct {
	Order           string
	Purpose         string
	Notes           string
	IfNonEngaged    bool
	IfEngaged       bool
	CavalryOnly     bool
	DefendIfEngaged bool
	ShakenIfEngaged bool
}

type OrderArrival struct {
	Grids  uint16
	Delay  uint16
	DGrids uint16
}

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

func CreateGameData(gameData *db.Col) {

	// Create some DrillBooks
	gameData.Insert(DataMap("Drill", DrillBook{"Conscript", map[string]Drill{
		"Line":           Drill{5, 2, 0, 0},
		"MarchColumn":    Drill{8, 1, 0, 0},
		"AttackColumn":   Drill{7, 1, 1, 0},
		"ClosedColumn":   Drill{6, 1, 0, 0},
		"ScreenedColumn": Drill{5, 1, 1, 0}}}))

	gameData.Insert(DataMap("Drill", DrillBook{"Militia", map[string]Drill{
		"MarchColumn":  Drill{8, 1, 0, 0},
		"ClosedColumn": Drill{5, 1, 0, 0},
		"Screen":       Drill{4, 1, 2, 2}}}))

	gameData.Insert(DataMap("Drill", DrillBook{"Mob", map[string]Drill{
		"MarchColumn":   Drill{6, 1, 0, 0},
		"DisorderedMob": Drill{4, 2, 2, 0}}}))

	gameData.Insert(DataMap("Drill", DrillBook{"French", map[string]Drill{
		"Line":         Drill{7, 3, 0, 0},
		"MarchColumn":  Drill{10, 1, 0, 0},
		"AttackColumn": Drill{9, 2, 1, 1},
		"ClosedColumn": Drill{8, 1, 0, 1},
		"Square":       Drill{6, 1, 0, 1}}}))

	gameData.Insert(DataMap("Drill", DrillBook{"Light Infantry", map[string]Drill{
		"Skirmish":     Drill{7, 8, 4, 4},
		"Screen":       Drill{8, 6, 3, 3},
		"Line":         Drill{7, 3, 2, 1},
		"MarchColumn":  Drill{10, 1, 0, 1},
		"AttackColumn": Drill{9, 1, 1, 1},
		"ClosedColumn": Drill{8, 1, 0, 1},
		"Square":       Drill{5, 1, 0, 1}}}))

	gameData.Insert(DataMap("Drill", DrillBook{"Prussian", map[string]Drill{
		"Line":         Drill{7, 4, 0, 1},
		"Oblique":      Drill{6, 4, 0, 1},
		"ScreenedLine": Drill{7, 3, 1, 1},
		"MarchColumn":  Drill{9, 1, 0, 0},
		"AttackColumn": Drill{8, 2, 1, 1},
		"ClosedColumn": Drill{7, 2, 0, 0},
		"Square":       Drill{5, 1, 0, 0}}}))

	gameData.Insert(DataMap("Drill", DrillBook{"British", map[string]Drill{
		"Line":         Drill{7, 4, 0, 1},
		"ScreenedLine": Drill{7, 3, 1, 1},
		"MarchColumn":  Drill{10, 1, 0, 0},
		"AttackColumn": Drill{8, 1, 1, 1},
		"ClosedColumn": Drill{7, 1, 0, 0},
		"Square":       Drill{5, 1, 0, 0}}}))

	gameData.Insert(DataMap("Drill", DrillBook{"OldSchool", map[string]Drill{
		"Line":         Drill{6, 4, 0, 0},
		"Oblique":      Drill{5, 4, 0, 0},
		"ScreenedLine": Drill{6, 4, 0, 1},
		"MarchColumn":  Drill{9, 1, 0, 0},
		"AttackColumn": Drill{7, 1, 0, 0},
		"Square":       Drill{4, 1, 0, 0}}}))

	gameData.Insert(DataMap("Drill", DrillBook{"Russian", map[string]Drill{
		"Line":         Drill{7, 2, 0, 0},
		"MarchColumn":  Drill{9, 1, 0, 0},
		"AttackColumn": Drill{8, 1, 0, 0},
		"Square":       Drill{5, 1, 0, 0}}}))

	gameData.Insert(DataMap("Drill", DrillBook{"Austrian", map[string]Drill{
		"Line":         Drill{6, 6, 0, 1},
		"ScreenedLine": Drill{6, 5, 1, 1},
		"MarchColumn":  Drill{9, 1, 0, 0},
		"AttackColumn": Drill{8, 2, 1, 1},
		"ClosedColumn": Drill{7, 2, 0, 0},
		"Mass":         Drill{6, 3, 2, 0},
		"Square":       Drill{4, 1, 0, 0}}}))

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Add some Infantry

	// French Line
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1800, 1812, "30/32/34 Ligne", "Elite", "French", "4L 2E", 0, 1, "Musket", "Excellent", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1800, 1812, "57/84 Ligne", "Grenadier", "French", "4L 2E", 0, 1, "Musket", "Excellent", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1796, 1801, "18 Ligne", "Elite", "French", "3L 1E", 0, 1, "Musket", "Excellent", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1813, 1813, "135-156 Ligne", "Veteran", "French", "3L", 0, 2, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1814, 1814, "135-156 Ligne", "Veteran", "French", "2L", 0, 2, "Musket", "Poor", "Average", false}))

	gameData.Insert(DataMap("Infantry", Infantry{"France", 1791, 1792, "Veteran Ligne", "Veteran", "French", "4L 1S", 0, 2, "Musket", "Good", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1791, 1792, "Conscript Ligne", "Conscript", "French", "4L 1S", -1, 2, "Musket", "Average", "Average", false}))

	gameData.Insert(DataMap("Infantry", Infantry{"France", 1793, 1795, "Crack Ligne", "CrackLine", "French", "4L 1S", 0, 2, "Musket", "Good", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1793, 1795, "Veteran Ligne", "Veteran", "French", "4L 1S", -1, 2, "Musket", "Good", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1793, 1795, "Conscript Ligne", "Conscript", "French", "2L", -1, 2, "Musket", "Average", "Average", false}))

	gameData.Insert(DataMap("Infantry", Infantry{"France", 1796, 1804, "Elite Ligne", "Elite", "French", "4L 1S", 0, 2, "Musket", "Excellent", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1796, 1804, "Crack Ligne", "CrackLine", "French", "4L 1S", 0, 2, "Musket", "Good", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1796, 1804, "Veteran Ligne", "Veteran", "French", "4L 1S", 0, 2, "Musket", "Good", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1796, 1804, "Conscript Ligne", "Conscript", "Conscript", "3L", 0, 2, "Musket", "Average", "Average", false}))

	gameData.Insert(DataMap("Infantry", Infantry{"France", 1805, 1807, "Elite Ligne", "Elite", "French", "4L 1S", 0, 2, "Musket", "Excellent", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1805, 1807, "Crack Ligne", "CrackLine", "French", "4L 1S", 0, 2, "Musket", "Excellent", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1805, 1807, "Veteran Ligne", "Veteran", "French", "4L 1S", 0, 2, "Musket", "Average", "Good", false}))

	gameData.Insert(DataMap("Infantry", Infantry{"France", 1808, 1812, "Elite Ligne", "Elite", "French", "3L 1E", 0, 2, "Musket", "Excellent", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1808, 1812, "Crack Ligne", "CrackLine", "French", "3L 1E", 0, 2, "Musket", "Excellent", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1808, 1812, "Veteran Ligne", "Veteran", "French", "3L 1E", 0, 2, "Musket", "Good", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1808, 1812, "Regular Ligne", "Regular", "French", "3L 1E", 0, 2, "Musket", "Average", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1808, 1812, "Conscript Ligne", "Conscript", "Conscript", "4L", 0, 2, "Musket", "Poor", "Good", false}))

	gameData.Insert(DataMap("Infantry", Infantry{"France", 1813, 1814, "Veteran Ligne", "Veteran", "French", "2L 1E", 0, 2, "Musket", "Average", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1813, 1814, "Conscript Ligne", "Conscript", "Conscript", "3L", 0, 2, "Musket", "Poor", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1813, 1814, "Provisional Ligne", "Veteran", "French", "2L", 0, 2, "Musket", "Poor", "Poor", false}))

	gameData.Insert(DataMap("Infantry", Infantry{"France", 1815, 1815, "Elites", "Elite", "French", "2L 1E", 0, 2, "Musket", "Excellent", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1815, 1815, "Crack Ligne", "CrackLine", "French", "2L 1E", 0, 2, "Musket", "Excellent", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1815, 1815, "Veteran Ligne", "Veteran", "French", "2L 1E", 0, 2, "Musket", "Good", "Good", false}))

	// French Line in the Peninsula
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1807, 1807, "Veteran Ligne (Peninsula)", "Veteran", "French", "3L 1E", 0, 2, "Musket", "Average", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1807, 1807, "Ligne (Peninsula)", "Regular", "French", "3L 1E", 0, 2, "Musket", "Average", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1808, 1812, "Elite Ligne (Peninsula)", "Elite", "French", "3L 1E", 0, 2, "Musket", "Excellent", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1808, 1812, "Crack Ligne (Peninsula)", "CrackLine", "French", "3L 1E", 0, 2, "Musket", "Excellent", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1808, 1812, "Veteran Ligne (Peninsula)", "Veteran", "French", "3L 1E", 0, 2, "Musket", "Good", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1808, 1812, "Conscript Ligne (Peninsula)", "Conscript", "Conscript", "3L 1E", 0, 2, "Musket", "Poor", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1813, 1814, "Crack Ligne (Peninsula)", "CrackLine", "French", "3L 1E", 0, 2, "Musket", "Excellent", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1813, 1814, "Veteran Ligne (Peninsula)", "Veteran", "French", "3L 1E", 0, 2, "Musket", "Good", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1813, 1814, "Ligne (Peninsula)", "Regular", "French", "3L 1E", 0, 2, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1813, 1814, "Conscript Ligne (Peninsula)", "Conscript", "Conscript", "3L 1E", 0, 2, "Musket", "Poor", "Good", false}))

	// French Light Infantry
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1796, 1815, "9/10eme Legere", "Grenadier", "Light Infantry", "5E", 0, 1, "Musket", "Superior", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1811, 1815, "1/2Bn 11eme Legere", "Grenadier", "Light Infantry", "5E", 0, 1, "Musket", "Superior", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1811, 1815, "3/4Bn 11eme Legere", "Landwehr", "Light Infantry", "5E", 0, 2, "Musket", "Poor", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1805, 1811, "Tirailleurs du Po/Corses", "Grenadier", "Light Infantry", "5E", 0, 1, "Musket", "Superior", "Excellent", true}))

	gameData.Insert(DataMap("Infantry", Infantry{"France", 1791, 1792, "Crack Legere", "CrackLine", "Light Infantry", "3E", 0, 2, "Musket", "Good", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1791, 1792, "Veteran Legere", "Veteran", "Light Infantry", "3E", 0, 2, "Musket", "Good", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1793, 1795, "Elite Legere", "Elite", "Light Infantry", "5E", 0, 2, "Musket", "Excellent", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1793, 1795, "Crack Legere", "CrackLine", "Light Infantry", "5E", 0, 2, "Musket", "Excellent", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1793, 1795, "Veteran Legere", "Veteran", "Light Infantry", "5E", 0, 2, "Musket", "Good", "Good", false}))

	gameData.Insert(DataMap("Infantry", Infantry{"France", 1796, 1804, "Elite Legere", "Elite", "Light Infantry", "5E", 0, 2, "Musket", "Excellent", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1796, 1804, "Veteran Legere", "Veteran", "Light Infantry", "5E", 0, 2, "Musket", "Excellent", "Excellent", false}))

	gameData.Insert(DataMap("Infantry", Infantry{"France", 1805, 1807, "Elite Legere", "Elite", "French", "5E", 0, 2, "Musket", "Excellent", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1805, 1807, "Crack Legere", "CrackLine", "French", "5E", 0, 2, "Musket", "Excellent", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1807, 1807, "Veteran Legere (Peninsula)", "Veteran", "French", "5E", 0, 2, "Musket", "Good", "Good", false}))

	gameData.Insert(DataMap("Infantry", Infantry{"France", 1808, 1812, "Elite Legere", "Elite", "French", "4E", 0, 2, "Musket", "Excellent", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1808, 1812, "Crack Legere", "CrackLine", "French", "4E", 0, 2, "Musket", "Excellent", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1808, 1812, "Veteran Legere", "Elite", "French", "3E", 0, 2, "Musket", "Good", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1808, 1812, "Regular Legere", "CrackLine", "French", "3E", 0, 2, "Musket", "Average", "Average", false}))

	gameData.Insert(DataMap("Infantry", Infantry{"France", 1813, 1814, "Crack Legere", "CrackLine", "French", "3E", 0, 2, "Musket", "Excellent", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1813, 1814, "Veteran Legere", "Veteran", "French", "3E", 0, 2, "Musket", "Good", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1813, 1814, "Conscript Legere", "Conscript", "French", "3E", 0, 2, "Musket", "Poor", "Poor", false}))

	gameData.Insert(DataMap("Infantry", Infantry{"France", 1815, 1815, "Elite Legere", "Elite", "French", "3E", 0, 2, "Musket", "Excellent", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1815, 1815, "Veteran Legere", "Veteran", "French", "3E", 0, 2, "Musket", "Excellent", "Good", false}))

	// Other French Infantry
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1813, 1814, "Line Marines", "Elite", "French", "3L", -3, 0, "Musket", "Poor", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1810, 1812, "Croatian Regiment", "CrackLine", "French", "2L 1E", 0, 0, "Musket", "Good", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1805, 1815, "Swiss Regiment", "Elite", "French", "2L 1E", 0, 1, "Musket", "Good", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1805, 1812, "Irish Regiment", "Veteran", "French", "2L 1E", 0, 1, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1805, 1815, "Etranger", "Landwehr", "Conscript", "3L", 0, 0, "Musket", "Poor", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1792, 1815, "Joseph Napoleon Grenadiers", "Veteran", "French", "2L 1E", 0, 0, "Musket", "Poor", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1792, 1815, "Vistula Legion", "Grenadier", "French", "2L 1E", 0, 1, "Musket", "Good", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1805, 1806, "Foot Dragoon", "Veteran", "French", "2L 1E", -2, 0, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1792, 1815, "Sappeurs", "Grenadier", "French", "1E", -2, 0, "Musket", "Poor", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1807, 1813, "Neuchatel Regiment", "Veteran", "French", "2L 1E", 0, 0, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1796, 1805, "Converged Elites", "Elite", "French", "2L 1E", 0, 0, "Musket", "Excellent", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1806, 1814, "Converged Elites", "CrackLine", "French", "2L 1E", 0, 0, "Musket", "Excellent", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France", 1809, 1809, "Demi Brigades d'Elite", "Conscript", "Conscript", "3E", 0, 0, "Musket", "Average", "Average", false}))

	// French Imperial Guard
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1792, 1815, "1/2 Grenadiers", "OldGuard", "French", "3E", 0, 0, "Musket", "Superior", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1792, 1815, "1/2 Chasseurs", "OldGuard", "French", "3E", 0, 0, "Musket", "Superior", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1792, 1815, "Marine", "OldGuard", "French", "3E", 0, 0, "Musket", "Excellent", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1792, 1815, "Genies", "OldGuard", "French", "1E", 0, 0, "Musket", "Excellent", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1806, 1814, "Fusilier Grenadiers", "Guard", "French", "3L 1S", 0, 0, "Musket", "Excellent", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1806, 1814, "Fusilier Chasseurs", "Guard", "French", "3L 1S", 0, 0, "Musket", "Excellent", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1811, 1812, "3rd Grenadiers", "Grenadier", "French", "3L 1S", 0, 0, "Musket", "Good", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1815, 1815, "3/4 Guard Regiment", "Guard", "French", "3L 1S", 0, 0, "Musket", "Excellent", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1811, 1814, "Flanker Grenadiers", "Grenadier", "French", "3L 1S", 0, 0, "Musket", "Excellent", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1811, 1814, "Flanker Chasseurs", "Grenadier", "French", "3E", 0, 0, "Musket", "Excellent", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1809, 1812, "Young Guard", "Grenadier", "French", "2L 1E", 0, 0, "Musket", "Excellent", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1815, 1815, "Young Guard", "Grenadier", "French", "2L 1E", 0, 0, "Musket", "Excellent", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1813, 1813, "Young Guard (Elite)", "Grenadier", "French", "2L 1E", 0, 0, "Musket", "Excellent", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1813, 1813, "Young Guard", "Elite", "French", "1L 1E", 0, 0, "Musket", "Good", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1813, 1813, "Young Guard (Recruit)", "CrackLine", "French", "1L 1E", -1, 0, "Musket", "Good", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1813, 1813, "Young Guard (Elite)", "Elite", "French", "1L 1E", 0, 0, "Musket", "Good", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1813, 1813, "Young Guard", "CrackLine", "French", "2L 1E", -1, 0, "Musket", "Good", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1813, 1813, "Young Guard Sappeurs", "CrackLine", "French", "3E", -1, 0, "Musket", "Good", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1792, 1815, "Velites d'Florence", "Grenadier", "French", "3E", 0, 0, "Musket", "Excellent", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Guard", 1792, 1815, "Velites d'Turin", "Grenadier", "French", "3E", 0, 0, "Musket", "Excellent", "Good", true}))

	// French Royalist
	gameData.Insert(DataMap("Infantry", Infantry{"France Royalist", 1792, 1804, "Emigree", "Regular", "OldSchool", "2L 2S", 0, 0, "Musket", "Good", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"France Royalist", 1792, 1804, "Vendeen", "Landwehr", "Militia", "4E", 1, 0, "Musket", "Excellent", "Excellent", false}))

	// British
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1792, 1815, "Foot Guards", "Guards", "British", "4L 1E", 0, 0, "Musket", "Excellent", "Good", true}))

	gameData.Insert(DataMap("Infantry", Infantry{"British", 1792, 1815, "Highlanders", "Elite", "British", "4L 1E", 0, 0, "Musket", "Excellent", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1792, 1815, "Fusiliers", "Elite", "British", "4L 1E", 0, 0, "Musket", "Good", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1792, 1815, "Marines", "Elite", "British", "4L 1E", 0, 0, "Musket", "Good", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1792, 1815, "Light Infantry", "Elite", "British", "3E", 0, 1, "Musket", "Superior", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1792, 1815, "1st Line Bn", "CrackLine", "British", "3L 1E", 0, 0, "Musket", "Good", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1792, 1815, "2nd Line Bn", "Veteran", "British", "3L 1E", 0, 0, "Musket", "Good", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1812, 1815, "Canadian Militia", "Landwehr", "British", "1L 1E", 0, 0, "Musket", "Poor", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1792, 1815, "Rifles", "Elite", "British", "3E", 0, 1, "Rifle", "Superior", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1792, 1815, "95th Rifles", "Grenadier", "British", "5E", 0, 0, "Rifle", "Superior", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1792, 1815, "43/52 Line", "Grenadier", "British", "4L 1E", 0, 0, "Musket", "Excellent", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1809, 1811, "45/88 Line", "Elite", "British", "4L 1E", 0, 0, "Musket", "Excellent", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1812, 1815, "45/88 Line", "Grenadier", "British", "4L 1E", 0, 0, "Musket", "Excellent", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1809, 1815, "40/48 Line", "Elite", "British", "4L 1E", 0, 0, "Musket", "Excellent", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1812, 1815, "74/94 Line", "Elite", "British", "4L 1E", 0, 0, "Musket", "Excellent", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1815, 1815, "4/28/32/40 Line", "Elite", "British", "4L 1E", 0, 0, "Musket", "Excellent", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1815, 1815, "51/91/1 Line", "Elite", "British", "4L 1E", 0, 0, "Musket", "Excellent", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1792, 1815, "American Indian Allies", "Militia", "Rabble", "6S", -1, 0, "Musket", "Average", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1792, 1815, "KGL Light", "Elite", "British", "2L 2E", 0, 0, "Rifle", "Excellent", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1800, 1808, "KGL Line", "CrackLine", "British", "4L 1E", 0, 2, "Musket", "Good", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1809, 1811, "1/2/5 KGL Line", "Elite", "British", "4L 1E", 0, 2, "Musket", "Good", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1809, 1811, "KGL Line", "CrackLine", "British", "4L 1E", 0, 2, "Musket", "Good", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"British", 1812, 1815, "KGL Line", "Elite", "British", "4L 1E", 0, 2, "Musket", "Good", "Good", true}))

	// Russian
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1789, 1810, "Musketeer", "Veteran", "Russian", "3L", 0, 0, "Poor Musket", "", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1789, 1810, "Line Grenadier", "CrackLine", "Russian", "3L", 0, 0, "Poor Musket", "", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1811, 1812, "Line", "Veteran", "Russian", "3L", 0, 0, "Poor Musket", "", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1815, 1815, "Line", "Veteran", "Russian", "3L", 0, 0, "Poor Musket", "", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1813, 1814, "Veteran Line", "Veteran", "Russian", "2L", 0, 0, "Poor Musket", "", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1813, 1814, "Line", "Regular", "Russian", "2L", 0, 0, "Poor Musket", "", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1813, 1814, "Conscript Line", "Conscript", "Russian", "3L", 0, 0, "Poor Musket", "", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1789, 1810, "Veteran Jager", "Veteran", "Russian", "3O", 0, 0, "Poor Musket", "Average", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1789, 1810, "Jager", "Veteran", "Russian", "3O", 0, 0, "Poor Musket", "Average", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1789, 1810, "Conscript Jager", "Conscript", "Russian", "2E", 0, 0, "Poor Musket", "Poor", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1811, 1812, "Jager", "CrackLine", "Russian", "3O", 0, 0, "Poor Musket", "Average", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1815, 1815, "Jager", "CrackLine", "Russian", "3O", 0, 0, "Poor Musket", "Average", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1813, 1814, "Crack Jager", "CrackLine", "Russian", "3O", 0, 0, "Poor Musket", "Average", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1813, 1814, "Veteran Jager", "Veteran", "Russian", "3O", 0, 0, "Poor Musket", "Average", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1813, 1814, "Jager", "Regular", "Russian", "3O", 0, 0, "Poor Musket", "Average", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1813, 1814, "Conscript Jager", "Conscript", "Russian", "3O", 0, 0, "Poor Musket", "Poor", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1792, 1815, "Opolchenia", "Militia", "Militia", "3L", 0, 0, "Poor Musket", "", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1792, 1810, "Fusilier", "CrackLine", "Russian", "2L", 0, 0, "Poor Musket", "", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1792, 1810, "Grenadier", "Grenadier", "Russian", "2L", 0, 0, "Poor Musket", "", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1811, 1812, "Grenadier", "Grenadier", "Russian", "2L", 0, 0, "Poor Musket", "", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1815, 1815, "Grenadier", "Grenadier", "Russian", "2L", 0, 0, "Poor Musket", "", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1813, 1814, "Grenadier", "Elite", "Russian", "2L", 0, 0, "Poor Musket", "", "Excellent", true}))

	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1792, 1815, "Princess Catherine Regt", "CrackLine", "Russian", "3L", 0, 0, "Poor Musket", "Average", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1792, 1815, "Kexholm Regt", "Elite", "Russian", "3L", 0, 0, "Poor Musket", "Average", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1792, 1815, "Pennovsky Regt", "Elite", "Russian", "3L", 0, 0, "Poor Musket", "Average", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1805, 1812, "Kiev Regt", "Elite", "Russian", "3L", 0, 0, "Poor Musket", "Average", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1805, 1814, "Schusselberg Regt", "Elite", "Russian", "3L", 0, 0, "Poor Musket", "Average", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1807, 1812, "Triosk Regt", "Elite", "Russian", "3L", 0, 0, "Poor Musket", "Average", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1805, 1812, "Fangoria Regt", "Elite", "Russian", "3L", 0, 0, "Poor Musket", "Average", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1805, 1812, "Grouzin Regt", "Elite", "Russian", "3L", 0, 0, "Poor Musket", "Average", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1792, 1815, "Tomsk Regt", "Conscript", "Russian", "3L", 0, 0, "Poor Musket", "", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1792, 1815, "Apcherin Regt", "Conscript", "Russian", "3L", 0, 0, "Poor Musket", "", "Excellent", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia", 1811, 1815, "1/6/40 Jager", "Elite", "Russian", "3E", 0, 0, "Poor Musket", "Good", "Excellent", true}))

	// Russian Guard Infantry
	gameData.Insert(DataMap("Infantry", Infantry{"Russia Guard", 1792, 1812, "Guard Grenadier", "Guard", "Russian", "3L 1S", 0, 0, "Poor Musket", "Average", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia Guard", 1815, 1815, "Guard Grenadier", "Guard", "Russian", "3L 1S", 0, 0, "Poor Musket", "Average", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia Guard", 1813, 1814, "Guard Grenadier", "Grenadier", "Russian", "2L 1S", 0, 0, "Poor Musket", "Average", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia Guard", 1792, 1812, "Guard Jager", "Guard", "Russian", "3E", 0, 0, "Poor Musket", "Good", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia Guard", 1815, 1815, "Guard Jager", "Guard", "Russian", "3E", 0, 0, "Poor Musket", "Good", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Russia Guard", 1813, 1814, "Guard Jager", "Grenadier", "Russian", "2E", 0, 0, "Poor Musket", "Good", "Excellent", true}))

	// Prussia
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen", 1792, 1810, "Musketeer", "CrackLine", "OldSchool", "4L 1S", 0, 0, "Poor Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen", 1792, 1810, "Fusilier", "CrackLine", "OldSchool", "2L 3O", 0, 0, "Poor Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen", 1792, 1810, "Grenadier", "Grenadier", "OldSchool", "4L 1S", 0, 0, "Poor Musket", "Average", "Average", true}))

	gameData.Insert(DataMap("Infantry", Infantry{"Preussen", 1811, 1815, "1-12 Musketeer", "CrackLine", "Prussian", "4L 1S", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen", 1811, 1815, "1-12 Fusilier", "CrackLine", "Prussian", "2L 2E", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen", 1811, 1815, "Musketeer", "Regular", "Prussian", "4L", 0, 0, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen", 1811, 1815, "Line Grenadier", "Elite", "Prussian", "4L 1S", 0, 0, "Musket", "Good", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen", 1811, 1815, "Jager", "Elite", "Light Infantry", "2S", 0, 0, "Rifle", "Excellent", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen", 1815, 1815, "Westphalian Landwehr", "Militia", "Militia", "4L", 0, 0, "Musket", "", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen", 1813, 1813, "Landwehr", "Landwehr", "Conscript", "4L", 0, 0, "Musket", "", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen", 1815, 1815, "Landwehr", "Landwehr", "Conscript", "4L", 0, 0, "Musket", "", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen", 1814, 1814, "Veteran Landwehr", "Veteran", "Conscript", "4L", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen", 1814, 1814, "Landwehr", "Conscript", "Conscript", "4L", 0, 0, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen", 1815, 1815, "25/26/28/29 Line Regiment", "Veteran", "Prussian", "4L 1S", 0, 0, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen", 1815, 1815, "32 Regiment", "Landwehr", "Prussian", "4L", 0, 0, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen", 1811, 1812, "Russo German Legion", "Veteran", "Prussian", "4L 1S", 0, 0, "Poor Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen", 1813, 1814, "Russo German Legion", "CrackLine", "Prussian", "4L 1S", 0, 0, "Poor Musket", "Good", "Average", false}))

	// Prussian Guard
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen Guard", 179, 1810, "Leibguard", "Guard", "OldSchool", "4L", 0, 0, "Poor Musket", "", "Average", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen Guard", 1792, 1810, "Grenadier", "Grenadier", "OldSchool", "4L", 0, 0, "Poor Musket", "", "Average", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen Guard", 1811, 1815, "Grenadier", "Grenadier", "Prussian", "4L 1S", 0, 0, "Musket", "Good", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen Guard", 1811, 1815, "Jager", "Grenadier", "Light Infantry", "4E", 0, 0, "Rifle", "Excellent", "Excellent", true}))

	// Austrian
	gameData.Insert(DataMap("Infantry", Infantry{"Austria", 1792, 1808, "Line", "Regular", "OldSchool", "6L", 0, 0, "Musket", "", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Austria", 1809, 1815, "Line", "Regular", "Austrian", "6L 1S", 0, 0, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Austria", 1792, 1800, "Grenadier", "Grenadier", "OldSchool", "4L", 0, 0, "Musket", "", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Austria", 1801, 1805, "Elite Grenadier", "Grenadier", "OldSchool", "4L", 0, 0, "Musket", "", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Austria", 1801, 1805, "Grenadier", "CrackLine", "OldSchool", "4L", 0, 0, "Musket", "", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Austria", 1809, 1812, "Grenadier", "Grenadier", "Austrian", "4L 1S", 0, 0, "Musket", "Average", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Austria", 1813, 1814, "Grenadier", "Elite", "Austrian", "4L 2S", 0, 0, "Musket", "Average", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Austria", 1815, 1815, "Grenadier", "Grenadier", "Austrian", "4L 2S", 0, 0, "Musket", "Average", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Austria", 1792, 1808, "Jager", "CrackLine", "Light Infantry", "4E", 0, 0, "Rifle", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Austria", 1809, 1815, "Jager", "Grenadier", "Light Infantry", "4E", 0, 0, "Rifle", "Good", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Austria", 1792, 1815, "Freikorps", "Regular", "Light Infantry", "6O", 0, 0, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Austria", 1792, 1815, "Grenz", "Veteran", "Light Infantry", "4O", 0, 0, "Rifle", "Average", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Austria", 1792, 1815, "Insurrection", "Militia", "Militia", "4L", 0, 0, "Musket", "", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Austria", 1792, 1815, "#4 Line", "Grenadier", "Austrian", "6L", -1, 0, "Musket", "", "Average", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Austria", 1792, 1815, "#14 Line", "Elite", "Austrian", "6L", 0, 0, "Musket", "", "Average", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Austria", 1809, 1815, "1/3/11/19/46/59 Line", "CrackLine", "Austrian", "6L", 0, 0, "Musket", "", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Austria", 1811, 1815, "9/20/24/30/44/58/63 Line", "Conscript", "Conscript", "6L", 0, 0, "Rifle", "Good", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Austria", 1811, 1815, "#6 Grenz", "CrackLine", "Light Infantry", "4O", 0, 0, "Rifle", "Good", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Austria", 1809, 1815, "Charles Legion", "Veteran", "Light Infantry", "6O", 0, 0, "Rifle", "Average", "Average", false}))

	// Kingdom of Spain
	gameData.Insert(DataMap("Infantry", Infantry{"Spain", 1792, 1800, "Line", "Regular", "OldSchool", "4L", 0, 0, "Musket", "", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Spain", 1801, 1811, "Line", "Conscript", "OldSchool", "4L 1S", 0, 0, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Spain", 1812, 1814, "Line", "Regular", "OldSchool", "4L 2S", 0, 0, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Spain", 1792, 1814, "Grenadier", "CrackLine", "OldSchool", "2L 1E", 0, 0, "Musket", "Poor", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Spain", 1792, 1814, "Light Infantry", "Veteran", "Light Infantry", "3E", 0, 0, "Musket", "Average", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Spain", 1792, 1814, "Marines", "Landwehr", "Militia", "2L", 0, 0, "Musket", "", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Spain", 1792, 1814, "Levy", "Militia", "Militia", "2L", 0, 0, "Musket", "", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Spain", 1792, 1814, "Militia", "Rabble", "Militia", "2L", 0, 0, "Musket", "", "Poor", false}))

	// Kingdom of Spain - Guard
	gameData.Insert(DataMap("Infantry", Infantry{"Spain Guard", 1792, 1814, "Guard", "Elite", "OldSchool", "4L", 0, 0, "Musket", "", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Spain Guard", 1808, 1814, "Guard Joseph Napoleon", "CrackLine", "French", "3L 1E", 0, 0, "Musket", "Average", "Good", true}))

	// Portugal
	gameData.Insert(DataMap("Infantry", Infantry{"Portugal", 1800, 1808, "Line", "Militia", "Militia", "4L", 0, 0, "Musket", "", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Portugal", 1809, 1811, "Line", "Regular", "Militia", "4L", 0, 0, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Portugal", 1812, 1814, "Crack Line", "CrackLine", "British", "3L 1E", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Portugal", 1812, 1814, "Line", "Veteran", "British", "3L 1E", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Portugal", 1792, 1815, "Militia", "Militia", "Militia", "3L", 0, 0, "Musket", "", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Portugal", 1800, 1808, "Cacadore", "Militia", "Militia", "4L", 0, 0, "Musket", "", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Portugal", 1809, 1811, "Cacadore", "Veteran", "Light Infantry", "4E", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Portugal", 1812, 1814, "Cacadore", "CrackLine", "Light Infantry", "4E", 0, 0, "Musket", "Good", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Portugal", 1810, 1811, "1/3 Cacadore", "CrackLine", "Light Infantry", "4E", 0, 0, "Musket", "Excellent", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Portugal", 1812, 1814, "1/3 Cacadore", "Elite", "Light Infantry", "4E", 0, 0, "Musket", "Excellent", "Average", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Portugal", 1808, 1811, "Lusitanian Legion", "Regular", "British", "4L 1E", 0, 0, "Musket", "Average", "Average", false}))

	// Ottoman Empire
	gameData.Insert(DataMap("Infantry", Infantry{"Ottoman", 1790, 1815, "Veteran Janissaries", "Veteran", "Conscript", "5L", 0, 0, "Poor Musket", "Poor", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Ottoman", 1790, 1815, "Regular Janissaries", "Regular", "Conscript", "5L", 0, 0, "Poor Musket", "Poor", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Ottoman", 1790, 1815, "Janissaries", "Conscript", "Militia", "5L", 0, 0, "Poor Musket", "", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Ottoman", 1790, 1815, "Untrained Janissaries", "Landwehr", "Militia", "5L", 0, 0, "Poor Musket", "", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Ottoman", 1790, 1815, "Sekhans", "Landwehr", "Militia", "4L", 0, 0, "Poor Musket", "", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Ottoman", 1790, 1815, "Rayas", "Veteran", "Light Infantry", "2S", 0, 0, "Poor Musket", "Excellent", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Ottoman", 1790, 1815, "Martolo", "Regular", "Light Infantry", "3E", 0, 1, "Poor Musket", "Good", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Ottoman", 1790, 1815, "Derbant", "Militia", "Militia", "4E", 0, 0, "Poor Musket", "Poor", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Ottoman", 1790, 1815, "Fellihin", "Rabble", "Mob", "3L", 0, 0, "Poor Musket", "", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Ottoman", 1790, 1815, "Guard Janissaries", "CrackLine", "Conscript", "4L", -1, 0, "Poor Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Ottoman", 1790, 1815, "Crack Nizam-I Jadid", "CrackLine", "Conscript", "4L", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Ottoman", 1790, 1815, "Nizam-I Jadid", "Veteran", "Conscript", "4L", 0, 0, "Musket", "Average", "Average", false}))

	// Sweden
	gameData.Insert(DataMap("Infantry", Infantry{"Sweden", 1792, 1815, "Line", "Veteran", "Prussian", "3L", 0, 0, "Musket", "", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Sweden", 1792, 1815, "Light", "Veteran", "Prussian", "3L", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Sweden", 1792, 1815, "Jager", "CrackLine", "Prussian", "3E", 0, 0, "Musket", "Good", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Sweden", 1792, 1815, "Guard", "Grenadier", "Prussian", "3L", 0, 0, "Musket", "Average", "Average", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Sweden", 1792, 1815, "Guard Light", "Grenadier", "Prussian", "3E", 0, 0, "Musket", "Good", "Average", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Sweden", 1792, 1815, "Leib Grenadier", "Grenadier", "Prussian", "3L", 0, 0, "Musket", "Average", "Average", true}))

	// Denmark
	gameData.Insert(DataMap("Infantry", Infantry{"Denmark", 1792, 1815, "Line", "Veteran", "French", "3L", 0, 2, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Denmark", 1792, 1815, "Light", "Regular", "French", "2L 1E", 0, 2, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Denmark", 1792, 1815, "Guard", "Grenadier", "French", "2L", 0, 2, "Musket", "Average", "Average", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Denmark", 1792, 1815, "Leib Grenadier", "Elite", "French", "2L 2E", 0, 2, "Musket", "Average", "Average", false}))

	// United States
	gameData.Insert(DataMap("Infantry", Infantry{"United States", 1812, 1815, "Crack Line", "CrackLine", "French", "2E", 0, 1, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"United States", 1812, 1815, "Veteran Line", "Veteran", "French", "2E", 0, 1, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"United States", 1812, 1815, "Line", "Regular", "French", "2E", 0, 1, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"United States", 1812, 1815, "Volunteer", "Regular", "French", "2E", 0, 1, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"United States", 1812, 1815, "Northern Militia", "Militia", "French", "2E", 0, 2, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"United States", 1812, 1815, "Southern Militia", "Landwehr", "French", "2E", 0, 3, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"United States", 1812, 1815, "Kentucky Militia", "Conscript", "French", "2E", 0, 4, "Rifle", "Good", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"United States", 1812, 1815, "Marines", "Grenadier", "French", "2E", 0, 2, "Musket", "Good", "Average", true}))

	// Dutch Belgian
	gameData.Insert(DataMap("Infantry", Infantry{"Dutch Belgian", 1815, 1815, "Line", "Landwehr", "Conscript", "3L 1S", 0, 1, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Dutch Belgian", 1815, 1815, "Light", "Conscript", "Light Infantry", "2L 3S", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Dutch Belgian", 1815, 1815, "Militia", "Militia", "Militia", "4L", 0, 0, "Musket", "", "Average", false}))

	// Hannover
	gameData.Insert(DataMap("Infantry", Infantry{"Hannover", 1815, 1815, "Line", "Conscript", "Conscript", "4L", 0, 0, "Musket", "", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Hannover", 1815, 1815, "Landwehr", "Landwehr", "Militia", "4L", 0, 0, "Musket", "", "Average", false}))

	// Grand Duchy of Brunswick
	gameData.Insert(DataMap("Infantry", Infantry{"Brunswick", 1792, 1806, "Line", "Conscript", "Conscript", "3L", 0, 0, "Musket", "", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Brunswick", 1815, 1815, "Line", "Conscript", "Conscript", "3L", 0, 0, "Musket", "", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Brunswick", 1792, 1806, "Light", "Conscript", "Conscript", "2L 1E", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Brunswick", 1815, 1815, "Light", "Conscript", "Conscript", "2L 1E", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Brunswick", 1815, 1815, "Avant Guard", "Veteran", "Light Infantry", "3E", 0, 0, "Musket", "Good", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Brunswick", 1809, 1811, "Oels", "CrackLine", "Light Infantry", "3E", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Brunswick", 1812, 1814, "Oels", "Veteran", "Light Infantry", "3E", 0, 0, "Musket", "Average", "Average", false}))

	// Electorate of Hessen-Kassel
	gameData.Insert(DataMap("Infantry", Infantry{"Hessen-Kassel", 1792, 1806, "Line", "Veteran", "OldSchool", "4L", 0, 0, "Musket", "", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Hessen-Kassel", 1809, 1809, "Line", "Conscript", "Conscript", "3L", 0, 0, "Musket", "", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Hessen-Kassel", 1813, 1815, "Line", "Regular", "French", "2L 1E", 0, 0, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Hessen-Kassel", 1795, 1806, "Grenadier", "CrackLine", "OldSchool", "3L", 0, 0, "Musket", "", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Hessen-Kassel", 1809, 1809, "Grenadier", "Veteran", "OldSchool", "3L", 0, 0, "Musket", "", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Hessen-Kassel", 1793, 1806, "Jager", "Veteran", "OldSchool", "2L", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Hessen-Kassel", 1813, 1815, "Jager", "Veteran", "Light Infantry", "3E", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Hessen-Kassel", 1793, 1806, "Guard Regiment", "Elite", "OldSchool", "3L", 0, 0, "Musket", "", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Hessen-Kassel", 1813, 1815, "Guard Regiment", "CrackLine", "OldSchool", "2L 1E", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Hessen-Kassel", 1813, 1815, "Landwehr", "Landwehr", "Conscript", "4L", 0, 0, "Musket", "", "Average", false}))

	// Kingdom of Northern Italy
	gameData.Insert(DataMap("Infantry", Infantry{"Northern Italy", 1805, 1809, "Line", "Regular", "OldSchool", "4L", 1, 0, "Musket", "Poor", "GoodPoor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Northern Italy", 1810, 1812, "Line", "Veteran", "French", "2L 1E", 1, 0, "Musket", "Average", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Northern Italy", 1813, 1814, "Line", "Veteran", "French", "2L 1E", 1, 0, "Musket", "Average", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Northern Italy", 1814, 1814, "Conscript Line", "Conscript", "Conscript", "2L 1E", 1, 0, "Musket", "Poor", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Northern Italy", 1805, 1808, "Light", "Veteran", "Light Infantry", "3E", 1, 0, "Musket", "Average", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Northern Italy", 1809, 1812, "Light", "CrackLine", "Light Infantry", "3E", 1, 0, "Musket", "Good", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Northern Italy", 1792, 1815, "Milan Foot Guard", "CrackLine", "French", "4L 1S", 1, 0, "Musket", "Average", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Northern Italy", 1792, 1815, "Venetian Guard", "Regular", "French", "4L 1S", 1, 0, "Musket", "Poor", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Northern Italy", 1792, 1815, "Dalmation Regt", "Veteran", "French", "2L 1E", 1, 0, "Musket", "Average", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Northern Italy", 1792, 1815, "Foreign Regt", "Landwehr", "Conscript", "2L 1E", 1, 0, "Musket", "", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Northern Italy Guard", 1792, 1815, "Grenadier", "Grenadier", "French", "2L 1E", 1, 0, "Musket", "Good", "Excellent", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Northern Italy Guard", 1792, 1815, "Chasseur", "Grenadier", "Light Infantry", "1L 2E", 1, 0, "Musket", "Excellent", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Northern Italy Guard", 1792, 1815, "Velite", "Elite", "Light Infantry", "1L 2E", 1, 0, "Musket", "Good", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Northern Italy Guard", 1792, 1815, "Conscript", "Elite", "French", "3L", 1, 0, "Musket", "Good", "Good", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Northern Italy Guard", 1792, 1815, "Marine", "CrackLine", "French", "3L", 1, 0, "Musket", "Average", "Good", true}))

	// Grand Duchy of Warsaw
	gameData.Insert(DataMap("Infantry", Infantry{"Duchy of Warsaw", 1807, 1809, "Line", "Regular", "French", "2L 1E", 1, 0, "Musket", "Poor", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Duchy of Warsaw", 1810, 1812, "Crack Line", "CrackLine", "French", "2L 1E", 1, 0, "Musket", "Average", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Duchy of Warsaw", 1810, 1812, "Line", "Veteran", "French", "2L 1E", 1, 0, "Musket", "Average", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Duchy of Warsaw", 1813, 1814, "Line", "Elite", "French", "2L 1E", 1, 0, "Musket", "Average", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Duchy of Warsaw", 1813, 1813, "Guard Btn", "Grenadier", "French", "2L 1E", 1, 0, "Musket", "Good", "Good", false}))

	// Poland
	gameData.Insert(DataMap("Infantry", Infantry{"Poland", 1792, 1794, "Line", "CrackLine", "OldSchool", "3L", 0, -2, "Musket", "", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Poland", 1792, 1794, "Militia", "Landwehr", "Militia", "3L", 0, 0, "Musket", "", "Average", false}))

	// Switzerland
	gameData.Insert(DataMap("Infantry", Infantry{"Swiss", 1805, 1815, "Line", "Landwehr", "French", "2L 1E", 0, 1, "Musket", "Average", "Average", false}))

	// Kingdom of Holland
	gameData.Insert(DataMap("Infantry", Infantry{"Holland", 1806, 1810, "Line", "Conscript", "French", "2L 1E", 2, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Holland", 1806, 1810, "Light", "Veteran", "French", "2L 1E", 2, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Holland", 1806, 1810, "Grenadier Guard", "Grenadier", "French", "2L 1E", 2, 0, "Musket", "Average", "Average", false}))

	// Batavian Republic
	gameData.Insert(DataMap("Infantry", Infantry{"Batavia", 1795, 1798, "Line", "Landwehr", "OldSchool", "3L 1S", 0, 0, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Batavia", 1799, 1806, "Line", "Conscript", "OldSchool", "3L 1S", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Batavia", 1795, 1806, "Foreign Regt", "Conscript", "OldSchool", "2L 1S", 0, 0, "Musket", "", "Average", false}))

	// Kingdom of Two Sicilies
	gameData.Insert(DataMap("Infantry", Infantry{"Sicily", 1792, 1815, "Line", "Landwehr", "Conscript", "3L", 2, 0, "Musket", "", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Sicily", 1792, 1805, "Light", "Conscript", "Conscript", "2L 2S", 2, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Sicily", 1806, 1815, "Light", "Conscript", "Conscript", "3E", 2, 0, "Musket", "Average", "Average", false}))

	// Kingdom of Naples
	gameData.Insert(DataMap("Infantry", Infantry{"Naples", 1806, 1815, "Line", "Landwehr", "Conscript", "3L", 2, 0, "Musket", "", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Naples", 1806, 1815, "Light", "Conscript", "Conscript", "2L 2S", 2, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Naples", 1806, 1815, "Grenadier Guard", "Elite", "French", "2L 1E", 2, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Naples", 1806, 1815, "Guard Marine", "CrackLine", "French", "3L 1E", 2, 0, "Musket", "", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Naples", 1806, 1815, "Guard Velite", "CrackLine", "French", "3L 1E", 2, -1, "Musket", "", "Average", false}))

	// Papal States
	gameData.Insert(DataMap("Infantry", Infantry{"Papal States", 1792, 1815, "Papal Line", "Landwehr", "OldSchool", "3L", 0, 0, "Musket", "", "Average", false}))

	// Sardinia Piedmont
	gameData.Insert(DataMap("Infantry", Infantry{"Sardinia Piedmont", 1792, 1796, "Line", "Conscript", "OldSchool", "3L", 0, 0, "Musket", "", "Average", false}))

	// Helvetian Republic
	gameData.Insert(DataMap("Infantry", Infantry{"Helvetian", 1792, 1815, "Line", "Conscript", "French", "3L", 0, 0, "Musket", "Average", "Average", false}))

	// Cisalpine Republic
	gameData.Insert(DataMap("Infantry", Infantry{"Cisalpine", 1792, 1805, "Line", "Conscript", "Conscript", "3L", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Cisalpine", 1792, 1805, "Polish Legion", "CrackLine", "French", "3L", 0, 0, "Musket", "Average", "Average", false}))

	// Cispandane Republic
	gameData.Insert(DataMap("Infantry", Infantry{"Cispandane", 1792, 1805, "Line", "Conscript", "Conscript", "3L", 0, 0, "Musket", "Average", "Average", false}))

	// Irish Rebels
	gameData.Insert(DataMap("Infantry", Infantry{"Irish Rebel", 1792, 1815, "Rebels", "Landwehr", "Militia", "2L 2S", 0, 0, "Musket", "Average", "Average", false}))

	// Persian Empire
	gameData.Insert(DataMap("Infantry", Infantry{"Persian Empire", 1792, 1815, "Name", "Rating", "Drill", "4L", 0, 0, "Musket", "Average", "Average", false}))

	// Indian States
	gameData.Insert(DataMap("Infantry", Infantry{"Indian States", 1792, 1815, "Name", "Rating", "Drill", "4L", 0, 0, "Musket", "Average", "Average", false}))

	// Maharatta
	gameData.Insert(DataMap("Infantry", Infantry{"Maharatta", 1792, 1815, "Name", "Rating", "Drill", "4L", 0, 0, "Musket", "Average", "Average", false}))

	// East India Company
	gameData.Insert(DataMap("Infantry", Infantry{"East India Company", 1792, 1815, "Name", "Rating", "Drill", "4L", 0, 0, "Musket", "Average", "Average", false}))

	// Hyderabad
	gameData.Insert(DataMap("Infantry", Infantry{"Hyderabad", 1792, 1815, "Name", "Rating", "Drill", "4L", 0, 0, "Musket", "Average", "Average", false}))

	// Anhalt
	gameData.Insert(DataMap("Infantry", Infantry{"Anhalt", 1792, 1815, "Name", "Rating", "Drill", "4L", 0, 0, "Musket", "Average", "Average", false}))

	// Westphalia
	gameData.Insert(DataMap("Infantry", Infantry{"Westphalia", 1792, 1815, "Name", "Rating", "Drill", "4L", 0, 0, "Musket", "Average", "Average", false}))

	// Wurttemburg
	gameData.Insert(DataMap("Infantry", Infantry{"Wurttemburg", 1792, 1815, "Name", "Rating", "Drill", "4L", 0, 0, "Musket", "Average", "Average", false}))

	// Oldenburg
	gameData.Insert(DataMap("Infantry", Infantry{"Oldenburg", 1792, 1815, "Name", "Rating", "Drill", "4L", 0, 0, "Musket", "Average", "Average", false}))

	// Baden
	gameData.Insert(DataMap("Infantry", Infantry{"Baden", 1792, 1815, "Name", "Rating", "Drill", "4L", 0, 0, "Musket", "Average", "Average", false}))

	// Bavaria
	gameData.Insert(DataMap("Infantry", Infantry{"Bavaria", 1792, 1815, "Name", "Rating", "Drill", "4L", 0, 0, "Musket", "Average", "Average", false}))

	// Saxony
	gameData.Insert(DataMap("Infantry", Infantry{"Saxony", 1792, 1815, "Name", "Rating", "Drill", "4L", 0, 0, "Musket", "Average", "Average", false}))

	// Hessen-Darmstadt
	gameData.Insert(DataMap("Infantry", Infantry{"Hessen-Darmstadt", 1792, 1815, "Name", "Rating", "Drill", "4L", 0, 0, "Musket", "Average", "Average", false}))

	// Nassau
	gameData.Insert(DataMap("Infantry", Infantry{"Nassau", 1792, 1815, "Name", "Rating", "Drill", "4L", 0, 0, "Musket", "Average", "Average", false}))

	// Wurzburg
	gameData.Insert(DataMap("Infantry", Infantry{"Wurzburg", 1792, 1815, "Name", "Rating", "Drill", "4L", 0, 0, "Musket", "Average", "Average", false}))

	// Kleve-Berg
	gameData.Insert(DataMap("Infantry", Infantry{"Denmark", 1792, 1815, "Name", "Rating", "Drill", "4L", 0, 0, "Musket", "Average", "Average", false}))

	// Frankfurt
	gameData.Insert(DataMap("Infantry", Infantry{"Frankfurt", 1792, 1815, "Name", "Rating", "Drill", "4L", 0, 0, "Musket", "Average", "Average", false}))

	// Lippe
	gameData.Insert(DataMap("Infantry", Infantry{"Lippe", 1792, 1815, "Name", "Rating", "Drill", "4L", 0, 0, "Musket", "Average", "Average", false}))

	// Mecklenburg
	gameData.Insert(DataMap("Infantry", Infantry{"Mecklenburg", 1792, 1815, "Name", "Rating", "Drill", "4L", 0, 0, "Musket", "Average", "Average", false}))

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Add some Cavalry
	// French Chasseur
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1794, 1815, "1st Chasseur", "Elite", 22, 4, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1794, 1815, "5th Chasseur", "Elite", 22, 4, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1794, 1815, "7th Chasseur", "Grenadier", 24, 5, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1794, 1815, "23th Chasseur", "Grenadier", 24, 5, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1791, 1793, "Chasseur", "Regular", 12, 2, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1794, 1801, "Chasseur", "Veteran", 14, 3, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1802, 1807, "Chasseur", "CrackLine", 16, 4, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1808, 1812, "Chasseur", "CrackLine", 16, 5, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1813, 1814, "Chasseur", "Veteran", 14, 2, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1815, 1815, "Chasseur", "CrackLine", 16, 3, "Light", "Good"}))

	// French Hussar
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1794, 1815, "4/5/7 Hussar", "Grenadier", 24, 6, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1805, 1809, "10th Hussar", "Elite", 20, 6, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1811, 1814, "11th Hussar", "Veteran", 14, 4, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1791, 1793, "Hussar", "Veteran", 14, 3, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1794, 1801, "Hussar", "CrackLine", 16, 3, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1804, 1812, "Hussar", "Elite", 18, 4, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1813, 1814, "Hussar - In Spain", "Elite", 18, 3, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1813, 1814, "Hussar", "CrackLine", 16, 3, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1815, 1815, "Hussar", "Elite", 18, 4, "Light", "Good"}))

	// French Lancer
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1811, 1812, "1st-6th Lancer", "Elite", 22, 4, "Lancer", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1813, 1814, "1st-6th Lancer", "CrackLine", 18, 3, "Lancer", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1815, 1815, "1st-6th Lancer", "Elite", 22, 4, "Lancer", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1792, 1815, "7th-9th Lancer", "Grenadier", 26, 5, "Lancer", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1792, 1815, "Vistula Lancer", "Grenadier", 26, 6, "Lancer", ""}))

	// French Dragoons
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1791, 1801, "Dragoon", "Veteran", 14, 3, "Medium", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1804, 1807, "Dragoon", "CrackLine", 20, 4, "Medium", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1808, 1812, "Dragoon", "Elite", 22, 6, "Medium", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1813, 1815, "Dragoon", "Elite", 22, 4, "Medium", "Average"}))

	// French Cuirassier
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1811, 1814, "13th Cuirassier", "CrackLine", 22, 4, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1811, 1814, "14th Cuirassier", "CrackLine", 18, 4, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1791, 1801, "8th Cuirassier", "Elite", 24, 4, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1791, 1793, "Cavalarie", "CrackLine", 20, 3, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1794, 1801, "Cavalarie", "Elite", 22, 4, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1804, 1812, "Cuirassier", "Grenadier", 26, 6, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1813, 1814, "Cuirassier", "CrackLine", 22, 4, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1815, 1815, "Cuirassier", "Grenadier", 26, 4, "Heavy", ""}))

	// French Carabiniers
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1791, 1793, "Carabinier", "CrackLine", 20, 3, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1794, 1801, "Carabinier", "Elite", 24, 3, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1804, 1812, "Carabinier", "Grenadier", 28, 6, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1813, 1814, "Carabinier", "Elite", 22, 4, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France", 1815, 1815, "Carabinier", "Grenadier", 28, 4, "Heavy", ""}))

	// French Imperial Guard Cavalry
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1813, 1814, "Eclaireurs", "Elite", 20, 4, "Lancer", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1813, 1814, "Guards of Honour", "Elite", 20, 4, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1813, 1814, "Young Guard Sqn", "Elite", 22, 1, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1811, 1812, "Dutch Lancer", "Guard", 28, 6, "Lancer", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1813, 1814, "Dutch Lancer", "Elite", 24, 6, "Lancer", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1808, 1809, "Polish Guard Light Horse", "OldGuard", 33, 5, "Light", "Excellent"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1810, 1812, "Polish Guard Lancer", "OldGuard", 33, 8, "Lancer", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1813, 1814, "Polish Guard Lancer", "OldGuard", 28, 4, "Lancer", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1815, 1815, "Combined Guard Lancer", "Guard", 33, 8, "Lancer", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1796, 1801, "Guides", "OldGuard", 30, 3, "Light", "Excellent"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1804, 1807, "Guard Chasseur", "OldGuard", 33, 6, "Light", "Excellent"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1808, 1812, "Mamaluks", "OldGuard", 33, 8, "Light", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1813, 1814, "Mamaluks", "OldGuard", 28, 6, "Light", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1815, 1815, "Mamaluks", "OldGuard", 33, 10, "Light", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1806, 1808, "Empress Dragoons", "Guard", 26, 4, "Medium", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1809, 1812, "Empress Dragoons", "Guard", 28, 8, "Medium", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1813, 1814, "Empress Dragoons", "Guard", 26, 4, "Medium", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1815, 1815, "Empress Dragoons", "Guard", 28, 7, "Medium", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1802, 1808, "Guard Gendarmes", "OldGuard", 28, 3, "Heavy", "Excellent"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1809, 1812, "Guard Gendarmes", "OldGuard", 30, 3, "Heavy", "Excellent"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1813, 1814, "Guard Gendarmes", "OldGuard", 28, 1, "Heavy", "Excellent"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1815, 1815, "Guard Gendarmes", "OldGuard", 30, 1, "Heavy", "Excellent"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1800, 1800, "Grenadier a Cheval", "OldGuard", 30, 1, "Heavy", "Excellent"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1804, 1812, "Grenadier a Cheval", "OldGuard", 36, 8, "Heavy", "Excellent"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1813, 1814, "Grenadier a Cheval", "OldGuard", 30, 3, "Heavy", "Excellent"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1815, 1815, "Grenadier a Cheval", "OldGuard", 36, 7, "Heavy", "Excellent"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1812, 1813, "Guard Lithuanian Tartar", "Elite", 20, 2, "Lancer", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"France Guard", 1812, 1812, "3rd Guard Lancers", "Elite", 20, 3, "Lancer", "Good"}))

	// Great Britain
	gameData.Insert(DataMap("Cavalry", Cavalry{"Britain", 1792, 1815, "Light Dragoon", "Grenadier", 22, 4, "Light", "Excellent"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Britain", 1792, 1815, "Hussar", "Grenadier", 22, 4, "Light", "Excellent"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Britain", 1792, 1815, "Heavy Dragoon", "Grenadier", 26, 4, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Britain", 1792, 1815, "Dragoon Guards", "Grenadier", 26, 4, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Britain", 1792, 1815, "Life Guards", "Guard", 33, 3, "Medium", ""}))

	// Kingdom of Prussia
	gameData.Insert(DataMap("Cavalry", Cavalry{"Preussen", 1792, 1808, "Dragoon", "CrackLine", 18, 6, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Preussen", 1809, 1815, "Dragoon", "Veteran", 16, 5, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Preussen", 1792, 1815, "Leib Hussar", "Grenadier", 24, 6, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Preussen", 1792, 1815, "5th Hussars", "Grenadier", 24, 6, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Preussen", 1792, 1815, "Hussars", "Grenadier", 22, 6, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Preussen", 1792, 1807, "Towarczys", "CrackLine", 18, 8, "Lancer", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Preussen", 1792, 1808, "Kuirassier", "Grenadier", 26, 6, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Preussen", 1809, 1815, "Kuirassier", "Elite", 22, 5, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Preussen", 1808, 1815, "Uhlan", "Elite", 18, 5, "Lancer", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Preussen", 1809, 1815, "Jager zu Pferd", "Grenadier", 18, 2, "Light", "Excellent"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Preussen", 1813, 1813, "Landwehr", "Landwehr", 12, 3, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Preussen", 1814, 1815, "Landwehr", "Conscript", 14, 3, "Light", ""}))

	// Prussian Guard
	gameData.Insert(DataMap("Cavalry", Cavalry{"Preussen Guard", 1792, 1808, "Guard Kuirassier", "Guard", 30, 6, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Preussen Guard", 1809, 1815, "Guard du Corps", "Guard", 28, 4, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Preussen Guard", 1809, 1815, "Guard Uhlan", "Grenadier", 24, 5, "Lancer", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Preussen Guard", 1809, 1815, "Guard Hussar", "Grenadier", 26, 5, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Preussen Guard", 1809, 1815, "Guard Dragoon", "Elite", 22, 5, "Medium", ""}))

	// Russian
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia", 1792, 1815, "St Petersburg Dragoon", "Elite", 22, 4, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia", 1792, 1815, "Dragoon", "Veteran", 16, 4, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia", 1813, 1815, "Horse Jaeger", "Veteran", 16, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia", 1791, 1804, "Hussar", "Elite", 20, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia", 1805, 1811, "Hussar", "Grenadier", 22, 6, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia", 1812, 1812, "Hussar", "Grenadier", 22, 6, "Lancer", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia", 1813, 1815, "Hussar", "Elite", 20, 4, "Lancer", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia", 1791, 1812, "Uhlan", "Elite", 18, 6, "Lancer", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia", 1813, 1815, "Uhlan", "Veteran", 16, 6, "Lancer", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia", 1791, 1812, "Kuirassier", "Grenadier", 26, 6, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia", 1813, 1815, "Kuirassier", "Elite", 24, 4, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia", 1792, 1815, "Don Cossack", "Conscript", 11, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia", 1792, 1815, "Cossack", "Landwehr", 9, 4, "Light", "Poor"}))

	// Russian Guard Cavalry
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia Guard", 1791, 1812, "Guard Dragoon", "Elite", 24, 6, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia Guard", 1813, 1815, "Guard Dragoon", "Elite", 22, 5, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia Guard", 1791, 1812, "Guard Hussar", "Guard", 30, 8, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia Guard", 1813, 1815, "Guard Hussar", "Guard", 26, 5, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia Guard", 1791, 1812, "Guard Uhlan", "Grenadier", 26, 8, "Lancer", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia Guard", 1813, 1815, "Guard Uhlan", "Grenadier", 24, 5, "Lancer", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia Guard", 1791, 1812, "Guard Kuirassier", "Guard", 30, 1, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia Guard", 1813, 1815, "Guard Kuirassier", "Guard", 26, 4, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia Guard", 1791, 1812, "Chevalier Guard", "Guard", 33, 8, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia Guard", 1813, 1815, "Chevalier Guard", "Guard", 30, 6, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia Guard", 1791, 1812, "Horse Guards", "Guard", 33, 8, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Russia Guard", 1813, 1815, "Horse Guards", "Guard", 30, 5, "Heavy", ""}))

	// Austrian
	gameData.Insert(DataMap("Cavalry", Cavalry{"Austria", 1791, 1815, "OReilly Chevauleger", "Grenadier", 28, 8, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Austria", 1791, 1815, "Elite Chevauleger", "Elite", 24, 8, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Austria", 1791, 1812, "Chevauleger", "Elite", 18, 8, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Austria", 1813, 1815, "Chevauleger", "CrackLine", 16, 6, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Austria", 1791, 1815, "Blankenstein Hussar", "Grenadier", 24, 8, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Austria", 1791, 1812, "Hussar", "Grenadier", 22, 8, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Austria", 1813, 1815, "Hussar", "Elite", 20, 8, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Austria", 1791, 1815, "Insurrection", "Conscript", 12, 4, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Austria", 1791, 1815, "Erz. Johan Dragoon", "Elite", 22, 6, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Austria", 1791, 1815, "Dragoon", "CrackLine", 16, 6, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Austria", 1791, 1812, "Kuirassier", "Grenadier", 24, 6, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Austria", 1813, 1815, "Kuirassier", "Elite", 22, 5, "Heavy", ""}))

	// Sweden
	gameData.Insert(DataMap("Cavalry", Cavalry{"Sweden", 1792, 1815, "Light Dragoons", "CrackLine", 16, 5, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Sweden", 1792, 1815, "Mounted Jager", "CrackLine", 16, 5, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Sweden", 1792, 1815, "Hussar", "Elite", 18, 5, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Sweden", 1792, 1815, "Dragoon", "Veteran", 16, 5, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Sweden", 1792, 1815, "Kuirassier", "Elite", 22, 5, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Sweden", 1792, 1815, "Carabinier", "Elite", 24, 5, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Sweden", 1792, 1815, "Leib Guard", "Grenadier", 22, 5, "Medium", "Average"}))

	// Ottoman Empire
	gameData.Insert(DataMap("Cavalry", Cavalry{"Ottoman", 1792, 1815, "Suvarileris", "Grenadier", 16, 6, "Lancer", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Ottoman", 1792, 1815, "Suvarileri Guard", "Grenadier", 24, 6, "Lancer", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Ottoman", 1792, 1815, "Sipahis Heavy", "Grenadier", 22, 6, "Medium", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Ottoman", 1792, 1815, "Sipahis Elite", "Elite", 20, 6, "Lancer", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Ottoman", 1792, 1815, "Sipahis Crack", "CrackLine", 16, 6, "Lancer", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Ottoman", 1792, 1815, "Sipahis Veteran", "Veteran", 14, 6, "Lancer", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Ottoman", 1792, 1815, "Sipahis", "Regular", 12, 6, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Ottoman", 1792, 1815, "Djellis", "CrackLine", 16, 6, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Ottoman", 1792, 1815, "Yoruks", "Landwehr", 10, 6, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Ottoman", 1792, 1815, "Arab Cavalry", "Militia", 7, 4, "Light", "Average"}))

	// Persian Empire
	gameData.Insert(DataMap("Cavalry", Cavalry{"Persian", 1792, 1815, "Nazam Atli", "Veteran", 16, 6, "Lancer", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Persian", 1792, 1815, "Ristalische", "Landwehr", 10, 6, "Lancer", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Persian", 1792, 1815, "Sakhlu", "Militia", 8, 4, "Light", ""}))

	// Indian States
	gameData.Insert(DataMap("Cavalry", Cavalry{"Maharatta", 1792, 1815, "Line", "Regular", 14, 4, "Lancer", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Maharatte", 1792, 1815, "Militia", "Conscript", 12, 4, "Light", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Mysorean", 1792, 1815, "Kuzzaks", "Conscript", 12, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Mysorean", 1792, 1815, "Silahdars", "Conscript", 14, 4, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Mysorean", 1792, 1815, "Sawar Askars", "Conscript", 14, 4, "Lancer", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"East India Company", 1792, 1815, "Madras", "CrackLine", 20, 4, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Nizam of Hyderbad", 1792, 1815, "Line", "Conscript", 12, 4, "Light", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Nizam of Hyderbad", 1792, 1815, "Levi", "Militia", 7, 4, "Light", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Nizam of Hyderbad", 1792, 1815, "Jats", "CrackLine", 18, 4, "Lancer", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Nizam of Hyderbad", 1792, 1815, "Sikhs", "CrackLine", 18, 4, "Lancer", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Nizam of Hyderbad", 1792, 1815, "Afghani Tribes", "Veteran", 14, 4, "Light", ""}))

	// Revolutionary States
	gameData.Insert(DataMap("Cavalry", Cavalry{"Poland", 1792, 1794, "Line", "CrackLine", 18, 4, "Lancer", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Emigree", 1792, 1815, "Line", "CrackLine", 16, 4, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Batavian Republic", 1792, 1815, "Line", "Regular", 14, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Helvetian Republic", 1792, 1815, "Line", "Regular", 14, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Two Sicilies", 1792, 1797, "Chevauleger", "Elite", 20, 4, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Two Sicilies", 1792, 1797, "Dragoon", "Elite", 20, 4, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Sardinia-Piedmont", 1792, 1796, "Chevauleger", "Regular", 14, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Sardinia-Piedmont", 1792, 1796, "Dragoon", "Regular", 16, 4, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Cisalpine Republic", 1796, 1805, "Hussar", "CrackLine", 16, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Cisalpine Republic", 1796, 1803, "Polish Legion", "CrackLine", 18, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Cispandane", 1796, 1805, "Chasseur", "CrackLine", 20, 4, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Irish Rebels", 1792, 1815, "Rebels", "Regular", 12, 4, "Light", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Misc Revolutionary", 1792, 1804, "Line", "Regular", 14, 4, "Light", ""}))

	// Denmark (French Allied)
	gameData.Insert(DataMap("Cavalry", Cavalry{"Denmark", 1792, 1815, "Light Dragoon", "Regular", 14, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Denmark", 1792, 1815, "Hussar", "Veteran", 16, 4, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Denmark", 1792, 1815, "Dragoon", "Regular", 14, 4, "Medium", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Denmark", 1792, 1815, "Heavy Cavalry", "CrackLine", 18, 4, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Denmark", 1792, 1815, "Leib Heavy Cavalry", "Elite", 20, 4, "Medium", ""}))

	// Kingdom Northern Italy
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Nth Italy", 1812, 1812, "Chasseur", "CrackLine", 16, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Nth Italy", 1813, 1814, "Chasseur", "Veteran", 14, 3, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Nth Italy", 1805, 1812, "Dragoon Napoleone", "Grenadier", 26, 4, "Medium", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Nth Italy", 1813, 1814, "Dragoon Napoleone", "Elite", 22, 3, "Medium", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Nth Italy", 1805, 1812, "Dragoon Regina", "Elite", 22, 4, "Medium", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Nth Italy", 1813, 1814, "Dragoon Regina", "CrackLine", 20, 3, "Medium", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Nth Italy", 1805, 1812, "Guards of Honour", "Grenadier", 26, 4, "Medium", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Nth Italy", 1813, 1814, "Guards of Honour", "Grenadier", 24, 3, "Medium", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Nth Italy", 1805, 1812, "Guard Dragoon", "Grenadier", 26, 4, "Medium", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Nth Italy", 1813, 1814, "Guard Dragoon", "Grenadier", 24, 3, "Medium", "Poor"}))

	// Kingdom of Holland (French Allied)
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Holland", 1806, 1810, "Hussar", "Veteran", 14, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Holland", 1806, 1810, "Cuirassier", "CrackLine", 20, 4, "Heavy", ""}))

	// Kingdom of Naples
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Naples", 1792, 1815, "Chasseur", "Conscript", 12, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Naples", 1792, 1815, "Lancer", "Veteran", 14, 4, "Lancer", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Naples", 1792, 1815, "Guard Velite", "Elite", 22, 4, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Naples", 1792, 1815, "Guard of Honour", "Elite", 20, 4, "Light", "Average"}))

	// Kingdom of Spain (under so called 'King' Joseph)
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Spain (Joseph)", 1792, 1815, "Line", "Veteran", 14, 3, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kingdom of Spain (Joseph)", 1792, 1815, "Guard", "CrackLine", 18, 3, "Light", "Poor"}))

	// Grand Duchy of Warsaw
	gameData.Insert(DataMap("Cavalry", Cavalry{"Duchy of Warsaw", 1807, 1812, "Chasseur", "Elite", 18, 4, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Duchy of Warsaw", 1813, 1814, "Chasseur", "Elite", 16, 4, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Duchy of Warsaw", 1807, 1812, "Hussar", "Elite", 22, 4, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Duchy of Warsaw", 1813, 1814, "Hussar", "Elite", 20, 4, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Duchy of Warsaw", 1807, 1812, "Uhlan", "Elite", 22, 4, "Lancer", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Duchy of Warsaw", 1813, 1814, "Uhlan", "Elite", 20, 4, "Lancer", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Duchy of Warsaw", 1807, 1812, "Cuirassier", "Grenadier", 26, 3, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Duchy of Warsaw", 1813, 1814, "Cuirassier", "Grenadier", 24, 2, "Heavy", ""}))

	// Rhine Confederation - Bavaria
	gameData.Insert(DataMap("Cavalry", Cavalry{"Bavaria", 1792, 1815, "Chevauleger", "CrackLine", 16, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Bavaria", 1792, 1815, "Dragoon", "CrackLine", 16, 4, "Light", "Poor"}))

	// Rhine Confederation - Wurttemburg
	gameData.Insert(DataMap("Cavalry", Cavalry{"Wurttemburg", 1792, 1815, "Chasseur", "CrackLine", 16, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Wurttemburg", 1792, 1815, "Chevauleger", "CrackLine", 16, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Wurttemburg", 1792, 1815, "Dragoon", "CrackLine", 16, 4, "Light", "Poor"}))

	// Rhine Confederation - Westphalia
	gameData.Insert(DataMap("Cavalry", Cavalry{"Westphalia", 1792, 1815, "Chasseur", "CrackLine", 16, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Westphalia", 1806, 1812, "Hussar", "Elite", 18, 4, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Westphalia", 1813, 1814, "Hussar", "Elite", 16, 4, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Westphalia", 1806, 1812, "Cuirassier", "Grenadier", 26, 4, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Westphalia", 1813, 1814, "Cuirassier", "Grenadier", 24, 4, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Westphalia", 1806, 1812, "Guard Chevauleger", "Grenadier", 24, 4, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Westphalia", 1813, 1814, "Guard Chevauleger", "Grenadier", 22, 4, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Westphalia", 1806, 1812, "Guard du Corps", "Grenadier", 24, 4, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Westphalia", 1813, 1814, "Guard du Corps", "Grenadier", 22, 4, "Medium", ""}))

	// Rhine Confederation - Saxony
	gameData.Insert(DataMap("Cavalry", Cavalry{"Saxony", 1792, 1815, "Chevauleger", "Grenadier", 26, 6, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Saxony", 1792, 1815, "Hussar", "Grenadier", 26, 5, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Saxony", 1792, 1815, "Cuirassier", "Guard", 30, 5, "Heavy", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Saxony", 1792, 1815, "Carabinier", "Guard", 30, 5, "Heavy", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Saxony", 1792, 1815, "Guard du Corps", "Guard", 33, 5, "Heavy", "Poor"}))

	// Rhine Confederation - Various Minor States
	gameData.Insert(DataMap("Cavalry", Cavalry{"Anhalt", 1813, 1813, "Chasseur", "Conscript", 12, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kleve-Berg", 1807, 1809, "Chevauleger", "Elite", 22, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kleve-Berg", 1810, 1813, "Lancer", "Grenadier", 24, 4, "Lancer", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Baden", 1792, 1815, "Light Dragoon", "Elite", 20, 5, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Baden", 1792, 1815, "Hussar", "Grenadier", 26, 5, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Hesse-Darmstadt", 1792, 1815, "Chevauleger", "Grenadier", 26, 4, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Wurzburg", 1792, 1815, "Dragoon", "CrackLine", 18, 3, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Nassau", 1792, 1815, "Chevauleger", "Veteran", 14, 3, "Light", "Average"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Mecklemburg", 1813, 1814, "Hussar", "Elite", 20, 4, "Light", "Good"}))

	// Kings German Legion (British Allied)
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kings German Legion", 1792, 1815, "Light Dragoon", "Grenadier", 24, 4, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kings German Legion", 1792, 1815, "1st Hussar", "Grenadier", 26, 4, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kings German Legion", 1792, 1815, "2nd/3rd Hussar", "Grenadier", 24, 4, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Kings German Legion", 1792, 1815, "Heavy Dragoon", "Grenadier", 30, 4, "Medium", ""}))

	// Kingdom of Portugal
	gameData.Insert(DataMap("Cavalry", Cavalry{"Portugal", 1807, 1811, "Dragoon", "Landwehr", 12, 3, "Medium", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Portugal", 1812, 1814, "Dragoon", "Conscript", 14, 3, "Light", "Poor"}))

	// Kingdom of Spain
	gameData.Insert(DataMap("Cavalry", Cavalry{"Spain", 1807, 1814, "Militia", "Militia", 4, 3, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Spain", 1807, 1814, "Dragoons Del Rey", "Conscript", 12, 3, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Spain", 1807, 1814, "Dragoon", "Landwehr", 8, 3, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Spain", 1807, 1814, "Hussar", "Landwehr", 8, 3, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Spain", 1807, 1814, "Lancer", "Landwehr", 8, 3, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Spain", 1807, 1808, "Guard Cavalry", "Conscript", 14, 3, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Spain", 1807, 1814, "Cavalry", "Landwehr", 6, 3, "Heavy", ""}))

	// Dutch Belgian
	gameData.Insert(DataMap("Cavalry", Cavalry{"Netherlands", 1815, 1815, "Light Dragoon", "Conscript", 12, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Netherlands", 1815, 1815, "Hussar", "Conscript", 12, 4, "Light", "Poor"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Netherlands", 1815, 1815, "Carabinier", "Veteran", 14, 4, "Medium", ""}))

	// Brunswick
	gameData.Insert(DataMap("Cavalry", Cavalry{"Brunswick", 1808, 1815, "Hussar", "Elite", 18, 4, "Light", "Good"}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Brunswick", 1815, 1815, "Uhlan", "Elite", 18, 1, "Lancer", "Average"}))

	// Grand Duchy of Hanover
	gameData.Insert(DataMap("Cavalry", Cavalry{"Hanover", 1792, 1805, "Cavalry", "CrackLine", 16, 3, "Medium", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Hanover", 1815, 1815, "Hussar", "Landwehr", 10, 5, "Light", "Poor"}))

	// Electorate of Hessen-Kassel
	gameData.Insert(DataMap("Cavalry", Cavalry{"Hessen-Kassel", 1792, 1805, "Guarde du Corps", "Elite", 20, 3, "Heavy", ""}))
	gameData.Insert(DataMap("Cavalry", Cavalry{"Hessen-Kassel", 1792, 1805, "Line", "CrackLine", 16, 4, "Medium", ""}))

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Add some Artillery
	// French Line
	gameData.Insert(DataMap("Artillery", Artillery{"France", 1792, 1815, "Line Reserve", "CrackLine", 1, "12pdr", "6\"", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"France", 1791, 1806, "Line", "CrackLine", 1, "8pdr", "5.5\"", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"France", 1807, 1815, "Line", "CrackLine", 1, "6pdr", "5.5\"", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"France", 1812, 1812, "Regimental", "Veteran", 1, "4pdr", "", 1, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"France", 1813, 1814, "Regimental", "Veteran", 1, "6pdr", "", 1, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"France", 1791, 1809, "Horse", "Elite", 1, "8pdr", "5.5\"", 3, true}))
	gameData.Insert(DataMap("Artillery", Artillery{"France", 1810, 1815, "Horse", "Elite", 1, "6pdr", "5.5\"", 3, true}))

	// French Guard
	gameData.Insert(DataMap("Artillery", Artillery{"France Guard", 1812, 1815, "Young Guard", "Grenadier", 1, "6pdr", "5.5\"", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"France Guard", 1804, 1805, "Guard Horse", "OldGuard", 0, "8pdr", "5.5\"", 4, true}))
	gameData.Insert(DataMap("Artillery", Artillery{"France Guard", 1806, 1806, "Volante", "OldGuard", 0, "8pdr", "5.5\"", 3, true}))
	gameData.Insert(DataMap("Artillery", Artillery{"France Guard", 1807, 1815, "Volante", "OldGuard", 0, "6pdr", "5.5\"", 4, true}))
	gameData.Insert(DataMap("Artillery", Artillery{"France Guard", 1808, 1815, "Guard Reserve", "OldGuard", 0, "12pdr", "6\"", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"France Guard", 1810, 1815, "Guard Divisional", "OldGuard", 0, "6pdr", "5.5\"", 4, true}))

	// British
	gameData.Insert(DataMap("Artillery", Artillery{"Britain", 1792, 1809, "Royal Foot", "Grenadier", 1, "6pdr", "5.5\"", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Britain", 1810, 1815, "Royal Foot", "Grenadier", 1, "9pdr", "5.5\"", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Britain", 1792, 1809, "Royal Horse", "Grenadier", 1, "6pdr", "5.5\"", 3, true}))
	gameData.Insert(DataMap("Artillery", Artillery{"Britain", 1810, 1815, "Royal Horse", "Grenadier", 1, "9pdr", "5.5\"", 3, true}))

	// Russian
	gameData.Insert(DataMap("Artillery", Artillery{"Russia Guard", 1792, 1810, "Guard", "Guard", 1, "12pdr", "18pdr L", 5, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Russia Guard", 1811, 1815, "Guard", "Guard", 1, "12pdr", "18pdr L", 6, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Russia Guard", 1792, 1810, "Guard Horse", "Grenadier", 1, "6pdr", "9pdr L", 5, true}))
	gameData.Insert(DataMap("Artillery", Artillery{"Russia Guard", 1811, 1815, "Guard Horse", "Grenadier", 1, "6pdr", "9pdr L", 4, true}))

	gameData.Insert(DataMap("Artillery", Artillery{"Russia", 1792, 1815, "Line Heavy", "Elite", 2, "12pdr", "18pdr L", 6, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Russia", 1792, 1815, "Line Light", "Elite", 2, "6pdr", "9pdr L", 6, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Russia", 1792, 1809, "Battalion Guns", "Veteran", 2, "6pdr", "", 1, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Russia", 1792, 1815, "Line", "CrackLine", 2, "6pdr", "9pdr L", 6, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Russia", 1792, 1815, "Flying Cossack", "Conscript", 3, "2pdr", "", 5, true}))

	// Prussian
	gameData.Insert(DataMap("Artillery", Artillery{"Preussen Guard", 1792, 1815, "Guard Horse", "Grenadier", 1, "6pdr", "7pdr", 4, true}))
	gameData.Insert(DataMap("Artillery", Artillery{"Preussen", 1792, 1815, "Line", "Veteran", 2, "6pdr", "10pdr", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Preussen", 1792, 1815, "Reserve", "Veteran", 2, "12pdr", "10pdr", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Preussen", 1792, 1807, "Battalion Guns", "CrackLine", 3, "3pdr", "", 1, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Preussen", 1792, 1815, "Horse", "CrackLine", 2, "6pdr", "7pdr", 3, true}))

	// Austrian
	gameData.Insert(DataMap("Artillery", Artillery{"Austria", 1792, 1815, "Line", "CrackLine", 2, "6pdr", "7pdr", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Austria", 1792, 1815, "Reserve", "CrackLine", 2, "12pdr", "10pdr", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Austria", 1792, 1815, "Brigade", "CrackLine", 2, "6pdr", "", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Austria", 1792, 1800, "Battalion Guns", "CrackLine", 2, "6pdr", "", 1, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Austria", 1792, 1800, "Grenz Bn Guns", "CrackLine", 2, "3pdr", "", 1, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Austria", 1792, 1815, "Kavallarie", "CrackLine", 2, "6pdr", "7pdr", 3, true}))

	//Minor Powers
	gameData.Insert(DataMap("Artillery", Artillery{"Sweden", 1792, 1815, "Line", "CrackLine", 2, "6pdr", "5.5\"", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Sweden", 1792, 1815, "Reserve", "CrackLine", 2, "12pdr", "", 3, false}))

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Command and Control Tables

	// Initiative Modifiers
	gameData.Insert(DataMap("InitTable", InitTable{"Each Bold Leader per engaged ME", 1}))
	gameData.Insert(DataMap("InitTable", InitTable{"CA by Army Commander", 3}))
	gameData.Insert(DataMap("InitTable", InitTable{"CA by Wing Commander", 2}))
	gameData.Insert(DataMap("InitTable", InitTable{"CA by Corps Commander", 1}))

	// Corps Orders
	gameData.Insert(DataMap("CorpsOrder", CorpsOrder{"Attack", []string{"Attack", "Bombard", "Support/Intercept", "March"}, "At least 1 ME must attempt to stay engaged."}))
	gameData.Insert(DataMap("CorpsOrder", CorpsOrder{"Engaged", []string{"Attack", "Bombard", "Defend", "Support/Intercept", "March", "Rest", "Redeploy", "BreakOff"}, "At least 1 ME must attempt to stay engaged."}))
	gameData.Insert(DataMap("CorpsOrder", CorpsOrder{"Defend", []string{"Defend", "Support/Intercept", "March", "Rest", "Redeploy", "BreakOff"}, "At least 1 ME must have a Defend order"}))
	gameData.Insert(DataMap("CorpsOrder", CorpsOrder{"Maneuver", []string{"Support/Intercept", "March"}, ""}))
	gameData.Insert(DataMap("CorpsOrder", CorpsOrder{"Withdraw", []string{"Rearguard", "Defend", "Support/Intercept", "March", "Rest", "BreakOff"}, "The Corps must try to have most of its MEs disengaged until it reaches the destination"}))

	// ME Orders
	gameData.Insert(DataMap("MEOrder", MEOrder{"Attack", "The ME is to engage the enemy", "The ME advances to contact", true, false, false, false, false}))
	gameData.Insert(DataMap("MEOrder", MEOrder{"Engaged", "Attacking ME is to take the objective", "The ME is to fight through to the objective grid", false, true, false, false, false}))
	gameData.Insert(DataMap("MEOrder", MEOrder{"Bombard", "The ME conducts softening up of the objective", "The ME will advance to within 3 grids of the objective and conduct bombardment and skirmish attacks. Pending order of Attack after a specified period.", true, false, false, true, false}))
	gameData.Insert(DataMap("MEOrder", MEOrder{"Defend", "The ME is to hold its ground", "The ME must remain within 1 grid of the defended objective", true, true, false, false, false}))
	gameData.Insert(DataMap("MEOrder", MEOrder{"Support/Intercept", "The ME is to support another ME", "Intercept order may be activated when enemy is within 3 grids", true, false, true, true, false}))
	gameData.Insert(DataMap("MEOrder", MEOrder{"Maneuver", "The ME is to march to a new position", "Will move up to 2 grids off line to avoid contact. Revert to Defend or BreakOff if engaged", true, false, false, true, true}))
	gameData.Insert(DataMap("MEOrder", MEOrder{"RearGuard", "The ME is to fight a delaying action", "Half the units of the ME may fall back 1 grid during GT movement", false, true, false, false, false}))
	gameData.Insert(DataMap("MEOrder", MEOrder{"BreakOff", "The ME is to attempt to disenage and withdraw to a new position", "Receive a full GT movement to disengage. Convert to defend when objective is reached", false, true, false, true, false}))
	gameData.Insert(DataMap("MEOrder", MEOrder{"Screen", "The ME is to screen the advance and conduct reconnaissance", "Convert to Defend when enemy is at 2 grids. Choose pending order of Attack, RearGuard or BreakOff on contact", true, false, true, false, false}))
	gameData.Insert(DataMap("MEOrder", MEOrder{"ReDeploy", "The ME is to perform a general change of facing, formation and relative position", "Shaken if engaged. Recieve 3D6 GT adjustments per half hour", true, false, false, true, true}))
	gameData.Insert(DataMap("MEOrder", MEOrder{"Rest", "The ME is to rest and rally", "Shaken if engaged.", true, false, false, true, true}))

	// Order Arrival Delay
	gameData.Insert(DataMap("OrderArrival", OrderArrival{2, 0, 1}))
	gameData.Insert(DataMap("OrderArrival", OrderArrival{6, 1, 4}))
	gameData.Insert(DataMap("OrderArrival", OrderArrival{12, 2, 8}))
	gameData.Insert(DataMap("OrderArrival", OrderArrival{21, 3, 15}))
	gameData.Insert(DataMap("OrderArrival", OrderArrival{30, 4, 21}))
	gameData.Insert(DataMap("OrderArrival", OrderArrival{39, 5, 27}))
	gameData.Insert(DataMap("OrderArrival", OrderArrival{48, 6, 34}))
	gameData.Insert(DataMap("OrderArrival", OrderArrival{57, 7, 40}))

}
