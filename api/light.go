package api

type Light struct {
	ColorCapabilities int         `json:"colorcapabilities"`
	Ctmax             int         `json:"ctmax"`
	Ctmin             int         `json:"ctmin"`
	LastAnnounced     string      `json:"lastannounced"`
	LastSeen          string      `json:"lastseen"`
	Etag              string      `json:"etag"`
	HasColor          bool        `json:"hascolor"`
	ManufacturerName  string      `json:"manufacturername"`
	ModelId           string      `json:"modelid"`
	Name              string      `json:"name"`
	PointSymbol       interface{} `json:"pointsymbol"`
	State             LightState  `json:"state"`
	SwVersion         string      `json:"swversion"`
	Type              string      `json:"type"`
	UniqueId          string      `json:"uniqueid"`
}

type LightState struct {
	Alert          string    `json:"alert"`
	Bri            int       `json:"bri"`
	ColorloopSpeed int       `json:"colorloopspeed"`
	Ct             int       `json:"ct"`
	Effect         string    `json:"effect"`
	Hue            int       `json:"hue"`
	On             bool      `json:"on"`
	Reachable      bool      `json:"reachable"`
	Sat            int       `json:"sat"`
	TransitionTime int       `json:"transitiontime"`
	Xy             []float32 `json:"xy"`
}
