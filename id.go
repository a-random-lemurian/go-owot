package owot

import (
	"encoding/json"
	"time"
)

// go-owot parses all IDs as strings, even if they are purely numeric.
// This is future-proofing in anticipation of IP address-based hashes
// being used as IDs.

type ID string

func (id *ID) UnmarshalJSON(data []byte) error {
	if len(data) >= 2 && data[0] == '"' && data[len(data)-1] == '"' {
		*id = ID(data[1 : len(data)-1])
		return nil
	}
	*id = ID(string(data))
	return nil
}

func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(id))
}

type UnixMillis time.Time

func (t *UnixMillis) UnmarshalJSON(data []byte) error {
	var ms int64
	if err := json.Unmarshal(data, &ms); err != nil {
		return err
	}
	*t = UnixMillis(time.UnixMilli(ms))
	return nil
}

func (t UnixMillis) Time() time.Time {
	return time.Time(t)
}
