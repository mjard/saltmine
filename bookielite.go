package main

import (
	"database/sql"
	"errors"
	"io"
	"log"

	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"code.google.com/p/go.crypto/bcrypt"
)

const (
	EventClosed   = 0
	EventOpened   = 1
	EventFinished = 2
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

func (b *LiteBookie) sqlexec(query string, args ...interface{}) (result sql.Result, err error) {
	stmt, err := b.db.Prepare(query)
	if err != nil {
		b.elog.Println(err)
		return nil, errors.New("Temporary database issue")
	}
	defer stmt.Close()

	result, err = stmt.Exec(args...)
	if err != nil {
		b.elog.Println(err)
		return nil, errors.New("Unable to complete")
	}
	return result, err
}

func (b *LiteBookie) Open(path string) (err error) {
	b.db, err = sql.Open("sqlite3", path)
	if err != nil {
		b.elog.Println(err)
	}
	return err
}

func (b *LiteBookie) UserRegister(user, email, password string, balance int) (err error) {
	const query = "INSERT INTO user(name, email, password, balance) VALUES(?,?,?,?);"

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		b.elog.Println(err)
		return errors.New("Temporary Registration Failure")
	}

	_, err = b.sqlexec(query, user, email, hash, balance)
	// could be a lie, fix this by recording what stage the error occurred
	if err != nil {
		return errors.New("Username or Email address not unique")
	}

	return nil
}

func (b *LiteBookie) UserLogin(user, password string) (err error) {
	const query = "SELECT password FROM user WHERE name=?;"
	const update = "UPDATE user SET last_login=current_timestamp WHERE name=?;"

	var hash []byte
	err = b.db.QueryRow(query, user).Scan(&hash)
	if err != nil {
		b.elog.Println(err)
		return errors.New("Invalid User")
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		b.elog.Println(err)
		return errors.New("Invalid Password")
	}

	_, err = b.sqlexec(update, user)
	if err != nil {
		return errors.New("Temporarly Unable to Login")
	}

	return err
}

func (b *LiteBookie) EventCreate(streamId int) (err error) {
	const query = `INSERT INTO event(stream, status) VALUES(?, 
        (SELECT id FROM eventstatus WHERE code=?));`

	_, err = b.sqlexec(query, streamId, EventOpened)
	if err != nil {
		b.elog.Println(err)
		return errors.New("Unable to create event")
	}

	return err
}

func (b *LiteBookie) EventUpdateStatus(streamId, code int) (err error) {
	const query = `UPDATE event SET status=
        (SELECT id FROM eventstatus WHERE code=?) WHERE stream=?`

	_, err = b.sqlexec(query, code, streamId)
	if err != nil {
		b.elog.Println(err)
		return errors.New("Unable to update event")
	}

	return err
}

func (b *LiteBookie) EventOpen(streamId int) {
	b.EventUpdateStatus(streamId, EventOpened)
}

func (b *LiteBookie) EventFinalize(streamId int) {
	b.EventUpdateStatus(streamId, EventFinished)
}

func (b *LiteBookie) EventCancel(streamId int) {
	b.EventUpdateStatus(streamId, EventFinished)
}

func (b *LiteBookie) EventList() {
}

// this is going to take a slice of Rankings eventually
func (b *LiteBookie) EventBet(eventId, participantId, rank, bet int) {
}

func (b *LiteBookie) StreamCreate() {
}

func (b *LiteBookie) StreamList() {
}
