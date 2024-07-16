package disks

type DiskType struct {
	BootDisk       bool    `json:"bootDisk,omitempty"`
	Datastore      string  `json:"datastore,omitempty"`
	DiskNo         int     `json:"diskNo,omitempty"`
	DiskPrefix     string  `json:"diskPrefix,omitempty"`
	Id             int     `json:"id,omitempty"`
	Lazy           bool    `json:"lazy,omitempty"`
	LocalDisk      bool    `json:"localDisk,omitempty"`
	Multipathed    bool    `json:"multipathed,omitempty"`
	Overhead       float64 `json:"overhead,omitempty"`
	Recover        bool    `json:"recover,omitempty"`
	SizeInGB       int     `json:"sizeInGB,omitempty"`
	SourceDiskId   int     `json:"sourceDiskId,omitempty"`
	SourceDiskName string  `json:"sourceDiskName,omitempty"`
	SystemDisk     bool    `json:"systemDisk,omitempty"`
	Thin           bool    `json:"thin,omitempty"`
	UserAdded      bool    `json:"userAdded,omitempty"`
}
