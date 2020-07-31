package model

type User struct {
	ID       string `db:"id,omitempty" json:"id"`
	Name     string `db:"name"  json:"name"`
	Password string `db:"password"  json:"password"`
}

