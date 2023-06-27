<template>
  <el-drawer
      title="添加客户端"
      :visible="getDrawerStatus"
      :before-close="handleClose"
      :destroy-on-close="true"
      size="40%">
    <el-tabs :tab-position="tabPosition" style="margin: 1px 20px">

      <el-tab-pane label="钉钉 · 群机器人">
        <clientDingtalk v-bind="$attrs"></clientDingtalk>
      </el-tab-pane>

      <el-tab-pane label="飞书 · 群机器人">
        <!--<clientFeishu v-bind:getClientList="getClientList"></clientFeishu>-->
        <clientFeishu v-bind="$attrs"></clientFeishu>
      </el-tab-pane>

      <el-tab-pane label="企微 · 群机器人">
        <!--<clientWechatRobot v-bind:getClientList="getClientList"></clientWechatRobot>-->
        <clientWechatRobot v-bind="$attrs"></clientWechatRobot>
      </el-tab-pane>

      <el-tab-pane label="企微 · 应用号">
        <!--<clientWechat v-bind:getClientList="getClientList"></clientWechat>-->
        <clientWechat v-bind="$attrs"></clientWechat>
      </el-tab-pane>


      <el-tab-pane label="其它" style="text-align: left">
        <!--<clientOther v-bind:getClientList="getClientList"></clientOther>-->
        <clientOther v-bind="$attrs"></clientOther>
      </el-tab-pane>

    </el-tabs>
  </el-drawer>
</template>

<script>
import clientDingtalk from "./clients/clientDingtalk.vue";
import clientWechat from "./clients/clientWechat.vue";
import clientOther from "./clients/clientOther.vue";
import clientFeishu from "./clients/clientFeishu.vue";
import clientWechatRobot from "@/components/clients/clientWechatRobot.vue";

export default {
  name: "cDrawer",
  data() {
    return {
      tabPosition: "top",
    }
  },
  components: {
    clientDingtalk,
    clientWechat,
    clientOther,
    clientFeishu,
    clientWechatRobot,
  },
  computed: {
    getDrawerStatus: function () {
      return this.$store.state.DrawerStatus;
    }
  },
  methods: {
    handleClose(done) {
      this.$confirm('确认关闭？').then(_ => {
        console.log(_);
        this.$store.commit("updateDrawerStatus", false);
        done();
      }).catch(_ => {
        console.log(_);
      });
    },
  }
}
</script>


