<template>
  <AppLayout>
    <div class="auth-page">
      <div class="auth-card card-elevated">
        <div class="auth-icon">
          <span class="material-symbols-outlined" style="font-size:32px;color:var(--md-primary)">person_add</span>
        </div>
        <h2 class="headline-medium" style="text-align:center;margin-bottom:8px">注册账号</h2>
        <p class="body-medium text-on-surface-variant" style="text-align:center;margin-bottom:24px">注册成功后需等待管理员审批</p>

        <form class="auth-form" @submit.prevent="onSubmit">
          <label>
            用户名
            <input v-model.trim="username" autocomplete="username" placeholder="请输入用户名" />
          </label>
          <label>
            密码
            <input v-model="password" type="password" autocomplete="new-password" placeholder="请输入密码" />
          </label>
          <button :disabled="loading" type="submit" class="btn-filled" style="width:100%;margin-top:8px">
            {{ loading ? '提交中...' : '提交注册' }}
          </button>
        </form>

        <p v-if="errorText" class="text-error body-small" style="margin-top:12px;text-align:center">{{ errorText }}</p>

        <hr class="divider" style="margin:20px 0" />

        <p class="body-medium text-on-surface-variant" style="text-align:center">
          已有账号？
          <router-link to="/login" class="text-primary label-large">去登录</router-link>
        </p>
      </div>
    </div>
  </AppLayout>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import AppLayout from '../components/AppLayout.vue'
import { register } from '../lib/api'

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
    await register({ username: username.value, password: password.value })
    await router.push('/pending-approval')
  } catch (e) { errorText.value = (e as { message?: string }).message ?? '注册失败' }
  finally { loading.value = false }
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
