<template>
<!--    <el-row style="margin-top: 20px;margin-bottom: 20px">-->
        <el-row >
            <el-col span="11" offset="1">
                <CConfigMapShow></CConfigMapShow>
            </el-col>

            <el-col span="10" offset="1">
            <el-card class="box-card" id="MessageTemplateCard" shadow="always">
                <div slot="header" class="clearfix">
                    <span style="padding-left: 50px">消息模板</span>
                    <el-button style="float: right; padding: 3px 0; margin-left: 15px" type="text" v-on:click="pushTemplateData()">更新模板</el-button>
                    <!--                    <el-button style="float: right; padding: 3px 0;" type="text" v-on:click="pullTemplateData()">拉取</el-button>-->
                </div>
                <div>
                    <el-input
                        type="textarea"
                        :autosize="{ minRows: 10, maxRows: 200}"
                        placeholder="请输入Golang语法的模板内容"
                        v-model="textarea"
                        resize="none"
                    >
                    </el-input>
                </div>
            </el-card>
        </el-col>
    </el-row>
</template>

<script>
import CConfigMapShow from "@/views/bak/cConfigMapShow";
export default {
    name: "ViewTemplate",
    data() {
        return {
            textarea: '点击拉取按钮，获取示例模板',
        }
    },
    components:{
        CConfigMapShow,
    },
    methods: {
        printData: function () {
            console.log(this.textarea);
        },
        //一键粘贴（纯原生js document对象实现）
        copyCode: function () {
            const val = document.getElementById('templateContent');
            // let val = this.textarea;
            window.getSelection().selectAllChildren(val);
            document.execCommand("Copy");
            alert("已复制好，可以贴粘...");
        },
        handlePaste(e) {
            let clip = e.clipboardData.getData('Text');
            this.text = clip;
            console.log(clip);
        },
        //推送数据
        pushTemplateData: function () {

            console.log(this.textarea);

            this.axios.post('/v1/web/template',
                {
                    "message_template": this.textarea,
                }
            )
                .then(response => {
                    // alert("模板已存储在库中...");
                    console.log(response.data);
                    this.$message.success("数据库更新成功...")
                })
                .catch(err => {
                    console.log(err);
                });
        },
        //拉取数据
        pullTemplateData: function () {
            this.axios.get('/v1/web/template')
                .then(response => {
                    this.textarea = response.data["MessageTemplate"];
                    console.log(response.data);
                })
                .catch(err => {
                    console.log(err);
                });
        }
    }, created() {
        //修改步骤条的值
        this.$store.commit("updateStepsActive",2);
        //拉取数据
        this.pullTemplateData();
    }
}
</script>

<style scoped>
#MessageTemplateContent {
    width: 100%;
    /*margin-left: 20%;*/
}

.templateContentClass >>> input {
    background-color: #2c3e50;
}
</style>
