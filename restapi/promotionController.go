package restapi

import (
	"encoding/json"
	"net/http"
	"promotion/models"
	u "promotion/utils"
)

var RedeemPromotion = func(w http.ResponseWriter, r *http.Request) {

	promotion := &models.Promotion{}

	err := json.NewDecoder(r.Body).Decode(promotion)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := models.RedeemPromoCode(promotion)
	u.Respond(w, resp)
}

