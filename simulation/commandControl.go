package simulation

import (
	"fmt"
	"github.com/steveoc64/ActionFront/list"
	"github.com/steveoc64/tiedot/db"
	"log"
)

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

func stringSliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

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
