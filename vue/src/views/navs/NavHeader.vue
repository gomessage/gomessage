<template>
  <!--#303133-->
  <!--#000000-->
  <el-menu
    :collapse-transition="true"
    :default-active="$route.name"
    active-text-color="#ffd04b"
    background-color="#161823"
    mode="horizontal"
    router
    style="height: 100%;"
    text-color="#fff"
    @select="handleSelect"
  >
    <!-- background-color="#303133" text-color="#fff" active-text-color="#ffd04b" collapse-transition=true style="height: 100%">-->

    <!--左侧logo-->
    <el-menu-item>
      <router-link to="/">
        <img height="90%" src="../../assets/logo.svg">
        &nbsp;&nbsp;
        <span style="font-size: 15px;">GoMessage · 消息转发器</span>
      </router-link>
    </el-menu-item>

    <el-menu-item style="padding: 0;">
      <router-link to="/">
        <el-button
          type="text"
          icon="el-icon-location"
          style="color: #fff;margin-left: 10px;">
          {{ getStoreNamespace }}
        </el-button>
      </router-link>
    </el-menu-item>


    <!--右侧按钮-->
    <el-row justify="end" type="flex">

      <router-link to="/">
        <el-menu-item index="1-1"><i class="el-icon-position"></i>消息入口</el-menu-item>
      </router-link>


      <router-link to="/request_data">
        <el-menu-item index="2-1"><i class="el-icon-set-up"></i>数据渲染</el-menu-item>
      </router-link>


      <router-link to="/data_client">
        <el-menu-item index="3-1"><i class="el-icon-chat-dot-square"></i>接收客户端</el-menu-item>
      </router-link>

      <!--            <el-link href="https://gitee.com/gomessage/gomessage#gomessage" :underline="false" type="primary" target="_blank">-->
      <!--                <el-menu-item><i class="el-icon-notebook-2"></i>文档</el-menu-item>-->
      <!--            </el-link>-->

      <el-menu-item>
        <el-link :underline="false" href="https://github.com/gomessage/gomessage#gomessage" target="_blank"
                 type="primary"><i class="el-icon-document"></i>文档
        </el-link>
      </el-menu-item>

      <!--            <el-link href="https://gitee.com/gomessage/gomessage#gomessage" :underline="false" type="primary" target="_blank">-->
      <!--                <el-menu-item><i class="el-icon-notebook-2"></i>文档</el-menu-item>-->
      <!--            </el-link>-->

    </el-row>


  </el-menu>
</template>

<script>

export default {
  name: "NavHeader",
  data() {
    return {
      // activeIndex: this.getActiveIndex,
    };
  },
  components: {
    // CSteps,
  },
  computed: {
    // 计算属性：动态获取vuex中的值
    getStoreNamespace: function () {
      return this.$store.getters.getNamespace
    }
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    }
  }
}
</script>

<style scoped>
.router-link-active {
  text-decoration: none;
}
</style>
