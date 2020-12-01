package database

import "github.com/ednailson/serasa-challenge/domain"

type Database interface {
	ReadByDocument(document string) ([]domain.Negativation, error)
	Save(negativation domain.Negativation) (string, error)
}
