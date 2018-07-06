package sql

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-07-06

import (
	"database/sql"

	"github.com/belfinor/Helium/log"
	_ "github.com/lib/pq"
)

type Row = *sql.Rows

type DB struct {
	dbh *sql.DB
	tx  *sql.Tx
}

type ROW_CALLBACK func(r Row) error

func New(conf *Config) *DB {

	db, err := sql.Open(conf.Driver, conf.Connect)
	if err != nil {
		log.Error("connect usersdb error: " + err.Error())
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Error("database ping error:" + err.Error())
		db.Close()
		return nil
	}

	tr, e := db.Begin()
	if e != nil {
		log.Error("begin transaction error: " + e.Error())
		db.Close()
		return nil
	}

	return &DB{dbh: db, tx: tr}
}

func (db *DB) Commit() {
	e := db.tx.Commit()

	db.tx, e = db.dbh.Begin()
	if e != nil {
		log.Error("begin transaction error: " + e.Error())
		panic(e)
	}
}

func (db *DB) Rollback() {
	e := db.tx.Rollback()

	db.tx, e = db.dbh.Begin()
	if e != nil {
		log.Error("begin transaction error: " + e.Error())
		panic(e)
	}
}

func (db *DB) Close() {
	if db.tx != nil {
		db.tx.Rollback()
	}
	db.dbh.Close()
}

func (db *DB) Exec(query string, args ...interface{}) {
	if _, err := db.tx.Exec(query, args...); err != nil {
		log.Error(err)
		panic(err)
	}
}

func (db *DB) SelectInt(query string, args ...interface{}) (int64, bool) {
	var val int64
	if err := db.tx.QueryRow(query, args...).Scan(&val); err != nil {
		if err != sql.ErrNoRows {
			log.Error(err)
			panic(err)
		}
		return 0, false
	}
	return val, true
}

func (db *DB) SelectFloat(query string, args ...interface{}) (float64, bool) {
	var val float64
	if err := db.tx.QueryRow(query, args...).Scan(&val); err != nil {
		if err != sql.ErrNoRows {
			log.Error(err)
			panic(err)
		}
		return 0, false
	}
	return val, true
}

func (db *DB) SelectString(query string, args ...interface{}) (string, bool) {
	var val string
	if err := db.tx.QueryRow(query, args...).Scan(&val); err != nil {
		if err != sql.ErrNoRows {
			log.Error(err)
			panic(err)
		}
		return "", false
	}
	return val, true
}

func (db *DB) SelectRow(query string, args []interface{}, res ...interface{}) bool {
	if err := db.tx.QueryRow(query, args...).Scan(res...); err != nil {
		if err != sql.ErrNoRows {
			log.Error(err)
			panic(err)
		}
		return false
	}

	return true
}

func (db *DB) Select(query string, args []interface{}, fn ROW_CALLBACK) {
	rows, err := db.tx.Query(query, args...)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = fn(rows); err != nil {
			log.Error(err)
			panic(err)
		}
	}
}
