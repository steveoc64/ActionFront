package simulation

import (
	"fmt"
	"github.com/steveoc64/ActionFront/dice"
	"github.com/steveoc64/tiedot/db"
)

func Initiative(col *db.Col, params map[string]interface{}) map[string]interface{} {

	// Go through the whole InitTable table
	adderA := 0
	adderB := 0

	// Manually apply the parameters
	if params["AA"].(bool) {
		adderA += 3
	}
	if params["AW"].(bool) {
		adderA += 3
	}
	adderA += int(params["ACCA"].(float64))
	adderA += int(params["AB"].(float64))

	// Repeat for Side B
	if params["BA"].(bool) {
		adderB += 3
	}
	if params["BW"].(bool) {
		adderB += 2
	}
	adderB += int(params["BCCA"].(float64))
	adderB += int(params["BB"].(float64))

	// Roll the dice
	dA := dice.DieRoll()
	dB := dice.DieRoll()

	params["DiceA"] = fmt.Sprint(dA, "+", adderA, "=", dA+adderA)
	params["DiceB"] = fmt.Sprint(dB, "+", adderB, "=", dB+adderB)

	if dA+adderA >= dB+adderB {
		params["Result"] = "Side A wins the initiative"
	} else {
		params["Result"] = "Side B wins the initiative"
	}

	return params
}
