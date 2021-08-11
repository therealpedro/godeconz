package api

const GET_DISCOVER_GATEWAY = "http://phoscon.de/discover"

const CONTENTTYPE_JSON = "application/json"

const URL_WEBSOCKET_SERVER = "ws://%s:443"

// Endpoint paths for interaction with lights
const (
	URL_LIGHTS_GET          = "http://%s:%d/api/%s/lights"
	URL_LIGHT_STATE_GET     = "http://%s:%d/api/%s/lights/%d"
	URL_LIGHT_STATE_PUT     = "http://%s:%d/api/%s/lights/%d/state"
	URL_LIGHT_ATTRIBUTE_PUT = "http://%s:%d/api/%s/lights/%d"
	URL_LIGHT_DELETE        = "http://%s:%d/api/%s/lights/%d"
	URL_LIGHT_GROUPS_DELETE = "http://%s:%d/api/%s/lights/%d/groups"
	URL_LIGHT_SCENES_DELETE = "http://%s:%d/api/%s/lights/%d/scenes"
)

// Endpoint paths for interaction with groups
const (
	URL_GROUP_CREATE_POST    = "http://%s:%d/api/%s/groups"
	URL_GROUPS_GET           = "http://%s:%d/api/%s/groups"
	URL_GROUP_ATTRIBUTES_GET = "http://%s:%d/api/%s/groups/%s"
	URL_GROUP_ATTRIBUTES_PUT = "http://%s:%d/api/%s/groups/%s"
	URL_GROUP_STATE_PUT      = "http://%s:%d/api/%s/groups/%s/action"
	URL_GROUP_DELETE         = "http://%s:%d/api/%s/groups/%s"
)

// Endpoint paths for interaction with the gateway configuration
const (
	URL_ACQUIRE_APIKEY_POST  = "http://%s:%d/api"
	URL_APIKEY_DELETE        = "http://%s:%d/api/%s/config/whitelist/%s"
	URL_CURRENT_CONFIG_GET   = "http://%s:%d/api/%s/config"
	URL_FULLSTATE_GET        = "http://%s:%d/api/%s"
	URL_CONFIG_PUT           = "http://%s:%d/api/%s/config"
	URL_UPDATE_SOFTWARE_POST = "http://%s:%d/api/%s/config/update"
	URL_UPDATE_FIRMWARE_POST = "http://%s:%d/api/%s/config/updatefirmware"
	URL_RESET_GATEWAY_POST   = "http://%s:%d/api/%s/config/reset"
	URL_CHANGE_PASSWORD_PUT  = "http://%s:%d/api/%s/config/password"
	URL_RESET_DELETE         = "http://%s:%d/api/%s/config/password"
)
