package models

type DockerCompose struct {
	Version  string             `json:"version" yaml:"version" mapstructure:"version,omitempty"`
	Services map[string]Service `json:"services" yaml:"services" mapstructure:"services,omitempty"`
	Networks map[string]Network `json:"networks" yaml:"networks,omitempty" mapstructure:"networks,omitempty"`
}

type Service struct {
	Image         string   `json:"image,omitempty" yaml:"image,omitempty" mapstructure:"image,omitempty"`
	Command       string   `json:"command,omitempty" yaml:"command,omitempty" mapstructure:"command,omitempty"`
	Build         string   `json:"build,omitempty" yaml:"build,omitempty" mapstructure:"build,omitempty"`
	ContainerName string   `json:"container_name,omitempty" yaml:"container_name,omitempty" mapstructure:"container_name,omitempty"`
	Ports         []string `json:"ports,omitempty" yaml:"ports,omitempty" mapstructure:"ports,omitempty"`
	Networks      []string `json:"networks,omitempty" yaml:"networks,omitempty" mapstructure:"networks,omitempty"`
}

type Network struct{}
