<script>
export default {
  name: 'login',
  data() {
    return {
      activeName: 'login',
      user_info: {
        username: "",
        password: "",
        password2: "",
        email: "",
        phone: ""
      },
    }
  }
}
</script>

<template>
  <el-card class="shadow" style="width: 30%;margin-left: 35%;margin-top: 150px">
    <!--标签页-->
    <el-tabs v-model="activeName" tab-position="top">

      <!--标签页1-->
      <el-tab-pane label="登录" name="login">

        <!--登录-->
        <el-form label-position="right" :model="user_info">

          <el-form-item label="账号">
            <el-input v-model="user_info.username"></el-input>
          </el-form-item>

          <el-form-item label="密码">
            <el-input v-model="user_info.password"></el-input>
          </el-form-item>
        </el-form>

      </el-tab-pane>

      <!--标签页2-->
      <el-tab-pane label="注册" name="register">

        <!--登录-->
        <el-form label-position="right" label-width="80px" :model="user_info">

          <el-form-item label="账号">
            <el-input v-model="user_info.username"></el-input>
          </el-form-item>

          <el-form-item label="密码">
            <el-input v-model="user_info.password"></el-input>
          </el-form-item>

          <el-form-item label="密码">
            <el-input v-model="user_info.password2"></el-input>
          </el-form-item>

          <el-form-item label="邮箱">
            <el-input v-model="user_info.email"></el-input>
          </el-form-item>

          <el-form-item label="手机">
            <el-input v-model="user_info.phone"></el-input>
          </el-form-item>
        </el-form>

      </el-tab-pane>
    </el-tabs>
  </el-card>
</template>

<style scoped>
.shadow {
  -webkit-box-shadow: #ccc 0 30px 30px;
  -moz-box-shadow: #ccc 0 30px 30px;
  box-shadow: #ccc 0 30px 30px;
  /*transform: scale(1.01, 1.01);*/
}
</style>
