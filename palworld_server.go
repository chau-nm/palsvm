package palsvm

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

// ServerConfig holds the configuration for Palworld server
type ServerConfig struct {
	ServerDir  string // Directory where server is installed
	ServerExec string // Server executable file name
	LogFile    string // Log file name for server output
	SteamCmd   string // Path to SteamCMD executable
	AppID      string // Steam App ID for Palworld
}

// DefaultConfig returns the default server configuration
func DefaultConfig() *ServerConfig {
	return &ServerConfig{
		ServerDir:  "/home/palworld/server",
		ServerExec: "./PalServer.sh",
		LogFile:    "server.log",
		SteamCmd:   "./steamcmd.sh",
		AppID:      "2394010",
	}
}

// ServerManager manages Palworld server operations
type ServerManager struct {
	config *ServerConfig
}

// NewServerManager creates a new ServerManager instance
// If config is nil, it uses DefaultConfig
func NewServerManager(config *ServerConfig) *ServerManager {
	if config == nil {
		config = DefaultConfig()
	}
	return &ServerManager{
		config: config,
	}
}

// runCommand is a helper function to execute shell commands
// with stdout and stderr redirected to os.Stdout and os.Stderr
func (sm *ServerManager) runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Start launches the Palworld server in background using nohup
// Returns error if server fails to start or log file cannot be created
func (sm *ServerManager) Start() error {
	fmt.Println("[Starting Palworld server...]")

	cmd := exec.Command("nohup",
		sm.config.ServerExec,
		"-useperfthreads",
		"-NoAsyncLoadingThread",
		"-UseMultithreadForDS",
	)

	cmd.Dir = sm.config.ServerDir

	// Redirect output to server.log
	logPath := fmt.Sprintf("%s/%s", sm.config.ServerDir, sm.config.LogFile)
	logfile, err := os.Create(logPath)
	if err != nil {
		return fmt.Errorf("failed to create log file: %w", err)
	}
	defer logfile.Close()

	cmd.Stdout = logfile
	cmd.Stderr = logfile

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	fmt.Printf("Server started with PID: %d\n", cmd.Process.Pid)
	return nil
}

// Stop terminates the running Palworld server using pkill
// Returns error if pkill command fails
func (sm *ServerManager) Stop() error {
	fmt.Println("[Stopping Palworld server...]")

	err := sm.runCommand("pkill", "PalServer")
	if err != nil {
		fmt.Println("Server was not running or pkill error occurred.")
		return err
	}

	fmt.Println("Server stopped.")
	return nil
}

// Restart stops and then starts the server with a 3-second delay
// Returns error if Start fails
func (sm *ServerManager) Restart() error {
	if err := sm.Stop(); err != nil {
		log.Printf("Warning: Stop returned error: %v", err)
	}

	fmt.Println("Waiting 3 seconds before restart...")
	time.Sleep(3 * time.Second)

	return sm.Start()
}

// Status checks if the server is running and returns its PID
// Returns: (isRunning bool, pid string, error)
func (sm *ServerManager) Status() (bool, string, error) {
	out, err := exec.Command("pidof", "PalServer").Output()

	if err != nil || len(out) == 0 {
		return false, "", nil
	}

	pid := strings.TrimSpace(string(out))
	return true, pid, nil
}

// PrintStatus prints the server status to stdout
// Shows whether server is running and its PID if available
func (sm *ServerManager) PrintStatus() {
	running, pid, err := sm.Status()
	if err != nil {
		fmt.Printf("Error checking status: %v\n", err)
		return
	}

	if running {
		fmt.Printf("Server is running with PID: %s\n", pid)
	} else {
		fmt.Println("Server is NOT running.")
	}
}

// Update updates the Palworld server via SteamCMD
// Downloads and validates the latest server files from Steam
// Returns error if SteamCMD command fails
func (sm *ServerManager) Update() error {
	fmt.Println("[Updating Palworld server via SteamCMD...]")

	err := sm.runCommand(sm.config.SteamCmd,
		"+login", "anonymous",
		"+app_update", sm.config.AppID, "validate",
		"+quit",
	)

	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	fmt.Println("Update completed.")
	return nil
}

// TailLog displays the server log in real-time using tail -f
// This is a blocking operation that continues until interrupted
// Returns error if tail command fails
func (sm *ServerManager) TailLog() error {
	fmt.Println("[Showing live log...]")

	logPath := fmt.Sprintf("%s/%s", sm.config.ServerDir, sm.config.LogFile)
	cmd := exec.Command("tail", "-f", logPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// GetLogPath returns the full path to the server log file
func (sm *ServerManager) GetLogPath() string {
	return fmt.Sprintf("%s/%s", sm.config.ServerDir, sm.config.LogFile)
}

// IsRunning is a convenience method that returns true if server is running
// This is a quick check without returning PID or errors
func (sm *ServerManager) IsRunning() bool {
	running, _, _ := sm.Status()
	return running
}
