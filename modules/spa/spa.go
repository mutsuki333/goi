package spa

import (
	"embed"
	"net/http"
	"strings"

	"github.com/mutsuki333/goi/modules/log"
)

type SpaFS struct {
	FS embed.FS
	// Root the root path to prefix used in fs
	Root       string
	PublicPath string
}

func (fs SpaFS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.Replace(r.URL.Path, "//", "/", -1)
	path := fs.Root + strings.TrimPrefix(r.URL.Path, fs.PublicPath)
	data, err := fs.FS.ReadFile(path)
	if err == nil {
		contentType := ""
		switch {
		case strings.HasSuffix(path, "css"):
			contentType = "text/css; charset=utf-8"
		case strings.HasSuffix(path, "js"):
			contentType = "text/javascript; charset=utf-8"
		default:
			contentType = http.DetectContentType(data)
		}
		log.Debug(r.URL.Path, " : ", contentType, " => OK!")
		w.Header().Set("Content-Type", contentType)
		w.Write(data)
	} else {
		data, err = fs.FS.ReadFile(fs.Root + "/index.html")
		if err != nil {
			log.Error(r.URL.Path + " => " + err.Error())
			http.Error(w, err.Error(), 404)
			return
		}
		log.Debug(r.URL.Path, " => Not Found!")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(data)
	}
}
