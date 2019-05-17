package src

import (
	"net/http"

	"github.com/valyala/fasthttp"
)

//Importfile é responsável por orquestrar as chamadas necessárias para realizar a importação do arquivo.
func Importfile(ctx *fasthttp.RequestCtx) {

	stringConnection := GetStringConnection()
	csvData := Importcsv()
	err := Insert(csvData, stringConnection)
	if err != nil {
		ctx.Error("Erro ao importar arquivo", 500)
		return
	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.SetStatusCode(http.StatusOK)
	ctx.WriteString("Arquivo importado com sucesso.")
}
