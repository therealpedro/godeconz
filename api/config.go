package api

type Config struct {
	ApiVersion          string                  `json:"apiversion"`
	Backup              Config_Backup           `json:"backup"`
	BridgeId            string                  `json:"bridgeid"`
	DataStoreVersion    string                  `json:"datastoreversion"`
	DeviceName          string                  `json:"devicename"`
	Dhcp                bool                    `json:"dhcp"`
	FactoryNew          bool                    `json:"factorynew"`
	FwVersion           string                  `json:"fwversion"`
	Gateway             string                  `json:"gateway"`
	InternetServices    Config_InternetServices `json:"internetservices"`
	IpAddress           string                  `json:"ipaddress"`
	LinkButton          bool                    `json:"linkbutton"`
	Mac                 string                  `json:"mac"`
	ModeId              string                  `json:"modeid"`
	Name                string                  `json:"name"`
	Netmask             string                  `json:"netmask"`
	Networkopenduration int                     `json:"networkopenduration"`
	Ntp                 string                  `json:"ntp"`
	PanId               int                     `json:"panid"`
	PortalConnection    string                  `json:"portalconnection"`
	PortalServices      bool                    `json:"portalservices"`
	PortalState         Config_PortalState      `json:"portalstate"`
	ProxyAddress        string                  `json:"proxyaddress"`
	ProxyPort           int                     `json:"proxyport"`
	RfConnected         bool                    `json:"rfconncected"`
	StartkitId          string                  `json:"startkitid"`
	SwUpdate            interface{}             `json:"swupdate"` //TODO Implement later
	SwVersion           string                  `json:"swversion"`
	TimeFormat          string                  `json:"timeformat"`
	TimeZone            string                  `json:"timezone"`
	UTC                 string                  `json:"UTC"`
	Uuid                string                  `json:"uuid"`
	WebsocketNotifyAll  bool                    `json:"websocketnotifyall"`
	WebsocketPort       int                     `json:"websocketport"`
	Whitelist           Config_Whitelist        `json:"whitelist"`
	ZigbeeChannel       int                     `json:"zigbeechannel"`
}

type Config_Backup struct {
	ErrorCode int    `json:"errorcode"`
	Status    string `json:"status"`
}

type Config_InternetServices struct {
	RemoteAccess string `json:"remoteaccess"`
}

type Config_PortalState struct {
	Communication string `json:"communication"`
	Incoming      bool   `json:"incoming"`
	Outgoing      bool   `json:"outgoing"`
	Signedon      bool   `json:"signedon"`
}

type Config_Whitelist struct {
	Client map[string]Config_DeconzClient
}

type Config_DeconzClient struct {
	CreateDate  string `json:"create date"`
	LastUseDate string `json:"last use date"`
	Name        string `json:"name"`
}
