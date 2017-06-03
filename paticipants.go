package twiligo

import (
	"net/http"

	"github.com/pkg/errors"
)

type Paticipant struct {
	AccountSid             string `json:"account_sid"`
	CallSid                string `json:"call_sid"`
	ConferenceSid          string `json:"conference_sid"`
	DateCreated            string `json:"date_created"`
	DateUpdated            string `json:"date_updated"`
	EndConferenceOnExit    bool   `json:"end_conference_on_exit"`
	Muted                  bool   `json:"muted"`
	Hold                   bool   `json:"hold"`
	StartConferenceOnEnter bool   `json:"start_conference_on_enter"`
	URI                    string `json:"uri"`
}

func (c Client) Paticipant(conferenceSid, callSid string) (Paticipant, error) {
	var out Paticipant

	endpoint := base + "/Accounts/" + c.accountSid + "/Conferences/" + conferenceSid + "/Participants/" + callSid
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return Paticipant{}, errors.Wrap(err, "request twilio hold")
	}
	req.SetBasicAuth(c.accountSid, c.authToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.client.Do(req)
	if err != nil {
		return Paticipant{}, errors.Wrap(err, "request to twilio")
	}

	err = unmarshal(res, &out)

	return out, err
}
