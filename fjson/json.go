package fjson

import (
	"github.com/json-iterator/go"
)

var ConfigCompatibleWithStandardLibrary = jsoniter.Config{
	EscapeHTML:             true,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
	UseNumber:              true,
}.Froze()
