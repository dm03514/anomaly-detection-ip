package blacklist

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_ServeHTTP_DecodeError(t *testing.T) {
	h := Handler{}
	recorder := httptest.NewRecorder()
	h.ServeHTTP(recorder, &http.Request{
		Body: ioutil.NopCloser(bytes.NewReader([]byte("invalid"))),
	})
	resp := recorder.Result()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected code %q received %q",
			http.StatusInternalServerError, resp.StatusCode)
	}
}

func TestHandler_ServeHTTP_BlacklistTestError(t *testing.T) {
	h := Handler{
		Blacklist: &StubBlacklist{
			Result: Result{},
			Error:  fmt.Errorf("error error"),
		},
	}
	recorder := httptest.NewRecorder()
	ipRequest := Payload{}
	payload, err := ipRequest.Bytes()
	if err != nil {
		t.Error(err)
	}

	h.ServeHTTP(recorder, &http.Request{
		Body: ioutil.NopCloser(bytes.NewReader(payload)),
	})
	resp := recorder.Result()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected code %q received %q",
			http.StatusInternalServerError, resp.StatusCode)
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if "error error\n" != string(bs) {
		t.Errorf("expected %q received %q", "error error", string(bs))
	}
}

func TestHandler_ServeHTTP_BlacklistResult(t *testing.T) {
	metadata := map[string]string{
		"Name":        "Hello",
		"LastRefresh": time.Now().String(),
	}
	blacklistResult := Result{
		Found: true,
		Providers: []Provider{
			{
				Name:     "Hello",
				Metadata: metadata,
			},
		},
	}

	h := Handler{
		Blacklist: &StubBlacklist{
			Result: blacklistResult,
			Error:  nil,
		},
	}
	recorder := httptest.NewRecorder()
	ipRequest := Payload{}
	payload, err := ipRequest.Bytes()
	if err != nil {
		t.Error(err)
	}

	h.ServeHTTP(recorder, &http.Request{
		Body: ioutil.NopCloser(bytes.NewReader(payload)),
	})
	resp := recorder.Result()
	if http.StatusOK != resp.StatusCode {
		t.Errorf("expected code %q received %q",
			http.StatusOK, resp.StatusCode)
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	resultPayload, err := blacklistResult.Bytes()
	if err != nil {
		t.Error(err)
	}
	if string(resultPayload)+"\n" != string(bs) {
		t.Errorf("expected %q received %q", string(resultPayload), string(bs))
	}
}
