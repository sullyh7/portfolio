package main

import (
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/sullyh7/portfolio/view"
	"github.com/sullyh7/portfolio/view/notfound"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() *echo.Echo {
	r := echo.New()
	// r.Use(middleware.RequestID)
	// r.Use(middleware.RealIP)
	// r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)

	// r.Use(middleware.Timeout())
	assetHandler := http.FileServer(view.GetPublicAssetsFileSystem())
	r.RouteNotFound("/*", func(c echo.Context) error {
		return app.render(c, 404, notfound.Index())
	})
	r.GET("/public/*", echo.WrapHandler(http.StripPrefix("/public/", assetHandler)))
	r.GET("/health", func(c echo.Context) error {
		return c.String(200, "ok")
	})
	r.GET("/", app.homeHandler)

	return r
}

func (a *application) run(mux *echo.Echo) error {
	srv := &http.Server{
		Addr:         a.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("server has started at %s", a.config.addr)
	return srv.ListenAndServe()
}

func (a *application) render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
