package model

type Product struct { // amazonの商品ひとつの情報
	Name              string
	Price             int
	isPrime           bool
	ThumbnailImageURL string
}
