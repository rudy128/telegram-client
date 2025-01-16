package src

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Define the arguments for the Python script
	phoneNumber := os.Getenv("TG_PHONE")
	apiID := os.Getenv("APP_ID")
	apiHash := os.Getenv("APP_HASH")
	userID := os.Getenv("RECEPTIONIST_ID")

	// Construct the command to run the Python script
	cmd := exec.Command("python3", "find_username.py", phoneNumber, apiID, apiHash, userID)

	// Capture the output from the Python script
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running the Python script:", err)
		return
	}

	// Print the output from the Python script
	username = out.String()
	// fmt.Println("Result from Python script:", result)
}
