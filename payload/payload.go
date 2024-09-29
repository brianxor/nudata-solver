package payload

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/brianxor/nudata-solver/internal"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

type Website struct {
	ReqScriptUrl        string
	ReqScriptHeaders    map[string]string
	Version             string
	Mode                string
	IndentedJson        bool
	EncodeBase64Payload bool
}

var supportedWebsites = map[string]*Website{
	"kohls": {
		ReqScriptUrl: "https://fc.kohls.com//2.2/w/w-756138/init/js/",
		ReqScriptHeaders: map[string]string{
			"accept":          "*/*",
			"x-dynatrace":     "MT_3_2_15764307333791_4-0_02ae9e77-04af-45d9-9002-a561bb6830e4_14_13_99",
			"user-agent":      "NuDetectSDK/2.7.5 (iOS; iOS 18.0.0; en_US)",
			"accept-language": "en-GB,en;q=0.9",
			"accept-encoding": "gzip, deflate, br",
		},
		Version:             "2.7.5",
		Mode:                "LoginNew",
		IndentedJson:        true,
		EncodeBase64Payload: true,
	},
}

type Solver interface {
	Solve() (*Solution, error)
}

type Session struct {
	sessionId  string
	website    *Website
	httpClient *http.Client
}

func NewSolver(websiteName string, proxy string) (Solver, error) {
	if websiteName == "" {
		return nil, fmt.Errorf("website name is missing")
	}

	website, ok := supportedWebsites[websiteName]

	if !ok {
		return nil, fmt.Errorf("website name %s is not supported", websiteName)
	}

	var parsedProxy *url.URL

	var parseProxyErr error

	if proxy != "" {
		parsedProxy, parseProxyErr = url.Parse(proxy)

		if parseProxyErr != nil {
			return nil, parseProxyErr
		}
	}

	sessionUuid := internal.GenerateUuid(true)
	sessionTimestamp := time.Now().UnixMilli()

	sessionId := fmt.Sprintf("%s+%d", sessionUuid, sessionTimestamp)

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	if parsedProxy != nil {
		httpClient.Transport = &http.Transport{
			Proxy: http.ProxyURL(parsedProxy),
		}
	}

	session := &Session{
		sessionId:  sessionId,
		website:    website,
		httpClient: httpClient,
	}

	return session, nil
}

type NudataPayload struct {
	NdsPmd string `json:"nds-pmd"`
	Sid    string `json:"sid"`
}

type Solution struct {
	NdsPmd    string
	Sid       string
	SolveTime string
}

func (s *Session) Solve() (*Solution, error) {
	startTime := time.Now()

	initialPayload, err := s.initialPayload()

	if err != nil {
		return nil, err
	}

	widgetToken, err := s.getWidgetToken(initialPayload)

	if err != nil {
		return nil, err
	}

	finalPayload, err := s.finalPayload(widgetToken)

	if err != nil {
		return nil, err
	}

	solveTime := time.Since(startTime)

	nudataPayload := &NudataPayload{
		NdsPmd: finalPayload,
		Sid:    s.sessionId,
	}

	var nudataPayloadJson []byte
	var marshalErr error

	if s.website.IndentedJson {
		nudataPayloadJson, marshalErr = json.MarshalIndent(nudataPayload, "", "  ")

		if marshalErr != nil {
			return nil, marshalErr
		}
	} else {
		nudataPayloadJson, marshalErr = json.Marshal(nudataPayload)

		if marshalErr != nil {
			return nil, marshalErr
		}
	}

	nudataPayloadStr := string(nudataPayloadJson)

	if s.website.EncodeBase64Payload {
		nudataPayloadStr = base64.StdEncoding.EncodeToString([]byte(nudataPayloadStr))
	}

	solution := &Solution{
		NdsPmd:    nudataPayloadStr,
		Sid:       s.sessionId,
		SolveTime: solveTime.String(),
	}

	return solution, nil
}

func (s *Session) initialPayload() (string, error) {
	type payloadStructure struct {
		R   int    `json:"r"`
		Sid string `json:"sid"`
		Jsv string `json:"jsv"`
		Wpp int    `json:"wpp"`
		Ls  struct {
		} `json:"ls"`
		Wp string `json:"wp"`
	}

	initialPayload := &payloadStructure{
		R:   internal.GenerateRandomInt(1_000, 1_000_000),
		Sid: s.sessionId,
		Jsv: s.website.Version,
		Wpp: 1,
		Ls:  struct{}{},
		Wp:  s.website.Mode,
	}

	initialPayloadJson, err := json.Marshal(initialPayload)

	if err != nil {
		return "", err
	}

	initialPayloadStr := string(initialPayloadJson)

	return initialPayloadStr, nil
}

func (s *Session) finalPayload(widgetToken string) (string, error) {
	type widgetDataStructure struct {
		Mpmiv int      `json:"mpmiv"`
		Mhs   []string `json:"mhs"`
		Msc   string   `json:"msc"`
		Didtz int      `json:"didtz"`
		Miui  string   `json:"miui"`
		Mpi   string   `json:"mpi"`
		Mbmf  string   `json:"mbmf"`
		Mpmv  int      `json:"mpmv"`
		Mbm   string   `json:"mbm"`
		Mid   bool     `json:"mid"`
		Msm   int64    `json:"msm"`
		Ua    string   `json:"ua"`
		Mie   bool     `json:"mie"`
		Ic    string   `json:"ic"`
		Dit   bool     `json:"dit"`
		Wkr   int      `json:"wkr"`
		Sr    string   `json:"sr"`
		Mso   string   `json:"mso"`
		Mbb   string   `json:"mbb"`
		Mhbcs float64  `json:"mhbcs"`
		Midfv string   `json:"midfv"`
		Ipr   string   `json:"ipr"`
		Mbp   string   `json:"mbp"`
		Mul   string   `json:"mul"`
	}

	type payloadStructure struct {
		Sid        string               `json:"sid"`
		WidgetData *widgetDataStructure `json:"widgetData"`
		Wt         string               `json:"wt"`
	}

	mpmv := internal.GetRandomItem(iosMajorVersions)
	msm := internal.GenerateRandomInt(2_103_562_240, 12_103_562_240)
	wkr := internal.GenerateRandomInt(1_000, 1_000_000)
	sr := internal.GetRandomItem(iPhoneResolutions)
	mhbcs := internal.GenerateRandomFloat(0.25, 0.95)
	midpv := internal.GenerateUuid(true)

	widgetData := &widgetDataStructure{
		Mpmiv: defaultMpmiv,
		Mhs:   defaultMhs,
		Msc:   defaultMsc,
		Didtz: deviceTimezone,
		Miui:  deviceType,
		Mpi:   operatingSystem,
		Mbmf:  deviceManufacturer,
		Mpmv:  mpmv,
		Mbm:   deviceModel,
		Mid:   false,
		Msm:   int64(msm),
		Ua:    s.website.Version,
		Mie:   false,
		Ic:    defaultIc,
		Dit:   false,
		Wkr:   wkr,
		Sr:    sr,
		Mso:   defaultMso,
		Mbb:   deviceManufacturer,
		Mhbcs: mhbcs,
		Midfv: midpv,
		Ipr:   "",
		Mbp:   deviceModel,
		Mul:   userLocale,
	}

	finalPayload := &payloadStructure{
		Sid:        s.sessionId,
		WidgetData: widgetData,
		Wt:         widgetToken,
	}

	finalPayloadJson, err := json.Marshal(finalPayload)

	if err != nil {
		return "", err
	}

	finalPayloadStr := string(finalPayloadJson)

	encodedFinalPayload := internal.Rot13(finalPayloadStr)

	return encodedFinalPayload, nil
}

func (s *Session) getWidgetToken(initialPayload string) (string, error) {
	queryParams := url.Values{}
	queryParams.Set("q", initialPayload)

	req, err := http.NewRequest(http.MethodGet, s.website.ReqScriptUrl, nil)

	if err != nil {
		return "", err
	}

	for k, v := range s.website.ReqScriptHeaders {
		req.Header.Set(k, v)
	}

	req.URL.RawQuery = queryParams.Encode()

	resp, err := s.httpClient.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	var jsonBody []byte

	matches := ndwtiPattern.FindSubmatch(body)

	if len(matches) > 1 {
		jsonBody = matches[1]
	}

	type responseStructure struct {
		Wi string `json:"wi"`
		Co struct {
			UseNdx bool `json:"useNdx"`
		} `json:"co"`
		Wmd struct {
			Ipr struct {
				Fm []interface{} `json:"fm"`
				Lm bool          `json:"lm"`
				Tl int           `json:"tl"`
				Pd struct {
					Mn string `json:"mn"`
					Iq string `json:"iq"`
				} `json:"pd"`
				Il int `json:"il"`
			} `json:"ipr"`
			Wk struct {
				R string `json:"r"`
			} `json:"wk"`
			Di struct {
				Rt int `json:"rt"`
				Ut int `json:"ut"`
			} `json:"di"`
			Af []interface{} `json:"af"`
		} `json:"wmd"`
		Fd struct {
			Ipr string `json:"ipr"`
			Bi  string `json:"bi"`
			Wt  string `json:"wt"`
		} `json:"fd"`
		Gf []interface{} `json:"gf"`
	}

	var widgetTokenResponse responseStructure

	if unmarshalErr := json.Unmarshal(jsonBody, &widgetTokenResponse); unmarshalErr != nil {
		return "", unmarshalErr
	}

	widgetToken := widgetTokenResponse.Fd.Wt

	return widgetToken, nil
}

var ndwtiPattern = regexp.MustCompile(`ndwti\((.*?)\)`)
