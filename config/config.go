package config

import (
	"fmt"

	"github.com/kikiluvv/clockr/db"
)

func ShowConfig() {
	dbData, _ := db.Load()
	fmt.Println("Pay period:", dbData.Config.PayPeriod)
	fmt.Println("Start of period:", dbData.Config.StartOfPeriod)
	fmt.Println("End of period:", dbData.Config.EndOfPeriod)
}

func SetConfig(field, value string) {
	dbData, _ := db.Load()
	switch field {
	case "pay_period":
		dbData.Config.PayPeriod = value
	case "start_of_period":
		dbData.Config.StartOfPeriod = value
	case "end_of_period":
		dbData.Config.EndOfPeriod = value
	default:
		fmt.Println("Unknown config field:", field)
		return
	}
	db.Save(dbData)
	fmt.Println("Config updated:", field, "=", value)
}
