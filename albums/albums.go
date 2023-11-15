package albums

import (
	"errors"
)

type album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAll() []album {
	return albums
}

func getAlbumById(id int) (*album, error) {
	for _, a := range albums {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("album not found")
}
