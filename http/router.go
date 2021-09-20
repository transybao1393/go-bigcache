package http

type Router struct{}

func InitRouter() IRouter {
	return &Router{}
}

func (r *Router) GET() {
	return
}
func (r *Router) POST() {
	return
}
