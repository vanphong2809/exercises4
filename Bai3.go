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

// func PrintUser(buffchan chan *database.Data, wg *sync.WaitGroup) {
// 	for {
// 		select {
// 		case data := <-buffchan:
// 			fmt.Printf("Line %v - %v - %v\n", data.Iden, data.User.Id, data.User.Name)
// 			wg.Done()
// 		}
// 	}
// }

// func GetNameOfUser() error {
// 	buffchan := make(chan *database.Data, 100)
// 	defer close(buffchan)
// 	var wg sync.WaitGroup

// 	for i := 0; i < 2; i++ {
// 		go PrintUser(buffchan, &wg)
// 	}
// 	err := db.ScanforRow(buffchan, &wg)
// 	if err != nil {
// 		return err
// 	}
// 	wg.Wait()
// 	return nil

// }
