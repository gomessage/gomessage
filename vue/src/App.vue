<template>
  <div id="app">

    <!--容器布局：全局-->
    <el-container id="MyContainer-Container">

      <!--Header导航栏-->
      <el-header id="MyContainer-Header">
        <NavHeader/>
      </el-header>


      <el-container>
        <!--左侧导航栏-->
        <el-aside id="MyContainer-Aside" width="275px" v-if="showHeader">
          <NavAside/>
        </el-aside>


        <el-main id="MyContainer-Main">
          <!-- 路由匹配到的组件将渲染在这里 -->
          <router-view></router-view>
        </el-main>

      </el-container>

      <!--底部导航栏-->
      <el-footer id="MyContainer-Footer">
        <NavFooter/>
      </el-footer>

    </el-container>


  </div>
</template>

<script>
import NavFooter from "@/views/navs/NavFooter.vue";
import NavAside from "@/views/navs/NavAside.vue";
import NavHeader from "@/views/navs/NavHeader.vue";

export default {
  name: 'app',
  data() {
    return {
      // isRouterAlive: true,
      showHeader: false
    }
  },
  components: {
    NavHeader, NavAside, NavFooter
  },
  //在 src/App.vue 文件中的 provide 函数用于定义可以被所有后代组件注入的依赖。这里的 provide 函数返回一个对象，其中包含一个 reload 方法。这意味着任何子组件都可以通过注入机制访问 reload 方法，而不需要通过属性传递（prop drilling）。
  //在这个上下文中，reload 方法用于重新加载路由视图。当调用 reload 方法时，它会先将 isRouterAlive 设置为 false，然后在下一个事件循环中将其设置回 true。这种方式实际上是在重置 <router-view> 组件，使得路由加载的组件可以重新渲染。这通常用于处理需要在路由级别强制刷新页面内容的情况。
  //通过 provide 提供的 reload 方法，任何子组件都可以直接调用这个方法来触发路由视图的重载，而无需从顶层组件一层层传递方法或状态。这样做可以大大简化组件间的通信和状态管理。
  // provide() {
  //   return {
  //     reload: this.reload,
  //   }
  // },
  methods: {
    // reload() {
    //   this.isRouterAlive = false; //先关闭，
    //   this.$nextTick(function () {
    //     this.isRouterAlive = true; //再打开
    //   });
    // }
  },
  watch: {
    $route: { // 监控路由变化
      immediate: true,
      handler(to) {
        // 根据路由路径判断是否显示Header
        this.showHeader = !['/login',].includes(to.path);
      }
    }
  },
  created() {
    //在页面加载时读取sessionStorage里的状态信息
    if (sessionStorage.getItem("store")) {
      this.$store.replaceState(Object.assign({}, this.$store.state, JSON.parse(sessionStorage.getItem("store"))))
    }
    //在页面刷新时将vuex里的信息保存到sessionStorage里
    window.addEventListener("beforeunload", () => {
      sessionStorage.setItem("store", JSON.stringify(this.$store.state))
    })
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

/*布局容器：全局*/
#MyContainer-Container {
  height: 100vh;
}

/*布局容器：导航栏*/
#MyContainer-Header {
  padding: 0;
  margin: 0;
}

/*布局容器：内容栏*/
#MyContainer-Main {
  padding: 0;
  margin: 0;
  /*定时消息里使用了绝对定位，因此这里需要设定为相对定位。*/
  position: relative;
  //border: 3px solid red;
}

/*布局容器：左侧栏*/
#MyContainer-Aside {
  padding: 0;
  margin: 0;
}

/*布局容器：底部栏*/
#MyContainer-Footer {
  padding: 0;
  margin: 0;
  height: auto !important;
}

</style>
