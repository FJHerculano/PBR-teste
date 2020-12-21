// Arquivo de estruturação de modelo de personagens presentes no projeto

package models

import "gopkg.in/mgo.v2/bson"
// modelo basico para crud, buscando passar o que se é necessario para a sua criação

type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	ThubImage   string        `bson:"thumb_image" json:"thumb_image"`
	Description string        `bson:"description" json:"description"`
	Active      bool          `bson:"active" json:"active"`
}