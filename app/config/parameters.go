package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ClusterConfiguration defines parameters required to connect to a given
// Elasticsearch server
type ClusterConfiguration struct {
	Name        string `json:"name,omitempty" yaml:"name"`
	Description string `json:"description,omitempty" yaml:"description"`
	URL         string `json:"url,omitempty" yaml:"url"`
	ClusterID   string `json:"cluster_id,omitempty" yaml:"cluster_id"`
	Username    string `json:"username,omitempty" yaml:"username"`
	Password    string `json:"password,omitempty" yaml:"password"`
}

// Configuration holds parameters required for the application ability to connect
// to Elasticsearch servers
type Configuration struct {
	Settings map[string]interface{} `json:"settings" yaml:"settings"`
	Clusters []ClusterConfiguration `json:"clusters" yaml:"clusters"`
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

// ClusterNames returns the names of all configured clusters
func (c Configuration) ClusterNames() []string {
	var names []string

	for _, n := range c.Clusters {
		names = append(names, n.Name)
	}

	return names
}

// Cluster will return the configuration for a given clusterName or an error if
// not found among the configured ones
func (c Configuration) Cluster(clusterName string) (ClusterConfiguration, error) {
	for _, n := range c.Clusters {
		if n.Name == clusterName {
			return n, nil
		}
	}

	return ClusterConfiguration{}, fmt.Errorf("requested configuraion for cluster %s was not found", clusterName)
}
