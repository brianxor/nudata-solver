package payload

var defaultMhs = []string{"mimt"}

const (
	defaultMpmiv = 0
	defaultMsc   = "--"
	defaultMso   = "--"
	defaultIc    = "0,no;"
)
const (
	deviceTimezone     = -180
	deviceManufacturer = "Apple"
	deviceModel        = "iPhone"
	userLocale         = "en-US"
	deviceType         = "phone"
	operatingSystem    = "ios"
)

var iosMajorVersions = []int{
	15,
	16,
	17,
	18,
}

var iPhoneResolutions = []string{
	"1334x750",  // iPhone 6
	"1920x1080", // iPhone 6 Plus
	"1334x750",  // iPhone 6s
	"1920x1080", // iPhone 6s Plus
	"1334x750",  // iPhone 7
	"1920x1080", // iPhone 7 Plus
	"1334x750",  // iPhone 8
	"1920x1080", // iPhone 8 Plus
	"2436x1125", // iPhone X
	"1792x828",  // iPhone XR
	"2436x1125", // iPhone XS
	"2688x1242", // iPhone XS Max
	"1792x828",  // iPhone 11
	"2436x1125", // iPhone 11 Pro
	"2688x1242", // iPhone 11 Pro Max
	"2532x1170", // iPhone 12
	"2340x1080", // iPhone 12 mini
	"2532x1170", // iPhone 12 Pro
	"2778x1284", // iPhone 12 Pro Max
	"2532x1170", // iPhone 13
	"2340x1080", // iPhone 13 mini
	"2532x1170", // iPhone 13 Pro
	"2778x1284", // iPhone 13 Pro Max
	"2532x1170", // iPhone 14
	"2778x1284", // iPhone 14 Plus
	"2556x1179", // iPhone 14 Pro
	"2796x1290", // iPhone 14 Pro Max
	"2556x1179", // iPhone 15
	"2796x1290", // iPhone 15 Plus
	"2556x1179", // iPhone 15 Pro
	"2796x1290", // iPhone 15 Pro Max
	"2556x1179", // iPhone 16
	"2796x1290", // iPhone 16 Plus
	"2622x1206", // iPhone 16 Pro
	"2868x1320", // iPhone 16 Pro Max
}
