var myChart = echarts.init(document.getElementById('cpu'));
var memChart = echarts.init(document.getElementById('mem'));
var memUsage = echarts.init(document.getElementById('disk'));
var swapUsage = echarts.init(document.getElementById('swap'));
var netChart = echarts.init(document.getElementById('net'));

cpuUsageOption = {
    tooltip: {
        formatter: '{a} <br/>{b} : {c}%'
    },
    series: [
        {
            name: 'CPU占用率',
            type: 'gauge',
            detail: {formatter: '{value}%'},
            data: [{value: 50, name: '占用率'}]
        }
    ]
};



diskOption = {
    title: {
        text: '硬盘使用率',
        subtext: '(单位GB)',
        left: 'center'
    },
    tooltip: {
        trigger: 'item',
        formatter: '{a} <br/>{b} : {c} GB ({d}%)'
    },
    legend: {
        orient: 'vertical',
        left: 'left',
        data: ['已使用', '未使用']
    },
    series: [
        {
            name: '访问来源',
            type: 'pie',
            radius: '55%',
            center: ['50%', '60%'],
            data: [
                {value: 5.7, name: '已使用'},
                {value: 10.3, name: '未使用'},
            ],
            emphasis: {
                itemStyle: {
                    shadowBlur: 10,
                    shadowOffsetX: 0,
                    shadowColor: 'rgba(0, 0, 0, 0.5)'
                }
            }
        }
    ],
    color : [
        '#30475e',
        '#fd8c04'
    ],
};

swapOption = {
    title: {
        text: 'Swap使用率',
        subtext: '(单位MB)',
        left: 'center'
    },
    tooltip: {
        trigger: 'item',
        formatter: '{a} <br/>{b} : {c} MB ({d}%)'
    },
    legend: {
        orient: 'vertical',
        left: 'left',
        data: ['已使用', '未使用']
    },
    series: [
        {
            name: '访问来源',
            type: 'pie',
            radius: '55%',
            center: ['50%', '60%'],
            data: [
                {value: 77, name: '已使用'},
                {value: 22, name: '未使用'},
            ],
            emphasis: {
                itemStyle: {
                    shadowBlur: 10,
                    shadowOffsetX: 0,
                    shadowColor: 'rgba(0, 0, 0, 0.5)'
                }
            }
        }
    ],
    color : [
        '#f05454',
        '#16697a'
    ],
};

function randomData() {
    now = new Date(+now + oneDay);
    value = value + Math.random() * 21 - 10;
    return {
        name: now.toString(),
        value: [
            [now.getFullYear(), now.getMonth() + 1, now.getDate()].join('/'),
            Math.round(value)
        ]
    };
}

function randomData1() {
    now = new Date(+now + oneDay);
    value = value + Math.random() * 21 - 10;
    return {
        name: now.toString(),
        value: [
            [now.getFullYear(), now.getMonth() + 1, now.getDate()].join('/'),
            Math.round(value+100)
        ]
    };
}

var data = [];
var data1 = [];
var now = +new Date(1997, 9, 3);
var oneDay = 24 * 3600 * 1000;
var value = Math.random() * 1000;
for (var i = 0; i < 1000; i++) {
    data.push(randomData());
    data1.push(randomData1());
}

netOption = {
    title: {
        text: '动态数据 + 时间坐标轴'
    },
    tooltip: {
        trigger: 'axis',
        formatter: function (params) {
            params = params[0];
            var date = new Date(params.name);
            return date.getDate() + '/' + (date.getMonth() + 1) + '/' + date.getFullYear() + ' : ' + params.value[1] + 'KB';
        },
        axisPointer: {
            animation: false
        }
    },
    xAxis: {
        type: 'time',
        splitLine: {
            show: false
        }
    },
    yAxis: {
        type: 'value',
        boundaryGap: [0, '100%'],
        splitLine: {
            show: false
        }
    },
    series: [{
        name: '模拟数据',
        type: 'line',
        showSymbol: false,
        hoverAnimation: false,
        data: data
    },{
        name: '模拟数据1',
        type: 'line',
        showSymbol: false,
        hoverAnimation: false,
        data: data1
    }]
};


memUsage.setOption(diskOption, true)
swapUsage.setOption(swapOption, true)

setInterval(function () {
    cpuUsageOption.series[0].data[0].value = (Math.random() * 100).toFixed(2) - 0;
    myChart.setOption(cpuUsageOption, true);
    memChart.setOption(cpuUsageOption, true);

    for (var i = 0; i < 5; i++) {
        data.shift();
        data1.shift();
        data.push(randomData());
        data1.push(randomData1());
    }

    netOption.series[0].data = data
    netOption.series[0].data1 = data1

    netChart.setOption(netOption, true);
},2000);
