package amplitude

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func ExampleEvent() {
	keyResp, bodyResp, server := mockServer("event")
	defer server.Close()

	client := New("s3cr3ts")
	client.endpoint = server.URL

	client.Event(Event{
		UserId:    "0000001",
		EventType: "joined",
	})

	key := <-keyResp
	body := <-bodyResp
	fmt.Printf("Key: %s\n%s", string(key), string(body))

	// Output:
	// Key: s3cr3ts
	// {
	//   "event_type": "joined",
	//   "user_id": "0000001"
	// }
}

func ExampleIdentify() {
	keyResp, bodyResp, server := mockServer("identification")
	defer server.Close()

	client := New("s3cr3ts")
	client.endpoint = server.URL

	client.Identify(Identify{
		UserId: "0000001",
		UserProperties: map[string]interface{}{
			"name":  "Art Vandelay",
			"email": "art@vandelayindustries.com",
		},
	})

	key := <-keyResp
	body := <-bodyResp
	fmt.Printf("Key: %s\n%s", string(key), string(body))

	// Output:
	// Key: s3cr3ts
	// {
	//   "user_id": "0000001",
	//   "user_properties": {
	//     "email": "art@vandelayindustries.com",
	//     "name": "Art Vandelay"
	//   }
	// }
}

func mockServer(msgKey string) (chan []byte, chan []byte, *httptest.Server) {
	key, body := make(chan []byte, 1), make(chan []byte, 1)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		k := r.FormValue("api_key")
		id := r.FormValue(msgKey)

		var v interface{}
		err := json.Unmarshal([]byte(id), &v)
		if err != nil {
			panic(err)
		}

		b, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			panic(err)
		}

		key <- []byte(k)
		body <- b
	}))

	return key, body, server
}
