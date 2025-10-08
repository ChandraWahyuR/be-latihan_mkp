package model

import "time"

type GetJadwalTayang struct {
	ID         string     `json:"id"`
	MovieID    string     `json:"movie_id"`
	NameStudio string     `json:"name_studio"`
	Starting   *time.Time `json:"starting"`
	Ending     *time.Time `json:"ending"`
}

type CreateJadwalTayang struct {
	ID         string     `json:"id"`
	MovieID    string     `json:"movie_id"`
	NameStudio string     `json:"name_studio"`
	Starting   *time.Time `json:"starting"`
	Ending     *time.Time `json:"ending"`
}

type EditJadwalTayang struct {
	ID         string     `json:"id"`
	MovieID    string     `json:"movie_id"`
	NameStudio string     `json:"name_studio"`
	Starting   *time.Time `json:"starting"`
	Ending     *time.Time `json:"ending"`
}
