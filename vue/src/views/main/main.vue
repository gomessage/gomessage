<script>

import NavFooter from "@/views/navs/NavFooter.vue";
import NavHeader from "@/views/navs/NavHeader.vue";
import NavAside from "@/views/navs/NavAside.vue";

export default {
  name: 'main',
  components: {
    NavHeader,
    NavAside,
    NavFooter
  },
  created() {
    // //在页面加载时读取sessionStorage里的状态信息
    // if (sessionStorage.getItem("store")) {
    //   this.$store.replaceState(Object.assign({}, this.$store.state, JSON.parse(sessionStorage.getItem("store"))))
    // }
    // //在页面刷新时将vuex里的信息保存到sessionStorage里
    // window.addEventListener("beforeunload", () => {
    //   sessionStorage.setItem("store", JSON.stringify(this.$store.state))
    // })
  }
}

</script>

<template>
  <!--容器布局：全局-->
  <el-container id="MyContainer-Container">

    <!--容器布局：Header-->
    <el-header id="MyContainer-Header">
      <NavHeader/>
    </el-header>


    <el-container>
      <!--容器布局：Aside-->
      <el-aside id="MyContainer-Aside" width="275px">
        <NavAside/>
      </el-aside>


      <!--容器布局：Main-->
      <el-main id="MyContainer-Main">
        <!-- 路由匹配到的组件将渲染在这里 -->
        <router-view></router-view>
      </el-main>

    </el-container>

    <!--容器布局：Footer-->
    <el-footer id="MyContainer-Footer">
      <NavFooter/>
    </el-footer>

  </el-container>
</template>

<style scoped>

</style>
