// Package language exposes the flux parser as an interface.
package fluxlang

import (
	"github.com/influxdata/flux/ast"
	"github.com/influxdata/flux/parser"
	"github.com/influxdata/influxdb"
)

// DefaultService is the default language service.
var DefaultService influxdb.FluxLanguageService = defaultService{}

type defaultService struct{}

func (d defaultService) Parse(source string) *ast.Package {
	return parser.ParseSource(source)
}
