package twiligo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetPaticipantObject(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()
	base = server.URL

	mux.HandleFunc("/Accounts/ACb04f7f58411fb4e1130389a1b90f8bca/Conferences/CFbbe46ff1274e283f7e3ac1df0072ab39/Participants/CA386025c9bf5d6052a1d1ea42b4d16662", func(w http.ResponseWriter, r *http.Request) {
		response := `{
			"account_sid": "ACb04f7f58411fb4e1130389a1b90f8bca",
			"call_sid": "CA386025c9bf5d6052a1d1ea42b4d16662",
			"conference_sid": "CFbbe46ff1274e283f7e3ac1df0072ab39",
			"date_created": "Wed, 18 Aug 2010 20:20:10 +0000",
			"date_updated": "Wed, 18 Aug 2010 20:20:10 +0000",
			"end_conference_on_exit": true,
			"muted": false,
			"hold": false,
			"start_conference_on_enter": true,
			"uri": "/2010-04-01/Accounts/ACb04f7f58411fb4e1130389a1b90f8bca/Conferences/CFbbe46ff1274e283f7e3ac1df0072ab39/Participants/CA386025c9bf5d6052a1d1ea42b4d16662.json"
		}`
		fmt.Fprint(w, response)
	})

	c := NewClient("ACb04f7f58411fb4e1130389a1b90f8bca", "AAAA", nil)
	paticipant, err := c.Paticipant("CFbbe46ff1274e283f7e3ac1df0072ab39", "CA386025c9bf5d6052a1d1ea42b4d16662")
	if err != nil {
		t.Error(err)
		return
	}

	expected := Paticipant{
		AccountSid:          "ACb04f7f58411fb4e1130389a1b90f8bca",
		CallSid:             "CA386025c9bf5d6052a1d1ea42b4d16662",
		ConferenceSid:       "CFbbe46ff1274e283f7e3ac1df0072ab39",
		DateCreated:         "Wed, 18 Aug 2010 20:20:10 +0000",
		DateUpdated:         "Wed, 18 Aug 2010 20:20:10 +0000",
		EndConferenceOnExit: true,
		Muted:               false,
		Hold:                false,
		StartConferenceOnEnter: true,
		URI: "/2010-04-01/Accounts/ACb04f7f58411fb4e1130389a1b90f8bca/Conferences/CFbbe46ff1274e283f7e3ac1df0072ab39/Participants/CA386025c9bf5d6052a1d1ea42b4d16662.json",
	}

	if !reflect.DeepEqual(paticipant, expected) {
		t.Error("%#v\n", paticipant)
	}
}
