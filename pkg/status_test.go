package pkg

import (
	"reflect"
	"testing"
)

var (
	ubus_t *Ubus_call_sys_slice
)

func set_data() {
	ubus_t = &Ubus_call_sys_slice{
		Board: Board{
			Kernel:   "4.4.60",
			Hostname: "OpenWrt",
			Model:    "Qualcomm Technologies, Inc. IPQ807x/AP-HK02",
			Release: Release{
				Distribution: "OpenWrt",
				Version:      "OpenWrt",
				Revision:     "bc5ea22b2+r49254",
				Codename:     "chaos_calmer",
				Target:       "ipq/ipq807x_64",
				Description:  "OpenWrt Chaos Calmer 15.05.1",
			},
		},
		Info: Info{
			Uptime:    1000,
			Localtime: 1647961745,
			Load:      []string{"100", "200", "300"},
			Memory: Memory{
				Total:    1000,
				Free:     1000,
				Shared:   1000,
				Buffered: 1000,
			},
			Swap: Swap{
				Total: 1000,
				Free:  1000,
			},
		},
	}
}

func TestParser_Status(t *testing.T) {
	tests := []struct {
		name string
		want Stauts_info_s
	}{
		{
			name: "Test Parser Status",
			want: Stauts_info_s{
				Hostname:         "OpenWrt",
				Model:            "Qualcomm Technologies, Inc. IPQ807x/AP-HK02",
				Firmware_Version: "OpenWrt Chaos Calmer 15.05.1 bc5ea22b2+r49254",
				Kernel_Version:   "4.4.60",
				Local_Time:       "2022-03-22 15:09:05",
			},
		},
	}
	set_data()
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if got := Parser_Status_Unit_Test(ubus_t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser_Status_Unit_Test() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTestParser_Status(b *testing.B) {
	set_data()
	for i := 0; i < b.N; i++ {
		Parser_Status_Unit_Test(ubus_t)
	}
}
