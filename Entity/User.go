package entity

import "time"

type User struct {
	Id           int64
	Name, Alamat string
	Age          int64
	Married      bool
	Date         time.Time
}
