package main

import (
	"bufio"
	"errors"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/labstack/gommon/log"
	"golang.org/x/net/context"

	"google.golang.org/appengine"
)

const (
	nounLen      = 76216
	adjectiveLen = 26664
)

type handler struct {
	noun, adjective []string
}

func (h *handler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.Write([]byte(index))
	case "/wordlist":
		// adjective 1
		adj1 := rand.Intn(adjectiveLen)
		// adjective 2
		adj2 := rand.Intn(adjectiveLen)
		// noun
		noun3 := rand.Intn(adjectiveLen)
		w.Write([]byte(h.adjective[adj1] + " " + h.adjective[adj2] + " " + h.noun[noun3]))
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

func readWords(filename string, lines int) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("failed to open " + filename + ": " + err.Error())
	}
	defer f.Close()

	words := make([]string, lines)
	var c int
	s := bufio.NewScanner(f)
	for s.Scan() {
		words[c] = s.Text()
		c++
	}
	return words, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var err error
	h := &handler{}

	h.noun, err = readWords("noun.txt", nounLen)
	if err != nil {
		log.Errorf("failed to read noun: %s", err.Error())
		return
	}

	h.adjective, err = readWords("adjective.txt", adjectiveLen)
	if err != nil {
		log.Errorf("failed to read adjective: %s", err.Error())
		return
	}

	http.Handle("/", h)
	appengine.Main()
}

const index = `
<html>
<head>
	<title>Generate a sentence with 3 random words</title>
	<script>
	window.addEventListener('load', _ => {
		document.getElementById('submit').addEventListener('click', _ => {
			document.getElementById('string').textContent = "generating... (it will take about 5 seconds or little more)";
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
<div>Generate a sentence with 3 random words</div>
<div><button id="submit">generate</button></div>
<div id="string"></div>
</body>
</html>
`
