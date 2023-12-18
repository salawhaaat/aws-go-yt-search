package config

type Config struct {
	SearchAPI string
	ApiKey    string
}

var Cfg = Config{
	SearchAPI: "https://www.googleapis.com/youtube/v3/search",
	ApiKey:    "yourAPIKey",
}
