package rollup

import (
	"net/http"
	"path/filepath"
	"io/fs"
	"log"

	aa_log "github.com/aaronland/go-log/v2"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
)

type RollupCSSHandlerOptions struct {
	FS fs.FS
	Paths map[string][]string
	Logger *log.Logger
}

func RollupCSSHandler(opts *RollupCSSHandlerOptions) (http.Handler, error) {

	m := minify.New()
	m.AddFunc("text/css", css.Minify)

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		path := req.URL.Path
		fname := filepath.Base(path)

		js_paths, exists := opts.Paths[fname]

		if !exists {
			http.Error(rsp, "Not found", http.StatusNotFound)
			return
		}

		rsp.Header().Set("Content-type", "text/css")

		for _, path := range js_paths {

			r, err := opts.FS.Open(path)

			if err != nil {
				aa_log.Error(opts.Logger, "Failed to open %s for reading, %v", path, err)
				http.Error(rsp, err.Error(), http.StatusInternalServerError)
				return
			}

			defer r.Close()

			err = m.Minify("text/css", rsp, r)

			if err != nil {
				aa_log.Error(opts.Logger, "Failed to minify %s, %v", path, err)				
				http.Error(rsp, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		return
	}

	return http.HandlerFunc(fn), nil
}
