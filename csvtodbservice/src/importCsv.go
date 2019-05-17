package src

import (
	"encoding/csv"
	"fmt"
	"os"
)

//Importcsv é responsável por fazer a importação dos dados do arquivo, separando as colunas conforme configuração
func Importcsv() [][]string {

	csvFile, err := os.Open("./base_teste.csv")

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.Comma = ' ' // separador utilizado

	reader.FieldsPerRecord = -1

	reader.TrimLeadingSpace = true

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return csvData

}
