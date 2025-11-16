package owot_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/a-random-lemurian/go-owot"
)

const privateMessage = `{"nickname":"port and starboard elizabeth","realUsername":"lemuria","id":9086,"message":"h","registered":true,"location":"page","op":false,"admin":false,"staff":false,"color":"#008ec4","kind":"chat","privateMessage":"to_me"}`

func TestParsePrivateMessage(t *testing.T) {
	var msg *owot.MessageChat
	if err := json.Unmarshal([]byte(privateMessage), &msg); err != nil {
		t.Fatal(err)
	}
	actual := msg.PrivateMessage.String()
	if actual != "to_me" {
		t.Fatal(fmt.Errorf("expected parsed value of to_me, got %s", actual))
	}
}

const publicMessage = `{"kind":"chat","nickname":"port and starboard elizabeth","realUsername":"lemuria","id":2520,"message":"h","registered":true,"location":"page","op":false,"admin":false,"staff":false,"color":"#008ec4","date":1763273428379}`

func TestParsePublicMessage(t *testing.T) {
	var msg *owot.MessageChat
	if err := json.Unmarshal([]byte(publicMessage), &msg); err != nil {
		t.Fatal(err)
	}
	if msg.PrivateMessage != nil {
		t.Fatal(fmt.Errorf("msg.PrivateMessage should be nil, but it was %+v", msg.PrivateMessage))
	}
}
