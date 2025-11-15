package service

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateUsername(fullName string) string {
	parts := strings.Fields(fullName)

	var firstName, lastName string

	if len(parts) > 0 {
		firstName = strings.ToLower(parts[0])
	}
	if len(parts) > 1 {
		lastName = strings.ToLower(parts[1])
	}

	if lastName == "" {
		lastName = "user"
	}

	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(10000)

	randomStr := fmt.Sprintf("%04d", randomNum)

	return fmt.Sprintf("%s_%s_%s", firstName, lastName, randomStr)
}
