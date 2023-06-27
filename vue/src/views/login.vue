<template>
  <div style="width: 50%;margin-left: 25%;margin-top: 100px;height: 500px">
    <h3>GoMessage 用户登录</h3>
    <el-divider></el-divider>

    <el-row>
      <el-col :span="12" :offset="6">
        <el-form label-position="right" :model="user_info">
          <el-form-item label="账号">
            <el-input v-model="user_info.username"></el-input>
          </el-form-item>
          <el-form-item label="密码">
            <el-input v-model="user_info.password" type="password" @keyup.enter.native="onSubmit" show-password></el-input>
          </el-form-item>
          <br>
          <el-form-item>
            <el-button type="primary" @click="onSubmit" style="width: 150px">登录</el-button>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>

  </div>
</template>

<script>
import {login} from "@/service/requests";

export default {
  name: 'login',
  data() {
    return {
      activeName: 'login',
      user_info: {
        username: "",
        password: "",
      },
      imageUrl: "../assets/image22222.png"
    }
  },
  methods: {
    onSubmit: function () {

      //TODO：后端逻辑还没处理好，暂时先注释掉
      // 先注销一下登录（这里不管返回的是什么，不管注销成功与否都交给后端来处理，前端不用管）
      // logout({"demo": "demo"}).then(resp => {
      //   console.log(resp)
      // })

      // 登录
      login(this.user_info).then(resp => {
        if (resp.data.code === 1) {

          this.$store.commit("updateToken", resp.data.result.token);
          this.$router.push("/main/")

        } else if (resp.data.code === 0) {
          this.$message.error("登录失败")
        }
      })
    },

    //参数n为休眠时间，单位为毫秒:
    sleep: function (n) {
      let start = new Date().getTime();
      // eslint-disable-next-line no-constant-condition
      while (true) {
        if (new Date().getTime() - start > n) {
          break;
        }
      }
    },
  },
}
</script>

<style scoped>
.shadow {
  -webkit-box-shadow: #ccc 0 30px 30px;
  -moz-box-shadow: #ccc 0 30px 30px;
  box-shadow: #ccc 0 30px 30px;
  /*transform: scale(1.01, 1.01);*/
}
</style>
