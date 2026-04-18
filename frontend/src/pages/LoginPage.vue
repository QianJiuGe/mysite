<template>
  <AppLayout>
    <div class="auth-page">
      <div class="auth-card card-elevated">
        <div class="auth-icon">
          <span class="material-symbols-outlined" style="font-size:32px;color:var(--md-primary)">login</span>
        </div>
        <h2 class="headline-medium" style="text-align:center;margin-bottom:8px">登录</h2>
        <p class="body-medium text-on-surface-variant" style="text-align:center;margin-bottom:24px">使用账号登录以访问更多功能</p>

        <form class="auth-form" @submit.prevent="onSubmit">
          <label>
            用户名
            <input v-model.trim="username" autocomplete="username" placeholder="请输入用户名" />
          </label>
          <label>
            密码
            <input v-model="password" type="password" autocomplete="current-password" placeholder="请输入密码" />
          </label>
          <button :disabled="loading" type="submit" class="btn-filled" style="width:100%;margin-top:8px">
            {{ loading ? '登录中...' : '登录' }}
          </button>
        </form>

        <p v-if="errorText" class="text-error body-small" style="margin-top:12px;text-align:center">{{ errorText }}</p>

        <hr class="divider" style="margin:20px 0" />

        <p class="body-medium text-on-surface-variant" style="text-align:center">
          还没有账号？
          <router-link to="/register" class="text-primary label-large">去注册</router-link>
        </p>
      </div>
    </div>
  </AppLayout>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import AppLayout from '../components/AppLayout.vue'
import { login } from '../lib/api'
import { setAuth } from '../lib/auth'

const router = useRouter()
const username = ref('')
const password = ref('')
const loading = ref(false)
const errorText = ref('')

const onSubmit = async () => {
  errorText.value = ''
  if (!username.value || !password.value) { errorText.value = '请输入用户名和密码'; return }
  try {
    loading.value = true
    const res = await login({ username: username.value, password: password.value })
    setAuth(res.token, res.role, username.value)
    await router.push(res.role === 'admin' ? '/admin' : '/blog')
  } catch (e) {
    const err = e as { code?: string; message?: string }
    if (err.code === 'FORBIDDEN_PENDING_APPROVAL') { await router.push('/pending-approval'); return }
    errorText.value = err.message ?? '登录失败'
  } finally { loading.value = false }
}
</script>

<style scoped>
.auth-page {
  display: flex; justify-content: center;
  padding-top: 64px;
}

.auth-card {
  width: 100%; max-width: 420px;
}

.auth-icon {
  width: 56px; height: 56px; border-radius: var(--md-radius-lg);
  background: var(--md-primary-container);
  display: flex; align-items: center; justify-content: center;
  margin: 0 auto 20px;
}

.auth-form {
  display: flex; flex-direction: column; gap: 16px;
}
</style>
