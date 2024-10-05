package main

import (
	"encoding/csv"
	"os"
	"strconv"

	db_pkg "github.com/zedosoad1995/pokemon-wordle/config/db"
	"github.com/zedosoad1995/pokemon-wordle/config/env"
	poke_questions "github.com/zedosoad1995/pokemon-wordle/constants/pokemon/questions"
	"github.com/zedosoad1995/pokemon-wordle/models/pokemon"
	"github.com/zedosoad1995/pokemon-wordle/utils"
)

func main() {
	env.LoadEnvs()
	db := db_pkg.Init()

	pokemons, _ := pokemon.GetPokemons(db)

	labels := poke_questions.ValidQuestions
	N := len(labels)

	matrix := make([][]int, N)
	for i := range matrix {
		matrix[i] = make([]int, N)
	}

	for i, row := range labels {
		for j, col := range labels {
			rowLabel, rowVal := utils.ExtractLabelAndValue(row)
			colLabel, colVal := utils.ExtractLabelAndValue(col)

			resPokemons := pokemons.Filter(poke_questions.AllQuestions[rowLabel](rowVal).Condition, poke_questions.AllQuestions[colLabel](colVal).Condition)
			matrix[i][j] = len(resPokemons)
		}
	}

	file, err := os.Create("data/answersCountMatrix.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := append([]string{""}, labels...)
	if err := writer.Write(header); err != nil {
		panic(err)
	}

	for i, row := range matrix {
		rowData := []string{labels[i]}
		for _, val := range row {
			rowData = append(rowData, strconv.Itoa(val))
		}
		if err := writer.Write(rowData); err != nil {
			panic(err)
		}
	}
}
