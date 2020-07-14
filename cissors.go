package cissors

// Location describes location of the Rule in a CIS benchmark
type Location struct {
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
}

// Rule describes a CIS benchmark rule
type Rule struct {
	ID       string            `yaml:"id"`
	Name     string            `yaml:"name"`
	Scored   bool              `yaml:"scored"`
	Location []Location        `yaml:"location,omitempty"`
	Sections map[string]string `yaml:"-,inline"`
}
