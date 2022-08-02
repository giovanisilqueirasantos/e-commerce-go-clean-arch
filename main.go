package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	_authPresentation "github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/auth/presentation"
	_authRepo "github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/auth/repository"
	_authService "github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/auth/service"
	_authUsecase "github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/auth/usecase"
	_authValidator "github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/auth/validator"
	_codeRepo "github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/code/repository"
	_codeService "github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/code/service"
	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/config"
	_messageService "github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/message/service"
	_productPresentation "github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/product/presentation"
	_productRepo "github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/product/repository"
	_productUsecase "github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/product/usecase"
	_tokenService "github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/token/service"
	_userRepo "github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/user/repository"
	_userValidator "github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/user/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	conf, err := config.GetConf("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := sql.Open(`mysql`, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.Database.User, conf.Database.Pass, conf.Database.Host, conf.Database.Port, conf.Database.Name))

	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.Ping()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()

	e.Use(middleware.CORS())

	authRepo := _authRepo.NewAuthMysqlRepository(dbConn)
	codeRepo := _codeRepo.NewCodeMysqlRepository(dbConn)
	userRepo := _userRepo.NewUserMysqlRepository(dbConn)
	productRepo := _productRepo.NewProductMysqlRepository(dbConn)

	authService := _authService.NewAuthService()
	codeService := _codeService.NewCodeService(codeRepo)
	messageService := _messageService.NewMessageService()
	tokenService := _tokenService.NewTokenService()

	authValidator := _authValidator.NewAuthValidator()
	userValidator := _userValidator.NewUserValidator()

	authUsecase := _authUsecase.NewAuthUseCase(authService, tokenService, codeService, messageService, authRepo, userRepo)
	productUsecase := _productUsecase.NewProductUseCase(productRepo)

	_authPresentation.NewAuthHandler(e, authUsecase, authValidator, userValidator)
	_productPresentation.NewProductHandler(e, productUsecase, tokenService)

	log.Fatal(e.Start(conf.Server.Address))
}
