<template>
  <div>
    <!--开启或关闭渲染的开关-->
    <el-row>
      <el-col style="height: 60px;background-color: #cccccc">
        <el-switch
          v-model="renders"
          inactive-text="渲染功能 · 关闭"
          active-text="渲染功能 · 打开"
          style="height: 100%"
          width="40"
        >
        </el-switch>
      </el-col>
    </el-row>

    <!--渲染相关的功能设置-->
    <el-row
      style="padding: 20px 0;margin-left: 0;margin-right: 0"
      :gutter="20"
      v-loading="isRenders"
      element-loading-text='当前通道已关闭【上层渲染模式】，仅保留【下层转发功能】继续工作，过境数据将"原封不动的"送达至目标客户端，可在"数据分析页面"查看转发记录与统计！'
      element-loading-spinner="el-icon-info"
      element-loading-background="rgba(0, 0, 0, 0.9)"
    >

      <!--数据劫持与格式化、用户变量映射-->
      <el-col span="12">
        <DataFormat class="shadow"></DataFormat>
        <br>
        <DataMap class="shadow"></DataMap>
      </el-col>

      <!--消息模板-->
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
  data() {
    return {
      renders: false,
      dialogVisible: true
    }
  },
  computed: {
    isRenders: function () {
      return !this.renders
    }
  },
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
