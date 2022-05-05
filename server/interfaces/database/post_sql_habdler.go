package database

import (
	"server/entities"
)

type PostSqlHandler interface {
	AddPost(string, entities.Post) error
	GetAllPosts(string) (entities.Posts, error)
	Delete(string, int64) error
	UpdatePost(string, entities.Post) error
}