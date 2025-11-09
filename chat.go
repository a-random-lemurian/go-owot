// {"kind":"chat","nickname":"port and starboard elizabeth","message":"test","location":"page","color":"#008ec4"}
package owot

import (
	"context"
	"fmt"

	"golang.org/x/time/rate"
)

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

	chat := map[string]string{}

	if location != "global" && location != "page" {
		return fmt.Errorf("unsupported location: %s", location)
	}

	chat["kind"] = "chat"
	chat["nickname"] = nickname
	chat["message"] = message
	chat["location"] = location
	chat["color"] = "#008ec4"

	o.ws.WriteJSON(chat)

	return nil
}
