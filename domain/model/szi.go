package model

type Szi struct {
	ID   string `db:"id,omitempty" json:"id"`
	Name string `db:"name" json:"name"`
}
