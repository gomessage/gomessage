<template>
  <!--#161823-->
  <!--这里的背景色修改完成之后，要在对地下的#NavFooter-Footer中修改一下线条背景色，不然底部会出现一条白线-->
  <el-menu id="NavFooter-Footer" class="el-menu-demo" mode="horizontal" @select="handleSelect"
           background-color="#161823" text-color="#fff" active-text-color="#ffd04b" :collapse-transition="true">
    <el-row type="flex" justify="center" id="my-row">
      <el-menu-item id="NavFooter-menu-item" disabled>
        <el-link href="https://github.com/gomessage/gomessage" type="primary">@ 2020~2023&nbsp;&nbsp;</el-link>
        <el-link href="https://opensource.org/licenses/MIT" type="primary">MIT License&nbsp;&nbsp;&nbsp;</el-link>
        <el-link href="https://github.com/gomessage/gomessage" type="primary">Github&nbsp;</el-link>
        or
        <el-link href="https://gitee.com/gomessage/gomessage" type="primary">&nbsp;Gitee&nbsp;&nbsp;</el-link>
        Version:
        <el-link href="https://github.com/gomessage/gomessage/releases" type="primary">v{{ version }}&nbsp;&nbsp;</el-link>
        Give feedback and Suggestions to the author:
        <el-link href="mailto:taycc3223@gmail.com" type="primary">Email</el-link>
      </el-menu-item>
    </el-row>
  </el-menu>
</template>

<script>
export default {
  name: "NavFooter",
  data() {
    return {
      version: "",
      activeIndex: '1',
      activeIndex2: '1'
    };
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    }
  },
  created() {
    this.version = process.env.VUE_APP_VERSION;
  }
}
</script>

<style scoped>

#NavFooter-Footer {
  padding: 0;
  margin: 0;
  /*下面这个颜色要和Footer的背景色保持一致，不然会在底部出现一条白线*/
  border-bottom-color: #161823;
  /*margin-bottom: 10px;*/
  /*background-color: #161823;*/
}
</style>
