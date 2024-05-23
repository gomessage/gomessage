<template>
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
    <el-menu-item style="text-decoration: none;border-bottom: none;">
      <router-link to="/">
        <img height="90%" src="../../assets/logo.svg">
        &nbsp;&nbsp;
        <span style="font-size: 15px;">GoMessage · 消息转发器</span>
      </router-link>
    </el-menu-item>

    <el-menu-item style="padding: 0;text-decoration: none;border-bottom: none;">
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

      <router-link to="/main/">
        <el-menu-item index="1-1"><i class="el-icon-position"></i>消息入口</el-menu-item>
      </router-link>


      <router-link to="/renders">
        <el-menu-item index="2-1"><i class="el-icon-set-up"></i>数据渲染</el-menu-item>
      </router-link>


      <router-link to="/clients">
        <el-menu-item index="3-1"><i class="el-icon-chat-dot-square"></i>接收客户端</el-menu-item>
      </router-link>

      <!--      <router-link to="/crontab">-->
      <!--        <el-menu-item index="4-1"><i class="el-icon-alarm-clock"></i>定时消息</el-menu-item>-->
      <!--      </router-link>-->

      <!--      <el-submenu index="4">-->
      <!--        <template slot="title"><i class="el-icon-s-operation"></i>其它功能</template>-->

      <!--        <router-link to="/main/crontab">-->
      <!--          <el-menu-item index="4-1"><i class="el-icon-alarm-clock"></i>定时消息</el-menu-item>-->
      <!--        </router-link>-->

      <!--        <router-link to="/main/clients">-->
      <!--          <el-menu-item index="4-2"><i class="el-icon-s-data"></i>数据统计</el-menu-item>-->
      <!--        </router-link>-->

      <!--        <router-link to="/main/clients">-->
      <!--          <el-menu-item index="4-3"><i class="el-icon-user"></i>人工语音</el-menu-item>-->
      <!--        </router-link>-->

      <!--        &lt;!&ndash;后台管理&ndash;&gt;-->
      <!--        <router-link to="/main/clients">-->
      <!--          <el-menu-item index="4-44"><i class="el-icon-setting"></i>后台管理</el-menu-item>-->
      <!--        </router-link>-->
      <!--      </el-submenu>-->


      <!--指向语雀的文档地址-->
      <el-menu-item>
        <el-link :underline="false" href="https://www.yuque.com/osoc/gomessage" target="_blank" type="primary">
          <i class="el-icon-document"></i>文档
        </el-link>
      </el-menu-item>


      <!--用户头像，下拉菜单-->
      <el-dropdown>

        <!--动态显示头像-->
        <el-menu-item>
          <el-avatar :src="image001" v-if="isToken"></el-avatar>
          <el-avatar v-else>User</el-avatar>
        </el-menu-item>

        <!--下拉菜单内容-->
        <el-dropdown-menu>

          <!--用户配置-->
          <el-dropdown-item v-if="isToken">
            <router-link to="/user">
              <el-button type="text">个人信息</el-button>
            </router-link>
          </el-dropdown-item>


          <!--如果存在token，就代表已经登录了，只显示退出登录按钮-->
          <el-dropdown-item v-if="isToken">
            <el-button type="text" @click="user_logout">退出登录</el-button>
          </el-dropdown-item>


          <!--如果不存在token，则显示"用户登录"按钮-->
          <!--          <el-dropdown-item v-else>-->
          <!--            <el-button type="text" @click="router2login">用户登录</el-button>-->
          <!--          </el-dropdown-item>-->

          <el-dropdown-item v-if="!isToken">
            <el-button type="text" @click="router2login">用户登录</el-button>
          </el-dropdown-item>


        </el-dropdown-menu>

      </el-dropdown>
    </el-row>


  </el-menu>
</template>

<script>
import image002 from "@/assets/image001.jpeg"
import {logout} from "@/service/requests";

export default {
  name: "NavHeader",
  data() {
    return {
      // activeIndex: this.getActiveIndex,
      image001: image002,
    };
  },
  computed: {
    // 计算属性：动态获取vuex中的值
    getStoreNamespace: function () {
      return this.$store.getters.getNamespace
    },
    isToken: function () {
      return this.$store.getters.getToken !== "";
    }
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    },
    user_logout() {
      logout({"demo": "demo"}).then(resp => {
        if (resp.data.code === 1) {
          if (this.$route.path !== '/login') {
            this.$router.push("/login");
          }
          this.$message.success("注销成功...");
          this.$store.commit("updateToken", "");
          // this.$nextTick(() => {
          // 确保所有DOM更新完成后执行路由跳转
          //
          //
          // });
        } else {
          this.$message.error("退出登录失败...")
        }
      })
    },
    router2login: function () {
      if (this.$route.path !== '/login') {
        this.$router.push("/login").catch((err) => {
          if (err.name !== 'NavigationDuplicated') {
            throw err;
          }
        })
      } else {
        console.log("已经在登录页面");
      }
    },
  }
}
</script>

<style scoped>
.router-link-active {
  text-decoration: none;
}
</style>
