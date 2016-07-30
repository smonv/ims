package message

// Error struct
type Error struct {
	Code int8
	Body string
}

// Response struct
type Response struct {
	Error Error       `json:"error"`
	Data  interface{} `json:"data"`
}
