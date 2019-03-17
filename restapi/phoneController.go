package restapi

import (
	"encoding/json"
	"net/http"
	"promotion/models"
	u "promotion/utils"
)

var RegisterPhone = func(w http.ResponseWriter, r *http.Request) {

	phone := &models.Phone{}
	err := json.NewDecoder(r.Body).Decode(phone) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := phone.Create() //Register phone
	u.Respond(w, resp)
}

var ConfirmPhone = func(w http.ResponseWriter, r *http.Request) {

	phone := &models.Phone{}
	err := json.NewDecoder(r.Body).Decode(phone) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.ConfirmPhone(phone) //Register phone
	u.Respond(w, resp)
}

var SendPromoCode = func(w http.ResponseWriter, r *http.Request) {

	phone := &models.Phone{}
	err := json.NewDecoder(r.Body).Decode(phone) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.SendSms(phone)
	u.Respond(w, resp)
}
