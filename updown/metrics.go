package updown

// Metrics is a type that an updown Check's metrics
type Metrics struct {
	Apdex    float32  `json:"apdex"`
	Requests Requests `json:"requests"`
	Timings  Timings  `json:"timings"`
}

// Requests is a type that represents the Requests subtype of updown Metrics
type Requests struct {
	Samples        uint           `json:"samples"`
	Failures       uint           `json:"failures"`
	Satisfied      uint           `json:"satisfied"`
	Tolerated      uint           `json:"tolerated"`
	ByResponseTime ByResponseTime `json:"by_response_time"`
}

// ByResponseTime is a type that represents the ByResponseTime subtype of Requests
type ByResponseTime struct {
	Under125  uint `json:"under125"`
	Under250  uint `json:"under250"`
	Under500  uint `json:"under500"`
	Under1000 uint `json:"under1000"`
	Under2000 uint `json:"under2000"`
	Under4000 uint `json:"under4000"`
}

// ToBuckets is a method that converts ByResponseTIme structs into Prometheus buckets
func (x *ByResponseTime) ToBuckets() map[float64]uint64 {
	return map[float64]uint64{
		125.0:  uint64(x.Under125),
		250.0:  uint64(x.Under250),
		500.0:  uint64(x.Under500),
		1000.0: uint64(x.Under1000),
		2000.0: uint64(x.Under2000),
		4000.0: uint64(x.Under4000),
	}
}

// Timings is a type that represents the Timings subtype of updown Metrics
type Timings struct {
	Redirect   uint `json:"redirect"`
	NameLookup uint `json:"namelookup"`
	Connection uint `json:"connection"`
	Handshake  uint `json:"handshake"`
	Response   uint `json:"response"`
	Total      uint `json:"total"`
}
