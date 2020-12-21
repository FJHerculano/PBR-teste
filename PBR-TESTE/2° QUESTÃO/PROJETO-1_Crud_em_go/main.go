// // Arquivo principal do projeto

// após a estruturação de pastas e arquivos base, foi instalado o "go get github.com/BurntSushi/toml gopkg.in/mgo.v2 github.com/gorilla/mux" onde o toml será utilizado 
// para trabalhar com parser e encoder, o mux para criação das nossas rotas e o mgo que é o nosso driver de utilização do mongodb

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/programadriano/go-restapi/config"
	. "github.com/programadriano/go-restapi/config/dao"
	movierouter "github.com/programadriano/go-restapi/router"
)

var dao = MoviesDAO{}
var config = Config{}
//Buscando iniciar uma conexão com banco de dados
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}
// A partir desse ponto, informamos as rotas a serem utilizadas ao longo dos testes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/movies", movierouter.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/movies/{id}", movierouter.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/movies", movierouter.Create).Methods("POST")
	r.HandleFunc("/api/v1/movies/{id}", movierouter.Update).Methods("PUT")
	r.HandleFunc("/api/v1/movies/{id}", movierouter.Delete).Methods("DELETE")

// servidor estara rodando na porta , localhost:3737
	var port = ":3737"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}