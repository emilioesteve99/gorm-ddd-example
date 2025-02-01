package authDomainConverters

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	authDomainCommands "gorm-ddd-example/src/auth/domain/command"
	authDomainModels "gorm-ddd-example/src/auth/domain/model"
	appErrors "gorm-ddd-example/src/common/application/model"
	domainConverters "gorm-ddd-example/src/common/domain/converter"
	"gorm-ddd-example/src/config"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
	"time"
)

type LoginCommandToJwtAuthResponseConverter struct {
	secretKey string
}

func NewLoginCommandToJwtAuthResponseConverter(cfg config.Config) domainConverters.ConverterWithExtraArgs[authDomainCommands.LoginCommand, userDomainModels.User, authDomainModels.AuthResponse] {
	return &LoginCommandToJwtAuthResponseConverter{
		secretKey: cfg.Secret,
	}
}

func (c *LoginCommandToJwtAuthResponseConverter) Convert(command authDomainCommands.LoginCommand, user userDomainModels.User, ctx context.Context) (authDomainModels.AuthResponse, error) {
	compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(command.Password))
	if compareErr != nil {
		return authDomainModels.AuthResponse{}, appErrors.AppError{
			Code:    appErrors.UnauthorizedCode,
			Message: appErrors.UnauthorizedMsg,
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, signErr := token.SignedString([]byte(c.secretKey))
	if signErr != nil {
		return authDomainModels.AuthResponse{}, appErrors.BuildUnknownError(signErr)
	}

	return authDomainModels.AuthResponse{
		AccessToken: tokenString,
	}, nil
}
