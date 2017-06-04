package twiligo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetConferencesByFriendlyName(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()
	base = server.URL

	mux.HandleFunc("/Accounts/AC1111/Conferences.json", func(w http.ResponseWriter, r *http.Request) {
		response := `{
			"conferences": [
				{
					"sid": "CFb86aba72677cbf0f18f616768abe68ca",
					"friendly_name": "CAc10577ebfc7b2deb081fed19bb1f2ba6",
					"status": "completed",
					"date_created": "Thu, 01 Jun 2017 16:36:11 +0000",
					"date_updated": "Thu, 01 Jun 2017 16:39:25 +0000"
				}
			]
		}`
		fmt.Fprint(w, response)
	})

	c := NewClient("AC1111", "AAAA", nil)
	conf, err := c.Conference("Jenny")
	if err != nil {
		t.Error(err)
		return
	}

	expected := Conferences{
		Conferences: []conference{
			{
				SID:          "CFb86aba72677cbf0f18f616768abe68ca",
				FriendlyName: "CAc10577ebfc7b2deb081fed19bb1f2ba6",
				Status:       "completed",
				DateCreated:  "Thu, 01 Jun 2017 16:36:11 +0000",
				DateUpdated:  "Thu, 01 Jun 2017 16:39:25 +0000",
			},
		},
	}

	if !reflect.DeepEqual(conf, expected) {
		t.Error("%#v\n", conf)
	}
}
