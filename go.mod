module github.com/nsip/n3-sif2json

go 1.15

replace (
	github.com/nsip/n3-sif2json/SIFSpec/3.4.6 => ../SIFSpec/3.4.6
	github.com/nsip/n3-sif2json/SIFSpec/3.4.7 => ../SIFSpec/3.4.7
)

require (
	github.com/basgys/goxml2json v1.1.0
	github.com/cdutwhu/debog v0.2.10
	github.com/cdutwhu/gotil v0.1.5
	github.com/cdutwhu/n3-util v0.3.5
	github.com/clbanning/mxj v1.8.4
	github.com/davecgh/go-spew v1.1.1
	github.com/go-xmlfmt/xmlfmt v0.0.0-20191208150333-d5b6f63a941b
	github.com/labstack/echo-contrib v0.9.0
	github.com/labstack/echo/v4 v4.1.17
	github.com/nats-io/jwt v1.0.1 // indirect
	github.com/nats-io/nats.go v1.10.0
	github.com/nats-io/nkeys v0.2.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/peterbourgon/mergemap v0.0.0-20130613134717-e21c03b7a721
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/sys v0.0.0-20200909081042-eff7692f9009 // indirect
)
