package pkg

type Stauts_info_s struct {
	Hostname         string
	Model            string
	Firmware_Version string
	Kernel_Version   string
	Local_Time       string
}

/*
 * ubus call system
 */

type Ubus_call_sys_slice struct {
	Board Board
	Info  Info
}

/*
 * ubus call system board
 */

type Board struct {
	Kernel   string  `json:"kernel"`
	Hostname string  `json:"hostname"`
	Model    string  `json:"model"`
	Release  Release `json:"release"`
}

type Release struct {
	Distribution string `json:"distribution"`
	Version      string `json:"version"`
	Revision     string `json:"revision"`
	Codename     string `json:"codename"`
	Target       string `json:"target"`
	Description  string `json:"description"`
}

/*
 * ubus call system info
 */

type Info struct {
	Uptime    int64    `json:"uptime"`
	Localtime int64    `json:"localtime"`
	Load      []string `json:"load"`
	Memory    Memory   `json:"memory"`
	Swap      Swap     `json:"swap"`
}

type Memory struct {
	Total    int64 `json:"total"`
	Free     int64 `json:"free"`
	Shared   int64 `json:"shared"`
	Buffered int64 `json:"buffered"`
}

type Swap struct {
	Total int64 `json:"total"`
	Free  int64 `json:"free"`
}
