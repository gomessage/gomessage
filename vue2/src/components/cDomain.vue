<template>
    <el-card class="box-card" shadow="always" style="width: 100%" id="yinying">
        <div slot="header">
            <span style="padding-left: 80px">消息推送地址</span>
            <el-button style="float: right; padding: 3px 0;" type="text" v-on:click="copyCode()">一键复制</el-button>
        </div>

        <div>
            <!--下面这一行是代码框，不能换行；如果换行第二行的空格就显示在代码框里面了-->
            <pre id="DomainUrlStyle"><code id="DomainUrlContent">{{ myDomain }}{{ showNamespace }}</code></pre>
        </div>
        <!--        <ul v-for="o in 4" :key="o" class="item">-->
        <!--            <li>{{'列表内容 ' + o }}</li>-->
        <!--        </ul>-->

    </el-card>

</template>

<script>
export default {
  name: "cDomain",
  data() {
    return {
      myDomain: "",
    }
  },
  computed: {
    // 计算属性：动态获取vuex中的值
    getStoreNamespace: function () {
      return this.$store.getters.getNamespace
    },
    showNamespace: function () {
      if (this.getStoreNamespace === "default") {
        return "message"
      } else {
        return this.getStoreNamespace
      }
    }
  },
  methods: {
    getDomain: function () {
      let domain = window.location.href;
      let newDomain = domain.split("#/")[0];
      console.log(newDomain);
      // let namespace = ;
      this.myDomain = newDomain + "go/";
    },
    //一键粘贴（纯原生js document对象实现）
    copyCode: function () {
      const val = document.getElementById('DomainUrlContent');
      window.getSelection().selectAllChildren(val);
      document.execCommand("Copy");
      this.$message('复制成功...');
    },
  },
  created() {
    this.getDomain();
  }
}
</script>

<style scoped>
#DomainUrlStyle {
    /*text-align: left;*/
    background-color: #2f2f35;
    /*font-family: Monaco,  Arial, Helvetica, sans-serif;*/
    font-family: Menlo, Monaco, Consolas, "Andale Mono WT", "Andale Mono", "Lucida Console", "Lucida Sans Typewriter", "DejaVu Sans Mono", "Bitstream Vera Sans Mono", "Liberation Mono", "Nimbus Mono L", "Courier New", Courier, monospace;
    font-size: 20px;
    line-height: 50px;
    color: violet;
    overflow-x: auto;
    /*min-height: 20rem;*/
}

#DomainUrlContent {

}

.item {
    margin-top: 18px;
    text-align: left;
}

#yinying {
    -webkit-box-shadow: #ccc 0px 20px 20px;
    -moz-box-shadow: #ccc 0px 20px 20px;
    box-shadow: #ccc 0px 20px 20px;
    /*transform: scale(1.01, 1.01);*/
}
</style>
