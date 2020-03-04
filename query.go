package influxdb

import "github.com/influxdata/flux/ast"

// TODO(desa): These files are possibly a temporary. This is needed
// as a part of the source work that is being done.
// See https://github.com/influxdata/platform/issues/594 for more info.

// SourceQuery is a query for a source.
type SourceQuery struct {
	Query string `json:"query"`
	Type  string `json:"type"`
}

// FluxLanguageService is a service for interacting with flux code.
type FluxLanguageService interface {
	// Parse will take flux source code and produce a package.
	Parse(source string) *ast.Package
}
