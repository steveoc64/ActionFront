package main

import (
	"flag"
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"net/http"
)

var (
	httpListen = flag.String("http", "localhost:8080", "host:port to listen on")
)

func main() {

	flag.Parse()

	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "unitEditor", "as such")
	})

	//m.Run()
	fmt.Println("ActionFront Unit Editor starting on ", *httpListen)
	http.ListenAndServe(*httpListen, m)
}
