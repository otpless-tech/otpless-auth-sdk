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

type UserDetail struct{}
