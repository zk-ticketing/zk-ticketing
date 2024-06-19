package repos

import (
	"github.com/NFTGalaxy/zk-ticketing-server/repos/events"
	"github.com/NFTGalaxy/zk-ticketing-server/repos/registrations"
	"github.com/NFTGalaxy/zk-ticketing-server/repos/users"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Client struct {
	DBConnPool    *pgxpool.Pool
	Events        *events.Queries
	Registrations *registrations.Queries
	Users         *users.Queries
}

func NewClient(pool *pgxpool.Pool) *Client {
	return &Client{
		DBConnPool:    pool,
		Events:        events.New(pool),
		Registrations: registrations.New(pool),
		Users:         users.New(pool),
	}
}
