package models

type Band struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Country string `db:"country"`
	Status  string `db:"status"`
	FromIn  int    `db:"from_in"`
	Genre   string `db:"genre"`
	Theme   string `db:"theme"`
	Active  string `db:"active"`
}
