package models

type USGSResponse struct {
	Type     string       `json:"type"`
	Metadata Metadata     `json:"metadata"`
	Features []Earthquake `json:"features"`
	Bbox     []float64    `json:"bbox"`
}
type Metadata struct {
	Generated int64  `json:"generated"`
	URL       string `json:"url"`
	Title     string `json:"title"`
	Status    int    `json:"status"`
	API       string `json:"api"`
	Count     int    `json:"count"`
}
type QuakeProperties struct {
	Mag     float64     `json:"mag"`
	Place   string      `json:"place"`
	Time    int64       `json:"time"`
	Updated int64       `json:"updated"`
	Tz      interface{} `json:"tz"`
	URL     string      `json:"url"`
	Detail  string      `json:"detail"`
	Felt    interface{} `json:"felt"`
	Cdi     interface{} `json:"cdi"`
	Mmi     interface{} `json:"mmi"`
	Alert   interface{} `json:"alert"`
	Status  string      `json:"status"`
	Tsunami int         `json:"tsunami"`
	Sig     int         `json:"sig"`
	Net     string      `json:"net"`
	Code    string      `json:"code"`
	Ids     string      `json:"ids"`
	Sources string      `json:"sources"`
	Types   string      `json:"types"`
	Nst     int         `json:"nst"`
	Dmin    float64     `json:"dmin"`
	Rms     float64     `json:"rms"`
	Gap     int         `json:"gap"`
	Magtype string      `json:"magType"`
	Type    string      `json:"type"`
	Title   string      `json:"title"`
}
type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
type Earthquake struct {
	ID         string          `json:"id"`
	Type       string          `json:"type"`
	Properties QuakeProperties `json:"properties"`
	Geometry   Geometry        `json:"geometry"`
}

type Dog struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
