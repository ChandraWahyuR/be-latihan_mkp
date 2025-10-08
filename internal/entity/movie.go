package entity

type Movie struct {
	ID    string
	Title string
}

type MovieAndStudio struct {
	ID           string
	Title        string
	JadwalTayang []JadwalTayang
}
