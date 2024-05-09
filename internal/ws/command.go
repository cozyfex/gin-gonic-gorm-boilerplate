package ws

type Command struct {
	Action  string `json:"action"`
	Room    string `json:"room"`
	Message string `json:"message"`
}
