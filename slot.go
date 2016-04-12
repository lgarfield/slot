package main

import(
	"fmt"
	"github.com/go-martini/martini"
	"html"
	// "log"
	"net/http"
	"slot/client"
	"slot/config"
	//"slot/web"
	"strings"
	// "time"
	"regexp"
)

func main() {
	m := martini.Classic()
	m.Use(MapEncoder)

	// start mysql
	dbmap := config.InitDb()
	defer dbmap.Db.Close()

	m.Get("/client", func(rw http.ResponseWriter, rq *http.Request) error {
		rq.ParseForm()
		fmt.Fprintln(rw, html.EscapeString(rq.URL.Host))
		fmt.Fprintf(rw, "Hello, %q, -=- %v", html.EscapeString(rq.URL.Path), html.EscapeString(rq.URL.Host))
		err := client.Get(rw, rq, dbmap)

		return err
	})

	m.Post("/client", func(rw http.ResponseWriter, rq *http.Request) error {
		err := client.Post(rw, rq, dbmap)

		return err
	})

	m.Run()
}

// The regex to check for the requested format (allows an optional tralling slash.)
var rxExt = regexp.MustCompile(`(\.(?:text|json))\/?$`)

// MapEncoder intercepts the request's URL, detects the required format, and injects the correct encoder dependency for this request.
// It writes the URL to remove the format extension, so that routes can be defined without it.
func MapEncoder(c martini.Context, w http.ResponseWriter, r *http.Request) {
	// Get the format extension
	matches := rxExt.FindStringSubmatch(r.URL.Path)
	ft := ".json"

	if len(matches) > 1 {
		l := len(r.URL.Path) - len(matches[1])
		if strings.HasSuffix(r.URL.Path, "/") {
			l--
		}
		r.URL.Path = r.URL.Path[:1]
		ft = matches[1]
	}

	switch ft {
	case ".text":
		text := textEncoder{}
		c.MapTo(text, (*Encoder)(nil))
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	default:
		json := jsonEncoder{}
		c.MapTo(json, (*Encoder)(nil))
		w.Header().Set("Content-Type", "application/json")
	}
}
