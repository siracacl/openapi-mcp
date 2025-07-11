package config

import "os"

// APIKeyLocation defines where the API key is to be placed in the request.
type APIKeyLocation string

const (
	APIKeyLocationHeader APIKeyLocation = "header"
	APIKeyLocationQuery  APIKeyLocation = "query"
	APIKeyLocationPath   APIKeyLocation = "path"
	APIKeyLocationCookie APIKeyLocation = "cookie"
)

// Config holds the application's configuration.
type Config struct {
	SpecPath          string
	APIKey            string
	APIKeyFromEnvVar  string
	APIKeyName        string
	APIKeyLocation    APIKeyLocation
	IncludeTags       []string
	ExcludeTags       []string
	IncludeOperations []string
	ExcludeOperations []string
	ServerBaseURL     string
	BaseURLHeader     string // New field for the dynamic base URL header
	BasicAuthHeader   string // New field for the dynamic basic auth header
	DefaultToolName   string
	DefaultToolDesc   string
	CustomHeaders     string
}

// GetAPIKey resolves the API key from direct value or environment variable.
func (c *Config) GetAPIKey() string {
	if c.APIKey != "" {
		return c.APIKey
	}
	if c.APIKeyFromEnvVar != "" {
		return os.Getenv(c.APIKeyFromEnvVar)
	}
	return ""
}
