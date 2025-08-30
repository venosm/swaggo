package main

import (
	"github.com/venosm/swaggo"
	"github.com/venosm/swaggo/testdata/delims/api"
	_ "github.com/venosm/swaggo/testdata/delims/docs"
)

func ReadDoc() string {
	doc, _ := swag.ReadDoc("CustomDelims")
	return doc
}

// @title Swagger Example API
// @version 1.0
// @description Testing custom template delimeters
// @termsOfService http://swagger.io/terms/

func main() {
	api.MyFunc()
}
