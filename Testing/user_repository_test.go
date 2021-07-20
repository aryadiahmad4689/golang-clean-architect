package testing

import (
	"context"
	"fmt"
	database "golang-clean-architecture/Database"
	entity "golang-clean-architecture/Entity"
	repository "golang-clean-architecture/Repository"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsert(t *testing.T) {
	db, err := database.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := repository.NewUser(db)

	ctx := context.Background()
	user := entity.User{
		Name:    "Ariadi Ahmad",
		Alamat:  "Makassar",
		Age:     23,
		Married: false,
		Date:    time.Now(),
	}
	result, err := userRepo.Store(ctx, user)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Name)
	// assert.Contains(t, "ariadi", result)

}

func TestFindUser(t *testing.T) {
	db, err := database.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := repository.NewUser(db)

	result, errFind := userRepo.FindId(context.Background(), int64(4))
	if errFind != nil {
		fmt.Println(errFind)
	}

	fmt.Println(result)

}

func TestDeleteUser(t *testing.T) {
	db, err := database.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := repository.NewUser(db)

	message := userRepo.Delete(context.Background(), int64(2))

	if err != nil {
		panic(err)
	}

	fmt.Println(message)
}
