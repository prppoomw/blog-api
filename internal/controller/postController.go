package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prppoomw/blog-api/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PostController struct {
	postService domain.PostUsecase
}

func NewPostController(postService domain.PostUsecase) *PostController {
	return &PostController{postService: postService}
}

func (ctrl *PostController) GetPost(c *gin.Context) {
	slug := c.Query("slug")
	post, err := ctrl.postService.GetPost(c, slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	if post == nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Post Not Found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (ctrl *PostController) CreatePost(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not Authenticated"})
		return
	}
	var post domain.Post
	err := c.ShouldBind(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	post.UserId = userId
	res, e := ctrl.postService.CreatePost(c, &post)
	if e != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: e.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (ctrl *PostController) DeletePost(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not Authenticated"})
		return
	}
	postId := c.Query("id")
	id, err := bson.ObjectIDFromHex(postId)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	err = ctrl.postService.DeletePost(c, id, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Post has been deleted")
}

func (ctrl *PostController) GetPostList(c *gin.Context) {
	var queryReq domain.PostListQueryRequest
	e := c.ShouldBind(&queryReq)
	if e != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: e.Error()})
		return
	}
	res, err := ctrl.postService.GetPostList(c, &queryReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
