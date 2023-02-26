<template>
  <el-row style="margin-top: 20px;margin-bottom: 20px;">
    <el-col :offset="4" :span="16">
      <!--右侧抽屉：添加客户端-->
      <CDrawer v-bind:getClientList="getClientList"></CDrawer>
      <!--右侧抽屉：显示客户端详情-->
      <CDrawerOneDataInfo v-bind:oneClientInfo="clientOneInfo" v-bind:thisClose="thisClose" v-bind:visibleStatus="visibleStatus"></CDrawerOneDataInfo>

      <!--中间卡片：客户端列表-->
      <el-card style="box-shadow: #ccc 0px 10px 10px;">
        <div slot="header" class="clearfix">
          <span style="padding-left: 50px">客户端 · 列表</span>
          <el-button style="float: right; padding: 3px 0" type="text" v-on:click="getClientList">刷新列表</el-button>
        </div>
        <!-- ------------------------------------------------------- -->

        <!--Table表格-->
        <el-table
            :border="true"
            :data="clientList"
            :header-cell-style="{background:'#2f2f35',color:'#fff'}"
            :stripe="true"
            style="width: 100%"
        >

          <el-table-column label="客户端名称" prop="client_name"></el-table-column>
          <el-table-column label="客户端描述" prop="client_description"></el-table-column>
          <el-table-column label="客户端类型" prop="client_annotation"></el-table-column>

          <el-table-column label="是否激活" prop="is_active">
            <template slot-scope="scope">
              <el-checkbox v-model="scope.row.is_active">激活</el-checkbox>
            </template>
          </el-table-column>

          <el-table-column fixed="right" label="操作" width="100">
            <template slot-scope="scope">
              <el-button size="small" type="text" @click.native.prevent="getOneClient(scope.$index, clientList)">详情</el-button>
              <el-button size="small" type="text" @click.native.prevent="deleteOneClient(scope.$index, clientList)">移除</el-button>
            </template>
          </el-table-column>

        </el-table>

        <br>
        <br>
        <br>


        <el-button round type="info" v-on:click="saveActiveStatus">保存激活状态</el-button>
        <el-button round type="primary" v-on:click="addClient">添加新客户端</el-button>
      </el-card>


    </el-col>
  </el-row>

</template>

<script>
import CDrawer from "@/components/cDrawer";
import CDrawerOneDataInfo from "@/components/cDrawerOneDataInfo";
import {deleteClientOne, getClient, getClientOne, putClientOne} from '@/service/requests'

export default {
  name: "ViewClient",
  data() {
    return {
      visibleStatus: false, //传递给子组件用的变量，用来控制是否显示OneDataInfo的抽屉


      clientList: [
        {
          id: 1,
          client_name: "示例客户端1",
          client_description: "示例数据，随时可删~",
          is_active: false,
          client_type: "dingtalk",
          client_annotation: "钉钉·机器人",
        },
        {
          id: 2,
          client_name: "示例客户端2",
          client_description: "示例数据，随时可删~",
          is_active: false,
          client_type: "wechat",
          client_annotation: "企业微信·应用号",
        },
        {
          id: 3,
          client_name: "示例客户端3",
          client_description: "示例数据，随时可删~",
          is_active: true,
          client_type: "feishu",
          client_annotation: "飞书·机器人",
        },
      ],

      clientOneInfo: {  //传递给子组件用的变量，传过去一个json对象，那边会自动展开来显示的
        id: null,
        client_name: null,
        client_description: null,
        is_active: null,
        client_type: "dingtalk",
        client_annotation: "",
        client_info: {}
      },
    }
  },
  components: {
    CDrawer,
    CDrawerOneDataInfo,
  },
  methods: {
    //根据客户端类型，获取对应的客户端注解名称，转换成人类可读的样子
    getAnnotation: function (client_type) {
      let type_map = {
        "dingtalk": "钉钉·机器人",
        "wechat": "企业微信·应用号",
        "feishu": "飞书·机器人",
      }
      return type_map[client_type]
    },

    //保存激活状态
    saveActiveStatus: function () {
      this.clientList.forEach(client => {
        putClientOne(this.$store.getters.getNamespace, client.id, client).then(response => {
          console.log(response.data);
        }).catch(err => {
          console.log(err);
        });
      });
      this.$message.success("激活状态保存成功...");
    },

    //添加客户端：打开抽屉的动作
    addClient: function () {
      this.$store.commit("updateDrawerStatus", true);
    },

    //关闭抽屉
    thisClose: function () {
      this.visibleStatus = false;
    },

    // 删除一行数据：跟后端交互，然后刷新表格
    deleteOneClient(index, rows) {
      let id = rows[index].id;
      deleteClientOne(this.$store.getters.getNamespace, id, null).then(response => {
        if (response.data.code === 1) {
          this.$message.success("删除一行数据成功...");
          rows.splice(index, 1);
        } else {
          this.$message.error("删除数据失败...");
        }
      }).catch(err => {
        console.log(err);
      });
    },


    //获取单个客户端的详情
    getOneClient: function (index, rows) {
      let id = rows[index].id;
      getClientOne(this.$store.getters.getNamespace, id, null).then(response => {
        console.log(response.data.result);
        this.clientOneInfo = response.data.result;
      }).catch(err => {
        console.log(err);
      });
      this.visibleStatus = true;
    },

    //拉取客户端列表
    getClientList: function () {
      getClient(this.$store.getters.getNamespace, null).then(response => {
        if (response.data.result.length === 0) {
          console.log("数据库里没有数据");
          this.$message.error("数据库中没有数据，显示Demo数据...");
        } else {
          this.clientList = response.data.result;
          this.clientList.forEach(client => {
            client.client_annotation = this.getAnnotation(client.client_type)
          })
        }
      }).catch(err => {
        console.log(err);
      });
    },


  },
  created() {
    //修改步骤条的值
    this.$store.commit("updateStepsActive", 3);
    //拉取客户端列表
    this.getClientList();
  }
}
</script>

<style scoped>
#ClientContent {
  width: 60%;
  margin-left: 20%;
}


#subButton {
  width: 100%;
}


.ClientBoxCard {
  /*width: 45rem;*/
  /* margin-top: 1.5rem; */
  /* margin-left: 1.5rem; */
}

.ClientGridContent {
  /*圆角*/
  border-radius: 4px;
  /*最小高度*/
  min-height: 36px;
  margin-bottom: 20px;
}
</style>
