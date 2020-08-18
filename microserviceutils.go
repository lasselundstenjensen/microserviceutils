package microserviceutils

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"io"
	"net/http/httputil"
	"fmt"
)

func getBody(r interface{}) []byte {
	switch v := r.(type) {
	case *http.Request:
		return readBodyContent(v.Body)
	case *http.Response:
		return readBodyContent(v.Body)
	default:
		break
	}
	return []byte{}
}

func readBodyContent(stream io.ReadCloser) []byte {
	defer stream.Close()
	body, readErr := ioutil.ReadAll(stream)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body
}

func Unmarshal(r interface{}, o interface{}) {
	error := json.Unmarshal(getBody(r), &o)
	if error != nil {
		log.Fatal(error)
	}
}

func UnmarshalBytes(b []byte, o interface{}) {
	error := json.Unmarshal(b, &o)
	if error != nil {
		log.Fatal(error)
	}
}

func Marshal(o interface{}) []byte {
	value, error := json.Marshal(o)
	if error != nil {
		log.Fatal(error)
	}
	return value
}

func MarshalIndent(o interface{}) []byte {
	value, error := json.MarshalIndent(o, "", "    ")
	if error != nil {
		log.Fatal(error)
	}
	return value
}

func Dump(r interface{}) {
	switch v := r.(type) {
	case *http.Request:
		dump, err := httputil.DumpRequest(v, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(dump))
		break
	case *http.Response:
		dump, err := httputil.DumpResponse(v, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(dump))
		break
	default:
		break
	}
}
