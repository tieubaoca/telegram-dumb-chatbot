package services

var _ SDService = &sdService{}

type SDService interface {
}

type sdService struct{}

func NewSDService() SDService {
	return &sdService{}
}
