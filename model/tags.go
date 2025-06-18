package model

type Tags struct {
	Id   int    `grom:"type:itn;primary_key"`
	Name string `grom:"type:varchar(225)"`
}
