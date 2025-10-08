package model

import "time"

type CreateMovie struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type EditMovie struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type GetMovieResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type MovieAndStudio struct {
	ID        string      `json:"id"`
	Title     string      `json:"title"`
	JamTayang []JamTayang `json:"jam_tayang"`
}

type JamTayang struct {
	MovieID    string     `json:"movie_id"`
	NameStudio string     `json:"name_studio"`
	Starting   *time.Time `json:"starting"`
	Ending     *time.Time `json:"ending"`
}
