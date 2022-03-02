package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

type Album struct {
    ID     int64 `json:"id"`
    Title  string `json:"title"`
    Artist string `json:"artist"`
    Price  float32 `json:"price"`
}

var db *sql.DB

func main() {
    
    var err error
    db, err = sql.Open("mysql", "root:mayank99@tcp(127.0.0.1:3306)/recordings")
    if err != nil {
        panic(err)
    }
    fmt.Print("Successfully Created")
    defer db.Close()

	albID, err := addAlbum(Album{
        Title:  "The Modern Sound of Betty Carter",
        Artist: "Betty Carter",
        Price:  49.99,
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("ID of added album: %v\n", albID)


    router := httprouter.New()
	router.GET("/",index)
    log.Fatal(http.ListenAndServe(":8080",router))

}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!")
    w.Header().Set("Content-Type","application/json")
    albums, err := albumsByArtist(db, "John Coltrane")
    if err != nil {
        log.Fatal(err)
    }
    // fmt.Printf("Albums found: %v\n", albums)
    db.Close()
    // m := Album{1, "Hello", "mayan",29.99}
    json.NewEncoder(w).Encode(albums)
    
}


func addAlbum(alb Album) (int64, error) {
    result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    return id, nil
}

func albumsByArtist(db *sql.DB, name string) ([]Album, error) {

    var albums []Album
    fmt.Printf("%v", db)
    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }

    defer rows.Close()
    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
        albums = append(albums, alb)
    }

    return albums, nil
}


// package main

// import (
//     "database/sql"
//     "fmt"
//     "log"
//     "os"

//     "github.com/go-sql-driver/mysql"
// )
// var db *sql.DB

// func main() {
//     cfg := mysql.Config{
//         User:   os.Getenv("root"),
//         Passwd: os.Getenv("mayank99"),
//         Net:    "tcp",
//         Addr:   "127.0.0.1:3306",
//         DBName: "recordings",
// 		AllowNativePasswords: true,
//     }
//     var err error
//     db, err = sql.Open("mysql", cfg.FormatDSN())
//     if err != nil {
//         log.Fatal(err)
//     }

//     pingErr := db.Ping()
//     if pingErr != nil {
//         log.Fatal(pingErr)
//     }
//     fmt.Println("Connected!")
// }

