package domain

type Coordinate struct {
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	Timestamp int64   `json:"timestamp"`
}
