package otplessAuthSdk

type PhoneDetail struct {
	Mode        string `json:"mode"`
	PhoneNumber string `json:"phone_number"`
	CountryCode string `json:"country_code"`
	AuthState   string `json:"auth_state"`
}

type EmailDetail struct {
	Email     string `json:"email"`
	Mode      string `json:"mode"`
	AuthState string `json:"auth_state"`
}

type AuthenticationDetail struct {
	Phone PhoneDetail `json:"phone"`
	Email EmailDetail `json:"email"`
}

type UserDetailResult struct {
	Aud                  string               `json:"aud"`
	AuthTime             int64                `json:"auth_time"`
	AuthenticationDetail AuthenticationDetail `json:"authentication_details"`
	Email                string               `json:"email"`
	Exp                  int64                `json:"exp"`
	Iat                  int64                `json:"iat"`
	Iss                  string               `json:"iss"`
	Name                 string               `json:"name"`
	PhoneNumber          string               `json:"phone_number"`
	Sub                  string               `json:"sub"`
	Success              bool                 `json:"success"`
}

type UserDetail struct{}
