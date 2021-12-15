<template>
  <div class="login-container">
    <el-form
      :model="loginForm"
      status-icon
      :rules="rules"
      ref="loginForm"
      label-width="0px"
      class="demo-loginForm"
    >
      <h3 class="title">系统登录</h3>
      <el-form-item prop="username">
        <el-input
          placeholder="用户名"
          autocomplete="on"
          v-model.number="loginForm.username"
        ></el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input
          type="password"
          v-model="loginForm.password"
          autocomplete="off"
          placeholder="密码"
        ></el-input>
      </el-form-item>
      <el-checkbox style="width:100%;" v-model="checked" class="rememberme"
        >记住密码</el-checkbox
      >
      <el-form-item>
        <el-button
          :loading="loading"
          style="width:100%;"
          type="primary"
          @click="submitForm()"
          >登陆</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { login } from "network/login";

import { base64Encode } from "utils/base64";

export default {
  data() {
    return {
      loginForm: {
        username: "",
        password: "",
      },
      rules: {
        username: [
          { required: true, message: "please enter username", trigger: "blur" },
        ],
        password: [
          { required: true, message: "please enter password", trigger: "blur" },
        ],
      },
      checked: false,
      loading: false,
    };
  },
  methods: {
    submitForm() {
      this.$refs.loginForm.validate((valid) => {
        if (valid) {
          this.loading = true;
          let userInfo = {
            username: this.loginForm.username,
            password: this.loginForm.password,
          };

          login(userInfo).then((res) => {
            if (res.status === 200) {
              sessionStorage.setItem("token", res.token);
              
              this.$message({
                message: "登陆成功",
                type: "success",
              });
              this.$router.push("/admin");
              this.loading = false;
            } else {
              this.$message({
                message: res.message,
                type: "error",
              });
              this.loading = false;
            }
          });
        } else {
          // console.log("error submit!!");
          return false;
        }
      });
    },
  },
};
</script>

<style scoped>
.login-container {
  width: 100%;
  height: 100%;
}
.demo-loginForm {
  /* 兼容 chrome 或 safari */
  -webkit-border-radius: 5px;
  border-radius: 5px;
  margin: 180px auto;
  width: 350px;
  padding: 35px 35px 15px;
  background: #fff;
  border: 1px solid #eaeaea;
  box-shadow: 0 0 25px #cac6c6;
}
.title {
  text-align: center;
}
label.el-checkbox.rememberme {
  margin: 0px 0px 15px;
  /* text-align: center; */
}
</style>
