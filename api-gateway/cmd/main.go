package main

import (
	"github.com/casbin/casbin/v2"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/casbin/casbin/v2/util"
	"github.com/kafka_example/api-gateway/api"
	"github.com/kafka_example/api-gateway/config"
	"github.com/kafka_example/api-gateway/pkg/logger"
	"github.com/kafka_example/api-gateway/services"
)

func main() {
	cfg := config.Load()
	log := logger.New("debug", "chatapp")

	grpcConn, err := services.NewServiceManager(cfg)
	if err != nil {
		log.Error("error grpc conn: ", logger.Error(err))
	}
	casbinEnforcer, err := casbin.NewEnforcer(cfg.AuthFilePath, cfg.CsvFilePath)

	err = casbinEnforcer.LoadPolicy()
	if err != nil {
		log.Error("casbin error load policy", logger.Error(err))
		return
	}

	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManagerImpl).AddMatchingFunc("keyMatch", util.KeyMatch)
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManagerImpl).AddMatchingFunc("keyMatch3", util.KeyMatch3)

	server := api.New(&api.Options{
		Cfg:            cfg,
		Log:            log,
		CasbinEnforcer: casbinEnforcer,
		ServiceManager: grpcConn,
	})
	if err := server.Run(cfg.HttpPort); err != nil {
		log.Fatal("error running in port: 8080", logger.Error(err))
	}
}
