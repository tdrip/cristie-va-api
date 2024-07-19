package ldap

type AuthSource struct {
	// The id of the LDAP Auth
	Id int32 `json:"id"`
	// The name of the domain
	DomainName string `json:"domain_name"`
	// If the LDAPAuth is using anonymous access
	AnonymousAccess bool `json:"anonymous_access"`
	// The username to access the LDAP
	Username string `json:"username,omitempty"`
	// The password to access the LDAP
	Password string `json:"password,omitempty"`
	// INTERNAL
	PasswordVersion int32 `json:"password_version,omitempty"`
	// The Primary address of the LDAP/Active Directory Server. In a large environment this can be ommitted and the servers are looked up via DNS
	PrimaryAddr string `json:"primary_addr,omitempty"`
	// The Secondary address of the LDAP/Active Directory server. In a large environment this can be ommitted and the servers are looked up via DNS
	SecondaryAddr string `json:"secondary_addr,omitempty"`
	// The type of server you are connecting to AD/LDAP
	Type string `json:"type"`
	// The User Base distinguished name
	UserBaseDn string `json:"user_base_dn,omitempty"`
	// The Group Base distinguished name
	GroupBaseDn string `json:"group_base_dn,omitempty"`
	// The Search Base distinguished name
	SearchBaseDn string `json:"search_base_dn,omitempty"`
	// The Tenant id
	Tenantid int32 `json:"tenantid"`
	// Strictly enforce SSL cerficiate integrity
	ForceVerification bool `json:"force_verification"`
	// Use secure LDAP protocol
	SecureLdap bool `json:"secure_ldap"`
	// The filter to apply to user searches.                     The standard search uses the following filter: (&(objectClass=user)(objectClass=person)(&(!(objectClass=computer))))                     If provided, the search will use this filter: (&(objectClass=user)(YOUR_FILTER)(objectClass=person)(&(!(objectClass=computer))))
	UserSearchFilter string `json:"user_search_filter,omitempty"`
	// The filter to apply to group searches                     The standard search uses the following filter: (objectClass=group)                     If provided, the search will use this filter: (&(objectClass=group)(YOUR_FILTER))
	GroupSearchFilter string `json:"group_search_filter,omitempty"`
}
