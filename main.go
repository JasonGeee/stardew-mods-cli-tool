package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	psCommand := `Get-ChildItem -Path C:\ -Recurse -Directory -ErrorAction SilentlyContinue | Where-Object { $_.FullName -like "*\Stardew Valley\Mods" } | Select-Object -ExpandProperty FullName`

	// Create the command to run PowerShell with the desired command
	cmd := exec.Command("powershell", "-Command", psCommand)

	// Run the command and capture the output
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error outputting command: ", err)
		return
	}

	// Convert the output to a string and trim any whitespace
	result := strings.TrimSpace(string(output))

	// Print the result
	fmt.Println("Found directory path:", result)
}
