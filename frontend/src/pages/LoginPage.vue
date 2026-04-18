<template>
  <div class="login-shell">
    <div class="aurora one"></div>
    <div class="aurora two"></div>
    <section class="login-grid">
      <aside class="brand-pane">
        <p class="eyebrow">MYSITE USER HUB</p>
        <h1>欢迎回来</h1>
        <p class="lead">先登录，再进入你的个人空间。管理端账号也从这里进入。</p>
        <ul class="feature-list">
          <li>注册后自动进入待审批状态</li>
          <li>管理员审批后即可使用首页能力</li>
          <li>后续将扩展内容管理与站点运营模块</li>
        </ul>
      </aside>

      <section class="auth-panel card-glass">
        <h2>账号登录</h2>
        <form class="auth-form" @submit.prevent="onSubmit">
          <label>
            用户名
            <input v-model.trim="username" autocomplete="username" placeholder="例如 admin" />
          </label>
          <label>
            密码
            <input v-model="password" type="password" autocomplete="current-password" placeholder="请输入密码" />
          </label>
          <button :disabled="loading" type="submit" class="btn-primary">{{ loading ? '登录中...' : '登录' }}</button>
        </form>
        <p v-if="errorText" class="error" style="margin-top: 12px">{{ errorText }}</p>
        <p class="muted" style="margin-top: 16px"><router-link to="/register">还没有账号？去注册</router-link></p>
      </section>
    </section>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

import { login } from '../lib/api'
import { setAuth } from '../lib/auth'

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
    const res = await login({ username: username.value, password: password.value })
    setAuth(res.token, res.role)
    if (res.role === 'admin') {
      await router.push('/admin')
      return
    }
    await router.push('/home')
  } catch (e) {
    const err = e as { code?: string; message?: string }
    if (err.code === 'FORBIDDEN_PENDING_APPROVAL') {
      await router.push('/pending-approval')
      return
    }
    errorText.value = err.message ?? '登录失败'
  } finally {
    loading.value = false
  }
}
</script>
