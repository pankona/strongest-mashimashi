package phragen

import (
	"bufio"
	"errors"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
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

type apiv1Handler struct {
	nouns, adjectives []string
}

const (
	apiv1Prefix = "/api/v1/"
)

func getNumFromQuery(values url.Values) (int, error) {
	if len(values["num"]) == 0 {
		return 0, errors.New("num is not specified in query")
	}

	num, err := strconv.Atoi(values["num"][0])
	if err != nil {
		return 0, errors.New("invalid num [" + values["num"][0] + "] is specified. must be a integer")
	}

	if num <= 0 || num >= 6 {
		return 0, errors.New("invalid num [" + values["num"][0] + "] is specified. must be in a range from 1 to 5")
	}

	return num, nil
}

func (h *apiv1Handler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case apiv1Prefix + "phrase":
		n, err := getNumFromQuery(r.URL.Query())
		if err != nil {
			n = 3 // default value
		}

		var phrase string
		// -1 loop count since noun must be reserved
		for i := 0; i < n-1; i++ {
			phrase += h.adjectives[rand.Intn(len(h.adjectives))] + " "
		}
		phrase += h.nouns[rand.Intn(len(h.nouns))]
		w.Write([]byte(phrase))
	}
}

func (h *apiv1Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func Generate(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		ctx = context.Background()
	)
	rand.Seed(time.Now().UnixNano())
	h := &apiv1Handler{}

	h.nouns, err = loadWords("noun.txt", nounLen)
	if err != nil {
		log.Errorf("failed to read noun: %v", err)
		return
	}

	h.adjectives, err = loadWords("adjective.txt", adjectiveLen)
	if err != nil {
		log.Errorf("failed to read adjective: %v", err)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.Get(ctx, w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("This method is not supported"))
	}
}
