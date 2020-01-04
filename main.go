//https://medium.com/baixada-nerd/criando-um-crud-simples-em-go-3640d3618a67

package main

import(
    "database/sql"
    "net/http"
    "text/template"

    _ "github.com/go-sql-driver/mysql"
)

type Names struct{
Id    int
Name  string
Email string

}

func dbConn()(db *sql.DB){
    dbDriver := "mysql"
    dbUser :=""
    dbPass :=""
    dbName :=""
    
    db, err := sql.OPen(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())

    
    }
    return db
}

func main(){
    http.HandleFunc("/",Index)
    http.HandleFunc("/show",Show)
    htttp.HandleFunc("/new",New)
    http.HandleFunc("/edit"Edit)

    http.HandleFunc("/insert",insert)
    http.HandleFunc("/update",Update)
    http.HandleFunc("/delete",Delete)
    
    http.ListenAndServe(":9000",nil)

var tmpl = template.Must(templete.ParseGlob("tmpl/*"))


}

}




    
