/*
 * NxFs
 *
 * Simple file access APIs for the Entando Nx subsystem
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nxsiteman

import (
	"os"
)

const envVarBrowsableFs = "BROWSABLE_FS"
const fsBaseDir = "./browsableFS"

var browsableFsPath = ""

//Response return a ImplResponse struct filled
func Response(code int, body interface{}) ImplResponse {
	return ImplResponse{Code: code, Body: body}
}

//ErrorResponse return a ImplResponse struct filled with an error
func ErrorResponse(code int, errorCode string, errorMessage string) *ImplResponse {
	return &ImplResponse{Code: code, Body: &Error{Code: errorCode, Message: errorMessage}}
}

// GetBrowsableFsRootPath - return the root path of the file system to browse
func GetBrowsableFsRootPath() string {
	if "" == browsableFsPath {
		browsableFsPath = os.Getenv(envVarBrowsableFs)
		if "" == browsableFsPath {
			browsableFsPath = fsBaseDir
		}
	}
	return browsableFsPath
}
