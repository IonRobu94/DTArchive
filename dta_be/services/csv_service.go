package services

import (
	"dta_be/models"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

// this CSV must be downloaded. Later will be added to a DB
var baseUrl = "static/archive/"

func LoadBandsFomCSV() []models.Band {

	file, err := os.Open(baseUrl + "bands.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var bands []models.Band
	for _, record := range records[1:] { // Skip the header row
		id, err := strconv.Atoi(record[0])
		if err != nil {
			panic(err)
		}
		var fromIn int
		if record[4] != "N/A" {
			tempFromIn, err := strconv.Atoi(record[4])
			if err != nil {
				log.Panicf("Error converting value: %v", err)
				fromIn = -1
			} else {
				fromIn = tempFromIn
			}
		}

		if err != nil {
			panic(err)
		}

		band := models.Band{
			ID:      id,
			Name:    record[1],
			Country: record[2],
			Status:  record[3],
			FromIn:  fromIn,
			Genre:   record[5],
			Theme:   record[6],
			Active:  record[7],
		}
		bands = append(bands, band)
	}
	return bands
}
