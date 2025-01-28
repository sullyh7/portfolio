package view

import (
	"context"
	"embed"
	"io"
	"io/fs"
	"net/http"

	"github.com/a-h/templ"
)

//go:embed public/styles.css
var publicAssets embed.FS

// GetPublicAssetsFileSystem returns a http.FileSystem for the public assets so that
// we can embed them into the binary so it is self-contained.
func GetPublicAssetsFileSystem() http.FileSystem {
	fsys, err := fs.Sub(publicAssets, "public")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}

func Unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}
