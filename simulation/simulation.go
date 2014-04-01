package simulation

import (
	"fmt"
	"github.com/steveoc64/ActionFront/list"
	"github.com/steveoc64/tiedot/db"
	//	"log"
	"math"
	"math/rand"
)

func dieRoll() int {

	d1 := rand.Intn(9)
	d2 := rand.Intn(9)
	return 2 + d1 + d2
}

// For a given set of parameters, calculate the GTMove, and return this as a result set
func GTMove(col *db.Col, params map[string]interface{}) map[string]interface{} {

	retval := make(map[string]interface{})

	retval["METype"] = params["METype"]
	retval["DeploymentState"] = params["DeploymentState"]
	retval["Terrain"] = params["Terrain"]
	retval["Weather"] = params["Weather"]
	retval["Accumulated"] = params["Accumulated"]
	retval["Forced"] = params["Forced"]
	retval["MarchOrder"] = params["MarchOrder"]
	retval["Diagonal"] = params["Diagonal"]
	retval["Distance"] = 0
	retval["Inches"] = 0

	var baseMove float64

	// get the GT Movement record for this METype
	GTMoves, _ := list.Get(col, "GTMove")
	for _, myMove := range GTMoves.Data.([]interface{}) {
		GTMove := myMove.(map[string]interface{})
		if GTMove["METype"] == params["METype"] {
			// We now have the correct GT Move record
			switch params["DeploymentState"] {
			case "Deployed":
				baseMove = GTMove["D1"].(float64)
				if params["MarchOrder"].(bool) {
					baseMove += 4
				}
			case "Bde Out":
				baseMove = GTMove["D2"].(float64)
			case "Deploying":
				baseMove = GTMove["D3"].(float64)
			case "Condensed Col":
				baseMove = GTMove["D4"].(float64)
			case "Regular Col":
				baseMove = GTMove["D5"].(float64)
			case "Extended Col":
				baseMove = GTMove["D6"].(float64)
			}

			// Lets see if we have a forced march on our hands
			if params["Forced"].(bool) {
				for _, fmove := range GTMoves.Data.([]interface{}) {
					checkf := fmove.(map[string]interface{})
					if checkf["METype"] == "Forced March" {
						switch params["DeploymentState"] {
						case "Deployed":
							baseMove += checkf["D1"].(float64)
						case "Bde Out":
							baseMove += checkf["D2"].(float64)
						case "Deploying":
							baseMove += checkf["D3"].(float64)
						case "Condensed Col":
							baseMove += checkf["D4"].(float64)
						case "Regular Col":
							baseMove += checkf["D5"].(float64)
						case "Extended Col":
							baseMove += checkf["D6"].(float64)
						}

					}
				}
			}

			acc := params["Accumulated"].(float64)
			turns := 1.0

			// Get the appropriate weather modifier
			w, _ := list.Get(col, "Weather")
			for _, myWeather := range w.Data.([]interface{}) {
				Weather := myWeather.(map[string]interface{})
				if Weather["Code"] == params["Weather"] {
					// We now have the appropriate weather as well
					//log.Println("Weather does this ", Weather)

					baseMove = baseMove * Weather["Move"].(float64) / 10.0
					//log.Println("Weather alters base move to ", baseMove)
				}
			}

			baseMove *= turns
			inchesPerGrid := 10.0
			if params["Diagonal"].(bool) {
				inchesPerGrid = 15.0
			}
			retval["Inches"] = math.Trunc(baseMove)
			retval["Distance"] = math.Trunc((baseMove + acc) / inchesPerGrid)
			retval["Accumulated"] = math.Trunc(math.Mod(baseMove+acc, inchesPerGrid))
		}
	}

	return retval
}

// For a given set of parameters, calculate the Deployment stats, and return this as a result set
func Deployment(col *db.Col, params map[string]interface{}) map[string]interface{} {

	retval := make(map[string]interface{})

	retval["DepRating"] = params["DepRating"]
	retval["DepState"] = params["DepState"]
	retval["MarchCol"] = params["MarchCol"]
	retval["Darkness"] = params["Darkness"]
	retval["Choke"] = params["Choke"]
	retval["Mud"] = params["Mud"]
	retval["Fog"] = params["Fog"]
	retval["Grids"] = params["Grids"]
	retval["Dice"] = 0
	retval["DieMods"] = 0
	retval["Result"] = ""

	// get the adjustment bonus for this rating

	adjust := 0

	DepMods, _ := list.Get(col, "DeploymentMod")

	for _, depMod := range DepMods.Data.([]interface{}) {
		myDepMod := depMod.(map[string]interface{})

		val := int(myDepMod["Value"].(float64))

		// Adjust for Type of unit
		if myDepMod["Descr"] == params["DepRating"] {
			adjust += val
		}

		// Adjust for other known conditions
		if params["Mud"].(bool) && myDepMod["Code"] == "MUD" {
			adjust += val
		}
		if params["Fog"].(bool) && myDepMod["Code"] == "FOG" {
			adjust += val
		}
		if params["Choke"].(bool) && myDepMod["Code"] == "CP" {
			adjust += val
		}
		if params["Darkness"].(bool) && myDepMod["Code"] == "DK" {
			adjust += val
		}
		if params["Grids"].(float64) != 0 && myDepMod["Code"] == "MV" {
			adjust += (int(params["Grids"].(float64)) * val)
		}
	}

	retval["DieMods"] = adjust
	d := dieRoll()
	Score := d + adjust

	retval["Dice"] = fmt.Sprintf("%d + %d = %d", d, adjust, Score)

	// Convert the DepState to a number
	depState := 0
	switch params["DepState"] {
	case "Deployed":
		depState = 1
	case "Bde Out":
		depState = 2
	case "Deploying":
		depState = 3
	case "Condensed Col":
		depState = 4
	case "Regular Col":
		depState = 5
	case "Extended Col":
		depState = 6
	}

	direction := -1
	if params["MarchCol"].(bool) {
		direction = 1
	}
	change := 0

	// Compare the adjusted die roll to the score needed
	if Score >= 1 {
		change = 1
		if Score >= 10 {
			change = 2
			if Score >= 16 {
				change = 3
			}
		}
	}
	change *= direction
	depState += change
	if depState < 1 {
		depState = 1
	}
	if depState > 6 {
		depState = 6
	}
	if change == 0 {
		retval["Result"] = "No Change"
	} else {
		resString := params["DepState"].(string) + " -> "
		switch depState {
		case 1:
			retval["DepState"] = "Deployed"
		case 2:
			retval["DepState"] = "Bde Out"
		case 3:
			retval["DepState"] = "Deploying"
		case 4:
			retval["DepState"] = "Condensed Col"
		case 5:
			retval["DepState"] = "Regular Col"
		case 6:
			retval["DepState"] = "Extended Col"
		}
		retval["Result"] = resString + retval["DepState"].(string)
	}

	return retval
}
