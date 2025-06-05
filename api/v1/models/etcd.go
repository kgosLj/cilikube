package models

// EtcdOptions 定义了 etcd 相关的选项（该配置项只适合运行在 Pod 中的 Etcd）
type EtcdOptions struct {
	ClusterName      string   `json:"clusterName" binding:"required"`
	EtcdPodNamespace string   `json:"etcdNamespace,omitempty"`                               // etcd 所在的命名空间
	EtcdPodLabel     string   `json:"etcdLabel,omitempty"`                                   // etcd 标签
	EtcdPodContainer string   `json:"etcdContainer,omitempty"`                               // etcd 容器名称
	EndPoints        []string `json:"endpoints,omitempty"`                                   // etcd 端点
	EtcdCa           string   `json:"etcdCa,omitempty"`                                      // etcd CA证书
	EtcdCert         string   `json:"etcdCert,omitempty"`                                    // etcd 客户端证书
	EtcdKey          string   `json:"etcdKey,omitempty"`                                     // etcd 客户端秘钥
	EtcdBackPath     string   `json:"etcdBackPath,omitempty"`                                // etcd 备份路径
	EtcdBackName     string   `json:"etcdBackName,omitempty"`                                // etcd 备份名称
	EtcdBackType     string   `json:"etcdBackType,omitempty" example:"local-storage, s3..."` // etcd 备份类型
	EtcdBackSize     float64  `json:"etcdBackSize,omitempty"`                                // etcd 备份大小
}

// EtcdBackRequest 创建 etcd 备份请求
type EtcdBackRequest struct {
	ClusterName  string `json:"clusterName" binding:"required"`
	EtcdBackName string `json:"etcdBackName,omitempty"`
	Description  string `json:"description,omitempty"`
}

// EtcdBackResponse 创建 etcd 备份响应
type EtcdBackResponse struct {
	ClusterName  string `json:"clusterName" binding:"required"`
	EtcdBackName string `json:"etcdBackName" binding:"required"`
	Description  string `json:"description,omitempty"`
	Status       string `json:"status" binding:"required"`
	StartTime    string `json:"startTime" binding:"required"`
	EndTime      string `json:"endTime,omitempty" binding:"required"`
	Size         string `json:"size,omitempty" binding:"required"`
	Error        string `json:"error,omitempty"`
}
