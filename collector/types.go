package collector

// System is a type that represents a Prometheus Exporter system
type System struct {
	Namespace string
	Subsystem string
	Version   string
}
