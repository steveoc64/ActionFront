package simulation

import (
	"fmt"
	"github.com/steveoc64/ActionFront/list"
	"github.com/steveoc64/tiedot/db"
	"log"
	"math"
	"math/rand"
)

func dieRoll() int {
	d1 := rand.Intn(9)
	d2 := rand.Intn(9)
	return 2 + d1 + d2
}

func shootDice(ammoOut int) (int, bool) {
	d1 := rand.Intn(9) + 1
	d2 := rand.Intn(9) + 1

	if d1 <= ammoOut {
		return d1 + d2, true
	} else {
		return d1 + d2, false
	}
}

func percent(p int) bool {

	d := rand.Intn(99)
	if d < p {
		return true
	}
	return false
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

					baseMove = baseMove * Weather["Move"].(float64) / 10.0
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

// For a given set of parameters, calculate the Tactical Move stats, and return this as a result set
func TacMove(col *db.Col, params map[string]interface{}) map[string]interface{} {

	retval := make(map[string]interface{})

	retval["UnitType"] = params["UnitType"]
	retval["DrillBook"] = params["DrillBook"]
	retval["Formation"] = params["Formation"]
	retval["FormationTo"] = params["FormationTo"]
	retval["Terrain"] = params["Terrain"]
	retval["Extra"] = params["Extra"]
	retval["Mud"] = params["Mud"]
	retval["Marsh"] = params["Marsh"]
	retval["Diagonal"] = params["Diagonal"]
	retval["LoWall"] = params["LoWall"]
	retval["HiWall"] = params["HiWall"]
	retval["LtWood"] = params["LtWood"]
	retval["HvWood"] = params["HvWood"]
	retval["Weather"] = params["Weather"]
	retval["Accumulated"] = params["Accumulated"]
	retval["Trained"] = params["Trained"]
	retval["Disorder"] = false
	retval["Fire"] = false
	retval["Quadrants"] = 0
	retval["Inches"] = 0
	retval["Frontage"] = 0
	retval["SK"] = ""

	// Get the TacMove record for this unit type
	TacMoves, _ := list.Get(col, "TacMove")
	Drills, _ := list.Get(col, "Drill")
	FormationChanges, _ := list.Get(col, "FormationChange")
	baseMove := float64(10)
	multiplier := float64(1)
	adder := float64(0)
	frontage := float64(1)
	isInf := false
	isCav := false
	isArt := false
	isSK := false
	disorder := params["Disorder"].(bool)
	canFire := true

	if params["Formation"].(string) == "Skirmish" {
		isSK = true
	}

	switch params["UnitType"].(string) {
	case "Artillery":
		isArt = true
	case "Cavalry", "LightCav":
		isCav = true
	default:
		isInf = true
	}

	// Adjust for the unit type
	for _, tacMove := range TacMoves.Data.([]interface{}) {
		myTacMove := tacMove.(map[string]interface{})

		// Adjust for Type of unit
		if myTacMove["UnitType"] == params["UnitType"] {
			baseMove = myTacMove["Move"].(float64)
		}
	}

	// Adjust for the drill type and formation, applies to infantry only
	if isInf {
		for _, drill := range Drills.Data.([]interface{}) {
			myDrill := drill.(map[string]interface{})
			if myDrill["Name"] == params["DrillBook"] {

				// Get the Drill entry for the start formation
				entries := myDrill["Entries"].(map[string]interface{})
				startDrill := entries[params["Formation"].(string)].(map[string]interface{})

				frontage = startDrill["FR"].(float64)
				retval["SK"] = fmt.Sprintf("%d / %d", int(startDrill["OO"].(float64)), int(startDrill["SK"].(float64)))
				multiplier = startDrill["EF"].(float64) / 10
			}
		}

		// If changing formation - do some major adjustments
		ffrom := params["Formation"].(string)
		fto := params["FormationTo"].(string)
		if ffrom == "Line" && fto != "Line Left" ||
			ffrom != "Line" && ffrom != fto {

			for _, fchange := range FormationChanges.Data.([]interface{}) {
				myfc := fchange.(map[string]interface{})
				era := "DIV"
				if params["DrillBook"].(string) == "Old School" {
					era = "AR"
				}
				if myfc["Era"].(string) == era &&
					myfc["From"].(string) == params["Formation"].(string) &&
					myfc["To"].(string) == params["FormationTo"].(string) {
					log.Println("Change record that applies", myfc)

					fceffect := float64(0)
					switch params["Trained"].(string) {
					case "UnTrained":
						fceffect = myfc["Untrained"].(float64)
					default:
						fceffect = myfc["Trained"].(float64)
					}

					switch fceffect {
					case -3:
						multiplier = 0
						canFire = false
						disorder = true
					case -2:
						multiplier = 0
						canFire = false
					case -1:
						multiplier = 0
						canFire = true
					case 0:
						// Not allowed
						retval["FormationTo"] = retval["Formation"]
					default:
						baseMove = fceffect
					}
				}
			}

		}
	}

	// Adjust for the terrain
	switch params["Terrain"].(string) {
	case "Marchfeld":
		multiplier *= 1.1
	case "Rolling":
		multiplier *= 1.0
	case "Rough":
		multiplier *= 0.8
		frontage -= 1
	case "Hill":
		if isCav {
			multiplier *= 0.3
			frontage -= 1
			canFire = false
		} else if isArt {
			multiplier *= 0.2
			frontage = 1
			canFire = false
		} else {
			multiplier *= 0.4
			frontage -= 2
			disorder = true
		}
	case "Town":
		frontage = 1
		if isCav {
			multiplier *= 0.2
			canFire = false
		} else if isArt {
			multiplier *= 0.1
			canFire = false
		} else {
			multiplier *= 0.5
			disorder = true
		}
	}

	// Adjust for terrain effects
	if params["LtWood"].(bool) {
		if isArt {
			multiplier *= 0.5
		} else if isCav {
			multiplier *= 0.5
			disorder = true
		} else {
			multiplier *= 0.7
			if !disorder {
				disorder = percent(40)
			}
			frontage -= 2
			if canFire {
				canFire = percent(60)
			}
		}
	}
	if params["HvWood"].(bool) {
		if isArt {
			multiplier *= 0
		} else if isCav {
			multiplier *= 0
		} else {
			multiplier *= 0.5
			if !disorder {
				disorder = percent(80)
			}
			frontage = 1
			if canFire {
				canFire = percent(30)
			}
		}
	}
	if params["Mud"].(bool) {
		if isArt {
			multiplier *= 0.2
			canFire = false
		} else if isCav {
			multiplier *= 0.3
			if !disorder {
				disorder = percent(50)
			}
			canFire = false
		} else {
			multiplier *= 0.5
			if !disorder {
				disorder = percent(30)
			}
			frontage -= 1
			if canFire {
				canFire = percent(60)
			}
		}
	}
	if params["Marsh"].(bool) {
		if isArt {
			multiplier *= 0
		} else if isCav {
			multiplier *= 0.6
			if !disorder {
				disorder = percent(50)
			}
			canFire = false
		} else {
			multiplier *= 0.7
			if !disorder {
				disorder = percent(50)
			}
			frontage = 2
			if canFire {
				canFire = percent(80)
			}
		}
	}
	if params["LoWall"].(bool) {
		if isArt {
			multiplier *= 0
		} else if isCav {
			adder -= 5
			canFire = false
		} else {
			adder -= 1
			if !disorder {
				disorder = percent(30)
			}
			if canFire {
				canFire = percent(90)
			}
		}
	}
	if params["HiWall"].(bool) {
		if isArt {
			multiplier *= 0
		} else if isCav {
			multiplier *= 0
		} else {
			adder -= 3
			if !disorder {
				disorder = percent(60)
			}
			if canFire {
				canFire = percent(60)
			}
		}
	}

	// If infantry, and not in Skirmish order, apply march pace
	if isInf && !isSK {
		switch params["Extra"].(float64) {
		case 0:
			if disorder {
				// small chance of recovering any disorder if advancing at a slow march
				disorder = percent(70)
			}
		case 1:
			adder += 1
			// small chance of becoming disordered, depending on terrain
			switch params["Terrain"].(string) {
			case "Marchfeld":
				if !disorder {
					disorder = percent(10)
				}
			case "Rolling":
				if !disorder {
					disorder = percent(20)
				}
			case "Rough":
				if !disorder {
					disorder = percent(30)
				}
			}

		case 3:
			adder += 3
			// Good chance of becoming disordered, depending on terrain
			switch params["Terrain"].(string) {
			case "Marchfeld":
				if !disorder {
					disorder = percent(30)
				}
			case "Rolling":
				if !disorder {
					disorder = percent(60)
				}
			case "Rough":
				if !disorder {
					disorder = percent(90)
				}
			}

		}
	}

	if isCav {
		switch params["Extra"].(float64) {
		case 0:
			adder -= 4
			// slow trotting pace, to ensure good order
			if disorder {
				disorder = percent(60)
			}
		case 1:
			if !disorder {
				disorder = percent(20)
			}
		case 3:
			adder += 6
			if !disorder {
				disorder = percent(40)
			}
		}
	}

	// Perform final calculations
	startInches := params["Accumulated"].(float64)
	inches := (baseMove * multiplier) + adder
	retval["Inches"] = math.Trunc(inches)
	inchesPerQuad := float64(5)
	if params["Diagonal"].(bool) {
		inchesPerQuad = 7
	}

	totalInches := inches + startInches
	quads := totalInches / inchesPerQuad
	if isSK {
		disorder = false
	}
	retval["Quadrants"] = math.Trunc(quads)
	retval["Accumulated"] = math.Trunc(math.Mod(totalInches, inchesPerQuad))
	retval["Disorder"] = disorder
	retval["Frontage"] = frontage
	retval["Fire"] = canFire

	return retval
}

// For a given set of parameters, calculate the Tactical Move stats, and return this as a result set
func VolleyFire(col *db.Col, params map[string]interface{}) map[string]interface{} {

	retval := make(map[string]interface{})

	retval["Rating"] = params["Rating"]
	retval["FirstFire"] = params["FirstFire"]
	retval["OppFire"] = params["OppFire"]
	retval["FSquare"] = params["FSquare"]
	retval["Disordered"] = params["Disordered"]
	retval["Shaken"] = params["Shaken"]
	retval["Ammo"] = params["Ammo"]
	retval["Hits"] = params["Hits"]
	retval["Fatigue"] = params["Fatigue"]
	retval["Range"] = params["Range"]
	retval["Bases"] = params["Bases"]
	retval["LtWood"] = params["LtWood"]
	retval["HvWood"] = params["HvWood"]
	retval["MdWood"] = params["MdWood"]
	retval["Rain"] = params["Rain"]
	retval["HRain"] = params["HRain"]
	retval["Cover"] = params["Cover"]
	retval["Enfilade"] = params["Enfilade"]
	retval["TargetF"] = params["TargetF"]
	retval["Dice"] = 0
	retval["Effect"] = ""
	retval["EffectHits"] = 0
	retval["EffectAmmo"] = ""

	adder := float64(0)

	// Go through the whole FireMods table

	firerRating := params["Rating"].(string)
	firerFatigue := params["Fatigue"].(float64)
	if firerFatigue > 4 {
		firerFatigue = 4
	}
	firerHits := params["Hits"].(float64)
	if firerHits > 6 {
		firerHits = 6
	}

	FireMods, _ := list.Get(col, "FireMod")
	for _, fmod := range FireMods.Data.([]interface{}) {
		myFireMod := fmod.(map[string]interface{})

		code := myFireMod["Code"].(string)
		val := myFireMod["Value"].(float64)
		switch code {
		case "FF":
			if params["FirstFire"].(bool) {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "DIS":
			if params["Disordered"].(bool) {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "SHK":
			if params["Shaken"].(bool) {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "SQ":
			if params["FSquare"].(bool) {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "AMD":
			if params["Ammo"].(float64) == 1 {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "AME":
			if params["Ammo"].(float64) == 2 {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "FLW":
			if params["LtWood"].(bool) {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "FMW":
			if params["MdWood"].(bool) {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "FHW":
			if params["HvWood"].(bool) {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "ENL":
			if params["Enfilade"].(bool) {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "TSQ":
			if params["TargetF"].(string) == "Square" {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "ART":
			if params["TargetF"].(string) == "Artillery" {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "TCOL":
			if params["TargetF"].(string) == "Column" {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "TCC":
			if params["TargetF"].(string) == "ClosedCol" {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "OO":
			if params["TargetF"].(string) == "OpenOrder" {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "SK":
			if params["TargetF"].(string) == "Skirmish" {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "CAV":
			if params["TargetF"].(string) == "Cavalry" {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "OPP":
			if params["OppFire"].(bool) {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "RN":
			if params["Rain"].(bool) {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "HR":
			if params["HRain"].(bool) {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "C1":
			if params["Cover"].(float64) == 1 {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "C2":
			if params["Cover"].(float64) == 2 {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "C3":
			if params["Cover"].(float64) == 3 {
				adder += val
				log.Println("apply", code, val, adder)
			}
		case "FTG":
			adder += (val * firerFatigue)
			log.Println("apply", code, val, "*", firerFatigue, adder)
		case "HIT":
			adder += (val * firerHits)
			log.Println("apply", code, val, "*", firerHits, adder)
		default:
			if code == firerRating {
				adder += val
				log.Println("apply", code, val, adder)
			}
		}
	}

	d, ammoOut := shootDice(1)
	retval["EffectAmmo"] = ammoOut

	if ammoOut {
		switch params["Ammo"].(float64) {
		case 0:
			retval["Ammo"] = 1
		case 1:
			retval["Ammo"] = 2
		}
	}
	d2 := int(adder)
	dice := d + d2
	retval["Dice"] = fmt.Sprintf("%d  +%d  (%d)", d, d2, dice)

	fid := 1

	if dice >= 1 {
		fid = 2
		if dice >= 5 {
			fid = 3
			if dice >= 9 {
				fid = 4
				if dice >= 12 {
					fid = 5
					if dice >= 15 {
						fid = 6
						if dice >= 19 {
							fid = 7
							if dice >= 23 {
								fid = 8
								if dice >= 29 {
									fid = 9
									if dice >= 34 {
										fid = 10
									}
								}
							}
						}
					}
				}
			}
		}
	}

	FireEffects, _ := list.Get(col, "FireEffect")
	for _, feffect := range FireEffects.Data.([]interface{}) {
		myFireEffect := feffect.(map[string]interface{})
		if fid == int(myFireEffect["ID"].(float64)) {
			retval["Effect"] = myFireEffect["Descr"]
		}
	}
	percentDamage := float64(0)
	FireCharts, _ := list.Get(col, "FireChart")
	for _, fchart := range FireCharts.Data.([]interface{}) {
		myFireChart := fchart.(map[string]interface{})
		if fid == int(myFireChart["ID"].(float64)) {
			percentDamage = myFireChart["SmallArms"].(float64)
		}
	}

	numBases := params["Bases"].(float64)
	damage := float64(0)
	switch params["Range"].(float64) {
	case 0:
		damage = percentDamage * 10 * numBases
	case 1:
		damage = percentDamage * 5 * numBases
	case 2:
		damage = percentDamage * 1 * numBases
	}

	fullHits := int(math.Trunc(damage / 100))
	partialHits := int(math.Mod(damage, 100))
	log.Println(damage, fullHits, partialHits)
	extraHit := 0
	if percent(partialHits) {
		extraHit = 1
	}

	//retval["EffectHits"] = fmt.Sprintf("%f (%d, %d)", damage, fullHits, extraHit)
	retval["EffectHits"] = fullHits + extraHit
	if dice < -5 {
		retval["EffectHits"] = 0
	}

	// No longer has first fire advantage
	retval["FirstFire"] = false

	return retval
}
