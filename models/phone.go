package models

import (
	"github.com/jinzhu/gorm"
	"math/rand"
	"promotion/clients/smsclient"
	u "promotion/utils"
	"strconv"
	"strings"
	"time"
)

//a struct to rep phone
type Phone struct {
	gorm.Model
	Phone            string    `json:"phone"`
	ConfirmationCode int       `json:"confirmation_code"`
	Confirmed        bool      `json:"confirmed"`
	PromoCode        string    `json:"promo_code"`
	Redeemed         bool      `json:"redeemed"`
	RedeemDate       time.Time `json:"redeem_date"`
	RedeemProductId  uint      `json:"redeem_product_id"`
}

func (phone *Phone) Validate() (map[string]interface{}, bool) {

	if !(strings.HasPrefix(phone.Phone, "+") || strings.HasPrefix(phone.Phone, "00")) {
		return u.Message(false, "Phone number should be in the international format starting with either 00 or '+'"), false
	}

	if strings.HasPrefix(phone.Phone, "00") {
		phone.Phone = strings.Replace(phone.Phone, "00", "+", 1)
	}

	if !IsEuropean(phone.Phone[1:4]) {
		return u.Message(false, "Phone number is not European"), false
	}

	return u.Message(true, "Requirement passed"), true
}

func (phone *Phone) Create() map[string]interface{} {

	if resp, ok := phone.Validate(); !ok {
		return resp
	}

	err := GetDB().Table("phones").Where("phone = ?", phone.Phone).First(phone).Error
	if err == nil && phone.Confirmed {
		return u.Message(false, "Phone already registered")
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry")
	}

	phone.ConfirmationCode = rand.Intn(10000)
	smsBody := "Please go to CompanyXYZ.com/confirm in order to confirm your phone" +
		" number with the following confirmation code:" + strconv.Itoa(phone.ConfirmationCode)
	if resp, ok := smsclient.SendSms(phone.Phone, smsBody); !ok {
		return resp
	}

	if err == gorm.ErrRecordNotFound {
		GetDB().Create(phone)
	} else {
		GetDB().Save(phone)
	}

	if phone.ID <= 0 {
		return u.Message(false, "Failed to register phone, connection error.")
	}

	response := u.Message(true, smsBody)
	return response
}

func ConfirmPhone(phone *Phone) map[string]interface{} {

	if resp, ok := phone.Validate(); !ok {
		return resp
	}

	dbPhone := &Phone{}
	err := GetDB().Table("phones").Where("phone = ?", phone.Phone).First(dbPhone).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Account not found! You can register for a free account on CompanyXYZ.com/register. Sincerely CompanyXYZ!")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	if dbPhone.ConfirmationCode != phone.ConfirmationCode {
		return u.Message(false, "Not matching confirmation code, please try again")
	}

	// Success, phone is confirmed
	dbPhone.Confirmed = true
	GetDB().Save(dbPhone)

	response := u.Message(true, "Phone has been registered")
	response["phone"] = phone
	return response
}

func SendSms(phone *Phone) map[string]interface{} {

	if resp, ok := phone.Validate(); !ok {
		return resp
	}

	err := GetDB().Table("phones").Where("phone = ?", phone.Phone).First(phone).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Account not found! You can register for a free account on CompanyXYZ.com/register. Sincerely CompanyXYZ!")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	if !phone.Confirmed {
		return u.Message(false, "Phone is not yet confirmed")
	}

	afternoon := time.Now().Hour() > 12

	var greeting string
	if afternoon {
		phone.PromoCode = "PM456"
		greeting = "Hello! Your promocode is " + phone.PromoCode
	} else {
		phone.PromoCode = "AM123"
		greeting = "Good morning! Your promocode is " + phone.PromoCode
	}

	GetDB().Save(&phone)

	if resp, ok := smsclient.SendSms(phone.Phone, greeting); !ok {
		return resp
	}
	return u.Message(true, greeting)
}
