package main

import (
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/net/context"

	client "github.com/pankona/oxford-dict-api-client-go/client"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

type handler struct {
	client *client.APIClient
}

func randomChar() string {
	n := rand.Intn(26)
	return string('a' + n)
}

func (h *handler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.Write([]byte(index))
	case "/wordlist":
		wl, _, err := h.client.WordlistApi.WordlistSourceLangFiltersAdvancedGet(
			context.Background(),
			"en",
			"lexicalCategory=adjective",
			appID,
			appKey,
			map[string]interface{}{
				"wordLength": ">5,<10",
				"prefix":     randomChar(),
			},
		)
		if err != nil {
			log.Debugf(ctx, "%s", err.Error())
			return
		}
		word1 := wl.Results[rand.Intn(len(wl.Results))].Word

		wl, _, err = h.client.WordlistApi.WordlistSourceLangFiltersAdvancedGet(
			context.Background(),
			"en",
			"lexicalCategory=adjective",
			appID,
			appKey,
			map[string]interface{}{
				"wordLength": ">5,<10",
				"prefix":     randomChar(),
			},
		)
		if err != nil {
			log.Debugf(ctx, "%s", err.Error())
			return
		}
		word2 := wl.Results[rand.Intn(len(wl.Results))].Word

		wl, _, err = h.client.WordlistApi.WordlistSourceLangFiltersAdvancedGet(
			context.Background(),
			"en",
			"lexicalCategory=noun",
			appID,
			appKey,
			map[string]interface{}{
				"wordLength": ">5,<10",
				"prefix":     randomChar(),
			},
		)
		if err != nil {
			log.Debugf(ctx, "%s", err.Error())
			return
		}
		word3 := wl.Results[rand.Intn(len(wl.Results))].Word
		w.Write([]byte(word1 + " " + word2 + " " + word3))
	}
}

func (h *handler) Post(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("post!"))
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	cfg := client.NewConfiguration()
	cfg.HTTPClient = urlfetch.Client(ctx)
	h.client = client.NewAPIClient(cfg)
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
	rand.Seed(time.Now().Unix())
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
<div><button id="submit">generate</button></div>
<div id="string"></div>
</body>
</html>
`
