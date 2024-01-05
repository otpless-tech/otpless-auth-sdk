package otplessAuthSdk

type UserDetailResult struct {
	Success             bool   `json:"success"`
	AuthTime            int64  `json:"auth_time"`
	Email               string `json:"email"`
	Name                string `json:"name"`
	PhoneNumber         string `json:"phone_number"`
	CountryCode         string `json:"country_code"`
	NationalPhoneNumber string `json:"national_phone_number"`
}

type MagicLinkResponse struct {
	RequestIds []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"requestIds"`
}

type SendOTPRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Channel     string `json:"channel"`
	Hash        string `json:"hash"`
	OrderId     string `json:"orderId"`
	Expiry      int    `json:"expiry"`
	OtpLength   int    `json:"otpLength"`
}

type SendOTPResponse struct {
	OrderID string `json:"orderId"`
}

type ResendOTPRequest struct {
	OrderID string `json:"orderId"`
}

type ResendOTPResponse struct {
	OrderID string `json:"orderId"`
}

type VerifyOTPRequest struct {
	OrderID     string `json:"orderId"`
	OTP         string `json:"otp"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type VerifyOTPResponse struct {
	IsOTPVerified bool   `json:"isOTPVerified"`
	Reason        string `json:"reason,omitempty"`
}
