package src

import (
	"github.com/valyala/fasthttp" //pacote utilizado no lugar do http pois o intervalo de inserção de dados é menor
)

//HandleRequests é responsável pela disponibilização do service, assim como a rota '/persons'.
func HandleRequests() {

	m := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/persons":
			Importfile(ctx)
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}
	port := ":8080"

	fasthttp.ListenAndServe(port, m)
}
