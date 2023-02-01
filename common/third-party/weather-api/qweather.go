/**
 * @author:       wangxuebing
 * @fileName:     qweather.go
 * @date:         2023/2/1 14:47
 * @description:  和风天气API接口
 */

package weather_api

const (
	QWeatherGeoDomain           = "https://geoapi.qweather.com"
	QWeatherApiDomain           = "https://api.qweather.com"
	QWeatherApiDevDomain        = "https://devapi.qweather.com"
	QWeather_GeoLookup_Url      = "/v2/city/lookup"         //[GeoDomain]城市搜索URL
	QWeather_GeoTop_Url         = "/v2/city/top"            //[GeoDomain]热门城市查询URL
	QWeather_GeoPoi_Url         = "/v2/poi/lookup"          //[GeoDomain]POI搜索
	QWeather_GeoPoiRange_Url    = "/v2/poi/range"           //[GeoDomain]POI范围搜索
	QWeather_Now_Url            = "/v7/weather/now"         //[ApiDomain]实时天气
	QWeather_Day_Url            = "/v7/weather/{0}"         //[ApiDomain]每日天气预报[参数值: <3d:3天预报;7d:7天预报;10d:10天预报;15d:15天预报;30d:30天预报>]
	QWeather_Hour_Url           = "/v7/weather/{0}"         //[ApiDomain]逐小时天气预报[参数值: <24h:未来24小时;72h:未来72小时;168h:未来168小时>]
	QWeather_Indices_Url        = "/v7/indices/{0}"         //[ApiDomain]天气指数预报[参数值:<1d:当天生活指数;3d:未来3天生活指数>]
	QWeather_Air_Url            = "/v7/air/now"             //[ApiDomain]实时空气质量
	QWeather_DayAir_Url         = "/v7/air/5d"              //[ApiDomain]空气质量每日预报
	QWeather_SolarRadiation_Url = "/v7/solar-radiation/{0}" //[ApiDomain]太阳辐射逐小时预报[参数值: <24h:未来24小时;72h:未来72小时>]
	Qweather_Warning_Url        = "/v7/warning/now"         //[ApiDomain]天气灾害预警
	QWeather_WarningCity_Url    = "/v7/warning/list"        //[ApiDomain]天气预警城市列表
)

type (
	QWeatherCommonReferResp struct {
		Source  []string `json:"source"`  //原始数据来源，或数据源说明，可能为空
		License []string `json:"license"` //数据许可或版权声明，可能为空
	}
)

// GeoApi Params
type (
	// QWeatherGeoCityLookupResp 城市搜索Resp
	QWeatherGeoCityLookupResp struct {
		Code     string                         `json:"code"`
		Location []*QWeatherGeoCityLocationResp `json:"location"`
		Refer    QWeatherCommonReferResp        `json:"refer"`
	}
	// QWeatherGeoCityTopResp 热门城市查询Resp
	QWeatherGeoCityTopResp struct {
		Code        string                         `json:"code"`
		TopCityList []*QWeatherGeoCityLocationResp `json:"topCityList"`
		Refer       QWeatherCommonReferResp        `json:"refer"`
	}

	QWeatherGeoCityLocationResp struct {
		Name      string `json:"name"`      //地区/城市名称
		Id        string `json:"id"`        //地区/城市ID
		Lat       string `json:"lat"`       //地区/城市纬度
		Lon       string `json:"lon"`       //地区/城市经度
		Adm1      string `json:"adm1"`      //地区/城市所属一级行政区域
		Adm2      string `json:"adm2"`      //地区/城市的上级行政区划名称
		Country   string `json:"country"`   //地区/城市所属国家名称
		Tz        string `json:"tz"`        //地区/城市所在时区
		UtcOffset string `json:"utcOffset"` //地区/城市目前与UTC时间偏移的小时数
		IsDst     string `json:"isDst"`     //地区/城市是否当前处于夏令时。1 表示当前处于夏令时，0 表示当前不是夏令时
		Type      string `json:"type"`      //地区/城市的属性
		Rank      string `json:"rank"`      //地区评分
		FxLink    string `json:"fxLink"`    //该地区的天气预报网页链接，便于嵌入你的网站或应用
	}
)

// 城市天气
type (
	// QWeatherNowResp 实时天气Resp
	QWeatherNowResp struct {
		Code       string                  `json:"code"`
		UpdateTime string                  `json:"updateTime"` //当前API的最近更新时间
		FxLink     string                  `json:"fxLink"`     //当前数据的响应式页面，便于嵌入网站或应用
		Now        QWeatherRealDataResp    `json:"now"`
		Refer      QWeatherCommonReferResp `json:"refer"`
	}
	// QWeatherDailyResp 每日天气预报Resp
	QWeatherDailyResp struct {
		Code       string                   `json:"code"`
		UpdateTime string                   `json:"updateTime"` //当前API的最近更新时间
		FxLink     string                   `json:"fxLink"`     //当前数据的响应式页面，便于嵌入网站或应用
		Daily      []*QWeatherDailyDataResp `json:"daily"`
		Refer      QWeatherCommonReferResp  `json:"refer"`
	}
	// QWeatherHourResp 逐小时天气预报
	QWeatherHourResp struct {
		Code       string                   `json:"code"`
		UpdateTime string                   `json:"updateTime"` //当前API的最近更新时间
		FxLink     string                   `json:"fxLink"`     //当前数据的响应式页面，便于嵌入网站或应用
		Hourly     []*QWeatherDailyDataResp `json:"hourly"`
		Refer      QWeatherCommonReferResp  `json:"refer"`
	}

	QWeatherRealDataResp struct {
		ObsTime   string `json:"obs_time"`   //数据观测时间
		Temp      string `json:"temp"`       //温度，默认单位：摄氏度
		FeelsLike string `json:"feels_like"` //体感温度，默认单位：摄氏度
		Icon      string `json:"icon"`       //天气状况和图标的代码，图标可通过天气状况和图标下载
		Text      string `json:"text"`       //天气状况的文字描述，包括阴晴雨雪等天气状态的描述
		Wind360   string `json:"wind360"`    //风向360角度
		WindDir   string `json:"windDir"`    //风向
		WindScale string `json:"windScale"`  //风力等级
		WindSpeed string `json:"windSpeed"`  //风速，公里/小时
		Humidity  string `json:"humidity"`   //相对湿度，百分比数值
		Precip    string `json:"precip"`     //当前小时累计降水量，默认单位：毫米
		Pressure  string `json:"pressure"`   //大气压强，默认单位：百帕
		Vis       string `json:"vis"`        // 能见度，默认单位：公里
		Cloud     string `json:"cloud"`      //云量，百分比数值。可能为空
		Dew       string `json:"dew"`        //露点温度。可能为空
	}
	QWeatherDailyDataResp struct {
		FxDate         string `json:"fxDate"`         //预报日期
		Sunrise        string `json:"sunrise"`        //日出时间，在高纬度地区可能为空
		Sunset         string `json:"sunset"`         //日落时间，在高纬度地区可能为空
		Moonrise       string `json:"moonrise"`       //当天月升时间，可能为空
		Moonset        string `json:"moonset"`        //当天月落时间，可能为空
		MoonPhase      string `json:"moonPhase"`      //月相名称
		MoonPhaseIcon  string `json:"moonPhaseIcon"`  //月相图标代码，图标可通过天气状况和图标下载
		TempMax        string `json:"tempMax"`        //预报当天最高温度
		TempMin        string `json:"tempMin"`        //预报当天最低温度
		IconDay        string `json:"iconDay"`        //预报白天天气状况的图标代码，图标可通过天气状况和图标下载
		TextDay        string `json:"textDay"`        //预报白天天气状况文字描述，包括阴晴雨雪等天气状态的描述
		IconNight      string `json:"iconNight"`      //预报夜间天气状况的图标代码，图标可通过天气状况和图标下载
		TextNight      string `json:"textNight"`      //预报晚间天气状况文字描述，包括阴晴雨雪等天气状态的描述
		Wind360Day     string `json:"wind360Day"`     //预报白天风向360角度
		WindDirDay     string `json:"windDirDay"`     //预报白天风向
		WindScaleDay   string `json:"windScaleDay"`   //预报白天风力等级
		WindSpeedDay   string `json:"windSpeedDay"`   //预报白天风速，公里/小时
		Wind360Night   string `json:"wind360Night"`   //预报夜间风向360角度
		WindDirNight   string `json:"windDirNight"`   //预报夜间当天风向
		WindScaleNight string `json:"windScaleNight"` //预报夜间风力等级
		WindSpeedNight string `json:"windSpeedNight"` //预报夜间风速，公里/小时
		Humidity       string `json:"humidity"`       //相对湿度，百分比数值
		Precip         string `json:"precip"`         //预报当天总降水量，默认单位：毫米
		Pressure       string `json:"pressure"`       //大气压强，默认单位：百帕
		Vis            string `json:"vis"`            //能见度，默认单位：公里
		Cloud          string `json:"cloud"`          //云量，百分比数值。可能为空
		UvIndex        string `json:"uvIndex"`        //紫外线强度指数
	}
	QWeatherHourDataResp struct {
		FxTime    string `json:"fxTime"`    //预报时间
		Temp      string `json:"temp"`      //温度，默认单位：摄氏度
		Icon      string `json:"icon"`      //天气状况和图标的代码，图标可通过天气状况和图标下载
		Text      string `json:"text"`      //天气状况的文字描述，包括阴晴雨雪等天气状态的描述
		Wind360   string `json:"wind360"`   //风向360角度
		WindDir   string `json:"windDir"`   //风向
		WindScale string `json:"windScale"` //风力等级
		WindSpeed string `json:"windSpeed"` //风速，公里/小时
		Humidity  string `json:"humidity"`  //相对湿度，百分比数值
		Pop       string `json:"pop"`       //当前小时累计降水量，默认单位：毫米
		Precip    string `json:"precip"`    //逐小时预报降水概率，百分比数值，可能为空
		Pressure  string `json:"pressure"`  //大气压强，默认单位：百帕
		Cloud     string `json:"cloud"`     //云量，百分比数值。可能为空
		Dew       string `json:"dew"`       //露点温度。可能为空
	}
)

// 预警
type (
	// QWeatherWarningNowResp 天气灾害预警Resp
	QWeatherWarningNowResp struct {
		Code       string                  `json:"code"`
		UpdateTime string                  `json:"updateTime"` //当前API的最近更新时间
		FxLink     string                  `json:"fxLink"`     //当前数据的响应式页面，便于嵌入网站或应用
		Warning    QWeatherWarningDataResp `json:"warning"`
		Refer      QWeatherCommonReferResp `json:"refer"`
	}
	// QWeatherWarningLocListResp 天气预警城市列表Resp
	QWeatherWarningLocListResp struct {
		Code           string                        `json:"code"`
		UpdateTime     string                        `json:"updateTime"`
		WarningLocList []*QWeatherWarningLocDataResp `json:"warningLocList"`
	}

	QWeatherWarningDataResp struct {
		Id        string `json:"id"`        //本条预警的唯一标识，可判断本条预警是否已经存在
		Sender    string `json:"sender"`    //预警发布单位，可能为空
		PubTime   string `json:"pubTime"`   //预警发布时间
		Title     string `json:"title"`     //预警信息标题
		StartTime string `json:"startTime"` //预警开始时间，可能为空
		EndTime   string `json:"endTime"`   //预警结束时间，可能为空
		Status    string `json:"status"`    //预警信息的发布状态
		Level     string `json:"level"`     //预警等级（弃用），该字段已弃用，使用severity和severityColor代替
		Type      string `json:"type"`      //预警类型ID
		TypeName  string `json:"typeName"`  //预警类型名称
		Text      string `json:"text"`      //预警详细文字描述
		Related   string `json:"related"`   //与本条预警相关联的预警ID，当预警状态为cancel或update时返回。可能为空
	}

	QWeatherWarningLocDataResp struct {
		LocationId string `json:"locationId"`
	}
)

// 天气指数
type (
	// QWeatherIndicesResp 天气指数预报Resp
	QWeatherIndicesResp struct {
		Code       string                     `json:"code"`
		UpdateTime string                     `json:"updateTime"`
		FxLink     string                     `json:"fxLink"`
		Daily      []*QWeatherIndicesDataResp `json:"daily"`
		Refer      QWeatherCommonReferResp    `json:"refer"`
	}
	QWeatherIndicesDataResp struct {
		Date     string `json:"date"`     //预报日期
		Type     string `json:"type"`     //生活指数类型ID
		Name     string `json:"name"`     //生活指数类型的名称
		Level    string `json:"level"`    //生活指数预报等级
		Category string `json:"category"` //生活指数预报级别名称
		Text     string `json:"text"`     //生活指数预报的详细描述，可能为空
	}
)

// 空气质量
type (
	// QWeatherAirResp 实时空气质量
	QWeatherAirResp struct {
		Code       string                        `json:"code"`
		UpdateTime string                        `json:"updateTime"`
		Now        QWeatherAirNowDataResp        `json:"now"`
		Station    []*QWeatherAirStationDataResp `json:"station"`
		Refer      QWeatherCommonReferResp       `json:"refer"`
	}
	// QWeatherDailyAirResp 空气质量每日预报
	QWeatherDailyAirResp struct {
		Code       string                      `json:"code"`
		UpdateTime string                      `json:"updateTime"`
		Daily      []*QWeatherDailyAirDataResp `json:"daily"`
		Refer      QWeatherCommonReferResp     `json:"refer"`
	}
	QWeatherAirNowDataResp struct {
		PubTime  string `json:"pubTime"`  //空气质量数据发布时间
		Aqi      string `json:"aqi"`      //空气质量指数
		Level    string `json:"level"`    //空气质量指数等级
		Category string `json:"category"` //空气质量指数级别
		Primary  string `json:"primary"`  //空气质量的主要污染物，空气质量为优时，返回值为NA
		Pm10     string `json:"pm10"`     //PM10
		Pm2p5    string `json:"pm2p5"`    //PM2.5
		No2      string `json:"no2"`      //二氧化氮
		So2      string `json:"so2"`      //二氧化硫
		Co       string `json:"co"`       //一氧化碳
		O3       string `json:"o3"`       //臭氧
	}
	QWeatherAirStationDataResp struct {
		PubTime  string `json:"pubTime"`  //空气质量数据发布时间
		Id       string `json:"id"`       //监测站ID
		Name     string `json:"name"`     //监测站名称
		Aqi      string `json:"aqi"`      //空气质量指数
		Level    string `json:"level"`    //空气质量指数等级
		Category string `json:"category"` //空气质量指数级别
		Primary  string `json:"primary"`  //空气质量的主要污染物，空气质量为优时，返回值为NA
		Pm10     string `json:"pm10"`     //PM10
		Pm2p5    string `json:"pm2p5"`    //PM2.5
		No2      string `json:"no2"`      //二氧化氮
		So2      string `json:"so2"`      //二氧化硫
		Co       string `json:"co"`       //一氧化碳
		O3       string `json:"o3"`       //臭氧
	}

	QWeatherDailyAirDataResp struct {
		FxDate   string `json:"fxDate"`   //预报日期
		Aqi      string `json:"aqi"`      //空气质量指数
		Level    string `json:"level"`    //空气质量指数等级
		Category string `json:"category"` //空气质量指数级别
		Primary  string `json:"primary"`  //空气质量的主要污染物，空气质量为优时，返回值为NA
	}
)
