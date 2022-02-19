package aws

type ServiceI interface {
}

type Service struct {
}

func InitializeAwsService() *Service {
	return &Service{}
}
