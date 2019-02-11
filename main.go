package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/trytwice/gushiwen/pkg/db"
	"github.com/trytwice/gushiwen/pkg/spider"
	"github.com/trytwice/gushiwen/router"
	"github.com/urfave/cli"
)

var (
	baseURLPoet   = "https://so.gushiwen.org/authors/Default.aspx?p=%d"
	baseURLPoetry = "https://www.gushiwen.org/shiwen/default.aspx?page=%d&type=0&id=0"
)

func main() {
	app := cli.NewApp()
	app.Name = "gushiwen"
	app.Usage = "gushiwen api"
	app.Version = "1.1.1"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "table, t",
			Usage: "creat table if not exists.",
		},
		cli.BoolFlag{
			Name:  "poet",
			Usage: "get poet data from website.",
		},
		cli.BoolFlag{
			Name:  "poetry",
			Usage: "get poetry data from website.",
		},
		cli.BoolFlag{
			Name:  "run",
			Usage: "run web server.",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.Bool("table") {
			err := db.CreatDB()
			if err != nil {
				return err
			}
			return nil
		}
		if c.Bool("poet") {
			for i := 1; i <= 1000; i++ {
				url := fmt.Sprintf(baseURLPoet, i)
				err := spider.GetPoet(url)
				if err != nil {
					fmt.Println(err)
				}
			}
			return nil
		}
		if c.Bool("poetry") {
			for i := 1; i <= 1000; i++ {
				url := fmt.Sprintf(baseURLPoetry, i)
				err := spider.GetPoetry(url)
				if err != nil {
					fmt.Println(err)
				}
			}
			return nil

		}
		if c.Bool("run") {
			routers := router.Router{}
			g := gin.New()
			routers.InitRouter(g)
			err := g.Run(":18080")
			return err
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
