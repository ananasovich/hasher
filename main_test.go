package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestParseArgs(t *testing.T) {
	cases := []struct {
		name     string
		args     []string
		expected []string
	}{
		{
			name:     "one arg success",
			args:     []string{"http://yandex.com"},
			expected: []string{"http://yandex.com"},
		},
		{
			name:     "few args success",
			args:     []string{"adjust.com", "http://yandex.com", "reddit.com/r/notfunny"},
			expected: []string{"http://adjust.com", "http://yandex.com", "http://reddit.com/r/notfunny"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := make(chan url.URL, 20)
			parseArgs(c.args, result)
			for _, u := range c.expected {
				res := <-result
				if u != res.String() {
					t.Errorf("expected %v insted got %v", u, res)
				}
			}

		})
	}
}

func TestGetHash(t *testing.T) {
	s := httptest.NewServer(&handlerMock{})
	defer s.Close()
	s2 := httptest.NewServer(&handlerMock{body: "some body"})
	defer s2.Close()

	cases := []struct {
		name     string
		arg      string
		expected string
	}{
		{
			name:     "success empty body",
			arg:      s.URL,
			expected: "d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			name:     "success with some body",
			arg:      s2.URL,
			expected: "328c30fae61cd119cd177c061d1ac11f",
		},
		{
			name:     "silent fail",
			arg:      "broken",
			expected: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := getHash(c.arg)
			if c.expected != result {
				t.Errorf("expected %v insted got %v", c.expected, result)
			}
		})
	}
}

type handlerMock struct {
	body string
}

func (h *handlerMock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.body))
}
