module github.com/wmsx/group_api

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/wmsx/group_svc v0.0.0-20200904055907-64203c72b0b4
	github.com/wmsx/pkg v0.0.0-20200904032714-e9c2f482fcbc
	github.com/wmsx/xconf v0.0.0-20200721142237-75926266fd1a
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0