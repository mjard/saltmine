package main

import (
	"database/sql"
	"errors"
	"io"
	"log"

	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"code.google.com/p/go.crypto/bcrypt"
)

type LiteBookie struct {
	db   *sql.DB
	elog *log.Logger
	tlog *log.Logger
}

func NewLiteBookie(eWriter, tWriter io.Writer) (b *LiteBookie) {
	b = &LiteBookie{}
	b.elog = log.New(eWriter, "[Bookie] ", log.Lmicroseconds|log.Lshortfile)
	b.tlog = log.New(tWriter, "[T] ", log.Lmicroseconds)

	return b
}

func (b *LiteBookie) Open(path string) (err error) {
	b.db, err = sql.Open("sqlite3", path)
	if err != nil {
		b.elog.Println(err)
	}
	return err
}

func (b *LiteBookie) UserRegister(user, email, password string) (err error) {
	tx, err := b.db.Begin()
	if err != nil {
		b.elog.Println(err)
		return errors.New("Database failure")
	}

	sql := "INSERT INTO user(name, email, password) VALUES(?,?,?);"
	stmt, err := tx.Prepare(sql)
	if err != nil {
		b.elog.Println(err)
		return errors.New("Database failure")
	}
	defer stmt.Close()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		b.elog.Println(err)
		return errors.New("Temporary Registration Failure")
	}

	_, err = stmt.Exec(user, email, hash)
	if err != nil {
		b.elog.Println(err)
		return errors.New("Username or Email address not unique")
	}
	tx.Commit()

	return nil
}

func (b *LiteBookie) UserLogin(user, password string) {
}

func (b *LiteBookie) EventCreate() {
}

func (b *LiteBookie) EventOpen() {
}

func (b *LiteBookie) EventClose() {
}

func (b *LiteBookie) EventCancel() {
}

func (b *LiteBookie) EventList() {
}

func (b *LiteBookie) EventBet() {
}

func (b *LiteBookie) StreamCreate() {
}

func (b *LiteBookie) StreamList() {
}
