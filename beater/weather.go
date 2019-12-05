package beater

import (
	"encoding/json"
	"net/http"
)

type WeatherSummary struct {
	Observations struct {
		Notice []struct {
			Copyright     string `json:"copyright"`
			CopyrightURL  string `json:"copyright_url"`
			DisclaimerURL string `json:"disclaimer_url"`
			FeedbackURL   string `json:"feedback_url"`
		} `json:"notice"`
		Header []struct {
			RefreshMessage string `json:"refresh_message"`
			ID             string `json:"ID"`
			MainID         string `json:"main_ID"`
			Name           string `json:"name"`
			StateTimeZone  string `json:"state_time_zone"`
			TimeZone       string `json:"time_zone"`
			ProductName    string `json:"product_name"`
			State          string `json:"state"`
		} `json:"header"`
		Data []struct {
			SortOrder         int         `json:"sort_order"`
			Wmo               int         `json:"wmo"`
			Name              string      `json:"name"`
			HistoryProduct    string      `json:"history_product"`
			LocalDateTime     string      `json:"local_date_time"`
			LocalDateTimeFull string      `json:"local_date_time_full"`
			AifstimeUtc       string      `json:"aifstime_utc"`
			Lat               float64     `json:"lat"`
			Lon               float64     `json:"lon"`
			ApparentT         float64     `json:"apparent_t"`
			Cloud             string      `json:"cloud"`
			CloudBaseM        int         `json:"cloud_base_m"`
			CloudOktas        int         `json:"cloud_oktas"`
			CloudTypeID       int         `json:"cloud_type_id"`
			CloudType         string      `json:"cloud_type"`
			DeltaT            float64     `json:"delta_t"`
			GustKmh           int         `json:"gust_kmh"`
			GustKt            int         `json:"gust_kt"`
			AirTemp           float64     `json:"air_temp"`
			Dewpt             float64     `json:"dewpt"`
			Press             float64     `json:"press"`
			PressQnh          float64     `json:"press_qnh"`
			PressMsl          float64     `json:"press_msl"`
			PressTend         string      `json:"press_tend"`
			RainTrace         string      `json:"rain_trace"`
			RelHum            int         `json:"rel_hum"`
			SeaState          string      `json:"sea_state"`
			SwellDirWorded    string      `json:"swell_dir_worded"`
			SwellHeight       interface{} `json:"swell_height"`
			SwellPeriod       interface{} `json:"swell_period"`
			VisKm             string      `json:"vis_km"`
			Weather           string      `json:"weather"`
			WindDir           string      `json:"wind_dir"`
			WindSpdKmh        int         `json:"wind_spd_kmh"`
			WindSpdKt         int         `json:"wind_spd_kt"`
		} `json:"data"`
	} `json:"observations"`
}

func GetWeatherSummary(url string) (WeatherSummary, error) {
	var weather WeatherSummary
	resp, err := http.Get(url)
	if err != nil {
		return weather, err
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&weather)
	return weather, nil
}
