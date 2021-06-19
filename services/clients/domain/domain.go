package domain

type Domain struct {
	clientsService ClientsService
}

func NewDomain(clientsRepo ClientsRepository) *Domain {
	return &Domain{clientsService: NewService(clientsRepo)}
}

func (d *Domain) RegisterNewClient(newClient NewClient) (*Client, *ClientError) {
	return d.clientsService.AddNewClient(newClient)
}
