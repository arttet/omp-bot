package basket

import "github.com/ozonmp/omp-bot/internal/model/buy"

type BasketService interface {
	Describe(basketID uint64) (*buy.Basket, error)
	List(cursor uint64, limit uint64) ([]buy.Basket, error)
	Create(basket buy.Basket) (uint64, error)
	Update(basketID uint64, basket buy.Basket) error
	Remove(basketID uint64) (bool, error)
}

type service struct{}

func NewService() BasketService {
	return &service{}
}

func (s *service) Describe(subdomainID uint64) (*buy.Basket, error) {
	return nil, nil
}

func (s *service) List(cursor uint64, limit uint64) ([]buy.Basket, error) {
	return nil, nil
}

func (s *service) Create(object buy.Basket) (uint64, error) {
	return 0, nil
}

func (s *service) Update(subdomainID uint64, subdomain buy.Basket) error {
	return nil
}

func (s *service) Remove(subdomainID uint64) (bool, error) {
	return false, nil
}
