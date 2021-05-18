package test

import (
	"fmt"
	_ "helloBee/routers"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	fmt.Printf("testing, TestBeego, Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestGet(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	beego.BeeApp.Handlers.ServeHTTP(w, req)
	t.Log("Get:", w.Body.String())
}

func TestPost(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", nil)
	beego.BeeApp.Handlers.ServeHTTP(w, req)
	t.Log("Post:", w.Body.String())
}

func TestPut(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/", nil)
	beego.BeeApp.Handlers.ServeHTTP(w, req)
	t.Log("Put:", w.Body.String())
}

func TestDelete(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/", nil)
	beego.BeeApp.Handlers.ServeHTTP(w, req)
	t.Log("Delete:", w.Body.String())
}
