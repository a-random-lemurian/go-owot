package owot

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type OwotConn struct {
	ws         *websocket.Conn
	HandleRaw  func([]byte)
	HandleChat func(*MessageChat)
}

// Connect to the OWOT server.
//
// The URL argument must be a valid WebSocket URL, such as <wss://ourworldoftext.com/go-owot/ws/>
func Dial(url string) (*OwotConn, error) {
	headers := http.Header{}
	ws, _, err := websocket.DefaultDialer.Dial(url, headers)

	if err != nil {
		return nil, err
	}

	conn := &OwotConn{
		ws: ws,
	}

	conn.initFuncs()

	return conn, nil
}

// Set all handler functions to nil.
func (o *OwotConn) initFuncs() {
	o.HandleRaw = func(b []byte) {}
	o.HandleChat = func(mc *MessageChat) {}
}

// Begin reading messages from the websocket. Blocks until an error is encountered
// when reading messages from the websocket.
func (o *OwotConn) Run() error {
	// Partially marshal the JSON file into a "partial" struct
	// that reads the "kind" field, before fully parsing it and
	// routing it to the appropriate function.
	type partial struct {
		Kind string `json:"kind"`
	}

	var message partial

	for {
		// For HandleRaw(), we want to pass messages along exactly the way they are,
		// so we don't use ReadJSON().
		_, msgBytes, err := o.ws.ReadMessage()
		// TODO better error handling
		if err != nil {
			return err
		}
		o.HandleRaw(msgBytes)

		err = json.Unmarshal(msgBytes, &message)
		if err != nil {
			log.Printf("%v", err)
			continue
		}

		switch message.Kind {
		case "chat":
			var chat MessageChat
			err = json.Unmarshal(msgBytes, &chat)
			if err == nil {
				o.HandleChat(&chat)
			}
		default:
			continue
		}

		if err != nil {
			return err
		}
	}
}
