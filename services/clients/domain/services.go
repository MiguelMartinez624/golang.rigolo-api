package domain

import (
	"github.com/rigolo-api/common/value_objects"
	"log"
)

type ClientsService interface {
	AddNewClient(newClient NewClient) (*Client, *ClientError)

	GetAllClients() (clientList []Client, clientError *ClientError)
}

type Service struct {
	clientsRepo ClientsRepository
}

func NewService(clientsRepo ClientsRepository) *Service {
	return &Service{clientsRepo: clientsRepo}
}

func (s *Service) AddNewClient(newClient NewClient) (*Client, *ClientError) {
	err := newClient.IsValid()
	if err != nil {
		return nil, err
	}

	if clientStored, err := s.clientsRepo.FindClientByPhone(string(newClient.Phone)); err != nil {
		return nil, err
	} else if clientStored != nil {
		log.Println(clientStored)
		return nil, DuplicatedPhoneError
	}

	clientToStore := Client{
		ID:     value_objects.NewID(),
		Phone:  newClient.Phone,
		Email:  newClient.Email,
		Name:   newClient.Name,
		Region: newClient.Region,
	}

	return s.clientsRepo.SaveClient(clientToStore)
}

func (s *Service) GetAllClients() (clientList []Client, clientError *ClientError) {
	return s.clientsRepo.GetAllClients()
}
