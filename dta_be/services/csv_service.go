package services

import (
	"dta_be/models"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

var baseUrl = "static/archive/"

func GetAllData() (bands []models.Band, albums []models.Album, reviews []models.Review) {
	bands = mapRecordsToBands(loadDataFromCSV("bands"))
	albums = mapRecordsToAlbums(loadDataFromCSV("albums"))
	reviews = mapRecordsToReviews(loadDataFromCSV("reviews"))

	return bands, albums, reviews
}

func mapRecordsToReviews(reviewsRecords [][]string) (reviews []models.Review) {
	for _, record := range reviewsRecords[1:] { // Skip the header row
		id, err := strconv.Atoi(record[0])
		if err != nil {
			panic(err)
		}
		albumId, err := strconv.Atoi(record[1])
		if err != nil {
			panic(err)
		}
		score, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			panic(err)
		}
		review := models.Review{
			ID:      id,
			AlbumID: albumId,
			Title:   record[2],
			Score:   score,
			Content: record[4],
		}
		reviews = append(reviews, review)
	}
	return reviews
}
func mapRecordsToAlbums(albumsRecords [][]string) (albums []models.Album) {
	for _, record := range albumsRecords[1:] { // Skip the header row
		id, err := strconv.Atoi(record[0])
		if err != nil {
			panic(err)
		}
		bandId, err := strconv.Atoi(record[1])
		if err != nil {
			panic(err)
		}
		year, err := strconv.Atoi(record[3])
		if err != nil {
			panic(err)
		}
		album := models.Album{
			ID:     id,
			BandID: bandId,
			Title:  record[2],
			Year:   year,
		}
		albums = append(albums, album)
	}
	return albums
}

func mapRecordsToBands(bandsRecords [][]string) (bands []models.Band) {
	for _, record := range bandsRecords[1:] { // Skip the header row
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

func loadDataFromCSV(dataType string) (records [][]string) {
	// dataType should be "bands", "albums", "reviews"
	file, err := os.Open(baseUrl + dataType + ".csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	records, err = reader.ReadAll()
	if err != nil {
		panic(err)
	}
	return records
}

func LoadBandsFromCSV() []models.Band {

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
