package weatherapi

var (
	Condition = map[string]string{
		"clear":                  "ясно",
		"partly-cloudy":          "малооблачно",
		"cloudy":                 "облачно с прояснениями",
		"overcast":               "пасмурно",
		"drizzle":                "морось",
		"light-rain":             "небольшой дождь",
		"rain":                   "дождь",
		"moderate-rain":          "умеренно сильный дождь",
		"heavy-rain":             "сильный дождь",
		"continuous-heavy-rain":  "длительный сильный дождь",
		"showers":                "ливень",
		"wet-snow":               "дождь со снегом",
		"light-snow":             "небольшой снег",
		"snow":                   "снег",
		"snow-showers":           "снегопад",
		"hail":                   "град",
		"thunderstorm":           "гроза",
		"thunderstorm-with-rain": "дождь с грозой",
		"thunderstorm-with-hail": "гроза с градом",
	}

	Daytime = map[string]string{"d": "светлое время суток", "n": "темное время суток"}

	Season = map[string]string{"summer": "лето", "autumn": "осень", "winter": "зима", "spring": "весна"}

	WindDir = map[string]string{
		"nw": "северо-западное",
		"n":  "северное",
		"ne": "северо-восточное",
		"e":  "восточное",
		"se": "юго-восточное",
		"s":  "южное",
		"sw": "юго-западное",
		"w":  "западное",
		"с":  "штиль",
	}

	MoonText = map[string]string{
		"moon-code-0":  "полнолуние",
		"moon-code-1":  "убывающая луна",
		"moon-code-2":  "убывающая луна",
		"moon-code-3":  "убывающая луна",
		"moon-code-4":  "последняя четверть",
		"moon-code-5":  "убывающая луна",
		"moon-code-6":  "убывающая луна",
		"moon-code-7":  "убывающая луна",
		"moon-code-8":  "новолуние",
		"moon-code-9":  "растущая луна",
		"moon-code-10": "растущая луна",
		"moon-code-11": "растущая луна",
		"moon-code-12": "первая четверть",
		"moon-code-13": "растущая луна",
		"moon-code-14": "растущая луна",
		"moon-code-15": "растущая луна",
	}

	PartName = map[string]string{
		"night":   "ночь",
		"morning": "утро",
		"day":     "день",
		"evening": "вечер",
	}
)

type WeatherQueryT struct {
	ServerTimeUnix uint32 `json:"now"`    // Время сервера в Unix формате
	ServerTimeDT   string `json:"now_dt"` // Время сервера в формате Data time
	Info           struct {
		Url       string  `json:"url"`
		Lattitude float32 `json:"lat"`
		Longitude float32 `json:"lon"`
	} `json:"info"` //
	Fact struct {
		TimeUnix     uint32  `json:"obs_time"`    // Время замера погодных данных в Unix формате
		Temp         int16   `json:"temp"`        // Температура в (С)
		TempFeelLike int16   `json:"feel_like"`   // Температура по ощущеметру
		Icon         string  `json:"icon"`        // Название иконки https://yastatic.net/weather/i/icons/funky/dark/< icon >.svg
		Condition   string  `json:"condition"`  // Код погодного описания
		WindSpeed   float32 `json:"wind_speed"`  // Скорость ветра (в м/с)
		WindGust    float32 `json:"wind_gust"`   // Скорость порывов ветра (в м/с)
		WindDir     string  `json:"wind_dir"`    // Направление ветра
		PressureMM  uint16  `json:"pressure_mm"` // Давление (в мм рт.ст.).
		PressurePas  uint16  `json:"pressure_pa"` // Давление (в гексопаскалях).
		Humidity     uint16  `json:"humidity_mm"` // Влажность воздуха
		Daytime      string  `json:"daytime"`     // Время суток
		Polar        bool    `json:"polar"`       // Признак того, что время суток, указанное в поле daytime является полярным.
		Season       string  `json:"season"`      // Время года в данном населенном пункте
	} `json:"fact"` //	Объект содержит информацию о погоде на данный момент
	Forecast struct {
		DateUnix uint32 `json:"date_ts"`   // Дата прогноза погоды в формате Unixtime
		Date     string `json:"date"`      // Дата прогноза в формате ГГГГ-ММ-ДД
		Week     uint16 `json:"week"`      // Порядковый номер недели
		Sunrise  string `json:"sunrise"`   // Время восхода Солнца, локальное время (может отсутствовать для полярныйх регионов)
		Sunset   string `json:"sunset"`    // Время заката Солнца, локальное время (может отсутствовать для полярныйх регионов)
		MoonCode uint8  `json:"moon_code"` // Код фазы луны
		MoonText string `json:"moon_text"` // Текстовый код фазы луны
		Parts    []struct {
			PartName    string  `json:"part_name"`   // Название времени суток
			TempMin     int16   `json:"temp_min"`    // Минимальная температура для времени суток
			TempMax     int16   `json:"temp_max"`    // Максимальная температура для времени суток
			TempAvg     int16   `json:"temp_avg"`    // Средняя температура для времени суток
			WindSpeed   float32 `json:"wind_speed"`  // Скорость ветра (в м/с)
			WindGust    float32 `json:"wind_gust"`   // Скорость порывов ветра (в м/с)
			WindDir     string  `json:"wind_dir"`    // Направление ветра
			PressureMM uint16  `json:"pressure_mm"` // Давление (в мм рт.ст.).
			PressurePa uint16  `json:"pressure_pa"` // Давление (в гексопаскалях).
			Humidity    uint16  `json:"humidity_mm"` // Влажность воздуха
			PrecMM     float32 `json:"prec_mm"`     // Прогнозируемое количество осадков
			PrecProb   uint16  `json:"prec_prob"`   // Вероятность выпадения осадков
			PrecPeriod uint16  `json:"prec_period"` // Прогнозируемый период осадков
			Icon        string  `json:"icon"`        // Название иконки https://yastatic.net/weather/i/icons/funky/dark/< icon >.svg
			Condition   string  `json:"conditions"`  // Код погодного описания
			FeelsLike   int16   `json:"feels_like"`  // Температура по ощущеметру для всемени суток
			Daytime     string  `json:"daytime"`     // Время суток
			Polar       bool    `json:"polar"`       // Признак того, что время суток, указанное в поле daytime является полярным.
		} `json:"parts"` // Прогнозы по временам суток
	} `json:"forecast"` //	Объект содержит данные прогноза погоды

}
