package main

import (
	"dc_haur/src/starter"
)

// @title           HAUR Swagger API
// @version         1.0
// @description     This is a HAUR server.
// @termsOfService  http://swagger.io/terms/
// @license.url   	http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.basic  BasicAuth
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	starter.StartApp()
}
