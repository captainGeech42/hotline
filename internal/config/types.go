package config

type Http struct {
	Domain         string `yaml:"domain"`
	Port           int    `yaml:"port"`
	DefaultReponse string `yaml:"default_reponse"`
}

type Dns struct {
	Port               int    `yaml:"port"`
	DefaultAResponse   string `yaml:"default_A_response"`
	DefaultTXTResponse string `yaml:"default_TXT_response"`
}

type Callback struct {
	Http Http `yaml:"http"`
	Dns  Dns  `yaml:"dns"`
}

type App struct {
	Port int `yaml:"port"`
}

type Server struct {
	Callback Callback `yaml:"callback"`
	App      App      `yaml:"app"`
}

type Client struct {
	ServerUrl string `yaml:"server_url"`
}

type Config struct {
	Server Server `yaml:"server"`
	Client Client `yaml:"client"`
}
