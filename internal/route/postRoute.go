package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prppoomw/blog-api/internal/controller"
	"github.com/prppoomw/blog-api/internal/domain"
	"github.com/prppoomw/blog-api/internal/repository"
	"github.com/prppoomw/blog-api/internal/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewPostRoute(timeout time.Duration, db mongo.Database, privateGroup *gin.RouterGroup, publicGroup *gin.RouterGroup) {
	r := repository.NewPostRepository(&db, domain.CollectionPosts)
	s := service.NewPostService(r, timeout)
	c := controller.NewPostController(s)

	publicGroup.GET("/post", c.GetPost)
	privateGroup.POST("/post", c.CreatePost)
	privateGroup.DELETE("/post", c.DeletePost)
	publicGroup.GET("/post/search", c.GetPostList)
}
