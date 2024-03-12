package configs

type Config struct {
	Servers            map[string]Server `mapstructure:"servers"`
	JWTSecretKey       string            `mapstructure:"jwt_secret_key"`
	RecaptchaSecretKey string            `mapstructure:"recaptcha_secret_key"`
	APIUrl             string            `mapstructure:"api_url"`
	TelegramBot        TelegramBot       `mapstructure:"telegram_bot"`
	SwaggerPort        int               `mapstructure:"swagger_port"`
	TestMode           bool              `mapstructure:"test_mode"`
}

type Server struct {
	Address    string     `mapstructure:"address,omitempty"`
	PostgresDB PostgresDB `mapstructure:"postgres_db"`
	MongoDB    MongoDB    `mapstructure:"mongo_db"`
}

type PostgresDB struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

type MongoDB struct {
	Database string `mapstructure:"database"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Username string `mapstructure:"username"`
}

type TelegramBot struct {
	Enable bool   `mapstructure:"enable"`
	Token  string `mapstructure:"token"`
}
