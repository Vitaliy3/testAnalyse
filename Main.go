package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/mysql"
	"log"
	"os"
	"testSystem/irt"
)

var Rdb *reform.DB
var Db *sql.DB

func main() {
	var err error

	Db, err = sql.Open("mysql", "root:kripper@/FTest")
	if err != nil {
		log.Print(err)
	}
	logger := log.New(os.Stderr, "SQL: ", log.Flags())
	Rdb = reform.NewDB(Db, mysql.Dialect, reform.NewPrintfLogger(logger.Printf))
	if Rdb == nil {
		panic("Rdb is NIL")
	}
	fmt.Println("Dddd")
	binar := irt.BinaryMatrix{}
	binar.TcBinaryMatrix()
		binar.Ffull()

	fmt.Println(len(binar.Value), "x", len(binar.Value[0]))
	//zadanie := len(binar.Number)
	//people := len(binar.Value)
	//for i := 0; i < zadanie; i++ {
	//	for j := 0; j < people; j++ {
	//		fmt.Print( "-",  ":", v.Value[j][i], "	")
	//	}
	//	fmt.Println()
	//}
	//var k irt.PIniter
	//kz := irt.InitPreparedness{}
	//kz.P(k)

}
