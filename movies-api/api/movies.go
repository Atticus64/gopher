package movies

import (
	"database/sql"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	"net/http"
	"encoding/json"
	"log"
)

type Movie struct {
	Name string
	Genre string
	id int
}

func GetRouter() *http.ServeMux {
	var MoviesRouter = http.NewServeMux()
	MoviesRouter.HandleFunc("POST /movies", CreateMovie)
	return MoviesRouter
}

func CreateMovie(w http.ResponseWriter, req *http.Request) {
	var mov Movie

	err := json.NewDecoder(req.Body).Decode(&mov)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mov.Name == "" || mov.Genre == "" {
		log.Println("Empty")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing field Genre"))
		return
	} else {
		log.Printf("name: %s, genre: %s \n", mov.Name, mov.Genre)
	}
	db, err := sql.Open("sqlite3", "file:movie.db")
	if (err != nil) {
		log.Printf(err.Error())
		return
	}


	_, error := db.Exec(`insert into movie(name, genre) values(?, ?)`, mov.Name, mov.Genre)
	if (error != nil) {
		log.Printf(error.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Movie created"))
}



