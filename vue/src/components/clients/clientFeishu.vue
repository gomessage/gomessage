<template>
  <el-form
      label-position="right"
      label-width="100px"
      :model="client"
      :rules="feishuRules"
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
    <el-form-item label="标题名称:">
      <el-input v-model="client.client_info.robot_keyword" placeholder="标题中要包含机器人的放行关键词"></el-input>
    </el-form-item>

    <!--标题颜色-->
    <el-form-item label="标题颜色:">
      <el-select v-model="client.client_info.title_color" placeholder="请选择">
        <el-option v-for="item in all_title_color" :key="item.color" :label="item.label" :value="item.color">
        </el-option>
      </el-select>
    </el-form-item>

    <!--机器人url列表-->
    <div v-for="(list, index) in client.client_info.robot_url_list" :key="index">
      <el-form-item label="机器人URL:">
        <el-input v-model="list.url" placeholder="从飞书上粘贴而来的机器人URL地址" style="width: 85%"></el-input>
        <el-button type="danger" icon="el-icon-delete" circle size="mini" v-on:click="del(index)" style="margin-left: 10px"></el-button>
      </el-form-item>
    </div>

    <!--动态增删输入框-->
    <div style="text-align: center">
      <el-button type="text" icon="el-icon-plus" size="mini" @click="add">再添加一个机器人</el-button>
    </div>

    <!--最下的文字提醒-->
    <p id="textStype">
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
  name: "clientFeishu",
  data() {
    return {
      client: {
        client_name: "",
        client_description: "",
        client_type: "feishu", //客户端类型，与后端的约定，禁止修改
        is_active: false,
        client_info: {
          robot_keyword: "",
          robot_url_list: [
            {
              url: "",
            },
          ],
          title_color: "grey",
        },
        typeDescription: "飞书·机器人",
      },
      feishuRules: {
        client_name: [
          {required: true, message: "name不能为空", trigger: "blur"},
        ],
      },
      all_title_color: [
        {
          color: "grey",
          label: "灰色（默认）"
        },
        {
          color: "blue",
          label: "蓝色"
        },
        {
          color: "wathet",
          label: "浅蓝色"
        },
        {
          color: "turquoise",
          label: "松石绿"
        },
        {
          color: "green",
          label: "绿色"
        },
        {
          color: "yellow",
          label: "黄色"
        },
        {
          color: "orange",
          label: "橘色"
        },
        {
          color: "red",
          label: "红色"
        },
        {
          color: "carmine",
          label: "胭脂红"
        },
        {
          color: "violet",
          label: "紫罗兰色"
        },
        {
          color: "purple",
          label: "紫色"
        },
        {
          color: "indigo",
          label: "靛青"
        },

      ],
    }
  }, props: {
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

    //动态删除输入框
    del: function (index) {
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
#textStype {
  font-size: 11px;
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  line-height: 1.7;
  color: #C0C4CC;
}
</style>
