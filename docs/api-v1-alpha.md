example header

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
| **id** | *string* | unique name of app | `"example"` |


#### Curl Example

```bash
$ curl -n -X POST https://<your-iruka-server>.com/apps \
  -H "Content-Type: application/json" \
 \
  -d '{
  "id": "example"
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
$ curl -n -X DELETE https://<your-iruka-server>.com/apps/$APP_ID_OR_NAME \
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
$ curl -n https://<your-iruka-server>.com/apps/$APP_ID_OR_NAME
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
$ curl -n https://<your-iruka-server>.com/apps
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


#### Curl Example

```bash
$ curl -n -X PATCH https://<your-iruka-server>.com/apps/$APP_ID_OR_NAME \
  -H "Content-Type: application/json" \
 \
  -d '{
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


