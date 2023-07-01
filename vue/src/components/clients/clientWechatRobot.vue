<template>
  <el-form :label-position="labelPosition" :model="client" :rules="clientRules" label-width="100px" style="text-align: left">
    <el-form-item label="客户端名称:" prop="client_name">
      <el-input v-model="client.client_name" placeholder=""></el-input>
    </el-form-item>

    <el-form-item label="客户端描述:">
      <el-input v-model="client.client_description" placeholder=""></el-input>
    </el-form-item>

    <el-form-item label="客户端类型:">
      <el-input v-model="client.typeDescription" disabled></el-input>
    </el-form-item>

    <el-divider content-position="center">机器人</el-divider>

    <!--<el-form-item label="放行关键字:">-->
    <!--  <el-input v-model="client.client_info.robot_keyword" placeholder="要和机器人界面设置的放行关键词保持一直"></el-input>-->
    <!--</el-form-item>-->

    <div v-for="(list, index) in client.client_info.robot_url_list" :key="index">
      <el-form-item label="机器人URL:">
        <el-input v-model="list.url" placeholder="机器人URL地址" style="width: 85%"></el-input>
        <el-button
          circle
          icon="el-icon-delete"
          size="mini"
          style="margin-left: 10px"
          type="danger"
          v-on:click="del(index)">
        </el-button>
      </el-form-item>
    </div>
    <div style="text-align: center">
      <el-button icon="el-icon-plus" size="mini" type="text" @click="add">再添加一个机器人</el-button>
    </div>
    <p id="textStype">
      此处可以添加多个机器人，推送消息时会从中随机挑选一个URL来使用，可以避免单个机器人消息推送时（每分钟）的次数限制，避免重要报警信息被漏送的可能。</p>

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
  name: "clientWechatRobot",
  data() {
    return {
      labelPosition: "right", //表格标题右对齐

      client: {
        client_name: "",
        client_description: "",
        client_type: "wechat_robot", //客户端类型，与后端的约定，禁止修改
        is_active: false,
        client_info: {
          robot_keyword: "",
          robot_url_list: [
            {
              url: "",
            },
          ],
        },
        typeDescription: "企业微信·群机器人",
      },
      clientRules: {
        client_name: [
          {required: true, message: "name不能为空", trigger: "blur"},
        ],
      }
    }
  }, props: {
    cli_GetClientList: Function,
  },
  methods: {
    onSubmit: function () {
      postClient(this.$store.getters.getNamespace, this.client).then(response => {
        if (response.data.code === 1) {
          this.$message.success("添加成功...")
          this.cli_GetClientList();
          this.$store.commit("updateDrawerStatus", false);

        } else {
          this.$message.error("添加失败...");
          this.cli_GetClientList();
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
      this.client.client_info.robot_url_list.push(cope);
      // console.log(this.mapList2);
      for (let i = 0; i < this.client.client_info.robot_url_list.length; i++) {
        console.log(this.client.client_info.robot_url_list[i]);
      }
    },
    del: function (index) {
      this.client.client_info.robot_url_list.splice(index, 1);
      // console.log(this.mylist);
    },
    closeDrawer: function () {
      this.cli_GetClientList();
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
