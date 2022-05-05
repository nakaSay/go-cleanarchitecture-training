package controllers

import (
	"server/usecase"
	"server/interfaces/database"
	"server/entities"
	"fmt"
	"strconv"
)

type PostController struct {
	Interactor usecase.PostInteractor
}

func NewPostController(postSqlhandler database.PostSqlHandler) *PostController {
	return &PostController{
		Interactor: usecase.PostInteractor{
			PostRepository: &database.PostRepository{
				PostSqlHandler: postSqlhandler,
			},
		},
	}
}

func (controller *PostController) Create(c Context) error {
	post := entities.Post {
		Text: c.FormValue("text"),
	}
	err := controller.Interactor.Add(post)
	if err != nil {
		fmt.Printf("Interactor Add() error: %s", err)
    }
	return err
}

func (controller *PostController) Index() (entities.Posts, error) {
	posts, err := controller.Interactor.Posts()
	if err != nil {
		fmt.Printf("Interactor Posts() error: %s", err)
	}
	return posts, err
}

func (controller *PostController) Remove(c Context) error {
	idQueryParam := c.QueryParam("id")
	id, _ := strconv.ParseInt(idQueryParam, 10, 64)
	err := controller.Interactor.Remove(id)
	if err != nil {
		fmt.Printf("Interactor Remove() error: %s", err)
	}
	return err
}

func (controller *PostController) Update(c Context) error {
	id, _ := strconv.ParseInt(c.FormValue("id"), 10, 64)
	post := entities.Post {
		ID: id,
		Text: c.FormValue("text"),
	}
	err := controller.Interactor.Update(post)
	if err != nil {
		fmt.Printf("Interactor Update() error: %s", err)
	}
	return err
}