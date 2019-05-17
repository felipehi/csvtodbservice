# CSVTODBSERVICE
serviço em GO que recebe um arquivo csv/txt de entrada, persistindo no banco de dados relacional

# Compilar
Além de ter o Go instalado no sistema operacional é necessário executar:

export GOPATH=/src <br/>
go get github.com/valyala/fasthttp <br/>
go get github.com/lib/pq <br/>


# Executar <br/>
Basta executar o arquivo main e ele ficará ouvindo na porta 8080 por novas requisições. <br/>

go run main.go <br/>

# Uso
Enviar um comando POST no endpoint http://localhost:8080/persons <br/>
O retorno será uma mensagem com a informação: <br/>
mensagem: "Erro ao importar arquivo" - o arquivo encontrou um erro ao ser importado. <br/>
mensagem: "Arquivo importado com sucesso." - o arquivo foi importado com sucesso. <br/>

