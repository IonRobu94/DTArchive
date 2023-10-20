package models

type Album struct {
	ID     int    `db:"id"`
	BandID int    `db:"band"` // foreign key to Band
	Title  string `db:"title"`
	Year   int    `db:"year"`
}
