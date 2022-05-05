package database

import (
	"server/entities"
)

type PostRepository struct {
	PostSqlHandler
}

func (repo *PostRepository) Store(post entities.Post) error {
	err := repo.AddPost("Post", post)
	return err
}

func (repo *PostRepository) FindAll() (posts entities.Posts, err error) {
    posts, err = repo.GetAllPosts("Post")
    return posts, err
}

func (repo *PostRepository) Remove(id int64) error {
    err := repo.Delete("Post", id)
    return err
}

func (repo *PostRepository) Update(post entities.Post) error {
    err := repo.UpdatePost("Post", post)
    return err
}