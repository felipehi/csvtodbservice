package src

import (
	"log"
	"regexp"
)

//CleanChar realiza a limpeza dos dados, retirando os caracteres especiais e mantendo apenas os alfanuméricos.
func CleanChar(fieldchar string) string {

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedfield := reg.ReplaceAllString(fieldchar, "")
	return processedfield
}
