package mongo

type ServiceI interface {
}

type Service struct {
}

func InitializeMongoService() *Service {
	return &Service{}
}
