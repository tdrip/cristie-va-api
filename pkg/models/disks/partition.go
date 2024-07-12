package disks

type Partition struct {
	BootPartition   bool     `json:"bootPartition,omitempty"`
	Efi             bool     `json:"efi,omitempty"`
	Filesystem      string   `json:"filesystem,omitempty"`
	Guid            string   `json:"guid,omitempty"`
	Id              int32    `json:"id,omitempty"`
	Identifier      string   `json:"identifier,omitempty"`
	MountCount      int32    `json:"mountCount,omitempty"`
	MountPoints     []string `json:"mountPoints,omitempty"`
	MountedFolder   bool     `json:"mountedFolder,omitempty"`
	Msr             bool     `json:"msr,omitempty"`
	Name            string   `json:"name,omitempty"`
	PartitionNumber int32    `json:"partitionNumber,omitempty"`
	Raid            bool     `json:"raid,omitempty"`
	RaidName        string   `json:"raidName,omitempty"`
	Selected        bool     `json:"selected,omitempty"`
	Size            int64    `json:"size,omitempty"`
	SizeGB          int32    `json:"sizeGB,omitempty"`
	SizeKB          int64    `json:"sizeKB,omitempty"`
	SizeMB          int32    `json:"sizeMB,omitempty"`
	SystemPartition bool     `json:"systemPartition,omitempty"`
	Valid           bool     `json:"valid,omitempty"`
}
