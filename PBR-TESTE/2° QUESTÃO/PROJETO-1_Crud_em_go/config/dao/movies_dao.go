// Local destinado a buscar uma interação com o mongoDB, um "repositorio"
// utilização de docker para conexão de banco de dados 
package dao

import (
	"log"

	. "github.com/programadriano/go-restapi/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
//Criando uma estruct para se conectar com o banco de dados
type MoviesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database
//passando a COLLECTION que mapearemos para o nosso repositorio 
const (
	COLLECTION = "movies"
)
//Criando uma session com o banco de dados , utilizasse o metodo Dial do pacote mgo, pois em caso de erro  ele "matara" a session, em caso de sucesso... 
//retorna o valor para a variavel db 
func (m *MoviesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}
// Inicio de metodos, que buscasse retorno de um json.
func (m *MoviesDAO) GetAll() ([]Movie, error) {
	var movies []Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}
//Função similar a anterior, poréns trazendo apenas um intem especificado por id 
func (m *MoviesDAO) GetByID(id string) (Movie, error) {
	var movie Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}
//Metodo de requisição de criação/ registro
func (m *MoviesDAO) Create(movie Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}
// Metodo de exclusão de dados 
func (m *MoviesDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}
// metodo de atualização/ edição dos dados
func (m *MoviesDAO) Update(id string, movie Movie) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &movie)
	return err
}