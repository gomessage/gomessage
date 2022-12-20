<template>
    <el-form :label-position="labelPosition" label-width="100px" :model="client" style="text-align: left">
        <el-form-item label="客户端名称:">
            <el-input v-model="client.client_name" placeholder=""></el-input>
        </el-form-item>

        <el-form-item label="客户端描述:">
            <el-input v-model="client.client_description" placeholder=""></el-input>
        </el-form-item>

        <el-form-item label="客户端类型:">
            <el-input v-model="client.typeDescription" disabled></el-input>
        </el-form-item>

        <el-divider content-position="center">机器人</el-divider>

        <el-form-item label="放行关键字:">
            <el-input v-model="client.client_info.robot_keyword" placeholder="要和机器人界面设置的放行关键词保持一直"></el-input>
        </el-form-item>

        <div v-for="(list, index) in client.client_info.robot_url" :key="index">
            <el-form-item label="机器人URL:">
                <el-input v-model="list.url" placeholder="从钉钉上粘贴而来的机器人URL地址" style="width: 85%"></el-input>
                <el-button type="danger" icon="el-icon-delete" circle size="mini" v-on:click="del(index)" style="margin-left: 10px"></el-button>
            </el-form-item>
        </div>
        <div style="text-align: center">
            <el-button type="text" icon="el-icon-plus" size="mini" @click="add">再添加一个机器人</el-button>
        </div>
        <p id="textStype">此处可以添加多个机器人，推送消息时会从中随机挑选一个URL来使用，可以避免单个机器人消息推送时（每分钟）的次数限制，避免重要报警信息被漏送的可能。</p>

        <br><br>
        <el-form-item>
            <el-button type="primary" @click="onSubmit">立即创建</el-button>
            <el-button @click="closeDrawer">取消</el-button>
        </el-form-item>
    </el-form>
</template>

<script>
import {postClient} from '@/service/requests'

export default {
    name: "clientDingtalk",
    data() {
        return {
            labelPosition: "right", //表格标题右对齐

            client: {
                client_name: "",
                client_description: "",
                client_type: "dingtalk", //客户端类型，与后端的约定，禁止修改
                is_active: false,
                client_info: {
                    robot_keyword: "",
                    robot_url: [
                        {
                            url: "",
                        },
                    ],
                },
                typeDescription: "钉钉·机器人",


            },
        }
    }, props: {
        getClientList: Function,
    },
    methods: {
        onSubmit: function () {
            postClient(this.client).then(response => {
                if (response.data.result) {
                    this.$message.success("添加成功...")
                    this.getClientList();
                    this.$store.commit("updateDrawerStatus", false);

                } else {
                    this.$message.error("添加失败...");
                    this.getClientList();
                    this.$store.commit("updateDrawerStatus", false);
                }
            }).catch(err => {
                console.log(err);
            });
        },
        // 添加输入框
        add: function () {
            let cope = {
                url: "",
            };
            this.client.client_info.robot_url.push(cope);
            // console.log(this.mapList2);
            for (let i = 0; i < this.client.client_info.robot_url.length; i++) {
                console.log(this.client.client_info.robot_url[i]);
            }
        },
        del: function (index) {
            this.client.client_info.robot_url.splice(index, 1);
            // console.log(this.mylist);
        },
        closeDrawer: function () {
            this.getClientList();
            this.$store.commit("updateDrawerStatus", false);
        },
    }
}
</script>

<style scoped>
#textStype {
    font-size: 11px;
    font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
    line-height: 1.7;
    color: #C0C4CC;
}
</style>
