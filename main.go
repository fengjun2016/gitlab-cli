package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"golang.org/x/sync/errgroup"
	"time"
	"log"
)

var (
	g errgroup.Group
)

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
			},
		)
	})

	return e
}

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	server := &http.Server{
		Addr: ":8081",
		Handler: router01(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}