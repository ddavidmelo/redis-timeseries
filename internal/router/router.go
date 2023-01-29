package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
)

type Handler struct {
	ginHandler     *gin.Engine
	grpcwebHandler *grpcweb.WrappedGrpcServer
}

func InitRouter(grpcWebServer *grpcweb.WrappedGrpcServer) *Handler {
	router := gin.New()
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST, OPTIONS, GET, PUT"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Grpc-Web, X-User-Agent"},
		AllowCredentials: true,
	}))

	router.Use(static.Serve("/", static.LocalFile("./ui/dist", false)))
	router.NoRoute(func(c *gin.Context) {
		c.File("./ui/dist/index.html")
	})

	return &Handler{
		ginHandler:     router,
		grpcwebHandler: grpcWebServer,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	if contentType == "application/grpc-web-text" {
		h.grpcwebHandler.ServeHTTP(w, req)
		return
	}
	h.ginHandler.ServeHTTP(w, req)
}
