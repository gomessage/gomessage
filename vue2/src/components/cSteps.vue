<template>
  <div style="padding-top: 20px;padding-bottom: 20px">
    <el-steps :active="getActive" finish-status="success" process-status="success" space="30%" align-center="true">
      <el-step title="进入数据" @click.native="stepRouter('/')"></el-step>
      <el-step title="创建模板变量" @click.native="stepRouter('/request_data')"></el-step>
      <el-step title="编写模板内容" @click.native="stepRouter('/data_template')"></el-step>
      <el-step title="设定接收客户端" @click.native="stepRouter('/data_client')"></el-step>
    </el-steps>
  </div>
</template>

<script>
export default {
  name: "cSteps",
  data() {
    return {
      active: 0,
    }
  },
  computed: {
    getActive: function () {
      console.log("StepsActive当前值为：" + this.$store.state.StepsActive);
      return this.$store.state.StepsActive;
    }
  },
  methods: {
    stepRouter: function (routerPath) {
      this.$router.push(routerPath);
    }
  },


// created() {
//     this.getActive();
// }

// methods: {
//     next() {
//         if (this.active++ > 2) this.active = 0;
//     }
// }
}
</script>

<style scoped>

</style>
