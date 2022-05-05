package usecase

import (
	"server/entities"
)

type PostInteractor struct {
	PostRepository PostRepository
}

func (interactor *PostInteractor) Add(post entities.Post) error {
	err := interactor.PostRepository.Store(post)
	return err
}

func (interactor *PostInteractor) Posts() (posts entities.Posts ,err error) {
	posts, err = interactor.PostRepository.FindAll()
	return posts, err
}

func (interactor *PostInteractor) Remove(id int64) error {
	err := interactor.PostRepository.Remove(id)
	return err
}

func (interactor *PostInteractor) Update(post entities.Post) error {
	err := interactor.PostRepository.Update(post)
	return err
}