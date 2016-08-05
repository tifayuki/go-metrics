package metrics

import "github.com/prometheus/client_golang/prometheus"

// Register adds all the metrics in the provided namespace to the global
// metrics registry
func Register(n Namespace) {
	for _, m := range n.getMetrics() {
		prometheus.MustRegister(m)
	}
}

// Deregister removes all the metrics in the provided namespace from the
// global metrics registry
func Deregister(n Namespace) {
	for _, m := range n.getMetrics() {
		prometheus.Unregister(m)
	}
}
