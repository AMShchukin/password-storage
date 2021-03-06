// Package spec provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package spec

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// AllPasswordData defines model for AllPasswordData.
type AllPasswordData = []struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// LoginDataRequest defines model for LoginDataRequest.
type LoginDataRequest struct {
	Password *string `json:"password,omitempty"`
	User     *string `json:"user,omitempty"`
}

// LoginDataResponse defines model for LoginDataResponse.
type LoginDataResponse struct {
	SessionId *string `json:"sessionId,omitempty"`
}

// NewPasswordDataRequest defines model for NewPasswordDataRequest.
type NewPasswordDataRequest struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewPasswordDataResponse defines model for NewPasswordDataResponse.
type NewPasswordDataResponse struct {
	Key *string `json:"key,omitempty"`
}

// PasswordDataByKey defines model for PasswordDataByKey.
type PasswordDataByKey struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// PostApiV1CreateNewPasswordDataJSONBody defines parameters for PostApiV1CreateNewPasswordData.
type PostApiV1CreateNewPasswordDataJSONBody = NewPasswordDataRequest

// PostApiV1SignInJSONBody defines parameters for PostApiV1SignIn.
type PostApiV1SignInJSONBody = LoginDataRequest

// PostApiV1CreateNewPasswordDataJSONRequestBody defines body for PostApiV1CreateNewPasswordData for application/json ContentType.
type PostApiV1CreateNewPasswordDataJSONRequestBody = PostApiV1CreateNewPasswordDataJSONBody

// PostApiV1SignInJSONRequestBody defines body for PostApiV1SignIn for application/json ContentType.
type PostApiV1SignInJSONRequestBody = PostApiV1SignInJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create new password data
	// (POST /api/v1/createNewPasswordData)
	PostApiV1CreateNewPasswordData(ctx echo.Context) error
	// Get All Password Data
	// (GET /api/v1/getAllPasswordData)
	GetApiV1GetAllPasswordData(ctx echo.Context) error
	// Get password data by key
	// (GET /api/v1/getPasswordDataByKey/{key})
	GetApiV1GetPasswordDataByKeyKey(ctx echo.Context, key string) error
	// Login
	// (POST /api/v1/signIn)
	PostApiV1SignIn(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostApiV1CreateNewPasswordData converts echo context to params.
func (w *ServerInterfaceWrapper) PostApiV1CreateNewPasswordData(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostApiV1CreateNewPasswordData(ctx)
	return err
}

// GetApiV1GetAllPasswordData converts echo context to params.
func (w *ServerInterfaceWrapper) GetApiV1GetAllPasswordData(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetApiV1GetAllPasswordData(ctx)
	return err
}

// GetApiV1GetPasswordDataByKeyKey converts echo context to params.
func (w *ServerInterfaceWrapper) GetApiV1GetPasswordDataByKeyKey(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "key" -------------
	var key string

	err = runtime.BindStyledParameterWithLocation("simple", false, "key", runtime.ParamLocationPath, ctx.Param("key"), &key)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter key: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetApiV1GetPasswordDataByKeyKey(ctx, key)
	return err
}

// PostApiV1SignIn converts echo context to params.
func (w *ServerInterfaceWrapper) PostApiV1SignIn(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostApiV1SignIn(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/api/v1/createNewPasswordData", wrapper.PostApiV1CreateNewPasswordData)
	router.GET(baseURL+"/api/v1/getAllPasswordData", wrapper.GetApiV1GetAllPasswordData)
	router.GET(baseURL+"/api/v1/getPasswordDataByKey/:key", wrapper.GetApiV1GetPasswordDataByKeyKey)
	router.POST(baseURL+"/api/v1/signIn", wrapper.PostApiV1SignIn)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RVwY7TMBD9FWvgGDUpe8ttl5VWKxBaAeKCOHiTaettYhvbaRVVlYAVB07wKXvgULGC",
	"/QXnj5CdhnSbVKpQe6pjz8ybeTPzuoBE5FJw5EZDvACdTDCn/niaZVdU67lQ6Tk11F0xg7l/k0pIVIah",
	"/5pi6X5MKRFi0EYxPoZlADOaFdjzsgyaG3F9g4mB9oIqRUv3/VKMGXe4r/Fjgdp0UeU6uV7oQqPaE3kD",
	"SUvBNXahNGrNBL9M9wz5Cueb1O0s4RDEdbB2FdEP1hdyM95Z+aL2O3Tm7orxkXDGKepEMWmY4BCD/WEf",
	"7IrYn/a++u5+/lSfqy/VrX2wd8T+ItXX6pO9s7/tyt7bVfWtuoUADDMZ/q/zDJWusYeDaBC5QoRETiWD",
	"GE4G0eAEApDUTHztIZUsnA3DRCE1uMW/p0rUzXaEUbMeHLgS2pxK9m74vNcvAFWPyZlIPbeJ4Aa5D0Sl",
	"zFjiQ4U32iXabKo7PVU4ghiehO0qh+s9DndMomff4TGFKcRGFegv6tnxVT6LhsfLYj2jPo3HzX87QZK5",
	"jSRzqokukgS1HhVZVg78FOkiz6kqIYaaRsJxThopIKlj0pk1PRqj6dGxMfb05wLr9lx0XTrURAejZhuq",
	"h5LTLGsrPG8qbIm4QEOczVXHZoOFzk6HiymWy33Y6Lg6RXD7oGiOBpWG+P32Dk+9CXNHtzcQAKc5/nto",
	"uRkJlVPTKkXQUY4PR2S/q3Q9/Dt+H00YuS7J1Ntut2GnWdMJzcb8ku8hEm9qw+OoQufPdS89iI6Bfwgl",
	"8OFciOXfAAAA///jA9OZzwgAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
