# Best practices

## Endpoint functions

The structure for handling requests from users and giving response

```go
flow := func() error {
    HERE WE CALL USECASES in some cases calling adapters
    but mostly adapters are being called in usecases

    this is the last station of your error trace
    (ex.: not found <- repo.GetUser <- usecase.Info)

    you also can wrap the last error but our flow is needed to convert
    the error to api error - check below

}


switch err := flow(); {
case err == nil: (all good, do nothing)
case errors.Is(err, InvalidArgument):
    c.JSON(http.StatusBadRequest, "invalid argument")
    This error might only occur in Flow where we parse request

... list of all cases and lastly
default:
    c.JSON(http.StatusInternalServerError, "internal error")
    IF WE GET INTERNAL ERROR AND
NOT HANDLED ALL THE POSSIBLE ERRORS WE MUST LOG IT
}


```
