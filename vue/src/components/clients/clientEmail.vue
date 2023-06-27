<template>
  <!-- 表单开始 -->
  <el-form :label-position="labelPosition" :model="client" :rules="clientRules" label-width="100px" style="text-align: left">
    <!-- 客户端名称输入框 -->
    <el-form-item label="客户端名称:" prop="client_name">
      <el-input v-model="client.client_name" placeholder=""></el-input>
    </el-form-item>

    <!-- 客户端描述输入框 -->
    <el-form-item label="客户端描述:">
      <el-input v-model="client.client_description" placeholder=""></el-input>
    </el-form-item>

    <!-- 客户端类型输入框（禁止修改） -->
    <el-form-item label="客户端类型:">
      <el-input v-model="client.typeDescription" disabled></el-input>
    </el-form-item>

    <!-- 分割线 -->
    <el-divider content-position="center">机器人</el-divider>

    <!-- 邮箱标题输入框 -->
    <el-form-item label="邮箱标题:">
      <el-input v-model="client.client_info.robot_keyword" placeholder="例如：GoMessage消息通知"></el-input>
    </el-form-item>

    <!-- 邮箱地址输入框，可以添加多个 -->
    <div v-for="(list, index) in client.client_info.robot_url_list" :key="index">
      <el-form-item label="邮箱地址:">
        <el-input v-model="list.url" placeholder="例如：admin@example.com" style="width: 85%"></el-input>
        <!-- 删除按钮 -->
        <el-button circle icon="el-icon-delete" size="mini" style="margin-left: 10px" type="danger"
                   v-on:click="del(index)"></el-button>
      </el-form-item>
    </div>
    <!-- 添加邮箱按钮 -->
    <div style="text-align: center">
      <el-button icon="el-icon-plus" size="mini" type="text" @click="add">再添加一个邮箱</el-button>
    </div>
    <br>
    <!-- 提示信息 -->
    <p id="textStype">此处可以添加多个邮箱，GoMessage会同时向多个邮箱推送内容。</p>

    <br><br>
    <!-- 提交和取消按钮 -->
    <el-form-item>
      <el-button type="primary" @click="onSubmit">立即创建</el-button>
      <el-button @click="closeDrawer">取消</el-button>
    </el-form-item>
  </el-form>
  <!-- 表单结束 -->
</template>

<script>
// 导入请求函数
import {postClient} from '@/service/requests'

export default {
  name: "clientEmail",
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
          robot_url_list: [
            {
              url: "",
            },
          ],
        },
        typeDescription: "电子邮件",
      },
      clientRules: {
        client_name: [
          {required: true, message: "name不能为空", trigger: "blur"},
        ],
      }
    }
  }, props: {
    getClientList: Function,
  },
  methods: {
    // 提交表单函数
    onSubmit: function () {
      postClient(this.$store.getters.getNamespace, this.client).then(response => {
        if (response.data.code === 1) {
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
    // 添加输入框函数
    add: function () {
      let cope = {
        url: "",
      };
      this.client.client_info.robot_url_list.push(cope);
      for (let i = 0; i < this.client.client_info.robot_url_list.length; i++) {
        console.log(this.client.client_info.robot_url_list[i]);
      }
    },
    // 删除输入框函数
    del: function (index) {
      this.client.client_info.robot_url_list.splice(index, 1);
    },
    // 关闭抽屉函数
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
