package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vanphong2809/exercises4/database"
)

var db *database.DB = new(database.DB)

//Viết hàm: sau khi tạo user thì insert user_id vào user_point với số điểm 10.
func InsertUserToPoint(user database.User) error {
	err1 := db.InsertUser(user)
	if err1 != nil {
		return err1
	}
	point := database.Point{UserId: user.Id, Points: 10}
	err2 := db.InsertPoint(&point)
	if err2 != nil {
		return err2
	}
	return nil
}
func main() {
	if err := db.ConnectDb(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
	// tao database
	// db.Createtable()
	// if err2 := db.Createtable(); err2 != nil {
	// 	fmt.Println("fail")
	// } else {
	// 	fmt.Println("create database success")
	// }
	// now := time.Now().Unix()
	// user1 := database.User{"user1", "Nguyen Cong Phuong", 28091995, now, now}
	// err1 := db.InsertUser(user1)
	// if err1 != nil {
	// 	panic(err1)
	// }

	//Update User
	// user := database.User{}
	// user.Name = "Phong"
	// selector := database.User{Id: "User2"}
	// err2 := db.UpdateUser(&selector, &user)
	// if err2 != nil {
	// 	panic(err2)
	// }

	//ListUser
	// list, err3 := db.ListUsers()
	// if err3 != nil {
	// 	panic(err3)
	// }
	// if len(list) > 0 {
	// 	for _, item := range list {
	// 		fmt.Println(*item)
	// 	}
	// } else {
	// 	fmt.Println("khong co ban ghi nao")
	// }

	//FindUserById
	// fmt.Println("Tim user co id=user1")
	// user, err4 := db.FindUser("user1")
	// if err4 != nil {
	// 	panic(err4)
	// }
	// fmt.Println(*user)
	//Viết hàm: sau khi tạo user thì insert user_id vào user_point với số điểm 10.
	// user1 := database.User{Id: "user3", Name: "Nguyen Tuan Anh", Birth: 22112000, Created: time.Now().Unix(), UpdatedAt: time.Now().Unix()}
	// InsertUserToPoint(user1)

	//Bai2
	// user1 := &database.User{}
	// user1.Id = "user3"
	// user1.Birth = 23091997
	// db.Bai2(user1)

	//Bai3
	// InsertUser()
	db.ScanTableUser()

}
