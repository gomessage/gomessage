<template>
  <el-drawer
    title="添加客户端"
    :visible="getDrawerStatus"
    :before-close="handleClose"
    :destroy-on-close="true">
    <el-tabs :tab-position="tabPosition" style="margin: 20px">

      <el-tab-pane label="钉钉">
        <clientDingtalk v-bind:getClientList="getClientList"></clientDingtalk>
      </el-tab-pane>

      <el-tab-pane label="飞书">
        <clientFeishu v-bind:getClientList="getClientList"></clientFeishu>
      </el-tab-pane>

      <el-tab-pane label="企微 · 群机器人">
        <clientWechatRobot v-bind:getClientList="getClientList"></clientWechatRobot>
      </el-tab-pane>

      <el-tab-pane label="企微 · 应用号">
        <clientWechat v-bind:getClientList="getClientList"></clientWechat>
      </el-tab-pane>



      <el-tab-pane label="其它" style="text-align: left">
        <clientOther v-bind:getClientList="getClientList"></clientOther>
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
  }, props: {
    getClientList: Function,
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
    // getClientList2: function () {
    //     this.$parent.getClientList();
    // }
  }
}
</script>

<style scoped>

</style>
