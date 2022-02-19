package jwt

type ServiceI interface {
}

type Service struct {
}

func InitializeJwtService() *Service {
	return &Service{}
}
