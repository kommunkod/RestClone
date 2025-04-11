package config

import (
	"encoding/json"
	"log"
	"maps"
	"os"
	"slices"
	"strings"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

var validExtensions = []string{"yaml", "yml", "json", "env"}

type Server struct {
	Listen    string `json:"listen" yaml:"listen" env:"RC_LISTEN" env-default:":8080" description:"Server listen address"`
	ListenTLS string `json:"listenTls" yaml:"listenTls" env:"RC_LISTEN_TLS" env-default:":8443" description:"Server TLS listen address"`

	TLS  Tls  `json:"tls" yaml:"tls" description:"TLS configuration"`
	Auth Auth `json:"auth" yaml:"auth" description:"Authentication configuration"`

	Timeout int `json:"timeout" yaml:"timeout" env:"RC_TIMEOUT" env-default:"30" description:"Server timeout in seconds"`

	ExtraEnv map[string]string `description:"Extra environment variables"`

	logger *log.Logger `json:"-" yaml:"-" description:"Logger"`
}

func Init() (*Server, error) {
	s := &Server{
		Listen:    ":8080",
		ListenTLS: ":8443",
		TLS:       Tls{},
		Auth:      Auth{},
		Timeout:   30,
		ExtraEnv:  make(map[string]string),
	}

	s.logger = log.Default()
	s.logger.SetPrefix("[RestClone] ")

	s.logger.Println("Initializing server configuration...")

	return s, s.Load()
}

func (s *Server) Println(logs ...any) {
	s.logger.Println(logs...)
}

func (s *Server) Printf(format string, logs ...any) {
	s.logger.Printf(format, logs...)
}

// Load configuration from environment variables, files, etc.
// Priority order: ENV > YAML/JSON > DEFAULT
func (s *Server) Load() error {
	configFiles := map[string][]byte{}

	for _, file := range os.Args[1:] {
		splitSuffix := strings.Split(strings.ToLower(file), ".")
		if len(splitSuffix) < 2 {
			s.logger.Printf("Skipping file %s: not a valid YAML/JSON file\n", file)
			continue
		}

		fileSuffix := splitSuffix[len(splitSuffix)-1]

		if fileSuffix == "env" {
			err := godotenv.Load(file)
			if err != nil {
				s.logger.Printf("Error loading .env file %s: %v\n", file, err)
			}
			continue
		}

		if slices.Contains(validExtensions, fileSuffix) {
			read, err := os.ReadFile(file)
			if err != nil {
				s.logger.Printf("Error reading file %s: %v\n", file, err)
				continue
			}

			configFiles[file] = read
		} else {
			s.logger.Printf("Skipping file %s: not a valid YAML/JSON/ENV file\n", file)
		}

	}

	extras, err := env.UnmarshalFromEnviron(s)
	if err != nil {
		return err
	}

	s.ExtraEnv = make(map[string]string)
	maps.Copy(s.ExtraEnv, extras)

	// Load from any files provided in args
	for file, content := range configFiles {
		if strings.HasSuffix(file, ".json") {
			if err := json.Unmarshal(content, s); err != nil {
				s.logger.Printf("Error unmarshalling JSON file %s: %v\n", file, err)
				continue
			}
		} else if strings.HasSuffix(file, ".yaml") || strings.HasSuffix(file, ".yml") {
			if err := yaml.Unmarshal(content, s); err != nil {
				s.logger.Printf("Error unmarshalling YAML file %s: %v\n", file, err)
				continue
			}
		}
	}

	return nil
}
