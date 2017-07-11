package main

import (
	"testing"

	"github.com/get-ion/ion"
	"github.com/get-ion/ion/httptest"
)

func TestSessionsEncodeDecode(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app, httptest.URL("http://example.com"))

	es := e.GET("/set").Expect()
	es.Status(ion.StatusOK)
	es.Cookies().NotEmpty()
	es.Body().Equal("All ok session setted to: ion")

	e.GET("/get").Expect().Status(ion.StatusOK).Body().Equal("The name on the /set was: ion")
	// delete and re-get
	e.GET("/delete").Expect().Status(ion.StatusOK)
	e.GET("/get").Expect().Status(ion.StatusOK).Body().Equal("The name on the /set was: ")
	// set, clear and re-get
	e.GET("/set").Expect().Body().Equal("All ok session setted to: ion")
	e.GET("/clear").Expect().Status(ion.StatusOK)
	e.GET("/get").Expect().Status(ion.StatusOK).Body().Equal("The name on the /set was: ")
}
