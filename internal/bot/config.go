package bot

type Database struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Server       string `json:"server"`
	Database     string `json:"database"`
	MaxOpenConns int    `json:"maxOpenConns"`
	MaxIdleConns int    `json:"maxIdleConns"`
	MaxIdleTime  string `json:"maxIdleTime"`
}

type Config struct {
	ApiKey    string   `json:"apiKey"`
	SecretKey string   `json:"secretKey"`
	Pairs     []string `json:"pairs"`
	DB        Database `json:"database"`
}
