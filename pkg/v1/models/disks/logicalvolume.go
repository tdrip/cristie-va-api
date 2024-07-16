package disks

type Volume struct {
	BootVolume   bool   `json:"bootVolume,omitempty"`
	FullName     string `json:"fullName,omitempty"`
	Name         string `json:"name,omitempty"`
	ParentVG     string `json:"parentVG,omitempty"`
	Selected     bool   `json:"selected,omitempty"`
	Size         int64  `json:"size,omitempty"`
	SizeGB       int32  `json:"sizeGB,omitempty"`
	SizeKB       int64  `json:"sizeKB,omitempty"`
	SizeMB       int32  `json:"sizeMB,omitempty"`
	SystemVolume bool   `json:"systemVolume,omitempty"`
	Valid        bool   `json:"valid,omitempty"`
}

type LogicalVolume struct {
	Volume
	Amended     bool     `json:"amended,omitempty"`
	Id          int32    `json:"id,omitempty"`
	MountPoints []string `json:"mountPoints,omitempty"`
}
