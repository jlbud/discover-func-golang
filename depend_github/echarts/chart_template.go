package echarts

// 1
// 雷达图
var RadarChart = `{
    tooltip: {},
    legend: {
        width:10,
        right:10,
        icon:'circle',
        data: ['预算分配（Allocated Budget）', '实际开销（Actual Spending）']
    },
    radar: {
        name: {
            textStyle: {
                color: '#999',
                borderRadius: 3,
                padding: [3, 5]
           }
        },
        indicator: [
           { name: '销售', max: 6500},
           { name: '管理', max: 16000},
           { name: '信息技术', max: 30000},
           { name: '客服', max: 38000},
           { name: '研发', max: 52000},
           { name: '市场', max: 25000}
        ]
    },
    color:['#f1be2c','#6a3cf1'],
    series: [{
        name: '预算 vs 开销（Budget vs spending）',
        type: 'radar',
        areaStyle: {normal: {
            
        }},
        data : [
            {
                value : [4300, 10000, 28000, 35000, 50000, 19000],
                name : '预算分配（Allocated Budget）'
            },
             {
                value : [5000, 14000, 28000, 31000, 42000, 21000],
                name : '实际开销（Actual Spending）'
            }
        ]
    }]
}`

// 2
// 负数
// 柱状图
// "-w", "600", "-h", "250"
var NegativeHistogramChart = `{
      grid: {
        top: 0,
        bottom: 0,
        left: 45,
        right: 5,
        height: 220,
        width: 'calc(100% - 200px)',
      },
   xAxis: [
        {
          type: 'value',
          max: 12,
          min: -12,
          offset: 2,
          maxInterval: 2,
          splitLine: {
            lineStyle: {
              type: 'dotted',
              color: '#ddd'
            }
          },
          axisLine: {
            lineStyle: {
              type: 'dotted',
              color: '',
              width: 0,
            },
          },
          axisLabel: {
            textStyle: {
              color: '#111',
              fontWeight: 400,
              fontSize: 12
            }
          },
        }
      ],

    yAxis: {
        type: 'category',
        axisLine: { show: false },
        axisTick: { show: false },
        offset: 2,
        axisLabel: {
          textStyle: {
            color: '#111',
            fontWeight: 400,
            fontSize: 12
          }
        },
        splitLine: { show: false },
        data: ['授权', '教练', '支持', '命令'],
    },
    series : [
        {
            name:'生活费',
            type:'bar',
            stack: '总量',
            data:[
                {value: -4, label: {normal: {position: 'right'}}},
                {value: -9, label: {normal: {position: 'right'}}},
                {value: 2, label: {normal: {position: 'right'}}},
                {value: -2, label: {normal: {position: 'right'}}},
                {value: 6, label: {normal: {position: 'right'}}},

            ],
            itemStyle:{
                 emphasis: {
              barBorderRadius: 2
            },
            normal: {
              barBorderRadius: 2,
              color:'#A674FF'
            }
            }
        }
    ]
}`

// 3
// 风险
// 柱状图
var RiskHistogramChart = `{
    legend: {
        show: true,
        icon: 'circle',
        data: ['员工平均得分'],
        bottom: 16,
        itemWidth: 10,
        itemHeight: 10
    },
    grid: {
        x: 40,
        y: 40,
        x2: 60,
        y2: 70
    },
    tooltip: {
        padding: [16, 16],
        confine: true,
        trigger: 'axis',
        backgroundColor: '#fff',
        textStyle: {
            color: '#111',
            fontSize: 12
        },
        extraCssText: 'box-shadow:0px4px12px0pxrgba(0,0,0,0.1);text-align:left;border-radius:8px;'
    },
    xAxis: {
        data: ['分裂型', '边缘型', '反社会型', '表演型', '依赖型', '偏执型', '自恋型'],
        axisTick: {
            show: false
        },
        axisLine: {
            show: true,
            lineStyle: {
                color: '#eee'
            }
        },
        axisLabel: {
            interval: 0,
            color: '#111',
            fontWeight: 'bold'
        },
        axisPointer: {
            show: true,
            type: 'shadow',
            label: {
                show: false
            },
            shadowStyle: {
                color: 'rgba(52,76,93,0.03)'
            }
        }
    },
    yAxis: {
        max: 10,
        axisLine: {
            show: false
        },
        axisTick: {
            lineStyle: {
                color: 'yellow',
                type: 'dashed',
                opacity: 0.1
            }
        },
        axisLabel: {
            textStyle: {
                color: '#999'
            }
        },
        splitLine: {
            lineStyle: {
                color: '#eee',
                type: 'dashed'
            }
        }
    },
    series: [{
        type: 'bar',
        barWidth: 28.4,
        name: '员工平均得分',
        itemStyle: {
            normal: {
                barBorderRadius: 2,
                color: {
                    x: 0,
                    y: 0,
                    x2: 0,
                    y2: 1,
                    type: 'linear',
                    global: false,
                    colorStops: [{
                        offset: 0,
                        color: '#A674FF'
                    },
                    {
                        offset: 1,
                        color: '#593DFF'
                    }]
                }
            }
        },
        data: ['4.0', '9.0', '10.0', '7.0', '7.0', '9.0', '6.0'],
        markLine: {
            silent: true,
            symbol: 'none',
            lineStyle: {
                type: 'dashed'
            },
            data: [{
                lineStyle: {
                    color: '#00A950'
                },
                label: {
                    formatter: '低风险(<3分)',
                    lineHeight: 18
                },
                yAxis: 3
            },
            {
                lineStyle: {
                    color: '#F15656'
                },
                label: {
                    formatter: '高风险(>8分)',
                    lineHeight: 18
                },
                yAxis: 8
            }]
        }
    }]
}`

// 4
// 进度条
// "900", "-h", "30"
var ProgressBarChart = `{
    tooltip: {
        trigger: 'axis',
        axisPointer: {
            type: 'shadow'
        }
    },
    grid: {
        left: '13%',
        top: '3%',
        height: 20,
        width: 300
    },
    xAxis: {
        type: '',
        show: false,
    },
    yAxis: {
        show: true,
        type: 'category',
        data: ['【合作精神】:8分'],
        axisLabel: {
            fontWeight: 600
        },
        axisLine: {
            show: false
        },
        axisTick: {
            show: false
        }
    },
    series: [{
        type: 'bar',
        data: [60],
        tooltip: {
            show: false
        },
        barMinHeight: 30,
        barWidth: 10,
        barMaxWidth: 100,
        z: 10,
        itemStyle: {
            normal: {
                barBorderRadius: 20,
                color:'#FFC311',
                label: {
                    show: false,
                    position: 'right',
                    formatter: '{c}%',
                    textStyle: {
                        color: '#000'
                    }
                }
            }
        }
    },
    {
        type: 'bar',
        data: [100],
        tooltip: {
            show: false
        },
        barMinHeight: 30,
        barWidth: 10,
        barMaxWidth: 100,
        barGap: '-100%',
        itemStyle: {
            normal: {
                barBorderRadius: 20,
                color: '#CCCCCC',
                label: {
                    show: false,
                    position: 'right',
                    testStyle: {
                        color: '#000'
                    }
                }
            }
        }
    }]
}`

// 5
// 折线柱状图
var BrokenLineHistogramChart = `{
      grid: {
        top: 30,
        bottom: 0,
        left: 45,
        right: 5,
        height: 220,
        width: 400,

      },
   xAxis: [
        {
          type: 'value',
          max: 12,
          min: 0,
          position: 'top',
          offset: 2,
          maxInterval: 2,
          splitLine: {
            lineStyle: {
              type: 'dotted',
              color: '#ddd'
            }
          },
          axisLine: {
            lineStyle: {
              type: 'dotted',
              color: '',
              width: 0,
            },
          },
          axisLabel: {
            textStyle: {
              color: '#111',
              fontWeight: 400,
              fontSize: 12
            }
          },
        }
      ],

    yAxis: {
        type: 'category',
        axisLine: { show: false },
        axisTick: { show: false },
        offset: 2,
        axisLabel: {
          textStyle: {
            color: '#111',
            fontWeight: 400,
            fontSize: 12
          }
        },
        splitLine: { show: false },
        data: ['授权', '教练', '支持', '命令','你好'],
    },
    series : [
        {
            name:'生活费',
            type:'bar',
            stack: '总量',
            data:[
                {value: 4, label: {show: true, position: 'right', color: '#111'}},
                {value: 9, label: {show: true, position: 'right', color: '#111'}},
                {value: 2, label: {show: true, position: 'right', color: '#111'}},
                {value: 2, label: {show: true, position: 'right', color: '#111'}},
                {value: 6, label: {show: true, position: 'right', color: '#111'}},
            ],
            itemStyle:{
                 emphasis: {
              barBorderRadius: 2
            },
            normal: {
              barBorderRadius: 2,
              color:'#A674FF'
            },
            label: {
                show: true,
                position: 'left',
                color: '#111'
            },
            },
            markLine: {
                 silent: true,
          symbol: 'none',
          lineStyle: {
            type: 'dashed',
          },
            }
        }
    ]
}`
