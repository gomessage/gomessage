<template>
  <div>
    <!--开启或关闭渲染的开关-->
    <el-row>
      <el-col style="height: 60px;background-color: #cccccc">
        <el-switch
          v-model="thisRenders"
          inactive-text="基础转发模式"
          active-text="高级渲染模式"
          style="height: 100%"
          width="40"
          @change="updateNamespaceIsRenders"
        >
        </el-switch>
      </el-col>
    </el-row>

    <!--渲染相关的功能设置-->
    <el-row
      style="padding: 20px 0;margin-left: 0;margin-right: 0"
      :gutter="20"
      v-loading="getNamespaceIsRenders"
      element-loading-text='每个通道都可以独立开启【高级渲染模式】借用GoMessage的动态算法把过境数据实时【渲染为人类可读】的信息；若不开启此模式，则当前通道只会将数据"原封不动"的送达至目标客户端。'
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
      thisRenders: false,
      dialogVisible: true
    }
  },
  computed: {
    getNamespaceIsRenders: function () {
      // 从vuex中读取当前是否需要开启渲染模式
      // eslint-disable-next-line vue/no-side-effects-in-computed-properties
      return  !this.$store.getters.getNamespaceIsRenders
    }
  },
  components: {
    DataFormat,
    DataMap,
    CTemplate,
  },
  methods: {
    updateNamespaceIsRenders: function () {
      namespace = this.$store.getters.getNamespace

    }
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
