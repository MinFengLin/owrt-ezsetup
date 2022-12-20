package pkg

type Wifi_info struct {
	Wifi_device int
	Wifi_info_s Wifi_info_s
}

type Wifi_info_s struct {
	SSID    string
	PWD     string
	ENC     string
	CHANNEL string
	MODE    string
}
