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

func NewPostRoute(timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	r := repository.NewPostRepository(&db, domain.CollectionPosts)
	s := service.NewPostService(r, timeout)
	c := controller.NewPostController(s)

	group.GET("/post", c.GetPost)
	group.POST("/post", c.CreatePost)
	group.DELETE("/post", c.DeletePost)
	group.GET("/post/search", c.GetPostList)
}
