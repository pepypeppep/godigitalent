package main

import (
	"database/sql"
	"embed"
	"fmt"
	"godigitalent"
	"godigitalent/mysqldata"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/joho/godotenv"
)

// Embed a directory
//go:embed template/*
var embedDirStatic embed.FS

func getAllFilenames(fs *embed.FS, path string) (out []string, err error) {
	if len(path) == 0 {
		path = "."
	}
	entries, err := fs.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		fp := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			res, err := getAllFilenames(fs, fp)
			if err != nil {
				return nil, err
			}
			out = append(out, res...)
			continue
		}
		out = append(out, fp)
	}
	return
}

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	return db, err
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		ReadBufferSize: 5000,
	})

	db, err := GetDatabase()
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to PlanetScale!")

	server := godigitalent.Server{
		App:     app,
		DB:      db,
		Queries: mysqldata.New(db),
	}

	server.Routes()

	// Or extend your config for customization
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "*",
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// }))

	// Access file "image.png" under `static/` directory via URL: `http://<server>/static/image.png`.
	// Without `PathPrefix`, you have to access it via URL:
	// `http://<server>/static/static/image.png`.
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(embedDirStatic),
		PathPrefix: "template",
		Browse:     true,
	}))

	// app.Use("/", func(c *fiber.Ctx) error {
	// 	list, err := getAllFilenames(&embedDirStatic, "template")
	// 	if err != nil {
	// 		return c.JSON(fiber.Map{
	// 			"message": err.Error(),
	// 		})
	// 	}

	// 	return c.JSON(list)
	// })

	log.Fatal(app.Listen(":3000"))
}
