<template>
  <div id="app">
    <router-view v-if="isRouterAlive"></router-view>
  </div>
</template>

<script>
export default {
  name: 'app',
  provide() {
    return {
      reload: this.reload,
    }
  },
  data() {
    return {
      isRouterAlive: true
    }
  },
  components: {},
  created() {
    //在页面加载时读取sessionStorage里的状态信息
    if (sessionStorage.getItem("store")) {
      this.$store.replaceState(Object.assign({}, this.$store.state, JSON.parse(sessionStorage.getItem("store"))))
    }
    //在页面刷新时将vuex里的信息保存到sessionStorage里
    window.addEventListener("beforeunload", () => {
      sessionStorage.setItem("store", JSON.stringify(this.$store.state))
    })
  },
  methods: {
    reload() {
      this.isRouterAlive = false; //先关闭，
      this.$nextTick(function () {
        this.isRouterAlive = true; //再打开
      });
    }
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
  /*border: 3px solid red;*/
  /*margin-top: 60px;*/

  /*下面这5个属性，是为了让element-ui绝对占满屏幕，但是使用了布局容器后，就不再需要这个干巴巴的写了*/
  /*position: absolute;*/
  /*top: 0;*/
  /*left: 0;*/
  /*width: 100%;*/
  /*height: 100%;*/
}

/*布局容器：全局*/
#MyContainer-Container {
  height: 100vh;
  /*background-color: red;*/
}

/*布局容器：导航栏*/
#MyContainer-Header {
  padding: 0;
  margin: 0;
  /*border: 3px solid red;*/
}

/*布局容器：内容栏*/
#MyContainer-Main {
  padding: 0;
  margin: 0;
  /*border: 3px solid red;*/
}

#MyContainer-Aside {
  padding: 0;
  margin: 0;
  /*height: 100%;*/
  /*border: 3px solid black;*/
}

/*布局容器：底部栏*/
#MyContainer-Footer {
  padding: 0;
  margin: 0;
  height: auto !important;
  /*background-color: #161823;*/
  /*border: 3px solid red;*/
}

/*全局设定el-row的样式，但是我打算用别的方式实现，这个先不删，留在这里备忘*/
/*.el-row {*/
/*    margin-bottom: 20px;*/

/*    &:last-child {*/
/*        margin-bottom: 0;*/
/*    }*/

/*}*/

/*.el-col {*/
/*    border-radius: 4px;*/
/*}*/

</style>
