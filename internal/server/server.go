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
	"github.com/gin-gonic/gin"
)

const pidFile = "/tmp/palsvm.pid"

var PalSvmServerConfig *config.PalsvmConfig

type Server struct {
	host    string
	port    string
	router  *gin.Engine
	srv     *http.Server
	running bool
	mu      sync.Mutex
}

func New(host, port string) *Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	return &Server{
		host:    host,
		port:    port,
		router:  router,
		running: false,
	}
}

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

	// loadServer config
	var err error
	PalSvmServerConfig, err = config.LoadPalsvmConfig()
	if err != nil {
		return err
	}

	fmt.Printf("PalsvmConfig: %+v\n", PalSvmServerConfig)

	// start Server
	go func() {
		s.srv.ErrorLog = log.New(io.Discard, "", 0)
		fmt.Printf("Server started on %s:%s (PID: %d)", s.host, s.port, os.Getpid())
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server error: %v", err)
		}
	}()

	s.running = true
	return nil
}

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

func (s *Server) Wait() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Server is running. Press Ctrl+C to stop")
	<-quit

	log.Println("Received shutdown signal")
}

func savePID() error {
	pid := os.Getpid()
	return os.WriteFile(pidFile, []byte(strconv.Itoa(pid)), 0644)
}

func removePID() {
	_ = os.Remove(pidFile)
}

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

func IsRunning() bool {
	pid, err := getPID()
	if err != nil {
		return false
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	err = process.Signal(syscall.Signal(0))
	return err == nil
}

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

	log.Printf("Sending SIGTERM to process %d", pid)

	if err := process.Signal(syscall.SIGTERM); err != nil {
		return fmt.Errorf("failed to send signal: %v", err)
	}

	time.Sleep(1 * time.Second)

	removePID()

	log.Println("Server stopped successfully")
	return nil
}
