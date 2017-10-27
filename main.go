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

func main() {
     flag.Parse()

     pipeline := falcore.NewPipeline()

     pipeline.Upstream.PushBack(helloFilter)

     server := falcore.NewServer(*port, pipeline)

     if err := server.ListenAndServe(); err != nil {
     	fmt.Println("failed to start server:", err)
     }
     
}

var helloFilter = falcore.NewRequestFilter(func(req *falcore.Request) *http.Response {
    return falcore.StringResponse(req.HttpRequest, 200, nil, `
<html>
  <body>
    <p>For <em>some</em> kind of test, visit 127.0.0.1:8000 and note in the browser's dev console that localhost:8000 doesn't bload</p>
    <script>
      var r = new XMLHttpRequest();
      r.addEventListener("load", function() { console.log(this.responseText); });
      r.open("GET", "localhost:8000");
      r.send();
    </script>
  </body>
</html>
`)
})