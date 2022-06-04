package lib

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


const create_sql string = `CREATE TABLE IF NOT EXISTS  logs  (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        uname TEXT,
        ecount INTEGER
    );`

const dbname string = "./mydb.db"

func getCon() *sql.DB {
    
     db , err := sql.Open("sqlite3" , dbname)
     if err != nil{
        log.Fatalln(err)
     }
     return db


}

func CreateDatabase(){
    
    db , err := sql.Open("sqlite3" , dbname)

    if err != nil{
        fmt.Println("CREATE DB ERROR ->")
        log.Fatalln(err)
    }

    if _ , err := db.Exec(create_sql); err!=nil{
        fmt.Println("CREATE DB EXEC ERROR =>")
        log.Fatalln(err.Error())
    }
    
}

func CreateNewLog(dlog *Dlog){
    
    db := getCon()

    raw_sql , err:= db.Prepare("INSERT INTO logs(name,uname,ecount) VALUES(?,?,?)")

    if err != nil{
        fmt.Println("CREATE LOG ERROR->")
        log.Fatalln(err)
    }


   
   raw_sql.Exec(dlog.Name , dlog.Uname , len(dlog.Posts))

}
