<template>
  <div>

    <!--单独的一行，控制转发模式的开关-->
    <!--<el-row>-->
    <!--  &lt;!&ndash;开关&ndash;&gt;-->
    <!--  <el-col style="height: 60px;background-color: #cccccc">-->
    <!--    <el-switch-->
    <!--      v-model="thisRenders"-->
    <!--      inactive-text="关闭渲染"-->
    <!--      active-text="开启渲染"-->
    <!--      style="height: 100%"-->
    <!--      :width="40"-->
    <!--      @change="updateNamespaceRenders"-->
    <!--    ></el-switch>-->
    <!--  </el-col>-->
    <!--</el-row>-->

    <!--渲染相关的功能设置-->
    <el-row
      style="padding: 20px 0;margin-left: 0;margin-right: 0"
      :gutter="20"
      v-loading="dialogVisible"
      element-loading-text='启用【渲染功能】会把过境数据"渲染为人类可读"的信息；若不开启则只会将数据"原封不动"的送达至目标客户端！'
      element-loading-spinner="el-icon-info"
      element-loading-background="rgba(0, 0, 0, 0.95)"
    >

      <el-col :span="12">
        <!--劫持数据-->
        <DataFormat class="shadow"></DataFormat>
        <br>
        <!--变量映射-->
        <template v-if="configMapType===false">
          <!--用户填写变量-->
          <DataMap class="shadow"></DataMap>
        </template>

        <template v-else>
          <!--后端自动扁平化展开-->
          <DataMap2 class="shadow"></DataMap2>
        </template>
      </el-col>

      <el-col :span="12">
        <!--渲染模板-->
        <CTemplate class="shadow"></CTemplate>
      </el-col>

    </el-row>


  </div>
</template>

<script>
import DataFormat from "@/components/cCodeFormat";
import DataMap from "@/components/cConfigMap";
import DataMap2 from "@/components/cConfigMap2";
import CTemplate from "@/components/cTemplate";
import {getNamespaceOne, putNamespaceOne} from "@/service/requests";

export default {
  name: "ViewRequestData",
  data() {
    return {
      thisRenders: true, //是否渲染的开关样式
      dialogVisible: false, //遮罩层是否显示
      configMapType: false
    }
  },
  components: {
    DataFormat,
    DataMap,
    CTemplate,
    DataMap2,
  },
  computed: {
    //计算属性
    getThisRenders: function () {
      return this.$store.getters.getNamespaceInfo["is_renders"]
    }
  },
  methods: {
    // 获取通道是否开启渲染模式
    getNamespaceRenders: function () {
      let namespaceInfo = this.$store.getters.getNamespaceInfo
      getNamespaceOne(namespaceInfo.id, null).then(rsp => {
        this.thisRenders = rsp.data.result.is_renders;
        this.$store.commit("updateNamespaceInfo", rsp.data.result);
      })
    },

    // 更新通道的渲染模式
    updateNamespaceRenders: function () {
      // 获取一个namespace的结构体数据（这些数据之前被缓存在vuex中）
      let namespaceInfo = this.$store.getters.getNamespaceInfo
      namespaceInfo["is_renders"] = this.thisRenders

      // 提交数据到后端，修改数据状态
      putNamespaceOne(namespaceInfo.id, namespaceInfo).then(rsp => {
        this.thisRenders = rsp.data.result["is_renders"]
        this.$store.commit("updateNamespaceInfo", rsp.data.result)
      })
    }
  },
  watch: {
    // 如果thisRenders的值发生变化，则遮罩层的值也将会跟随变化
    thisRenders: {
      immediate: true,
      handler(newVal, oldVal) {
        console.log(newVal, oldVal);
        this.dialogVisible = !newVal
      }
    },
    // getThisRenders: {
    //   immediate: true,
    //   handler(newVal, oldVal) {
    //     console.log(newVal, oldVal)
    //     this.thisRenders = newVal
    //   }
    // }
  },
  created() {
    this.getNamespaceRenders()
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
