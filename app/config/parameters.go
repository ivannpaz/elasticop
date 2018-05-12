package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ClusterConfiguration defines parameters required to connect to a given
// Elasticsearch server
type ClusterConfiguration struct {
	Description string `json:"description,omitempty" yaml:"description"`
	URL         string `json:"url,omitempty"`
	ClusterID   string `json:"cluster_id,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
}

// Configuration holds parameters required for the application ability to connect
// to Elasticsearch servers
type Configuration struct {
	Clusters []ClusterConfiguration
}

// Load a given yml file and create the Configuration representation
func Load(fileName string) Configuration {
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	configuration := Configuration{}

	if err = yaml.Unmarshal(contents, &configuration); err != nil {
		panic(err)
	}

	return configuration
}
