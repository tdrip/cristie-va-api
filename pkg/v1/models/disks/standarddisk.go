package disks

type StandardDisk struct {
	BootDisk        bool        `json:"bootDisk,omitempty"`
	BytesPerSector  int32       `json:"bytesPerSector,omitempty"`
	Datastore       string      `json:"datastore,omitempty"`
	DiskNo          int32       `json:"diskNo,omitempty"`
	GptDisk         bool        `json:"gptDisk,omitempty"`
	Guid            string      `json:"guid,omitempty"`
	ImageBootDisk   bool        `json:"imageBootDisk,omitempty"`
	Lazy            bool        `json:"lazy,omitempty"`
	LocalDisk       bool        `json:"localDisk,omitempty"`
	MbrDisk         bool        `json:"mbrDisk,omitempty"`
	Name            string      `json:"name,omitempty"`
	Partitions      []Partition `json:"partitions,omitempty"`
	Selected        bool        `json:"selected,omitempty"`
	SignatureString string      `json:"signatureString,omitempty"`
	Size            int64       `json:"size,omitempty"`
	SizeGB          int32       `json:"sizeGB,omitempty"`
	SizeKB          int64       `json:"sizeKB,omitempty"`
	SizeMB          int32       `json:"sizeMB,omitempty"`
	SystemDisk      bool        `json:"systemDisk,omitempty"`
	Thin            bool        `json:"thin,omitempty"`
	Type_           string      `json:"type"`
	Valid           bool        `json:"valid,omitempty"`
}
