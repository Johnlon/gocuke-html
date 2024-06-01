package gocure

import (
	"gitlab.com/rodrigoodhin/gocure/embedded"
	"gitlab.com/rodrigoodhin/gocure/report/html"
)

type HTML struct {
	Config html.Data
}

type Embedded struct {
	Config embedded.Data
}
