package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// Version is manually updated when creating a new tag
var Version = "v0.1b"

const (
	LabolinkHl7ReciverExe  = "LABOLINK_HL7_RECIVER.exe"
	LabolinkHl7ReciverWExe = "LABOLINK_HL7_RECIVERw.exe"
	LabolinkHl7TranferExe  = "LABOLINK_HL7_TRANSFER.exe"
	LabolinkHl7TranferWExe = "LABOLINK_HL7_TRANSFERw.exe"
	InstallFile            = "install.bat"
	InstallTranferFile     = "install-tranfer.bat"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "rift",
		Short: "rift is a Golang code generator",
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
	}

	var rename = &cobra.Command{
		Use:   "rename [name]",
		Short: "Rename the Labolink SDK",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := strings.ToLower(args[0])
			// reciver
			path, err := filepath.Abs(LabolinkHl7ReciverExe)
			if err != nil {
				log.Fatal(err)
			}
			newName := strings.ReplaceAll(LabolinkHl7ReciverExe, "RECIVER", fmt.Sprintf("RECIVER_%s", strings.ToUpper(name)))
			if err := os.Rename(path, newName); err != nil {
				log.Fatal(err)
			}

			// reciver w
			path, err = filepath.Abs(LabolinkHl7ReciverWExe)
			if err != nil {
				log.Fatal(err)
			}
			newName = strings.ReplaceAll(LabolinkHl7ReciverWExe, "RECIVER", fmt.Sprintf("RECIVER_%s", strings.ToUpper(name)))
			if err := os.Rename(path, newName); err != nil {
				log.Fatal(err)
			}
			// tranfer
			path, err = filepath.Abs(LabolinkHl7TranferExe)
			if err != nil {
				log.Fatal(err)
			}
			newName = strings.ReplaceAll(LabolinkHl7TranferExe, "TRANSFER", fmt.Sprintf("TRANSFER_%s", strings.ToUpper(name)))
			if err := os.Rename(path, newName); err != nil {
				log.Fatal(err)
			}
			// tranfer
			path, err = filepath.Abs(LabolinkHl7TranferWExe)
			if err != nil {
				log.Fatal(err)
			}
			newName = strings.ReplaceAll(LabolinkHl7TranferWExe, "TRANSFER", fmt.Sprintf("TRANSFER_%s", strings.ToUpper(name)))
			if err := os.Rename(path, newName); err != nil {
				log.Fatal(err)
			}

			// install bat
			content, err := os.ReadFile(InstallFile)
			if err != nil {
				log.Fatal(err)
			}

			text := string(content)
			text = strings.ReplaceAll(text, "RECIVER", fmt.Sprintf("RECIVER_%s", strings.ToUpper(name)))
			text = strings.ReplaceAll(text, "TRANSFER", fmt.Sprintf("TRANSFER_%s", strings.ToUpper(name)))
			if err := os.WriteFile(InstallFile, []byte(text), os.ModePerm); err != nil {
				log.Fatal(err)
			}

			// install tranfer bat
			content, err = os.ReadFile(InstallTranferFile)
			if err != nil {
				log.Fatal(err)
			}

			text = string(content)
			text = strings.ReplaceAll(text, "RECIVER", fmt.Sprintf("RECIVER_%s", strings.ToUpper(name)))
			text = strings.ReplaceAll(text, "TRANSFER", fmt.Sprintf("TRANSFER_%s", strings.ToUpper(name)))
			if err := os.WriteFile(InstallTranferFile, []byte(text), os.ModePerm); err != nil {
				log.Fatal(err)
			}
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the current version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("rift version:", Version)
		},
	}

	rootCmd.AddCommand(rename)
	rootCmd.AddCommand(versionCmd)
	rootCmd.Execute()
}
