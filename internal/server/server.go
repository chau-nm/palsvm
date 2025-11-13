package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/chau-nm/palsvm/internal/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const (
	pidFile     = "/tmp/palsvm.pid"
	sessionName = "palsvm"
)

// Server represents the HTTP server instance with route handling and lifecycle management
type Server struct {
	host    string
	port    string
	router  *gin.Engine
	srv     *http.Server
	running bool
	mu      sync.Mutex
}

// New creates and initializes a new Server instance with the specified host and port.
// It sets up Gin in release mode, configures sessions, and loads templates.
func New(host, port string) *Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	setupSession(router)
	loadTemplates(router)

	return &Server{
		host:    host,
		port:    port,
		router:  router,
		running: false,
	}
}

// setupSession configures session management using cookie-based storage.
// Sessions are configured with 2-hour expiry, HttpOnly flag, and Lax SameSite policy.
func setupSession(router *gin.Engine) {
	store := cookie.NewStore([]byte(config.PalSvmServerConfig.ServerSecret))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   7200, // 2 hours in seconds
		HttpOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
	})
	router.Use(sessions.Sessions(sessionName, store))
}

// loadTemplates configures static file serving and loads HTML templates.
// Static assets are served from ./templates/assets, ./templates/css, and ./templates/js.
// HTML templates are loaded from ./templates/*.tmpl.
func loadTemplates(router *gin.Engine) {
	// Serve static files
	router.Static("/assets", "./templates/assets")
	router.Static("/css", "./templates/css")
	router.Static("/js", "./templates/js")

	// Load HTML templates
	router.LoadHTMLGlob("templates/*.tmpl")
}

// Start starts the HTTP server if not already running.
// It checks for existing server instances, saves the process PID, and starts listening.
// Returns an error if server is already running or if another instance exists.
func (s *Server) Start() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.running {
		return fmt.Errorf("server is already running")
	}

	if IsRunning() {
		return fmt.Errorf("another server instance is already running")
	}

	s.setupRoutes()

	s.srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.host, s.port),
		Handler: s.router,
	}

	if err := savePID(); err != nil {
		return fmt.Errorf("failed to save PID: %v", err)
	}

	// Start server in a goroutine to avoid blocking
	go func() {
		s.srv.ErrorLog = log.New(io.Discard, "", 0)
		fmt.Printf("Server started on %s:%s (PID: %d)\n\n", s.host, s.port, os.Getpid())
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	s.running = true
	return nil
}

// Stop gracefully shuts down the server with a 5-second timeout.
// It removes the PID file after successful shutdown.
// Returns an error if server is not running or shutdown fails.
func (s *Server) Stop() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.running {
		return fmt.Errorf("server is not running")
	}

	log.Println("Stopping server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %v", err)
	}

	s.running = false

	removePID()

	log.Println("Server stopped")
	return nil
}

// Wait blocks until an interrupt signal (SIGINT or SIGTERM) is received.
// Used to keep the server running until manual termination.
func (s *Server) Wait() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Server is running. Press Ctrl+C to stop")
	<-quit

	log.Println("Received shutdown signal")
}

// savePID writes the current process ID to the PID file.
// Returns an error if file write fails.
func savePID() error {
	pid := os.Getpid()
	return os.WriteFile(pidFile, []byte(strconv.Itoa(pid)), 0644)
}

// removePID deletes the PID file.
// Errors are ignored as this is a cleanup operation.
func removePID() {
	_ = os.Remove(pidFile)
}

// getPID reads and parses the process ID from the PID file.
// Returns the PID and an error if file reading or parsing fails.
func getPID() (int, error) {
	data, err := os.ReadFile(pidFile)
	if err != nil {
		return 0, err
	}

	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, fmt.Errorf("invalid PID in file: %v", err)
	}

	return pid, nil
}

// IsRunning checks if a server instance is currently running by verifying the PID file
// and sending a signal(0) to the process to test if it exists.
// Returns true if a server process is active, false otherwise.
func IsRunning() bool {
	pid, err := getPID()
	if err != nil {
		return false
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	// Send signal 0 to check if process exists without actually sending a signal
	err = process.Signal(syscall.Signal(0))
	return err == nil
}

// StopRunningServer stops any currently running server instance by sending SIGTERM.
// It reads the PID from file, sends termination signal, waits briefly, then cleans up.
// Returns an error if no server is running or if termination fails.
func StopRunningServer() error {
	if !IsRunning() {
		return fmt.Errorf("no server is running")
	}

	pid, err := getPID()
	if err != nil {
		return fmt.Errorf("failed to get PID: %v", err)
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("failed to find process: %v", err)
	}

	log.Printf("Sending SIGTERM to process %d\n", pid)

	if err := process.Signal(syscall.SIGTERM); err != nil {
		return fmt.Errorf("failed to send signal: %v", err)
	}

	// Wait for process to terminate
	time.Sleep(1 * time.Second)

	removePID()

	log.Println("Server stopped successfully")
	return nil
}
