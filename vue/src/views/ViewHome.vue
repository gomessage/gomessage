<template>
  <div>
    <el-row style="margin-top: 40px" class="shadow">
      <el-col :offset="7" :span="10">
        <Domain></Domain>
      </el-col>
    </el-row>
    <el-row>
      <el-col :offset="7" :span="10" style="margin-top: 150px">
        <el-button round @click="simulation">模拟 AlertManager 推送一条报警信息</el-button>
      </el-col>
    </el-row>
  </div>
</template>

<script>
// @ is an alias to /src
import Domain from "@/components/cDomain";
import {postDemoData} from '@/service/requests'

export default {
  name: 'ViewHome',
  components: {
    Domain,
  }, methods: {
    simulation: function () {
      let jsonData = require('./A2-template.json');
      postDemoData(this.$store.getters.getNamespace, jsonData).then(response => {
        console.log(response)
        this.$message.success("模拟消息推送成功...")
      }).catch(err => {
        console.log(err)
      });
    }
  },
  created() {
    //修改步骤条的值
    this.$store.commit("updateStepsActive", 0);
  }

}
</script>
<style scoped>

</style>
