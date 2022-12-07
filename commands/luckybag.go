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

	var blocks []slack.Block

	blocks = append(blocks, slack.NewSectionBlock(
		slack.NewTextBlockObject("mrkdwn", "令和最新版福袋っぴ！予算内で詰め込んだから買えっぴ！", false, false), nil, nil),
	)

	for _, row := range products {
		blocks = append(blocks, slack.NewSectionBlock(
			slack.NewTextBlockObject("mrkdwn", "<https://www.amazon.co.jp/dp/"+row.Asin+"|"+row.Name+">", false, false),
			[]*slack.TextBlockObject{
				slack.NewTextBlockObject("mrkdwn", humanize.Comma(int64(row.Price))+"円"+map[bool]string{true: " Prime配送", false: ""}[row.IsPrime], false, false),
			},
			slack.NewAccessory(
				slack.NewImageBlockElement(row.ThumbnailImageURL, row.Name),
			),
		))
		blocks = append(blocks, slack.NewDividerBlock())
	}
	_, _, _, err = o.c.SendMessage(slashCmd.ChannelID, slack.MsgOptionBlocks(blocks[:]...), slack.MsgOptionDisableLinkUnfurl(), slack.MsgOptionDisableMediaUnfurl())

	return err
}
