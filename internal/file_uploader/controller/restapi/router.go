package restapi

func (r *controller) RegisterRouter() {

	resource := r.Router.Group("/file-uploader", r.metricRequestCounter(), r.metricLatencyRequestChecker())

	v1 := resource.Group("/v1")

	v1.POST("/upload", r.uploadFileHandler())

}
