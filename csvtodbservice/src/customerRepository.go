package src

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
)

//Insert fará a inserção dos dados do arquivo no banco de dados.
func Insert(csvData [][]string, stringConnection string) error {

	for index, each := range csvData[1:] {

		cpf := CleanChar(each[0])
		private, err := strconv.Atoi(each[1])
		incompleto, err := strconv.Atoi(each[2])
		data := each[3]
		ticketmedio := strings.Replace(each[4], ",", ".", -1)
		ultcompra := strings.Replace(each[5], ",", ".", -1)
		cnpj1 := CleanChar(each[6])
		cnpj2 := CleanChar(each[7])

		if !CpfValid(cpf) {
			log.Println("O CPF da linha " + strconv.Itoa(index+1) + "  não é válido")
			continue
		}
		if !CnpjValid(cnpj1) {
			log.Println("O campo 'LOJA MAIS FREQUÊNTE' da linha " + strconv.Itoa(index+1) + " não é válido")
			continue
		}
		if !CnpjValid(cnpj2) {
			log.Println("O campo 'LOJA DA ÚLTIMA COMPRA' da linha " + strconv.Itoa(index+1) + " não é válido")
			continue
		}

		db, err := sql.Open("postgres", stringConnection)
		if err != nil {
			log.Fatal(err)
		}
		stmt := `insert into customer (cpf, private, incompleto, datultimacompra, compraticketmedio, ticketultcompra, lojamaisfrequente, lojaultimacompra) VALUES($1,$2,$3,$4,$5,$6,$7,$8);`

		_, err = db.Exec(stmt, cpf, private, incompleto, data, ticketmedio, ultcompra, cnpj1, cnpj2)
		if err != nil {
			return err
		}

		db.Close()
	}

	return nil
}

//CpfValid é responsável por validar o CPF. Caso não tenha 11 dígitos, será considerado inválido.
func CpfValid(cpf string) bool {

	if len(cpf) != 11 {
		return false
	}
	return true
}

//CnpjValid é responsável por validar o CNPJ. Caso não tenha 14 dígitos, será considerado inválido.
func CnpjValid(cnpj string) bool {

	if len(cnpj) != 14 {
		return false
	}
	return true
}
