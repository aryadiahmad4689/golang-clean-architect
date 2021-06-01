package testing

import (
	"fmt"
	database "golang-clean-architecture/Database"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnection(t *testing.T) {
	db, err := database.GetConnection()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}
