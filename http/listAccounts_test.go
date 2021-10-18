package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain"
)

func TestListAllAccounts(t *testing.T) {
	req, err := http.NewRequest("GET", "/accounts", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandleFunc("/accounts", Server.ListAll)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := domain.Account{}
	if rr.Body.String() == nil {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}
