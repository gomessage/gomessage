<template>
  <el-drawer title="客户端信息查看"
      :destroy-on-close="true"
      :visible.sync="cd_isShow"
      :before-close="cd_closeBeforeFunc">

    <!--钉钉：机器人-->
    <ClientDingtalk v-bind="$attrs" v-if="cd_clientType==='dingtalk'"></ClientDingtalk>
    <!--飞书：机器人-->
    <ClientFeishu v-bind="$attrs" v-else-if="cd_clientType==='feishu'"></ClientFeishu>
    <!--企微：机器人-->
    <ClientWechatRobot v-bind="$attrs" v-else-if="cd_clientType==='wechat_robot'"></ClientWechatRobot>
    <!--企微：应用号-->
    <ClientWechat v-bind="$attrs" v-else-if="cd_clientType==='wechat'"></ClientWechat>
  </el-drawer>
</template>

<script>
import ClientDingtalk from "@/components/clients/clientDingtalk.vue";
import ClientFeishu from "@/components/clients/clientFeishu.vue";
import ClientWechatRobot from "@/components/clients/clientWechatRobot.vue";
import ClientWechat from "@/components/clients/clientWechat.vue";

export default {
  name: "cDrawerOne",
  props: {
    cd_isShow: Boolean, //控制抽屉是否显示
    cd_closeBeforeFunc: Function, //关闭之前的回调函数，这里调用一个父级函数，把状态回写到父级的变量中
    cd_clientType: String, //客户端类型（根据不同的客户端类型，来决定显示哪一个子组件）
  },
  components: {
    ClientDingtalk,
    ClientFeishu,
    ClientWechatRobot,
    ClientWechat,
  },
}
</script>


