package ChargeOperation

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"net/http"
	"os"
)

//CreateCharge
func CreateCharge(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")

	decoder := json.NewDecoder(r.Body)
	var param *stripe.ChargeParams
	err := decoder.Decode(&param)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	param.SetSource("tok_visa")

	ch, err := charge.New(param)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}
	bytes, _ := json.Marshal(ch)
	WriteJsonResponse(w, bytes, http.StatusCreated)
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
		WriteErrorResponse(w, err)
		return
	}

	ch, err := charge.Capture(charge_id, param)
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}
	bytes, _ := json.Marshal(ch)
	WriteJsonResponse(w, bytes, http.StatusOK)
}

func GetResultCreated() int {
	return http.StatusCreated
}

func GetResultOK() int {
	return http.StatusOK
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
