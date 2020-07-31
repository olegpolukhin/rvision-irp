package model

type SziManufacture struct {
	ID              string `db:"id,omitempty" json:"id"`
	Name            string `db:"name" json:"name"`
	PermissionLevel uint16 `db:"permission_level" json:"-"`
}

