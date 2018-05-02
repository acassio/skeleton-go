package app

import (
	"github.com/acassio/base-app-go/conf"
	"github.com/acassio/base-app-go/handler"
	"github.com/acassio/base-app-go/lib/cache"
	"github.com/acassio/base-app-go/lib/contx"
	"github.com/acassio/base-app-go/lib/cors"
	"github.com/acassio/base-app-go/model"
	"github.com/go-macaron/binding"
	mcache "github.com/go-macaron/cache"
	"github.com/go-macaron/gzip"
	"github.com/go-macaron/i18n"
	"github.com/go-macaron/session"
	"github.com/go-macaron/toolbox"
	"gopkg.in/macaron.v1"
)

//SetupMiddlewares configures the middlewares using in each web request
func SetupMiddlewares(app *macaron.Macaron) {
	app.Use(macaron.Logger())
	app.Use(macaron.Recovery())
	app.Use(gzip.Gziper())
	app.Use(toolbox.Toolboxer(app, toolbox.Options{
		HealthCheckers: []toolbox.HealthChecker{
			new(handler.AppChecker),
		},
	}))

	app.Use(macaron.Static("public"))
	app.Use(i18n.I18n(i18n.Options{
		Directory: "locale",
		Langs:     []string{"pt-BR", "en-US"},
		Names:     []string{"Português do Brasil", "American English"},
	}))

	//Cache in memory
	app.Use(mcache.Cacher(
		cache.Option(conf.Cfg.Section("").Key("cache_adapter").Value()),
	))

	app.Use(session.Sessioner())
	app.Use(contx.Contexter())
	app.Use(cors.Cors())
}

//SetupRoutes defines the routes the Web Application will respond
func SetupRoutes(app *macaron.Macaron) {
	app.Get("", func() string {
		return "App Base running!"
	})

	app.Get("/user", handler.GetUser)
	app.Post("/user", binding.Bind(model.User{}), handler.CreateUser)

}
