<template>
  <el-form
      label-position="right"
      label-width="100px"
      :model="client"
      :rules="wechatRules"
      style="text-align: left">

    <!--客户端ID：只有show时才会显示-->
    <el-form-item label="客户端ID:" prop="id" v-if="cli_OperateType==='show'">
      <el-input v-model="client.id" disabled></el-input>
    </el-form-item>

    <!--客户端所属通道：只有show时才会显示-->
    <el-form-item label="所属通道:" prop="namespace" v-if="cli_OperateType==='show'">
      <el-input v-model="client.namespace" disabled></el-input>
    </el-form-item>

    <!--客户端名称-->
    <el-form-item label="客户端名称:" prop="client_name">
      <el-input v-model="client.client_name" placeholder=""></el-input>
    </el-form-item>

    <!--客户端描述-->
    <el-form-item label="客户端描述:">
      <el-input v-model="client.client_description" placeholder=""></el-input>
    </el-form-item>

    <!--客户端类型-->
    <el-form-item label="客户端类型:">
      <el-input v-model="client.client_type" disabled></el-input>
    </el-form-item>

    <!--客户端是否激活：只有show时才会显示-->
    <el-form-item label="是否激活:" prop="is_active" v-if="cli_OperateType==='show'">
      <el-radio-group v-model="client.is_active">
        <el-radio :label="true">激活</el-radio>
        <el-radio :label="false">未激活</el-radio>
      </el-radio-group>
    </el-form-item>

    <!--分割线-->
    <el-divider content-position="center">应用号</el-divider>

    <!--企业ID-->
    <el-form-item label="企业ID:">
      <el-input v-model="client.client_info.corp_id" placeholder="请联系企业微信管理员获取"></el-input>
    </el-form-item>

    <!--应用AgentId-->
    <el-form-item label="应用AgentId:">
      <el-input v-model="client.client_info.agent_id" placeholder="请联系企业微信管理员获取"></el-input>
    </el-form-item>

    <!--应用Secret-->
    <el-form-item label="应用Secret:">
      <el-input v-model="client.client_info.secret" placeholder="请联系企业微信管理员获取" show-password></el-input>
    </el-form-item>

    <!--接收用户-->
    <el-form-item label="接收用户:">
      <el-input
          type="textarea"
          :autosize="{ minRows: 4, maxRows: 6}"
          v-model="client.client_info.touser"
          placeholder="可以填写多个用户账号，用 | 分割开 （例如：aaa|bbb）"
      >
      </el-input>
    </el-form-item>

    <br><br>

    <!--如果操作类型为show，则显示"修改"按钮-->
    <el-form-item v-if="this.cli_OperateType==='show'">
      <el-button type="info" @click="updateClient" round>立即修改</el-button>
      <el-button type="danger" @click="deleteClient" round>删除</el-button>
    </el-form-item>

    <!--如果操作类型为create，则显示"创建"按钮-->
    <el-form-item v-else>
      <el-button type="primary" @click="createClient">立即创建</el-button>
      <el-button @click="closeDrawer">取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import {deleteClientOne, postClient, putClientInfoOne} from '@/service/requests'

export default {
  name: "clientWechat",
  data() {
    return {
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
    cli_GetClientList: Function, //获取客户端列表：用来刷新客户端列表
    cli_OneClientObject: Object, //单个客户端详情：里面包含了单个客户端的所有详情
    cli_OperateType: String, //操作类型：只能是show或create
  },
  methods: {
    //提交表单（创建一个客户端）
    createClient: function () {
      let secret = this.client.client_info.secret;
      if (secret.length <= 5) {
        this.$message.error("应用Secret的输入长度不符合要求...");
      } else {
        postClient(this.$store.getters.getNamespace, this.client).then(response => {
          if (response.data.result) {
            this.$message.success("添加成功...");
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
      }
    },

    //修改客户端信息
    updateClient: function () {
      putClientInfoOne(this.$store.getters.getNamespace, this.cli_OneClientObject.id, this.client).then(resp => {
        if (resp.data.code === 1) {
          this.$message.success("数据更新成功")
          this.cli_GetClientList();
          //show客户端与create客户端不同，因此"不需要全局控制"抽屉的显示状态。
          //TODO：这里欠缺一个"自动收回抽屉"的逻辑，明天再继续写
        }
      }).catch(err => {
        console.log(err)
      })
    },

    //删除客户端
    deleteClient: function () {
      deleteClientOne(this.$store.getters.getNamespace, this.cli_OneClientObject.id, null).then(response => {
        if (response.data.code === 1) {
          this.$message.success("删除客户端成功...");
          this.cli_GetClientList();
          //show客户端与create客户端不同，因此"不需要全局控制"抽屉的显示状态。
          //TODO：这里欠缺一个"自动收回抽屉"的逻辑，明天再继续写
        } else {
          this.$message.error("删除失败...");
        }
      }).catch(err => {
        console.log(err);
      });
    },

    //关闭抽屉
    closeDrawer: function () {
      this.cli_GetClientList();
      this.$store.commit("updateDrawerStatus", false);
    },
  },

  created() {
    //如果操作类型为show，则为表单client填充数据，否则就还让client表单保持为空的样子
    if (this.cli_OperateType === "show") {
      this.client = this.cli_OneClientObject;
    }
  },
  watch: {
    //监听pops传入的值是否发生了变化，则刷新表单内容
    cli_OneClientObject: {
      handler(newVal, oldVal) {
        console.log(newVal, oldVal)
        if (this.cli_OperateType === "show") {
          this.client = this.cli_OneClientObject;
        }
      },
      deep: true
    }
  }

}
</script>

<style scoped>

</style>
