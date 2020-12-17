$(document).ready(function() {
    var cpuUsage = echarts.init(document.getElementById('cpu-usage'));
    var memUsage = echarts.init(document.getElementById('mem-usage'));
    var swapUsage = echarts.init(document.getElementById('swap-usage'));

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


    cpuUsage.setOption(cpuUsageOption, true);
    memUsage.setOption(memUsageOption, true);
    swapUsage.setOption(swapUsageOption, true);
})