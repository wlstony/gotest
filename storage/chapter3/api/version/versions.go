package version

import (
	"database/sql"
	"encoding/json"
	"github.com/api/database"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strings"
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
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//from := 0
	//size := 1000
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	ms := make([]database.Metadata, 0)
	s := 0
	//for {
		rows, e := _db.Query("select * from objects where name=?", name)
		if e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		for rows.Next() {
			ms = append(ms, database.Metadata{})
			se := rows.Scan(&ms[s].Name, &ms[s].Version, &ms[s].Size, &ms[s].Hash)
			if se != nil {
				log.Println(e)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			b, _ := json.Marshal(ms[s])
			w.Write(b)
			w.Write([]byte("\n"))
			s++
		}
		//from += size
	//}
}
