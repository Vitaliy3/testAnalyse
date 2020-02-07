package irt

import (
	"database/sql"
	"gopkg.in/reform.v1"
)

var Rdb *reform.DB
var Db *sql.DB
type DifficultFull struct {
	index         []int
	Difficult     []int
	StandartError []int
	Difficult2    []int
}