package main

import (
	"backend-ccff/api"
	"backend-ccff/internal/env"
)

func main() {
	c := env.NewConfiguration()
	api.Start(c.App.Port, c.App.ServiceName, c.App.LoggerHttp, c.App.AllowedDomains)
}
