package exchange

import (
	"github.com/pscompsci/bot/pkg/models"
	"github.com/pscompsci/eventbus"

	ws "github.com/aopoltorzhicky/go_kraken/websocket"
)

type Kraken struct {
	apiKey    string
	secretKey string
	bus       *eventbus.EventBus
}

func NewKraken(api, secret string, bus *eventbus.EventBus) *Kraken {
	return &Kraken{
		apiKey:    api,
		secretKey: secret,
		bus:       bus,
	}
}

func (k *Kraken) CollectCandles(pairs []string, interval int64) error {
	kraken := ws.NewKraken(ws.ProdBaseURL)
	if err := kraken.Connect(); err != nil {
		return err
	}

	if err := kraken.SubscribeCandles(pairs, interval); err != nil {
		return err
	}

	for {
		update := <-kraken.Listen()
		switch message := update.Data.(type) {
		case ws.Candle:
			k.bus.Publish("kraken", models.Candle{
				Pair:    update.Pair,
				EndTime: message.EndTime,
				Open:    message.Open,
				High:    message.High,
				Low:     message.Low,
				Close:   message.Close,
				Volume:  message.Volume,
			})
		}
	}
}
