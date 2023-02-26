<template>

    <el-card shadow="always">
        <div slot="header">
            <span style="padding-left: 100px">过境数据</span>
            <el-button style="float: right; padding: 3px 0; margin-left: 10px" type="text" v-on:click="getServerData()">劫持数据</el-button>
            <el-button style="float: right; padding: 3px 0;" type="text" v-on:click="copyCode()">一键复制</el-button>

        </div>

        <!--        <el-row style="color: #C0C4CC;font-size: 12px;line-height: 1.7;padding-bottom: 20px">-->
        <!--            <el-col span="12" style="text-align: left">数据过境时间：{{ codeUpdateTime }}</el-col>-->
        <!--            <el-col span="8" offset="4" style="text-align: right">查看文档</el-col>-->
        <!--        </el-row>-->

        <div>
            <!--            <div id="codeTime">数据过境时间：{{ codeUpdateTime }}</div>-->
            <!--下面这一行是代码框，不能换行；如果换行第二行的空格就显示在代码框里面了-->
            <pre id="CodeStyle"><code id="CodeContent">{{ codeJsonContent }}</code></pre>
        </div>

    </el-card>

</template>

<script>
import {queryJson} from '@/service/requests'

export default {
    name: "cCodeFormat",
    data() {
        return {
            codeJsonContent: "点击 <劫持数据> 可以看到推送进GoMessage的原始数据...\n\n每次 <劫持数据> 都可以实时抓取到最新的一条数据...\n\n此处对数据进行了格式化与对齐，您可以把数据 <一键复制> 到其它地方使用...",
            codeUpdateTime: null,
            arr2: [],
        };
    },
    methods: {
        //获取数据（这个函数已经被弃用）
        // getData: function () {
        //     const jsonData = require('./aaa.json'); //这里拿到的是一个json对象
        //     // const jsonObject = JSON.parse(jsonData); //如果从后端过来拿到的只是json字符串，则使用这个先转成对象
        //     const jsonShow = JSON.stringify(jsonData, null, 4) //格式化json字符串，便于后面进行展示
        //     return jsonShow
        // },
        //
        //修改文本（这个函数已经被弃用）
        // updateText: function () {
        //     this.codeJsonContent = this.getData()
        // },


        //一键粘贴（纯原生js document对象实现）
        copyCode: function () {
            const val = document.getElementById('CodeContent');
            window.getSelection().selectAllChildren(val);
            document.execCommand("Copy");
            // alert("已复制好，可以贴粘...");
            this.$message('复制成功...');
        },


        //获取服务端数据（服务端已经在后台劫持好"过路"数据了）
        getServerData: function () {
            // 这里如果使用funciton(response){}这个种模式，那么this指的就不是vue实例了，
            // 如果下面这样使用箭头函数，就可以实现上下文穿透，this就可以代指vue实例本身了，然后就有$router方法了
            queryJson(this.$store.getters.getNamespace, null).then(response => {
              console.log(response.data);
              const jsonData = response["data"]["request_json"];
              const t = response["data"]["time"];

              if (jsonData === null || jsonData.length === 0) {
                this.$message({
                  showClose: false,
                  message: "没有数据进入GoMessage服务，无法向您展示数据。"
                });
              } else {
                    this.codeJsonContent = JSON.stringify(jsonData, null, 2);
                    this.codeUpdateTime = this.dateFormat(new Date(t));
                }
            }).catch(function (error) {
                console.log(error);
            });
        },


        //时间处理格式化
        // dateFormat: function (date) {
        //     console.log(date)
        //     const daters = date
        //     if (daters != null) {
        //         const dateMat = new Date(daters);
        //         const year = dateMat.getFullYear();
        //         const month = dateMat.getMonth() + 1;
        //         const day = dateMat.getDate();
        //         const hh = dateMat.getHours();
        //         const mm = dateMat.getMinutes();
        //         const ss = dateMat.getSeconds();
        //         return year + "-" + (month < 10 ? "0" : "") + month + "-" + (day < 10 ? "0" : "") + day + (hh < 10 ? "0" : "") + " " + hh + ":" + (mm < 10 ? "0" : "") + mm + ":" + (ss < 10 ? "0" : "") + ss;
        //     }
        // },


    },
}
</script>

<style scoped>
/*#DataFormatCard {*/
/*    width: 100%;*/
/*}*/

#CodeStyle {
    text-align: left;
    background-color: #2f2f35;
    /*font-family: Monaco,  Arial, Helvetica, sans-serif;*/
    /*font-family: "Andale Mono WT", "Andale Mono", "Lucida Console", "Lucida Sans Typewriter", "DejaVu Sans Mono", "Bitstream Vera Sans Mono", "Liberation Mono", "Nimbus Mono L", "Courier New", Courier, monospace;*/
    font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
    font-size: 13px;
    line-height: 22px;
    color: violet;
    overflow-x: auto;
    min-height: 20rem;
}

#CodeContent {

}

#codeTime {
    text-align: left;
    margin: 10px 0;
}
</style>
