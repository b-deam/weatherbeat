package beater

import (
	"fmt"
	"strings"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/b-deam/weatherbeat/config"
)

// weatherbeat configuration.
type weatherbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of weatherbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &weatherbeat{
		done:   make(chan struct{}),
		config: c,
	}

	if (strings.HasPrefix(bt.config.DataFeedUrl, "http://reg.bom.gov.au/fwo/") && strings.HasSuffix(bt.config.DataFeedUrl, ".json")) != true {
		return nil, fmt.Errorf("The `data.feed.url` must be a JSON URL from the BOM's Data Feeds (http://www.bom.gov.au/catalogue/data-feeds.shtml), instead we got: ", bt.config.DataFeedUrl)
	}

	return bt, nil
}

// Run starts weatherbeat.
func (bt *weatherbeat) Run(b *beat.Beat) error {
	logp.Info("weatherbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		weatherData, err := GetWeatherSummary(bt.config.DataFeedUrl)
		if err != nil {
			logp.Warn("Failed to collect data from: ", bt.config.DataFeedUrl, " and got: ", err)
		}

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"type":                   b.Info.Name,
				"notice.copyright":       weatherData.Observations.Notice[0].Copyright,
				"notice.copyright_url":   weatherData.Observations.Notice[0].CopyrightURL,
				"notice.disclaimer_url":  weatherData.Observations.Notice[0].DisclaimerURL,
				"notice.feedback_url":    weatherData.Observations.Notice[0].FeedbackURL,
				"header.refresh_message": weatherData.Observations.Header[0].RefreshMessage,
				"header.id":              weatherData.Observations.Header[0].ID,
				"header.main_id":         weatherData.Observations.Header[0].MainID,
				"header.name":            weatherData.Observations.Header[0].Name,
				"header.statetimezone":   weatherData.Observations.Header[0].StateTimeZone,
				"header.timezone":        weatherData.Observations.Header[0].TimeZone,
				"header.product_name":    weatherData.Observations.Header[0].ProductName,
				"header.state":           weatherData.Observations.Header[0].State,
				"sort_order":             weatherData.Observations.Data[0].SortOrder,
				"wmo":                    weatherData.Observations.Data[0].Wmo,
				"name":                   weatherData.Observations.Data[0].Name,
				"history_product":        weatherData.Observations.Data[0].HistoryProduct,
				"local_date_time":        weatherData.Observations.Data[0].LocalDateTime,
				"local_date_time_full":   weatherData.Observations.Data[0].LocalDateTimeFull,
				"aifs_time_utc":          weatherData.Observations.Data[0].AifstimeUtc,
				"lat":                    weatherData.Observations.Data[0].Lat,
				"lon":                    weatherData.Observations.Data[0].Lon,
				"apparent_t":             weatherData.Observations.Data[0].ApparentT,
				"cloud":                  weatherData.Observations.Data[0].Cloud,
				"cloud_base_m":           weatherData.Observations.Data[0].CloudBaseM,
				"cloud_oktas":            weatherData.Observations.Data[0].CloudOktas,
				"cloud_type_id":          weatherData.Observations.Data[0].CloudTypeID,
				"cloud_type":             weatherData.Observations.Data[0].CloudType,
				"delta_t":                weatherData.Observations.Data[0].DeltaT,
				"gust_kmh":               weatherData.Observations.Data[0].GustKmh,
				"gust_kt":                weatherData.Observations.Data[0].GustKt,
				"air_temp":               weatherData.Observations.Data[0].AirTemp,
				"dew_pt":                 weatherData.Observations.Data[0].Dewpt,
				"press":                  weatherData.Observations.Data[0].Press,
				"press_qnh":              weatherData.Observations.Data[0].PressQnh,
				"press_msl":              weatherData.Observations.Data[0].PressMsl,
				"press_tend":             weatherData.Observations.Data[0].PressTend,
				"rain_trace":             weatherData.Observations.Data[0].RainTrace,
				"rel_hum":                weatherData.Observations.Data[0].RelHum,
				"sea_state":              weatherData.Observations.Data[0].SeaState,
				"swell_dir_worded":       weatherData.Observations.Data[0].SwellDirWorded,
				"swell_height":           weatherData.Observations.Data[0].SwellHeight,
				"swell_period":           weatherData.Observations.Data[0].SwellPeriod,
				"vis_km":                 weatherData.Observations.Data[0].VisKm,
				"weather":                weatherData.Observations.Data[0].Weather,
				"wind_dir":               weatherData.Observations.Data[0].WindDir,
				"wind_spd_kmh":           weatherData.Observations.Data[0].WindSpdKmh,
				"windspd_kt":             weatherData.Observations.Data[0].WindSpdKt,
			},
		}
		bt.client.Publish(event)
		logp.Info("Event sent!")
	}
}

// Stop stops weatherbeat.
func (bt *weatherbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
