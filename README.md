# microserviceutils

Utility package wrapping frequently used web service operations into easy to use functions.

When having web service/micro service functions like this:
```go
func myFunction(w http.ResponseWriter, r *http.Request) {
```
you can simply use the wrapper methods for dumping the content of the request/reponse to get quick diagnostics about requests being received, but also for marshalling objects that need to be sent via a http.ResponseWriter.

```go
microserviceutils.Dump(r)
```

which prints something like this to the console (for debugging/diagnostics), including what is in the body of the request/response, if anything:

```
GET /run HTTP/1.1
Host: localhost:25000
Accept: */*
Accept-Encoding: gzip, deflate, br
Connection: keep-alive
Postman-Token: 56c5e0d8-0b8e-42af-9a86-140d6c37c015
User-Agent: PostmanRuntime/7.26.2
```
