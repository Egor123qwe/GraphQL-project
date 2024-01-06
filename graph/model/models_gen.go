// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Query struct {
}

type User struct {
	ID         string   `json:"id"`
	Email      string   `json:"email"`
	Name       *string  `json:"name,omitempty"`
	Age        int      `json:"age"`
	University string   `json:"university"`
	Course     *int     `json:"course,omitempty"`
	Hobbies    []string `json:"hobbies,omitempty"`
}
