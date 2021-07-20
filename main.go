package main

import (
	"context"
	database "golang-clean-architecture/Database"
	route "golang-clean-architecture/Route"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// migrate := flag.Bool("migrate", false, "to migrate table")
	// if *migrate {
	// 	MigrateTable()
	// 	return
	// }
	MigrateTable()
	mux := route.Route()
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func MigrateTable() {
	db, _ := database.GetConnection()
	defer db.Close()
	ctx := context.Background()
	res, _ := db.ExecContext(ctx, "CREATE table IF NOT EXISTS user(id int PRIMARY key auto_increment,nama VARCHAR(100),alamat VARCHAR(100), age int, married boolean, date date)")
	res.RowsAffected()
	// return err
}
