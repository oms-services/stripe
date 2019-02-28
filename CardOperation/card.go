package CardOperation

import (
	"encoding/json"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/card"
	"net/http"
	"os"
)

func CreateCard(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")

	decoder := json.NewDecoder(r.Body)
	var param *stripe.CardParams
	err := decoder.Decode(&param)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	card, err := card.New(param)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}
	bytes, _ := json.Marshal(card)
	WriteJsonResponse(w, bytes, http.StatusCreated)
}

func GetResultCreated() int {
	return http.StatusCreated
}

func WriteErrorResponse(w http.ResponseWriter, err error) {
	msgbytes, _ := json.Marshal(err)
	WriteJsonResponse(w, msgbytes, http.StatusBadRequest)
}

func WriteJsonResponse(w http.ResponseWriter, bytes []byte, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
