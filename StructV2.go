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
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
}
