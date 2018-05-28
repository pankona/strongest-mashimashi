package main

import (
	"context"
	"net/http"

	"google.golang.org/appengine"
)

type handler struct {
}

func (h *handler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.Write([]byte(index))
	case "/wordlist":
		w.Write([]byte("wordlist is underconstruction..."))
	}
}

func (h *handler) Post(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("post!"))
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	switch r.Method {
	case http.MethodGet:
		h.Get(ctx, w, r)
	case http.MethodPost:
		h.Post(ctx, w, r)
	default:
		w.Write([]byte("unimplemented method!"))
	}
}

func main() {
	h := &handler{}
	http.Handle("/", h)
	appengine.Main()
}

const index = `
<html>
<head>
	<title>hello mashimashi</title>
	<script>
	window.addEventListener('load', _ => {
		document.getElementById('submit').addEventListener('click', _ => {
			fetch('/wordlist?num=3', {
				method: 'GET'
			}).then(response => {
				if (response.ok) {
					return response.text();
				} else {
					throw new Error();
				}
			}).then(text => {
				document.getElementById('string').textContent = text;
			}).catch(error => {
				console.log(error);
			});
		});
	});
	</script>
</head>
<body>
<div>hello mashimashi!</div>
<div><button id="submit">click me</button></div>
<div id="string"></div>
</body>
</html>
`
