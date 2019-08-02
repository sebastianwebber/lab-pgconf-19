package model

// Department contains a department information
type Department struct {
	tableName struct{} `sql:"department"` // default values are the same

	ID   int
	Name string
}
