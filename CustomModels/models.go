package CustomModels

type WeatherResponse struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	ResolvedAddress string  `json:"resolvedAddress"`
}
