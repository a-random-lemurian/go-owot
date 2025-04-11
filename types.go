package owot

// {"kind":"chat","nickname":"ztunedd","realUsername":"discor","id":8980,"message":"i can read that smh","registered":true,"location":"global","op":false,"admin":false,"staff":false,"color":"#5865F2","customMeta":{"discordUserId":"1128340150477398106"},"rankName":"Discord","rankColor":"#05A72C5","date":1744398348259}

type MessageChat struct {
	Kind string
	Nickname string
	RealUsername string
	
	// The user's ID. Rendered as a string due to futureproofing for
	// IP-based hash IDs.
	Id ID

	Message string
	Registered bool
	Location string
	Op bool
	Admin bool
	Staff bool
	Color string
	CustomMeta interface{}
	RankName string
	RankColor string

	// Over the wire, the date format is in milliseconds.
	Date UnixMillis
}
