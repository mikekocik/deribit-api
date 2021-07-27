package deribit

import (
	"log"
	"strings"

	"github.com/KyberNetwork/deribit-api/models"
	jsoniter "github.com/json-iterator/go"
)

func (c *Client) subscriptionsProcess(event *Event) {
	if c.debugMode {
		log.Printf("channel: %s %v", event.Channel, string(event.Data))
	}
	switch {
	case event.Channel == "announcements":
		var notification models.AnnouncementsNotification
		err := jsoniter.Unmarshal(event.Data, &notification)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		c.Emit(event.Channel, &notification)
	case strings.HasPrefix(event.Channel, "book"):
		count := strings.Count(event.Channel, ".")
		if count == 2 {
			// book.BTC-PERPETUAL.raw, book.BTC-PERPETUAL.100ms
			if strings.HasSuffix(event.Channel, ".raw") {
				var notification models.OrderBookRawNotification
				err := jsoniter.Unmarshal(event.Data, &notification)
				if err != nil {
					log.Printf("%v", err)
					return
				}
				c.Emit(event.Channel, &notification)
			} else {
				var notification models.OrderBookNotification
				err := jsoniter.Unmarshal(event.Data, &notification)
				if err != nil {
					log.Printf("%v", err)
					return
				}
				c.Emit(event.Channel, &notification)
			}
		} else if count == 4 {
			// book.ETH-PERPETUAL.100.1.100ms, ...
			var notification models.OrderBookGroupNotification
			err := jsoniter.Unmarshal(event.Data, &notification)
			if err != nil {
				log.Printf("%v", err)
				return
			}
			c.Emit(event.Channel, &notification)
		}
	case strings.HasPrefix(event.Channel, "deribit_price_index"):
		var notification models.DeribitPriceIndexNotification
		err := jsoniter.Unmarshal(event.Data, &notification)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		c.Emit(event.Channel, &notification)
	case strings.HasPrefix(event.Channel, "deribit_price_ranking"):
		var notification models.DeribitPriceRankingNotification
		err := jsoniter.Unmarshal(event.Data, &notification)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		c.Emit(event.Channel, &notification)
	case strings.HasPrefix(event.Channel, "estimated_expiration_price"):
		var notification models.EstimatedExpirationPriceNotification
		err := jsoniter.Unmarshal(event.Data, &notification)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		c.Emit(event.Channel, &notification)
	case strings.HasPrefix(event.Channel, "markprice.options"):
		var notification models.MarkpriceOptionsNotification
		err := jsoniter.Unmarshal(event.Data, &notification)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		c.Emit(event.Channel, &notification)
	case strings.HasPrefix(event.Channel, "perpetual"):
		var notification models.PerpetualNotification
		err := jsoniter.Unmarshal(event.Data, &notification)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		c.Emit(event.Channel, &notification)

	case strings.HasPrefix(event.Channel, "quote"):
		var notification models.QuoteNotification
		err := jsoniter.Unmarshal(event.Data, &notification)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		c.Emit(event.Channel, &notification)
	case strings.HasPrefix(event.Channel, "ticker"):
		var notification models.TickerNotification
		err := jsoniter.Unmarshal(event.Data, &notification)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		c.Emit(event.Channel, &notification)
	case strings.HasPrefix(event.Channel, "trades"):
		count := strings.Count(event.Channel, ".")
		if count == 2 {
			// trades.BTC-PERPETUAL.raw
			var notification models.TradesNotification
			err := jsoniter.Unmarshal(event.Data, &notification)
			if err != nil {
				log.Printf("%v", err)
				return
			}
			c.Emit(event.Channel, &notification)
		}
		if count == 3 {
			// trades.future.ETH.100ms
			var notification models.TradesKindCurrencyNotification
			err := jsoniter.Unmarshal(event.Data, &notification)
			if err != nil {
				log.Printf("%v", err)
				return
			}
			c.Emit(event.Channel, &notification)
		}
	case strings.HasPrefix(event.Channel, "user.changes"):
		var notification models.UserChangesNotification
		err := jsoniter.Unmarshal(event.Data, &notification)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		c.Emit(event.Channel, &notification)
	case strings.HasPrefix(event.Channel, "user.orders"):
		if strings.HasSuffix(event.Channel, ".raw") {
			var notification models.Order
			err := jsoniter.Unmarshal(event.Data, &notification)
			if err != nil {
				log.Printf("%v", err)
				return
			}
			c.Emit(event.Channel, &notification)
		} else {
			var notification models.UserOrderNotification
			err := jsoniter.Unmarshal(event.Data, &notification)
			if err != nil {
				log.Printf("%v", err)
				return
			}
			c.Emit(event.Channel, &notification)
		}
	case strings.HasPrefix(event.Channel, "user.portfolio"):
		var notification models.PortfolioNotification
		err := jsoniter.Unmarshal(event.Data, &notification)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		c.Emit(event.Channel, &notification)
	case strings.HasPrefix(event.Channel, "user.trades"):
		var notification models.UserTradesNotification
		err := jsoniter.Unmarshal(event.Data, &notification)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		c.Emit(event.Channel, &notification)
	default:
		log.Printf("%v", string(event.Data))
	}
}
