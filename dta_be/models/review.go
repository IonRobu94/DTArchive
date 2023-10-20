package models

type Review struct {
	ID      int     `db:"id"`
	AlbumID int     `db:"album"` // foreign key to Album
	Title   string  `db:"title"`
	Score   float64 `db:"score"`
	Content string  `db:"content"`
}
