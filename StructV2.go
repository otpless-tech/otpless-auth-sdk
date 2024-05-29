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

type VerifyCodeRequest struct {
	Code string `json:"code"`
}

type VerifyCodeResponse struct {
	RequestId   string       `json:"requestId,omitempty"`
	Message     string       `json:"message,omitempty"`
	Description string       `json:"description,omitempty"`
	UserDetails *UserDetails `json:"userDetails,omitempty"`
}

type UserDetails struct {
	Token      string     `json:"token,omitempty"`
	Timestamp  string     `json:"timestamp,omitempty"`
	Identities []Identity `json:"identities,omitempty"`
	Network    Network    `json:"network,omitempty"`
	DeviceInfo DeviceInfo `json:"deviceInfo,omitempty"`
}

type Identity struct {
	IdentityType  string   `json:"identityType,omitempty"`
	IdentityValue string   `json:"identityValue,omitempty"`
	Channel       string   `json:"channel,omitempty"`
	Methods       []string `json:"methods,omitempty"`
	Verified      bool     `json:"verified,omitempty"`
	VerifiedAt    string   `json:"verifiedAt,omitempty"`
}

type Network struct {
	IP         string     `json:"ip,omitempty"`
	Timezone   string     `json:"timezone,omitempty"`
	IPLocation IPLocation `json:"ipLocation,omitempty"`
}

type IPLocation struct {
	City         LocationName `json:"city,omitempty"`
	Subdivisions LocationName `json:"subdivisions,omitempty"`
	Country      LocationName `json:"country,omitempty"`
	Continent    LocationCode `json:"continent,omitempty"`
	Latitude     float64      `json:"latitude,omitempty"`
	Longitude    float64      `json:"longitude,omitempty"`
	PostalCode   string       `json:"postalCode,omitempty"`
}

type LocationName struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

type LocationCode struct {
	Code string `json:"code,omitempty"`
}

type DeviceInfo struct {
	UserAgent string `json:"userAgent,omitempty"`
}
type InitiateOAuthRequest struct {
	Channel     string                 `json:"channel"`
	Channels    []string               `json:"channels"`
	RedirectURI string                 `json:"redirectURI"`
	Expiry      int                    `json:"expiry"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type InitiateOAuthResponse struct {
	RequestId string `json:"requestId,omitempty"`
	Link      string `json:"link,omitempty"`
}

type CheckStatusRequest struct {
	RequestId string `json:"requestId"`
}

type CheckStatusResponse struct {
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
}
