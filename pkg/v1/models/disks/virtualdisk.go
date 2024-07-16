package disks

type VirtualDisk struct {
	PhysicalDisk
	PhysicalDisks []PhysicalDisk `json:"physicalDisks,omitempty"`
}
