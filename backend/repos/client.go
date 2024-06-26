package repos

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zk-ticketing/zk-ticketing/backend/repos/email_credentials"
	"github.com/zk-ticketing/zk-ticketing/backend/repos/events"
	"github.com/zk-ticketing/zk-ticketing/backend/repos/registrations"
	"github.com/zk-ticketing/zk-ticketing/backend/repos/ticket_credentials"
	"github.com/zk-ticketing/zk-ticketing/backend/repos/users"
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
