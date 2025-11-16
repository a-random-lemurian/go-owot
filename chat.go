// {"kind":"chat","nickname":"port and starboard elizabeth","message":"test","location":"page","color":"#008ec4"}
package owot

import (
	"context"
	"fmt"

	"golang.org/x/time/rate"
)

type PrivateMessageType int

const (
	PMTFromMe = iota
	PMTToMe
)

func (p PrivateMessageType) String() string {
	switch p {
	case PMTFromMe:
		return "from_me"
	case PMTToMe:
		return "to_me"
	default:
		return ""
	}
}

type ChatLocation int

const (
	CLPage = iota
	CLGlobal
)

func (c ChatLocation) String() string {
	switch c {
	case CLPage:
		return "page"
	case CLGlobal:
		return "global"
	default:
		return ""
	}
}

type MessageChat struct {
	Kind         string
	Nickname     string
	RealUsername string

	// The user's ID. Rendered as a string due to futureproofing for
	// IP-based hash IDs.
	Id ID

	Message        string
	Registered     bool
	Location       ChatLocation
	Op             bool
	Admin          bool
	Staff          bool
	Color          string
	CustomMeta     map[string]any
	RankName       string
	RankColor      string
	PrivateMessage PrivateMessageType

	// Over the wire, the date format is in milliseconds.
	Date UnixMillis
}

var defaultChatRatelimiter = rate.NewLimiter(rate.Limit(2), 1) // two messages per second

// Send a chat message.
//
// - nickname   : The nickname to use when sending the message. If this string is empty, it defaults to your username.
// - message    : The message to send.
// - location   : Either "page" or "global".
// - color      : Hex code of nickname color. Only works when logged in.
// - customMeta : Custom JSON data to send with the message.
func (o *OwotConn) Chat(nickname string, message string, location string, color string, customMeta any) error {
	ctx := context.Background()
	o.chatRate.Wait(ctx)

	chat := map[string]any{}

	if location != "global" && location != "page" {
		return fmt.Errorf("unsupported location: %s", location)
	}

	chat["kind"] = "chat"
	chat["nickname"] = nickname
	chat["message"] = message
	chat["location"] = location
	chat["color"] = "#008ec4"
	chat["customMeta"] = customMeta

	o.ws.WriteJSON(chat)

	return nil
}
