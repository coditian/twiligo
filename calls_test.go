package twiligo

import "testing"
import "fmt"
import "net/http"
import "net/http/httptest"
import "reflect"

func TestMakeCall(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()
	base = server.URL

	mux.HandleFunc("/Accounts/AC1111/Calls.json", func(w http.ResponseWriter, r *http.Request) {
		response := `{
			"sid": "CAa346467ca321c71dbd5e12f627deb854",
			"date_created": "Thu, 19 Aug 2010 00:12:15 +0000",
			"date_updated": "Thu, 19 Aug 2010 00:12:15 +0000",
			"account_sid": "AC1111",
			"to": "%s",
			"formatted_to": "(415) 555-1212",
			"from": "%s",
			"formatted_from": "(415) 867-5309",
			"phone_number_sid": "PNd6b0e1e84f7b117332aed2fd2e5bbcab",
			"status": "queued"
		}`
		fmt.Fprint(w, response)
	})

	c := NewClient("AC1111", "AAAA", nil)
	call, err := c.MakeCall("+12345", "client:jenny", "http://10.0.0.1")
	if err != nil {
		t.Error(err)
	}

	expected := Calls{
		Sid:            "CAa346467ca321c71dbd5e12f627deb854",
		DateCreated:    "Thu, 19 Aug 2010 00:12:15 +0000",
		DateUpdated:    "Thu, 19 Aug 2010 00:12:15 +0000",
		ParentCallSid:  "",
		AccountSid:     "AC1111",
		To:             "%s",
		FormattedTo:    "(415) 555-1212",
		From:           "%s",
		FormattedFrom:  "(415) 867-5309",
		PhoneNumberSid: "PNd6b0e1e84f7b117332aed2fd2e5bbcab",
		Status:         "queued",
		StartTime:      "",
		EndTime:        "",
		Duration:       "",
		Price:          0,
		Direction:      "",
		AnsweredBy:     "",
		APIVersion:     "",
		ForwardedFrom:  "",
		CallerName:     "",
		URI:            "",
		SubresourceUris: struct {
			Notifications string "json:\"notifications\""
			Recordings    string "json:\"recordings\""
		}{
			Notifications: "",
			Recordings:    "",
		},
	}

	if !reflect.DeepEqual(call, expected) {
		t.Error(call)
	}
}
