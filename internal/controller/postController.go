package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imagekit-developer/imagekit-go"
	"github.com/prppoomw/blog-api/internal/config"
	"github.com/prppoomw/blog-api/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PostController struct {
	postService domain.PostUsecase
	cfg         *config.Config
}

func NewPostController(postService domain.PostUsecase, cfg *config.Config) *PostController {
	return &PostController{
		postService: postService,
		cfg:         cfg,
	}
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
	pageStr := c.Query("page")
	limitStr := c.Query("limit")
	queryReq.Category = c.Query("category")
	queryReq.Author = c.Query("author")
	queryReq.Search = c.Query("search")

	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid page number"})
			return
		}
		queryReq.Page = page
	} else {
		queryReq.Page = 1
	}

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid limit value"})
			return
		}
		queryReq.Limit = limit
	} else {
		queryReq.Limit = 6
	}

	res, err := ctrl.postService.GetPostList(c, &queryReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (ctrl *PostController) Upload(c *gin.Context) {
	ik := imagekit.NewFromParams(imagekit.NewParams{
		PrivateKey:  ctrl.cfg.ImgkitPrivateKey,
		PublicKey:   ctrl.cfg.ImgkitPublicKey,
		UrlEndpoint: ctrl.cfg.ImgkitUrlEndpoint,
	})
	resp := ik.SignToken(imagekit.SignTokenParam{})
	log.Println(resp)
	c.JSON(http.StatusOK, gin.H{
		"signature": resp.Signature,
		"expire":    resp.Expires,
		"token":     resp.Token,
	})
}
