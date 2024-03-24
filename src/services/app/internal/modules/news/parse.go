package news

import (
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
)

const (
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"
)

func Parse(url string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	fp.UserAgent = userAgent

	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, errors.Wrapf(err, "failed getting and parsing rss feed")
	}

	return feed, nil
}
