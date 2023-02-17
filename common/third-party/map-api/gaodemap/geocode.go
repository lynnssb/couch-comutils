/**
 * @author:       wangxuebing
 * @fileName:     geocode.go
 * @date:         2023/2/10 13:16
 * @description:
 */

package gaodemap

type (
	CommonNtParam struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
	CommonLniParam struct {
		Location string `json:"location"`
		Name     string `json:"name"`
		Id       string `json:"id"`
	}
)

type (
	GeocodeReqParam struct {
		Key      string `json:"key"`
		Address  string `json:"address"`
		City     string `json:"city"`
		Sig      string `json:"sig"`
		Output   string `json:"output"`
		Callback string `json:"callback"`
	}
	GeocodesData struct {
		Country  string `json:"country"`
		Province string `json:"province"`
		City     string `json:"city"`
		CityCode string `json:"citycode"`
		District string `json:"district"`
		Street   string `json:"street"`
		Number   string `json:"number"`
		AdCode   string `json:"adcode"`
		Location string `json:"location"`
		Level    string `json:"level"`
	}
	GeocodeRespParam struct {
		Status   int          `json:"status"`
		Count    int          `json:"count"`
		Info     string       `json:"info"`
		Geocodes GeocodesData `json:"geocodes"`
	}
)

type (
	ReGeoReqParam struct {
		Key        string `json:"key"`
		Location   string `json:"location"`
		PoiType    string `json:"poitype"`
		Radius     int    `json:"radius"`
		Extensions string `json:"extensions"`
		RoadLevel  string `json:"roadlevel"`
		Sig        string `json:"sig"`
		Output     string `json:"output"`
		Homeorcorp string `json:"homeorcorp"`
	}
	ReGeoRespParam struct {
		Status    int           `json:"status"`
		Info      string        `json:"info"`
		InfoCode  string        `json:"infocode"`
		Regeocode ReGeoCodeData `json:"regeocode"`
	}
	ReGeoCodeData struct {
		FormattedAddress string                    `json:"formattedaddress"`
		AddressComponent ReGeoCodeAddressComponent `json:"address_component"`
	}
	ReGeoCodeAddressComponent struct {
		Country       string                `json:"country"`
		Province      string                `json:"province"`
		City          []string              `json:"city"`
		CityCode      string                `json:"citycode"`
		District      string                `json:"district"`
		AdCode        string                `json:"adcode"`
		TownShip      string                `json:"township"`
		TownCode      string                `json:"towncode"`
		Neighborhood  CommonNtParam         `json:"neighborhood"`
		Building      CommonNtParam         `json:"building"`
		StreetNumber  ReGeoCodeStreetNUmber `json:"streetnumber"`
		BusinessAreas []string              `json:"businessareas"`
		Posi          []string              `json:"posi"`
		Aois          []string              `json:"aois"`
		Roads         []string              `json:"roads"`
		Roadinters    []string              `json:"roadinters"`
	}
	ReGeoCodeStreetNUmber struct {
		Street    string `json:"street"`
		Number    string `json:"number"`
		Location  string `json:"location"`
		Direction string `json:"direction"`
		Distance  string `json:"distance"`
	}
)
