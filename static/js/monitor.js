$(document).ready(function() {
    var cpuUsage = echarts.init(document.getElementById('cpu-usage'));
    var memUsage = echarts.init(document.getElementById('mem-usage'));
    var swapUsage = echarts.init(document.getElementById('swap-usage'));

    var ws = new WebSocket("ws://" + document.location.host + "/ws"); 
    //连接打开时触发 
    ws.onopen = function(evt) {  
        console.log("Connection open ...");  
        ws.send("Hello WebSockets!");  
    };  
    //接收到消息时触发  
    ws.onmessage = function(evt) { 
        var msg = JSON.parse(evt.data);

        cpuUsageOption.series[0].data[0].value = msg.CPU.Load.Percent.toFixed(1)
        memUsageOption.series[0].data[0].value = msg.Memory.UsedPercent.toFixed(1)

        cpuUsage.setOption(cpuUsageOption, true);
        memUsage.setOption(memUsageOption, true);

        $("#temp").text(msg.CPU.Temp+" C°");
        $("#freq").text(msg.CPU.Freq.Curfreq+" MHz");
        $("#idle").text(msg.CPU.Load.Idle.toFixed(1)+"%");
        $("#user").text(msg.CPU.Load.User.toFixed(1)+"%");
        $("#sys").text(msg.CPU.Load.Sys.toFixed(1)+"%");
        $("#nice").text(msg.CPU.Load.Nice.toFixed(1)+"%");
        $("#iow").text(msg.CPU.Load.Iowait.toFixed(1)+"%");
        $("#irq").text(msg.CPU.Load.Irq.toFixed(1)+"%");
        //console.log("Received Message: " + evt);
        console.log(msg);
    };  
    //连接关闭时触发  
    ws.onclose = function(evt) {  
        console.log("Connection closed.");  
    }; 

    cpuUsageOption = {
        tooltip: {
            formatter: '{a} <br/>{b} : {c}%'
        },
        series: [
            {
                name: 'CPU占用率',
                type: 'gauge',
                detail: {formatter: '{value}%'},
                data: [{value: 0, name: 'CPU占用率'}]
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
                data: [{value: 0, name: '内存占用率'}]
            }
        ]
    };

    swapUsageOption = {
        title: {
            text: 'Swap使用率',
            subtext: '(单位MB)',
            left: 'center'
        },
        tooltip: {
            trigger: 'item',
            formatter: '{a} <br/>{b}: {c} ({d}%)'
        },
        legend: {
            orient: 'vertical',
            left: 10,
            data: ['已使用', '未使用']
        },
        series: [
            {
                name: '访问来源',
                type: 'pie',
                radius: ['50%', '70%'],
                avoidLabelOverlap: false,
                label: {
                    show: false,
                    position: 'center'
                },
                emphasis: {
                    label: {
                        show: true,
                        fontSize: '30',
                        fontWeight: 'bold'
                    }
                },
                labelLine: {
                    show: false
                },
                data: [
                    {value: host.Mem.SwapTotal - host.Mem.SwapFree, name: '已使用'},
                    {value: host.Mem.SwapFree, name: '未使用'},
                ]
            }
        ]
    };

    var net_traffic = new Array();
    var net_traffic_option = new Array();

    $.each(host.Interface.Interfaces, function(i,j){
        net_traffic.push(echarts.init(document.getElementById('network-'+(j.Name))));
    
        net_traffic_option.push({
            title: {
                text: '',
            },
            color: ["#f05454", "#006699"],
            legend: {
                icon:'roundRect',//标记图标，方形
                textStyle: {
    
                }
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
                nameTextStyle:{
                    padding:[15,0,0,0],
                    fontSize:14,
                },
                type: 'time',
                maxInterval: 3600*1*1000,
                splitLine: { 
                    show: false
                }
            },
            yAxis: {
                name:'流量',
                nameLocation:'center',
                nameGap:1,
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
                showSymbol: false,
                hoverAnimation: false,
                data: new Array(),
                
            },
            {
                name: '出口网速',
                type: 'line',
                showSymbol: false,
                hoverAnimation: false,
                data: new Array(),
            }]
        });
    })


    cpuUsage.setOption(cpuUsageOption, true);
    memUsage.setOption(memUsageOption, true);
    swapUsage.setOption(swapUsageOption, true);

    $.each(net_traffic,function(key,val){
        val.setOption(net_traffic_option[key], true);
    });  
})