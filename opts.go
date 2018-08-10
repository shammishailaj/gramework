package gramework

import "github.com/valyala/fasthttp"

// OptUseServer sets fasthttp.Server instance to use
func OptUseServer(s *fasthttp.Server) func(*App) {
	return func(a *App) {
		if a != nil && s != nil {
			a.server = s
			a.server.Handler = a.handler()
		}
	}
}

// OptMaxRequestBodySize sets new MaxRequestBodySize in the server used at the execution time.
// All OptUseServer will overwrite this setting 'case OptUseServer replaces the whole server instance
// with a new one.
func OptMaxRequestBodySize(new int) func(*App) {
	return func(a *App) {
		if a != nil && a.server != nil {
			a.server.MaxRequestBodySize = new
		}
	}
}