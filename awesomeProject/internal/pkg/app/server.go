package app

import (
	"awesomeProject/internal/app/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Application struct {
	repo *repository.Repository
}

func NewApplication(dsn string) (*Application, error) {
	repo, err := repository.New(dsn) // str for bd
	if err != nil {
		return nil, err
	}
	return &Application{
		repo: repo,
	}, nil
}

func (a *Application) StartServer() {
	log.Println("Server start up")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		id := c.Query("id") // получаем из запроса query string

		if id != "" {
			log.Printf("id recived %s\n", id)
			/*intID, err := strconv.Atoi(id) // пытаемся привести это к чиселке
			if err != nil {                // если не получилось
				log.Printf("cant convert id %v", err)
				c.Error(err)
				return
			}

			fmt.Println("int id = ", intID)*/// выкидываем приведение к целому типу

			product, err := a.repo.GetProductByID(id)
			if err != nil { // если не получилось
				log.Printf("cant get product by id %v", err)
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"product_price": product.Price,
				"prodict_name":  product.Code,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("templates/*")

	r.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test.tmpl", gin.H{
			"title": "Main website",
			"test":  []string{"a", "b"},
		})
	})

	r.GET("/wl", func(c *gin.Context) {
		id := c.Query("itemCount")
		products, err := a.repo.GetProductBelowID(id)

		for _, currentProduct := range products {
			fmt.Println("debug products", currentProduct.ID, currentProduct.Code, currentProduct.Price, currentProduct.Image)
		}

		if err != nil {
			c.Error(err)
			return
		}

		c.HTML(http.StatusOK, "wl.tmpl", gin.H{
			"products": products,
		})
	})

	r.Static("/image", "./resources")
	r.Static("/styles", "./resources")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}
