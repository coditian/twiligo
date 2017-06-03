package twiligo

import (
	"net/http"

	"github.com/pkg/errors"
)

type Conferences struct {
	FirstPageURI string `json:"first_page_uri"`
	End          int    `json:"end"`
	Conferences  []struct {
		AccountSid      string `json:"account_sid"`
		APIVersion      string `json:"api_version"`
		DateCreated     string `json:"date_created"`
		DateUpdated     string `json:"date_updated"`
		FriendlyName    string `json:"friendly_name"`
		Region          string `json:"region"`
		SID             string `json:"sid"`
		Status          string `json:"status"`
		URI             string `json:"uri"`
		SubresourceUris struct {
			Participants string `json:"participants"`
		} `json:"subresource_uris"`
	} `json:"conferences"`
	PageSize int    `json:"page_size"`
	Start    int    `json:"start"`
	Page     int    `json:"page"`
	URI      string `json:"uri"`
}

func (c Client) Conference(friendlyName string) (Conferences, error) {
	var out Conferences

	endpoint := base + "/Accounts/" + c.accountSid + "/Conferences.json?FriendlyName=" + friendlyName
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return Conferences{}, errors.Wrap(err, "request twilio hold")
	}
	req.SetBasicAuth(c.accountSid, c.authToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.client.Do(req)
	if err != nil {
		return Conferences{}, errors.Wrap(err, "request to twilio")
	}

	err = unmarshal(res, &out)

	return out, err
}
