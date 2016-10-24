package nim

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateServer(t *testing.T) {
	Convey("generate server from raml", t, func() {
		var apiDef raml.APIDefinition
		err := raml.ParseFile("../fixtures/server_resources/deliveries.raml", &apiDef)
		So(err, ShouldBeNil)

		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		ns := Server{
			Title:      apiDef.Title,
			APIDef:     &apiDef,
			APIDocsDir: "apidocs",
			Dir:        targetDir,
		}
		err = ns.Generate()
		So(err, ShouldBeNil)

		rootFixture := "./fixtures/server/delivery"
		checks := []struct {
			Result   string
			Expected string
		}{
			{"main.nim", "main.nim"},
			{"deliveries_api.nim", "deliveries_api.nim"},
		}

		for _, check := range checks {
			s, err := testLoadFile(filepath.Join(targetDir, check.Result))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile(filepath.Join(rootFixture, check.Expected))
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)
		}

		Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}
