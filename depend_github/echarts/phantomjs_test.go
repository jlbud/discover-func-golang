package echarts

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
)

var (
	PhantomJs        = "/Users/kevin/Software/phantomjs-2.1.1-macosx/bin/phantomjs"
	EchartsConvertJs = "/Users/kevin/Software/echartsconvert/echarts-convert.js"
	TempImgRootPath  = "/Users/kevin/go_mod/github/discover-func-golang/depend_github/echarts/img"
)

func GetChart(chartIndex int) {
	switch chartIndex {
	case RadarChartIndex:
	case NegativeHistogramChartIndex:
	case RiskHistogramChartIndex:
	case ProgressBarChartIndex:
	case BrokenLineHistogramChartIndex:
	}
}

func TestPhantomjs(t *testing.T) {

	repBlance := strings.Replace(RiskHistogramChartOption, " ", "", -1)
	chart := strings.Replace(repBlance, "\n", "", -1)
	cmd := exec.Command(PhantomJs, "--debug=true", EchartsConvertJs, "-o", chart, "-t", "file", "-f", TempImgRootPath+"/1.png", "-w", "900", "-h", "300")
	cmd.Stdout = os.Stdout
	cmd.Run()

	decodePNG("1.png")
}

func decodePNG(name string) []byte {
	fBuf, _ := ioutil.ReadFile("/Users/kevin/go_mod/github/discover-func-golang/depend_github/echarts/img/" + name)
	// 数据缓存
	bufStore := make([]byte, 5000000)
	// 文件转base64
	base64.StdEncoding.Encode(bufStore, fBuf)
	// 解压
	dist, _ := base64.StdEncoding.DecodeString(string(bufStore))
	_ = ioutil.WriteFile("/Users/kevin/go_mod/github/discover-func-golang/depend_github/echarts/img/"+"buf.png", dist, 0666)
	return bufStore
}

func decodePNG1(name string) string {
	ff, _ := os.Open("/Users/kevin/go_mod/github/discover-func-golang/depend_github/img/" + name)
	defer ff.Close()
	sourcebuffer := make([]byte, 500000)
	n, _ := ff.Read(sourcebuffer)
	//base64压缩
	sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])
	return sourcestring
	//解压
	//dist, _ := base64.StdEncoding.DecodeString(string(sourcestring))
	//写入新文件
	//f, _ := os.OpenFile("./2.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	//defer f.Close()
	//f.Write(dist)
}

type AdviceData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	//Desc    []string `json:"desc,omitempty"`
	DescStr string `json:"desc,omitempty"`
}

type StaffLabels struct {
	StaffId              int      `json:"staff_id"`
	JobChoiceValueLabels []string `json:"job_choice_value_labels,omitempty"`
	KeyExperienceLabels  []string `json:"key_experience_labels,omitempty"`
	Labels               []string `json:"labels"`
}

func TestAA(t *testing.T) {
	data := &AdviceData{
		//Desc:    []string{"1", "2"},
		DescStr: "11",
	}
	b, _ := json.Marshal(data)
	t.Log(string(b))
}

func TestBB(t *testing.T) {
	data := &StaffLabels{
		Labels: []string{},
	}
	b, _ := json.Marshal(data)
	t.Log(string(b))
}
