package main

import (
	"strconv"
	"time"

	"github.com/vanphong2809/exercises4/database"
)

//insert 100 ban ghi
func InsertUser() {
	for i := 0; i < 100; i++ {
		user := database.User{}
		user.Id = "u" + strconv.Itoa(i)
		user.Name = "name" + strconv.Itoa(i)
		user.Birth = 28091998
		user.Created = time.Now().Unix()
		user.UpdatedAt = time.Now().Unix()
		db.InsertUser(user)
	}
}
