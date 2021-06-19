package domain

type ClientsRepository interface {
	FindClientByPhone(phone string) (*Client, *ClientError)

	SaveClient(client Client) (*Client, *ClientError)
}
