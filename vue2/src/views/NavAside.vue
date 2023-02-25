<template>
  <el-menu
      background-color="#f0fcff"
      class="el-menu-vertical-demo"
      default-active="0"
      style="height: 100%"
      @close="handleClose"
      @open="handleOpen"
  >

    <!--导航1-->
    <!--<el-submenu index="3">-->
    <!--    <template slot="title">-->
    <!--        <i class="el-icon-location"></i>-->
    <!--        <span>导航一</span>-->
    <!--    </template>-->
    <!--    <el-menu-item-group>-->
    <!--        <template slot="title">分组一</template>-->
    <!--        <el-menu-item index="1-1">选项1</el-menu-item>-->
    <!--        <el-menu-item index="1-2">选项2</el-menu-item>-->
    <!--    </el-menu-item-group>-->
    <!--    <el-menu-item-group title="分组2">-->
    <!--        <el-menu-item index="1-3">选项3</el-menu-item>-->
    <!--    </el-menu-item-group>-->
    <!--    <el-submenu index="1-4">-->
    <!--        <template slot="title">选项4</template>-->
    <!--        <el-menu-item index="1-4-1">选项1</el-menu-item>-->
    <!--    </el-submenu>-->
    <!--</el-submenu>-->

    <!--导航1-->
    <!--<el-menu-item index="1" style="text-align: left">-->
    <!--    <i class="el-icon-location"></i>-->
    <!--    <span slot="title">Default</span>-->
    <!--</el-menu-item>-->

    <el-menu-item v-for="(oneNs,index) in namespaceList" :key="index" :index="index" style="text-align: left">
      <!--<i class="el-icon-location" v-if="oneNs.name === 'default'"></i>-->
      <!--<i class="el-icon-menu" v-else></i>-->
      <i class="el-icon-menu"></i>
      <span slot="title">{{ oneNs.name }}</span>
    </el-menu-item>

    <!--导航2-->


    <!--&lt;!&ndash;导航3&ndash;&gt;-->
    <!--<el-menu-item index="3" disabled>-->
    <!--    <i class="el-icon-document"></i>-->
    <!--    <span slot="title">导航三</span>-->
    <!--</el-menu-item>-->

    <!--&lt;!&ndash;导航4&ndash;&gt;-->
    <!--<el-menu-item index="4">-->
    <!--    <i class="el-icon-setting"></i>-->
    <!--    <span slot="title">导航四</span>-->
    <!--</el-menu-item>-->
  </el-menu>
</template>

<script>
import {getNamespace} from "@/service/requests";

export default {
  name: "NavAside",
  data() {
    return {
      namespaceList: [
        {"name": "test1"},
        {"name": "prometheus"},
        {"name": "zabbix"},
        {"name": "kubernetes"},
        {"name": "elasticsearch"},
        {"name": "MyRobot"},
      ]
    }
  },
  methods: {
    pullNamespace: function () {
      getNamespace().then(response => {
        if (response.data.code === 1) {
          this.namespaceList = response.data.result;
          // this.$message.success("成功");
        }
      }).catch(err => {
        console.log(err);
      })
    }
  },
  //created生命周期
  created() {
    this.pullNamespace();
  }
}
</script>

<style scoped>
#NavAside-sty {
  height: 100%;
  background-color: aqua;
}
</style>
