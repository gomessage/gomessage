<template>
  <el-form :label-position="labelPosition" label-width="100px" :model="client" :rules="wechatRules" style="text-align: left">
    <el-form-item label="客户端名称:" prop="client_name">
      <el-input v-model="client.client_name" placeholder=""></el-input>
    </el-form-item>

    <el-form-item label="客户端描述:">
      <el-input v-model="client.client_description" placeholder=""></el-input>
    </el-form-item>

    <el-form-item label="客户端类型:">
      <el-input v-model="client.typeDescription" disabled></el-input>
    </el-form-item>

    <el-divider content-position="center">应用号</el-divider>

    <el-form-item label="企业ID:">
      <el-input v-model="client.client_info.corp_id" placeholder="请联系企业微信管理员获取"></el-input>
    </el-form-item>

    <el-form-item label="应用AgentId:">
      <el-input v-model="client.client_info.agent_id" placeholder="请联系企业微信管理员获取"></el-input>
    </el-form-item>

    <el-form-item label="应用Secret:">
      <el-input v-model="client.client_info.secret" placeholder="请联系企业微信管理员获取" show-password></el-input>
    </el-form-item>

    <el-form-item label="接收用户:">
      <el-input
        type="textarea"
        :autosize="{ minRows: 4, maxRows: 6}"
        v-model="client.client_info.touser"
        placeholder="可以填写多个用户账号，用 | 分割开 （例如：aaa|bbb）"
      >
      </el-input>
    </el-form-item>

    <el-form-item>
      <el-button type="primary" @click="onSubmit">立即创建</el-button>
      <el-button @click="closeDrawer">取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import {postClient} from '@/service/requests'

export default {
  name: "clientWechat",
  data() {
    return {
      labelPosition: "right",
      client: {
        client_name: "",
        client_description: "",
        client_type: "wechat", //客户端类型，与后端的约定，禁止修改
        is_active: false,
        client_info: {
          corp_id: "",
          agent_id: "",
          secret: "",
          touser: ""
        },
        typeDescription: "企业微信·应用号",
      },
      wechatRules: {
        client_name: [
          {required: true, message: "name不能为空", trigger: "blur"},
        ],
      },
    }
  }, props: {
    getClientList: Function,
  },
  methods: {
    onSubmit: function () {
      let secret = this.client.client_info.secret;
      if (secret.length <= 15) {
        this.$message.error("应用Secret的输入长度不符合要求...");
      } else {
        postClient(this.client).then(response => {
          if (response.data.result) {
            this.$message.success("添加成功...");
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
      }
    },
    closeDrawer: function () {
      this.getClientList();
      this.$store.commit("updateDrawerStatus", false);
    },
  }
}
</script>

<style scoped>

</style>
