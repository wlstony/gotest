package database

import (
	"database/sql"
	"errors"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var _db *sql.DB

func init() {
	//sql.Register("mysql", mysql.MySQLDriver{})
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	_db = db
}

type Metadata struct {
	Name    string
	Version int
	Size    int64
	Hash    string
}

func getMetadata(name string, versionId int) (meta Metadata, e error) {
	rows, e := _db.Query("select * from objects where name=? and version=?", name, versionId)
	if e != nil {
		return
	}
	for rows.Next() {
		e = rows.Scan(&meta.Name, &meta.Version, &meta.Size, &meta.Hash)
		if e != nil {
			return
		}
	}
	return
}

type hit struct {
	Source Metadata `json:"_source"`
}
type searchResult struct {
	Hits struct {
		Total int
		Hits  []hit
	}
}

func SearchLatestVersion(name string) (meta Metadata, e error) {
	rows, err := _db.Query("select * from objects where `name`=?", name)
	if err != nil {
		return meta, err
	}
	for rows.Next() {
		re := rows.Scan(&meta.Name, &meta.Version, &meta.Size, &meta.Hash)
		if re != nil {
			return meta, re
		}
	}
	return
}
func GetMetadata(name string, version int) (meta Metadata, e error) {
	if version == 0 {
		return SearchLatestVersion(name)
	}
	return getMetadata(name, version)
}
func PutMetadata(name string, version int, size int64, hash string) error {
	r, e := _db.Exec("insert into objects(`name`, `version`, `size`, `hash`) values(?, ?, ?, ?) ", name, version, size, hash)
	if e != nil {
		return e
	}
	af, er := r.RowsAffected()
	if af <= 0 {
		return errors.New(er.Error())
	}
	return nil
}

func AddVersion(name, hash string, size int64) error {
	version, e := SearchLatestVersion(name)
	if e != nil {
		return e
	}
	return PutMetadata(name, version.Version+1, size, hash)
}