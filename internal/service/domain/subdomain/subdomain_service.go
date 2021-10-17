package subdomain

import "github.com/ozonmp/omp-bot/internal/model/domain"

type SubdomainService interface {
	Describe(subdomainID uint64) (*domain.Subdomain, error)
	List(cursor uint64, limit uint64) ([]domain.Subdomain, error)
	Create(object domain.Subdomain) (uint64, error)
	Update(subdomainID uint64, subdomain domain.Subdomain) error
	Remove(subdomainID uint64) (bool, error)
}

type service struct{}

func NewService() SubdomainService {
	return &service{}
}

func (s *service) Describe(
	subdomainID uint64,
) (
	*domain.Subdomain,
	error,
) {

	return nil, nil
}

func (s *service) List(
	cursor uint64,
	limit uint64,
) (
	[]domain.Subdomain,
	error,
) {

	return nil, nil
}

func (s *service) Create(
	object domain.Subdomain,
) (
	uint64,
	error,
) {

	return 0, nil
}

func (s *service) Update(
	subdomainID uint64,
	subdomain domain.Subdomain,
) error {

	return nil
}

func (s *service) Remove(
	subdomainID uint64,
) (
	bool,
	error,
) {

	return false, nil
}
