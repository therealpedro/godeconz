package api

type State struct {
	Config Config           `json:"config"`
	Lights map[string]Light `json:"lights"`
	Groups map[string]Group `json:"groups"`
	//TODO Rules
	//TODO Schedules
}
