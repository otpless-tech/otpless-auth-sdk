package otplessAuthSdk

import "time"

const (
	OTPLESS_KEY_API     = "https://otpless.com/.well-known/openid-configuration"
	HTTP_TIMEOUT        = 30 * time.Second
	OIDC_AUTH_TOKEN_API = "https://oidc.otpless.app/auth/userInfo"
	MAGIC_LINK_URL      = "https://auth.otpless.app/auth/v1/authorize"
	OTP_BASE_URL        = "https://user-auth.otpless.app/auth/otp"
)
