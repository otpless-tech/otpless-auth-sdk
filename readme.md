# Merchant Integration Documentation(Backend GoLang Auth SDK)

---

> ## A. OTPLessAuth Dependency
>
> install Below dependency in your project's

```go
go get github.com/otpless-tech/otpless-auth-sdk@latest
```

you can also get latest version of dependency
at https://pkg.go.dev/github.com/otpless-tech/otpless-auth-sdk/pkg

---

---

> ## B. OTPLessAuth

The `OTPLessAuth` pkg provides methods to integrate OTPLess authentication into your GoLang backend application. This
documentation explains the usage of the class and its methods.

### Methods:

---

> ### 1. decodeIdToken

---

This method help to resolve `idToken(JWT token)` which is issued by `OTPLess` which return user detail
from that token also this method verify that token is valid, token should not expired and
issued by only `otpless.com`

##### Method Signature:

```go
DecodeIDToken(idToken, clientID, clientSecret, audience string)
```

#### Method Params:

| Params       | Data type | Mandatory | Constraints | Remarks                                                                      |
| ------------ | --------- | --------- | ----------- | ---------------------------------------------------------------------------- |
| idToken      | String    | true      |             | idToken which is JWT token which you get from `OTPLess` by exchange code API |
| clientId     | String    | true      |             | Your OTPLess `Client Id`                                                     |
| clientSecret | String    | true      |             | Your OTPLess `Client Secret`                                                 |
| audience     | String    | false     |             | None                                                                         |

#### Return

Return:
Object Name: UserDetailResult

> ### 2. VerifyCode

---

This method help to resolve `token` which is issued by `OTPLess` which return user detail
from that token also this method verify that token is valid, token should not expired and
issued by only `otpless.com`

##### Method Signature:

```go
VerifyCode(code, clientID, clientSecret)
```

#### Method Params:

| Params       | Data type | Mandatory | Constraints | Remarks                           |
| ------------ | --------- | --------- | ----------- | --------------------------------- |
| code         | String    | true      |             | code which you get from `OTPLess` |
| clientId     | String    | true      |             | Your OTPLess `Client Id`          |
| clientSecret | String    | true      |             | Your OTPLess `Client Secret`      |

#### Return

Return:
Object Name: UserDetailResult

> ### 3. VerifyAuthToken

---

This method help to resolve `token` which is issued by `OTPLess` which return user detail
from that token also this method verify that token is valid, token should not expired and
issued by only `otpless.com`

##### Method Signature:

```go
VerifyAuthToken(token, clientID, clientSecret string)
```

#### Method Params:

| Params       | Data type | Mandatory | Constraints | Remarks                            |
| ------------ | --------- | --------- | ----------- | ---------------------------------- |
| token        | String    | true      |             | token which you get from `OTPLess` |
| clientId     | String    | true      |             | Your OTPLess `Client Id`           |
| clientSecret | String    | true      |             | Your OTPLess `Client Secret`       |

#### Return

Return:
Object Name: UserDetailResult

```json
userDetail := UserDetailResult{
    Success:     true,
    AuthTime:  1697473071,
    Email:       "dev****@gmail.com",
    Name:        "Dhaval From OTP-less",
    PhoneNumber: "+9193*****",
    country_code: "+91",
    national_phone_number: "95******",
}
```

> ### UserDetail Object Fields:
>
> `success` (boolean): This will be `true` in case of method successfully performed operation.<br> > `authTime` (Long, required): The time when authentication was completed.<br> > `phoneNumber` (String, required): The user's phone number.<br> > `countryCode` (String, required): The country code of user's phone number.<br> > `nationalPhoneNumber` (String, required): The user's phone number without country code.<br> > `email` (String, required): The user's email address.<br> > `name` (String, required): The user's full name.<br>

> ### 4. Generate Magic link

---

The Authorization Endpoint initiates the authentication process by sending a `magic link` to the user's WhatsApp or email, based on the provided contact information. This link is used to verify the identity of the user. Upon the user's action on this link, they are redirected to the specified URI with an authorization code included in the redirection.

##### Method Signature:

```Go
GenerateMagicLink(mobileNumber, email, clientID, clientSecret, redirectUri, channel string) (*MagicLinkResponse, error)
```

#### Method Params:

| Params        | Data type | Mandatory | Constraints                                       | Remarks                                                                                               |
| ------------- | --------- | --------- | ------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| channel       | String    | false     | if no channel given WHATSAPP is chosen as default | WHATSAPP/SMS                                                                                          |
| mobile_number | String    | false     | At least one required                             | The user's mobile number for authentication in the format: country code + number (e.g., 91XXXXXXXXXX) |
| email         | String    | false     | At least one required                             | The user's email address for authentication.                                                          |
| redirect_uri  | String    | true      |                                                   | The URL to which the user will be redirected after authentication. This should be URL-encoded         |
| clientId      | String    | true      |                                                   | Your OTPLess `Client Id`                                                                              |
| clientSecret  | String    | true      |                                                   | Your OTPLess `Client Secret`                                                                          |

#### Return

Return:
Object Name: RquestIds

```json
&{RequestIds:[{Type:MOBILE Value:c36b678aef104691b15f93910acfee48 DestinationUri:https://wa.me/91971***XX**} {Type:EMAIL Value:39250d56e8da4f4cb86224929bf76a2d}]}
```

### Error case:

`success` (boolean): This will be `false`. The method is failed to perform.<br>
`errorMessage` (String): The message contains error information.<br>

> ## C. OTP class

These methods enable you to send, resend and verify OTP.

### Methods:

---

> ### 1. Send OTP

---

##### Method Signature:

```go
SendOTP(req SendOTPRequest, clientID, clientSecret string) (*SendOTPResponse, error)

```

#### Method Params:

| Params       | Data type | Mandatory | Constraints                                                     | Remarks                                                           |
| ------------ | --------- | --------- | --------------------------------------------------------------- | ----------------------------------------------------------------- |
| phoneNumber  | String    | true      |                                                                 | Mobile Number of your users                                       |
| email        | String    | true      |                                                                 | Mail Id of your users                                             |
| orderId      | String    | false     |                                                                 | An Merchant unique id for the request.                            |
| expiry       | Int       | false     |                                                                 | OTP expiry in sec                                                 |
| hash         | String    | false     |                                                                 | An Hash will be used to auto read OTP.                            |
| clientId     | String    | true      |                                                                 | Your OTPLess `Client Id`                                          |
| clientSecret | String    | true      |                                                                 | Your OTPLess `Client Secret`                                      |
| otpLength    | Integer   | false     | 4 or 6 only allowed                                             | Allowes you to send OTP in 4/6 digit. default will be 6 digit.    |
| channel      | String    | false     | SMS/WHATSAPP/ALL (if no channel given SMS is chosen as default) | Allowes you to send OTP on WhatsApp/SMS/Both. default will be SMS |

#### Return

Return:
Object Name: SendOTPResponse

```json
{
  "orderId": "Otp_ED5F709E1C6B41EB8C0595C7968354EB"
}
```

### 2. Resend OTP

##### Method Signature:

```go
    ResendOTP(orderID, clientID, clientSecret string) (*ResendOTPResponse, error)
```

#### Method Params:

| Params       | Data type | Mandatory | Constraints | Remarks                                |
| ------------ | --------- | --------- | ----------- | -------------------------------------- |
| orderId      | String    | true      |             | An Merchant unique id for the request. |
| clientId     | String    | true      |             | Your OTPLess `Client Id`               |
| clientSecret | String    | true      |             | Your OTPLess `Client Secret`           |

#### Return

Return:
Object Name: ResendOTPResponse

```json
{
  "orderId": "DP0000111"
}
```

---

### 3. Verify OTP

##### Method Signature:

```go
    VerifyOTP(orderID, otp, email, phoneNumber, clientID, clientSecret string) (*VerifyOTPResponse, error)
```

#### Method Params:

| Params       | Data type | Mandatory | Constraints | Remarks                                     |
| ------------ | --------- | --------- | ----------- | ------------------------------------------- |
| email        | String    | true      |             | An email on which OTP has been sent.        |
| phoneNumber  | String    | true      |             | An phone number on which OTP has been sent. |
| orderId      | String    | true      |             | An Merchant unique id for the request.      |
| otp          | String    | true      |             | OTP value.                                  |
| clientId     | String    | true      |             | Your OTPLess `Client Id`                    |
| clientSecret | String    | true      |             | Your OTPLess `Client Secret`                |

#### Return

Return:
Object Name: VerifyOTPResponse

- `reason` (String): The will be errorMessage in case of OTP doesn't verified

```json
{
  "isOTPVerified": true
}
```

### Example of usage

```go
package main

import (
	"fmt"

	otplessAuthSdk "github.com/otpless-tech/otpless-auth-sdk"
)

func main() {
	clientID := "your_client_id"
	clientSecret := "your_client_secret"
	code := "some_code"

	result, err := otplessAuthSdk.VerifyCode(clientID, clientSecret, code)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}

```

This method allows you to decode and verify OTPLess ID tokens and retrieve user information for integration into your
backend GoLang application.
