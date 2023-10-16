# Merchant Integration Documentation(Backend GoLang Auth SDK)

---

> ## A. OTPLessAuth

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
VerifyCode(code, clientID, clientSecret)
```

#### Method Params:

| Params       | Data type | Mandatory | Constraints | Remarks                                                                      |
| ------------ | --------- | --------- | ----------- | ---------------------------------------------------------------------------- |
| idToken      | String    | true      |             | idToken which is JWT token which you get from `OTPLess` by exchange code API |
| clientId     | String    | true      |             | Your OTPLess `Client Id`                                                     |
| clientSecret | String    | true      |             | Your OTPLess `Client Secret`                                                 |

#### Return

Return:
Object Name: UserDetailResult

```json
userDetail := UserDetailResult{
Aud:  "kp***i",
AuthTime:  1697473071,
AuthenticationDetail: AuthenticationDetail{
Phone: PhoneDetail{
Mode:        "",
PhoneNumber: "",
CountryCode: "",
AuthState:   "",
},
Email: EmailDetail{
Email:     "dev****@gmail.com",
Mode:      "",
AuthState: "",
},
},
Email:       "dev****@gmail.com",
Exp:         1697456877000,
Iat:         1697453277000,
Iss:         "https://otpless.com",
Name:        "Dhaval From OTP-less",
PhoneNumber: "+9193*****",
Sub:         "579",
Success:     true,
}
```
> ### UserDetail Object Fields:
>
> `success` (boolean): This will be `true` in case of method successfully performed operation.<br> > `iss` (String, required): The issuer, which should be "https://otpless.com."<br> > `sub` (String, required): The subject, typically the user ID.<br> > `aud` (String, required): The audience, your OTPLess Client ID.<br> > `exp` (Long, required): The expiration time of the token.<br> > `iat` (Long, required): The issuance time of the token.<br> > `authTime` (Long, required): The time when authentication was completed.<br> > `phoneNumber` (String, required): The user's phone number.<br> > `email` (String, required): The user's email address.<br> > `name` (String, required): The user's full name.<br> > `authenticationDetails` (AuthenticationDetails, required): Authentication details containing information about phone
> verification and email verification.

#### AuthenticationDetails Object Fields:

`phone` (PhoneAuthentication, required): Information about phone verification.<br>
`email` (EmailAuthentication, required): Information about email verification.

#### PhoneAuthentication Object Fields:

`mode` (String): The mode of phone verification.<br>
`phoneNumber` (String): The verified phone number.<br>
`countryCode` (String): The country code of the phone number.<br>
`authState` (String): The authentication state of the phone number.

#### EmailAuthentication Object Fields:

`email` (String): The verified email address.<br>
`mode` (String): The mode of email verification.<br>
`authState` (String): The authentication state of the email.<br>

### Error case:

`success` (boolean): This will be `false`. The method is failed to perform.<br>
`errorMessage` (String): The message contains error information.<br>

### Example of usage

```go
import OTPLessAuthSdk

func main() {
	
    code := "your_id_token_here"
    clientID := "your_client_id"
    clientSecret := "your_client_secret"
    userDetail := UserDetail{}
    result, err := userDetail.VerifyCode(code, clientID, clientSecret)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
	
    fmt.Printf("User Detail Struct: %+v\n", result)
}


```

This method allows you to decode and verify OTPLess ID tokens and retrieve user information for integration into your
backend GoLang application.
