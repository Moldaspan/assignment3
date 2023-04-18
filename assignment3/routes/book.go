package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moldaspan/assignment3/controllers"
)

func BookRoute(r *gin.Engine) {
	r.GET("books/", controllers.GetBooks)
	r.GET("books/:id", controllers.GetBook)
	r.GET("books/search", controllers.SearchBookByTitle)
	r.GET("books/sort", controllers.SortedBooks)
	r.POST("books/", controllers.CreateBook)
	r.PUT("books/:id", controllers.UpdateBook)
	r.DELETE("books/:id", controllers.DeleteBook)
}
