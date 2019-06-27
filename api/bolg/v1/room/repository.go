package room

import (
	"github.com/bolg-developers/BoLG-Server/infra"
	"github.com/bolg-developers/BoLG-Server/internal/id"
	"github.com/bolg-developers/BoLG-Server/server/config"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"strconv"
	"sync"
)

type Repository interface {
	Create(*Room) error
	List() (Rooms, error)
	Get(string) (*Room, error)
}

type repository struct {
	sync.Mutex
	rooms Rooms
}

func NewRepository() *repository {
	repo := &repository{
		rooms: make(Rooms, 0, config.Env().MaxRoomCount),
	}
	return repo
}

func (repo *repository) Create(r *Room) error {
	repo.Lock()
	defer repo.Unlock()
	if len(repo.rooms) > config.Env().MaxRoomCount {
		return errors.New("ルーム数が最大数に達しています")
	}
	r.Id = uuid.New().String()
	repo.rooms = append(repo.rooms, r)
	return nil
}

func (repo *repository) List() (Rooms, error) {
	return repo.rooms, nil
}

func (repo *repository) Get(id string) (*Room, error) {
	for _, r := range repo.rooms {
		if r.Id == id {
			return r, nil
		}
	}
	return nil, errors.New("not found")
}

type PlayerRepository interface {
	Create(*Player) error
	List(string) (Players, error)
}

type playerRepository struct {
	sync.Mutex
	roomIdDbMap map[string]*redis.Client
}

func NewPlayerRepository() *playerRepository {
	return &playerRepository{roomIdDbMap: make(map[string]*redis.Client)}
}

func (pr *playerRepository) Create(p *Player) error {
	pr.Lock()
	if _, ok := pr.roomIdDbMap[p.RoomId]; !ok {
		cm, ok := infra.ClientManager.Get()
		if !ok {
			pr.Unlock()
			return errors.New("clientがすべて使用されています")
		}
		pr.roomIdDbMap[p.RoomId] = cm
	}
	pr.Unlock()

	// playerを作成
	pipe := pr.roomIdDbMap[p.RoomId].TxPipeline()
	if err := pipe.HSet(id.ConvertString(p.IrId), "name", p.Name).Err(); err != nil {
		return err
	}
	if err := pipe.HSet(id.ConvertString(p.IrId), "hp", p.Hp).Err(); err != nil {
		return err
	}
	if err := pipe.HSet(id.ConvertString(p.IrId), "power", p.Power).Err(); err != nil {
		return err
	}
	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}

func (pr *playerRepository) List(roomId string) (Players, error) {
	irIdList := pr.roomIdDbMap[roomId].Keys("*").Val()
	players := make(Players, len(irIdList))
	for i, irId := range irIdList {
		playerMap := pr.roomIdDbMap[roomId].HGetAll(irId).Val()
		hp, err := strconv.Atoi(playerMap["hp"])
		if err != nil {
			return nil, err
		}
		power, err := strconv.Atoi(playerMap["power"])
		if err != nil {
			return nil, err
		}
		irId_, err := strconv.Atoi(irId)
		if err != nil {
			return nil, err
		}
		players[i] = &Player{
			RoomId: roomId,
			Name:   playerMap["name"],
			IrId:   int32(irId_),
			Hp:     int32(hp),
			Power:  int32(power),
		}
	}
	return players, nil
}
