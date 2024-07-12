package backupservers

type ConfigRequest struct {
	BackupType          string   `json:"backupType,omitempty"`
	ClientDomain        string   `json:"clientDomain,omitempty"`
	ClientName          string   `json:"clientName,omitempty"`
	EncryptionPasswords []string `json:"encryptionPasswords,omitempty"`
	Password            string   `json:"password,omitempty"`
	ServerId            int      `json:"serverId,omitempty"`
	Username            string   `json:"username,omitempty"`
}
