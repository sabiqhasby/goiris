package main

import (
	"goiris/configs"
	"goiris/handler"
	"goiris/migrations"

	"github.com/kataras/iris/v12"
)

type PingResponse struct {
	Message string `json:"message"`
}

func main() {
	db, err := configs.InitDB()
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	migrations.Migrate(db)
	app := iris.New()
	app.Use(myMiddleware)

	booksApi := app.Party("/books")
	{
		booksApi.Use(iris.Compression)
		booksApi.Get("/", handler.ShowBooks(db))

		booksApi.Get("/list", func(ctx iris.Context) {
			ctx.StatusCode(iris.StatusOK)
			ctx.JSON(iris.Map{"status": 200, "message": "hellllo"})
		})
	}

	/* Same as:
	   app.Handle("GET", "/ping", func(ctx iris.Context) {
	       ctx.JSON(iris.Map{
	           "message": "pong",
	       })
	   })
	*/

	// Listens and serves incoming http requests
	// on http://localhost:8080.
	app.Listen(":8080")
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
