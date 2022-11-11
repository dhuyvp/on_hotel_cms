package models

type Person struct {
	Name    *string `db:"Name" json:"Name"`
	Age     *int    `db:"Age" json:"Age"`
	Logo    *string `db:"Logo" json:"Logo"`
	IsChild *bool   `db:"IsChild" json:"IsChild"`
}
