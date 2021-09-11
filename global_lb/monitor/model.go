package monitor

// MonitorTemplateDataIcmp icmp 健康检查模板
type MonitorTemplateDataIcmp struct {
	Interval           int64  `json:"interval"`
	Timeout            int64  `json:"timeout"`
	ProbeInterval      int64  `json:"probe_interval"`
	ProbeTimeout       int64  `json:"probe_timeout"`
	ProbeAttempts      int64  `json:"probe_attempts"`
	IgnoreDownResponse string `json:"ignore_down_response"`
	Transparent        bool   `json:"transparent"`
}

// MonitorTemplateDataHttp http 健康检查模板
type MonitorTemplateDataHttp struct {
	Interval           int64  `json:"interval"`
	Timeout            int64  `json:"timeout"`
	ProbeTimeout       int64  `json:"probe_timeout"`
	IgnoreDownResponse string `json:"ignore_down_response"`
	SendString         string `json:"send_string"`
	ReceiveString      string `json:"receive_string"`
	UserName           string `json:"user_name"`
	Password           string `json:"password"`
	Reverse            bool   `json:"reverse"`
	Transparent        bool   `json:"transparent"`
}

// MonitorTemplateDataHttps https 健康检查模板
type MonitorTemplateDataHttps struct {
	Interval           int64  `json:"interval"`
	Timeout            int64  `json:"timeout"`
	ProbeTimeout       int64  `json:"probe_timeout"`
	IgnoreDownResponse string `json:"ignore_down_response"`
	SendString         string `json:"send_string"`
	ReceiveString      string `json:"receive_string"`
	CipherList         string `json:"cipher_list"`
	UserName           string `json:"user_name"`
	Password           string `json:"password"`
	Compatibility      string `json:"compatibility"` // 暂无
	ClientCertificate  string `json:"client_certificate"`
	ClientKey          string `json:"client_key"`
	Reverse            bool   `json:"reverse"`
	Transparent        bool   `json:"transparent"`
}

// MonitorTemplateDataTcp tcp 健康检查模板
type MonitorTemplateDataTcp struct {
	Interval           int64  `json:"interval"`
	Timeout            int64  `json:"timeout"`
	ProbeTimeout       int64  `json:"probe_timeout"`
	IgnoreDownResponse string `json:"ignore_down_response"`
	SendString         string `json:"send_string"`
	ReceiveString      string `json:"receive_string"`
	Reverse            bool   `json:"reverse"`
	Transparent        bool   `json:"transparent"`
}

// tcp syn 健康检查模板(tcp half open)
type MonitorTemplateDataTcpSyn struct {
	Interval           int64  `json:"interval"`
	Timeout            int64  `json:"timeout"`
	ProbeInterval      int64  `json:"probe_interval"`
	ProbeTimeout       int64  `json:"probe_timeout"`
	ProbeAttempts      int64  `json:"probe_attempts"`
	IgnoreDownResponse string `json:"ignore_down_response"`
	Transparent        bool   `json:"transparent"`
}

// MonitorTemplateDataUdp udp 健康检查模板
type MonitorTemplateDataUdp struct {
	Interval           int64  `json:"interval"`
	Timeout            int64  `json:"timeout"`
	ProbeInterval      int64  `json:"probe_interval"`
	ProbeTimeout       int64  `json:"probe_timeout"`
	ProbeAttempts      int64  `json:"probe_attempts"`
	IgnoreDownResponse string `json:"ignore_down_response"`
	SendString         string `json:"send_string"`
	Transparent        bool   `json:"transparent"`
	Debug              string `json:"debug"`
}
