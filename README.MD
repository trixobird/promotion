# Coding Challenge
Company XYZ, a well-established ecommerce retailer has decided to give a discount coupon that can be redeemed for free discounts on products listed on its website CompanyXYZ.com. However, to stop abuse we want to provide the discount code only in exchange for a telephone number, sending the code in an SMS message. Furthermore, the coupon issuing needs to be restricted to customers living in Europe which is the region the company ships its products to, and the promo campaign is associated with the company’s terms and conditions which are mandatory.
Your task is to create API endpoints or microservices that handle the above business scenario. The mandatory endpoints for this coding challenge are listed below.
1. Endpoint ‘sms-promotion’. The request body will consist of a single field, “phone”. You should then send an SMS to the provided phone number, with one of two messages:
If the user’s local time is in the morning, the message should read “Good morning! Your promocode is AM123”.
If the user’s local time is in the afternoon or later, the message should read “Hello! Your promocode is PM456”.
Only one of the two can be redeemed per phone number.
2. Endpoint ‘redeem-promotion’. The request body will consist of three fields, “phone”, “promocode” and “productId”. You should then send an SMS to the provided phone number, with one of three messages:
If the promo code has never been redeemed, the message should read “Congratulations! You have redeemed promo code: {promocode} on product: {productid} at {redeemDate}. Sincerely CompanyXYZ!”.
If the promo has already been redeemed, the message should read “You have already redeemed promo code: {promocode} on product: {productid} at {redeemDate}. Sincerely CompanyXYZ!”.
If the phone provided is not registered, the message should read “Account not found! You can register for a free account on CompanyXYZ.com/register. Sincerely CompanyXYZ!”.
You can either create a register endpoint to register new phone numbers or work with a pre-registered set of phone numbers. The choice is up to you.

The endpoints should return suitable HTTP responses (you can get creative here), including handling errors appropriately.
This task is preferably coded in Golang but candidates can use any language that they feel most comfortable with. Our focus will not be on language syntax but more on design, coding style and logic.

Trial third party services can be used for the coding challenge for a trial period without cost
## Installation
### Prerequisites
* Postgresql 
* Golan
* An IDE (I used GoLand)
### Database set up
* Create a user with user/pass 'promotion' (or change the relevant property in .env)
* Create a database named promotion (or change the relevant property in .env)
## Requests
Only the happy path is presented here. Order matters!
### Register Phone
```
curl -X POST \
  http://localhost:6060/api/register \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 1db1e678-d4ae-43db-8559-85be84d21191' \
  -H 'cache-control: no-cache' \
  -d '{"phone":"+357XXXXXX"}'
  ```
### Confirm Phone
```
curl -X POST \
  http://localhost:6060/api/confirm \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: dbda94ac-e5f5-45e8-91c3-ea2f210eeea8' \
  -H 'cache-control: no-cache' \
  -d '{"phone":"+357XXXXXX","confirmation_code":1234}'
  ```
### Send Code
```
curl -X POST \
  http://localhost:6060/api/sms-promotion \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: eeabb802-fe89-42b1-bff7-1568e4f3cd8b' \
  -H 'cache-control: no-cache' \
  -d '{"phone":"+357XXXXXX"}'
  ```
### Redeem Promotion
```
curl -X POST \
  http://localhost:6060/api/redeem-promotion \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 34d8b07f-5bfb-4210-9a17-30d0c3c9413a' \
  -H 'cache-control: no-cache' \
  -d '{
    "phone": "+357XXXXXX",
    "promo_code": "PM456",
    "product_id": 1
}'
```
## Business Error Codes
### Register Phone
* Phone number should be in the international format starting with either 00 or '+'
* Phone number is not European
* Phone already registered
### Confirm Phone
* Phone number should be in the international format starting with either 00 or '+'
* Phone number is not European
* Account not found! You can register for a free account on CompanyXYZ.com/register. Sincerely CompanyXYZ!
* Not matching confirmation code, please try again
* Phone has been registered
### Send Code
* Phone number should be in the international format starting with either 00 or '+'
* Phone number is not European
* Account not found! You can register for a free account on CompanyXYZ.com/register. Sincerely CompanyXYZ!
* Phone is not yet confirmed
### Redeem Promotion
* Phone number should be on the payload
* Phone number should be in the international format starting with either 00 or '+'
* Phone number is not European
* Promotion code should be on the payload
* Product Id is not recognized
* Account not found! You can register for a free account on CompanyXYZ.com/register. Sincerely CompanyXYZ!
* Phone is not yet confirmed
* You have already redeemed promo code: {promocode} on product: {productid} at {redeemDate}. Sincerely CompanyXYZ!
* The phone does not have a promocode registered. Get yours on CompanyXYZ.com/sms-promotion
* The promocode is not correct. Please try again 
