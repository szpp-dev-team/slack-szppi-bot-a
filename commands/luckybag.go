package commands

import (
	"log"
	"strconv"
	"strings"

	"github.com/slack-go/slack"
	"github.com/szpp-dev-team/szpp-slack-bot/luckyBag"
)

type SubHandlerLuckyBag struct {
	c *slack.Client
}

func NewSubHandlerLuckyBag(c *slack.Client) *SubHandlerLuckyBag {
	return &SubHandlerLuckyBag{c}
}

func (o *SubHandlerLuckyBag) Name() string {
	return "lucky-bag"
}

func (o *SubHandlerLuckyBag) Handle(slashCmd *slack.SlashCommand) error {
	q := strings.Join(strings.Fields(slashCmd.Text)[1:], "")
	log.Println(q)
	price, err := strconv.Atoi(q)
	if err != nil {
		return err
	}

	products, err := luckyBag.MakeLuckyBag(price)
	if err != nil {
		return err
	}

	resp := "令和最新版福袋っぴ！予算内で詰め込んだから買えっぴ！\n"
	for _, row := range products {
		resp += row.Name[:30]
		if len(row.Name) > 30 {
			resp += "..."
		}
		resp += "\n"
		resp += strconv.Itoa(row.Price) + "\n"
		resp += strconv.Itoa(row.Price) + "円"
		if row.IsPrime {
			resp += "(Prime配送)"
		}
		resp += "\n"
	}

	_, _, _, err = o.c.SendMessage(slashCmd.ChannelID, slack.MsgOptionText(resp, false))

	return err
}
