package main

import (
	"github.com/JesseNicholas00/BeliMang/controllers"
	authCtrl "github.com/JesseNicholas00/BeliMang/controllers/auth"
	imageCtrl "github.com/JesseNicholas00/BeliMang/controllers/image"
	merchantCtrl "github.com/JesseNicholas00/BeliMang/controllers/merchant"
	"github.com/JesseNicholas00/BeliMang/middlewares"
	authRepo "github.com/JesseNicholas00/BeliMang/repos/auth"
	merchantRepo "github.com/JesseNicholas00/BeliMang/repos/merchant"
	authSvc "github.com/JesseNicholas00/BeliMang/services/auth"
	merchantSvc "github.com/JesseNicholas00/BeliMang/services/merchant"
	"github.com/JesseNicholas00/BeliMang/types/role"
	"github.com/JesseNicholas00/BeliMang/utils/ctxrizz"
	"github.com/JesseNicholas00/BeliMang/utils/logging"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/jmoiron/sqlx"
)

func initControllers(
	cfg ServerConfig,
	db *sqlx.DB,
	uploader *manager.Uploader,
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
		dbRizzer,
		cfg.jwtSecretKey,
		cfg.bcryptSaltCost,
	)
	authController := authCtrl.NewAuthController(authService)
	ctrls = append(ctrls, authController)

	adminMw := middlewares.NewAuthMiddleware(authService, role.Admin)

	merchantRepository := merchantRepo.NewMerchantRepository(dbRizzer)
	merchantService := merchantSvc.NewMerchantService(merchantRepository, dbRizzer)

	merchantController := merchantCtrl.NewMerchantController(
		merchantService,
		adminMw,
	)
	ctrls = append(ctrls, merchantController)

	imageCtrl := imageCtrl.NewImageController(uploader, cfg.awsS3BucketName, adminMw)
	ctrls = append(ctrls, imageCtrl)

	return
}
