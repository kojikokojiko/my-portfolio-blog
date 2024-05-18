package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	apiKey := os.Getenv("MICROCMS_API_KEY")
	endpoint := os.Getenv("MICROCMS_ENDPOINT")

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	r.GET("/api/blogs", func(c *gin.Context) {
		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/blogs", endpoint), nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create request"})
			return
		}
		req.Header.Add("X-MICROCMS-API-KEY", apiKey)

		response, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch articles"})
			return
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			c.JSON(response.StatusCode, gin.H{"error": "Error from microCMS"})
			return
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read response body"})
			return
		}

		var result struct {
			Contents []Article `json:"contents"`
		}
		if err := json.Unmarshal(body, &result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to parse response body"})
			return
		}

		// フロントエンドのインターフェースに合わせて必要なフィールドを抽出
		articles := make([]Article, len(result.Contents))
		for i, article := range result.Contents {
			articles[i] = Article{
				ID:      article.ID,
				Title:   article.Title,
				Content: article.Content,
			}
		}

		c.JSON(http.StatusOK, articles)
	})

	r.GET("/api/blogs/:id", func(c *gin.Context) {
		id := c.Param("id")
		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/blogs/%s", endpoint, id), nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create request"})
			return
		}
		req.Header.Add("X-API-KEY", apiKey)

		response, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch article"})
			return
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			c.JSON(response.StatusCode, gin.H{"error": "Error from microCMS"})
			return
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read response body"})
			return
		}

		var article Article
		if err := json.Unmarshal(body, &article); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to parse response body"})
			return
		}

		c.JSON(http.StatusOK, Article{
			ID:      article.ID,
			Title:   article.Title,
			Content: article.Content,
		})
	})

	r.Run(":8080")
}
