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
	BootDelay                 int                             `json:"bootDelay,omitempty"`
	BootOrder                 int                             `json:"bootOrder,omitempty"`
	CmLiteTarget              *CmLiteTarget                   `json:"cmLiteTarget,omitempty"`
	DhwConfig                 *DhwConfiguration               `json:"dhwConfig,omitempty"`
	Dissimilar                bool                            `json:"dissimilar,omitempty"`
	DnsEntryRequest           *networks.DnsEntryRequest       `json:"dnsEntryRequest,omitempty"`
	DrNetworkBondingConfigs   []networks.DrNetworkBonding     `json:"drNetworkBondingConfigs,omitempty"`
	DrNetworkConfig           []networks.NetworkInformation   `json:"drNetworkConfig,omitempty"`
	Enabled                   bool                            `json:"enabled,omitempty"`
	EnhancedTestingConfig     *etesting.EnhancedTestingConfig `json:"enhancedTestingConfig,omitempty"`
	Events                    []models.Event                  `json:"events,omitempty"`
	Id                        int                             `json:"id,omitempty"`
	IsoTrigger                bool                            `json:"isoTrigger,omitempty"`
	LastRunTime               time.Time                       `json:"lastRunTime,omitempty"`
	OptType                   string                          `json:"optType,omitempty"`
	OrchestrationBlockId      int                             `json:"orchestrationBlockId,omitempty"`
	PauseForValidation        bool                            `json:"pauseForValidation,omitempty"`
	PostRecoveryNetworkConfig []networks.NetworkInformation   `json:"postRecoveryNetworkConfig,omitempty"`
	ProductType               string                          `json:"productType,omitempty"`
	RebootBlockId             int                             `json:"rebootBlockId,omitempty"`
	RecoverToSource           bool                            `json:"recoverToSource,omitempty"`
	RecoveryStatus            *recovery.Status                `json:"recoveryStatus,omitempty"`
	ReportBlockId             int                             `json:"reportBlockId,omitempty"`
	ScheduleId                int                             `json:"scheduleId,omitempty"`
	ScriptBlockId             int                             `json:"scriptBlockId,omitempty"`
	SendStatistics            bool                            `json:"sendStatistics,omitempty"`
	Source                    *BackupSource                   `json:"source,omitempty"`
	SourceClient              string                          `json:"sourceClient,omitempty"`
	SourceMacAddress          string                          `json:"sourceMacAddress,omitempty"` //missing from swagger
	SourceServer              string                          `json:"sourceServer,omitempty"`
	Status                    string                          `json:"status,omitempty"`
	StoragePools              []disks.StoragePool             `json:"storagePools,omitempty"`
	Target                    *HypervisorTarget               `json:"target,omitempty"`
	TargetBiosUuid            string                          `json:"targetBiosUuid,omitempty"`
	TargetConfig              *TargetConfig                   `json:"targetConfig,omitempty"`
	TargetIP                  string                          `json:"targetIP,omitempty"`
	TargetMAC                 string                          `json:"targetMAC,omitempty"`
	TargetVmCreated           bool                            `json:"targetVmCreated,omitempty"`
	Type                      string                          `json:"type"`.                     //missing from swagger
	ValidationPauseDuration   int                             `json:"validationPauseDuration,omitempty"`
	VmHostType                string                          `json:"vmHostType,omitempty"`
	VpcConfig                 *VpcConfig                      `json:"vpcConfig,omitempty"`
	VpcEnabled                bool                            `json:"vpcEnabled,omitempty"`
	WebbootMapping            *webboot.Mapping                `json:"webbootMapping,omitempty"`
}
