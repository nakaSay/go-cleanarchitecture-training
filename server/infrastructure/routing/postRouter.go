package routing

import (
	"github.com/labstack/echo/v4"
	"server/interfaces/controllers"
	"net/http"
	"server/infrastructure/sqlhandler"
)

func SetPostRouting(e *echo.Echo) {

	postController := controllers.NewPostController(sqlhandler.NewSqlHandler())

	e.POST("/createPost", func(c echo.Context) error {
		err := postController.Create(c)
		if err != nil {
			return c.JSON(500, err)
		}
		return c.JSON(http.StatusOK, "OK!")
	})

	e.GET("/posts", func(c echo.Context) error {
		docs, err := postController.Index()
		if err != nil {
			return c.JSON(500, docs)
		}
		return c.JSON(http.StatusOK, docs)
	})

	e.GET("/removePost", func(c echo.Context) error {
		err := postController.Remove(c)
		if err != nil {
			return c.JSON(500, err)
		}
		return c.JSON(http.StatusOK, "OK!")
	})

	e.POST("/updatePost", func(c echo.Context) error {
		err := postController.Update(c)
		if err != nil {
			return c.JSON(500, err)
		}
		return c.JSON(http.StatusOK, "OK!")
	})
}