package restapi

func (r *controller) RegisterRouter() {

	resource := r.Router.Group("/newsletter", r.metricRequestCounter(), r.metricLatencyRequestChecker())

	v1 := resource.Group("/v1")

	v1.POST("/subscribers", r.subscribeHandler())
}
