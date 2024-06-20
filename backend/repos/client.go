package repos

import (
	"github.com/NFTGalaxy/zk-ticketing-server/repos/email_credentials"
	"github.com/NFTGalaxy/zk-ticketing-server/repos/events"
	"github.com/NFTGalaxy/zk-ticketing-server/repos/registrations"
	"github.com/NFTGalaxy/zk-ticketing-server/repos/ticket_credentials"
	"github.com/NFTGalaxy/zk-ticketing-server/repos/users"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Client struct {
	DBConnPool        *pgxpool.Pool
	EmailCredentials  *email_credentials.Queries
	Events            *events.Queries
	Registrations     *registrations.Queries
	TicketCredentials *ticket_credentials.Queries
	Users             *users.Queries
}

func NewClient(pool *pgxpool.Pool) *Client {
	return &Client{
		DBConnPool:        pool,
		EmailCredentials:  email_credentials.New(pool),
		Events:            events.New(pool),
		Registrations:     registrations.New(pool),
		TicketCredentials: ticket_credentials.New(pool),
		Users:             users.New(pool),
	}
}
