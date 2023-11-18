package albums

import (
	"errors"
)

type Album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var db = []Album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAll() []Album {
	return db
}

func nextId() int {
	return len(db) + 1
}

func GetById(id int) (*Album, error) {
	for _, a := range db {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("album not found")
}

func CreateAlbum(newAlbum Album) (*Album, string) {
	var errors string

	if newAlbum.Title == "" {
		errors += "Title can't be empty"
	}

	if newAlbum.Artist == "" {
		errors += "Artist can't be empty"
	}

	if errors != "" {
		return nil, errors
	}

	newAlbum.ID = nextId()
	db = append(db, newAlbum)

	return &newAlbum, ""
}
