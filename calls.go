package twiligo

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

type Calls struct {
	Sid             string  `json:"sid"`
	DateCreated     string  `json:"date_created"`
	DateUpdated     string  `json:"date_updated"`
	ParentCallSid   string  `json:"parent_call_sid"`
	AccountSid      string  `json:"account_sid"`
	To              string  `json:"to"`
	FormattedTo     string  `json:"formatted_to"`
	From            string  `json:"from"`
	FormattedFrom   string  `json:"formatted_from"`
	PhoneNumberSid  string  `json:"phone_number_sid"`
	Status          string  `json:"status"`
	StartTime       string  `json:"start_time"`
	EndTime         string  `json:"end_time"`
	Duration        string  `json:"duration"`
	Price           float64 `json:"price"`
	Direction       string  `json:"direction"`
	AnsweredBy      string  `json:"answered_by"`
	APIVersion      string  `json:"api_version"`
	ForwardedFrom   string  `json:"forwarded_from"`
	CallerName      string  `json:"caller_name"`
	URI             string  `json:"uri"`
	SubresourceUris struct {
		Notifications string `json:"notifications"`
		Recordings    string `json:"recordings"`
	} `json:"subresource_uris"`
}

// MakeCall function make a call to twilio
func (c Client) MakeCall(from, to, callback string) (Calls, error) {
	var out Calls
	v := url.Values{}
	v.Set("From", from)
	v.Set("To", to)
	v.Set("Url", callback)

	endpoint := base + "/Accounts/" + c.accountSid + "/Calls.json"

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return Calls{}, errors.Wrap(err, "new request")
	}

	res, err := c.client.Do(req)
	if err != nil {
		return Calls{}, errors.Wrap(err, "client do")
	}

	err = unmarshal(res, &out)

	return out, err
}
