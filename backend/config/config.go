package config

import (
	"github.com/martinisecurity/trusty/pkg/gserver"
)

const (
	// WFEServerName specifies server name for Web Front End
	WFEServerName = "wfe"
	// CISServerName specifies server name for Certificate Information
	CISServerName = "cis"
	// CAServerName specifies server name for Certification Authority
	CAServerName = "ca"
)

// Configuration contains the user configurable data for the service
type Configuration struct {

	// Region specifies the Region / Datacenter where the instance is running
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Environment specifies the environment where the instance is running: prod|stage|dev
	Environment string `json:"environment,omitempty" yaml:"environment,omitempty"`

	// ServiceName specifies the service name to be used in logs, metrics, etc
	ServiceName string `json:"service,omitempty" yaml:"service,omitempty"`

	// ClusterName specifies the cluster name
	ClusterName string `json:"cluster,omitempty" yaml:"cluster,omitempty"`

	// Metrics specifies the metrics pipeline configuration
	Metrics Metrics `json:"metrics" yaml:"metrics"`

	// Audit contains configuration for the audit logger
	Audit Logger `json:"audit" yaml:"audit"`

	// Logs contains configuration for the logger
	Logs Logger `json:"logs" yaml:"logs"`

	// LogLevels specifies the log levels per package
	LogLevels []RepoLogLevel `json:"log_levels" yaml:"log_levels"`

	// CryptoProv specifies the configuration for crypto providers
	CryptoProv CryptoProv `json:"crypto_provider" yaml:"crypto_provider"`

	// CaSQL specifies the configuration for SQL provider
	CaSQL SQL `json:"ca_sql" yaml:"ca_sql"`

	// JWT specifies configuration file for the JWT provider
	JWT string `json:"jwt_provider" yaml:"jwt_provider"`

	// Authority specifies configuration file for CA
	Authority string `json:"authority" yaml:"authority"`

	// RegistrationAuthority contains configuration info for RA
	RegistrationAuthority *RegistrationAuthority `json:"ra" yaml:"ra"`

	// HTTPServers specifies a list of servers that expose HTTP or gRPC services
	HTTPServers map[string]*gserver.HTTPServerCfg `json:"servers" yaml:"servers"`

	// TODO: refactor
	// TrustyClient specifies configurations for the client to connect to the cluster
	TrustyClient TrustyClient `json:"trusty_client" yaml:"trusty_client"`

	// Tasks specifies array of tasks
	Tasks []Task `json:"tasks" yaml:"tasks"`

	// CertsMonitor specifies the configuration for cert monitor
	CertsMonitor CertsMonitor `json:"certs_monitor" yaml:"certs_monitor"`
}

// CryptoProv specifies the configuration for crypto providers
type CryptoProv struct {

	// Default specifies the location of the configuration file for default provider
	Default string `json:"default,omitempty" yaml:"default,omitempty"`

	// Providers specifies the list of locations of the configuration files
	Providers []string `json:"providers,omitempty" yaml:"providers,omitempty"`

	// PKCS11Manufacturers specifies the list of supported manufactures of PKCS11 tokens
	PKCS11Manufacturers []string `json:"pkcs11_manufacturers,omitempty" yaml:"pkcs11_manufacturers,omitempty"`
}

// CertsMonitor specifies configurations for monitoring certs expiry
type CertsMonitor struct {

	// Locations specifies the list of files to monitor. It may have a prefix with cert type, in format {type}:{location}
	Locations []string `json:"locations" yaml:"locations"`
}

// Task specifies configuration of a single task.
type Task struct {

	// Name specifies the name of the task.
	Name string `json:"name" yaml:"name"`

	// Schedule specifies the schedule of this task.
	Schedule string `json:"schedule" yaml:"schedule"`

	// Args specifies parameters for the task.
	Args []string `json:"args" yaml:"args"`
}
