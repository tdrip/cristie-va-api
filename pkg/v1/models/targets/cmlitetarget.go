package targets

type CmLiteTarget struct {
	Disk         string `json:"disk,omitempty"`
	HypervisorId int32  `json:"hypervisorId,omitempty"`
	Id           int32  `json:"id,omitempty"`
	StorageId    int32  `json:"storageId,omitempty"`
	TargetType   string `json:"targetType,omitempty"`
}