package main

import (
	"log"
)

func main() {
	log.Printf("windstatsd started")

	c := ReadConfig()

	log.Printf("staring with config %v", c)

	mc := NewMetClient("http://88.97.23.70:81")
	dbc := NewDBClient(c.DBAdrr, c.DBName, c.DBUser, c.DBPass)
	defer dbc.Close()

	mc.GetSystemInfo()
	mc.GetConnectionStatus()
	ms, err := mc.GetMeasurement()
	if err != nil {
		log.Fatalf("Error getting measurements: %v", err)
	}

	for _, m := range ms {
		if err := dbc.Insert("wind",
			map[string]string{"sourceId": "1"},
			map[string]interface{}{"speed": m.Speed, "direction": m.Direction}); err != nil {
			log.Printf("Error writing data %v : %v", m, err)
		}
	}
}
