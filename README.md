# microserviceutils

Utility package for wrapping frequently used web service operations into easy to use functions.

When having web service/micro service functions like this:
```go
func myFunction(w http.ResponseWriter, r *http.Request) {
```
you can simply use the wrapper methods for dumping the content of the request/reponse to get quick diagnostics about requests being received, but also for marshalling objects that need to be sent via a http.ResponseWriter.
