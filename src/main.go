package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// Rectangle represents a rectangle with its coordinates and dimensions
type Rectangle struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

// Intersection represents an intersecting rectangle with timestamp
type Intersection struct {
	Rectangle
	Time string `json:"time"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./rectangles.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the intersections table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS intersections (
		x INTEGER,
		y INTEGER,
		width INTEGER,
		height INTEGER,
		time TEXT
	)`)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", intersectHandler).Methods("POST")
	router.HandleFunc("/", getIntersectionsHandler).Methods("GET")

	server := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting web server on port 8080...")
	log.Fatal(server.ListenAndServe())
}

func intersectHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received POST request")

	var data struct {
		Main  Rectangle   `json:"main"`
		Input []Rectangle `json:"input"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Failed to decode JSON:", err)
		return
	}

	for _, rect := range data.Input {
		if intersects(data.Main, rect) {
			intersection := Intersection{
				Rectangle: rect,
				Time:      time.Now().Format("2006-01-02 15:04:05"),
			}
			// Save intersection to database
			saveIntersection(intersection)
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func getIntersectionsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received GET request")

	// Retrieve intersections from database
	intersections := getIntersections()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(intersections)
}

func intersects(rect1, rect2 Rectangle) bool {
	if rect1.X >= rect2.X+rect2.Width || rect2.X >= rect1.X+rect1.Width {
		return false
	}
	if rect1.Y >= rect2.Y+rect2.Height || rect2.Y >= rect1.Y+rect1.Height {
		return false
	}
	return true
}

func saveIntersection(intersection Intersection) {
	stmt, err := db.Prepare("INSERT INTO intersections(x, y, width, height, time) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		intersection.X,
		intersection.Y,
		intersection.Width,
		intersection.Height,
		intersection.Time,
	)
	if err != nil {
		log.Println(err)
		return
	}
}

func getIntersections() []Intersection {
	rows, err := db.Query("SELECT x, y, width, height, time FROM intersections")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	intersections := []Intersection{}

	for rows.Next() {
		var intersection Intersection
		err := rows.Scan(
			&intersection.X,
			&intersection.Y,
			&intersection.Width,
			&intersection.Height,
			&intersection.Time,
		)
		if err != nil {
			log.Println(err)
			continue
		}

		intersections = append(intersections, intersection)
	}

	return intersections
}
