package persistency

import (
	"github.com/rigolo-api/services/clients/domain"
	"gorm.io/gorm"
)

type PostgresClientsRepository struct {
	db *gorm.DB
}

func NewPostgresClientsRepository(db *gorm.DB) *PostgresClientsRepository {
	return &PostgresClientsRepository{db: db}

}

func (p PostgresClientsRepository) FindClientByPhone(phone string) (*domain.Client, *domain.ClientError) {
	var entity domain.Client
	result := p.db.Where("phone = ?", phone).Find(&entity)
	if result.Error != nil {
		return nil, domain.NewPersistenceLayerError(result.Error)
	}
	var empty domain.Client
	if entity == empty {
		return nil, nil
	}

	return &entity, nil
}

func (p PostgresClientsRepository) SaveClient(client domain.Client) (*domain.Client, *domain.ClientError) {

	err := p.db.Create(client).Error
	if err != nil {
		return nil, domain.NewPersistenceLayerError(err)
	}

	return &client, nil
}
