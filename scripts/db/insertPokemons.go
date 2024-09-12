package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	db_pkg "github.com/zedosoad1995/pokemon-wordle/config/db"
	"github.com/zedosoad1995/pokemon-wordle/config/env"
	"github.com/zedosoad1995/pokemon-wordle/models/pokemon"
)

func main() {
	env.LoadEnvs()

	file, err := os.Open("data/pokemon_kaggle.csv")
	if err != nil {
		panic("Couldn't open the csv file")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		log.Fatalf("couldn't read the header row: %v", err)
		return
	}
	headerMap := make(map[string]int)
	for i, header := range headers {
		headerMap[header] = i
	}

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("couldn't read the csv file: %v", err)
		return
	}

	db := db_pkg.Init()
	db.Delete(&pokemon.Pokemon{})

	for _, record := range records {
		baseTotal, _ := strconv.ParseUint(record[headerMap["base_total"]], 10, 16)
		height, _ := strconv.ParseFloat(record[headerMap["height_m"]], 64)
		name := record[headerMap["name"]]
		pokedexNum, _ := strconv.ParseUint(record[headerMap["pokedex_number"]], 10, 16)
		type1 := nilIfEmpty(record[headerMap["type1"]])
		type2 := nilIfEmpty(record[headerMap["type2"]])
		weight, _ := strconv.ParseFloat(record[headerMap["weight_kg"]], 64)
		gen, _ := strconv.ParseUint(record[headerMap["generation"]], 10, 8)
		isLegendary, _ := stringToBool(record[headerMap["is_legendary"]])

		pokemon := pokemon.Pokemon{
			BaseTotal:   uint16(baseTotal),
			Height:      height,
			Name:        name,
			PokedexNum:  uint16(pokedexNum),
			Type1:       type1,
			Type2:       type2,
			Weight:      weight,
			Gen:         uint8(gen),
			IsLegendary: isLegendary,
		}
		db.Create(&pokemon)
	}
}

func stringToBool(s string) (bool, error) {
	switch s {
	case "1":
		return true, nil
	case "0":
		return false, nil
	default:
		return false, fmt.Errorf("invalid input: %s", s)
	}
}

func nilIfEmpty(s string) *string {
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return nil
	}
	return &trimmed
}
