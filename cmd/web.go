package cmd

import (
	v1 "github.com/alextuan1024/requiem/web/v1"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"net/http"
)

var web *cli.Command

var r *gin.Engine

func init() {
	r = gin.New()

	web = &cli.Command{
		Name:      "web",
		Usage:     "start requiem web server",
		UsageText: "requiem web --port [port]",
		Action: func(c *cli.Context) error {
			r.Use(gin.Logger())
			r.Use(gin.Recovery())

			r.GET("/ping", func(c *gin.Context) {
				c.String(http.StatusOK, "pong")
			})
			v1.RegisterHandlers(r, v1.GetApi())
			port := c.String("port")
			return r.Run(port)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "port",
				Aliases: []string{"p"},
				EnvVars: []string{"HTTP_LISTEN"},
				Usage:   "specify port using `--port`",
			},
		},
	}
	app.Commands = append(app.Commands, web)
}
