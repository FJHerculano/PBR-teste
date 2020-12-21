//Arquivo destinado a dados  á respeito da configuração

// após a estruturação de pastas e arquivos base, foi instalado o "go get github.com/BurntSushi/toml gopkg.in/mgo.v2 github.com/gorilla/mux" onde o toml será utilizado 
// para trabalhar com parser e encoder, o mux para criação das nossas rotas e o mgo que é o nosso driver de utilização do mongodb
//Importando pacotes instalados 

//nome do pacote 
package config 

//De fato a importação dos pacotes
import(
	"log"

	"github.com/BurntSushi/toml"

)
//Criação de estruct para passar dados de conexão db
type Config struct{
	Server string
	Database string
}

// Agora realiza leitura do arquivo config.toml, passando o dados para a nossa struct
func(c*Config) Read( ){
	if _, err := toml.DecodeFile("config.toml", &c); err != nil{
		log.Fatal(err)
	}
}