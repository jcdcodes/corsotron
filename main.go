package main

import (
	"flag"
	"fmt"
	"github.com/fitstar/falcore"
	"net/http"
)

var (
	port = flag.Int("port", 8000, "http listen port")
)

type CorsHeaderResponseFilter struct {
}

func (f *CorsHeaderResponseFilter) FilterResponse(req *falcore.Request, res *http.Response) {
	fmt.Println("setting a response header")
	res.Header.Set("Access-Control-Allow-Origin", "http://127.0.0.1:8000")
}

func main() {
	flag.Parse()

	pipeline := falcore.NewPipeline()

	pipeline.Upstream.PushBack(helloFilter)
	pipeline.Downstream.PushBack(corsFilter)

	server := falcore.NewServer(*port, pipeline)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("failed to start server:", err)
	}
}

var corsFilter = new(CorsHeaderResponseFilter)

var helloFilter = falcore.NewRequestFilter(func(req *falcore.Request) *http.Response {
	return falcore.StringResponse(req.HttpRequest, 200, nil, `
<html>
  <body>
    <p>For <em>some</em> kind of test, visit 127.0.0.1:8000 and note in the browser's dev console that localhost:8000 doesn't bload</p>
    <script>
      var r = new XMLHttpRequest();
      r.addEventListener("load", function() { console.log("monkey: "+this.responseText); });
      r.open("GET", "http://localhost:8000");
      r.setRequestHeader("Content-Type", "text/plain")
      r.send();
    </script>
  </body>
</html>
`)
})
