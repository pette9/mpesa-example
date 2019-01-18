package service

import (
	"encoding/json"
	"io/ioutil"
	"mpesa-example/utils"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pette9/mpesa"

	"github.com/sirupsen/logrus"
)

var mpesaClient *mpesa.Client

//MakeRequest ...
func MakeRequest(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)["request"]
	var params interface{}
	switch requestName := request; requestName {
	default:
		utils.WriteErrorResponse(w, http.StatusBadRequest, "invalid request name")
		return
	case "b2b":
		b2bRequestParams := mpesa.B2BQueryParameters{}
		json.NewDecoder(r.Body).Decode(&b2bRequestParams)
		params = b2bRequestParams
	case "b2c":
		b2cRequestParams := mpesa.B2CQueryParameters{}
		json.NewDecoder(r.Body).Decode(&b2cRequestParams)
		params = b2cRequestParams
	case "c2b-register-url":
		c2bRegisterURLParams := mpesa.C2BRegisterURLParameters{}
		json.NewDecoder(r.Body).Decode(&c2bRegisterURLParams)
		params = c2bRegisterURLParams
	case "c2b-simulate":
		c2bSimulateParams := mpesa.C2BSimulateParameters{}
		json.NewDecoder(r.Body).Decode(&c2bSimulateParams)
		params = c2bSimulateParams
	case "lipa-na-mpesa-online-payment":
		lnmopParams := mpesa.LipaNaMpesaOnlinePaymentParameters{}
		json.NewDecoder(r.Body).Decode(&lnmopParams)
		params = lnmopParams
	case "lipa-na-mpesa-online-query":
		lnmoqParams := mpesa.LipaNaMpesaOnlineQueryParameters{}
		json.NewDecoder(r.Body).Decode(&lnmoqParams)
		params = lnmoqParams
	case "reversal":
		reversalParams := mpesa.ReversalParameters{}
		json.NewDecoder(r.Body).Decode(&reversalParams)
		params = reversalParams
	case "transaction-status":
		transactionStatusParams := mpesa.TransactionStatusParameters{}
		json.NewDecoder(r.Body).Decode(&transactionStatusParams)
		params = transactionStatusParams
	}
	mpesaToken, err := mpesaClient.Authenticate()
	if err != nil {
		logrus.Errorln(err.Error())
		w.WriteHeader(http.StatusForbidden)
		return
	}
	response, err := mpesaClient.MakeRequest(mpesaToken["access_token"], request, params)
	if err != nil {
		logrus.Errorln(err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	data, _ := json.Marshal(response)
	utils.WriteJSONResponse(w, http.StatusCreated, data)
}

//Results ...
func Results(w http.ResponseWriter, r *http.Request) {
	response, _ := ioutil.ReadAll(r.Body)
	logrus.Infoln(string(response))
}
