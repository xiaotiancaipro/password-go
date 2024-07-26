<script setup lang="ts">

import {ref, onMounted, onUnmounted} from 'vue';
import router from "@/router";

const username = ref("");
const password = ref("");
const isLoggedIn = ref(false);

const getUsername = () => {
  return "root"
}

const getPassword = () => {
  return "root"
}

const setLoggedIn = (flag: boolean) => {
  isLoggedIn.value = flag
}

const login = () => {
  if (username.value === getUsername() && password.value === getPassword()) {
    localStorage.setItem("username", getUsername());
    localStorage.setItem("password", getPassword());
    setLoggedIn(true);
    router.replace("/home")
  } else {
    alert("登录失败, 请正确输入账号和密码!");
  }
};

const keyDown = (e) => {
  if (e.keyCode === 13 || e.keyCode === 100) {
    login();
  }
};

onMounted(() => {
  window.addEventListener("keydown", keyDown);
});

onUnmounted(() => {
  window.removeEventListener("keydown", keyDown);
});

</script>

<template>
  <div class="main-wrap">
    <form method="post" class="login-form">
      <h1 class="login-title">用户登陆</h1>
      <input class="username" type="text" name="username" placeholder="账号" autocomplete="off" v-model="username"/>
      <input class="password" type="password" name="password" placeholder="密码" v-model="password"/>
      <input class="input-button" type="button" value="登 陆" @click="login" @keyup.enter="keyDown"/>
    </form>
  </div>
</template>

<style scoped>

.main-wrap {
  position: fixed;
  width: 100%;
  height: 100%;
  left: 0;
  top: 0;
  background: linear-gradient(45deg, #facff1, #6db9db);
}

.login-form {
  width: 400px;
  height: 300px;
  background: #fff;
  position: absolute;
  top: 40%;
  left: 50%;
  transform: translate(-50%, -50%);
  overflow: hidden;
  border-radius: 6px;
}

.login-title {
  line-height: 240%;
  text-align: center;
  background: linear-gradient(90deg, rgba(246, 60, 47, 0.6), rgba(128, 58, 242, 0.6));
  color: #fff;
  font-size: 20px;
}

.username {
  position: absolute;
  top: 35%;
  left: 50%;
  transform: translate(-50%, -50%);
  display: block;
  width: 300px;
  height: 35px;
  border: 1px solid #d0d0d0;
  border-radius: 4px;
  outline: none;
  text-indent: 10px;
}

.password {
  position: absolute;
  top: 55%;
  left: 50%;
  transform: translate(-50%, -50%);
  display: block;
  width: 300px;
  height: 35px;
  border: 1px solid #d0d0d0;
  border-radius: 4px;
  outline: none;
  text-indent: 10px;
}

.input-button {
  position: absolute;
  top: 70%;
  left: 50%;
  transform: translate(-50%, -50%);
  display: block;
  margin: 5% auto 0;
  width: 300px;
  height: 35px;
  border: 0;
  border-radius: 4px;
  background: rgba(247, 60, 45, 0.6);
  color: #fff;
  cursor: pointer;
}

</style>
