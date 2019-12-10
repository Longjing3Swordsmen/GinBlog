package main

import (
	"GinBlog/db"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
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
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
