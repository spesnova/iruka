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

## App

An app represents the program that you would like to deploy and run on iruka

### Attributes

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **id** | *uuid* | unique identifier of app | `"01234567-89ab-cdef-0123-456789abcdef"` |
| **name** | *string* | unique name of app | `"example"` |
| **web_url** | *string* | web URL of app | `"https://example.irukaapp.com/"` |
| **created_at** | *date-time* | when app was created | `"2012-01-01T12:00:00Z"` |
| **updated_at** | *date-time* | when app was updated | `"2012-01-01T12:00:00Z"` |

### App Create

Create a new app.

```
POST /apps
```

#### Optional Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **name** | *string* | unique name of app | `"example"` |


#### Curl Example

```bash
$ curl -n -X POST https://<your-iruka-server>.com/api/v1-alpha/apps \
  -H "Content-Type: application/json" \
 \
  -d '{
  "name": "example"
}'
```


#### Response Example

```
HTTP/1.1 201 Created
```

```json
{
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "web_url": "https://example.irukaapp.com/",
  "created_at": "2012-01-01T12:00:00Z",
  "updated_at": "2012-01-01T12:00:00Z"
}
```

### App Delete

Delete an existing app.

```
DELETE /apps/{app_id_or_name}
```


#### Curl Example

```bash
$ curl -n -X DELETE https://<your-iruka-server>.com/api/v1-alpha/apps/$APP_ID_OR_NAME \
  -H "Content-Type: application/json" \
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "web_url": "https://example.irukaapp.com/",
  "created_at": "2012-01-01T12:00:00Z",
  "updated_at": "2012-01-01T12:00:00Z"
}
```

### App Info

Info for existing app.

```
GET /apps/{app_id_or_name}
```


#### Curl Example

```bash
$ curl -n https://<your-iruka-server>.com/api/v1-alpha/apps/$APP_ID_OR_NAME
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "web_url": "https://example.irukaapp.com/",
  "created_at": "2012-01-01T12:00:00Z",
  "updated_at": "2012-01-01T12:00:00Z"
}
```

### App List

List existing apps.

```
GET /apps
```


#### Curl Example

```bash
$ curl -n https://<your-iruka-server>.com/api/v1-alpha/apps
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
[
  {
    "id": "01234567-89ab-cdef-0123-456789abcdef",
    "name": "example",
    "web_url": "https://example.irukaapp.com/",
    "created_at": "2012-01-01T12:00:00Z",
    "updated_at": "2012-01-01T12:00:00Z"
  }
]
```

### App Update

Update an existing app.

```
PATCH /apps/{app_id_or_name}
```

#### Optional Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **name** | *string* | unique name of app | `"example"` |


#### Curl Example

```bash
$ curl -n -X PATCH https://<your-iruka-server>.com/api/v1-alpha/apps/$APP_ID_OR_NAME \
  -H "Content-Type: application/json" \
 \
  -d '{
  "name": "example"
}'
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "web_url": "https://example.irukaapp.com/",
  "created_at": "2012-01-01T12:00:00Z",
  "updated_at": "2012-01-01T12:00:00Z"
}
```


## Container

Container encapsulate running processes of an app on iruka.

### Attributes

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **id** | *uuid* | unique identifier of container | `"01234567-89ab-cdef-0123-456789abcdef"` |
| **app_id** | *uuid* | unique identifier of app the container is belong to | `"01234567-89ab-cdef-0123-456789abcdef"` |
| **name** | *string* | unique name of container | `"example.web.1"` |
| **image** | *string* | resource URI of the Docker image (including tag) of the container | `"quay.io/spesnova/example:latest"` |
| **size** | *string* | container size (default “1X”) | `"2X"` |
| **command** | *string* | command used to start this process | `"bundle exec rails server"` |
| **ports** | *array* | expose ports | `[80,8080]` |
| **type** | *string* | type of process (either "web", "worker", "timer", or "run") | `"web"` |
| **desired_state** | *string* | desired state of process (either exited or up) | `"up"` |
| **state** | *string* | current state of process (either exited or up) | `"up"` |
| **created_at** | *date-time* | when container was created | `"2012-01-01T12:00:00Z"` |
| **updated_at** | *date-time* | when container was updated | `"2012-01-01T12:00:00Z"` |

### Container Create

Create and run a new container.

```
POST /apps/{app_id_or_name}/containers
```

#### Optional Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **image** | *string* | resource URI of the Docker image (including tag) of the container | `"quay.io/spesnova/example:latest"` |
| **size** | *string* | container size (default “1X”) | `"2X"` |
| **command** | *string* | command used to start this process | `"bundle exec rails server"` |
| **type** | *string* | type of process (either "web", "worker", "timer", or "run") | `"web"` |
| **ports** | *array* | expose ports | `[80,8080]` |


#### Curl Example

```bash
$ curl -n -X POST https://<your-iruka-server>.com/api/v1-alpha/apps/$APP_ID_OR_NAME/containers \
  -H "Content-Type: application/json" \
 \
  -d '{
  "image": "quay.io/spesnova/example:latest",
  "size": "2X",
  "command": "bundle exec rails server",
  "type": "web",
  "ports": [
    80,
    8080
  ]
}'
```


#### Response Example

```
HTTP/1.1 201 Created
```

```json
{
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "app_id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example.web.1",
  "image": "quay.io/spesnova/example:latest",
  "size": "2X",
  "command": "bundle exec rails server",
  "ports": [
    80,
    8080
  ],
  "type": "web",
  "desired_state": "up",
  "state": "up",
  "created_at": "2012-01-01T12:00:00Z",
  "updated_at": "2012-01-01T12:00:00Z"
}
```

### Container Delete

Delete an existing container.

```
DELETE /apps/{app_id_or_name}/containers/{container_id_or_name}
```


#### Curl Example

```bash
$ curl -n -X DELETE https://<your-iruka-server>.com/api/v1-alpha/apps/$APP_ID_OR_NAME/containers/$CONTAINER_ID_OR_NAME \
  -H "Content-Type: application/json" \
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "app_id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example.web.1",
  "image": "quay.io/spesnova/example:latest",
  "size": "2X",
  "command": "bundle exec rails server",
  "ports": [
    80,
    8080
  ],
  "type": "web",
  "desired_state": "up",
  "state": "up",
  "created_at": "2012-01-01T12:00:00Z",
  "updated_at": "2012-01-01T12:00:00Z"
}
```

### Container Info

Info for existing container.

```
GET /apps/{app_id_or_name}/containers/{container_id_or_name}
```


#### Curl Example

```bash
$ curl -n https://<your-iruka-server>.com/api/v1-alpha/apps/$APP_ID_OR_NAME/containers/$CONTAINER_ID_OR_NAME
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "app_id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example.web.1",
  "image": "quay.io/spesnova/example:latest",
  "size": "2X",
  "command": "bundle exec rails server",
  "ports": [
    80,
    8080
  ],
  "type": "web",
  "desired_state": "up",
  "state": "up",
  "created_at": "2012-01-01T12:00:00Z",
  "updated_at": "2012-01-01T12:00:00Z"
}
```

### Container List

List existing containers.

```
GET /apps/{app_id_or_name}/containers
```


#### Curl Example

```bash
$ curl -n https://<your-iruka-server>.com/api/v1-alpha/apps/$APP_ID_OR_NAME/containers
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
[
  {
    "id": "01234567-89ab-cdef-0123-456789abcdef",
    "app_id": "01234567-89ab-cdef-0123-456789abcdef",
    "name": "example.web.1",
    "image": "quay.io/spesnova/example:latest",
    "size": "2X",
    "command": "bundle exec rails server",
    "ports": [
      80,
      8080
    ],
    "type": "web",
    "desired_state": "up",
    "state": "up",
    "created_at": "2012-01-01T12:00:00Z",
    "updated_at": "2012-01-01T12:00:00Z"
  }
]
```

### Container Update

Update options and restart an existing container.

```
PATCH /apps/{app_id_or_name}/containers/{container_id_or_name}
```

#### Optional Parameters

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **image** | *string* | resource URI of the Docker image (including tag) of the container | `"quay.io/spesnova/example:latest"` |
| **size** | *string* | container size (default “1X”) | `"2X"` |
| **command** | *string* | command used to start this process | `"bundle exec rails server"` |
| **type** | *string* | type of process (either "web", "worker", "timer", or "run") | `"web"` |
| **ports** | *array* | expose ports | `[80,8080]` |


#### Curl Example

```bash
$ curl -n -X PATCH https://<your-iruka-server>.com/api/v1-alpha/apps/$APP_ID_OR_NAME/containers/$CONTAINER_ID_OR_NAME \
  -H "Content-Type: application/json" \
 \
  -d '{
  "image": "quay.io/spesnova/example:latest",
  "size": "2X",
  "command": "bundle exec rails server",
  "type": "web",
  "ports": [
    80,
    8080
  ]
}'
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "app_id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example.web.1",
  "image": "quay.io/spesnova/example:latest",
  "size": "2X",
  "command": "bundle exec rails server",
  "ports": [
    80,
    8080
  ],
  "type": "web",
  "desired_state": "up",
  "state": "up",
  "created_at": "2012-01-01T12:00:00Z",
  "updated_at": "2012-01-01T12:00:00Z"
}
```



