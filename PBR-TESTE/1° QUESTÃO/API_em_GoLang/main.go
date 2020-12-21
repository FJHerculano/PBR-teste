package main
//importando os pacotes que serão utilizados no projeto
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
//Inicializando rotas, passando o valor true para StrictSlash
func main() {
	rotas := mux.NewRouter().StrictSlash(true)
// Rota de requisição get para listar todos as pessoas, primeiro passamos a rota, segundo o nome do metodo , em seguida a rota de post que serve para criar pessoa
	rotas.HandleFunc("/", getAll).Methods("GET")
	rotas.HandleFunc("/persons", create).Methods("POST")
	var port = ":3838"  // Utilizaremos essa porta para realizar os testes de requisição via -> localhost:3838
	fmt.Println("Server running in port:", port) // Terá uma impressão va console da porta utilizada em caso de sucesso ao startar a aplicação
	log.Fatal(http.ListenAndServe(port, rotas))  // passando rotas para o listen http

}
//Criação da struct / modulo Person, e trazendo uma inicialização dele com dois valores
type Person struct {
	Name string
}

var persons = []Person{

	Person{Name: "Herculano"},
	Person{Name: "papai-noel"},
}
//Criação do metodo getAll com os parametro de requisição e resposta, eles serven para retornar os valores do array que criamos à cima.
func getAll(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(persons)
}
//setando o content-type do header, configurando  header em si
func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//criando variavel com valor do nosso modelo estrutural person
	var p Person
//recebendo os valores do nosso request e passando para duas variaveis que são : body e err.
	body, err := ioutil.ReadAll(r.Body)
// Até a linha 61 buscamos algumas validações ... tais como, se o erro é diferente de nil, se o arquivo veio vazio , e retornamos para o usuario
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &p); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	json.Unmarshal(body, &p)

	persons = append(persons, p)
// valido lembrar que como não estar sendo utilizado um banco de dados nesta api estaremos utilizando o array criado.
// Por fim finalizamos retornando um status 201, mostrando o sucesso na requisição.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		panic(err)
	}
}