package testing

import (
	"context"
	"fmt"
	database "golang-clean-architecture/Database"
	entity "golang-clean-architecture/Entity"
	repository "golang-clean-architecture/Repository"
	service "golang-clean-architecture/Service"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestAdd(t *testing.T) {
	db, _ := database.GetConnection()
	repositoryUser := repository.NewUser(db)
	service := service.NewUserService(&repositoryUser)
	user := entity.User{
		Name:    "Ariadi Ahmad",
		Alamat:  "Makassar",
		Age:     23,
		Married: false,
		Date:    time.Now(),
	}
	userService, _ := service.AddUser(context.Background(), user)

	fmt.Println(userService)
}

func TestFind(t *testing.T) {
	db, err := database.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := repository.NewUser(db)
	service := service.NewUserService(&userRepo)
	userService, _ := service.FindUseer(context.Background(), 3)

	fmt.Println(userService)

}

func TestDelete(t *testing.T) {
	db, err := database.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := repository.NewUser(db)
	service := service.NewUserService(&userRepo)
	userService := service.DeleteUser(context.Background(), int64(6))

	fmt.Println(userService)
}
