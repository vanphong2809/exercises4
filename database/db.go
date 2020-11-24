package database

import (
	"xorm.io/xorm"
)

type DB struct {
	engine *xorm.Engine
}

var tables []interface{}

//ConnectDb open connection to db
func (d *DB) ConnectDb() error {
	engine, err := xorm.NewEngine("mysql", "root:password@tcp(0.0.0.0:3306)/test?charset=utf8")
	d.engine = engine
	return err
}

func Init() {
	tables = append(tables, new(User), new(Point))
}

func (d *DB) Createtable() error {
	Init()
	err := d.engine.CreateTables(tables...)
	if err != nil {
		return err
	}
	return nil
}
// func (d *DB) Sync2() error {
// 	Init()
// 	err := d.engine.Sync2(tables...)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
