package disks

type PhysicalDisk struct {
	DeviceId int    `json:"deviceId,omitempty"`
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Selected bool   `json:"selected,omitempty"`
	Size     int64  `json:"size,omitempty"`
	SizeGB   int    `json:"sizeGB,omitempty"`
	SizeKB   int64  `json:"sizeKB,omitempty"`
	SizeMB   int    `json:"sizeMB,omitempty"`
	Valid    bool   `json:"valid,omitempty"`
}
