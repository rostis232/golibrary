package models

import (
	"github.com/lib/pq"
)

type LibraryItem struct{
	ID int64 `db:"id"`
	Title string `db:"title"`
	ShortDescEN string `db:"short_desc_en"`
	ShortDescUK string `db:"short_desc_uk"`
	FullDescriptionEN string `db:"description_en"`
	FullDescriptionUK string `db:"description_uk"`
	URL string `db:"url"`
	Type int64 `db:"type"`
	Difficulty int64 `db:"difficulty"`
	Language pq.Int64Array `db:"language"`
}

type Type struct{
	ID int64 `db:"id"`
	TitleEN string `db:"title_en"`
	TitleUK string `db:"title_uk"`
}

type Difficulty struct{
	ID int64 `db:"id"`
	TitleEN string `db:"title_en"`
	TitleUK string `db:"title_uk"`
}

type Language struct{
	ID int64 `db:"id"`
	TitleEN string `db:"title_en"`
	TitleUK string `db:"title_uk"`
}