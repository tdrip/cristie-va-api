package targets

type CmLiteTarget struct {
	Disk         string `json:"disk,omitempty"`
	HypervisorId int    `json:"hypervisorId,omitempty"`
	Id           int    `json:"id,omitempty"`
	StorageId    int    `json:"storageId,omitempty"`
	TargetType   string `json:"targetType,omitempty"`
}