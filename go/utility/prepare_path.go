package utility

import (
	vs "github.com/voxgig-sdk/inkspan-document-sdk/go/utility/struct"

	"github.com/voxgig-sdk/inkspan-document-sdk/go/core"
)

func preparePathUtil(ctx *core.Context) string {
	point := ctx.Point

	var parts []any
	if p := vs.GetProp(point, "parts"); p != nil {
		if pl, ok := p.([]any); ok {
			parts = pl
		}
	}

	path := vs.Join(parts, "/", true)
	return path
}
