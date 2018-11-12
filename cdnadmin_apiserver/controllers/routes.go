package http

import (
	"cdn-upload-admin/cdnadmin_apiserver/api/v1"
	"net/http"

	"github.com/gorilla/mux"
)

// Route : define Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Middleware  mux.MiddlewareFunc
}

// Routes : define Routes is []Route struct
type Routes []Route

// NewRouter : router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		r := router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name)
		if route.Middleware != nil {
			r.Handler(route.Middleware(handler))
		} else {
			r.Handler(handler)
		}
	}
	return router
}

var routes = Routes{
	Route{"Index", "GET", "/", Index, nil},
	Route{"Health", "GET", "/health", Health, nil},
	Route{"Version", "GET", "/version", Version, nil},

	// AUTH
	Route{"v1.Auth", "POST", "/api/v1/auth", v1.UserAuth, nil},

	// PROJECT
	Route{"v1.GetProjectList", "GET", "/api/v1/project/list", v1.GetProjectList, v1.TokenMiddleware},
	Route{"v1.AddProject", "POST", "/api/v1/project/add", v1.AddProject, v1.TokenMiddleware},
	Route{"v1.UpdateProject", "POST", "/api/v1/project/update", v1.UpdateProject, v1.TokenMiddleware},
	Route{"v1.DeleteProject", "GET", "/api/v1/project/delete", v1.DeleteProject, v1.TokenMiddleware},

	// PLATFORM
	Route{"v1.GetPlatformList", "GET", "/api/v1/platform/list", v1.GetPlatformList, v1.TokenMiddleware},
	Route{"v1.AddPlatform", "POST", "/api/v1/platform/add", v1.AddPlatform, v1.TokenMiddleware},
	Route{"v1.UpdatePlatform", "POST", "/api/v1/platform/update", v1.UpdatePlatform, v1.TokenMiddleware},
	Route{"v1.DeletePlatform", "GET", "/api/v1/platform/delete", v1.DeletePlatform, v1.TokenMiddleware},

	// FILESUFFIX
	Route{"v1.GetFileSuffixList", "GET", "/api/v1/suffix/list", v1.GetFileSuffixList, v1.TokenMiddleware},
	Route{"v1.AddFileSuffix", "POST", "/api/v1/suffix/add", v1.AddFileSuffix, v1.TokenMiddleware},
	Route{"v1.UpdateFileSuffix", "POST", "/api/v1/suffix/update", v1.UpdateFileSuffix, v1.TokenMiddleware},
	Route{"v1.DeleteFileSuffix", "GET", "/api/v1/suffix/delete", v1.DeleteFileSuffix, v1.TokenMiddleware},

	// CDNFILE
	Route{"v1.GetCDNFileList", "GET", "/api/v1/cdnfile/list", v1.GetCDNFileList, v1.TokenMiddleware},
	Route{"v1.AddCDNFile", "POST", "/api/v1/cdnfile/add", v1.AddCDNFile, v1.TokenMiddleware},
	Route{"v1.UpdateCDNFile", "POST", "/api/v1/cdnfile/update", v1.UpdateCDNFile, v1.TokenMiddleware},
	Route{"v1.DeleteCDNFile", "GET", "/api/v1/cdnfile/delete", v1.DeleteCDNFile, v1.TokenMiddleware},
	// CDN PUSH
	Route{"v1.CDNPush", "POST", "/api/v1/cdnfile/push", v1.CDNPush, v1.TokenMiddleware},

	// USER
	Route{"v1.GetUserList", "GET", "/api/v1/user/list", v1.GetUserList, v1.TokenMiddleware},
	Route{"v1.AddUser", "POST", "/api/v1/user/add", v1.AddUser, v1.TokenMiddleware},
	Route{"v1.UpdateUser", "POST", "/api/v1/user/update", v1.UpdateUser, v1.TokenMiddleware},
	Route{"v1.UpdateUserPwd", "POST", "/api/v1/user/updatePwd", v1.UpdateUserPwd, v1.TokenMiddleware},
	Route{"v1.ResetUserPwd", "POST", "/api/v1/user/reset", v1.ResetUserPwd, v1.TokenMiddleware},
	Route{"v1.SetUserAdmin", "POST", "/api/v1/user/setAdmin", v1.SetUserAdmin, v1.TokenMiddleware},
	Route{"v1.DeleteUser", "GET", "/api/v1/user/delete", v1.DeleteUser, v1.TokenMiddleware},
}
