package bind_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andviro/goldie"

	"github.com/go-noodle/bind"
	"github.com/go-noodle/noodle"
)

type bindTestCase struct {
	Payload    string
	Middleware noodle.Middleware
}

type testStruct struct {
	A int    `json:"a" form:"a" schema:"a"`
	B string `json:"b" form:"b" schema:"b"`
}

var bindBodyTestCases = []bindTestCase{
	{"alskdjasdklj", bind.Form(testStruct{})},
	{"a=1&b=Ololo", bind.Form(testStruct{})},
	{`{"a": 1, "b": "Ololo"}`, bind.JSON(testStruct{})},
	{"{}", bind.JSON(testStruct{})},
	{"", bind.JSON(testStruct{})},
}

func TestBind_Body(t *testing.T) {
	buf := new(bytes.Buffer)
	for _, tc := range bindBodyTestCases {
		w := httptest.NewRecorder()
		r, err := http.NewRequest("POST", "", bytes.NewBuffer([]byte(tc.Payload)))
		if err != nil {
			t.Fatal(err)
		}
		fmt.Fprintf(buf, "---\n")
		fmt.Fprintf(buf, "payload: %s\n", tc.Payload)
		n := tc.Middleware(func(w http.ResponseWriter, r *http.Request) {
			data, err := bind.Get(r)
			fmt.Fprintf(buf, "data: %#v\n", data)
			fmt.Fprintf(buf, "error: %+v\n", err)
		})
		n(w, r)
	}
	goldie.Assert(t, "bind-body", buf.Bytes())
}

var bindQueryTestCases = []bindTestCase{
	{"a=1&b=Ololo", bind.Query(testStruct{})},
	{"c=1&a=2", bind.Query(testStruct{})},
	{"b=1&a=qwe", bind.Query(testStruct{})},
}

func TestBind_Query(t *testing.T) {
	buf := new(bytes.Buffer)
	for _, tc := range bindQueryTestCases {
		w := httptest.NewRecorder()
		r, err := http.NewRequest("GET", "?"+tc.Payload, nil)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Fprintf(buf, "---\n")
		fmt.Fprintf(buf, "payload: %s\n", tc.Payload)
		n := tc.Middleware(func(w http.ResponseWriter, r *http.Request) {
			data, err := bind.Get(r)
			fmt.Fprintf(buf, "data: %#v\n", data)
			fmt.Fprintf(buf, "error: %+v\n", err)
		})
		n(w, r)
	}
	goldie.Assert(t, "bind-query", buf.Bytes())
}
