package database

import (
	"errors"
)

type Point struct {
	UserId    string `json: "user_id"`
	Points    int64  `json: "points"`
	MaxPoints int64  `json: "max_points"`
}

func (db *DB) InsertPoint(p *Point) error {
	c, err := db.engine.Insert(p)
	if err != nil {
		return err
	}
	if c == 0 {
		return errors.New("Can not insert")
	}
	return err
}

//Bai2 tạo 1 transaction khi update `birth` thành công thì cộng 10 điểm vào `point`
// sau đó sửa lại `name ` thành `$name + "updated "` nếu 1 quá trình fail thì rollback, xong commit (CreateSesson)
func (db *DB) Bai2(user *User) error {
	session := db.engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}
	//cap nhat birth trong bang User
	selectorUser := &User{Id: user.Id}
	updatorUser := &User{Id: user.Id}
	find, err := session.Get(updatorUser)
	if err != nil {
		session.Rollback()
		return err
	}
	if !find {
		session.Rollback()
		return errors.New("Not found user!")
	}
	updatorUser.Birth = user.Birth
	if _, err := session.Update(updatorUser, selectorUser); err != nil {
		session.Rollback()
		return err
	}
	//Cong 10 diem vao point
	selectorPoint := &Point{UserId: user.Id}
	updatorPoint := &Point{UserId: user.Id}
	_, err2 := session.Get(updatorPoint)
	if err2 != nil {
		session.Rollback()
		return err2
	}
	updatorPoint.Points += 10
	if _, err := session.Update(updatorPoint, selectorPoint); err != nil {
		session.Rollback()
		return err
	}
	//sau đó sửa lại `name ` thành `$name + "updated "`
	updatorUser.Name += "updated"
	if _, err := session.Update(updatorUser, selectorUser); err != nil {
		session.Rollback()
		return err
	}
	session.Commit()
	return nil
}
