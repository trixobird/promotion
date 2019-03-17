package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

//a struct to rep phone
type TwoDigitCode struct {
	gorm.Model
	Code uint8
}

type ThreeDigitCode struct {
	gorm.Model
	Code uint16
}

func IsEuropean(phoneCode string) bool {

	// Russia has one digit
	if phoneCode[:1] == "7" {
		return true
	}

	c := make(chan bool)
	go searchTwoDigits(phoneCode, c)
	go searchThreeDigits(phoneCode, c)

	return <-c || <-c
}

func searchTwoDigits(code string, c chan bool) {

	phoneCodeInt, errConv := strconv.ParseUint(code[:2], 10, 8)
	if errConv != nil {
		c <- false
	}
	err := GetDB().Table("two_digit_codes").Where("code = ?", phoneCodeInt).First(&Phone{}).Error
	if err == nil {
		c <- true
	}
	c <- false
}

func searchThreeDigits(code string, c chan bool) {

	phoneCodeInt, errConv := strconv.ParseUint(code, 10, 16)
	if errConv != nil {
		c <- false
	}
	err := GetDB().Table("three_digit_codes").Where("code = ?", phoneCodeInt).First(&Phone{}).Error
	if err == nil {
		c <- true
	}
	c <- false
}

func SeedTwoDigits() {
	GetDB().Create(&TwoDigitCode{Code: 30})
	GetDB().Create(&TwoDigitCode{Code: 31})
	GetDB().Create(&TwoDigitCode{Code: 32})
	GetDB().Create(&TwoDigitCode{Code: 33})
	GetDB().Create(&TwoDigitCode{Code: 34})
	GetDB().Create(&TwoDigitCode{Code: 36})
	GetDB().Create(&TwoDigitCode{Code: 39})
	GetDB().Create(&TwoDigitCode{Code: 40})
	GetDB().Create(&TwoDigitCode{Code: 43})
	GetDB().Create(&TwoDigitCode{Code: 44})
	GetDB().Create(&TwoDigitCode{Code: 45})
	GetDB().Create(&TwoDigitCode{Code: 46})
	GetDB().Create(&TwoDigitCode{Code: 47})
	GetDB().Create(&TwoDigitCode{Code: 48})
	GetDB().Create(&TwoDigitCode{Code: 49})
	GetDB().Create(&TwoDigitCode{Code: 41})
	GetDB().Create(&TwoDigitCode{Code: 90})
}

func SeedThreeDigits() {
	GetDB().Create(&ThreeDigitCode{Code: 350})
	GetDB().Create(&ThreeDigitCode{Code: 351})
	GetDB().Create(&ThreeDigitCode{Code: 352})
	GetDB().Create(&ThreeDigitCode{Code: 353})
	GetDB().Create(&ThreeDigitCode{Code: 354})
	GetDB().Create(&ThreeDigitCode{Code: 356})
	GetDB().Create(&ThreeDigitCode{Code: 357})
	GetDB().Create(&ThreeDigitCode{Code: 358})
	GetDB().Create(&ThreeDigitCode{Code: 359})
	GetDB().Create(&ThreeDigitCode{Code: 370})
	GetDB().Create(&ThreeDigitCode{Code: 371})
	GetDB().Create(&ThreeDigitCode{Code: 372})
	GetDB().Create(&ThreeDigitCode{Code: 385})
	GetDB().Create(&ThreeDigitCode{Code: 420})
	GetDB().Create(&ThreeDigitCode{Code: 421})
	GetDB().Create(&ThreeDigitCode{Code: 423})
	GetDB().Create(&ThreeDigitCode{Code: 355})
	GetDB().Create(&ThreeDigitCode{Code: 376})
	GetDB().Create(&ThreeDigitCode{Code: 374})
	GetDB().Create(&ThreeDigitCode{Code: 994})
	GetDB().Create(&ThreeDigitCode{Code: 375})
	GetDB().Create(&ThreeDigitCode{Code: 387})
	GetDB().Create(&ThreeDigitCode{Code: 298})
	GetDB().Create(&ThreeDigitCode{Code: 995})
	GetDB().Create(&ThreeDigitCode{Code: 383})
	GetDB().Create(&ThreeDigitCode{Code: 389})
	GetDB().Create(&ThreeDigitCode{Code: 373})
	GetDB().Create(&ThreeDigitCode{Code: 377})
	GetDB().Create(&ThreeDigitCode{Code: 382})
	GetDB().Create(&ThreeDigitCode{Code: 374})
	GetDB().Create(&ThreeDigitCode{Code: 378})
	GetDB().Create(&ThreeDigitCode{Code: 381})
	GetDB().Create(&ThreeDigitCode{Code: 373})
	GetDB().Create(&ThreeDigitCode{Code: 380})
	GetDB().Create(&ThreeDigitCode{Code: 379})
}
