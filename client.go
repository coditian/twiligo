package twiligo

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

var (
	base          = "https://api.twilio.com/2010-04-01"
	DefaultClient = &http.Client{
		Transport: &http.Transport{
			TLSHandshakeTimeout: 4 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				MinVersion:         tls.VersionTLS10,
				MaxVersion:         tls.VersionTLS10,
			},
			DisableCompression: false,
			DisableKeepAlives:  true,
		},
	}
)

type Client struct {
	accountSid string
	authToken  string
	client     *http.Client
}

func NewClient(accountSid, authToken string, client *http.Client) Client {
	if client == nil {
		client = DefaultClient
	}
	return Client{
		accountSid: accountSid,
		authToken:  authToken,
		client:     client,
	}
}

func unmarshal(res *http.Response, v interface{}) error {
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "read body")
	}

	if res.StatusCode >= 400 {
		ext := Exception{}
		err = json.Unmarshal(b, &ext)
		if err != nil {
			return errors.Wrap(err, "unmarshal exception,"+string(b))
		}
		return errors.New(ext.Message)
	}

	if v == nil {
		return nil
	}

	err = json.Unmarshal(b, v)
	if err != nil {
		return errors.Wrap(err, "unmarshal")
	}

	return nil
}

type Exception struct {
	Status   int    `json:"status"`
	Message  string `json:"message"`
	Code     int    `json:"code"`
	MoreInfo string `json:"more_info"`
}
