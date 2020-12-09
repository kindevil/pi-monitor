$(document).ready(function() {
var myChart = echarts.init(document.getElementById('cpu'));
var memChart = echarts.init(document.getElementById('mem'));
var memUsage = echarts.init(document.getElementById('disk'));
var swapUsage = echarts.init(document.getElementById('swap'));

cpuUsageOption = {
    tooltip: {
        formatter: '{a} <br/>{b} : {c}%'
    },
    series: [
        {
            name: 'CPU占用率',
            type: 'gauge',
            detail: {formatter: '{value}%'},
            data: [{value: 50, name: 'CPU占用率'}]
        }
    ]
};

memUsageOption = {
    tooltip: {
        formatter: '{a} <br/>{b} : {c}%'
    },
    series: [
        {
            name: '内存占用率',
            type: 'gauge',
            detail: {formatter: '{value}%'},
            data: [{value: 50, name: '内存占用率'}]
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
            [now.getHours(), now.getMinutes() + 1, now.getSeconds()].join('/'),
            Math.round(value)
        ]
    };
}

var traffic_in = new Array();
var traffic_out = new Array();
var net_traffic = new Array();
var net_traffic_option = new Array();

for (i=0;i<net_count;i++) {
    net_traffic.push(echarts.init(document.getElementById('net-traffic-'+(i+1))));

    net_traffic_option.push({
        title: {
            text: '24小时流量监控图',
        },
        legend: {
            icon:'rect',//标记图标，方形
        },
        tooltip: { 
            trigger: 'axis',
            formatter: function (params) {
                var tmpparams = params[1];
                params = params[0];
                if (tmpparams.value[1] > 1024 || params.value[1] > 1024 ) {
                    return params.value[0] + '</br>'+params.seriesName+':' + (params.value[1]/1024).toFixed(2) + "Mb/s"+'</br>'+tmpparams.seriesName+':' + (tmpparams.value[1]/1024).toFixed(2) + "Mb/s";
                } else {
                    return params.value[0] + '</br>'+params.seriesName+':' + params.value[1] + "kb/s"+'</br>'+tmpparams.seriesName+':' + tmpparams.value[1] + "kb/s";
                }
                return params.value[0] + '</br>'+params.seriesName+':' + params.value[1] + "kb/s";
            },
            axisPointer: {
                animation: false
            },
        },
        xAxis: {
            name:'时间',
            nameGap:30,
            nameTextStyle:{
                padding:[15,0,0,0],
                fontSize:14,
            },
            type: 'time',
            maxInterval: 3600*2*1000,
            splitLine: { 
                show: false
            },
        },
        yAxis: {
            name:'流量',
            nameLocation:'center',
            nameGap:50,
            type: 'value',
            min:0,
            splitLine: {
                show: true
            },
            axisLabel: {
                formatter: function(value, index) {
                    if (value > 1024) {
                        return parseInt(value/1024)+"M";
                    }
                    return value+"k";
                },
            }
        },
        series: [{
            name: '入口网速',
            type: 'line',
            //step:true,
            showSymbol: false,
            hoverAnimation: false,
            data: new Array(),
            //areaStyle: {},
            // lineStyle:{
            //     opacity:0,
            // },
            itemStyle:{
                color:'rgb(0,204,0)',
            },
            
        },
        {
            name: '出口网速',
            type: 'line',
            showSymbol: false,
            hoverAnimation: false,
            data: new Array(),
            itemStyle:{
                opacity:0.7,
                color:'rgb(0,0,225)',
            }
        }]
    });
}

//时间格式化
Date.prototype.Format = function (fmt) {  
    var o = {  
        "M+": this.getMonth() + 1, //月份   
        "d+": this.getDate(), //日   
        "H+": this.getHours(), //小时   
        "m+": this.getMinutes(), //分   
        "s+": this.getSeconds(), //秒   
        "q+": Math.floor((this.getMonth() + 3) / 3), //季度   
        "S": this.getMilliseconds() //毫秒   
    };  
    if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));  
    for (var k in o)  
    if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));  
    return fmt;  
}

setInterval(function () {
    var currtime = new Date().Format("yyyy-MM-dd HH:mm:ss");

    axios.get('/api/get')
        .then(function(response) {
            console.log(response.data);
            cpuUsageOption.series[0].data[0].value = response.data.CPU.Load.Percent
            memUsageOption.series[0].data[0].value = response.data.Memory.UsedPercent.toFixed(2)
            swapOption.series[0].data[0].value = response.data.Memory.SwapTotal - response.data.Memory.SwapFree
            swapOption.series[0].data[1].value = response.data.Memory.SwapFree

            $("#core").text(response.data.CPU.Cores);
            $("#temp").text(response.data.CPU.Temp+"C°");
            $("#freq").text(response.data.CPU.Freq.Curfreq+"MHz");
            $("#idle").text(response.data.CPU.Load.Idle+"%");
            $("#user").text(response.data.CPU.Load.User+"%");
            $("#sys").text(response.data.CPU.Load.Sys+"%");
            $("#nice").text(response.data.CPU.Load.Nice+"%");
            $("#iow").text(response.data.CPU.Load.Iowait+"%");
            $("#irq").text(response.data.CPU.Load.Irq+"%");
            $("#mem-total").text(response.data.Memory.Total+"MB");
            $("#mem-used").text(response.data.Memory.Used+"MB");
            $("#mem-free").text(response.data.Memory.Free+"MB");
            $("#mem-shared").text(response.data.Memory.Shared+"MB");
            $("#mem-cache").text(response.data.Memory.Cached+"MB");
            $("#mem-ava").text(response.data.Memory.Available+"MB");
            
            $.each(response.data.Net.Interface,function(key,val){
                var in_data = [];
                var out_data = [];
                in_data.push(currtime)
                in_data.push(val.Recv)
                out_data.push(currtime)
                out_data.push(val.Send)

                if (net_traffic_option[key].series[0].data.length >= 600) {
                    net_traffic_option[key].series[0].data.shift()
                    net_traffic_option[key].series[1].data.shift()
                }

                net_traffic_option[key].series[0].data.push(in_data)
                net_traffic_option[key].series[1].data.push(out_data)
            });
        })
        .catch(function(error){
            console.log(error);

            $.each(net_traffic_option,function(key,val){
                var in_data = [];
                var out_data = [];
                in_data.push(currtime)
                in_data.push(0)
                out_data.push(currtime)
                out_data.push(0)

                if (net_traffic_option[key].series[0].data.length >= 600) {
                    net_traffic_option[key].series[0].data.shift()
                    net_traffic_option[key].series[1].data.shift()
                }

                net_traffic_option[key].series[0].data.push(in_data)
                net_traffic_option[key].series[1].data.push(out_data)
            });
        })

    //cpuUsageOption.series[0].data[0].value = (Math.random() * 100).toFixed(2) - 0;
    myChart.setOption(cpuUsageOption, true);
    memChart.setOption(memUsageOption, true);
    memUsage.setOption(diskOption, true)
    swapUsage.setOption(swapOption, true)

    //net_traffic.setOption(net_traffic_option, true);

    $.each(net_traffic,function(key,val){
        val.setOption(net_traffic_option[key], true);
    });

    // $.each(net_traffic,function(key,val){

    //     for (var i = 0; i < 5; i++) {        
    //         traffic_in.shift();
    //         traffic_in.push(randomData());
    //         traffic_out.shift();
    //         traffic_out.push(randomData());
    //     }

    //     net_traffic_option[key].series[0].data = traffic_in
    //     net_traffic_option[key].series[1].data = traffic_out
    //     val.setOption(net_traffic_option[key], true);
    // });
    
},1000);

});