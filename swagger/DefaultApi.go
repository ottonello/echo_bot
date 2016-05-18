package swagger

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/dghubble/sling"
)

type DefaultApi struct {
	basePath string
}

func NewDefaultApi() *DefaultApi {
	return &DefaultApi{
		basePath: "http://sducidsettings1:8080/extension-service/ws",
	}
}

func NewDefaultApiWithBasePath(basePath string) *DefaultApi {
	return &DefaultApi{
		basePath: basePath,
	}
}

/**
 * Send a message
 * Send a message to a set of users
 * @param message Message to be sent
 * @return MessageResult
 */
//func (a DefaultApi) ApiV1MessagesPost (message Message) (MessageResult, error) {
func (a DefaultApi) ApiV1MessagesPost(message Message) (MessageResult, error) {

	_sling := sling.New().Post(a.basePath)

	// create path and map variables
	path := "/api/v1/messages"

	_sling = _sling.Path(path)

	// accept header
	accepts := []string{"application/json"}
	for key := range accepts {
		_sling = _sling.Set("Accept", accepts[key])
		break // only use the first Accept
	}

	// body params
	_sling = _sling.BodyJSON(message)

	var successPayload = new(MessageResult)

	// We use this map (below) so that any arbitrary error JSON can be handled.
	// FIXME: This is in the absence of this Go generator honoring the non-2xx
	// response (error) models, which needs to be implemented at some point.
	var failurePayload map[string]interface{}

	httpResponse, err := _sling.Receive(successPayload, &failurePayload)

	if err == nil {
		// err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
		if failurePayload != nil {
			// If the failurePayload is present, there likely was some kind of non-2xx status
			// returned (and a JSON payload error present)
			var str []byte
			str, err = json.Marshal(failurePayload)
			if err == nil { // For safety, check for an error marshalling... probably superfluous
				// This will return the JSON error body as a string
				err = errors.New(string(str))
			}
		} else {
			// So, there was no network-type error, and nothing in the failure payload,
			// but we should still check the status code
			if httpResponse == nil {
				// This should never happen...
				err = errors.New("No HTTP Response received.")
			} else if code := httpResponse.StatusCode; 200 > code || code > 299 {
				err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
			}
		}
	}

	return *successPayload, err
}

/**
 * Get an extension definition
 *
 * @param id Id of the extension to retrieve
 * @return Extension
 */
//func (a DefaultApi) InternalV1ExtensionsIdGet (id string) (Extension, error) {
func (a DefaultApi) InternalV1ExtensionsIdGet(id string) (Extension, error) {

	_sling := sling.New().Get(a.basePath)

	// create path and map variables
	path := "/internal/v1/extensions/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	_sling = _sling.Path(path)

	// accept header
	accepts := []string{"application/json"}
	for key := range accepts {
		_sling = _sling.Set("Accept", accepts[key])
		break // only use the first Accept
	}

	var successPayload = new(Extension)

	// We use this map (below) so that any arbitrary error JSON can be handled.
	// FIXME: This is in the absence of this Go generator honoring the non-2xx
	// response (error) models, which needs to be implemented at some point.
	var failurePayload map[string]interface{}

	httpResponse, err := _sling.Receive(successPayload, &failurePayload)

	if err == nil {
		// err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
		if failurePayload != nil {
			// If the failurePayload is present, there likely was some kind of non-2xx status
			// returned (and a JSON payload error present)
			var str []byte
			str, err = json.Marshal(failurePayload)
			if err == nil { // For safety, check for an error marshalling... probably superfluous
				// This will return the JSON error body as a string
				err = errors.New(string(str))
			}
		} else {
			// So, there was no network-type error, and nothing in the failure payload,
			// but we should still check the status code
			if httpResponse == nil {
				// This should never happen...
				err = errors.New("No HTTP Response received.")
			} else if code := httpResponse.StatusCode; 200 > code || code > 299 {
				err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
			}
		}
	}

	return *successPayload, err
}

/**
 * Update an extension
 *
 * @param id Id of the extension to modify
 * @param extension extension definition
 * @return Extension
 */
//func (a DefaultApi) InternalV1ExtensionsIdPut (id string, extension Extension) (Extension, error) {
func (a DefaultApi) InternalV1ExtensionsIdPut(id string, extension Extension) (Extension, error) {

	_sling := sling.New().Put(a.basePath)

	// create path and map variables
	path := "/internal/v1/extensions/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	_sling = _sling.Path(path)

	// accept header
	accepts := []string{"application/json"}
	for key := range accepts {
		_sling = _sling.Set("Accept", accepts[key])
		break // only use the first Accept
	}

	// body params
	_sling = _sling.BodyJSON(extension)

	var successPayload = new(Extension)

	// We use this map (below) so that any arbitrary error JSON can be handled.
	// FIXME: This is in the absence of this Go generator honoring the non-2xx
	// response (error) models, which needs to be implemented at some point.
	var failurePayload map[string]interface{}

	httpResponse, err := _sling.Receive(successPayload, &failurePayload)

	if err == nil {
		// err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
		if failurePayload != nil {
			// If the failurePayload is present, there likely was some kind of non-2xx status
			// returned (and a JSON payload error present)
			var str []byte
			str, err = json.Marshal(failurePayload)
			if err == nil { // For safety, check for an error marshalling... probably superfluous
				// This will return the JSON error body as a string
				err = errors.New(string(str))
			}
		} else {
			// So, there was no network-type error, and nothing in the failure payload,
			// but we should still check the status code
			if httpResponse == nil {
				// This should never happen...
				err = errors.New("No HTTP Response received.")
			} else if code := httpResponse.StatusCode; 200 > code || code > 299 {
				err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
			}
		}
	}

	return *successPayload, err
}

/**
 * Create a new extension
 *
 * @param extension New extension description
 * @return Extension
 */
//func (a DefaultApi) InternalV1ExtensionsPost (extension NewExtension) (Extension, error) {
func (a DefaultApi) InternalV1ExtensionsPost(extension NewExtension) (Extension, error) {

	_sling := sling.New().Post(a.basePath)

	// create path and map variables
	path := "/internal/v1/extensions"

	_sling = _sling.Path(path)

	// accept header
	accepts := []string{"application/json"}
	for key := range accepts {
		_sling = _sling.Set("Accept", accepts[key])
		break // only use the first Accept
	}

	// body params
	_sling = _sling.BodyJSON(extension)

	var successPayload = new(Extension)

	// We use this map (below) so that any arbitrary error JSON can be handled.
	// FIXME: This is in the absence of this Go generator honoring the non-2xx
	// response (error) models, which needs to be implemented at some point.
	var failurePayload map[string]interface{}

	httpResponse, err := _sling.Receive(successPayload, &failurePayload)

	if err == nil {
		// err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
		if failurePayload != nil {
			// If the failurePayload is present, there likely was some kind of non-2xx status
			// returned (and a JSON payload error present)
			var str []byte
			str, err = json.Marshal(failurePayload)
			if err == nil { // For safety, check for an error marshalling... probably superfluous
				// This will return the JSON error body as a string
				err = errors.New(string(str))
			}
		} else {
			// So, there was no network-type error, and nothing in the failure payload,
			// but we should still check the status code
			if httpResponse == nil {
				// This should never happen...
				err = errors.New("No HTTP Response received.")
			} else if code := httpResponse.StatusCode; 200 > code || code > 299 {
				err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
			}
		}
	}

	return *successPayload, err
}
