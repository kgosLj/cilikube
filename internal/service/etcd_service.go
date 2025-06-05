package service

// EtcdService 结构体中不需要 client 字段
type EtcdService struct {
}

func NewEtcdService() *EtcdService {
	return &EtcdService{}
}
