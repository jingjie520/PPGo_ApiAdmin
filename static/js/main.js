/*
* @Author: haodaquan
* @Date:   2017-09-22 10:50:33
* @Last Modified by:   haodaquan
* @Last Modified time: 2017-09-22 10:50:33
*/
var Snow = {
    getOnOrOff: function (val) {
        var color = val === "off" ? "#ff0000" : "#11cd6e",
            text = val === "off" ? "关闭" : "开启";

        console.log(val);

        return '<span style="color:' + color + '">' + text + '</span>';
    }
};