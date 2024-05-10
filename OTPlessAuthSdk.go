package otplessAuthSdk

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func DecodeIDToken(idToken, clientID, clientSecret, audience string) (*UserDetailResult, error) {
	oidcConfig, err := getConfig()
	if err != nil {
		return nil, err
	}

	publicKey, err := getPublicKey(oidcConfig["jwks_uri"].(string))
	if err != nil {
		return nil, err
	}

	n := publicKey["n"].(string)
	e := publicKey["e"].(string)

	decoded, err := decodeJWT(idToken, n, e, oidcConfig["issuer"].(string), audience)
	if err != nil {
		return nil, err
	}

	authTime, err := strconv.ParseInt(decoded["auth_time"].(string), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing auth_time: %v", err)
	}
	var userDetail UserDetailResult
	userDetail.Email = ""
	userDetail.Name = ""
	userDetail.PhoneNumber = ""
	userDetail.Success = true
	userDetail.AuthTime = authTime
	if phone_number, ok := decoded["phone_number"].(string); ok {
		userDetail.PhoneNumber = phone_number
	}
	if email, ok := decoded["email"].(string); ok {
		userDetail.Email = email
	}
	if name, ok := decoded["name"].(string); ok {
		userDetail.Name = name
	}
	if nationalPh, ok := decoded["national_phone_number"].(string); ok {
		userDetail.NationalPhoneNumber = nationalPh
	}
	if cCode, ok := decoded["country_code"].(string); ok {
		userDetail.CountryCode = cCode
	}

	return &userDetail, nil
}

func VerifyCode(code, clientID, clientSecret string) (*UserDetailResult, error) {
	audience := ""
	oidcConfig, err := getConfig()
	if err != nil {
		return nil, err
	}

	form := url.Values{}
	form.Set("code", code)
	form.Set("client_id", clientID)
	form.Set("client_secret", clientSecret)

	var httpClient = &http.Client{
		Timeout: HTTP_TIMEOUT,
	}
	response, err := httpClient.PostForm(oidcConfig["token_endpoint"].(string), form)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		var respJson map[string]interface{}
		json.Unmarshal(body, &respJson)

		return DecodeIDToken(respJson["id_token"].(string), clientID, clientSecret, audience)
	}

	return nil, errors.New(fmt.Sprintf("Request failed with status code %d", response.StatusCode))
}
func VerifyAuthToken(token, clientID, clientSecret string) (*UserDetailResult, error) {

	formData := url.Values{}
	formData.Set("token", token)
	formData.Set("client_id", clientID)
	formData.Set("client_secret", clientSecret)

	client := &http.Client{
		Timeout: HTTP_TIMEOUT, // Make sure you've defined this constant in your package
	}

	req, err := http.NewRequest("POST", OIDC_AUTH_TOKEN_API, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d %v", resp.StatusCode, string(body))
	}
	var decoded map[string]interface{}
	json.Unmarshal(body, &decoded)

	authTime, err := strconv.ParseInt(decoded["auth_time"].(string), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing auth_time: %v", err)
	}
	userDetail := UserDetailResult{
		Success:  true,
		AuthTime: authTime,
	}
	userDetail.Email = ""
	userDetail.Name = ""
	userDetail.PhoneNumber = ""

	if phoneNumber, ok := decoded["phone_number"].(string); ok {
		userDetail.PhoneNumber = phoneNumber
	}
	if email, ok := decoded["email"].(string); ok {
		userDetail.Email = email
	}
	if name, ok := decoded["name"].(string); ok {
		userDetail.Name = name
	}
	if countryCode, ok := decoded["country_code"].(string); ok {
		userDetail.CountryCode = countryCode
	}
	if nationalPhoneNumber, ok := decoded["national_phone_number"].(string); ok {
		userDetail.NationalPhoneNumber = nationalPhoneNumber
	}
	return &userDetail, nil
}
func GenerateMagicLink(mobileNumber, email, clientID, clientSecret, redirectUri, channel string) (*MagicLinkResponse, error) {
	params := url.Values{}
	params.Add("client_id", clientID)
	params.Add("client_secret", clientSecret)
	params.Add("redirect_uri", redirectUri)

	if mobileNumber != "" {
		params.Add("mobile_number", mobileNumber)
	}

	if email != "" {
		params.Add("email", email)
	}

	if channel != "" {
		params.Add("channel", channel)
	}

	fullURL := fmt.Sprintf("%s?%s", MAGIC_LINK_URL, params.Encode())

	response, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("error making the request: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var resp MagicLinkResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return &resp, nil
}
func decodeJWT(jwtToken, modulus, exponent, issuer, audience string) (jwt.MapClaims, error) {
	publicKey, err := createRSAPublicKey(modulus, exponent)
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(jwtToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Invalid JWT Token")
}

func createRSAPublicKey(modulusB64, exponentB64 string) (*rsa.PublicKey, error) {
	modulus, err := base64UrlDecode(modulusB64)
	if err != nil {
		return nil, err
	}

	exponent, err := base64UrlDecode(exponentB64)
	if err != nil {
		return nil, err
	}

	return &rsa.PublicKey{
		N: new(big.Int).SetBytes(modulus),
		E: int(new(big.Int).SetBytes(exponent).Int64()),
	}, nil
}

func base64UrlDecode(base64UrlString string) ([]byte, error) {
	padded := base64UrlString
	switch len(base64UrlString) % 4 {
	case 2:
		padded += "=="
	case 3:
		padded += "="
	}

	return base64.URLEncoding.DecodeString(padded)
}

func getConfig() (map[string]interface{}, error) {
	client := &http.Client{Timeout: HTTP_TIMEOUT}
	response, err := client.Get(OTPLESS_KEY_API)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var respJson map[string]interface{}
	if err := json.Unmarshal(body, &respJson); err != nil {
		return nil, err
	}

	return respJson, nil
}

func getPublicKey(url string) (map[string]interface{}, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var respJson map[string]interface{}
	json.Unmarshal(body, &respJson)

	if keys, present := respJson["keys"].([]interface{}); present && len(keys) > 0 {
		return keys[0].(map[string]interface{}), nil
	}

	return nil, errors.New("Unable to fetch public key")
}

func SendOTP(req SendOTPRequest, clientID, clientSecret string) (*SendOTPResponse, error) {
	url := OTP_BASE_URL + "/v1/send"
	headers := map[string]string{
		"clientId":     clientID,
		"clientSecret": clientSecret,
		"Content-Type": "application/json",
	}

	data := SendOTPRequest{
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Channel:     req.Channel,
		Hash:        req.Hash,
		OrderId:     req.OrderId,
		Expiry:      req.Expiry,
		OtpLength:   req.OtpLength,
		TemplateId:  req.TemplateId,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshalling data: %v", err)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d %v", response.StatusCode, string(body))
	}

	var otpResponse SendOTPResponse
	err = json.Unmarshal(body, &otpResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %v", err)
	}

	return &otpResponse, nil
}

func ResendOTP(orderID, clientID, clientSecret string) (*ResendOTPResponse, error) {
	url := OTP_BASE_URL + "/v1/resend"
	headers := map[string]string{
		"clientId":     clientID,
		"clientSecret": clientSecret,
		"Content-Type": "application/json",
	}

	data := ResendOTPRequest{
		OrderID: orderID,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshalling data: %v", err)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d %v", response.StatusCode, string(body))
	}

	var otpResponse ResendOTPResponse
	err = json.Unmarshal(body, &otpResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %v", err)
	}

	return &otpResponse, nil
}

func VerifyOTP(orderID, otp, email, phoneNumber, clientID, clientSecret string) (*VerifyOTPResponse, error) {
	url := OTP_BASE_URL + "/v1/verify"
	headers := map[string]string{
		"clientId":     clientID,
		"clientSecret": clientSecret,
		"Content-Type": "application/json",
	}

	data := VerifyOTPRequest{
		OrderID:     orderID,
		OTP:         otp,
		Email:       email,
		PhoneNumber: phoneNumber,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshalling data: %v", err)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d %v", response.StatusCode, string(body))
	}

	var otpResponse VerifyOTPResponse
	err = json.Unmarshal(body, &otpResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %v", err)
	}

	return &otpResponse, nil
}
