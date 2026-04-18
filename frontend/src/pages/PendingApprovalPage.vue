<template>
  <AppLayout>
    <div class="auth-page">
      <div class="auth-card card-elevated">
        <div class="auth-icon">
          <span class="material-symbols-outlined" style="font-size:32px;color:var(--md-tertiary)">hourglass_top</span>
        </div>
        <h2 class="headline-medium" style="text-align:center;margin-bottom:8px">待管理员审批</h2>
        <p class="body-large text-on-surface-variant" style="text-align:center;line-height:1.7;margin-bottom:8px">
          你的账号已注册成功，当前状态：
          <span class="chip chip-warning" style="height:24px;padding:0 10px;font-size:12px">pending</span>
        </p>
        <p class="body-medium text-on-surface-variant" style="text-align:center">请等待管理员审批后再登录。</p>

        <hr class="divider" style="margin:24px 0 16px" />

        <div style="text-align:center">
          <router-link to="/login" class="btn-text">
            <span class="material-symbols-outlined" style="font-size:18px">arrow_back</span>
            返回登录
          </router-link>
        </div>
      </div>
    </div>
  </AppLayout>
</template>

<script setup lang="ts">
import AppLayout from '../components/AppLayout.vue'
</script>

<style scoped>
.auth-page {
  display: flex; justify-content: center;
  padding-top: 64px;
}

.auth-card {
  width: 100%; max-width: 480px;
}

.auth-icon {
  width: 56px; height: 56px; border-radius: var(--md-radius-lg);
  background: var(--md-tertiary-container);
  display: flex; align-items: center; justify-content: center;
  margin: 0 auto 20px;
}
</style>
