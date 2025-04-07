package internal

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateAuthorHttp(usecase *CreateAuthor, log *Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var author InputCreateAuthor
		ctx := c.Request.Context()
		err := c.Bind(&author)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.DebugWithFields("user request", map[string]interface{}{
			"name":  author.Name,
			"email": author.Email,
			"phone": author.Phone,
		})
		err = usecase.Execute(ctx, author)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Author created"})
	}
}

func CreatePostHttp(usecase *PostCreate, log *Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var post InputPostCreate
		var ctx = c.Request.Context()
		err := c.Bind(&post)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		authorID := c.Param("author_id")
		log.DebugWithFields("user request", map[string]interface{}{
			"title":     post.Title,
			"content":   post.Content,
			"author_id": authorID,
		})
		authorIDInt, err := strconv.Atoi(authorID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = usecase.Execute(ctx, post, authorIDInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Post created"})
	}
}

func FindPostByAuthorHttp(usecase *FindPostByAuthor) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx = c.Request.Context()
		authorID := c.Param("author_id")
		authorIDInt, err := strconv.Atoi(authorID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		output, err := usecase.Execute(ctx, authorIDInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func FindPostHttp(usecase *FindPost) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx = c.Request.Context()
		output, err := usecase.Execute(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func EndCompetitionHttp(usecase *EndCompetition) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		winners, err := usecase.Execute(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Final competition executed", "winners": winners.Winners})
	}
}

func FindPostByIDHttp(usecase *FindPostByID) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx = c.Request.Context()
		postID := c.Param("post_id")
		postIDInt, err := strconv.Atoi(postID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		authorID := c.Param("author_id")
		authorIDInt, err := strconv.Atoi(authorID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		output, err := usecase.Execute(ctx, authorIDInt, postIDInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
