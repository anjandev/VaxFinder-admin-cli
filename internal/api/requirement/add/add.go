package add

import (
	"context"
	"moul.io/http2curl"
	"strconv"
	"vf-admin/internal/api"
	"vf-admin/internal/utils"
)

// HTTPOperation abstracts away the current HTTP operation
type HTTPOperation struct{}

var body api.CreateRequirementApiV1RequirementsPostJSONRequestBody
var authKey string

// SetAuthKey sets the authentication key to be used for the HTTP operation
func (HTTPOperation) SetAuthKey(key string) {
	authKey = key
}

// SetRequestURLArguments sets the appropriate url arguments for the HTTP operation
func (HTTPOperation) SetRequestURLArguments(args []string) error {
	return nil
}

// SetRequestBody sets the appropriate body for the HTTP operation
func (HTTPOperation) SetRequestBody(name, description string) error {
	body = api.CreateRequirementApiV1RequirementsPostJSONRequestBody{Name: name, Description: description}
	return nil
}

// GetDetails returns the details of the HTTP operation
func (HTTPOperation) GetDetails() (string, string, string) {
	return "add", "added", "requirement"
}

// GetVerboseResponseFieldNames returns the field names to be used when rendering the response as a table
func (HTTPOperation) GetVerboseResponseFieldNames() []string {
	return []string{"id", "name", "description", "created at"}
}

// GetResponseAsArray executes the HTTP operation and returns an array to be used when rendering the response as a table
func (HTTPOperation) GetResponseAsArray() ([][]string, error) {
	// Create the API client
	client, cErr := api.GetAPIClientFromKey(authKey)
	if cErr != nil {
		return nil, cErr
	}

	res, rErr := client.CreateRequirementApiV1RequirementsPostWithResponse(context.Background(), body)
	if rErr != nil {
		return nil, rErr
	}

	if res.StatusCode() != 200 {
		return nil, api.HandleHTTPError(res.StatusCode(), res.Body)
	}

	if res.JSON200 != nil {
		return [][]string{
			{
				strconv.Itoa(res.JSON200.Id), res.JSON200.Name, res.JSON200.Description, res.JSON200.CreatedAt.String(),
			},
		}, nil
	}

	return nil, nil
}

// GetAsCurlCommand returns the HTTP operation as a cURL command
func (HTTPOperation) GetAsCurlCommand(withKey bool) (*http2curl.CurlCommand, error) {
	// Create the HTTP Request (struct)
	req, rErr := api.NewCreateRequirementApiV1RequirementsPostRequest(utils.GetBaseURL(), body)
	if rErr != nil {
		return nil, rErr
	}
	// Attach auth key to request if it exists
	if authKey != "" && withKey {
		req.Header.Set("Authorization", "Bearer "+authKey)
	}
	return http2curl.GetCurlCommand(req)
}
