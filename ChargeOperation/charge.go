package ChargeOperation

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/heaptracetechnology/microservice-stripe/result"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

//CreateCharge
func CreateCharge(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")

	decoder := json.NewDecoder(r.Body)
	var param *stripe.ChargeParams
	err := decoder.Decode(&param)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	errr := param.SetSource("tok_visa")
	if errr != nil {
		result.WriteErrorResponse(nil, errr)
	}

	ch, err := charge.New(param)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	bytes, _ := json.Marshal(ch)
	result.WriteJsonResponse(w, bytes, http.StatusCreated)
}

//Capture Charge
func CaptureCharge(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")

	vars := mux.Vars(r)
	var charge_id = vars["charge"]

	decoder := json.NewDecoder(r.Body)
	var param *stripe.CaptureParams
	err := decoder.Decode(&param)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	ch, err := charge.Capture(charge_id, param)
	if err != nil {
		result.WriteErrorResponse(nil, err)
	}
	bytes, _ := json.Marshal(ch)
	result.WriteJsonResponse(w, bytes, http.StatusOK)
}
