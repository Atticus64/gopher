package main

import (
	"fmt"
	"database/sql"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	"net/http"
	"api.movies.dev/api"
)

func health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusTeapot)
	w.Write([]byte("OK!"))
}

func main() {
	db, _ := sql.Open("sqlite3", "file:movie.db")
	_, err := db.Exec(`create table if not exists movie(
		id integer primary key,
		name text not null,
		genre text not null
	)`)

	router := http.NewServeMux()

	router.HandleFunc("GET /health", health)
	router.Handle("/api/", http.StripPrefix("/api", movies.GetRouter()))

	server := http.Server {
		Addr: ":8080",
		Handler: router,
	}

	fmt.Printf("Listening at http://localhost%s\n", server.Addr)
	server.ListenAndServe()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	// res, err := db.Exec(`insert into movie(name, genre) values(?, ?)`, "IronMan", "Action")
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// 	return
	// }
	// count, err := res.RowsAffected()


}

