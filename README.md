# corsotron

A super simple falcore app that sets CORS headers so we can have a web
page at, say, http://127.0.0.1:8000, fetch a resource from
http://localhost:8000 without causing the browser to complain about
potential XSS attacks.

Also just goofing around with golang and falcore.

## Usage

    go build . && ./corsotron

and then point a browser to http://127.0.0.1:8000 and view the dev
tools console.  Instead of XSS-related complaints you should see the
HTML source of the page you're looking at.

Profit!
