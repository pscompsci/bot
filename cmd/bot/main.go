package main

import (
	"flag"

	"github.com/pscompsci/bot/internal/bot"
)

func main() {
	cfg := bot.Config{}

	flag.StringVar(&cfg.ApiKey, "apikey", "", "Kraken API Key")
	flag.StringVar(&cfg.SecretKey, "secret", "", "Kraken Secret Key")
	flag.Parse()

	cfg.Pairs = []string{"ADA/USD", "XBT/USD", "ETH/USD", "LTC/USD", "DOT/USD"}

	bot := bot.New(cfg)
	bot.Run()
}
