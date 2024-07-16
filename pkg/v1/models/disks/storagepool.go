package disks

type StoragePool struct {
	Id            int            `json:"id,omitempty"`
	Name          string         `json:"name,omitempty"`
	PhysicalDisks []PhysicalDisk `json:"physicalDisks,omitempty"`
	VirtualDisks  []VirtualDisk  `json:"virtualDisks,omitempty"`
}
