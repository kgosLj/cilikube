package handlers

import (
	"github.com/ciliverse/cilikube/internal/service"
	"github.com/ciliverse/cilikube/pkg/k8s"
	"github.com/gin-gonic/gin"
)

// EtcdHandler ...
type EtcdHandler struct {
	service        *service.EtcdService
	clusterManager *k8s.ClusterManager
}

func NewEtcdHandler(svc *service.EtcdService, cm *k8s.ClusterManager) *EtcdHandler {
	return &EtcdHandler{
		service:        svc,
		clusterManager: cm,
	}
}

func (h *EtcdHandler) ListKeys(c *gin.Context) {

}
