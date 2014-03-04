package gamedatadb

import (
	"encoding/json"
	"github.com/HouzuoGuo/tiedot/db"
	"log"
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

type EtatMajor struct {
	Nation string
	From   uint16
	To     uint16
	Rating string
	Value  int8
}

// Names of fields here shortened to help make the JSON daatbase more sensible
type Drill struct {
	EF uint8 // Efficiency. Range 1-10. value 1 = 10%, value 10 = 100%
	FR uint8 // Max frontage of this unit in line
	OO uint8 // How many Semi skirmish elements allowed
	SK uint8 // How many full skirmish elements allowed
}

type DrillBook struct {
	Name    string
	Entries map[string]Drill
}

type Equip struct {
	Name   string
	SK     int8
	Volley int8
	Close  int8
	Long   int8
}

type NationalOrg struct {
	Nation     string
	From       uint16
	To         uint16
	InfantryME string
	CavalryME  string
	Corps      string
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

type OrderActivation struct {
	Dice   uint8
	Points int8
}

type OrderActivationMod struct {
	Code        string
	Descr       string
	Points      int8
	CorpsPoints int8
}

type CommanderAction struct {
	Who    string
	Code   string
	Action string
	Cost   uint8
}

type CAScore struct {
	Code  string
	Descr string
	A1    uint8
	A2    uint8
	A3    uint8
	A4    uint8
}

// ME Morale Tables

type MEMoraleTest struct {
	Score   uint8
	Descr   string
	Broken  bool
	Retreat bool
	Shaken  bool
	Steady  bool
	Fatigue uint8
}

type MEMoraleMod struct {
	Code  string
	Descr string
	Value int8
}

type MEPanicTest struct {
	Rating  string
	Broken  uint8
	Shaken  uint8
	CarryOn uint8
}

type MEPanicMod struct {
	Code  string
	Descr string
	Value int8
}

type UnitMoraleTest struct {
	Rating string
	Pass   int8
}

type UnitMoraleMod struct {
	Code  string
	Descr string
	Value int8
}

type FireDisciplineTest struct {
	Rating string
	Pass   int8
	Fire   int8
}

type FireDisciplineMod struct {
	Code  string
	Descr string
	Value int8
}

type InitialBadMorale struct {
	Score int8
	Descr string
	Hits  uint8
	Stay  bool
}

type InitialBadMod struct {
	Code  string
	Descr string
	Value int8
}

type BonusImpulse struct {
	Score       uint8
	Descr       string
	Another     bool
	Fatigue     bool
	OneUnitOnly bool
	FFOnly      bool
}

type BonusImpulseMod struct {
	Code  string
	Descr string
	Value int8
}

type MEFatigue struct {
	Score             uint8
	Descr             string
	OnlyIfNotLastTurn bool
}

type MEFatigueMod struct {
	Code  string
	Descr string
	Value int8
}

type FatigueRecovery struct {
	Score   uint8
	Descr   string
	Recover uint8
}

type FatigueRecoveryMod struct {
	Code  string
	Descr string
	Value int8
}

type BadMoraleRec struct {
	Rating     string
	GoodMorale uint8
	TryAgain   uint8
}

type BadMoraleRecMod struct {
	Code  string
	Descr string
	Value int8
}

// Movement Tables

type GrandTacticalMove struct {
	METype string
	D1     uint8
	D2     uint8
	D3     uint8
	D4     uint8
	D5     uint8
	D6     uint8
}

////////////////////// NEW
type Deployment struct {
	Score  uint8
	Change uint8
}

type DeploymentMod struct {
	Code  string
	Descr string
	Value int8
}

type DeploymentState struct {
	State         uint8
	Descr         string
	Notes         string
	ContactFront  string
	ContactFlank  string
	ContactShaken bool
}

type TacMove struct {
	UnitType string
	Move     uint8
	LtWood   uint8
	HvWood   uint8
	Mud      uint8
	Marsh    uint8
	LoWall   int8
	HiWall   int8
}

type AdditionalMove struct {
	Terrain string
	D1      uint8
	D2      uint8
	D3      uint8
}

type ArtyMove struct {
	Class   string
	Guns    string
	Regular uint8
	Gallop  uint8
	Prolong uint8
}

type ArtyHorseLoss struct {
	Terrain string
	Loss    uint8
}

type BUAMove struct {
	Rating  string
	Ordered uint8 // score to occupy ordered, or exit ordered
	Exit    uint8 // exit disordered
}

type BUAMod struct {
	Code  string
	Descr string
	Value int8
}

type SKRelocate struct {
	Rating string
	Retire uint8
	Move   uint8
	Bold   uint8
}

type SKRelocateMod struct {
	Code  string
	Descr string
	Value int8
}

type SKSupport struct {
	Mode      string
	Marchfeld uint8
	Rolling   uint8
	Rough     uint8
}

type ArtyRelocate struct {
	Class string
	R6    uint8
	R5    uint8
	R4    uint8
	R3    uint8
	R2    uint8
	R1    uint8
	R0    uint8
	W6    uint8
	W5    uint8
	W4    uint8
	W3    uint8
	W2    uint8
	W1    uint8
	W0    uint8
}

type ArtyRelocateMod struct {
	Code  string
	Descr string
	Value int8
}

type FormationChange struct {
	Era       string
	From      string
	To        string
	Trained   int8
	Untrained int8
}

type FormSquare struct {
	Rating string
	From   string
	Grid0  int8
	Grid1  int8
	Grid1D int8
	Grid2  int8
}

type FormSquareMod struct {
	Code  string
	Descr string
	Value int8
}

type FireEffect struct {
	ID    uint8
	Dice  int8
	Descr string
}

type FireChart struct {
	ID        uint8
	SmallArms uint8
	LtArt     uint8
	MdArt     uint8
	MdHvArt   uint8
	HvArt     uint8
}

type FireMod struct {
	Code  string
	Descr string
	Value int8
}

type FireSKMod struct {
	Code  string
	Descr string
	Value int8
}

type SKEffect struct {
	ECode string
	Dice  uint8
	Descr string
}

type FireFight struct {
	Dice      int8
	Descr     string
	Fallback  bool
	HoldCover bool
	Disorder  bool
	Rout      bool
}
type FireFightMod struct {
	Code  string
	Descr string
	Value int8
}

type ArtRange struct {
	Weight string
	Short  uint8
	Medium uint8
	Long   uint8
}

type ArtMod struct {
	Code  string
	Descr string
	Value int8
}

type ArtLimber struct {
	Code  string
	Descr string
	Score uint8
}

type CounterBty struct {
	Score   uint8
	Horses  uint8
	Crew    uint8
	LHorses uint8
	LCrew   uint8
	Caisson bool
}

type Aflame struct {
	Cover     string
	TacMd     uint8
	TacHv     uint8
	BombardMd uint8
	BombardHv uint8
}

type DoubleTeam struct {
	Score uint8
}

type DoubleTeamMod struct {
	Code  string
	Value int8
}

type ArtFate struct {
	Situation string
	Score     uint8
}

type ArtFateMod struct {
	Code  string
	Descr string
	Value int8
}

type Bouncethru struct {
	Score    int8
	Light    uint8
	Medium   uint8
	MediumHv uint8
	Heavy    uint8
}

type BouncethruMod struct {
	Code  string
	Descr string
	Value int8
}

type DefFire struct {
	ID     uint8
	Target string
	Hits1  string
	Hits4  string
	Hits6  string
	Hits8  string
	Hits10 string
}

type DefFireNote struct {
	Code  string
	Descr string
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

	gameData.Insert(DataMap("Drill", DrillBook{"ClassA", map[string]Drill{
		"Line":         Drill{6, 3, 0, 0},
		"Skirmish":     Drill{15, 3, 0, 0},
		"MarchColumn":  Drill{12, 1, 0, 0},
		"AttackColumn": Drill{10, 2, 1, 1},
		"ClosedColumn": Drill{10, 1, 0, 1},
		"Square":       Drill{4, 1, 0, 1}}}))

	gameData.Insert(DataMap("Drill", DrillBook{"ClassB", map[string]Drill{
		"Line":         Drill{5, 3, 0, 0},
		"Skirmish":     Drill{13, 3, 0, 0},
		"MarchColumn":  Drill{10, 1, 0, 0},
		"AttackColumn": Drill{8, 2, 1, 1},
		"ClosedColumn": Drill{8, 1, 0, 1},
		"Square":       Drill{4, 1, 0, 1}}}))

	gameData.Insert(DataMap("Drill", DrillBook{"French", map[string]Drill{
		"Line":         Drill{7, 3, 0, 0},
		"MarchColumn":  Drill{10, 1, 0, 0},
		"AttackColumn": Drill{9, 2, 1, 1},
		"ClosedColumn": Drill{8, 1, 0, 1},
		"Square":       Drill{5, 1, 0, 1}}}))

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
	// Add some Equipment Types

	gameData.Insert(DataMap("Equip", Equip{"Musket", 0, 10, 5, 0}))
	gameData.Insert(DataMap("Equip", Equip{"Carbine", 1, 8, 2, 0}))
	gameData.Insert(DataMap("Equip", Equip{"Superior Musket", 1, 11, 6, 1}))
	gameData.Insert(DataMap("Equip", Equip{"Poor Musket", -1, 8, 4, 0}))
	gameData.Insert(DataMap("Equip", Equip{"Rifle", 3, 8, 6, 2}))
	gameData.Insert(DataMap("Equip", Equip{"Minie", 4, 9, 6, 3}))
	gameData.Insert(DataMap("Equip", Equip{"Bayonet", 0, 0, 0, 0}))
	gameData.Insert(DataMap("Equip", Equip{"Pike", 0, 0, 0, 0}))
	gameData.Insert(DataMap("Equip", Equip{"Breechloader", 4, 14, 7, 4}))
	gameData.Insert(DataMap("Equip", Equip{"Chasspot", 5, 15, 8, 6}))

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
	gameData.Insert(DataMap("Infantry", Infantry{"Preussen Guard", 1792, 1810, "Leibguard", "Guard", "OldSchool", "4L", 0, 0, "Poor Musket", "", "Average", true}))
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
	gameData.Insert(DataMap("Infantry", Infantry{"Persian Empire", 1792, 1815, "Guard (Djambazy)", "Regular", "French", "3L", 0, 0, "Musket", "Poor", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Persian Empire", 1792, 1815, "Line (Sarbaz)", "Landwehr", "Militia", "3L", 0, 0, "Musket", "Poor", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Persian Empire", 1792, 1815, "Militia (Tufendji)", "Rabble", "Mob", "3L", 0, 0, "Musket", "", "Poor", false}))

	// Indian States
	gameData.Insert(DataMap("Infantry", Infantry{"Indian States", 1792, 1815, "Arab Mercenaries", "Veteran", "Conscript", "3L", -2, 0, "Musket", "Average", "Average", false}))

	// Maharatta
	gameData.Insert(DataMap("Infantry", Infantry{"Maharatta", 1792, 1815, "Dupont Brigade", "Veteran", "British", "3L 1E", 0, 0, "Musket", "Poor", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Maharatta", 1792, 1815, "Sombroo Brigade", "Conscript", "Conscript", "3L 1O", 0, 0, "Musket", "Poor", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Maharatta", 1792, 1815, "Pohlman Brigade", "Regular", "British", "3L 1E", 0, 0, "Musket", "Poor", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Maharatta", 1792, 1815, "Other Brigades", "Conscript", "Conscript", "3L 1O", 0, 0, "Musket", "Poor", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Maharatta", 1792, 1815, "Mysorean", "Conscript", "Conscript", "3L 1O", 0, 0, "Musket", "Average", "Poor", false}))

	// East India Company
	gameData.Insert(DataMap("Infantry", Infantry{"East India Company", 1792, 1815, "Madras", "Regular", "British", "4L 1E", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"East India Company", 1792, 1815, "Bombay", "Veteran", "British", "4L 1E", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"East India Company", 1792, 1815, "Carnatcs", "Regular", "British", "4L 1E", 0, 0, "Musket", "Average", "Average", false}))

	// Hyderabad
	gameData.Insert(DataMap("Infantry", Infantry{"Hyderabad", 1792, 1815, "Trained Brides", "Conscript", "Conscript", "3L", 0, 0, "Musket", "Poor", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Hyderabad", 1792, 1815, "Levi", "Landwehr", "Militia", "3L", 0, 0, "Musket", "Poor", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Hyderabad", 1792, 1815, "Jats", "Militia", "Mob", "3L", 0, 0, "Musket", "", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Hyderabad", 1792, 1815, "Sikhs", "Veteran", "British", "3L", 0, 0, "Musket", "", "Poor", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Hyderabad", 1792, 1815, "Afgani Landwehr", "Rating", "Drill", "6S", 0, 0, "Rifle", "Good", "Good", false}))

	// Anhalt
	gameData.Insert(DataMap("Infantry", Infantry{"Anhalt", 1807, 1813, "Line", "Regular", "French", "2L 1E", 1, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Anhalt", 1795, 1802, "Jager", "Regular", "Light Infantry", "3O", 1, 0, "Musket", "Average", "Average", false}))

	// Westphalia
	gameData.Insert(DataMap("Infantry", Infantry{"Westphalia", 1792, 1815, "Line", "Regular", "French", "2L 1E", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Westphalia", 1792, 1815, "Light", "Conscript", "Light Infantry", "3E", 0, 0, "Musket", "Good", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Westphalia", 1792, 1815, "Grenadier Guard", "Veteran", "French", "2L 1E", 2, 0, "Musket", "Average", "Average", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Westphalia", 1792, 1815, "Guard Jager", "Regular", "Light Infantry", "3E", 2, 0, "Rifle", "Good", "Average", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Westphalia", 1792, 1815, "Guard Fusilier", "Regular", "Light Infantry", "3E", 2, 0, "Musket", "Good", "Average", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Westphalia", 1792, 1815, "Guard Conscript", "Conscript", "Conscript", "3L", 0, 0, "Musket", "Average", "Average", false}))

	// Wurttemburg
	gameData.Insert(DataMap("Infantry", Infantry{"Wurttemburg", 1792, 1815, "Line", "Veteran", "French", "3L", 0, 0, "Musket", "", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Wurttemburg", 1792, 1815, "Light", "Veteran", "Light Infantry", "3O", 0, 0, "Superior Musket", "Good", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Wurttemburg", 1792, 1815, "Leib Guard", "Elite", "French", "3L", 0, 0, "Superior Musket", "Good", "Good", true}))

	// Oldenburg
	gameData.Insert(DataMap("Infantry", Infantry{"Oldenburg", 1792, 1815, "Line", "Conscript", "Conscript", "3L", 0, 0, "Musket", "Average", "Average", false}))

	// Baden
	gameData.Insert(DataMap("Infantry", Infantry{"Baden", 1792, 1815, "Line", "Conscript", "French", "3L", 2, 0, "Superior Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Baden", 1792, 1815, "Light", "Regular", "French", "3E", 2, 0, "Superior Musket", "Average", "Average", false}))

	// Bavaria
	gameData.Insert(DataMap("Infantry", Infantry{"Bavaria", 1794, 1808, "Line", "Conscript", "Conscript", "3L", 1, 0, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Bavaria", 1809, 1812, "Line", "Veteran", "French", "2L 1E", 1, 0, "Musket", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Bavaria", 1813, 1815, "Line", "Conscript", "Conscript", "3L", 1, 0, "Rifle", "Poor", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Bavaria", 1794, 1812, "Light", "Veteran", "Light Infantry", "3E", 1, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Bavaria", 1813, 1815, "Light", "Conscript", "Light Infantry", "3E", 1, 0, "Rifle", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Bavaria", 1814, 1815, "Jager Corps", "CrackLine", "Light Infantry", "3E", 1, 0, "Rifle", "Good", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Bavaria", 1814, 1815, "Grenadier Guard", "Elite", "French", "2L 1E", 1, 0, "Musket", "Good", "Good", true}))

	// Saxony
	gameData.Insert(DataMap("Infantry", Infantry{"Saxony", 1794, 1805, "Line", "Regular", "OldSchool", "4L", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Saxony", 1806, 1806, "Line", "Veteran", "OldSchool", "4L", 1, 0, "Superior Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Saxony", 1807, 1812, "Line", "Regular", "French", "2L 1E", 1, 0, "Superior Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Saxony", 1813, 1813, "Line", "Landwehr", "Conscript", "3L", 1, 0, "Superior Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Saxony", 1792, 1815, "Light", "Veteran", "Light Infantry", "3E", 1, 0, "Superior Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Saxony", 1792, 1815, "Field Jager", "Elite", "Light Infantry", "2S", 1, 0, "Rifle", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Saxony", 1792, 1815, "Landwehr", "Landwehr", "Conscript", "3L", 0, 0, "Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Saxony", 1792, 1815, "Leib Grenadier Guard", "Grenadier", "French", "2L 1E", 0, 0, "Musket", "Average", "Average", true}))
	gameData.Insert(DataMap("Infantry", Infantry{"Saxony", 1792, 1815, "Grenadier", "CrackLine", "French", "2L 1E", 0, 0, "Musket", "Average", "Average", true}))

	// Hessen-Darmstadt
	gameData.Insert(DataMap("Infantry", Infantry{"Hessen-Darmstadt", 1792, 1815, "Line", "CrackLine", "French", "2L 1E", 0, 0, "Superior Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Hessen-Darmstadt", 1792, 1815, "Leib Regiment", "Elite", "French", "2L 1E", 0, 0, "Superior Musket", "Average", "Good", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Hessen-Darmstadt", 1792, 1815, "Guard Regiment", "Grenadier", "French", "2L 1E", 0, 0, "Superior Musket", "Good", "Good", true}))

	// Nassau
	gameData.Insert(DataMap("Infantry", Infantry{"Nassau", 1792, 1815, "Line", "Veteran", "French", "3L 1E", 1, 0, "Superior Musket", "Average", "Average", false}))

	// Wurzburg
	gameData.Insert(DataMap("Infantry", Infantry{"Wurzburg", 1792, 1815, "Line", "Veteran", "French", "2L 1E", 1, 0, "Musket", "Average", "Average", false}))

	// Kleve-Berg
	gameData.Insert(DataMap("Infantry", Infantry{"Kleve-Berg", 1806, 1813, "Line", "Regular", "French", "2L 1E", 1, 0, "Musket", "Average", "Average", false}))

	// Frankfurt
	gameData.Insert(DataMap("Infantry", Infantry{"Frankfurt", 1792, 1815, "Line", "Regular", "French", "2L 1E", 1, 0, "Musket", "Average", "Average", false}))

	// Lippe
	gameData.Insert(DataMap("Infantry", Infantry{"Lippe", 1807, 1813, "Line", "Regular", "French", "2L 1E", 1, 0, "Musket", "Average", "Average", false}))

	// Mecklenburg
	gameData.Insert(DataMap("Infantry", Infantry{"Mecklenburg", 1808, 1813, "Line", "Regular", "French", "2L 1E", 0, 0, "Superior Musket", "Average", "Average", false}))
	gameData.Insert(DataMap("Infantry", Infantry{"Mecklenburg", 1808, 1813, "Guard", "Veteran", "French", "1L 1E", 0, 0, "Sucperior Musket", "Average", "Average", false}))

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
	gameData.Insert(DataMap("Artillery", Artillery{"Sweden", 1792, 1815, "Horse", "Elite", 2, "6pdr", "8pdr", 3, true}))
	gameData.Insert(DataMap("Artillery", Artillery{"Sweden", 1792, 1815, "Reserve", "CrackLine", 2, "12pdr", "", 3, false}))

	gameData.Insert(DataMap("Artillery", Artillery{"Denmark", 1792, 1815, "Line", "Regular", 2, "6pdr", "7pdr", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Denmark", 1792, 1815, "Reserve", "Veteran", 2, "10pdr", "10pdr", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Denmark", 1792, 1815, "Horse", "Veteran", 2, "3pdr", "7pdr", 4, true}))

	gameData.Insert(DataMap("Artillery", Artillery{"United States", 1812, 1815, "Line", "CrackLine", 2, "6pdr", "5.5\"", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"United States", 1812, 1815, "Marine", "Guard", 2, "6pdr", "5.5\"", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"United States", 1812, 1815, "Horse", "Veteran", 2, "3pdr", "", 2, true}))

	gameData.Insert(DataMap("Artillery", Artillery{"Portugal", 1792, 1815, "Line", "Veteran", 2, "6pdr", "6\"", 3, false}))

	gameData.Insert(DataMap("Artillery", Artillery{"Spain", 1792, 1815, "Line", "Veteran", 3, "6pdr", "6\"", 3, false}))

	gameData.Insert(DataMap("Artillery", Artillery{"Dutch Belgium", 1815, 1815, "Line", "Conscript", 2, "6pdr", "5.5\"", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Dutch Belgium", 1815, 1815, "Horse", "Veteran", 2, "6pdr", "5.5\"", 4, true}))

	gameData.Insert(DataMap("Artillery", Artillery{"Hannover", 1815, 1815, "Line", "Conscript", 2, "6pdr", "5.5\"", 4, false}))

	gameData.Insert(DataMap("Artillery", Artillery{"Brunswick", 1815, 1815, "Line", "Conscript", 2, "6pdr", "", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Brunswick", 1815, 1815, "Horse", "Veteran", 2, "6pdr", "", 3, true}))

	gameData.Insert(DataMap("Artillery", Artillery{"Switzerland", 1812, 1812, "Line", "CrackLine", 1, "4pdr", "", 2, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Neuchatel", 1812, 1812, "Line", "CrackLine", 1, "4pdr", "", 1, false}))

	gameData.Insert(DataMap("Artillery", Artillery{"Northern Italy", 1792, 1815, "Line", "Veteran", 2, "6pdr", "5.5\"", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Northern Italy", 1792, 1815, "Reserve", "Veteran", 2, "12pdr", "6\"", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Northern Italy", 1812, 1812, "Regt Guns", "", 2, "3pdr", "", 1, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Northern Italy", 1792, 1815, "Horse", "CrackLine", 2, "6pdr", "5.5\"", 3, true}))
	gameData.Insert(DataMap("Artillery", Artillery{"Northern Italy", 1792, 1815, "Guard", "Grenadier", 2, "6pdr", "5.5\"", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Northern Italy", 1805, 1806, "Guard", "Grenadier", 2, "8pdr", "5.5\"", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Northern Italy", 1806, 1814, "Guard", "Grenadier", 2, "6pdr", "5.5\"", 3, false}))

	gameData.Insert(DataMap("Artillery", Artillery{"Holland", 1792, 1815, "Line", "Conscript", 2, "6pdr", "5.5\"", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Holland", 1792, 1815, "Heavy", "Conscript", 2, "12pdr", "6\"", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Holland", 1792, 1815, "Horse", "Regular", 2, "6pdr", "5.5\"", 3, true}))

	gameData.Insert(DataMap("Artillery", Artillery{"Naples", 1792, 1815, "Line", "Rating", 2, "6pdr", "6\"", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Naples", 1792, 1815, "Horse", "Rating", 2, "6pdr", "6\"", 3, true}))
	gameData.Insert(DataMap("Artillery", Artillery{"Naples", 1792, 1815, "Guard", "Rating", 2, "6pdr", "6\"", 3, false}))

	gameData.Insert(DataMap("Artillery", Artillery{"Warsaw", 1792, 1815, "Line", "Veteran", 2, "8pdr", "7pdr", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Warsaw", 1792, 1815, "Reserve", "Veteran", 2, "12pdr", "7pdr", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Warsaw", 1812, 1812, "Bn Guns", "Veteran", 2, "3pdr", "", 1, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Warsaw", 1792, 1815, "Horse", "CrackLine", 2, "6pdr", "7pdr", 3, true}))

	gameData.Insert(DataMap("Artillery", Artillery{"Ottoman", 1792, 1815, "Line", "Conscript", 3, "6pdr", "9pdr", 5, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Ottoman", 1792, 1815, "Topijis", "Conscript", 2, "6pdr", "9pdr", 5, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Ottoman", 1792, 1815, "French Mercenary", "Veteran", 2, "6pdr", "9pdr", 5, false}))

	gameData.Insert(DataMap("Artillery", Artillery{"Bavaria", 1792, 1815, "Line", "Veteran", 2, "6pdr", "7pdr", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Bavaria", 1792, 1815, "Reserve", "Veteran", 2, "12pdr", "7pdr", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Bavaria", 1812, 1812, "Battalion Guns", "Veteran", 2, "3pdr", "", 1, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Bavaria", 1792, 1815, "Horse", "Veteran", 2, "6pdr", "7pdr", 3, true}))

	gameData.Insert(DataMap("Artillery", Artillery{"Saxony", 1792, 1815, "Line", "Regular", 2, "6pdr", "7pdr", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Saxony", 1792, 1815, "Horse", "Regular", 2, "6pdr", "7pdr", 3, true}))
	gameData.Insert(DataMap("Artillery", Artillery{"Saxony", 1805, 1812, "Battalion Guns", "Veteran", 2, "4pdr", "", 1, false}))

	gameData.Insert(DataMap("Artillery", Artillery{"Hessen-Darmstadt", 1792, 1815, "Line", "CrackLine", 2, "6pdr", "7pdr", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Hessen-Darmstadt", 1792, 1815, "Horse", "CrackLine", 2, "6pdr", "7pdr", 3, true}))

	gameData.Insert(DataMap("Artillery", Artillery{"Westphalia", 1792, 1815, "Line", "Veteran", 2, "8pdr", "5.5\"", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Westphalia", 1792, 1815, "Horse", "Veteran", 2, "6pdr", "5.5\"", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Westphalia", 1792, 1815, "Battalion Guns", "Veteran", 2, "6pdr", "", 1, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Westphalia", 1792, 1815, "Guard", "Elite", 2, "6pdr", "5.5\"", 3, false}))

	gameData.Insert(DataMap("Artillery", Artillery{"Wurttemburg", 1792, 1815, "Line", "Veteran", 1, "6pdr", "7pdr", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Wurttemburg", 1792, 1815, "Heavy", "Veteran", 1, "12pdr", "7pdr", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Wurttemburg", 1792, 1815, "Horse", "Veteran", 1, "6pdr", "7pdr", 3, true}))
	gameData.Insert(DataMap("Artillery", Artillery{"Wurttemburg", 1792, 1815, "Guard", "Elite", 1, "6pdr", "7pdr", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Wurttemburg", 1792, 1815, "Guard Heavy", "Elite", 1, "12pdr", "7pdr", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Wurttemburg", 1792, 1815, "Guard Horse", "Elite", 1, "6pdr", "7pdr", 3, true}))

	gameData.Insert(DataMap("Artillery", Artillery{"Baden", 1792, 1815, "Line", "Veteran", 1, "6pdr", "7pdr", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Baden", 1792, 1815, "Heavy", "Veteran", 1, "12pdr", "7pdr", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Baden", 1792, 1815, "Horse", "CrackLine", 1, "6pdr", "7pdr", 3, true}))

	gameData.Insert(DataMap("Artillery", Artillery{"Mecklenburg", 1792, 1815, "Line", "Conscript", 3, "6pdr", "5.5\"", 4, false}))

	gameData.Insert(DataMap("Artillery", Artillery{"Kleve-Berg", 1792, 1815, "Line", "Veteran", 1, "6pdr", "7pdr", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Kleve-Berg", 1792, 1815, "Horse", "CrackLine", 1, "6pdr", "7pdr", 3, true}))

	gameData.Insert(DataMap("Artillery", Artillery{"Persian Empire", 1792, 1815, "Line (Zamburechki)", "CrackLine", 2, "6pdr", "", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Persian Empire", 1792, 1815, "Position", "CrackLine", 3, "12pdr", "7pdr", 6, false}))

	gameData.Insert(DataMap("Artillery", Artillery{"East India Company", 1792, 1815, "Line", "Veteran", 2, "6pdr", "5.5\"", 4, false}))

	gameData.Insert(DataMap("Artillery", Artillery{"Mysorean", 1792, 1815, "Line", "Regular", 3, "6pdr", "5.5\"", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Mysorean", 1792, 1815, "Heavy", "Regular", 3, "12pdr", "5.5\"", 4, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Mysorean", 1792, 1815, "Levi", "Landwehr", 3, "18pdr", "", 4, false}))

	gameData.Insert(DataMap("Artillery", Artillery{"Indian States", 1792, 1815, "Line", "Conscript", 3, "6pdr", "6\"", 3, false}))
	gameData.Insert(DataMap("Artillery", Artillery{"Indian States", 1792, 1815, "Levi", "Landwehr", 3, "18pdr", "", 3, false}))

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Staff Ratings

	gameData.Insert(DataMap("EtatMajor", EtatMajor{"France", 1792, 1795, "Average", 0}))
	gameData.Insert(DataMap("EtatMajor", EtatMajor{"France under Napoleon", 1796, 1800, "Good", 2}))
	gameData.Insert(DataMap("EtatMajor", EtatMajor{"France", 1796, 1800, "Average", 0}))
	gameData.Insert(DataMap("EtatMajor", EtatMajor{"France", 1801, 1814, "Good", 1}))
	gameData.Insert(DataMap("EtatMajor", EtatMajor{"France", 1815, 1815, "Average", 0}))
	gameData.Insert(DataMap("EtatMajor", EtatMajor{"Britain", 1792, 1815, "Average", 0}))
	gameData.Insert(DataMap("EtatMajor", EtatMajor{"Prussia", 1792, 1810, "Poor", -1}))
	gameData.Insert(DataMap("EtatMajor", EtatMajor{"Prussia", 1811, 1815, "Good", 1}))
	gameData.Insert(DataMap("EtatMajor", EtatMajor{"Austria", 1792, 1809, "Poor", -1}))
	gameData.Insert(DataMap("EtatMajor", EtatMajor{"Austria", 1810, 1815, "Average", 0}))
	gameData.Insert(DataMap("EtatMajor", EtatMajor{"Russia", 1792, 1805, "Poor", -1}))
	gameData.Insert(DataMap("EtatMajor", EtatMajor{"Russia", 1806, 1815, "Average", 0}))
	gameData.Insert(DataMap("EtatMajor", EtatMajor{"Spain", 1792, 1808, "Poor", -1}))
	gameData.Insert(DataMap("EtatMajor", EtatMajor{"Spain", 1809, 1815, "Average", 0}))
	gameData.Insert(DataMap("EtatMajor", EtatMajor{"Sweden", 1792, 1810, "Poor", -1}))
	gameData.Insert(DataMap("EtatMajor", EtatMajor{"Sweden", 1811, 1815, "Average", 0}))
	gameData.Insert(DataMap("EtatMajor", EtatMajor{"United States", 1812, 1815, "Average", 0}))
	gameData.Insert(DataMap("EtatMajor", EtatMajor{"Ancien Regimes", 1792, 1815, "Poor", -1}))
	gameData.Insert(DataMap("EtatMajor", EtatMajor{"Other Divisional", 1792, 1815, "Average", 0}))

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

	// Order Activation Points and Modifiers
	gameData.Insert(DataMap("OrderActivation", OrderActivation{0, -1}))
	gameData.Insert(DataMap("OrderActivation", OrderActivation{1, 0}))
	gameData.Insert(DataMap("OrderActivation", OrderActivation{3, 1}))
	gameData.Insert(DataMap("OrderActivation", OrderActivation{6, 2}))
	gameData.Insert(DataMap("OrderActivation", OrderActivation{8, 3}))
	gameData.Insert(DataMap("OrderActivation", OrderActivation{9, 4}))
	gameData.Insert(DataMap("OrderActivation", OrderActivation{11, 5}))
	gameData.Insert(DataMap("OrderActivation", OrderActivation{13, 6}))
	gameData.Insert(DataMap("OrderActivation", OrderActivation{16, 7}))
	gameData.Insert(DataMap("OrderActivation", OrderActivation{18, 8}))
	gameData.Insert(DataMap("OrderActivation", OrderActivation{19, 9}))
	gameData.Insert(DataMap("OrderActivation", OrderActivation{30, 10}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"CC1", "Both Commanders in same grid", 4, 6}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"CU1", "CA to urge order and commanders are within 2 grids", 3, 5}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"NLOS", "No Line of Sight between commanders", -1, -1}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"RVAN", "Receiving Commander has superior vantage point", 0, 2}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"CHAR", "Charismatic Commander activating Attack Order", 4, 4}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"INSP", "Inspirational Commander activating Attack Order", 2, 2}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"UINS", "Uninspiring Commander activating Attack Order", -2, -2}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"RETR", "ME Retreat Order during Corps withdrawal", 5, 0}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"TIRD", "ME is Tired, and ordered to March or Attack", -2, 0}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"BRK", "ME Break off order", 4, 0}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"CORP", "Corps Order", 0, 8}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"GRDB", "Form Grande Battery", 0, -6}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"GSTF", "Good Staff Work", 1, 2}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"PSTF", "Poor Staff Work", -1, -3}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"SNOW", "Snow or Heavy Rain", -2, -4}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"RAIN", "Miserable Rain", -1, -2}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"RIVL", "Commander Rivalry", -2, -12}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"ELIT", "Order to Elite ME", 4, 0}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"C1", "Order is from a Superior Commander", 3, 3}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"C2", "Order is from an Excellent Commander", 2, 2}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"C3", "Order is from a Good Commander", 1, 1}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"C4", "Order is from an Average Commander", 0, 0}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"C5", "Order is from a Poor Commander", -1, -1}))
	gameData.Insert(DataMap("OrderActivationMod", OrderActivationMod{"C6", "Order is from a Despicable Commander", -2, -2}))

	// Commander Actions
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Corps", "MV1", "Move 1 grid", 0}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Corps", "MV3", "Move 3 grids", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Corps", "MSG", "Send Message if Unengaged", 0}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Corps", "ORD", "Issue Orders", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Corps", "EMG", "Send Message if Engaged", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Corps", "BTY", "Commit battery from Corps reserve to ME reserve", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Corps", "BRR", "Form Brigade Reserve", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Corps", "AME", "Attach to ME within 2 grids", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Corps", "DME", "Detach from ME and move 2 grids", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Corps", "DEF", "Convert Maneuvre to Defend during GT movement", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Corps", "INI", "Boost Initiative for ME within 2 grids", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Corps", "FGB", "Form Grand Battery", 2}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Corps", "DGB", "Dissolve Grand Battery", 2}))

	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Initiative", "MV", "Move 1 grid if unattached", 0}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Initiative", "GI", "Gain Impetus", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Initiative", "CB", "Commit Battery from ME reserve", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Initiative", "RA", "Rally any 1 unit in same grid", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Initiative", "CC", "Commit subordinate to unit", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Initiative", "CS", "Commit sappers", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Initiative", "AT", "Attach to unit within 1 grid", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Initiative", "DT", "Detach from unit and move 1 grid", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Initiative", "SK", "Commit skirmishers to screen", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Initiative", "FM", "Follow Me", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Initiative", "WB", "Withdraw Battery to reserve", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Initiative", "BA", "Form bridge assault formation", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Initiative", "RG", "Form Battery from regimental guns", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Initiative", "RD", "Re-Mount Dragoons", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Initiative", "CR", "Commit reserve", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Initiative", "RS", "Resupply 1 unit", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"Initiative", "BG", "Form cavalry battle group", 1}))

	gameData.Insert(DataMap("CommanderAction", CommanderAction{"React", "MV", "Move 1 grid if unattached", 0}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"React", "FC", "Adjust facing or formation for 1 unit within 1 grid", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"React", "GI", "Attempt to gain impetus", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"React", "AT", "Attach to unit within 1 grid", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"React", "DT", "Detach from unit and move 1 grid", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"React", "PB", "Pull back 1 unit", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"React", "SK", "Commit skirmishers to screen", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"React", "RA", "Rally any 1 unit in same grid", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"React", "CV", "Order ME support cavalry to opportunity charge up to 1 grid", 1}))
	gameData.Insert(DataMap("CommanderAction", CommanderAction{"React", "BG", "Form cavalry battle group", 1}))

	gameData.Insert(DataMap("CAScore", CAScore{"A", "Army Commander", 6, 9, 12, 16}))
	gameData.Insert(DataMap("CAScore", CAScore{"B", "Corps Commander", 8, 11, 14, 18}))
	gameData.Insert(DataMap("CAScore", CAScore{"C", "Brigade Leader", 11, 15, 18, 20}))
	gameData.Insert(DataMap("CAScore", CAScore{"D", "Replacement", 13, 15, 18, 20}))

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// National Organisations

	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Austria", 1796, 1808, "Bde of 4-8 Bn", "Regt or Bde", "Ad Hoc columns of 1-4 Bde"}))
	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Austria", 1809, 1815, "Div of 1-3 Bde", "Bde of 2 Regt", "Corps of 1 Light, 2 Line Div"}))

	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Britain", 1792, 1807, "Bde of 2-4 Bn", "Bde of 2 Regt", "Ad Hoc 1-4 Bde"}))
	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Britain", 1808, 1814, "Div of 1-3 Bde", "Bde of 2 Regt", "Divisional"}))

	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Anglo Allied", 1815, 1815, "Div of 2-3 Bde", "Bdes", "Wing Commanders"}))

	gameData.Insert(DataMap("NationalOrg", NationalOrg{"France", 1792, 1804, "Div of 1-3 Demi Bde", "Div of 2-6 Regt", "Divisional"}))
	gameData.Insert(DataMap("NationalOrg", NationalOrg{"France", 1805, 1815, "Div of 6-15 Bn, 1-2 Bty", "Lt Bde of 2-3 Regt, Res Bde of 2-8 Regt, H Bty", "Full Corps Structure"}))

	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Preussen", 1792, 1807, "Bde of 4-6 Bn", "Bde of 1-2 Regt", "Ad Hoc columns"}))
	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Preussen", 1812, 1815, "Bde of 3 Line, 3 Res, 3 LW Bn, Bty, Cav Regt", "Bde of 2-3 Reg", "Corps of 4 Bde, Cav Bde, Res Artillery"}))

	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Russia", 1792, 1805, "Bde of 4-6 Bn", "Bde of 1-2 Regt", "Ad Hoc columns"}))
	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Russia", 1806, 1811, "Div of 2-3 Bde", "Bde of 1-2 Regt", "Ad Hoc columns"}))
	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Russia", 1812, 1815, "Div of 4 Line, 2 Jager Regt, 2-3 Bty", "Div of 2 Bde of 1-2 Regt", "Corps of 2 Div"}))

	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Spain", 1792, 1806, "Bde of 4-8 Bn", "Regt or Bde", "Ad Hoc column of 1-4 Bde"}))
	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Spain", 1807, 1815, "Div of 2-3 Bde", "Bde of 1-2 Regt", "Col of 1-2 Div, no formal Corps structure"}))

	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Turkey", 1792, 1815, "Bde of 4-10 Bn", "Command of 1-6 Regt", "Ad Hoc columns"}))

	gameData.Insert(DataMap("NationalOrg", NationalOrg{"German States", 1792, 1806, "Bde of 4-8 Bn", "Regt or Bde", "Ad Hoc columns of 1-4 Bde"}))
	gameData.Insert(DataMap("NationalOrg", NationalOrg{"German States", 1805, 1815, "Div of 2-3 Bde", "Bde of 1-3 Regt", "Integrated in French Corps"}))

	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Italian States", 1792, 1804, "Bde of 4-6 Bn", "Regt or Bde", "Ad Hoc columns of 1-4 Bde"}))
	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Italian States", 1805, 1815, "Div of 6-15 Bn, 1-2 Bty", "Lt Bde of 2-3 Regt, H Bty", "French Corps Sysstem"}))

	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Denmark", 1792, 1805, "Bde of 4-8 Bn", "Bde of 1-2 Regt", "Ad Hoc columns of 1-4 Bde"}))
	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Denmark", 1806, 1815, "Div of 2-3 Bde", "Bde of 1-2 Regt", "Col of 1-2 Div"}))

	gameData.Insert(DataMap("NationalOrg", NationalOrg{"American", 1812, 1815, "Bde of 4-8 Bn", "Bde of 1-2 Regt", "Column of 1-4 Bde"}))

	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Indian", 1792, 1815, "Bde of 4-10 Bn, some attached Cav", "Feudal levee of various size", "Ad Hoc column of 1-4 Bde"}))
	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Duchy of Warsaw", 1807, 1814, "Div of 6-15 Bn, 1-2 Bty", "Lt Bde of 2-3 Regt, Res Bde of 2-8 Regt, H Bty", "Full Corps Structure"}))

	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Sweden", 1792, 1809, "Bde of 4-8 Bn", "Bde of 1-3 Regt", "Ad Hoc column of 1-4 Bde"}))
	gameData.Insert(DataMap("NationalOrg", NationalOrg{"Sweden", 1810, 1815, "Div of 2-3 Bde", "Bde of 1-3 Regt", "Corps structure"}))

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Morale and Fatigue tables

	gameData.Insert(DataMap("MEMoraleTest", MEMoraleTest{0, "ME Breaks, and falls back 2 grids in Bad Morale", true, true, true, false, 1}))
	gameData.Insert(DataMap("MEMoraleTest", MEMoraleTest{6, "Retreat Shaken. Convert to BreakOff order, and fall back 2 grids", false, true, true, false, 1}))
	gameData.Insert(DataMap("MEMoraleTest", MEMoraleTest{9, "Shaken. Attacks without impetus fall back 2 grids, revert to Defend", false, false, true, false, 0}))
	gameData.Insert(DataMap("MEMoraleTest", MEMoraleTest{11, "Steady", false, false, false, true, 0}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"OldGuard", "OldGuard Morale Grade", 10}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"Guard", "Guard Morale Grade", 9}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"Grenadier", "Grenadier Morale Grade", 8}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"Elite", "Elite Morale Grade", 7}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"CrackLine", "CrackLine Morale Grade", 6}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"Veteran", "Veteran Morale Grade", 5}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"Regular", "Regular Morale Grade", 4}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"Conscript", "Conscript Morale Grade", 3}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"Landwehr", "Landwehr Morale Grade", 2}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"Militia", "Militia Morale Grade", 1}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"Rabble", "Rabble Morale Grade", 0}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"BADI", "Per Infantry Unit in Bad Morale", -2}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"BADC", "Per Cavalry Unit in Bad Morale", -3}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"BADA", "Per Artillery Unit in Bad Morale", -6}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"GOOD", "Per Unit in Good Morale with Full Ammo", 1}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"CAW", "Per Close Action won this turn", 2}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"CAD", "Per Close Action lost this turn", -2}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"Fatigue", "Per Fatigue level past Fresh", -1}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"SQP", "Adjacent ME elected to Sauve qui Peut", -4}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"SP", "Enemy Strongpoint within 3 grids has fallen in the last hour", 4}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"SPH", "Per Structure still held", 1}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"SPL", "Per Structure lost", -1}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"INTER", "Interpenetrated by another ME", -2}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"SHK", "Previously Shaken", -3}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"F1", "Hit in Flank during GT move, in March Column", -2}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"F2", "Hit in Flank during GT move, in Extended March Column", -8}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"F3", "Hit in Flank during GT move, in Regular March Column", -6}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"F4", "Hit in Flank during GT move, in Condensed March Column", -4}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"CF1", "Campaign Fatigue - Weary", -1}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"CF2", "Campaign Fatigue - Haggard", -2}))
	gameData.Insert(DataMap("MEMoraleMod", MEMoraleMod{"CF3", "Campaign Fatigue - Spent", -5}))

	gameData.Insert(DataMap("MEPanicTest", MEPanicTest{"OldGuard", 2, 3, 4}))
	gameData.Insert(DataMap("MEPanicTest", MEPanicTest{"Guard", 3, 5, 6}))
	gameData.Insert(DataMap("MEPanicTest", MEPanicTest{"Grenadier", 4, 6, 7}))
	gameData.Insert(DataMap("MEPanicTest", MEPanicTest{"Elite", 5, 6, 7}))
	gameData.Insert(DataMap("MEPanicTest", MEPanicTest{"CrackLine", 5, 7, 8}))
	gameData.Insert(DataMap("MEPanicTest", MEPanicTest{"Veteran", 6, 7, 8}))
	gameData.Insert(DataMap("MEPanicTest", MEPanicTest{"Regular", 7, 8, 9}))
	gameData.Insert(DataMap("MEPanicTest", MEPanicTest{"Conscript", 8, 9, 10}))
	gameData.Insert(DataMap("MEPanicTest", MEPanicTest{"Landwehr", 9, 10, 11}))
	gameData.Insert(DataMap("MEPanicTest", MEPanicTest{"Militia", 10, 11, 12}))
	gameData.Insert(DataMap("MEPanicTest", MEPanicTest{"Rabble", 11, 12, 13}))
	gameData.Insert(DataMap("MEPanicMod", MEPanicMod{"OG1", "Old Guard Broken within 3 grids", -5}))
	gameData.Insert(DataMap("MEPanicMod", MEPanicMod{"OG2", "Old Guard Shaken within 3 grids", -3}))
	gameData.Insert(DataMap("MEPanicMod", MEPanicMod{"50", "50%+ Units broken or destroyed", -6}))
	gameData.Insert(DataMap("MEPanicMod", MEPanicMod{"25", "25% Units broken or destroyed", -3}))
	gameData.Insert(DataMap("MEPanicMod", MEPanicMod{"Fatigue", "Each Fatigue level over Fresh", -1}))
	gameData.Insert(DataMap("MEPanicMod", MEPanicMod{"SHK", "Already Shaken", -2}))
	gameData.Insert(DataMap("MEPanicMod", MEPanicMod{"TRAP", "Enemy blocks line of retreat", -2}))
	gameData.Insert(DataMap("MEPanicMod", MEPanicMod{"SP", "Each fallen strongpoint within 3 grids", -2}))
	gameData.Insert(DataMap("MEPanicMod", MEPanicMod{"WTH", "ME is under a Withdraw Order", -1}))
	gameData.Insert(DataMap("MEPanicMod", MEPanicMod{"CF1", "Campaign Fatigue - Weary", -2}))
	gameData.Insert(DataMap("MEPanicMod", MEPanicMod{"CF2", "Campaign Fatigue - Haggard", -4}))
	gameData.Insert(DataMap("MEPanicMod", MEPanicMod{"CF3", "Campaign Fatigue - Spent", -6}))
	gameData.Insert(DataMap("MEPanicMod", MEPanicMod{"INTER", "Interpenetrated by another ME", -2}))
	gameData.Insert(DataMap("MEPanicMod", MEPanicMod{"GOOD", "All units of ME are in good morale", 2}))

	gameData.Insert(DataMap("UnitMoraleTest", UnitMoraleTest{"OldGuard", -2}))
	gameData.Insert(DataMap("UnitMoraleTest", UnitMoraleTest{"Guard", 0}))
	gameData.Insert(DataMap("UnitMoraleTest", UnitMoraleTest{"Grenadier", 1}))
	gameData.Insert(DataMap("UnitMoraleTest", UnitMoraleTest{"Elite", 2}))
	gameData.Insert(DataMap("UnitMoraleTest", UnitMoraleTest{"CrackLine", 3}))
	gameData.Insert(DataMap("UnitMoraleTest", UnitMoraleTest{"Veteran", 4}))
	gameData.Insert(DataMap("UnitMoraleTest", UnitMoraleTest{"Regular", 5}))
	gameData.Insert(DataMap("UnitMoraleTest", UnitMoraleTest{"Conscript", 6}))
	gameData.Insert(DataMap("UnitMoraleTest", UnitMoraleTest{"Landwehr", 7}))
	gameData.Insert(DataMap("UnitMoraleTest", UnitMoraleTest{"Militia", 8}))
	gameData.Insert(DataMap("UnitMoraleTest", UnitMoraleTest{"Rabble", 10}))
	gameData.Insert(DataMap("UnitMoraleMod", UnitMoraleMod{"C1", "Light Cover", 1}))
	gameData.Insert(DataMap("UnitMoraleMod", UnitMoraleMod{"C2", "Medium Cover", 2}))
	gameData.Insert(DataMap("UnitMoraleMod", UnitMoraleMod{"C3", "Heavy Cover", 3}))
	gameData.Insert(DataMap("UnitMoraleMod", UnitMoraleMod{"C4", "Very Heavy Cover", 4}))
	gameData.Insert(DataMap("UnitMoraleMod", UnitMoraleMod{"F1", "Enfilade Fire by Infantry at Close Range", -2}))
	gameData.Insert(DataMap("UnitMoraleMod", UnitMoraleMod{"F2", "Enfilade Fire by Infantry at Point Range", -5}))
	gameData.Insert(DataMap("UnitMoraleMod", UnitMoraleMod{"F3", "Enfilade Fire by Artillery within 1 grid", -5}))
	gameData.Insert(DataMap("UnitMoraleMod", UnitMoraleMod{"DIS", "Disordered", -2}))
	gameData.Insert(DataMap("UnitMoraleMod", UnitMoraleMod{"HIT", "Per Hit currently carried by unit", -1}))
	gameData.Insert(DataMap("UnitMoraleMod", UnitMoraleMod{"GC", "Charged by Guard unit", -2}))
	gameData.Insert(DataMap("UnitMoraleMod", UnitMoraleMod{"KL", "Unformed attempting to form Klumpen", -3}))
	gameData.Insert(DataMap("UnitMoraleMod", UnitMoraleMod{"HW", "Unit in Heavy Woods", -1}))
	gameData.Insert(DataMap("UnitMoraleMod", UnitMoraleMod{"CX", "Caisson explodes in same grid", -4}))
	gameData.Insert(DataMap("UnitMoraleMod", UnitMoraleMod{"BB", "Checking vs Bombardment only", 3}))
	gameData.Insert(DataMap("UnitMoraleMod", UnitMoraleMod{"L1", "Veteran in Line", -1}))
	gameData.Insert(DataMap("UnitMoraleMod", UnitMoraleMod{"L2", "Regular in Line", -2}))
	gameData.Insert(DataMap("UnitMoraleMod", UnitMoraleMod{"L3", "Conscript or lower in Line, per base of frontage", -1}))
	gameData.Insert(DataMap("UnitMoraleMod", UnitMoraleMod{"SQ", "Unit is in square", 3}))
	gameData.Insert(DataMap("UnitMoraleMod", UnitMoraleMod{"CC", "Unit is in closed column", 1}))

	gameData.Insert(DataMap("FireDisciplineTest", FireDisciplineTest{"OldGuard", 1, -2}))
	gameData.Insert(DataMap("FireDisciplineTest", FireDisciplineTest{"Guard", 5, 2}))
	gameData.Insert(DataMap("FireDisciplineTest", FireDisciplineTest{"Grenadier", 6, 4}))
	gameData.Insert(DataMap("FireDisciplineTest", FireDisciplineTest{"Elite", 7, 5}))
	gameData.Insert(DataMap("FireDisciplineTest", FireDisciplineTest{"CrackLine", 9, 6}))
	gameData.Insert(DataMap("FireDisciplineTest", FireDisciplineTest{"Veteran", 10, 8}))
	gameData.Insert(DataMap("FireDisciplineTest", FireDisciplineTest{"Regular", 11, 7}))
	gameData.Insert(DataMap("FireDisciplineTest", FireDisciplineTest{"Conscript", 12, 8}))
	gameData.Insert(DataMap("FireDisciplineTest", FireDisciplineTest{"Landwehr", 13, 9}))
	gameData.Insert(DataMap("FireDisciplineTest", FireDisciplineTest{"Militia", 17, 13}))
	gameData.Insert(DataMap("FireDisciplineTest", FireDisciplineTest{"Rabble", 20, 16}))
	gameData.Insert(DataMap("FireDisciplineMod", FireDisciplineMod{"SK", "Per hit from SK fire this turn", -1}))
	gameData.Insert(DataMap("FireDisciplineMod", FireDisciplineMod{"HIT", "Per hit carried in total", -1}))
	gameData.Insert(DataMap("FireDisciplineMod", FireDisciplineMod{"BG", "Battalion Guns attached", -1}))

	gameData.Insert(DataMap("InitialBadMorale", InitialBadMorale{12, "Halt in reserve area. Return to good morale when ME activates Rally order", 1, true}))
	gameData.Insert(DataMap("InitialBadMorale", InitialBadMorale{9, "Done for the day, march to the rear in good order", 2, false}))
	gameData.Insert(DataMap("InitialBadMorale", InitialBadMorale{5, "Hasty retreat to the rear with stragglers", 4, false}))
	gameData.Insert(DataMap("InitialBadMorale", InitialBadMorale{0, "Abandoned the field for the day, will reform in the morning well to the rear", 6, false}))
	gameData.Insert(DataMap("InitialBadMorale", InitialBadMorale{-4, "Dispersed in panic, will reform in 2 days well to the rear", 10, false}))
	gameData.Insert(DataMap("InitialBadMorale", InitialBadMorale{-20, "Cowards !, Deserters !", 12, false}))
	gameData.Insert(DataMap("InitialBadMod", InitialBadMod{"SQP", "ME Sauve Qui Peut", -5}))
	gameData.Insert(DataMap("InitialBadMod", InitialBadMod{"HIT", "Per Hit", -1}))
	gameData.Insert(DataMap("InitialBadMod", InitialBadMod{"FTG", "Per Fatigue over Fresh", -1}))
	gameData.Insert(DataMap("InitialBadMod", InitialBadMod{"RA", "Reluctant Allies", -2}))
	gameData.Insert(DataMap("InitialBadMod", InitialBadMod{"LC", "Lost Colours", -4}))
	gameData.Insert(DataMap("InitialBadMod", InitialBadMod{"ES", "Ejected from structure or strongpoint", -4}))
	gameData.Insert(DataMap("InitialBadMod", InitialBadMod{"L1", "Charismatic Leader in same grid", 10}))
	gameData.Insert(DataMap("InitialBadMod", InitialBadMod{"L2", "Inspirational Leader in same grid", 8}))
	gameData.Insert(DataMap("InitialBadMod", InitialBadMod{"L3", "Impersonal Leader in same grid", 2}))
	gameData.Insert(DataMap("InitialBadMod", InitialBadMod{"L4", "Uninspiring Leader in same grid", 1}))
	gameData.Insert(DataMap("InitialBadMod", InitialBadMod{"CF1", "Campaign Fatigue - Rested", 1}))
	gameData.Insert(DataMap("InitialBadMod", InitialBadMod{"CF2", "Campaign Fatigue - Weary", -1}))
	gameData.Insert(DataMap("InitialBadMod", InitialBadMod{"CF3", "Campaign Fatigue - Haggard", -2}))
	gameData.Insert(DataMap("InitialBadMod", InitialBadMod{"CF4", "Campaign Fatigue - Spent", -4}))

	gameData.Insert(DataMap("BonusImpulse", BonusImpulse{19, "ME receives bonus impulse", true, false, false, false}))
	gameData.Insert(DataMap("BonusImpulse", BonusImpulse{17, "ME receives bonus impulse at the cost of 1 fatigue", true, true, false, false}))
	gameData.Insert(DataMap("BonusImpulse", BonusImpulse{14, "ME receives bonus impulse for 1 unit only", true, false, true, false}))
	gameData.Insert(DataMap("BonusImpulse", BonusImpulse{12, "Another round of firefight and streetfight", false, false, false, true}))
	gameData.Insert(DataMap("BonusImpulse", BonusImpulse{7, "No effect", false, false, false, false}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"CA", "Commander Action by Army/Corps commander", 4}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"LA", "Leader Action by ME Leader", 2}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"CAW", "Per Close Action win this turn", 3}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"CAD", "Per Close Action loss this turn", -2}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"IMP", "ME has Impetus", 2}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"FTG", "Per Fatigue level over Fresh", -1}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"MV2", "ME Moved over 1 grid to engage", -3}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"DEF", "ME is on defend orders and started the turn unengaged", -2}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"RSM", "Rain, Snow or Mud", -2}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"FOG", "Fog or Smoke in same grid", -5}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"MEH", "Per unit in the ME holding area", -1}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"INT", "ME Interpenetrated by another ME", -5}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"SHK", "ME is Shaken", -4}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"FLG", "ME took a flag", 6}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"A1", "Took enemy light battery", 3}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"A2", "Took enemy heavy battery", 4}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"TWN", "Each town block captured", 4}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"SPL", "Own strongpoint captured", 7}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"C1", "Charismatic Army Commander attached", 5}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"C2", "Inspirational Army Commander attached", 3}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"C3", "Impersonal Army Commander attached", 2}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"C4", "Charismatic Corps Commander attached", 3}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"C5", "Inspirational Corps Commander attached", 2}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"C6", "Impersonal Corps Commander attached", 1}))
	gameData.Insert(DataMap("BonusImpulseMod", BonusImpulseMod{"C7", "Bold / Superior ME Leader attached", 1}))

	gameData.Insert(DataMap("MEFatigue", MEFatigue{15, "ME incurs 1 fatigue level", false}))
	gameData.Insert(DataMap("MEFatigue", MEFatigue{11, "ME incurs 1 fatigue if not fatigued last turn", true}))
	gameData.Insert(DataMap("MEFatigueMod", MEFatigueMod{"S1", "Took Strongpoint", -3}))
	gameData.Insert(DataMap("MEFatigueMod", MEFatigueMod{"F1", "Took Enemy Standard", -2}))
	gameData.Insert(DataMap("MEFatigueMod", MEFatigueMod{"1ST", "First turn in combat for the day", -6}))
	gameData.Insert(DataMap("MEFatigueMod", MEFatigueMod{"HT", "Extreme Heat", 6}))
	gameData.Insert(DataMap("MEFatigueMod", MEFatigueMod{"MUD", "Attacking in mud", 2}))
	gameData.Insert(DataMap("MEFatigueMod", MEFatigueMod{"NL", "ME took no losses this turn", -5}))
	gameData.Insert(DataMap("MEFatigueMod", MEFatigueMod{"BB", "ME checking for bombardment only", -2}))
	gameData.Insert(DataMap("MEFatigueMod", MEFatigueMod{"LW", "Leader wounded this turn", 1}))
	gameData.Insert(DataMap("MEFatigueMod", MEFatigueMod{"LK", "Leader killed", 2}))
	gameData.Insert(DataMap("MEFatigueMod", MEFatigueMod{"CK", "Corps commander killed this turn", 4}))
	gameData.Insert(DataMap("MEFatigueMod", MEFatigueMod{"FF", "Per Firefight", 1}))
	gameData.Insert(DataMap("MEFatigueMod", MEFatigueMod{"FM", "Forced March", 4}))
	gameData.Insert(DataMap("MEFatigueMod", MEFatigueMod{"LS", "Per lost standard this turn", 1}))
	gameData.Insert(DataMap("MEFatigueMod", MEFatigueMod{"TS", "Took structure other than strongpoint", -1}))
	gameData.Insert(DataMap("MEFatigueMod", MEFatigueMod{"EC", "Extreme Cold", 2}))
	gameData.Insert(DataMap("MEFatigueMod", MEFatigueMod{"CD", "Per Close Action Defeat", 3}))
	gameData.Insert(DataMap("MEFatigueMod", MEFatigueMod{"BM", "Each morale check caused by bombardment", 2}))
	gameData.Insert(DataMap("MEFatigueMod", MEFatigueMod{"BI", "Took a 2nd Impulse", 2}))
	gameData.Insert(DataMap("MEFatigueMod", MEFatigueMod{"CF", "Per campaign fatigue level", 1}))

	gameData.Insert(DataMap("FatigueRecovery", FatigueRecovery{22, "Full Recovery", 2}))
	gameData.Insert(DataMap("FatigueRecovery", FatigueRecovery{12, "Recovery", 1}))
	gameData.Insert(DataMap("FatigueRecovery", FatigueRecovery{8, "Resting (+2 next turn)", 0}))
	gameData.Insert(DataMap("FatigueRecoveryMod", FatigueRecoveryMod{"C1", "Campaign Fresh", 2}))
	gameData.Insert(DataMap("FatigueRecoveryMod", FatigueRecoveryMod{"C2", "Campaign Haggard", -2}))
	gameData.Insert(DataMap("FatigueRecoveryMod", FatigueRecoveryMod{"C3", "Campaign Spent", -4}))
	gameData.Insert(DataMap("FatigueRecoveryMod", FatigueRecoveryMod{"BB", "Per bombardment casualty this turn", -1}))
	gameData.Insert(DataMap("FatigueRecoveryMod", FatigueRecoveryMod{"RS", "Rested last turn", 2}))

	gameData.Insert(DataMap("BadMoraleRec", BadMoraleRec{"OldGuard", 9, 2}))
	gameData.Insert(DataMap("BadMoraleRec", BadMoraleRec{"Guard", 10, 3}))
	gameData.Insert(DataMap("BadMoraleRec", BadMoraleRec{"Grenadier", 11, 4}))
	gameData.Insert(DataMap("BadMoraleRec", BadMoraleRec{"Elite", 12, 5}))
	gameData.Insert(DataMap("BadMoraleRec", BadMoraleRec{"CrackLine", 13, 6}))
	gameData.Insert(DataMap("BadMoraleRec", BadMoraleRec{"Veteran", 14, 7}))
	gameData.Insert(DataMap("BadMoraleRec", BadMoraleRec{"Regular", 15, 8}))
	gameData.Insert(DataMap("BadMoraleRec", BadMoraleRec{"Conscript", 16, 9}))
	gameData.Insert(DataMap("BadMoraleRec", BadMoraleRec{"Landwehr", 17, 10}))
	gameData.Insert(DataMap("BadMoraleRec", BadMoraleRec{"Militia", 18, 12}))
	gameData.Insert(DataMap("BadMoraleRec", BadMoraleRec{"Rabble", 19, 14}))
	gameData.Insert(DataMap("BadMoraleRecMod", BadMoraleRecMod{"CF", "Per fatigue on Cavalry ME", -2}))
	gameData.Insert(DataMap("BadMoraleRecMod", BadMoraleRecMod{"MF", "Per fatigue on Mixed/Infantry ME", -1}))
	gameData.Insert(DataMap("BadMoraleRecMod", BadMoraleRecMod{"SL", "Units standard has been lost", -3}))

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Movement, Relocation and Formation Change Tables

	gameData.Insert(DataMap("GTMove", GrandTacticalMove{"A Infantry", 26, 10, 0, 36, 45, 54}))
	gameData.Insert(DataMap("GTMove", GrandTacticalMove{"Cavalry", 30, 14, 0, 40, 50, 60}))
	gameData.Insert(DataMap("GTMove", GrandTacticalMove{"B Infantry", 22, 8, 0, 28, 35, 44}))
	gameData.Insert(DataMap("GTMove", GrandTacticalMove{"Horse Arty", 16, 8, 0, 28, 40, 50}))
	gameData.Insert(DataMap("GTMove", GrandTacticalMove{"Class I Arty", 8, 4, 0, 26, 44, 54}))
	gameData.Insert(DataMap("GTMove", GrandTacticalMove{"Class II III Arty", 6, 3, 0, 20, 34, 42}))
	gameData.Insert(DataMap("GTMove", GrandTacticalMove{"Pontoon Train", 0, 0, 0, 0, 0, 32}))
	gameData.Insert(DataMap("GTMove", GrandTacticalMove{"Corps Baggage", 26, 10, 0, 36, 45, 54}))
	gameData.Insert(DataMap("GTMove", GrandTacticalMove{"Forced March", 26, 10, 0, 10, 10, 12}))

	gameData.Insert(DataMap("Deployment", Deployment{1, 1}))
	gameData.Insert(DataMap("Deployment", Deployment{10, 2}))
	gameData.Insert(DataMap("Deployment", Deployment{16, 3}))
	gameData.Insert(DataMap("DeploymentMod", DeploymentMod{"FL", "French Line 1800-1807", 6}))
	gameData.Insert(DataMap("DeploymentMod", DeploymentMod{"OG", "French Guard", 6}))
	gameData.Insert(DataMap("DeploymentMod", DeploymentMod{"F1", "French 1792-1799", 4}))
	gameData.Insert(DataMap("DeploymentMod", DeploymentMod{"F2", "French 1808-1812", 4}))
	gameData.Insert(DataMap("DeploymentMod", DeploymentMod{"F3", "French 1815", 4}))
	gameData.Insert(DataMap("DeploymentMod", DeploymentMod{"F4", "French Conscript 1809-1812", 2}))
	gameData.Insert(DataMap("DeploymentMod", DeploymentMod{"FA", "French Allied 1807", 2}))
	gameData.Insert(DataMap("DeploymentMod", DeploymentMod{"BR", "British", 2}))
	gameData.Insert(DataMap("DeploymentMod", DeploymentMod{"AU", "Austrian 1792-1805", -2}))
	gameData.Insert(DataMap("DeploymentMod", DeploymentMod{"RU", "Russian 1792-1805", -2}))
	gameData.Insert(DataMap("DeploymentMod", DeploymentMod{"PR", "Prussian 1792-1806", -3}))
	gameData.Insert(DataMap("DeploymentMod", DeploymentMod{"ML", "Militia", -5}))
	gameData.Insert(DataMap("DeploymentMod", DeploymentMod{"MUD", "Mud", -3}))
	gameData.Insert(DataMap("DeploymentMod", DeploymentMod{"FOG", "Fog", -2}))
	gameData.Insert(DataMap("DeploymentMod", DeploymentMod{"CP", "Deploying from chokepoint", -4}))
	gameData.Insert(DataMap("DeploymentMod", DeploymentMod{"MV", "Per Grid Moved this turn", -2}))
	gameData.Insert(DataMap("DeploymentMod", DeploymentMod{"DK", "Darkness", -5}))
	gameData.Insert(DataMap("DeploymentState", DeploymentState{1, "Fully Deployed", "The ME is fully deployed", "Half move if engaged", "May react", false}))
	gameData.Insert(DataMap("DeploymentState", DeploymentState{2, "Brigade Out", "1 Bde of the ME is deployed, with the remainder behind in column", "Deploy on the head of the column", "Recoil 1 grid DIS, or stand DIS and Shaken", false}))
	gameData.Insert(DataMap("DeploymentState", DeploymentState{3, "Deploying", "Lead 2 units deployed, skirmishers out. Remainder deploying from column", "Deploy on the head of the column", "Recoil 1 grid Shaken", true}))
	gameData.Insert(DataMap("DeploymentState", DeploymentState{4, "Closed Up Campaign Column", "Baggage to the rear, units prepare to deploy", "Units deploy 2 grids behind the head of column, Shaken", "Recoil 1 grid shaken", true}))
	gameData.Insert(DataMap("DeploymentState", DeploymentState{5, "Regular Campaign Column", "Units begin to close downt the distance between them but remain in road column", "Deploy 2 grids behind the head of the column, Shaken, +1 fatigue", "Recoil 1 grid shaken", true}))
	gameData.Insert(DataMap("DeploymentState", DeploymentState{6, "Extended Campaign Column", "Units are strung out on the march and widely dispersed, some stragglers and foragers. 1 Skirmisher unit may be formed on the head of the column.", "Deploy 2 grids from the head of the column, Shaken, +2 fatigue", "Broken", true}))

	gameData.Insert(DataMap("TacMove", TacMove{"Infantry", 10, 7, 5, 5, 7, -1, -3}))
	gameData.Insert(DataMap("TacMove", TacMove{"LightCav", 20, 10, 0, 10, 15, -5, 0}))
	gameData.Insert(DataMap("TacMove", TacMove{"Cavalry", 16, 8, 0, 5, 12, -5, 0}))
	gameData.Insert(DataMap("TacMove", TacMove{"Artillery", 6, 3, 0, 2, 0, 0, 0}))

	gameData.Insert(DataMap("AdditionalMove", AdditionalMove{"Marchfeld", 1, 2, 3}))
	gameData.Insert(DataMap("AdditionalMove", AdditionalMove{"Rolling", 2, 4, 6}))
	gameData.Insert(DataMap("AdditionalMove", AdditionalMove{"Rough", 3, 6, 9}))

	// Horse Light
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Guard", "2pdr", 13, 20, 2}))
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Guard", "3pdr", 13, 20, 2}))
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Guard", "4pdr", 13, 20, 2}))
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Guard", "6pdr", 13, 20, 2}))

	// Horse Medium
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Guard", "8pdr", 10, 16, 1}))
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Guard", "9pdr", 10, 16, 1}))

	// Class I light
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class I", "2pdr", 10, 16, 2}))
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class I", "3pdr", 10, 16, 2}))
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class I", "4pdr", 10, 16, 2}))
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class I", "6pdr", 10, 16, 2}))

	// Class I medium
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class I", "8pdr", 8, 13, 1}))
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class I", "9pdr", 8, 13, 1}))

	// Class I heavy
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class I", "10pdr", 6, 10, 1}))
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class I", "12pdr", 6, 10, 1}))
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class I", "18pdr", 6, 10, 1}))
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class I", "24pdr", 6, 10, 1}))

	// Class II light
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class II", "2pdr", 8, 13, 1}))
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class II", "3pdr", 8, 13, 1}))
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class II", "4pdr", 8, 13, 1}))
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class II", "6pdr", 8, 13, 1}))

	// Class II medium
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class II", "8pdr", 6, 10, 1}))
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class II", "9pdr", 6, 10, 1}))

	// Class II heavy
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class II", "10pdr", 5, 8, 1}))
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class II", "12pdr", 5, 8, 1}))
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class II", "18pdr", 5, 8, 1}))
	gameData.Insert(DataMap("ArtyMove", ArtyMove{"Class II", "24pdr", 5, 8, 1}))

	gameData.Insert(DataMap("ArtyHorseLoss", ArtyHorseLoss{"Marchfeld", 4}))
	gameData.Insert(DataMap("ArtyHorseLoss", ArtyHorseLoss{"Rolling", 6}))
	gameData.Insert(DataMap("ArtyHorseLoss", ArtyHorseLoss{"Rough", 8}))

	gameData.Insert(DataMap("BUAMove", BUAMove{"Excellent", 6, 4}))
	gameData.Insert(DataMap("BUAMove", BUAMove{"Good", 8, 5}))
	gameData.Insert(DataMap("BUAMove", BUAMove{"Average", 12, 6}))
	gameData.Insert(DataMap("BUAMove", BUAMove{"Poor", 14, 7}))

	gameData.Insert(DataMap("BUAMod", BUAMod{"CA", "Successful Commander Action", 2}))
	gameData.Insert(DataMap("BUAMod", BUAMod{"LA", "Successful Leader Action", 1}))
	gameData.Insert(DataMap("BUAMod", BUAMod{"UN", "Per unit that moved through the structure this turn", -2}))
	gameData.Insert(DataMap("BUAMod", BUAMod{"JN", "Janissaries", 2}))
	gameData.Insert(DataMap("BUAMod", BUAMod{"RU", "Russian Grenadiers", 3}))
	gameData.Insert(DataMap("BUAMod", BUAMod{"AU", "Austrian Grenadiers", 3}))

	gameData.Insert(DataMap("SKRelocate", SKRelocate{"Superior", 4, 8, 12}))
	gameData.Insert(DataMap("SKRelocate", SKRelocate{"Excellent", 5, 9, 13}))
	gameData.Insert(DataMap("SKRelocate", SKRelocate{"Good", 7, 11, 15}))
	gameData.Insert(DataMap("SKRelocate", SKRelocate{"Average", 8, 13, 16}))
	gameData.Insert(DataMap("SKRelocate", SKRelocate{"Poor", 11, 17, 23}))

	gameData.Insert(DataMap("SKRelocateMod", SKRelocateMod{"AM", "Ammo depleted", -6}))
	gameData.Insert(DataMap("SKRelocateMod", SKRelocateMod{"BD", "Unit is Bold", 6}))
	gameData.Insert(DataMap("SKRelocateMod", SKRelocateMod{"CL", "Charismatic Leader Attached", 6}))
	gameData.Insert(DataMap("SKRelocateMod", SKRelocateMod{"IL", "Inspirational Leader Attached", 4}))
	gameData.Insert(DataMap("SKRelocateMod", SKRelocateMod{"AL", "Average Leader Attached", 2}))
	gameData.Insert(DataMap("SKRelocateMod", SKRelocateMod{"UL", "Uninspiring Leader Attached", -1}))

	gameData.Insert(DataMap("SKSupport", SKSupport{"Normal", 9, 6, 12}))
	gameData.Insert(DataMap("SKSupport", SKSupport{"Bold", 30, 30, 18}))

	gameData.Insert(DataMap("ArtyRelocate", ArtyRelocate{"Guard", 6, 8, 10, 11, 13, 16, 21, 2, 2, 7, 10, 12, 15, 20}))
	gameData.Insert(DataMap("ArtyRelocate", ArtyRelocate{"Class I", 8, 9, 11, 12, 14, 17, 21, 2, 5, 8, 11, 13, 15, 20}))
	gameData.Insert(DataMap("ArtyRelocate", ArtyRelocate{"Class II", 9, 11, 12, 13, 15, 18, 22, 6, 7, 9, 12, 14, 16, 20}))
	gameData.Insert(DataMap("ArtyRelocate", ArtyRelocate{"Class III", 12, 13, 14, 15, 17, 19, 23, 7, 8, 10, 13, 15, 17, 20}))

	gameData.Insert(DataMap("ArtyRelocateMod", ArtyRelocateMod{"LA", "ME Commander Attached", 1}))
	gameData.Insert(DataMap("ArtyRelocateMod", ArtyRelocateMod{"CA", "Corps Commander Attached", 2}))
	gameData.Insert(DataMap("ArtyRelocateMod", ArtyRelocateMod{"AA", "Army Commander Attached", 3}))
	gameData.Insert(DataMap("ArtyRelocateMod", ArtyRelocateMod{"MD", "Mud", -2}))
	gameData.Insert(DataMap("ArtyRelocateMod", ArtyRelocateMod{"FT", "Per fatigue level", -1}))
	gameData.Insert(DataMap("ArtyRelocateMod", ArtyRelocateMod{"AT", "Per Attempt", 1}))

	// Result codes :
	// N = number of movement inches after formation change
	// 0 = Not allowed
	// -1 = entire turn, and may fire
	// -2 = entire turn
	// -3 = entire turn and disordered
	gameData.Insert(DataMap("FormationChange", FormationChange{"AR", "AttackColumn", "Line Left", 2, -1}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"AR", "AttackColumn", "Line Centre", -1, -3}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"AR", "AttackColumn", "Line Right", -1, -3}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"AR", "AttackColumn", "ClosedColumn", 6, 4}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"AR", "AttackColumn", "Square", 1, 0}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"AR", "ClosedColumn", "Line Left", -1, -2}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"AR", "ClosedColumn", "Line Centre", 0, 0}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"AR", "ClosedColumn", "Line Right", 0, 0}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"AR", "ClosedColumn", "AttackColumn", 6, 4}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"AR", "ClosedColumn", "Square", -2, -3}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"AR", "Line", "AttackColumn", 1, -2}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"AR", "Line", "ClosedColumn", -1, -2}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"AR", "Line", "Square", -1, -3}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"AR", "Square", "Line", 2, -1}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"AR", "Square", "AttackColumn", 2, -1}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"AR", "Square", "ClosedColumn", -2, -3}))

	gameData.Insert(DataMap("FormationChange", FormationChange{"DIV", "AttackColumn", "Line Left", 2, -1}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"DIV", "AttackColumn", "Line Centre", 1, -1}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"DIV", "AttackColumn", "Line Right", 2, -1}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"DIV", "AttackColumn", "ClosedColumn", 10, 6}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"DIV", "AttackColumn", "Square", 2, -1}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"DIV", "ClosedColumn", "Line Left", -1, -2}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"DIV", "ClosedColumn", "Line Centre", 3, 2}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"DIV", "ClosedColumn", "Line Right", -1, -2}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"DIV", "ClosedColumn", "AttackColumn", 10, 8}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"DIV", "ClosedColumn", "Square", -1, -2}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"DIV", "Line", "AttackColumn", 1, -2}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"DIV", "Line", "ClosedColumn", -1, -2}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"DIV", "Line", "Square", 1, -3}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"DIV", "Square", "Line", 2, -1}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"DIV", "Square", "AttackColumn", 4, 2}))
	gameData.Insert(DataMap("FormationChange", FormationChange{"DIV", "Square", "ClosedColumn", -2, -3}))

	gameData.Insert(DataMap("FormSquare", FormSquare{"OldGuard", "AttackColumn", 14, 6, 0, -4}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"OldGuard", "ClosedColumn", 15, 10, 8, 6}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"OldGuard", "Line", 17, 13, 11, 10}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Guard", "AttackColumn", 14, 6, 0, -4}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Guard", "ClosedColumn", 15, 10, 8, 6}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Guard", "Line", 17, 13, 11, 10}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Grenadier", "AttackColumn", 14, 6, 0, -4}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Grenadier", "ClosedColumn", 15, 10, 8, 6}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Grenadier", "Line", 17, 13, 11, 10}))

	gameData.Insert(DataMap("FormSquare", FormSquare{"Elite", "AttackColumn", 15, 10, 8, 2}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Elite", "ClosedColumn", 17, 12, 10, 8}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Elite", "Line", 22, 14, 13, 11}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"CrackLine", "AttackColumn", 15, 10, 8, 2}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"CrackLine", "ClosedColumn", 17, 12, 10, 8}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"CrackLine", "Line", 22, 14, 13, 11}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Veteran", "AttackColumn", 15, 10, 8, 2}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Veteran", "ClosedColumn", 17, 12, 10, 8}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Veteran", "Line", 22, 14, 13, 11}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Regular", "AttackColumn", 15, 10, 8, 2}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Regular", "ClosedColumn", 17, 12, 10, 8}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Regular", "Line", 22, 14, 13, 11}))

	gameData.Insert(DataMap("FormSquare", FormSquare{"Conscript", "AttackColumn", 17, 14, 12, 11}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Conscript", "ClosedColumn", 21, 16, 14, 12}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Conscript", "Line", 26, 17, 15, 14}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Landwehr", "AttackColumn", 17, 14, 12, 11}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Landwehr", "ClosedColumn", 21, 16, 14, 12}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Landwehr", "Line", 26, 17, 15, 14}))

	gameData.Insert(DataMap("FormSquare", FormSquare{"Militia", "AttackColumn", 21, 18, 16, 15}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Militia", "ClosedColumn", 26, 22, 18, 17}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Militia", "Line", 28, 26, 19, 18}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Rabble", "AttackColumn", 21, 18, 16, 15}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Rabble", "ClosedColumn", 26, 22, 18, 17}))
	gameData.Insert(DataMap("FormSquare", FormSquare{"Rabble", "Line", 28, 26, 19, 18}))

	gameData.Insert(DataMap("FormSquareMod", FormSquareMod{"D1", "Charge from Diagonal", -2}))
	gameData.Insert(DataMap("FormSquareMod", FormSquareMod{"D2", "Charged from Flank", -6}))
	gameData.Insert(DataMap("FormSquareMod", FormSquareMod{"D3", "Charged from Rear", -8}))
	gameData.Insert(DataMap("FormSquareMod", FormSquareMod{"DS", "Disordered", -4}))
	gameData.Insert(DataMap("FormSquareMod", FormSquareMod{"CA", "Commander Attached", 4}))
	gameData.Insert(DataMap("FormSquareMod", FormSquareMod{"LA", "Leader Attached", 2}))
	gameData.Insert(DataMap("FormSquareMod", FormSquareMod{"OC", "Opportunity Charge", -3}))

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Muskets and Cannons and things that go Bang

	gameData.Insert(DataMap("FireEffect", FireEffect{1, -5, "Terrible"}))
	gameData.Insert(DataMap("FireEffect", FireEffect{2, 1, "Erratic"}))
	gameData.Insert(DataMap("FireEffect", FireEffect{3, 5, "Ineffective"}))
	gameData.Insert(DataMap("FireEffect", FireEffect{4, 9, "Effective"}))
	gameData.Insert(DataMap("FireEffect", FireEffect{5, 12, "Deadly"}))
	gameData.Insert(DataMap("FireEffect", FireEffect{6, 15, "Withering"}))
	gameData.Insert(DataMap("FireEffect", FireEffect{7, 19, "Crushing"}))
	gameData.Insert(DataMap("FireEffect", FireEffect{8, 23, "Devestating"}))
	gameData.Insert(DataMap("FireEffect", FireEffect{9, 29, "Magnificent"}))
	gameData.Insert(DataMap("FireEffect", FireEffect{10, 34, "Extraordinaire"}))

	gameData.Insert(DataMap("FireChart", FireChart{1, 3, 0, 0, 1, 1}))
	gameData.Insert(DataMap("FireChart", FireChart{2, 5, 0, 1, 2, 2}))
	gameData.Insert(DataMap("FireChart", FireChart{3, 8, 1, 1, 2, 3}))
	gameData.Insert(DataMap("FireChart", FireChart{4, 12, 2, 2, 3, 4}))
	gameData.Insert(DataMap("FireChart", FireChart{5, 15, 2, 3, 4, 5}))
	gameData.Insert(DataMap("FireChart", FireChart{6, 18, 4, 4, 5, 6}))
	gameData.Insert(DataMap("FireChart", FireChart{7, 22, 4, 6, 7, 8}))
	gameData.Insert(DataMap("FireChart", FireChart{8, 26, 6, 8, 8, 10}))
	gameData.Insert(DataMap("FireChart", FireChart{9, 35, 7, 10, 10, 12}))
	gameData.Insert(DataMap("FireChart", FireChart{10, 42, 8, 11, 12, 14}))

	gameData.Insert(DataMap("FireMod", FireMod{"OldGuard", "OldGuard", 10}))
	gameData.Insert(DataMap("FireMod", FireMod{"Guard", "Guard", 8}))
	gameData.Insert(DataMap("FireMod", FireMod{"Grenadier", "Grenadier", 6}))
	gameData.Insert(DataMap("FireMod", FireMod{"Elite", "Elite", 4}))
	gameData.Insert(DataMap("FireMod", FireMod{"CrackLine", "CrackLine", 2}))
	gameData.Insert(DataMap("FireMod", FireMod{"Veteran", "Veteran", 0}))
	gameData.Insert(DataMap("FireMod", FireMod{"Regular", "Regular", -1}))
	gameData.Insert(DataMap("FireMod", FireMod{"Conscript", "Conscript", -2}))
	gameData.Insert(DataMap("FireMod", FireMod{"Landwehr", "Landwehr", -3}))
	gameData.Insert(DataMap("FireMod", FireMod{"Militia", "Militia", -4}))
	gameData.Insert(DataMap("FireMod", FireMod{"Rabble", "Rabble", -6}))
	gameData.Insert(DataMap("FireMod", FireMod{"FF", "First Fire", 3}))
	gameData.Insert(DataMap("FireMod", FireMod{"DIS", "Disordered", -5}))
	gameData.Insert(DataMap("FireMod", FireMod{"SHK", "Shaken", -5}))
	gameData.Insert(DataMap("FireMod", FireMod{"SQ", "Firing from Square", -4}))
	gameData.Insert(DataMap("FireMod", FireMod{"AMD", "Ammo Depleted", -6}))
	gameData.Insert(DataMap("FireMod", FireMod{"AME", "Ammo Exhausted", -10}))
	gameData.Insert(DataMap("FireMod", FireMod{"FLW", "Firing from Lt Woods", -1}))
	gameData.Insert(DataMap("FireMod", FireMod{"FMW", "Firing from Md Woods", -2}))
	gameData.Insert(DataMap("FireMod", FireMod{"FHW", "Firing from Hv Woods", -3}))
	gameData.Insert(DataMap("FireMod", FireMod{"ENL", "Enfilade vs Line", 5}))
	gameData.Insert(DataMap("FireMod", FireMod{"ENC", "Enfilade vs any Column", 5}))
	gameData.Insert(DataMap("FireMod", FireMod{"TSQ", "vs Square", 9}))
	gameData.Insert(DataMap("FireMod", FireMod{"ART", "vs Artillery", -10}))
	gameData.Insert(DataMap("FireMod", FireMod{"TCOL", "vs Column", 2}))
	gameData.Insert(DataMap("FireMod", FireMod{"TCC", "vs Closed Column", 5}))
	gameData.Insert(DataMap("FireMod", FireMod{"OO", "vs Open Order", -6}))
	gameData.Insert(DataMap("FireMod", FireMod{"SK", "vs Skirmish Order", -10}))
	gameData.Insert(DataMap("FireMod", FireMod{"OPP", "Opportunity Fire", -3}))
	gameData.Insert(DataMap("FireMod", FireMod{"CAV", "vs Cavalry", -8}))
	gameData.Insert(DataMap("FireMod", FireMod{"CAVM", "vs Cavalry charging uphill or in mud", -5}))
	gameData.Insert(DataMap("FireMod", FireMod{"RN", "Rain", -2}))
	gameData.Insert(DataMap("FireMod", FireMod{"HR", "Heavy Rain", -5}))
	gameData.Insert(DataMap("FireMod", FireMod{"C1", "Target in Lt Cover", -4}))
	gameData.Insert(DataMap("FireMod", FireMod{"C2", "Target in Md Cover", -8}))
	gameData.Insert(DataMap("FireMod", FireMod{"C3", "Target in Hv Cover", -12}))
	gameData.Insert(DataMap("FireMod", FireMod{"C1", "Target in SuperHv Cover", -16}))
	gameData.Insert(DataMap("FireMod", FireMod{"C1", "Target in SuperHv Cover", -16}))
	gameData.Insert(DataMap("FireMod", FireMod{"FTG", "Per Fatigue Level over Fresh", -2}))
	gameData.Insert(DataMap("FireMod", FireMod{"HIT", "Per Hit", -1}))

	gameData.Insert(DataMap("FireSKMod", FireSKMod{"1", "Superior Grade", 12}))
	gameData.Insert(DataMap("FireSKMod", FireSKMod{"2", "Excellent Grade", 8}))
	gameData.Insert(DataMap("FireSKMod", FireSKMod{"3", "Good Grade", 5}))
	gameData.Insert(DataMap("FireSKMod", FireSKMod{"4", "Average Grade", 0}))
	gameData.Insert(DataMap("FireSKMod", FireSKMod{"5", "Poor Grade", -5}))
	gameData.Insert(DataMap("FireSKMod", FireSKMod{"FTG", "Per Fatigue Level", -2}))
	gameData.Insert(DataMap("FireSKMod", FireSKMod{"CV", "Firing from Cover/Rough/Woods", 5}))
	gameData.Insert(DataMap("FireSKMod", FireSKMod{"SK", "Firing from Skirmish Order", 5}))

	gameData.Insert(DataMap("SKEffect", SKEffect{"T1", 2, "Column or Square"}))
	gameData.Insert(DataMap("SKEffect", SKEffect{"T2", 3, "Any Formed"}))
	gameData.Insert(DataMap("SKEffect", SKEffect{"T3", 4, "Artillery"}))
	gameData.Insert(DataMap("SKEffect", SKEffect{"T4", 5, "Open Order"}))
	gameData.Insert(DataMap("SKEffect", SKEffect{"T5", 6, "Skirmish Order"}))
	gameData.Insert(DataMap("SKEffect", SKEffect{"C1", 2, "Rolling/Rough Terrain"}))
	gameData.Insert(DataMap("SKEffect", SKEffect{"C2", 3, "Light Cover"}))
	gameData.Insert(DataMap("SKEffect", SKEffect{"C3", 4, "Medium Cover"}))
	gameData.Insert(DataMap("SKEffect", SKEffect{"C4", 5, "Heavy Cover"}))
	gameData.Insert(DataMap("SKEffect", SKEffect{"C5", 6, "Super Heavy Cover"}))

	gameData.Insert(DataMap("FireFight", FireFight{2, "Firefight Continues", false, true, false, false}))
	gameData.Insert(DataMap("FireFight", FireFight{13, "Fallback in good order if not in cover", true, true, false, false}))
	gameData.Insert(DataMap("FireFight", FireFight{17, "Fallback in good order", true, false, false, false}))
	gameData.Insert(DataMap("FireFight", FireFight{20, "Fallback in disorder", true, false, true, false}))
	gameData.Insert(DataMap("FireFight", FireFight{23, "Rout", true, false, true, true}))

	gameData.Insert(DataMap("FireFightMod", FireFightMod{"HIT", "Per Hit", -1}))
	gameData.Insert(DataMap("FireFightMod", FireFightMod{"NHIT", "Per Hit this turn", -2}))
	gameData.Insert(DataMap("FireFightMod", FireFightMod{"HITX", "Per Hit inflicted this turn", 2}))
	gameData.Insert(DataMap("FireFightMod", FireFightMod{"LCH", "Loser has Charismatic leader", -3}))
	gameData.Insert(DataMap("FireFightMod", FireFightMod{"LIN", "Loser has Inspirational leader", -2}))
	gameData.Insert(DataMap("FireFightMod", FireFightMod{"LAV", "Loser has Average/Poor leader", -1}))
	gameData.Insert(DataMap("FireFightMod", FireFightMod{"AMM", "Loser has Ammo depleted or exhausted", 3}))

	gameData.Insert(DataMap("ArtRange", ArtRange{"Heavy", 1, 2, 6}))
	gameData.Insert(DataMap("ArtRange", ArtRange{"MdHeavy", 1, 2, 5}))
	gameData.Insert(DataMap("ArtRange", ArtRange{"Medium", 1, 2, 4}))
	gameData.Insert(DataMap("ArtRange", ArtRange{"Light", 1, 2, 3}))

	gameData.Insert(DataMap("ArtMod", ArtMod{"MO", "Mixed Order Target", 1}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"SQ", "Target in Square", 9}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"RS", "Reverse Slope", -4}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"CAV", "Cavalry Target", 1}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"CW", "Cavalry Withdrawing", -6}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"CC", "Closed Column", 5}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"CO", "Other Columns", 2}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"OO", "Open Order", -6}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"G0", "Guard Artillery", 6}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"G1", "Class I", 3}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"G2", "Class II", 0}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"G3", "Class III", -4}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"RM", "Range - Medium", -3}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"RL", "Range - Long", -13}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"SC", "Screened by Skirmishers", -10}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"NAP", "Napoleon Himself attached to battery", 6}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"CR", "Corps Commander attached to battery", 3}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"RE", "Fresh out of Reserve", 10}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"FT", "Per Fatigue level beyond Fresh", -1}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"HC", "Heavy Counter Bty at Long Range", -3}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"LC", "Light/Medium Counter Bty at Long Range", -5}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"A1", "Ammo Depleted", -8}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"A2", "Ammo Exhausted", -13}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"T1", "Narchfeld Terrain", 3}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"3G", "3 Gun Section", 5}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"S1", "Defense of Bty vs Infantry", 5}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"S2", "Fire & Retire vs Infantry", -3}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"S3", "Followup Defense of Bty vs Infantry", -4}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"S4", "Followup Fire & Retire vs Infantry", -10}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"S5", "Support Fire vs Infantry", -2}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"S6", "Followup Support Fire vs Infantry", -6}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"S7", "Defense of Bty vs Cavalry", 4}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"S8", "Fire & Retire vs Cavalry", -5}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"S9", "Followup Defense of Bty vs Cavalry", -7}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"S10", "Followup Fire & Retire vs Cavalry", -13}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"S11", "Support Fire vs Cavalry", -3}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"S12", "Followup Support Fire vs Cavlary", -10}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"B1", "Bombardment - Grand Bty with rated leader", 5}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"B2", "Bombardment - Grand Bty without rated leader", 3}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"B3", "Bombardment - Short Range Fire", 3}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"B4", "Bombardment - Paced Rate of Fire", -4}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"B5", "Bombardment - Intensive Rate of Fire", 6}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"B5", "Bombardment - Target Moved 1 Grid", -7}))
	gameData.Insert(DataMap("ArtMod", ArtMod{"B5", "Bombardment - Target Moved more than 1 Grid", -10}))

	gameData.Insert(DataMap("ArtLimber", ArtLimber{"IC", "Infantry Close", 12}))
	gameData.Insert(DataMap("ArtLimber", ArtLimber{"IS", "Infantry Same Grid", 8}))
	gameData.Insert(DataMap("ArtLimber", ArtLimber{"IN", "Infantry Next Grid", 5}))
	gameData.Insert(DataMap("ArtLimber", ArtLimber{"CC", "Cavalry Close", 18}))
	gameData.Insert(DataMap("ArtLimber", ArtLimber{"CS", "Cavalry Same Grid", 12}))
	gameData.Insert(DataMap("ArtLimber", ArtLimber{"CN", "Cavalry Next Grid", 8}))

	gameData.Insert(DataMap("CounterBty", CounterBty{2, 2, 0, 3, 0, false}))
	gameData.Insert(DataMap("CounterBty", CounterBty{9, 1, 0, 1, 1, false}))
	gameData.Insert(DataMap("CounterBty", CounterBty{12, 0, 1, 0, 2, false}))
	gameData.Insert(DataMap("CounterBty", CounterBty{14, 1, 1, 2, 1, false}))
	gameData.Insert(DataMap("CounterBty", CounterBty{19, 1, 1, 2, 3, true}))

	gameData.Insert(DataMap("Aflame", Aflame{"Light", 17, 16, 16, 16}))
	gameData.Insert(DataMap("Aflame", Aflame{"Medium", 18, 17, 17, 17}))
	gameData.Insert(DataMap("Aflame", Aflame{"Heavy", 20, 18, 19, 18}))
	gameData.Insert(DataMap("Aflame", Aflame{"SuperHeavy", 21, 19, 20, 19}))

	gameData.Insert(DataMap("DoubleTeam", DoubleTeam{11}))
	gameData.Insert(DataMap("DoubleTeamMod", DoubleTeamMod{"British", -2}))
	gameData.Insert(DataMap("DoubleTeamMod", DoubleTeamMod{"French", 0}))
	gameData.Insert(DataMap("DoubleTeamMod", DoubleTeamMod{"French Guard", 4}))
	gameData.Insert(DataMap("DoubleTeamMod", DoubleTeamMod{"Horse Arty", 2}))
	gameData.Insert(DataMap("DoubleTeamMod", DoubleTeamMod{"Other", -3}))

	gameData.Insert(DataMap("ArtFate", ArtFate{"Friendly", 15}))
	gameData.Insert(DataMap("ArtFate", ArtFate{"Disable Enemy", 15}))
	gameData.Insert(DataMap("ArtFate", ArtFate{"Capture Enemy", 17}))
	gameData.Insert(DataMap("ArtFateMod", ArtFateMod{"CA", "Commander Action Used", 3}))
	gameData.Insert(DataMap("ArtFateMod", ArtFateMod{"NE", "Non Engaged", 2}))
	gameData.Insert(DataMap("ArtFateMod", ArtFateMod{"EN", "Engineers present", 3}))

	gameData.Insert(DataMap("Bouncethru", Bouncethru{5, 0, 0, 1, 2}))
	gameData.Insert(DataMap("Bouncethru", Bouncethru{9, 0, 1, 2, 3}))
	gameData.Insert(DataMap("Bouncethru", Bouncethru{14, 1, 2, 3, 4}))
	gameData.Insert(DataMap("Bouncethru", Bouncethru{18, 2, 3, 4, 5}))
	gameData.Insert(DataMap("Bouncethru", Bouncethru{23, 3, 4, 5, 6}))
	gameData.Insert(DataMap("BouncethruMod", BouncethruMod{"RF", "Rocky Fields", 2}))
	gameData.Insert(DataMap("BouncethruMod", BouncethruMod{"MF", "Marchfeld", 2}))
	gameData.Insert(DataMap("BouncethruMod", BouncethruMod{"LA", "Limbered Arty", 1}))
	gameData.Insert(DataMap("BouncethruMod", BouncethruMod{"DT", "Dense Target", 6}))
	gameData.Insert(DataMap("BouncethruMod", BouncethruMod{"B1", "Intensive Bombardment", 6}))
	gameData.Insert(DataMap("BouncethruMod", BouncethruMod{"HC", "Heavy Cover", -18}))
	gameData.Insert(DataMap("BouncethruMod", BouncethruMod{"RG", "Rough Ground", -2}))
	gameData.Insert(DataMap("BouncethruMod", BouncethruMod{"B2", "Paced Bombardment", -4}))
	gameData.Insert(DataMap("BouncethruMod", BouncethruMod{"TW", "Target in Town Block", -10}))
	gameData.Insert(DataMap("BouncethruMod", BouncethruMod{"MD", "Muddy conditions", -10}))
	gameData.Insert(DataMap("BouncethruMod", BouncethruMod{"LW", "Lt woods", -2}))
	gameData.Insert(DataMap("BouncethruMod", BouncethruMod{"CT", "Per Contour Difference", -2}))

	gameData.Insert(DataMap("DefFire", DefFire{1, "Infantry", "-1", "-2 FG", "-1 DFG", "-2 DEFG", "X"}))
	gameData.Insert(DataMap("DefFire", DefFire{2, "Infantry w/Impetus", "-1", "-2 G", "-1 DG", "-2 DG", "X"}))
	gameData.Insert(DataMap("DefFire", DefFire{3, "Infantry Storming", "", "G", "X", "X", "X"}))
	gameData.Insert(DataMap("DefFire", DefFire{4, "Shock Infantry", "-1", "-1 H", "DF", "DFG", "DER"}))
	gameData.Insert(DataMap("DefFire", DefFire{5, "Shock Infantry w/Impetus", "-1", "-1", "D", "-2 DG", "DER"}))
	gameData.Insert(DataMap("DefFire", DefFire{6, "Shock Infantry Storming", "", "", "X", "X", "X"}))
	gameData.Insert(DataMap("DefFire", DefFire{7, "Cavalry", "-1", "C", "DC", "DC", "X"}))
	gameData.Insert(DataMap("DefFire", DefFire{8, "Cavalry w/Impetus", "-1", "-1 D", "-2 D", "-3 C", "X"}))
	gameData.Insert(DataMap("DefFire", DefFire{9, "DUB Cavalry", "C", "DC", "A", "B", "X"}))
	gameData.Insert(DataMap("DefFire", DefFire{10, "DUB Cavalry w/Impetus", "-2", "-3 D", "A", "B", "X"}))
	gameData.Insert(DataMap("DefFire", DefFire{11, "Cossacks", "R", "DR", "DZ", "X", "X"}))

	gameData.Insert(DataMap("DefFireNote", DefFireNote{"-1", "Close Action at -1"}))
	gameData.Insert(DataMap("DefFireNote", DefFireNote{"-2", "Close Action at -2"}))
	gameData.Insert(DataMap("DefFireNote", DefFireNote{"-3", "Close Action at -3"}))
	gameData.Insert(DataMap("DefFireNote", DefFireNote{"D", "Disorders. If already disordered, Fallback"}))
	gameData.Insert(DataMap("DefFireNote", DefFireNote{"E", "If already disordered, Rout"}))
	gameData.Insert(DataMap("DefFireNote", DefFireNote{"F", "Fallback and fire"}))
	gameData.Insert(DataMap("DefFireNote", DefFireNote{"G", "If attacking guns, fire disordered and fallback"}))
	gameData.Insert(DataMap("DefFireNote", DefFireNote{"X", "Unit Routs (Inf 2 grid, Cav 3 grid), add 1-3 hits"}))
	gameData.Insert(DataMap("DefFireNote", DefFireNote{"R", "Retreat 1 Grid"}))
	gameData.Insert(DataMap("DefFireNote", DefFireNote{"Z", "Retreat 2 Grids"}))
	gameData.Insert(DataMap("DefFireNote", DefFireNote{"H", "Halt at close range, enter FireFight"}))
	gameData.Insert(DataMap("DefFireNote", DefFireNote{"C", "Cavalry, Fire and Retire 2 Grids"}))
	gameData.Insert(DataMap("DefFireNote", DefFireNote{"A", "Fire and Retire 2 Grids. 9 or less = Rout"}))
	gameData.Insert(DataMap("DefFireNote", DefFireNote{"B", "Fire and Retire 2 Grids. 12 or less = Rout"}))

	// Now create some indexes
	log.Println("Creating Index on Type")
	if err := gameData.Index([]string{"Type"}); err != nil {
		panic(err)
	}
}
