package main

import (
	"fmt"
	"log"
	"os"

	logs "mailscheduler/logger"
	"mailscheduler/routes"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Mailscheduler struct {
		Port string `yaml:"Port"`
	} `yaml:"mailscheduler"`
}

func NewConfig(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %v", err)
	}
	defer f.Close()
	var c Config
	err = yaml.NewDecoder(f).Decode(&c)
	if err != nil {
		return nil, fmt.Errorf("error decoding config file from yaml to go struct: %v", err)
	}
	// Ensure the port has the correct format
	if c.Mailscheduler.Port == "" {
		c.Mailscheduler.Port = "8080" // Default port if missing
	} else if c.Mailscheduler.Port[0] != ':' {
		c.Mailscheduler.Port = ":" + c.Mailscheduler.Port
	}
	return &c, nil
}

type Server struct {
	router *gin.Engine
	config *Config
}

func NewServer(config *Config) *Server {
	r := gin.Default()
	return &Server{
		router: r,
		config: config,
	}
}

func (s *Server) Start() {
	fmt.Println("Starting server on", s.config.Mailscheduler.Port)
	s.router.Run(s.config.Mailscheduler.Port)
}

func main() {
	newConfig, err := NewConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log, err := logs.NewLogger("logs/app.log")
	if err != nil {
		panic("Couldn't initialize logging")
	}
	// Inject dependencies into the server
	srv := NewServer(newConfig)
	routes.RegisterRoutes(srv.router, log)
	log.Info("server started")
	srv.Start()
}
