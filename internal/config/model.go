package config

// ServiceConfig represents the configuration needed to run this service
type ServiceConfig struct {
	SwordTaskManagerConfig *SwordTaskManager `yaml:"sword-task-manager"`
	RabbitMQConfig         *RabbitMQ         `yaml:"rabbitmq"`
	MySQLConfig            *MySQL            `yaml:"mysql"`
}

// SwordTaskManager represents the inner service configuration
type SwordTaskManager struct {
	Exchange    string            `yaml:"exchange"`
	RoutingKeys map[string]string `yaml:"routing_keys"`
	Type        string            `yaml:"type"`
	Durable     bool              `yaml:"durable"`
}

type RabbitMQ struct {
	URI               string `yaml:"uri"`
	HeartbeatInterval int    `yaml:"heartbeatInterval"`
}

type MySQL struct {
	URI      string `yaml:"uri"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
