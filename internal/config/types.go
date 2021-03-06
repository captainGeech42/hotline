package config

type Http struct {
	Port           int    `yaml:"port"`
	DefaultReponse string `yaml:"default_response"`
}

type Dns struct {
	Port               int    `yaml:"port"`
	DefaultAResponse   string `yaml:"default_A_response"`
	DefaultTXTResponse string `yaml:"default_TXT_response"`
	AcmeChallengePath  string `yaml:"acme_challenge_path"`
}

type Callback struct {
	Domain string `yaml:"domain"`
	Http   Http   `yaml:"http"`
	Dns    Dns    `yaml:"dns"`
}

type Web struct {
	Port int `yaml:"port"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
}

type Server struct {
	Callback Callback `yaml:"callback"`
	Web      Web      `yaml:"web"`
	Database Database `yaml:"db"`
}

type Client struct {
	ServerUrl string `yaml:"server_url"`
}

type Config struct {
	Server Server `yaml:"server"`
	Client Client `yaml:"client"`
}
