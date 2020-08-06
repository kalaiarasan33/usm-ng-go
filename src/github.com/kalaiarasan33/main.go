package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-contrib/static"
	"github.com/kalaiarasan33/usm-ng-go/controllers"
	"github.com/urfave/cli"

	"github.com/kalaiarasan33/usm-ng-go/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	DEFAULT_PORT = 3000
)

type binaryFileSystem struct {
	fs http.FileSystem
}

func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {

	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := b.fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

func BinaryFileSystem(assetFS func() *assetfs.AssetFS) *binaryFileSystem {
	fs := assetFS()
	return &binaryFileSystem{
		fs,
	}
}

func startServer(port int) {
	r := gin.Default()
	r.Use(cors.Default())

	models.ConnectDataBase()

	r.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Welcome": "User Management API"})
	})

	r.GET("/users", controllers.Findusers)
	r.POST("/createuser", controllers.Createuser)
	r.GET("/users/:id", controllers.Finduser)
	r.PATCH("/updateusers/:id", controllers.Updateuser)
	r.DELETE("/deleteusers/:id", controllers.DeleteUser)

	r.Use(static.Serve("/", BinaryFileSystem(assetFS)))

	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Server started on port %d", port)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "User Management"
	app.Compiled = time.Now()
	app.Flags = []cli.Flag{
		&cli.IntFlag{
			Name:  "port, p",
			Usage: "Server port",
			Value: DEFAULT_PORT,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start server",
			Action: func(c *cli.Context) error {
				port := c.Int("port")
				if port == 0 {
					port = DEFAULT_PORT
				}

				startServer(port)
				return nil
			},
		},
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Command not found %q !", command)
	}
	app.Run(os.Args)
}
