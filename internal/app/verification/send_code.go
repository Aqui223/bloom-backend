package verification

import (
	"fmt"

	"github.com/slipe-fun/skid-backend/internal/domain"
	"github.com/slipe-fun/skid-backend/internal/service"
	"github.com/slipe-fun/skid-backend/internal/service/logger"
)

func (v *VerificationApp) CreateAndSendCode(email string) error {
	code, err := service.GenerateNumericCode(6)
	if err != nil {
		logger.LogError(err.Error(), "verification-app")
		return domain.Failed("failed to generate numeric code")
	}

	createdCode, err := v.verification.Create(&domain.VerificationCode{
		Email: email,
		Code:  code,
	})
	if err != nil {
		logger.LogError(err.Error(), "verification-app")
		return domain.Failed("failed to create code")
	}

	sendEmailError := service.SendMail(
		fmt.Sprintf("Your code - %s", createdCode.Code),
		fmt.Sprintf("Hello! Your Bloom verification code - %s", createdCode.Code),
		email,
	)
	if sendEmailError != nil {
		logger.LogError(sendEmailError.Error(), "verification-app")
		v.verification.DeleteByEmailAndCode(email, createdCode.Code)
		return domain.Failed("failed to send email")
	}

	return nil
}
