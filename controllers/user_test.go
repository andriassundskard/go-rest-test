package controllers_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andriassundskard/go-rest-test/handlers"
)

var (
	server   *httptest.Server
	reader   io.Reader //Ignore this for now
	usersUrl string
)

func init() {
	server = httptest.NewServer(api.Handlers()) //Creating new server with the user handlers

	usersUrl = fmt.Sprintf("%s/user", server.URL) //Grab the address for the API endpoint
}

func TestCreateUser(t *testing.T) {
	userJson := `{"name": "Test Testsen", "gender": "Male", "age": 20}`

	reader = strings.NewReader(userJson) //Convert string to reader

	request, err := http.NewRequest("POST", usersUrl, reader) //Create request with JSON body

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	if res.StatusCode != 201 {
		t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
	}
}
