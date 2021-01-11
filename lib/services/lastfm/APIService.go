package lastfm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/go-querystring/query"
	"github.com/joho/godotenv"
)

// API holds methods for interacting with the Last.fm API
type API struct {
	baseURL string
	apiKey  string
}

func (lfm API) buildParams(method string, params interface{}) string {
	defaultParams := struct {
		APIKey string `url:"api_key"`
		Format string `url:"format"`
		Method string `url:"method"`
	}{
		APIKey: lfm.apiKey,
		Format: "json",
		Method: method,
	}

	fmt.Println(defaultParams)

	defaultValues, _ := query.Values(defaultParams)
	paramValues, _ := query.Values(params)

	if len(paramValues) > 0 {
		return fmt.Sprintf("%s&%s", defaultValues.Encode(), paramValues.Encode())
	}

	return defaultValues.Encode()
}

// MakeRequest calls the lastfm api with the given parameters
func (lfm API) MakeRequest(method string, params interface{}) *http.Response {
	queryparams := lfm.buildParams(method, params)

	log.Print(queryparams)

	resp, err := http.Get(lfm.baseURL + "?" + queryparams)

	if err != nil {
		log.Println("Error! ", err)
	}

	return resp
}

// ParseResponse parses a JSON respone from the last.fm api
func (lfm API) ParseResponse(response *http.Response, output interface{}) *ErrorResponse {
	defer response.Body.Close()

	responseBody, _ := ioutil.ReadAll(response.Body)
	errorResponse := &ErrorResponse{}

	json.Unmarshal(responseBody, output)
	json.Unmarshal(responseBody, errorResponse)

	if errorResponse.Error != 0 {
		return errorResponse
	}

	return nil
}

// UserInfo fetches a user's info from the last.fm API
func (lfm API) UserInfo(username string) (*ErrorResponse, *UserInfoResponse) {
	params := UserInfoParams{
		Username: username,
	}

	userInfo := &UserInfoResponse{}

	response := lfm.MakeRequest("user.getInfo", params)

	err := lfm.ParseResponse(response, userInfo)

	return err, userInfo
}

// ValidateUser validates that a given username exists in last.fm
func (lfm API) ValidateUser(username string) bool {
	err, _ := lfm.UserInfo(username)

	return err == nil
}

// CreateAPIService creates an instance of the lastfm api service object
func CreateAPIService() *API {
	godotenv.Load()

	apiKey := os.Getenv("LAST_FM_API_KEY")

	service := &API{
		baseURL: "http://ws.audioscrobbler.com/2.0/",
		apiKey:  apiKey,
	}

	return service
}