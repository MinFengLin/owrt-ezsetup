package pkg

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/digineo/go-uci"
)

func Wifi_device_count() (int, error) {
	interface_count, err := exec.Command("uci", "show", "|", "grep", "-ic", "wifi-device").Output()
	if err != nil {
		return -1, errors.New("WiFi interface counting error")
	}
	count, _ := strconv.Atoi(string(interface_count))

	return count, nil
}

func Call_uci_get(config, section, option string) []string {
	values, ok := uci.Get(config, section, option)
	if !ok {
		fmt.Printf("%s.%s.%s -> failed")
	}
	return values
}

func Parser_WiFi() bool {
	interface_count, err := Wifi_device_count()
	if err != nil {
		fmt.Println(err)
	}
	wifi_array := make([]Wifi_info, interface_count)
	wg := sync.WaitGroup{}
	wg.Add(interface_count)
	for index, value := range wifi_array {
		v := value
		v.Wifi_device = index
		go func(v *Wifi_info, wg *sync.WaitGroup) {
			v.Wifi_info_s.SSID = strings.Join(Call_uci_get("wireless", fmt.Sprintf("@wifi-iface[%d]", v.Wifi_device), "ssid"), "")
			v.Wifi_info_s.PWD = strings.Join(Call_uci_get("wireless", fmt.Sprintf("@wifi-iface[%d]", v.Wifi_device), "password"), "")
			v.Wifi_info_s.ENC = strings.Join(Call_uci_get("wireless", fmt.Sprintf("@wifi-iface[%d]", v.Wifi_device), "encryption"), "")
			v.Wifi_info_s.CHANNEL = strings.Join(Call_uci_get("wireless", fmt.Sprintf("wifi%d", v.Wifi_device), "channel"), "")
			v.Wifi_info_s.MODE = strings.Join(Call_uci_get("wireless", fmt.Sprintf("wifi%d", v.Wifi_device), "hwmode"), "")
			wg.Done()
		}(&v, &wg)
	}
	wg.Wait()
	return true
}

// wireless.wifi0=wifi-device
// wireless.wifi0.type='qcawificfg80211'
// wireless.wifi0.channel='auto'
// wireless.wifi0.macaddr='00:0b:6b:ee:86:de'
// wireless.wifi0.hwmode='11axa'
// wireless.wifi0.disabled='1'
// wireless.@wifi-iface[0]=wifi-iface
// wireless.@wifi-iface[0].device='wifi0'
// wireless.@wifi-iface[0].network='lan'
// wireless.@wifi-iface[0].mode='ap'
// wireless.@wifi-iface[0].ssid='OpenWrt'
// wireless.@wifi-iface[0].encryption='none'
// wireless.wifi1=wifi-device
// wireless.wifi1.type='qcawificfg80211'
// wireless.wifi1.channel='auto'
// wireless.wifi1.macaddr='00:0b:6b:ee:86:dd'
// wireless.wifi1.hwmode='11axg'
// wireless.wifi1.disabled='1'
// wireless.@wifi-iface[1]=wifi-iface
// wireless.@wifi-iface[1].device='wifi1'
// wireless.@wifi-iface[1].network='lan'
// wireless.@wifi-iface[1].mode='ap'
// wireless.@wifi-iface[1].ssid='OpenWrt'
// wireless.@wifi-iface[1].encryption='none'
