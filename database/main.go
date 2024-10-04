package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var db *pg.DB

func ReadyDatabase(options *pg.Options) (*pg.DB, error) {
    db = pg.Connect(options)

	return db, сreateSchema(db)
}

func FindUserById(id string) (*User, error) {
	user := &User{Id: id}
    err := db.Model(user).WherePK().Select()
    return user, err
}

func InsertUser(user User) (orm.Result, error) {
    return db.Model(&user).Insert()
}

func UpdateUser(user User) (orm.Result, error) {
    return db.Model(&user).Update()
}

func сreateSchema(db *pg.DB) error {
    models := []interface{}{
        (*User)(nil),
    }

    for _, model := range models {
        err := db.Model(model).CreateTable(&orm.CreateTableOptions{
            Temp: false,
            IfNotExists: true,
        })
        if err != nil {
            return err
        }
    }
    return nil
}