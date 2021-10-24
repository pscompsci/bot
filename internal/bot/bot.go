package bot

import (
	"fmt"

	"github.com/pscompsci/bot/internal/exchange"
	"github.com/pscompsci/eventbus"

	_ "github.com/lib/pq"
)

type bot struct {
	config   Config
	bus      *eventbus.EventBus
	exchange *exchange.Kraken
}

func New(cfg Config) (*bot, error) {
	bus := *eventbus.New()
	return &bot{
		config:   cfg,
		exchange: exchange.NewKraken(cfg.ApiKey, cfg.SecretKey, &bus),
		bus:      &bus,
	}, nil
}

func (b *bot) Run() {
	krakenChannel := make(chan eventbus.DataEvent)
	b.bus.Subscribe("kraken", krakenChannel)

	go b.exchange.CollectCandles(b.config.Pairs, 1)

	for {
		data := <-krakenChannel
		fmt.Println(data)
	}
}
