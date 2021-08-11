package api

type Event struct {
	Type     string                 `json:"t"`
	Event    string                 `json:"e"`
	Resource string                 `json:"r"`
	Id       string                 `json:"id"`
	UniqueId string                 `json:"uniqueid"`
	GroupId  string                 `json:"gid"`    // only for scene called events
	SceneId  string                 `json:"scid"`   // only for scene called events
	Config   map[string]interface{} `json:"config"` // only for changed events
	Name     string                 `json:"name"`   // only for changed events
	State    map[string]interface{} `json:"state"`  // only for changed events
	Group    map[string]interface{} `json:"group"`  // only for added events of a new group source
	Light    map[string]interface{} `json:"light"`  // only for added of a new light source
	Sensor   map[string]interface{} `json:"sensor"` // only for added of a new sensor source
}
