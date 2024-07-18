package models

import (
	"github.com/lib/pq"
)

type LibraryItem struct {
	ID              int64         `db:"id"`
	Title           string        `db:"title"`
	ShortDesc       string        `db:"short_desc"`
	FullDescription string        `db:"description"`
	URL             string        `db:"url"`
	Type            int64         `db:"type"`
	Difficulty      int64         `db:"difficulty"`
	Language        pq.Int64Array `db:"language"`
}

type Type struct {
	ID    int64  `db:"id"`
	Title string `db:"title"`
}

type Difficulty struct {
	ID    int64  `db:"id"`
	Title string `db:"title"`
}

type Language struct {
	ID    int64  `db:"id"`
	Title string `db:"title"`
}
