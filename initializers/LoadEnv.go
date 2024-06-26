package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

// ================================ Load values from .env ============================
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

}

//================================== END ==============================================
