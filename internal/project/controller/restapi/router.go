package restapi

func (r *controller) RegisterRouter() {

	resource := r.Router.Group("/project", r.metricRequestCounter(), r.metricLatencyRequestChecker())

	v1 := resource.Group("/v1")

	v1.POST("/discovery-session", r.discoverySessionRequestHandler())
	v1.GET("/discovery-session/calendly", r.calendlyDiscoverySessionOnScheduledHandler())

}
