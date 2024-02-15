package restapi

func (r *controller) RegisterRouter() {

	resource := r.Router.Group("/support", r.metricRequestCounter(), r.metricLatencyRequestChecker())

	v1 := resource.Group("/v1")

	v1.POST("/contact-us", r.sendContactFormHandler())

}
