package disks

type StandardDisk struct {
	BootDisk        bool        `json:"bootDisk,omitempty"`
	BytesPerSector  int         `json:"bytesPerSector,omitempty"`
	Datastore       string      `json:"datastore,omitempty"`
	DiskNo          int         `json:"diskNo,omitempty"`
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
	SizeGB          int         `json:"sizeGB,omitempty"`
	SizeKB          int64       `json:"sizeKB,omitempty"`
	SizeMB          int         `json:"sizeMB,omitempty"`
	SystemDisk      bool        `json:"systemDisk,omitempty"`
	Thin            bool        `json:"thin,omitempty"`
	Type            string      `json:"type"`
	Valid           bool        `json:"valid,omitempty"`
}
