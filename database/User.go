package database

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

type User struct {
	Id        string `json: "id"`
	Name      string `json: "name"`
	Birth     int64  `json: "birth"`
	Created   int64  `json: "created"`
	UpdatedAt int64  `json: "updated_at"`
}
type Data struct {
	Identity int
	User     User
}

func (db *DB) InsertUser(user User) error {
	affected, err := db.engine.Insert(&user)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("Insert fail")
	}
	return nil
}

func (db *DB) UpdateUser(selector, user *User) error {
	affected, err := db.engine.Update(user, selector)
	if affected == 0 {
		return errors.New("Khong the cap nhat")
	}
	return err
}

func (db *DB) ListUsers() ([]*User, error) {
	users := make([]*User, 0)
	err := db.engine.Find(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (db *DB) FindUser(id string) (*User, error) {
	user := &User{Id: id}
	ishas, err := db.engine.Get(user)
	if err != nil {
		return nil, err
	}
	if !ishas {
		return nil, errors.New("Not found")
	}
	return user, nil
}

func PrintUser(i int, wg *sync.WaitGroup, job chan *User) {
	for {
		select {
		case data := <-job:
			fmt.Println(i, data.Id, data.Name)
			wg.Done()
		}
	}
}
func (db *DB) ScanTableUser() error {

	jobs := make(chan *User, 100)
	defer close(jobs)
	var wg sync.WaitGroup

	// rows, err := db.engine.DB().Query("SELECT name from user")
	rows, err := db.engine.Rows(&User{})
	defer rows.Close()
	if err != nil {
		return err
	}

	user := new(User)
	for i := 0; i < 2; i++ {
		go PrintUser(i, &wg, jobs)
	}

	for rows.Next() {
		err := rows.Scan(user)
		if err != nil {
			log.Fatal(err)
		}
		if err == nil {
			// dataUser := &Data{Identity: i, User: *user}
			dataUser := *user
			// fmt.Println("Dia chi con tro user", &user.Id)
			// fmt.Println("Dia chi cua bien dataUser", &dataUser.Id)
			jobs <- &dataUser
			wg.Add(1)
		}
	}
	wg.Wait()
	return nil
}
