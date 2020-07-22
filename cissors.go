package cissors

// Location describes location of the Rule in a CIS benchmark
type Location struct {
	ID   string `yaml:"id" json:"id"`
	Name string `yaml:"name" json:"name"`
}

// Rule describes a CIS benchmark rule
type Rule struct {
	ID       string            `yaml:"id" json:"id"`
	Name     string            `yaml:"name" json:"name"`
	Scored   bool              `yaml:"scored" json:"scored"`
	Location []Location        `yaml:"location,omitempty" json:"location,omitempty"`
	Sections map[string]string `yaml:"-,inline" json:"sections"`
}
