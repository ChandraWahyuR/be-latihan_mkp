package entity

import "time"

type JadwalTayang struct {
	ID         string
	MovieID    string
	NameStudio string
	Starting   *time.Time
	Ending     *time.Time
}
