package otplessAuthSdk

import "time"

const (
	OTPLESS_KEY_API = "https://otpless.com/.well-known/openid-configuration"
	HTTP_TIMEOUT    = 30 * time.Second
	OIDC_AUTH_TOKEN_API = "https://oidc.otpless.app/auth/userInfo"
)