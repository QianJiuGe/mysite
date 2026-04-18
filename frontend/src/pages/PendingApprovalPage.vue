<template>
  <section class="auth-page card-glass">
    <h1>待管理员审批</h1>
    <p style="margin-top: 8px; color: var(--color-text-secondary); line-height: 1.7">
      你的账号已注册成功，当前状态：<strong style="color: var(--color-accent)">pending</strong>。
    </p>
    <p style="color: var(--color-text-secondary)">请等待管理员审批后再登录。</p>
    <p style="margin-top: 16px"><router-link to="/login">返回登录</router-link></p>
  </section>
</template>
