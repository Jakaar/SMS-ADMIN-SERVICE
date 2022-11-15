package model

type SmsDailyLimitation struct {
	Mobicom int `json:"mobicom"`
	Unitel  int `json:"unitel"`
	Skytel  int `json:"skytel"`
	Gmobile int `json:"gmobile"`
}
