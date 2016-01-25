package gokayako

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Kayako struct {
	ApiKey    string
	SecretKey string
	ApiUrl    string
	Client    *http.Client
}

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (kayako *Kayako) getSalt() string {
	return "01237729"
}

func (kayako *Kayako) buildQuery(base string, params map[string]string) (string, error) {

	u, err := url.Parse(kayako.ApiUrl)
	if err != nil {
		return "", err
	}

	query := u.Query()
	salt := kayako.getSalt()
	sig := ComputeHmac256(salt, kayako.SecretKey)

	query.Set("e", base)
	query.Set("apikey", kayako.ApiKey)
	query.Set("salt", salt)
	query.Set("signature", sig)

	for key, value := range params {
		query.Set(key, value)
	}

	u.RawQuery = query.Encode()

	return u.String(), nil

}

func getURLBody(client *http.Client, url string) ([]byte, error) {

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return []byte(body), nil

}

func (kayako *Kayako) buildAndGetBody(base string, params map[string]string) ([]byte, error) {

	query, err := kayako.buildQuery(base, params)
	if err != nil {
		return nil, err
	}

	body, err := getURLBody(kayako.Client, query)
	if err != nil {
		return nil, err
	}

	return body, nil

}
