package ws

type Hub struct {
	Rooms map[string]*Room
}

var instance *Hub

func NewHub() *Hub {
	if instance == nil {
		instance = &Hub{
			Rooms: make(map[string]*Room),
		}
	}
	return instance
}

func GetHubInstance() *Hub {
	return instance
}

func (h *Hub) Run() {
	// 필요에 따라 허브 관리 로직 추가
}
