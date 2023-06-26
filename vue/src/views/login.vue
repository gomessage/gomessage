<template>
  <el-card class="shadow" style="width: 50%;margin-left: 25%;margin-top: 100px;height: 600px">

    <el-row :gutter="20">
      <el-col :span="10">
        <img src="@/assets/image002.webp" alt="" width="350px">
      </el-col>

      <el-col :span="12">

        <h3 style="margin-top: 100px;margin-left: 0;padding-left: 0">GoMessage登录</h3>

        <br>

        <!--登录-->
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
  </el-card>
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

      login(this.user_info).then(resp => {
        if (resp.data.code === 1) {
          let token = resp.data.result.token;
          console.log(token)
          // this.$message.success("登录成功")
          // this.sleep(1000)
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
