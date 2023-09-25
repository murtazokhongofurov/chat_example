package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/kafka_example/api-gateway/api/docs"
	v1 "github.com/kafka_example/api-gateway/api/handler/v1"
	"github.com/kafka_example/api-gateway/api/middleware"
	"github.com/kafka_example/api-gateway/api/tokens"
	"github.com/kafka_example/api-gateway/config"
	"github.com/kafka_example/api-gateway/pkg/logger"
	"github.com/kafka_example/api-gateway/services"
	swaggerfile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	h "net/http"
)

type Options struct {
	Cfg            config.Config
	Log            logger.Logger
	CasbinEnforcer *casbin.Enforcer
	ServiceManager services.ServiceManagerI
}

// @title           ChatApp API
// @version         1.0
// @description     This is a sample server celler server.

// @contact.name   Murtazoxon Gofurov
// @contact.email  example@gmail.com

// @BasePath  /v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(opt *Options) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corConfig := cors.DefaultConfig()
	corConfig.AllowAllOrigins = true
	corConfig.AllowCredentials = true
	corConfig.AllowHeaders = []string{"*"}
	corConfig.AllowBrowserExtensions = true
	corConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corConfig))

	jwtHandler := tokens.JWTHandler{
		SigninKey: opt.Cfg.SigningKey,
		Log:       opt.Log,
	}
	handlerV1 := v1.New(&v1.HandlerV1Option{
		Cfg:            &opt.Cfg,
		Log:            opt.Log,
		JwtHandler:     jwtHandler,
		ServiceManager: opt.ServiceManager,
	})
	router.Use(middleware.NewAuth(opt.CasbinEnforcer, jwtHandler, config.Load()))

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(h.StatusOK, gin.H{
			"message": "Server is running!!!",
		})
	})

	router.MaxMultipartMemory = 8 << 20 // 8 Mib

	api := router.Group("/v1")
	//user
	api.POST("/user", handlerV1.PostUser)
	api.GET("/user/:id", handlerV1.GetUser)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfile.Handler, url))

	return router
}
