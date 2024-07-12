package disks

type VirtualDisk struct {
	DeviceId      int32          `json:"deviceId,omitempty"`
	Id            int32          `json:"id,omitempty"`
	Name          string         `json:"name,omitempty"`
	PhysicalDisks []PhysicalDisk `json:"physicalDisks,omitempty"`
	Selected      bool           `json:"selected,omitempty"`
	Size          int64          `json:"size,omitempty"`
	SizeGB        int32          `json:"sizeGB,omitempty"`
	SizeKB        int64          `json:"sizeKB,omitempty"`
	SizeMB        int32          `json:"sizeMB,omitempty"`
	Valid         bool           `json:"valid,omitempty"`
}
