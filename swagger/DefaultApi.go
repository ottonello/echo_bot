package swagger

import (
	"encoding/json"
	"errors"

	"github.com/dghubble/sling"
)

type DefaultApi struct {
	basePath string
}

func NewDefaultApi() *DefaultApi {
	return &DefaultApi{
		basePath: "http://sducidsettings1:8080",
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
	path := "/extension-service/ws/api/v1/messages"

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
