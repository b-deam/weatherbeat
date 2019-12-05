// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period      time.Duration `config:"period"`
	DataFeedUrl string        `config:"data.feed.url"` // The Data Feed URL for your local observation station - all available here: http://www.bom.gov.au/catalogue/data-feeds.shtml
}

var DefaultConfig = Config{
	Period:      1800 * time.Second,                                       // Only runs once every 30mins, as the observations are only updated at this cadence
	DataFeedUrl: "http://reg.bom.gov.au/fwo/IDS60901/IDS60901.94672.json", // Default Data Feed URL is for Adelaide Airport :-)
}
