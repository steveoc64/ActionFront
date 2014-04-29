package simulation

import (
	"fmt"
	"github.com/steveoc64/ActionFront/dice"
	"github.com/steveoc64/ActionFront/list"
	"github.com/steveoc64/tiedot/db"
	"log"
)

// Validator for Corps Orders
func CorpsOrder(col *db.Col, params map[string]interface{}) map[string]interface{} {

	CorpsOrder := params["CorpsOrder"].(string)

	myCorpsOrder := list.Lookup(col, "CorpsOrder", "Order")[CorpsOrder]

	params["Stipulation"] = myCorpsOrder["Stipulation"]
	params["MEOrders"] = myCorpsOrder["MEOrders"]

	var defCount = 0
	var attCount = 0
	var rgCount = 0
	var marchCount = 0
	var sptCount = 0

	// Check that ALL ME's have a valid order
	NumME := int(params["NumME"].(float64))
	for i := 0; i < NumME; i++ {
		MEOrder := params[fmt.Sprint("ME", i+1)].(string)
		switch MEOrder {
		case "March":
			marchCount++
		case "Defend":
			defCount++
		case "Screen", "Support", "Intercept":
			sptCount++
		case "Attack":
			attCount++
		case "Bombard":
			attCount++
			sptCount++
		case "BreakOff", "Rearguard":
			rgCount++
		}
		validOrder := false
		for _, order := range myCorpsOrder["MEOrders"].([]interface{}) {
			if order == MEOrder {
				validOrder = true
			}
		}
		if !validOrder {
			params["Result"] = fmt.Sprintf("ME%d does not have a valid order", i+1)
			params["ResultAccept"] = false
			return params
		}
	}

	params["Result"] = "Good Job Sir, Everything is in Order here !"
	params["ResultAccept"] = true

	switch CorpsOrder {
	case "Manoeuvre":
		if marchCount < 1 {
			params["Result"] = "There is some confusion over our destination, Sir !"
			params["ResultAccept"] = false
			return params
		}
		if sptCount < 1 {
			params["Result"] = "May I suggest we operate a Screen or Flank Support for the march, Sir ?"
			params["ResultAccept"] = true
		}
	case "Attack":
		if attCount < 1 {
			params["Result"] = "Sir, We Must Attack !"
			params["ResultAccept"] = false
			return params
		}
		if sptCount < 1 {
			params["Result"] = "Should we allocate anyone to Support the Attack, Sir ?"
			params["ResultAccept"] = true
		}
		if attCount > 2 {
			params["Result"] = "Jolly Good Show, Sir ..  Such an agressive move !"
			params["ResultAccept"] = true
		}
	case "Engaged":
		if attCount < 1 && defCount < 1 {
			params["Result"] = "Sir, We must either Defend our position or push the Attack forward !"
			params["ResultAccept"] = false
			return params
		}
		if attCount > 2 {
			params["Result"] = "Jolly Good Show, Sir ..  Victory, or Death !"
			params["ResultAccept"] = true
		}
	case "Defend":
		// At leat 1 ME must have a defend order
		if defCount < 1 {
			params["Result"] = "Sir, we Must Defend our position !"
			params["ResultAccept"] = false
			return params
		}
		if sptCount < 1 {
			params["Result"] = "Should we allocate anyone to Support the Line of Defence, Sir ?"
			params["ResultAccept"] = true
		}
		if defCount > 2 {
			params["Result"] = "Jolly Good Show, Sir ..  They wont get through that lot !"
			params["ResultAccept"] = true
		}
	case "Withdraw":
		if defCount < 1 && rgCount < 1 {
			params["Result"] = "Sir, we really should consider covering our withdrawal with a rearguard at the very least."
			params["ResultAccept"] = false
			return params
		}
	}

	return params
}

// Simple function to test whether a slice of strings contains a given string
func stringSliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Validator for ME Orders in different situations
func MEOrder(col *db.Col, params map[string]interface{}) map[string]interface{} {

	CorpsOrder := params["CorpsOrder"].(string)
	MEOrder := params["MEOrder"].(string)

	myCorpsOrder := list.Lookup(col, "CorpsOrder", "Order")[CorpsOrder]
	MEOrders := list.Lookup(col, "MEOrder", "Order")

	params["Stipulation"] = myCorpsOrder["Stipulation"]
	params["MEOrders"] = myCorpsOrder["MEOrders"]
	params["Purpose"] = ""
	params["Notes"] = ""

	isEngaged := params["Engaged"].(float64) == 1

	orders := make([]string, 0)
	hasDefend := false
	switch params["METype"] {
	case "Cavalry":
		for _, order := range myCorpsOrder["MEOrders"].([]interface{}) {
			log.Println(order)
			o := order.(string)
			if isEngaged {
				if MEOrders[o]["IfEngaged"].(bool) {
					orders = append(orders, o)
					if o == "Defend" {
						hasDefend = true
					}
				}
			} else {
				if MEOrders[o]["IfNonEngaged"].(bool) {
					orders = append(orders, o)
					if o == "Defend" {
						hasDefend = true
					}
				}
			}
		}

	case "Infantry":
		for _, order := range myCorpsOrder["MEOrders"].([]interface{}) {
			log.Println(order)
			o := order.(string)
			if !MEOrders[o]["CavalryOnly"].(bool) {
				if isEngaged {
					if MEOrders[o]["IfEngaged"].(bool) {
						orders = append(orders, o)
						if o == "Defend" {
							hasDefend = true
						}
					}
				} else {
					if MEOrders[o]["IfNonEngaged"].(bool) {
						orders = append(orders, o)
						if o == "Defend" {
							hasDefend = true
						}
					}
				}
			}
		}
	}
	if isEngaged && !hasDefend {
		// Add a defend order if not already there
		orders = append(orders, "Defend")
	}
	params["MEOrders"] = orders

	if MEOrder != "" {
		myMEOrder := MEOrders[MEOrder]
		if stringSliceContains(orders, MEOrder) {
			params["Purpose"] = myMEOrder["Purpose"]
			params["Notes"] = myMEOrder["Notes"]
			params["ResultDefend"] = myMEOrder["DefendIfEngaged"]
			params["ResultShaken"] = myMEOrder["ShakenIfEngaged"]
		}
	}

	return params
}

// Progress the activation of an order
func OrderActivation(col *db.Col, params map[string]interface{}) map[string]interface{} {

	OrderType := params["OrderType"].(string)
	Commander := params["Commander"].(string)
	Inspiration := params["Inspiration"].(string)
	Order := params["Order"].(string)
	Staff := params["Staff"].(float64)
	Grids := params["Grids"].(float64)
	Weather := params["Weather"].(float64)
	ActivationPoints := params["ActivationPoints"].(float64)
	CAUrge := params["CAUrge"].(bool)
	Vantage := params["Vantage"].(bool)
	NoLOS := params["NoLOS"].(bool)
	Rivalry := params["Rivalry"].(bool)
	Tired := params["Tired"].(bool)

	// Set default results
	params["Dice"] = ""
	params["ResultPoints"] = ""
	params["ResultActivated"] = false
	params["ResultFailed"] = false

	// Get the lookup records
	OrderActivation := list.Lookup(col, "OrderActivation", "Dice")

	// Apply all the modifiers
	adder := float64(0)
	Mods, _ := list.Get(col, "OrderActivationMod")

	Value := "Value"
	switch OrderType {
	case "Corps":
		Value = "CorpsValue"
	}
	for _, mod := range Mods.Data.([]interface{}) {
		myMod := mod.(map[string]interface{})

		code := myMod["Code"].(string)
		val := myMod[Value].(float64)
		switch code {
		case "BRK":
			if Order == "BreakOff" {
				adder += val
			}
		case "C1":
			if Commander == "Superior" {
				adder += val
			}
		case "C2":
			if Commander == "Excellent" {
				adder += val
			}
		case "C3":
			if Commander == "Good" {
				adder += val
			}
		case "C4":
			if Commander == "Average" {
				adder += val
			}
		case "C5":
			if Commander == "Poor" {
				adder += val
			}
		case "C6":
			if Commander == "Despicable" {
				adder += val
			}
		case "CC1":
			if Grids == 0 {
				adder += val
			}
		case "CHAR":
			if Inspiration == "Charismatic" && Order == "Attack" {
				adder += val
			}
		case "CORP":
			adder += val
		case "CU1":
			if CAUrge {
				adder += val
			}
		case "ELIT":
			if OrderType == "Elite" {
				adder += val
			}
		case "GRDB":
			if Order == "GB" {
				adder += val
			}
		case "GSTF":
			if Staff == 1 {
				adder += val
			}
		case "INSP":
			if Inspiration == "Inspirational" && Order == "Attack" {
				adder += val
			}
		case "NLOS":
			if NoLOS {
				adder += val
			}
		case "PSTF":
			if Staff == 3 {
				adder += val
			}
		case "RAIN":
			if Weather == 1 {
				adder += val
			}
		case "RETR":
			if Order == "Retreat" {
				adder += val
			}
		case "RIVL":
			if Rivalry {
				adder += val
			}
		case "RVAN":
			if Vantage {
				adder += val
			}
		case "SNOW":
			if Weather == 2 || Weather == 3 {
				adder += val
			}
		case "TIRD":
			if Tired {
				adder += val
			}
		case "GRID":
			adder += val * Grids
		case "UINS":
			if Inspiration == "Uninspiring" && Order == "Attack" {
				adder += val
			}
			if Inspiration == "Despicable" && Order == "Attack" {
				adder += val * 2
			}
		}
	}

	// Roll the Dice
	Dice := dice.DieRoll()
	TotalDice := Dice + int(adder)
	params["Dice"] = fmt.Sprintf("%d +%d (%d)", Dice, int(adder), TotalDice)

	fid := -1
	if TotalDice >= 1 {
		fid = 1
		if TotalDice >= 3 {
			fid = 3
			if TotalDice >= 6 {
				fid = 6
				if TotalDice >= 8 {
					fid = 8
					if TotalDice >= 9 {
						fid = 9
						if TotalDice >= 11 {
							fid = 11
							if TotalDice >= 13 {
								fid = 13
								if TotalDice >= 16 {
									fid = 16
									if TotalDice >= 18 {
										fid = 18
										if TotalDice >= 19 {
											fid = 19
											if TotalDice >= 20 {
												fid = 20

											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	if Dice == 2 || fid == -1 {
		params["ResultActivated"] = "Order has been lost, or disobeyed"
		params["ResultFailed"] = true
		params["ResultPoints"] = ""
		return params
	}

	DiceRecord := OrderActivation[fmt.Sprintf("%d", fid)]
	Points := DiceRecord["Points"].(float64)
	params["ResultPoints"] = Points

	ActivationPoints += Points
	if ActivationPoints > 10 {
		ActivationPoints = 10
		params["ResultActivated"] = true
	}
	params["ActivationPoints"] = ActivationPoints

	return params
}

// Attempt a Commander Action
func CommanderAction(col *db.Col, params map[string]interface{}) map[string]interface{} {

	Commander := params["Commander"].(string)
	Leader := params["Leader"].(string)
	Attachment := params["Attachment"].(float64)
	Action := params["Action"].(float64)

	CAScore := list.Lookup(col, "CAScore", "Code")[Commander]

	adder := float64(0)

	switch Leader {
	case "Charismatic":
		adder += 2
	case "Inspirational":
		adder += 1
	case "Average":
		adder += 0
	case "UnInspiring":
		adder += -1
	}

	switch Attachment {
	case 0:
		adder += 0
	case 1:
		adder += 3
	case 2:
		adder += -3
	}

	// Set default results
	params["Dice"] = ""
	params["Result"] = ""
	params["PassScore"] = ""
	params["NextPassScore"] = ""

	if Action > 0 {
		Field := fmt.Sprintf("A%d", int(Action))
		PassScore := int(CAScore[Field].(float64))
		params["PassScore"] = PassScore

		// Roll the Dice
		Dice := dice.DieRoll()
		TotalDice := Dice + int(adder)
		params["Dice"] = fmt.Sprintf("%d +%d (%d)", Dice, int(adder), TotalDice)

		if TotalDice >= PassScore {
			params["Result"] = "Pass"
			Action++
			if Action > 4 {
				Action = 0
			} else {
				Field2 := fmt.Sprintf("A%d", int(Action))
				params["NextPassScore"] = int(CAScore[Field2].(float64))
			}
		} else {
			params["Result"] = "Failed"
			Action = 0
		}
		params["Action"] = Action
	}

	return params
}

// Try for a Bonus Impulse move
func BonusImpulse(col *db.Col, params map[string]interface{}) map[string]interface{} {

	ArmyCommander := params["ArmyCommander"].(string)
	ArmyCA := params["ArmyCA"].(bool)
	CorpsCommander := params["CorpsCommander"].(string)
	CorpsCA := params["CorpsCA"].(bool)
	BoldLeader := params["BoldLeader"].(bool)
	MECA := params["MECA"].(bool)
	Defend := params["Defend"].(bool)
	Impetus := params["Impetus"].(bool)
	Shaken := params["Shaken"].(bool)
	Moved := params["Moved"].(bool)
	Fatigue := params["Fatigue"].(float64)
	Holding := params["Holding"].(float64)
	CAW := params["CAW"].(float64)
	CAD := params["CAD"].(float64)
	TookFlag := params["TookFlag"].(bool)
	Interp := params["Interp"].(bool)
	LostSP := params["LostSP"].(bool)
	TookSP := params["TookSP"].(bool)
	Town := params["Town"].(bool)
	TookHvBty := params["TookHvBty"].(bool)
	TookLtBty := params["TookLtBty"].(bool)
	Fog := params["Fog"].(bool)
	Mud := params["Mud"].(bool)

	// Set default results
	params["Dice"] = ""
	params["Result"] = ""
	params["ResultBonus"] = false
	params["ResultFatigue"] = false
	params["ResultOneUnit"] = false
	params["ResultFirefight"] = false

	// Get the lookup records
	BonusImpulse := list.Lookup(col, "BonusImpulse", "Score")

	// Apply all the modifiers
	adder := float64(0)
	Mods, _ := list.Get(col, "BonusImpulseMod")
	for _, mod := range Mods.Data.([]interface{}) {
		myMod := mod.(map[string]interface{})

		code := myMod["Code"].(string)
		val := myMod["Value"].(float64)
		switch code {
		case "A1":
			if TookLtBty {
				adder += val
			}
		case "A2":
			if TookHvBty {
				adder += val
			}
		case "C1":
			if ArmyCommander == "Charismatic" {
				adder += val
			}
		case "C2":
			if ArmyCommander == "Inspirational" {
				adder += val
			}
		case "C3":
			if ArmyCommander == "Impersonal" {
				adder += val
			}
		case "C4":
			if CorpsCommander == "Charismatic" {
				adder += val
			}
		case "C5":
			if CorpsCommander == "Inpsirational" {
				adder += val
			}
		case "C6":
			if CorpsCommander == "Impersonal" {
				adder += val
			}
		case "C7":
			if BoldLeader {
				adder += val
			}
		case "CA":
			if ArmyCA {
				adder += val
			}
			if CorpsCA {
				adder += val
			}
		case "CAD":
			adder += val * CAD
		case "CAW":
			adder += val * CAW
		case "DEF":
			if Defend {
				adder += val
			}
		case "FLG":
			if TookFlag {
				adder += val
			}
		case "FOG":
			if Fog {
				adder += val
			}
		case "FTG":
			adder += val * Fatigue
		case "IMP":
			if Impetus {
				adder += val
			}
		case "INT":
			if Interp {
				adder += val
			}
		case "LA":
			if MECA {
				adder += val
			}
		case "MEH":
			adder += val * Holding
		case "MV2":
			if Moved {
				adder += val
			}
		case "RSM":
			if Mud {
				adder += val
			}
		case "SHK":
			if Shaken {
				adder += val
			}
		case "SPL":
			if TookSP {
				adder += val
			}
			if LostSP {
				adder -= val
			}
		case "TWN":
			if Town {
				adder += val
			}

		}
	}

	// Roll the Dice
	Dice := dice.DieRoll()
	TotalDice := Dice + int(adder)
	params["Dice"] = fmt.Sprintf("%d +%d (%d)", Dice, int(adder), TotalDice)

	fid := 0
	if TotalDice >= 7 {
		fid = 7
		if TotalDice >= 12 {
			fid = 12
			if TotalDice >= 14 {
				fid = 14
				if TotalDice >= 17 {
					fid = 17
					if TotalDice >= 19 {
						fid = 19
					}
				}
			}
		}
	}

	DiceRecord := BonusImpulse[fmt.Sprintf("%d", fid)]

	params["Result"] = DiceRecord["Descr"]
	params["ResultBonus"] = DiceRecord["Bonus"]
	params["ResultFatigue"] = DiceRecord["Fatigue"]
	params["ResultOneUnit"] = DiceRecord["OneUnitOnly"]
	params["ResultFirefight"] = DiceRecord["FFOnly"]
	params["ResultAnotherME"] = DiceRecord["Another"]

	if DiceRecord["Fatigue"].(bool) {
		Fatigue++
		if Fatigue > 4 {
			Fatigue = 4
		}
		params["Fatigue"] = Fatigue
	}

	return params
}
