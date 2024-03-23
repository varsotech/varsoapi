package news

import (
	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
)

func Parse(url string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, errors.Wrapf(err, "failed getting and parsing rss feed")
	}

	return feed, nil
}
