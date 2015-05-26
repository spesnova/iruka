# iruka API v1 Alpha
The iruka API allows you to manage the state of the Docker containers and the CoreOS cluster using JSON over HTTP.

## Errors
Failing responses will have an appropriate status and a JSON body containing more details about a particular error

### Error Attributes

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **id** | **string** | id of error raised | `"bad_request"` |
| **message** | **string** | end user message of error raised | `"request invalid, validate usage and try again"` |
| **url** | **string** | reference url with more information about the error | `"https://github.com/spesnova/iruka/blob/master/docs/errors.md"` |
