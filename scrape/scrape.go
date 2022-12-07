package scrape

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/szpp-dev-team/szpp-slack-bot/model"

	"github.com/antchfx/htmlquery"
)

func Scrape() (model.ScrapeResult, error) {
	var results model.ScrapeResult
	url := "https://www.amazon.co.jp/s?k=%E4%BB%A4%E5%92%8C%E6%9C%80%E6%96%B0%E7%89%88&s=date-desc-rank"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return results, err
	}
	req.Header.Add("authority", "www.amazon.co.jp")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("accept-language", "ja")
	// req.Header.Add("cookie", "session-id=356-5905431-3456106; session-id-time=2082787201l; i18n-prefs=JPY; skin=noskin; ubid-acbjp=358-9165170-4915538; csm-hit=tb:E65Z1D1XHXN6DGJ7WT8S+s-E65Z1D1XHXN6DGJ7WT8S|1669978695005&t:1669978695005&adb:adblk_no; session-token=\"Uf3FTLGOg7tQymR8EETwCWtUAScBnBdVr2K62hP6mEMpcCJ+spHKeMhUFhg6g9aIDP7Hk+AEgoL36bcGwjgFWzdNE2x52HPlzsq3IVfbftSxcsOMlG3/ebf8L6AstS0IXqXbDyDO8Pz/92AkXlsrXJnJiUainZjWB5EE5CLJXt3NbFmSPwQvwm2frKtRENachtu5OLH+Hhq0+Uic+1bQJu0pZx1k7JfR2B7h4Bdmr90=\"; i18n-prefs=JPY; session-id=356-9384586-5935046; session-id-time=2082787201l; ubid-acbjp=355-5252746-6495124")
	req.Header.Add("device-memory", "8")
	req.Header.Add("downlink", "10")
	req.Header.Add("dpr", "2")
	req.Header.Add("ect", "4g")
	req.Header.Add("referer", "https://www.amazon.co.jp/")
	req.Header.Add("rtt", "100")
	req.Header.Add("sec-ch-device-memory", "8")
	req.Header.Add("sec-ch-dpr", "2")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"107\", \"Chromium\";v=\"107\", \"Not=A?Brand\";v=\"24\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-ch-ua-platform-version", "\"13.0.1\"")
	req.Header.Add("sec-ch-viewport-width", "494")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	req.Header.Add("viewport-width", "494")

	res, err := client.Do(req)
	if err != nil {
		return results, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return results, err
	}

	doc, err := htmlquery.Parse(strings.NewReader(string(body)))
	if err != nil {
		return results, err
	}

	list := htmlquery.Find(doc, "//div[@data-asin]")
	for _, n := range list {
		if htmlquery.SelectAttr(n, "data-asin") == "" || len(htmlquery.Find(n, "//img[@class='s-image']")) == 0 || len(htmlquery.Find(n, "//h2[@class='a-size-mini a-spacing-none a-color-base s-line-clamp-4']")) == 0 || len(htmlquery.Find(n, "//span[@class='a-price-whole']")) == 0 {
			continue
		}
		asin := htmlquery.SelectAttr(n, "data-asin")
		image_url := htmlquery.SelectAttr(htmlquery.Find(n, "//img[@class='s-image']")[0], "src")
		title := htmlquery.InnerText(htmlquery.Find(n, "//h2[@class='a-size-mini a-spacing-none a-color-base s-line-clamp-4']")[0])
		price, _ := strconv.Atoi(strings.Replace(strings.Replace(htmlquery.InnerText(htmlquery.Find(n, "//span[@class='a-price-whole']")[0]), "￥", "", -1), ",", "", -1))
		is_prime := len(htmlquery.Find(n, "//i[@aria-label='Amazon プライム']")) == 1
		results.Products = append(results.Products, model.Product{Asin: asin, Name: title, Price: price, IsPrime: is_prime, ThumbnailImageURL: image_url})
	}
	return results, nil
}
