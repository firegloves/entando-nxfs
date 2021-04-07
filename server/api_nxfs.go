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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer) Router {
	return &DefaultApiController{service: s}
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{
		{
			"ApiNxfsBrowseEncodedPathGet",
			strings.ToUpper("Get"),
			"/api/nxfs/browse/{EncodedPath}",
			c.ApiNxfsBrowseEncodedPathGet,
		},
		{
			"ApiNxfsObjectsEncodedPathDelete",
			strings.ToUpper("Delete"),
			"/api/nxfs/objects/{EncodedPath}",
			c.ApiNxfsObjectsEncodedPathDelete,
		},
		{
			"ApiNxfsObjectsEncodedPathGet",
			strings.ToUpper("Get"),
			"/api/nxfs/objects/{EncodedPath}",
			c.ApiNxfsObjectsEncodedPathGet,
		},
		{
			"ApiNxfsObjectsEncodedPathPublishPut",
			strings.ToUpper("Put"),
			"/api/nxfs/objects/{EncodedPath}/publish",
			c.ApiNxfsObjectsEncodedPathPublishPut,
		},
		{
			"ApiNxfsObjectsEncodedPathPut",
			strings.ToUpper("Put"),
			"/api/nxfs/objects/{EncodedPath}",
			c.ApiNxfsObjectsEncodedPathPut,
		},
		{
			"ApiNxfsObjectsEncodedPathUnpublishPut",
			strings.ToUpper("Put"),
			"/api/nxfs/objects/{EncodedPath}/unpublish",
			c.ApiNxfsObjectsEncodedPathUnpublishPut,
		},
	}
}

// ApiNxfsBrowseEncodedPathGet - Gets the list of objects in a directory
func (c *DefaultApiController) ApiNxfsBrowseEncodedPathGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	encodedPath := params["EncodedPath"]
	maxdepth, err := parseInt32Parameter(query.Get("maxdepth"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.ApiNxfsBrowseEncodedPathGet(r.Context(), encodedPath, maxdepth)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ApiNxfsObjectsEncodedPathDelete - Deletes an object
func (c *DefaultApiController) ApiNxfsObjectsEncodedPathDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	encodedPath := params["EncodedPath"]
	result, err := c.service.ApiNxfsObjectsEncodedPathDelete(r.Context(), encodedPath)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ApiNxfsObjectsEncodedPathGet - Gets an object
func (c *DefaultApiController) ApiNxfsObjectsEncodedPathGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	encodedPath := params["EncodedPath"]
	result, err := c.service.ApiNxfsObjectsEncodedPathGet(r.Context(), encodedPath)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ApiNxfsObjectsEncodedPathPublishPut - Publishes an object
func (c *DefaultApiController) ApiNxfsObjectsEncodedPathPublishPut(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	encodedPath := params["EncodedPath"]
	result, err := c.service.ApiNxfsObjectsEncodedPathPublishPut(r.Context(), encodedPath)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ApiNxfsObjectsEncodedPathPut - Creates or updates an object
func (c *DefaultApiController) ApiNxfsObjectsEncodedPathPut(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	encodedPath := params["EncodedPath"]
	fileObject := &FileObject{}
	if err := json.NewDecoder(r.Body).Decode(&fileObject); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.ApiNxfsObjectsEncodedPathPut(r.Context(), encodedPath, *fileObject)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ApiNxfsObjectsEncodedPathUnpublishPut - Publishes an object
func (c *DefaultApiController) ApiNxfsObjectsEncodedPathUnpublishPut(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	encodedPath := params["EncodedPath"]
	result, err := c.service.ApiNxfsObjectsEncodedPathUnpublishPut(r.Context(), encodedPath)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
