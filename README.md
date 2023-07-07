# NightBird
Mock server &amp; contract validator

Currently run integration-testing is painful due to the unexpected behaviour in external 3rd party API (sometimes this testing step takes so long to complete), also sometimes 3rd party API changes the responses without expectation.

So, we want to solve this problem by:
- Generating a mock server that serve the same response as the real 3rd party API's.
- A tool/service that validates the responses from the real 3rd party API and notify any change in their schema (trigger validation manually or in a regular basis).
- This tool is going to be fully configurable.
- The request/response parameters will be store in one place and it's going to be used for mocking server as for contract validation.

## Description:
This service implement something called `pellets` and `pellets-flows`. Basically each `pellet` is a json file with request and response information
and the `pellets-flows` contains the order of the calls to be made to the real API endpoints.

### Pellet:
* `id`:_id_ - id of the pellet.
* `name`: _string_ - name or description of the pellet.
* `url`: _string_ - URL to be mocked/tested.
* `request`:_object_ - request object to be used for mocked for integration tests and used for real API validation:
  - `headers`:_list_ - a list of headers to be mocked/validated.
  - `data`: _object_ - the payload of the request to be mocked/validated.
* `response`: _object_ - response object to be used for mocked for integration tests and used for real API validation:
  - `headers`: _list_ - a list of headers to be mocked/validated.
  - `data`: _object_ the payload of the request to be mocked/validated.
  - `response-code`: _int_ - response code to be validated/mocked (200, 201, 400, etc).
* `http-method`: _string_ - method to be used for mocking/validation, (POST, GET, PUT, UPDATE, PATCH).
* `project-name`:_string_ - name of the project that mainly will use this pellet (ej. service-social-communications).
* `is-active`: _bool_ - indicates if the pellet is active for use (only active pellets will be called for mocking and validation process).

#### example:
```json
{
  "id": 99,
  "name": "json-placeholder test",
  "url": "https://jsonplaceholder.typicode.com/posts",
  "request": {
    "headers": {
      "Content-type": "application/json; charset=UTF-8"
    },
    "data": {
      "title": "foo",
      "body": "bar",
      "userId": 1
    }
  },
  "response": {
    "headers": {
      "Location": "http://jsonplaceholder.typicode.com/posts/101",
      "Content-type": "application/json; charset=UTF-8"
    },
    "data": {
      "id": 101,
      "title": "foo",
      "body": "bar",
      "userId": 1
    },
    "response-code": 201
  },
  "http-method": "POST",
  "project-name": "jsonplaceholder-API",
  "is-active": true
}
```

### Pellet-flow:
* `id`:_int_ - identifier for pellet-flow.
* `name`:_string_ - name of the pellet-flow.
* `project_name`:_string_ - name of the project that depends on this pellet_flow.
* `flow`:_list_ - contains the order of the calls and the id of the pellet to be used. 
   - `order-id`:_int_ - order id of the pellet call.
   - `pellet_id`:_int_ - id of the pellet to be used as request for real API.
* `is-active`:_bool_ - indicate if the pellet_flow is active (only active pellet_flows will be call).

#### Example:
```json
{
  "id": 99,
  "name": "jsonplaceholder test flow",
  "project-name": "jsonplaceholder-API",
  "flow": [
    {
      "order-id": 1,
      "pellet-id": 99
    },
    {
      "order-id": 2,
      "pellet-id": 98
    }
  ],
  "is-active": true
}
```

> **Note**: In order to use a pellet, it needs to be located inside `./json-pellets/pellets` for single `pellets` 
> and `./json-pellets/pellets-flows` for `pellet-flows`

## Use:

### In order to start the service, just run 
```bash
└─(master ✹)──> make run-server   
```

### To use the mocker server:
just point your request url to:
```bash
localhost:8080/mocker/serve?destination-url=[REAL_API_URL]
```
like:
```bash
localhost:8080/mocker/serve?destination-url=https://api.linkedin.com/v2/posts
```
> **Important**: the [REAL_API_URL] needs to match with the `pellet` url that you want to
> use to mock, otherwise the call will not be successful.

### To use the API validator:
just make a post call to 
```
localhost:8080/validator/start
```
and check the logs in your terminal, it will show something like this:
```bash
2023/04/21 14:50:18 "POST http://localhost:8080/validator/start HTTP/1.1" from [::1]:56229 - 200 23B in 72.875µs
2023/04/21 14:50:24 [SUCCESS] -> pellet request validated successfully - pellet-id:99 pellet-flow-id:99 flow-order:1 flow-name:jsonplaceholder test flow
2023/04/21 14:50:24 [ERROR] -> expected response data is different from actual response data CompareExpectedResponseBody - actual: map[body:bar id:101 title:foo userId:1] | expected: &map[body:bar-bad id:10199 title:foo-bad userId:99]

```

## Next Steps:
- Complete Slack integration in `./utils/Notifier`
- Add React UI (a good approach can be this - [A Golang + ReactJS Application](https://medium.com/@madhanganesh/golang-react-application-2aaf3bca92b1))
- Add Unit testing
- Extend functionality to be able to get pellets from a NoSql database like mongo
- Update Dockerfile to create image service
- Create deployment configuration files
- Add more documentation
- Implement Metrics