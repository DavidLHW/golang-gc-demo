# API Endpoints

These endpoints allow interaction with the app.

### GET /config

Retrieve current environment configuration of the app, only concerned with the `GOMEMLIMIT` and `GOGC` values.

**Response**

Returns the current values of the environment variables.

```json
{
  "gomemlimit": "30MiB",
  "gogc": "100"
}
```

### POST /config

Updates environment variables, only concerned with the `GOMEMLIMIT` and `GOGC` values.

**Request Body**

|         Name | Required |  Type  | Description                                                                                                   |
| -----------: | :------: | :----: | ------------------------------------------------------------------------------------------------------------- |
| `gomemlimit` |    no    | string | The soft memory limit for the runtime in human readable bytes using SI standard (eg. "30MiB", "44kB", "17MB") |
|       `gogc` |    no    | string | The initial garbage collection target percentage, 100 would mean 100% which is 2x of live heap size           |

**Response**

Returns the updated values of the environment variables.

```json
{
  "gomemlimit": "30MiB",
  "gogc": "100"
}
```
