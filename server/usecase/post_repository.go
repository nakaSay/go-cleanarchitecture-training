package usecase

import (
	"server/entities"
)

type PostRepository interface {
	Store(entities.Post) error
	FindAll() (entities.Posts, error)
	Remove(int64) error
	Update(entities.Post) error
}