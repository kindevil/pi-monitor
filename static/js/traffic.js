$(document).ready(function() {
    var myChart = echarts.init(document.getElementById('net-traffic-1'));
    var traffic_in = [];
    var traffic_out = [];
    option = {
        title: {
            text: '24小时流量监控图',//图片标题
        },
        legend:{
            icon:'rect',//标记图标，方形
        },
        tooltip: { //focus显示内容
            trigger: 'axis',
            formatter: function (params) {
                var tmpparams = params[1];
                params = params[0];
                if(tmpparams)
                {
                    return params.value[0] + '</br>'+params.seriesName+':' + params.value[1] + "kb/s"+'</br>'+tmpparams.seriesName+':' + tmpparams.value[1] + "kb/s";
                }
                return params.value[0] + '</br>'+params.seriesName+':' + params.value[1] + "kb/s";
            },
            axisPointer: {
                animation: false
            },
        },
        xAxis: { //x轴,类型为日期格式，故在数据中添加了一个24小时的数组，以调整x坐标系显示数据
            name:'时间',
            nameGap:30,
            nameTextStyle:{
                padding:[15,0,0,0],
                fontSize:14,
            },
            type: 'time',
            maxInterval: 3600*2*1000,
            //min:data.time[0][0],
            //max:data.time[1][0],
            splitLine: { //显示分割线
                show: false
            },
        },
        yAxis: { //y轴，添加留白策略数据变多时y轴突然拉的比较长
            name:'bits per second(kb/s)',
            nameLocation:'center',
            nameGap:30,
            type: 'value',
            min:0,
            // boundaryGap: [0, '30%'],//坐标轴两边留白策略
            splitLine: {
                show: false
            }
        },
        series: [{	//数据
            name: '入口网速',
            type: 'line',
            step:true, //是否支持骤变，false有段数据为空时为渐变
            showSymbol: false,
            hoverAnimation: false,
            data: traffic_in,
            areaStyle: {},
            lineStyle:{
                opacity:0,
            },
            itemStyle:{
                color:'rgb(0,204,0)',
            },
            
        },
        {
            name: '出口网速',
            type: 'line',
            showSymbol: false,
            hoverAnimation: false,
            data: traffic_out,
            // lineStyle:{
            // 	opacity:0.7,
            // 	color:'rgb(0,0,225)'
            // },
            itemStyle:{
                opacity:0.7,
                color:'rgb(0,0,225)',
            }
        },
        {
            type: 'line',
            tooltip:{trigger:'none'}, 
            //data:data.time,//数据格式[[2019-07-04 15:20:12,-1],[2019-07-05 15:20:12,-1]]
        }]
    };
    
    //时间格式化问题
    // Date.prototype.Format = function (fmt) {  
    //     var o = {  
    //         "M+": this.getMonth() + 1, //月份   
    //         "d+": this.getDate(), //日   
    //         "H+": this.getHours(), //小时   
    //         "m+": this.getMinutes(), //分   
    //         "s+": this.getSeconds(), //秒   
    //         "q+": Math.floor((this.getMonth() + 3) / 3), //季度   
    //         "S": this.getMilliseconds() //毫秒   
    //     };  
    //     if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));  
    //     for (var k in o)  
    //     if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));  
    //     return fmt;  
    // }

    // setInterval(function () {
    //     var data = [];
    //     //var now = +new Date(1997, 9, 3);
    //     var currtime = new Date().Format("yyyy-MM-dd HH:mm:ss");
    //     for (var i =0; i < 10;i++) {
    //         data.push(currtime)
    //         data.push(123)
    //         traffic_in.push(data)
    //         traffic_out.push(data)
    //     }
    //     console.log(traffic_in)
    //     myChart.setOption(option);
    // }, 1000);
});
