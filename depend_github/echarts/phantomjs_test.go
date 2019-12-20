package echarts

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
)

//series结构体定义
type Series struct {
	Data          [][]interface{} `json:"data"`
	Name          string          `json:"name"`
	PointInterval int             `json:"pointInterval"`
}

//chart配置结构体定义
type ChartOption struct {
	Title struct {
		Margin int `json:"margin"`
		Style  struct {
			FontSize   string `json:"fontSize"`
			FontWeight string `json:"fontWeight"`
		} `json:"style"`
		Text string `json:"text"`
		X    int    `json:"x"`
	} `json:"title"`

	Chart struct {
		Type            string `json:"type"`
		BackgroundColor string `json:"backgroundColor"`
	} `json:"chart"`

	Credits struct {
		Enabled bool `json:"enabled"`
	} `json:"credits"`

	XAxis struct {
		Type                 string `json:"type"`
		DateTimeLabelFormats struct {
			Day string `json:"day"`
		} `json:"dateTimeLabelFormats"`
		TickInterval int `json:"tickInterval"`
	} `json:"xAxis"`
	YAxis struct {
		Labels struct {
			Format string `json:"format"`
			Style  struct {
				FontSize   string `json:"fontSize"`
				FontWeight string `json:"fontWeight"`
			} `json:"style"`
		} `json:"labels"`
		Title struct {
			Text string `json:"text"`
		} `json:"title"`
	} `json:"yAxis"`

	PlotOptions struct {
		Line struct {
			DataLabels struct {
				Enabled bool `json:"enabled"`
			} `json:"dataLabels"`
		} `json:"line"`
	} `json:"plotOptions"`

	Series []Series `json:"series"`

	Exporting struct {
		SourceWidth  int `json:"sourceWidth"`
		SourceHeight int `json:"sourceHeight"`
		Scale        int `json:"scale"`
	} `json:"exporting"`
}

type PieOption struct {
	Chart struct {
		BackgroundColor string `json:"backgroundColor"`
	} `json:"chart"`
	Colors  []string `json:"colors"`
	Credits struct {
		Enabled bool `json:"enabled"`
	} `json:"credits"`
	PlotOptions struct {
		Pie struct {
			DataLabels struct {
				Format string `json:"format"`
			} `json:"dataLabels"`
		} `json:"pie"`
	} `json:"plotOptions"`
	Series [1]struct {
		Data       [][]interface{} `json:"data"`
		DataLabels struct {
			Style struct {
				FontSize   string `json:"fontSize"`
				FontWeight string `json:"fontWeight"`
			} `json:"style"`
		} `json:"dataLabels"`
		Type string `json:"type"`
	} `json:"series"`
	Title struct {
		Margin int `json:"margin"`
		Style  struct {
			FontSize   string `json:"fontSize"`
			FontWeight string `json:"fontWeight"`
		} `json:"style"`
		Text string `json:"text"`
	} `json:"title"`
}

func NewChartOption() ChartOption {

	cht := ChartOption{}

	cht.Title.Margin = 30
	cht.Title.Style.FontSize = "18px"
	cht.Title.Style.FontWeight = "bold"
	cht.Title.X = -20

	cht.Chart.Type = "line"
	cht.Chart.BackgroundColor = "#f5f5f5"
	cht.Credits.Enabled = false

	cht.XAxis.Type = "datetime"
	cht.XAxis.DateTimeLabelFormats.Day = "%m月/%d日"
	cht.YAxis.Labels.Style.FontSize = "14px"
	cht.YAxis.Labels.Style.FontWeight = "bold"

	cht.PlotOptions.Line.DataLabels.Enabled = false

	cht.Exporting.Scale = 1
	cht.Exporting.SourceHeight = 400 //图片高度
	cht.Exporting.SourceWidth = 600  //图片宽度

	return cht
}

func NewPieOption() PieOption {
	pie := PieOption{}

	pie.Title.Margin = 30
	pie.Title.Style.FontSize = "18px"
	pie.Title.Style.FontWeight = "bold"

	pie.Credits.Enabled = false
	pie.Colors = []string{"#0067cc", "#30bfff", "#02fdff", "#4ad1d1", "#00b4cc", "#0193cd"} //饼图颜色
	pie.Chart.BackgroundColor = "#f5f5f5"                                                   //背景颜色
	pie.Series[0].Type = "pie"
	pie.Series[0].DataLabels.Style.FontSize = "14px"
	pie.Series[0].DataLabels.Style.FontWeight = "bold"

	return pie
}

// 雷达图
var chart1 = `{title:{text:'基础雷达图'},tooltip:{},legend:{width:10,right:10,icon:'circle',data:['预算分配（AllocatedBudget）','实际开销（ActualSpending）']},radar:{name:{textStyle:{color:'#999',borderRadius:3,padding:[3,5]}},indicator:[{name:'销售',max:6500},{name:'管理',max:16000},{name:'信息技术',max:30000},{name:'客服',max:38000},{name:'研发',max:52000},{name:'市场',max:25000}]},color:['#f1be2c','#6a3cf1'],series:[{name:'预算vs开销（Budgetvsspending）',type:'radar',areaStyle:{normal:{}},data:[{value:[4300,10000,28000,35000,50000,19000],name:'预算分配（AllocatedBudget）'},{value:[5000,14000,28000,31000,42000,21000],name:'实际开销（ActualSpending）'}]}]}`

// 有负数的图
var chart2 = `{grid:{top:0,bottom:0,left:45,right:5,height:220,width:'calc(100%-200px)',},xAxis:[{type:'value',max:12,min:-12,offset:2,maxInterval:2,splitLine:{lineStyle:{type:'dotted',color:'#ddd'}},axisLine:{lineStyle:{type:'dotted',color:'',width:0,},},axisLabel:{textStyle:{color:'#111',fontWeight:400,fontSize:12}},}],yAxis:{type:'category',axisLine:{show:false},axisTick:{show:false},offset:2,axisLabel:{textStyle:{color:'#111',fontWeight:400,fontSize:12}},splitLine:{show:false},data:['授权','教练','支持','命令'],},series:[{name:'生活费',type:'bar',stack:'总量',data:[{value:-4,label:{normal:{position:'right'}}},{value:-9,label:{normal:{position:'right'}}},{value:2,label:{normal:{position:'right'}}},{value:-2,label:{normal:{position:'right'}}},{value:6,label:{normal:{position:'right'}}},],itemStyle:{emphasis:{barBorderRadius:2},normal:{barBorderRadius:2,color:'#A674FF'}}}]}`

// 柱状图，高低风险
var chart3 = `{legend:{show:true,icon:'circle',data:['员工平均得分'],bottom:16,itemWidth:10,itemHeight:10},grid:{x:40,y:40,x2:60,y2:70},tooltip:{padding:[16,16],confine:true,trigger:'axis',backgroundColor:'#fff',textStyle:{color:'#111',fontSize:12},extraCssText:'box-shadow:0px4px12px0pxrgba(0,0,0,0.1);text-align:left;border-radius:8px;'},xAxis:{data:['分裂型','边缘型','反社会型','表演型','依赖型','偏执型','自恋型'],axisTick:{show:false},axisLine:{show:true,lineStyle:{color:'#eee'}},axisLabel:{interval:0,color:'#111',fontWeight:'bold'},axisPointer:{show:true,type:'shadow',label:{show:false},shadowStyle:{color:'rgba(52,76,93,0.03)'}}},yAxis:{max:10,axisLine:{show:false},axisTick:{lineStyle:{color:'yellow',type:'dashed',opacity:0.1}},axisLabel:{textStyle:{color:'#999'}},splitLine:{lineStyle:{color:'#eee',type:'dashed'}}},series:[{type:'bar',barWidth:28.4,name:'员工平均得分',itemStyle:{normal:{barBorderRadius:2,color:{x:0,y:0,x2:0,y2:1,type:'linear',global:false,colorStops:[{offset:0,color:'#A674FF'},{offset:1,color:'#593DFF'}]}}},data:['4.0','9.0','10.0','7.0','7.0','9.0','6.0'],markLine:{silent:true,symbol:'none',lineStyle:{type:'dashed'},data:[{lineStyle:{color:'#00A950'},label:{formatter:'低风险(<3分)',lineHeight:18},yAxis:3},{lineStyle:{color:'#F15656'},label:{formatter:'高风险(>8分)',lineHeight:18},yAxis:8}]}}]}`

// 进度条
var chart4 = `{tooltip:{trigger:'axis',axisPointer:{type:'shadow'}},grid:{left:'13%',top:'3%',height:20,width:300},xAxis:{type:'',show:false,},yAxis:{show:true,type:'category',data:['【合作精神】:8分'],axisLabel:{fontWeight:600},axisLine:{show:false},axisTick:{show:false}},series:[{type:'bar',data:[60],tooltip:{show:false},barMinHeight:30,barWidth:10,barMaxWidth:100,z:10,itemStyle:{normal:{barBorderRadius:20,color:'#FFC311',label:{show:false,position:'right',formatter:'{c}%',textStyle:{color:'#000'}}}}},{type:'bar',data:[100],tooltip:{show:false},barMinHeight:30,barWidth:10,barMaxWidth:100,barGap:'-100%',itemStyle:{normal:{barBorderRadius:20,color:'#CCCCCC',label:{show:false,position:'right',testStyle:{color:'#000'}}}}}]}`

// 和折线图
var chart5 = `{grid:{top:30,bottom:0,left:45,right:5,height:220,width:'calc(100%-200px)',},xAxis:[{type:'value',max:12,min:0,position:'top',offset:2,maxInterval:2,splitLine:{lineStyle:{type:'dotted',color:'#ddd'}},axisLine:{lineStyle:{type:'dotted',color:'',width:0,},},axisLabel:{textStyle:{color:'#111',fontWeight:400,fontSize:12}},}],yAxis:{type:'category',axisLine:{show:false},axisTick:{show:false},offset:2,axisLabel:{textStyle:{color:'#111',fontWeight:400,fontSize:12}},splitLine:{show:false},data:['授权','教练','支持','命令','你好'],},series:[{name:'生活费',type:'bar',stack:'总量',data:[{value:4,label:{show:true,position:'right',color:'#111'}},{value:9,label:{show:true,position:'right',color:'#111'}},{value:2,label:{show:true,position:'right',color:'#111'}},{value:2,label:{show:true,position:'right',color:'#111'}},{value:6,label:{show:true,position:'right',color:'#111'}},],itemStyle:{emphasis:{barBorderRadius:2},normal:{barBorderRadius:2,color:'#A674FF'},label:{show:true,position:'left',color:'#111'},},markLine:{silent:true,symbol:'none',lineStyle:{type:'dashed',},}}]}`

// /Users/kevin/Software/phantomjs-2.1.1-macosx/bin/phantomjs /Users/kevin/Software/echartsconvert/echarts-convert.js -t base64 -o "{xAxis:{type:'category',data:['Mon','Tue','Wed','Thu','Fri','Sat','Sun']},yAxis:{type:'value'},series:[{data:[820,932,901,934,1290,1330,1320],type:'line'}]}"
func TestPhantomjs(t *testing.T) {
	//f, err := os.Create("1.png")
	//defer os.Remove(f.Name())
	//defer f.Close()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	repBlance := strings.Replace(BrokenLineHistogramChart, " ", "", -1)
	chart := strings.Replace(repBlance, "\n", "", -1)
	cmd := exec.Command("/Users/kevin/Software/phantomjs-2.1.1-macosx/bin/phantomjs", "--debug=true", "/Users/kevin/Software/echartsconvert/echarts-convert.js", "-o", chart, "-t", "file", "-f", "/Users/kevin/go_mod/github/discover-func-golang/depend_github/echarts/img/1.png", "-w", "600", "-h", "300")
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
	_ = ioutil.WriteFile("/Users/kevin/go_mod/github/discover-func-golang/depend_github/echarts/img/" + "buf.png", dist, 0666)
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
