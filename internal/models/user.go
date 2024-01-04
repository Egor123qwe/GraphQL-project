package models

type User struct {
	Id    string `bson:"_id,omitempty"`
	Email string
	Name  string
	Age   int32
}
