package simulation

import (
	"fmt"
	"github.com/steveoc64/ActionFront/dice"
	"github.com/steveoc64/ActionFront/list"
	"github.com/steveoc64/tiedot/db"
	"log"
	"math"
)

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

	d, ammoOut := dice.ShootDice(1)
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
	dieScore := d + d2
	retval["Dice"] = fmt.Sprintf("%d  +%d  (%d)", d, d2, dieScore)

	fid := 1

	if dieScore >= 1 {
		fid = 2
		if dieScore >= 5 {
			fid = 3
			if dieScore >= 9 {
				fid = 4
				if dieScore >= 12 {
					fid = 5
					if dieScore >= 15 {
						fid = 6
						if dieScore >= 19 {
							fid = 7
							if dieScore >= 23 {
								fid = 8
								if dieScore >= 29 {
									fid = 9
									if dieScore >= 34 {
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
	if dice.Percent(partialHits) {
		extraHit = 1
	}

	//retval["EffectHits"] = fmt.Sprintf("%f (%d, %d)", damage, fullHits, extraHit)
	retval["EffectHits"] = fullHits + extraHit
	if dieScore < -5 {
		retval["EffectHits"] = 0
	}

	// No longer has first fire advantage
	retval["FirstFire"] = false

	return retval
}
