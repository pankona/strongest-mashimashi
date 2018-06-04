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
	// inspect word files in advance by "wc -l" command
	nounLen      = 76216
	adjectiveLen = 26664
)

type handler struct {
	noun, adjective []string
}

const (
	apiv1Prefix = "/api/v1"
)

func (h *handler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case apiv1Prefix + "/phrase":
		// adjective 1
		adj1 := rand.Intn(adjectiveLen)
		// adjective 2
		adj2 := rand.Intn(adjectiveLen)
		// noun
		noun3 := rand.Intn(adjectiveLen)
		w.Write([]byte(h.adjective[adj1] + " " + h.adjective[adj2] + " " + h.noun[noun3]))
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	switch r.Method {
	case http.MethodGet:
		h.Get(ctx, w, r)
	default:
		w.Write([]byte("unimplemented method!"))
	}
}

func loadWords(filename string, lines int) ([]string, error) {
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

	h.noun, err = loadWords("noun.txt", nounLen)
	if err != nil {
		log.Errorf("failed to read noun: %s", err.Error())
		return
	}

	h.adjective, err = loadWords("adjective.txt", adjectiveLen)
	if err != nil {
		log.Errorf("failed to read adjective: %s", err.Error())
		return
	}

	http.Handle("/", http.FileServer(http.Dir("./webapp/public")))
	http.Handle(apiv1Prefix, h)
	appengine.Main()
}
