package utils

import(
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// parse request body, needed for createBook function in controller
// request body is in json, so we need to unmarshall it as golang doesn't
// understand json
func ParseBody(r *http.Request, x interface{}){
	if body, err := ioutil.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return 
		}
	}
}