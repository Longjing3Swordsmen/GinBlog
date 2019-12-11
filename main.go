package main

import (
	"GinBlog/db"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init(){
	log.SetFlags(log.LstdFlags|log.Lshortfile)
}

func main() {
	app := &cli.App{
		Name:  "gblog",
		Usage: "blog deploy and remove",
		Action: func(c *cli.Context) error {
			fmt.Println("blog deploy and remove")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "init blog app",
				Action: func(c *cli.Context) error {
					err := db.CreateDB()
					db.InitDB()
					return err
				},
			},
			{
				Name:    "uninstall",
				Aliases: []string{"u"},
				Usage:   "uninstall blog app",
				Action: func(c *cli.Context) error {
					err := db.RecoverEnv()
					return err
				},
			},
			{
				Name: "run",
				Usage: "run blog",
				Action: func(c *cli.Context) error {
					startBlog()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func startBlog() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	srv := http.Server{
		Addr: ":8000",
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}
