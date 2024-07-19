package estate

import (
	"time"
)

type System struct {
	Addresses []string `json:"addresses,omitempty"`
	//BackupConfig        *BackupJobConfig  `json:"backupConfig,omitempty"`
	//BackupStatus        *BackupJobStatus  `json:"backupStatus,omitempty"`
	//BackupStorage       *BackupStorage    `json:"backupStorage,omitempty"`
	BiosUuid         string     `json:"biosUuid,omitempty"`
	BootTime         *time.Time `json:"bootTime,omitempty"`
	CbtDriverVersion string     `json:"cbtDriverVersion,omitempty"`
	CbtEnabled       bool       `json:"cbtEnabled,omitempty"`
	CommAddr         string     `json:"commAddr,omitempty"`
	CpuArch          string     `json:"cpuArch,omitempty"`
	CredentialId     int        `json:"credentialId,omitempty"`
	CredentialType   string     `json:"credentialType,omitempty"`
	//Credentials         *Credentials      `json:"credentials,omitempty"`
	DiscoveryType  string     `json:"discoveryType,omitempty"`
	DisrecDate     *time.Time `json:"disrecDate,omitempty"`
	Edition        string     `json:"edition,omitempty"`
	ForceNextFull  bool       `json:"forceNextFull,omitempty"`
	Hostname       string     `json:"hostname,omitempty"`
	HypervisorType string     `json:"hypervisorType,omitempty"`
	IsoCreated     *time.Time `json:"isoCreated,omitempty"`
	IsoCreatedOn   *time.Time `json:"isoCreatedOn,omitempty"`
	//LastBackupStatus    *BackupJobStatus  `json:"lastBackupStatus,omitempty"`
	LastUpdate          *time.Time      `json:"lastUpdate,omitempty"`
	LicenceSurrendered  bool            `json:"licenceSurrendered,omitempty"`
	MacAddress          string          `json:"macAddress,omitempty"`
	OperatingSystem     string          `json:"operatingSystem,omitempty"`
	Password            string          `json:"password,omitempty"`
	PasswordStored      bool            `json:"passwordStored,omitempty"`
	PasswordVersion     string          `json:"passwordVersion,omitempty"`
	Platform            string          `json:"platform,omitempty"`
	Port                int             `json:"port,omitempty"`
	ProductInfo         *ProductInfo    `json:"productInfo,omitempty"`
	RansomwareScan      *RansomwareScan `json:"ransomwareScan,omitempty"`
	RdpEnabled          bool            `json:"rdpEnabled,omitempty"`
	Recovery            bool            `json:"recovery,omitempty"`
	RecoveryEnvironment bool            `json:"recoveryEnvironment,omitempty"`
	//RestoreStatus       *RestoreJobStatus `json:"restoreStatus,omitempty"`
	SshKey          string     `json:"sshKey,omitempty"`
	SshKeyStored    bool       `json:"sshKeyStored,omitempty"`
	SshPort         int        `json:"sshPort,omitempty"`
	TestToolLastRun *time.Time `json:"testToolLastRun,omitempty"`
	TestToolLevel   int        `json:"testToolLevel,omitempty"`
	Username        string     `json:"username,omitempty"`
	Uuid            string     `json:"uuid,omitempty"`
}
