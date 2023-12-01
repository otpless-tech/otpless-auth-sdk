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
	SendTo  string `json:"sendTo"`
	OrderID string `json:"orderId"`
	Hash    string `json:"hash"`
}

type SendOTPResponse struct {
	OrderID string `json:"orderId"`
	RefID   string `json:"refId"`
}

type ResendOTPRequest struct {
	OrderID string `json:"orderId"`
}

type ResendOTPResponse struct {
	OrderID string `json:"orderId"`
	RefID   string `json:"refId"`
}

type VerifyOTPRequest struct {
	OrderID string `json:"orderId"`
	OTP     string `json:"otp"`
	SendTo  string `json:"sendTo"`
}

type VerifyOTPResponse struct {
	IsOTPVerified bool   `json:"isOTPVerified"`
	Reason        string `json:"reason,omitempty"`
}
