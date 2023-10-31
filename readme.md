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
