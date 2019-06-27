package room

import (
	"context"
	"github.com/bolg-developers/BoLG-Server/api/bolg/v1/bolgerr"
	"github.com/bolg-developers/BoLG-Server/genproto/bolg/v1"
	"github.com/bolg-developers/BoLG-Server/internal/id"
	"github.com/bolg-developers/BoLG-Server/internal/token"
	"github.com/bolg-developers/BoLG-Server/server/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	repo       Repository
	playerRepo PlayerRepository
	idMan      map[string]*id.Manager
}

func NewService(roomRepo Repository, playerRepo PlayerRepository) *Service {
	return &Service{
		repo:       roomRepo,
		playerRepo: playerRepo,
		idMan:      map[string]*id.Manager{},
	}
}

func (os *Service) CreateRoom(context.Context, *empty.Empty) (*bolg.Room, error) {
	r := new(Room)
	err := os.repo.Create(r)
	if err != nil {
		return nil, status.New(codes.ResourceExhausted, "ルームはこれ以上作成できません").Err()
	}
	os.idMan[r.Id] = id.NewManager(config.Env().MaxIrId)
	return &bolg.Room{Id: r.Id}, nil
}

func (os *Service) JoinRoom(c context.Context, req *bolg.JoinRoomRequest) (*bolg.JoinRoomResponse, error) {
	if _, err := os.repo.Get(req.RoomId); err != nil {
		// サーバー内部エラーは考慮していない
		return nil, bolgerr.InvalidArgument
	}
	playerId, ok := os.idMan[req.RoomId].Increment()
	if !ok {
		return nil, status.New(codes.ResourceExhausted, "プレイヤーはこれ以上作成できません").Err()
	}
	tkn, err := token.CreateJWT(jwt.MapClaims{
		config.RoomIdClaimsKey: req.RoomId,
		config.IrIdClaimsKey:   playerId,
	}, config.Env().TokenSecretKey)
	if err != nil {
		// TODO: logging
		return nil, bolgerr.Internal
	}

	if err := os.playerRepo.Create(&Player{
		RoomId: req.RoomId,
		Name:   req.Name,
		IrId:   int32(playerId),
		Power:  req.Power,
	}); err != nil {
		if _, ok := os.idMan[req.RoomId].Decrement(); !ok {
			// TODO: logging
			return nil, bolgerr.Internal
		}
		// TODO: logging
		return nil, bolgerr.Internal
	}

	players, err := os.playerRepo.List(req.RoomId)
	if err != nil {
		// TODO: logging
		return nil, bolgerr.Internal
	}

	playersRes := make([]*bolg.Player, len(players))
	for i, p := range players {
		playersRes[i] = &bolg.Player{
			Name: p.Name,
			Hp:   p.Hp,
		}
	}

	return &bolg.JoinRoomResponse{
		Players: playersRes,
		Token:   tkn,
	}, nil
}
