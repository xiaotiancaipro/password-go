<template>
  <div class="main-wrap">
    <el-form class="login-form" label-width="auto">
      <el-text class="login-title">PASSWORD-GO</el-text>
      <el-input class="username" v-model="username" placeholder="Username"/>
      <el-input class="password" v-model="password" type="password" show-password placeholder="Password"/>
      <el-button class="input-button" color="#409575" @click="login">Login</el-button>
    </el-form>
  </div>
</template>

<script setup lang="ts">

import {onMounted, onUnmounted, ref} from 'vue';
import router from "@/router";
import {ElMessage} from 'element-plus'

const username = ref("");
const password = ref("");
const isLoggedIn = ref(false);

const getUsername = () => {
  return "root";
};

const getPassword = () => {
  return "root";
};

const setLoggedIn = (flag: boolean) => {
  isLoggedIn.value = flag;
};

const login = () => {
  if (username.value === getUsername() && password.value === getPassword()) {
    localStorage.setItem("username", getUsername());
    localStorage.setItem("password", getPassword());
    setLoggedIn(true);
    router.replace("/home");
    ElMessage({
      message: "Login successful",
      type: "success",
    })
  } else {
    ElMessage({
      message: "Login failed, please enter your account and password correctly!",
      type: "error",
    });
  }
};

const keyDown = (e: KeyboardEvent): void => {
  if (e.key === "Enter") login();
};

onMounted(() => {
  window.addEventListener("keydown", keyDown);
});

onUnmounted(() => {
  window.removeEventListener("keydown", keyDown);
});

</script>

<style scoped>

.main-wrap {
  position: fixed;
  width: 100%;
  height: 100%;
  left: 0;
  top: 0;
  background: #FFFFFF;
}

.login-form {
  width: 25rem;
  height: 18.75rem;
  position: absolute;
  top: 30%;
  left: 50%;
  transform: translate(-50%, -50%);
  overflow: hidden;
}

.login-title {
  position: absolute;
  top: 15%;
  left: 48%;
  transform: translate(-50%, -50%);
  line-height: 240%;
  font-size: 1.25rem;
}

.username {
  position: absolute;
  top: 35%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 18.75rem;
  height: 2.1875rem;
  border-radius: 0.25rem;
}

.password {
  position: absolute;
  top: 55%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 18.75rem;
  height: 2.1875rem;
  border-radius: 0.25rem;
}

.input-button {
  position: absolute;
  top: 70%;
  left: 50%;
  transform: translate(-50%, -50%);
  margin: 5% auto 0;
  width: 18.75rem;
  height: 2.1875rem;
  border-radius: 0.25rem;
}

</style>
