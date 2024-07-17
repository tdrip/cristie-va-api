package targets

import (
	"time"

	models "github.com/tdrip/cristie-va-api/pkg/v1/models"
	disks "github.com/tdrip/cristie-va-api/pkg/v1/models/disks"
	etesting "github.com/tdrip/cristie-va-api/pkg/v1/models/etesting"
	networks "github.com/tdrip/cristie-va-api/pkg/v1/models/networks"
	recovery "github.com/tdrip/cristie-va-api/pkg/v1/models/recovery"
	webboot "github.com/tdrip/cristie-va-api/pkg/v1/models/webboot"
)

type SourceTarget struct {
	BootDelay                  *int                            `json:"bootDelay"`    // null required
	BootOrder                  *int                            `json:"bootOrder"`    // null required
	CmLiteTarget               *CmLiteTarget                   `json:"cmLiteTarget"` // null required
	DhwConfig                  *DhwConfiguration               `json:"dhwConfig"`    // null required
	Dissimilar                 bool                            `json:"dissimilar,omitempty"`
	DnsEntryRequest            *networks.DnsEntryRequest       `json:"dnsEntryRequest"`         // null required
	DrNetworkBondingConfigs    []networks.DrNetworkBonding     `json:"drNetworkBondingConfigs"` // null required
	DrNetworkConfigs           []networks.NetworkInformation   `json:"drNetworkConfig"`         // null required
	Enabled                    bool                            `json:"enabled,omitempty"`
	EnhancedTestingConfig      *etesting.EnhancedTestingConfig `json:"enhancedTestingConfig"` // null required
	Events                     []models.Event                  `json:"events,omitempty"`
	Id                         *int                            `json:"id"` // null required
	IsoTrigger                 bool                            `json:"isoTrigger,omitempty"`
	LastRunTime                *time.Time                      `json:"lastRunTime"` // null required
	OptType                    string                          `json:"optType,omitempty"`
	OrchestrationBlockId       int                             `json:"orchestrationBlockId,omitempty"`
	PauseForValidation         bool                            `json:"pauseForValidation,omitempty"`
	PostRecoveryNetworkConfigs []networks.NetworkInformation   `json:"postRecoveryNetworkConfig,omitempty"`
	ProductType                string                          `json:"productType,omitempty"`
	RebootBlockId              int                             `json:"rebootBlockId,omitempty"`
	RecoverToSource            bool                            `json:"recoverToSource,omitempty"`
	RecoveryStatus             *recovery.Status                `json:"recoveryStatus,omitempty"`
	ReportBlockId              int                             `json:"reportBlockId,omitempty"`
	ScheduleId                 int                             `json:"scheduleId,omitempty"`
	ScriptBlockId              int                             `json:"scriptBlockId,omitempty"`
	SendStatistics             bool                            `json:"sendStatistics,omitempty"`
	SendStats                  bool                            `json:"sendStats,omitempty"`
	Source                     *BackupSource                   `json:"source,omitempty"`
	SourceClient               string                          `json:"sourceClient,omitempty"`
	SourceMacAddress           string                          `json:"sourceMacAddress,omitempty"` //missing from swagger
	SourceServer               string                          `json:"sourceServer,omitempty"`
	Status                     string                          `json:"status,omitempty"`
	StoragePools               []disks.StoragePool             `json:"storagePools"`   // null required
	Target                     *HypervisorTarget               `json:"target"`         // null required
	TargetBiosUuid             *string                         `json:"targetBiosUuid"` // null required
	TargetConfig               *TargetConfig                   `json:"targetConfig,omitempty"`
	TargetIP                   string                          `json:"targetIP,omitempty"`
	TargetMAC                  *string                         `json:"targetMAC"` // null required
	TargetVmCreated            bool                            `json:"targetVmCreated,omitempty"`
	Type                       string                          `json:"type"` //missing from swagger
	ValidationPauseDuration    int                             `json:"validationPauseDuration,omitempty"`
	VmHostType                 string                          `json:"vmHostType,omitempty"`
	VpcConfig                  *VpcConfig                      `json:"vpcConfig"` // null required
	VpcEnabled                 bool                            `json:"vpcEnabled,omitempty"`
	WebbootMapping             *webboot.Mapping                `json:"webbootMapping"` // null required
}

func NewSourceTarget() SourceTarget {
	st := SourceTarget{}
	st.BootDelay = nil
	st.BootOrder = nil
	st.CmLiteTarget = nil
	st.DhwConfig = &DhwConfiguration{}
	st.DnsEntryRequest = nil
	st.DrNetworkBondingConfigs = []networks.DrNetworkBonding{}
	st.DrNetworkConfigs = []networks.NetworkInformation{}
	st.EnhancedTestingConfig = nil
	st.Id = nil
	st.PostRecoveryNetworkConfigs = []networks.NetworkInformation{}
	st.SendStats = true
	st.StoragePools = nil
	st.Target = nil
	st.VpcConfig = nil
	return st
}
