//  (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package models

type IDNameModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ListNetworksBody struct {
	Networks     []GetSpecificNetwork `json:"networks"`
	NetworkCount int                  `json:"networkCount"`
}

type CreateNetworkResponse struct {
	Success bool               `json:"success"`
	Errors  interface{}        `json:"errors"`
	Network GetSpecificNetwork `json:"network"`
}

type GetSpecificNetworkBody struct {
	Network GetSpecificNetwork `json:"network"`
}

type GetSpecificNetwork struct {
	ID                      int                             `json:"id" tf:"id,computed"`
	Name                    string                          `json:"name" tf:"name"`
	Zone                    IDNameModel                     `json:"zone"`
	DisplayName             string                          `json:"displayName" tf:"display_name,computed"`
	Type                    IDModel                         `json:"type"`
	TypeID                  int                             `json:"-" tf:"type_id,computed"`
	Owner                   IDNameModel                     `json:"owner"`
	Code                    string                          `json:"code" tf:"code,computed"`
	Category                string                          `json:"category"`
	ExternalID              string                          `json:"externalId" tf:"external_id,computed"`
	InternalID              string                          `json:"internalId" tf:"internal_id,computed"`
	UniqueID                string                          `json:"uniqueId" tf:"unique_id,computed"`
	ExternalType            string                          `json:"externalType"`
	RefType                 string                          `json:"refType"`
	RefID                   int                             `json:"refId"`
	DhcpServer              bool                            `json:"dhcpServer" tf:"dhcp_server"`
	Status                  string                          `json:"status" tf:"status,computed"`
	Visibility              string                          `json:"visibility" tf:"visibility"`
	EnableAdmin             bool                            `json:"enableAdmin"`
	ScanNetwork             bool                            `json:"scanNetwork" tf:"scan_network"`
	Active                  bool                            `json:"active" tf:"active"`
	DefaultNetwork          bool                            `json:"defaultNetwork"`
	AssignPublicIP          bool                            `json:"assignPublicIp"`
	ApplianceURLProxyBypass bool                            `json:"applianceUrlProxyBypass" tf:"appliance_url_proxy_bypass"`
	ZonePool                IDNameModel                     `json:"zonePool"`
	AllowStaticOverride     bool                            `json:"allowStaticOverride"`
	Tenants                 []IDNameModel                   `json:"tenants"`
	ResourcePermissions     CreateNetworkResourcePermission `json:"resourcePermissions" tf:"resource_permissions,sub"`
}

type CreateNetworkResourcePermission struct {
	All   bool      `json:"all" tf:"all"`
	Sites []IDModel `json:"sites" tf:"sites"`
}

type CreateNetworkRequest struct {
	Network             CreateNetwork                   `json:"network"`
	ResourcePermissions CreateNetworkResourcePermission `json:"resourcePermissions"`
}

type CreateNetwork struct {
	ID                      int                             `json:"-" tf:"id,computed"`
	Name                    string                          `json:"name" tf:"name"`
	Description             string                          `json:"description" tf:"description"`
	CloudID                 int                             `json:"-" tf:"cloud_id"`
	GroupID                 int                             `json:"-" tf:"group_id"`
	TypeID                  int                             `json:"-" tf:"type_id"`
	PoolID                  int                             `json:"-" tf:"pool_id"`
	Zone                    IDModel                         `json:"zone,omitempty"`
	Site                    IDModel                         `json:"site,omitempty"`
	Type                    IDModel                         `json:"type,omitempty"`
	Pool                    *IDModel                        `json:"pool,omitempty"`
	Cidr                    string                          `json:"cidr,omitempty" tf:"cidr"`
	Gateway                 string                          `json:"gateway,omitempty" tf:"gateway"`
	DNSPrimary              string                          `json:"dnsPrimary,omitempty" tf:"primary_dns"`
	DNSSecondary            string                          `json:"dnsSecondary,omitempty" tf:"secondary_dns"`
	Config                  CreateNetworkConfig             `json:"config,omitempty" tf:"config,sub"`
	DhcpServer              bool                            `json:"dhcpServer" tf:"dhcp_server"`
	ScanNetwork             string                          `json:"scanNetwork,omitempty" tf:"scan_network"`
	ApplianceURLProxyBypass string                          `json:"applianceUrlProxyBypass,omitempty" tf:"appliance_url_proxy_bypass"`
	NoProxy                 string                          `json:"noProxy,omitempty" tf:"no_proxy"`
	ScopeID                 string                          `json:"scopeId,omitempty" tf:"scode_id"`
	ResourcePermissions     CreateNetworkResourcePermission `json:"-" tf:"resource_permissions,sub"`
}

type CreateNetworkConfig struct {
	ConnectedGateway string `json:"connectedGateway" tf:"connected_gateway"`
	VlanIDs          string `json:"vlanIDs" tf:"vlan_ids"`
}

type GetNetworkTypesResponse struct {
	NetworkTypes []GetSpecificNetworkType `json:"networkTypes"`
}

type GetSpecificNetworkType struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

type GetNetworkPoolsResp struct {
	NetworkPools     []GetNetworkPool `json:"networkPools"`
	NetworkPoolCount int              `json:"networkPoolCount"`
	Meta             Meta             `json:"meta"`
}

type GetSpecificNetworkPool struct {
	NetworkPool GetNetworkPool `json:"networkPool"`
}

type GetNetworkPool struct {
	ID          int                      `json:"id"`
	Type        IDModel                  `json:"type"`
	Account     IDModel                  `json:"account"`
	Category    string                   `json:"category"`
	Code        string                   `json:"code"`
	Name        string                   `json:"name" tf:"name"`
	DisplayName string                   `json:"displayName" tf:"display_name,computed"`
	InternalID  interface{}              `json:"internalId"`
	ExternalID  string                   `json:"externalId"`
	PoolGroup   interface{}              `json:"poolGroup"`
	IPRanges    []GetNetworkPoolIPRanges `json:"ipRanges"`
}

type GetNetworkPoolIPRanges struct {
	ID           int    `json:"id"`
	StartAddress string `json:"startAddress"`
	EndAddress   string `json:"endAddress"`
}
