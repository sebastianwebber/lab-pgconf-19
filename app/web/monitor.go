package web

import (
	ginprom "github.com/zsais/go-gin-prometheus"
)

func init() {
	p := ginprom.NewPrometheus("gin")
	p.Use(router)
}
