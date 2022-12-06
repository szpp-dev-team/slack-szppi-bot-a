package commands

import (
	"log"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
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

	var options []slack.Block

	options = append(options, slack.NewSectionBlock(
		slack.NewTextBlockObject("mrkdwn", "令和最新版福袋っぴ！予算内で詰め込んだから買えっぴ！", false, false), nil, nil),
	)

	for _, row := range products {
		options = append(options, slack.NewSectionBlock(
			slack.NewTextBlockObject("mrkdwn", row.Name, false, false),
			[]*slack.TextBlockObject{
				slack.NewTextBlockObject("mrkdwn", humanize.Comma(int64(row.Price))+"円"+map[bool]string{true: "\nPrime配送", false: ""}[row.IsPrime], false, false),
			},
			slack.NewAccessory(
				slack.NewImageBlockElement(row.ThumbnailImageURL, row.Name),
			),
		))
		options = append(options, slack.NewDividerBlock())
	}
	_, _, _, err = o.c.SendMessage(slashCmd.ChannelID, slack.MsgOptionBlocks(options[:]...))

	return err
}
