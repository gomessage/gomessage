<template>
  <el-drawer
      title="客户端信息查看"
      :destroy-on-close="true"
      :visible.sync="isShow"
      :before-close="CloseBeforeFunc">

    <!--钉钉：机器人-->
    <ClientDingtalk v-bind="$attrs" v-if="OneClientType==='dingtalk'"></ClientDingtalk>

    <!--飞书：机器人-->
    <ClientFeishu v-bind="$attrs" v-else-if="OneClientType==='feishu'"></ClientFeishu>

    <!--企微：机器人-->
    <ClientWechatRobot v-bind="$attrs" v-else-if="OneClientType==='wechat_robot'"></ClientWechatRobot>

    <!--企微：应用号-->
    <ClientWechat v-bind="$attrs" v-else-if="OneClientType==='wechat'"></ClientWechat>

  </el-drawer>
</template>

<script>
import ClientDingtalk from "@/components/clients/clientDingtalk.vue";
import ClientFeishu from "@/components/clients/clientFeishu.vue";
import ClientWechatRobot from "@/components/clients/clientWechatRobot.vue";
import ClientWechat from "@/components/clients/clientWechat.vue";

export default {
  name: "cDrawerAll",
  props: {
    //控制抽屉是否显示
    isShow: Boolean,
    //关闭之前的回调函数，这里调用一个父级函数，把状态回写到父级的变量中
    CloseBeforeFunc: Function,
    //客户端类型（根据不同的客户端类型，来决定显示哪一个子组件）
    OneClientType: String,
  },
  components: {
    ClientDingtalk,
    ClientFeishu,
    ClientWechatRobot,
    ClientWechat,
  }
}
</script>


