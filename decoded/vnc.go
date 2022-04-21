package decoded

type Vnc struct {
	ConnectionFailedReason string              `json:"connection_failed_reason"`
	DesktopName            string              `json:"desktop_name"`
	PixelEncoding          []Vnc_PixelEncoding `json:"pixel_encoding"`
	ScreenInfo             Vnc_ScreenInfo      `json:"screen_info"`
	SecurityTypes          []Vnc_SecurityTypes `json:"security_types"`
	Version                string              `json:"version"`
}

type Vnc_SecurityTypes struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Vnc_ScreenInfo struct {
	Height      uint64                     `json:"height"`
	NameLen     uint64                     `json:"name_len"`
	PixelFormat Vnc_ScreenInfo_PixelFormat `json:"pixel_format"`
	Width       uint64                     `json:"width"`
}

type Vnc_ScreenInfo_PixelFormat struct {
	BigEndian    bool   `json:"big_endian"`
	BitsPerPixel uint64 `json:"bits_per_pixel"`
	BlueMax      uint64 `json:"blue_max"`
	BlueShift    uint64 `json:"blue_shift"`
	Depth        uint64 `json:"depth"`
	GreenMax     uint64 `json:"green_max"`
	GreenShift   uint64 `json:"green_shift"`
	Padding1     uint64 `json:"padding1"`
	Padding2     uint64 `json:"padding2"`
	Padding3     uint64 `json:"padding3"`
	RedMax       uint64 `json:"red_max"`
	RedShift     uint64 `json:"red_shift"`
	TrueColor    bool   `json:"true_color"`
}

type Vnc_PixelEncoding struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
