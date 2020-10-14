# microserviceutils

Utility package wrapping frequently used web service operations into easy to use functions.

### 1: Dump request/response header and body

When having web service/micro service functions like this:
```go
func myFunction(w http.ResponseWriter, r *http.Request) {
   ...
}
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

{
    "body": "This is a request."
}
```

### 2: Unmarshal request/response directly into instance of type

This can also be done for automatically unmarshalling the body content of a HTTP response or request into a given type, without having to do boilerplate code to first read the body of the request/response and then unmarshal it.

```go
t := MyType{}

err := microserviceutils.Unmarshal(r, t)
```

here 'r' is the request or response, and 't' is the type to unmarshal the request/response body JSON into.

