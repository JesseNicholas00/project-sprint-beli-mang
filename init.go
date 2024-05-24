package main

import (
	"github.com/JesseNicholas00/BeliMang/controllers"
	authCtrl "github.com/JesseNicholas00/BeliMang/controllers/auth"
	merchantCtrl "github.com/JesseNicholas00/BeliMang/controllers/merchant"
	"github.com/JesseNicholas00/BeliMang/middlewares"
	authRepo "github.com/JesseNicholas00/BeliMang/repos/auth"
	merchantRepo "github.com/JesseNicholas00/BeliMang/repos/merchant"
	authSvc "github.com/JesseNicholas00/BeliMang/services/auth"
	merchantSvc "github.com/JesseNicholas00/BeliMang/services/merchant"
	"github.com/JesseNicholas00/BeliMang/utils/ctxrizz"
	"github.com/JesseNicholas00/BeliMang/utils/logging"
	"github.com/jmoiron/sqlx"
)

func initControllers(
	cfg ServerConfig,
	db *sqlx.DB,
) (ctrls []controllers.Controller) {
	ctrlInitLogger := logging.GetLogger("main", "init", "controllers")
	defer func() {
		if r := recover(); r != nil {
			// add extra context to help debug potential panic
			ctrlInitLogger.Error("panic while initializing controllers: %s", r)
			panic(r)
		}
	}()

	dbRizzer := ctxrizz.NewDbContextRizzer(db)

	// withTxMw := middlewares.NewWithTxMiddleware(dbRizzer)

	authRepository := authRepo.NewAuthRepository(dbRizzer)
	authService := authSvc.NewAuthService(
		authRepository,
		cfg.jwtSecretKey,
		cfg.bcryptSaltCost,
	)
	authController := authCtrl.NewAuthController(authService)
	ctrls = append(ctrls, authController)

	authMw := middlewares.NewAuthMiddleware(authService)

	merchantRepository := merchantRepo.NewMerchantRepository(dbRizzer)
	merchantService := merchantSvc.NewMerchantService(merchantRepository)
	merchantController := merchantCtrl.NewMerchantController(
		merchantService,
		authMw,
	)
	ctrls = append(ctrls, merchantController)

	return
}
