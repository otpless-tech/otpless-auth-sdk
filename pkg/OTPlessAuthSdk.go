package otplessAuthSdk

import (
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

	"github.com/dgrijalva/jwt-go"
)

func (u UserDetail) DecodeIDToken(idToken, clientID, clientSecret, audience string) (*UserDetailResult, error) {
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
	authDetails := make(map[string]interface{})
	err = json.Unmarshal([]byte(decoded["authentication_details"].(string)), &authDetails)
	if err != nil {
		return nil, err
	}
	authTime, err := strconv.ParseInt(decoded["auth_time"].(string), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing auth_time: %v", err)
	}
	var userDetail UserDetailResult
	userDetail.Success = true
	userDetail.Iss = decoded["iss"].(string)
	userDetail.Sub = decoded["sub"].(string)
	userDetail.Aud = decoded["aud"].(string)
	userDetail.Exp = int64(decoded["exp"].(float64)) * 1000
	userDetail.Iat = int64(decoded["iat"].(float64)) * 1000
	userDetail.AuthTime = authTime
	userDetail.PhoneNumber = decoded["phone_number"].(string)
	userDetail.Email = decoded["email"].(string)
	userDetail.Name = decoded["name"].(string)

	if authDetails, ok := decoded["authentication_details"].(map[string]interface{}); ok {
		jsonData, err := json.Marshal(authDetails)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(jsonData, &userDetail.AuthenticationDetail)
		if err != nil {
			return nil, err
		}
	}
	return &userDetail, nil
}

func (u UserDetail) VerifyCode(code, clientID, clientSecret string) (*UserDetailResult, error) {
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

		return u.DecodeIDToken(respJson["id_token"].(string), clientID, clientSecret, audience)
	}

	return nil, errors.New(fmt.Sprintf("Request failed with status code %d", response.StatusCode))
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
