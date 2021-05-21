package phragen

import (
	"errors"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"golang.org/x/net/context"
)

type apiv1Handler struct {
	nouns, adjectives []string
}

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

func Generate(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	rand.Seed(time.Now().UnixNano())
	h := &apiv1Handler{
		nouns:      nouns,
		adjectives: adjectives,
	}

	switch r.Method {
	case http.MethodGet:
		h.Get(ctx, w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("This method is not supported"))
	}
}
