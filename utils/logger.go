package utils
/*import (
	"human-resource-management-system/internal/config"
	"log/slog"
	"os"
	"path/filepath"
)

func initializeLogger(cfg *config.Config) *slog.Logger {
    logDir := filepath.Dir(cfg.LogFilePath)
    if err := os.MkdirAll(logDir, 0755); err != nil {
        panic("Failed to create log directory: " + err.Error())
    }

    file, err := os.OpenFile(cfg.LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic("Failed to open log file: " + err.Error())
    }

    consoleHandler := slog.NewTextHandler(os.Stdout, nil)
    fileHandler := slog.NewJSONHandler(file, nil)
    
}*/

