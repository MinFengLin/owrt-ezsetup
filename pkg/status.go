package pkg

import (
	"encoding/json"
	"os/exec"
	"time"
)

func (ubus_c *Ubus_call_sys_slice) ubus_call_sys(cs string) {
	var Ubus_sys_b []byte

	switch {
	case cs == "board":
		Ubus_sys_b, _ = exec.Command("ubus", "call", "system", cs).Output()
		json.Unmarshal([]byte(string(Ubus_sys_b)), &ubus_c.Board)
	case cs == "info":
		Ubus_sys_b, _ = exec.Command("ubus", "call", "system", cs).Output()
		json.Unmarshal([]byte(string(Ubus_sys_b)), &ubus_c.Info)
	}
}

func time_stamp(time_data int64) string {
	tm_u := time.Unix(time_data, 0)
	return tm_u.Format("2006-01-02 15:04:05")
}

func Parser_Status() Stauts_info_s {
	var Status Stauts_info_s
	var ubus_call_data = &Ubus_call_sys_slice{}

	ubus_call_data.ubus_call_sys("board")
	ubus_call_data.ubus_call_sys("info")

	Status.Hostname = ubus_call_data.Board.Hostname
	Status.Model = ubus_call_data.Board.Model
	Status.Firmware_Version = ubus_call_data.Board.Release.Description + " " + ubus_call_data.Board.Release.Revision
	Status.Kernel_Version = ubus_call_data.Board.Kernel
	Status.Local_Time = time_stamp(ubus_call_data.Info.Localtime)

	return Status
}

/*
 * Unit_Test
 */

func Parser_Status_Unit_Test(ubus_call_data *Ubus_call_sys_slice) Stauts_info_s {
	var Status Stauts_info_s

	ubus_call_data.ubus_call_sys("board")
	ubus_call_data.ubus_call_sys("info")

	Status.Hostname = ubus_call_data.Board.Hostname
	Status.Model = ubus_call_data.Board.Model
	Status.Firmware_Version = ubus_call_data.Board.Release.Description + " " + ubus_call_data.Board.Release.Revision
	Status.Kernel_Version = ubus_call_data.Board.Kernel
	Status.Local_Time = time_stamp(ubus_call_data.Info.Localtime)

	return Status
}
