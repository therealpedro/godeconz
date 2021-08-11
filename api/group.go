package api

type Group struct {
	Action           Action        `json:"action"`
	DeviceMembership []interface{} `json:"devicemembership"`
	Etag             string        `json:"etag"`
	Hidden           bool          `json:"hidden"`
	Id               string        `json:"id"`
	Lights           []string      `json:"lights"`
	LightSequence    []string      `json:"lightsequence"`
	MultiDeviceIds   []string      `json:"multideviceids"`
	Name             string        `json:"name"`
	State            Group_State   `json:"state"`
	Type             string        `json:"type"`
	//TODO Scenes
}

type Action struct {
	Bri       int         `json:"bri"`
	Colormode string      `json:"colormode"`
	Ct        int         `json:"int"`
	Effect    string      `json:"effect"`
	Hue       int         `json:"hue"`
	On        bool        `json:"on"`
	Sat       int         `json:"sat"`
	Scene     interface{} `json:"scene"`
	Xy        []float32   `json:"xy"`
}

type Group_State struct {
	All_On bool `json:"all_on"`
	Any_On bool `json:"any_on"`
}
