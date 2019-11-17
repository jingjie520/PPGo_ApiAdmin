/*
* @Author: haodaquan
* @Date:   2017-09-22 10:50:33
* @Last Modified by:   haodaquan
* @Last Modified time: 2017-09-22 10:50:33
*/
var Snow = {
    getOnOrOff: function (val) {
        var color = "",
            text = val;

        switch (val) {
            case "off":
                color = "#ff0000";
                text = "关闭";
                break;
            case "on":
                color = "#11cd6e";
                text = "开启";
        }

        return '<span style="color:' + color + '">' + text + '</span>';
    },
    getStartOrStop: function (val) {
        var color = "",
            text = val;

        switch (val) {
            case "stop":
                color = "#ff0000";
                text = "停止";
                break;
            case "start":
                color = "#11cd6e";
                text = "开启";
        }

        return '<span style="color:' + color + '">' + text + '</span>';
    },
    getDemux: function (val) {
        var html = '';
        switch (val) {
            case 0 :
                html += '<span style="color: #ff6600">开始连接流</span>';
                break;
            case 1 :
                html += '<span style="color: #11cd6e">拉流成功</span>';
                break;
            case 2 :
                html += '<span style="color: #ff0000">拉流结束</span>';
                break;
            default:
                html += '<span style="color: #ff0000">拉流结束</span>';
        }
        return html;
    }
};