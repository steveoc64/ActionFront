package simulation

import (
	"fmt"
	"github.com/steveoc64/ActionFront/dice"
	"github.com/steveoc64/ActionFront/list"
	"github.com/steveoc64/tiedot/db"
	"log"
)

func LeaderDeath(col *db.Col, params map[string]interface{}) map[string]interface{} {

	adder := float64(0)
	Charmed := int(params["Charmed"].(float64))
	Situation := int(params["Situation"].(float64))
	Nation := params["Nation"].(float64)
	Hits := params["Hits"].(float64)
	POD := params["POD"].(bool)
	Foolish := params["Foolish"].(bool)
	Rifle := params["Rifle"].(bool)
	LoseCA := params["LoseCA"].(bool)

	MaxInjury := 0
	MinInjury := 0

	switch Charmed {
	case 0: // Naturally unlucky
		adder += -2
		MinInjury = 1
	case 1: // Clumsy
		adder += -1
		MinInjury = 2
	case 3: // Debonairre
		if Foolish {
			MinInjury = 3
		}
	case 4: // Charmed
		MaxInjury = 3
		POD = false
	}

	if Charmed != 3 {
		Foolish = false
	}

	switch Nation {
	case 1, 2: // French and British, lead from the front
		if Situation >= 3 {
			adder += -1
		}
	case 3: // Cautious
		adder += 1
	}

	if Rifle {
		adder += 1
	}

	if POD {
		adder += -2
	}

	/* Types of injury

	Escape   0
	Inspired 0
	Drunk    0
	Slowed   0
	Stunned  1
	Light    2
	Serious  3
	Critical 4
	Death    5
	Captured 6

	*/
	isCA := false
	Danger := false

	switch Situation {
	case 0: // No Danger
		if Charmed == 0 {
			MaxInjury = 4 // Naturally unlucky may get a critical injury
			Danger = true
		}
	case 1: // Within cannon range of enemy
		switch Charmed {
		case 0:
			MaxInjury = 3 // the closer to danger he is, the luckier he gets !
			Danger = true
		case 1:
			MaxInjury = 2
			Danger = true
		case 2, 3, 4:
			adder += 1
			Danger = true
		}
	case 2: // Part of engaged ME
		switch Charmed {
		case 0:
			MaxInjury = 2
		case 1:
			MaxInjury = 4
		case 2, 3, 4:
			MaxInjury = 5
			adder += 1
		}
		Danger = true

	case 3: // Attached to unit which is engaged
		adder -= Hits / 2
		Danger = true
		if Foolish {
			adder += -1
		}
		if Charmed == 0 {
			MaxInjury = 4
		}
	case 4: // Attached to unit involved in close action
		adder -= Hits / 1.5
		isCA = true
		Danger = true
		if Foolish {
			adder += -2
		}
		if Charmed == 0 {
			MaxInjury = 4
		}
	case 5: // Attached to unit involved in melee
		adder -= Hits
		adder += -1
		isCA = true
		Danger = true
		if Foolish {
			adder += -3
		}
		if Charmed == 0 {
			MaxInjury = 4
		}
	case 6: // Performed a follow me
		adder -= Hits / 1.5
		adder += -2
		isCA = true
		Danger = true
		if Foolish {
			adder += -3
		}
		if Charmed == 0 {
			MaxInjury = 4
		}
	}

	// Roll the Dice
	Dice := dice.DieRoll()
	TotalDice := Dice + int(adder)
	params["Dice"] = fmt.Sprintf("%d +%d (%d)", Dice, int(adder), TotalDice)

	params["Result"] = ""
	params["Severity"] = ""

	if POD {
		// If Premonition of Death, then allow Death to occur
		MaxInjury = 0
	}

	if Danger {
		Injuries := list.InjuryLookup(col)
		if TotalDice < 5 {

			if MaxInjury < MinInjury {
				// Prevent infinite loop from never finding an appropriate wound
				MaxInjury = 0
			}
			for gotOne := false; !gotOne; {
				Hi := dice.DieRoll()
				Lo := dice.D6()

				params["Result"] = fmt.Sprintf("Calculating %d %d", Hi, Lo)
				params["ResultSeverity"] = "Calculating ..."
				KeyVal := uint16(Hi*10 + Lo)
				if isCA {
					KeyVal += 1000
				}
				gotOne = true
				params["Result"] = Injuries[KeyVal]["Descr"]
				params["Severity"] = Injuries[KeyVal]["Severity"]

				if MinInjury > 0 {
					// Check that the injury is of the minimum type
					switch params["Severity"] {
					case "Escape", "Inspired", "Uninspired", "Drunk", "Slowed":
						gotOne = false
					case "Stunned":
						if MinInjury > 1 {
							gotOne = false
						}
					case "Light":
						if MinInjury > 2 {
							gotOne = false
						}
					case "Serious":
						if MinInjury > 3 {
							gotOne = false
						}
					case "Critical":
						if MinInjury > 4 {
							gotOne = false
						}
					case "Death":
						if MinInjury > 5 {
							gotOne = false
						}
					}
				}
				if gotOne && MaxInjury > 0 {
					// Check that the injury is of the maximum allowed type
					switch params["Severity"] {
					case "Escape", "Inspired", "Uninspired", "Drunk", "Slowed":
						gotOne = false
					case "Stunned":
						gotOne = false
					case "Light":
						if MaxInjury < 2 {
							gotOne = false
						}
					case "Serious":
						if MaxInjury < 3 {
							gotOne = false
						}
					case "Critical":
						if MaxInjury < 4 {
							gotOne = false
						}
					case "Death":
						if MaxInjury < 5 {
							gotOne = false
						}
					case "Captured":
						if MaxInjury < 6 {
							gotOne = false
						}
					}
				}
				if !gotOne {
					log.Println("Retry", params["Severity"], "Max:", MaxInjury, "Min:", MinInjury)
				}
			}
		} else if TotalDice == 5 && LoseCA {
			params["Result"] = "The General is captured during the close action"
			params["ResultSeverity"] = "Captured"
		} else {
			params["Result"] = "The General keeps out of harms way"
		}
	} else {
		params["Result"] = "The General is in no Danger at this stage"
		params["Severity"] = ""
		params["Dice"] = ""
	}

	return params
}