package infrasructure

import (
	"context"
	"fmt"
	database "golang-clean-architecture/Database"
	entity "golang-clean-architecture/Entity"
	repository "golang-clean-architecture/Repository"
	service "golang-clean-architecture/Service"
	"net/http"
	"strconv"
	"time"
)

func AddUser(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		panic(err)
	}

	age, _ := strconv.Atoi(r.PostForm.Get("age"))
	married, _ := strconv.ParseBool(r.PostForm.Get("married"))
	date, err := time.Parse("2006-01-02", r.PostFormValue("date"))

	if err != nil {
		panic(err)
	}

	user := entity.User{
		Name:    r.PostForm.Get("nama"),
		Alamat:  r.PostForm.Get("alamat"),
		Age:     int64(age),
		Married: married,
		Date:    date,
	}

	db, err := database.GetConnection()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	userRepo := repository.NewUser(db)
	service := service.NewUserService(&userRepo)
	result, err := service.AddUser(context.Background(), user)
	fmt.Println(result, err)
}
