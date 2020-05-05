package nethttp

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// put this FILE into github.com/opentracing-contrib/go-stdlib/nethttp

// SimpleTracer creates a new instance of Jaeger tracer.
// func SimpleTracer(serviceName string) opentracing.Tracer {
// 	cfg, _ := config.FromEnv()
// 	cfg.ServiceName = serviceName
// 	cfg.Sampler.Type = "const"
// 	cfg.Sampler.Param = 1

// 	// a quick hack to ensure random generators get different seeds, which are based on current time.
// 	time.Sleep(100 * time.Millisecond)
// 	tracer, _, _ := cfg.NewTracer()
// 	return tracer
// }

// TracerOptions :
type TracerOptions struct {
	tracer  opentracing.Tracer
	options []MWOption
}

// NewTrOpt :
func NewTrOpt(tr opentracing.Tracer, opts ...MWOption) *TracerOptions {
	return &TracerOptions{
		tracer:  tr,
		options: opts,
	}
}

// TraceMiddleware :
func (tropt *TracerOptions) TraceMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	opts := mwOptions{
		opNameFunc: func(r *http.Request) string {
			return "HTTP " + r.Method
		},
		spanFilter:   func(r *http.Request) bool { return true },
		spanObserver: func(span opentracing.Span, r *http.Request) {},
		urlTagFunc: func(u *url.URL) string {
			return u.String()
		},
	}
	for _, opt := range tropt.options {
		opt(&opts)
	}

	return func(c echo.Context) error {
		r, w := c.Request(), c.Response().Writer
		if !opts.spanFilter(r) {
			return next(c)
		}
		ctx, _ := tropt.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		sp := tropt.tracer.StartSpan(opts.opNameFunc(r), ext.RPCServerOption(ctx))
		ext.HTTPMethod.Set(sp, r.Method)
		ext.HTTPUrl.Set(sp, opts.urlTagFunc(r.URL))
		opts.spanObserver(sp, r)

		// set component name, use "net/http" if caller does not specify
		componentName := opts.componentName
		if componentName == "" {
			componentName = defaultComponentName
		}
		ext.Component.Set(sp, componentName)

		sct := &statusCodeTracker{ResponseWriter: w}
		r = r.WithContext(opentracing.ContextWithSpan(r.Context(), sp))

		defer func() {
			ext.HTTPStatusCode.Set(sp, uint16(sct.status))
			if sct.status >= http.StatusInternalServerError || !sct.wroteheader {
				ext.Error.Set(sp, true)
			}
			sp.Finish()
		}()

		resp := c.Response()
		resp.Writer = sct.wrappedResponseWriter()
		c.SetResponse(resp)
		c.SetRequest(r)
		return next(c)
	}
}

// TraceHandler :
func TraceHandler(tr opentracing.Tracer, h echo.HandlerFunc, optList ...MWOption) echo.HandlerFunc {
	opts := mwOptions{
		opNameFunc: func(r *http.Request) string {
			return "HTTP " + r.Method
		},
		spanFilter:   func(r *http.Request) bool { return true },
		spanObserver: func(span opentracing.Span, r *http.Request) {},
		urlTagFunc: func(u *url.URL) string {
			return u.String()
		},
	}
	for _, opt := range optList {
		opt(&opts)
	}
	return func(c echo.Context) error {
		r, w := c.Request(), c.Response().Writer
		if !opts.spanFilter(r) {
			return h(c)
		}
		ctx, _ := tr.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		sp := tr.StartSpan(opts.opNameFunc(r), ext.RPCServerOption(ctx))
		ext.HTTPMethod.Set(sp, r.Method)
		ext.HTTPUrl.Set(sp, opts.urlTagFunc(r.URL))
		opts.spanObserver(sp, r)

		// set component name, use "net/http" if caller does not specify
		componentName := opts.componentName
		if componentName == "" {
			componentName = defaultComponentName
		}
		ext.Component.Set(sp, componentName)

		sct := &statusCodeTracker{ResponseWriter: w}
		r = r.WithContext(opentracing.ContextWithSpan(r.Context(), sp))

		defer func() {
			ext.HTTPStatusCode.Set(sp, uint16(sct.status))
			if sct.status >= http.StatusInternalServerError || !sct.wroteheader {
				ext.Error.Set(sp, true)
			}
			sp.Finish()
		}()

		resp := c.Response()
		resp.Writer = sct.wrappedResponseWriter()
		c.SetResponse(resp)
		c.SetRequest(r)
		return h(c)
	}
}
