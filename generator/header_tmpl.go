package generator

var headerTmpl = `
package {{.}}
// Generated by https://github.com/cloudescape/gowsdl
// Do not modify
// Copyright (c) 2014, Cloudescape. All rights reserved.
import (
	"time"
	gowsdl "github.com/cloudescape/gowsdl/generator"
	{{/*range .Imports*/}}
		{{/*.*/}}
	{{/*end*/}}
)
`
