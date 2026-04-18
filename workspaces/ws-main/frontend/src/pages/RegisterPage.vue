<template>
  <section class="auth-page card-glass">
    <h1>注册账号</h1>
    <p class="muted">注册成功后需等待管理员审批。</p>
    <form class="auth-form" @submit.prevent="onSubmit">
      <label>
        用户名
        <input v-model.trim="username" autocomplete="username" />
      </label>
      <label>
        密码
        <input v-model="password" type="password" autocomplete="new-password" />
      </label>
      <button :disabled="loading" type="submit" class="btn-primary">{{ loading ? '提交中...' : '提交注册' }}</button>
    </form>
    <p v-if="errorText" class="error" style="margin-top: 12px">{{ errorText }}</p>
    <p style="margin-top: 16px"><router-link to="/login">返回登录</router-link></p>
  </section>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

import { register } from '../lib/api'

const router = useRouter()
const username = ref('')
const password = ref('')
const loading = ref(false)
const errorText = ref('')

const onSubmit = async () => {
  errorText.value = ''
  if (!username.value || !password.value) {
    errorText.value = '请输入用户名和密码'
    return
  }
  try {
    loading.value = true
    await register({ username: username.value, password: password.value })
    await router.push('/pending-approval')
  } catch (e) {
    const err = e as { message?: string }
    errorText.value = err.message ?? '注册失败'
  } finally {
    loading.value = false
  }
}
</script>
