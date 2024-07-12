package disks

type LogicalVolume struct {
	Amended      bool     `json:"amended,omitempty"`
	BootVolume   bool     `json:"bootVolume,omitempty"`
	FullName     string   `json:"fullName,omitempty"`
	Id           int32    `json:"id,omitempty"`
	MountPoints  []string `json:"mountPoints,omitempty"`
	Name         string   `json:"name,omitempty"`
	ParentVG     string   `json:"parentVG,omitempty"`
	Selected     bool     `json:"selected,omitempty"`
	Size         int64    `json:"size,omitempty"`
	SizeGB       int32    `json:"sizeGB,omitempty"`
	SizeKB       int64    `json:"sizeKB,omitempty"`
	SizeMB       int32    `json:"sizeMB,omitempty"`
	SystemVolume bool     `json:"systemVolume,omitempty"`
	Uuid         string   `json:"uuid,omitempty"`
	Valid        bool     `json:"valid,omitempty"`
}