package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "promotion/utils"
	"time"
)

type Promotion struct {
	gorm.Model
	Phone string `json:"phone"`
	PromoCode string `json:"promo_code"`
	ProductId uint `json:"product_id"`
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (promotion *Promotion) Validate() (map[string]interface{}, bool) {

	if promotion.Phone == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}

	phone := &Phone{Phone: promotion.Phone}
	if resp, ok := phone.Validate(); !ok {
		return resp, false
	}

	if promotion.PromoCode == "" {
		return u.Message(false, "Promotion code should be on the payload"), false
	}

	if promotion.ProductId <= 0 {
		return u.Message(false, "Product Id is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func RedeemPromoCode(promotion *Promotion) map[string]interface{} {

	if resp, ok := promotion.Validate(); !ok {
		return resp
	}

	phone := &Phone{}
	err := GetDB().Table("phones").Where("phone = ?", promotion.Phone).First(phone).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Account not found! You can register for a free account on CompanyXYZ.com/register. Sincerely CompanyXYZ!")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	if phone.Redeemed {
		return u.Message(false, "You have already redeemed promo code:"+phone.PromoCode+
			" on product:"+fmt.Sprint(phone.RedeemProductId)+" at "+phone.RedeemDate.Format("Mon, 02 Jan 2006 15:04:05")+
			". Sincerely CompanyXYZ!")
	}

	if phone.PromoCode == "" {
		return u.Message(false, "The phone does have a promocode registered. Get yours on CompanyXYZ.com/sms-promotion")
	}

	if phone.PromoCode != promotion.PromoCode {
		return u.Message(false, "The promocode is not correct. Please try again")
	}

	// Update db
	phone.RedeemDate = time.Now()
	phone.RedeemProductId = promotion.ProductId
	phone.Redeemed = true
	GetDB().Save(&phone)

	return u.Message(true, "Congratulations! You have redeemed promo code:"+phone.PromoCode+
		" on product:"+fmt.Sprint(phone.RedeemProductId)+" at "+phone.RedeemDate.Format("Mon, 02 Jan 2006 15:04:05")+
		". Sincerely CompanyXYZ!")
}
