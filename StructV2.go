package otplessAuthSdk

type SendOTPRequestV2 struct {
	PhoneNumber string                 `json:"phoneNumber,omitempty"`
	Email       string                 `json:"email,omitempty"`
	OtpLength   int                    `json:"otpLength,omitempty"`
	OtpHash     string                 `json:"otpHash,omitempty"`
	Channels    []string               `json:"channels,omitempty"`
	Expiry      int                    `json:"expiry,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type SendOTPResponseV2 struct {
	RequestId string `json:"requestId,omitempty"`
	Message   string `json:"message,omitempty"`
}

type VerifyOTPRequestV2 struct {
	Otp       string `json:"otp"`
	RequestId string `json:"requestId"`
}

type VerifyOTPResponseV2 struct {
	RequestId     string `json:"requestId,omitempty"`
	IsOTPVerified bool   `json:"isOTPVerified,omitempty"`
	Message       string `json:"message,omitempty"`
}

type InitiateOTPLinkRequest struct {
	PhoneNumber string                 `json:"phoneNumber,omitempty"`
	Email       string                 `json:"email,omitempty"`
	RedirectURI string                 `json:"redirectURI"`
	Channels    []string               `json:"channels"`
	Expiry      int                    `json:"expiry,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type InitiateOTPLinkResponse struct {
	RequestId string `json:"requestId,omitempty"`
	Message   string `json:"message,omitempty"`
}

type InitiateMagicLinkRequestV2 struct {
	PhoneNumber string                 `json:"phoneNumber,omitempty"`
	Email       string                 `json:"email,omitempty"`
	Channels    []string               `json:"channels"`
	RedirectURI string                 `json:"redirectURI"`
	Expiry      int                    `json:"expiry"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type InitiateMagicLinkResponseV2 struct {
	RequestId bool   `json:"requestId,omitempty"`
	Message   string `json:"message,omitempty"`
}
