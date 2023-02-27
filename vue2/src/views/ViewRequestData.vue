<template>
  <div>
    <el-row style="margin: 20px 0" :gutter="20">

      <el-col span="12">
        <DataFormat class="shadow"></DataFormat>
        <br>
        <DataMap class="shadow"></DataMap>
      </el-col>

      <el-col :span="12">
        <CTemplate class="shadow"></CTemplate>
      </el-col>
    </el-row>

  </div>
</template>

<script>
import DataFormat from "@/components/cCodeFormat";
import DataMap from "@/components/cConfigMap";
import CTemplate from "@/components/cTemplate";

export default {
  name: "ViewRequestData",
  components: {
    DataFormat,
    DataMap,
    CTemplate,
  },
  created() {
    //修改步骤条的值
    this.$store.commit("updateStepsActive", 1);
  }
}
</script>

<style scoped>
.shadow {
  -webkit-box-shadow: #ccc 0px 30px 30px;
  -moz-box-shadow: #ccc 0px 30px 30px;
  box-shadow: #ccc 0px 30px 30px;
  /*transform: scale(1.01, 1.01);*/
}
</style>
