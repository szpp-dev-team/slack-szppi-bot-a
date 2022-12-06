package model

type Product struct { // amazonの商品ひとつの情報
	Asin              string
	Name              string
	Price             int
	IsPrime           bool
	ThumbnailImageURL string
}

type ScrapeResult struct {
	Products []Product
}
