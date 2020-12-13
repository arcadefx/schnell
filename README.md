# schnell

This is the api server to handle login credential verification

## Install

$ go install

$ go build

## Run

$ ./schnell

## Testing

There is minimal testing due to time and learning 'testing' in Go.

See tests in "helpers" and "models"

Example test run:

$ cd models

$ go test -v

## Notes

* Learning/Re-learning Go and also learning "testing" in Go took time.
* The last time I touched Go was in 2017.
* I would create tests for all aspects of the server routes, controllers, models and helpers.

## Test and Cases if I had more time

### Cases
* Handle if request is coming from expected domain. Config driven for dev, stage and production.
* Handle if browser / ip address has changed.
* Create a cookie for the domain as a generic cookie was made in this App. It was made to see how to create a cookie.

### Tests
* Create tests for the controller methods.
* Create tests for routing.
* Create tests for http helper

