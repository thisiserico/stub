# stub
> A test double for `http` calls

## 🧐 Motivation
Often times in a microservice architecture, a service will –sadly– depend on many others.
When that happens, if the service is not ready to be run without its dependencies, a CI/CD pipeline can quickly become hard to maintain.
Not only the service dependencies are needed to run the SUT, also their own dependencies.
This ends up in a pipeline running tons of services from which the SUT needs, maybe, a single endpoint to work.

`stub` provides a way to easily mock those required endpoints and nothing else.
Point your SUT dependencies against `stub`, create request expectations and test your logic controlling the execution flow.

## 👩‍💻 How to use
You first need to create an `expectation`.
The expected `http` method is a must. The `path` and `headers` are optional, yet recommended.
For every expectation, you also specify what you'd expect as `status code`, `response` and `headers`.

Say you're running `stub` on port `8080`:
```
curl -XPUT http://localhost:8080/expectation -d '{
	"using_method": "get",
	"against_path": "/complete/path",
	"with_headers": {
		"x-correlation-id": ["abc0123"]
	},
	"returns_code": 200,
	"with_response": {
		"string": "value",
		"number": 24
	},
	"and_headers": {
		"x-correlation-id": ["abc0123"],
		"cache-control": ["no-store", "no-cache"]
	}
}'
```

Once the expectation is created, you can get the given response, status code and headers if the given method, path and headers match.

## 🧙‍♂️ How to run
### Using plain go
The easiest way to run `stub`, specially when writing tests.
```
go run ./cmd/stub
```

### Using docker
Run the `docker` image exposing the port `8080`:
```
docker run --rm -p 8080:8080 thisiserico/stub:v1.0.3
```

## 🤷‍♀️ What's missing?
Some handy features are not yet implemented. Feel free to contribute on them!
* Delete previously created expectation
* Delete all previously created expectations
* Test different scenarios
