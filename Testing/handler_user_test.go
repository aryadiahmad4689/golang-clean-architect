package testing

import (
	"fmt"
	infrasructure "golang-clean-architecture/Infrasructure"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestHandlerUser(t *testing.T) {
	requestBody := strings.NewReader("nama=Ariadi&alamat=Makassar&age=20&married=true&date=2020-02-02")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/hello", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	infrasructure.AddUser(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
