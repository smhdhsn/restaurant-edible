package service

// MenuService contains repositories that will be used within this service.
type MenuService struct{}

// NewMenuService creates a menu service with it's dependencies.
func NewMenuService() *MenuService {
	return &MenuService{}
}
