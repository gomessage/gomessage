<template>
  <el-row>
    <el-col :span="12" :offset="6" style="margin-top: 20px">
      <el-card>
        <div slot="header" class="clearfix">
          <span>修改密码</span>
        </div>

        <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">

          <el-form-item label="输入旧密码" prop="oldPass">
            <el-input type="password" v-model="ruleForm.oldPass" autocomplete="off"></el-input>
          </el-form-item>

          <el-form-item label="输入新密码" prop="pass">
            <el-input type="password" v-model="ruleForm.pass" autocomplete="off"></el-input>
          </el-form-item>

          <el-form-item label="确认新密码" prop="checkPass">
            <el-input type="password" v-model="ruleForm.checkPass" autocomplete="off"></el-input>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" @click="submitForm('ruleForm')">提交</el-button>
            <el-button @click="resetForm('ruleForm')">重置</el-button>
          </el-form-item>

        </el-form>

      </el-card>
    </el-col>
  </el-row>

</template>


<script>
import {updatePassword} from "@/service/requests";

export default {
  name: "UserProfile",
  data() {
    const checkAge = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('不能为空'));
      }
      // setTimeout(() => {
      //   if (!Number.isInteger(value)) {
      //     callback(new Error('请输入数字值'));
      //   } else {
      //     if (value < 18) {
      //       callback(new Error('必须年满18岁'));
      //     } else {
      //       callback();
      //     }
      //   }
      // }, 1000);
    };
    const validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'));
      } else {
        if (this.ruleForm.checkPass !== '') {
          this.$refs.ruleForm.validateField('checkPass');
        }
        callback();
      }
    };
    const validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'));
      } else if (value !== this.ruleForm.pass) {
        callback(new Error('两次输入密码不一致!'));
      } else {
        callback();
      }
    };
    return {
      ruleForm: {
        oldPass: '',
        pass: '',
        checkPass: '',
      },
      rules: {
        pass: [
          {validator: validatePass, trigger: 'blur'}
        ],
        checkPass: [
          {validator: validatePass2, trigger: 'blur'}
        ],
        oldPass: [
          {validator: checkAge, trigger: 'blur'}
        ]
      }
    };
  },
  methods: {
    submitForm(formName) {

      let pass = {
        "old_password": this.ruleForm.oldPass,
        "new_password": this.ruleForm.pass,
      }

      updatePassword(this.$store.getters.getUserID, pass).then(response => {
        if (response.data.code === 1) {
          if (this.$route.path !== '/login') {
            this.$router.push("/login");
          }
          this.$message.success("密码修改成功，请重新登录...");
          this.$store.commit("updateToken", "");
        }
      }).catch(err => {
        console.log(err);
        console.log(formName);
      })


      // this.$refs[formName].validate((valid) => {
      //   if (valid) {
      //

      //
      //   } else {
      //     console.log('error submit!!');
      //     return false;
      //   }
      // });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    }
  }
}
</script>


<style scoped>

</style>