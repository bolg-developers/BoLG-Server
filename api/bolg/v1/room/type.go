package room

type Room struct {
	Id               string
	IsUsing          bool
	IsBattleStarting bool
}

type Rooms []*Room

type Player struct {
	RoomId string
	Name   string `validate:"min=1,max=8"`
	IrId   int32  `validate:"min=1,max=15"`
	Hp     int32  `validate:"min=0"`
	Power  int32  `validate:"min=1,max=15"`
}

type Players []*Player
