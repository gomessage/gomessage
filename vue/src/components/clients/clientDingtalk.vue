<template>
  <el-form
      label-position="right"
      :model="client"
      :rules="clientRules"
      label-width="100px"
      style="text-align: left;padding: 0 20px">

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
    <el-divider content-position="center">机器人</el-divider>

    <!--关键字-->
    <el-form-item label="放行关键字:">
      <el-input
          v-model="client.client_info.robot_keyword"
          placeholder="要和机器人界面设置的放行关键词保持一直"
      ></el-input>
    </el-form-item>

    <!--机器人url列表-->
    <div v-for="(list, index) in client.client_info.robot_url_list" :key="index">
      <el-form-item label="机器人URL:">
        <el-input v-model="list.url" placeholder="从钉钉上粘贴而来的机器人URL地址" style="width: 85%"></el-input>
        <el-button
            circle
            icon="el-icon-delete"
            size="mini"
            style="margin-left: 10px"
            type="danger"
            v-on:click="input_delete(index)"
        ></el-button>
      </el-form-item>
    </div>

    <!--动态增删输入框-->
    <div style="text-align: center">
      <el-button
          icon="el-icon-plus"
          size="mini"
          type="text"
          @click="input_add">
        再添加一个机器人
      </el-button>
    </div>

    <!--最下文的文字提醒-->
    <p id="text_stype">
      此处可以添加多个机器人，推送消息时会从中随机挑选一个URL来使用，可以避免单个机器人消息推送时（每分钟）的次数限制，避免重要报警信息被漏送的可能。
    </p>

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
  name: "clientDingtalk",
  data() {
    return {
      client: {
        client_name: "", //客户端名称
        client_description: "", //客户端描述
        client_type: "dingtalk", //客户端类型，与后端的约定，禁止修改
        is_active: false, //是否激活
        client_info: { //客户端详情
          robot_url_list: [ //机器人URL，这是一个数组，可以存在多个URL地址
            {
              url: "",
            },
          ],
          robot_keyword: "", //放行关键字
        },
        typeDescription: "钉钉·机器人",
      },

      //规则验证
      clientRules: {
        client_name: [ //client_name字段
          {required: true, message: "name不能为空", trigger: "blur"},
        ],
      }
    }
  },
  //接收父级传递进来的属性
  props: {
    cli_GetClientList: Function, //获取客户端列表：用来刷新客户端列表
    cli_OneClientObject: Object, //单个客户端详情：里面包含了单个客户端的所有详情
    cli_OperateType: String, //操作类型：只能是show或create
  },
  methods: {
    //提交表单（创建一个客户端）
    createClient: function () {
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

    //动态添加输入框
    input_add: function () {
      let cope = {
        url: "",
      };
      this.client.client_info.robot_url_list.push(cope);
      for (let i = 0; i < this.client.client_info.robot_url_list.length; i++) {
        console.log(this.client.client_info.robot_url_list[i]);
      }
    },

    //动态删除输入框
    input_delete: function (index) {
      this.client.client_info.robot_url_list.splice(index, 1);
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
#text_stype {
  font-size: 11px;
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  line-height: 1.7;
  color: #C0C4CC;
}
</style>
