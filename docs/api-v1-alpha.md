sample

## App

FIXME

### Attributes

| Name | Type | Description | Example |
| ------- | ------- | ------- | ------- |
| **created_at** | *date-time* | when app was created | `"2012-01-01T12:00:00Z"` |
| **id** | *uuid* | unique identifier of app | `"01234567-89ab-cdef-0123-456789abcdef"` |
| **name** | *string* | unique name of app | `"name"` |
| **updated_at** | *date-time* | when app was updated | `"2012-01-01T12:00:00Z"` |

### App Create

Create a new app.

```
POST /apps
```


#### Curl Example

```bash
$ curl -n -X POST https://github.com/spesnova/iruka/apps \
  -H "Content-Type: application/json" \
 \
  -d '{
}'
```


#### Response Example

```
HTTP/1.1 201 Created
```

```json
{
  "created_at": "2012-01-01T12:00:00Z",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "name",
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
$ curl -n -X DELETE https://github.com/spesnova/iruka/apps/$APP_ID_OR_NAME \
  -H "Content-Type: application/json" \
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "created_at": "2012-01-01T12:00:00Z",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "name",
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
$ curl -n https://github.com/spesnova/iruka/apps/$APP_ID_OR_NAME
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
{
  "created_at": "2012-01-01T12:00:00Z",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "name",
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
$ curl -n https://github.com/spesnova/iruka/apps
```


#### Response Example

```
HTTP/1.1 200 OK
```

```json
[
  {
    "created_at": "2012-01-01T12:00:00Z",
    "id": "01234567-89ab-cdef-0123-456789abcdef",
    "name": "name",
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
$ curl -n -X PATCH https://github.com/spesnova/iruka/apps/$APP_ID_OR_NAME \
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
  "created_at": "2012-01-01T12:00:00Z",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "name",
  "updated_at": "2012-01-01T12:00:00Z"
}
```


