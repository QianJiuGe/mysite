<template>
  <div class="page home-shell">
    <header class="home-topbar">
      <div class="brand">MySpace</div>
      <div class="user-slot">
        <span class="welcome">{{ welcome || '已登录用户' }}</span>
        <button class="btn-subtle" @click="$emit('logout')">退出</button>
      </div>
    </header>

    <main class="home-content">
      <aside class="card-solid left-col">
        <h3>个人卡片</h3>
        <p>星标访客：12</p>
        <p>今日访问：36</p>
      </aside>

      <section class="card-solid main-col">
        <h3>动态占位区</h3>
        <p v-if="loading" class="loading-pulse">加载中...</p>
        <p v-if="errorText" class="error">{{ errorText }}</p>
        <div v-for="item in modules" :key="item" class="feed-item">模块：{{ item }}</div>
      </section>

      <aside class="card-solid right-col">
        <h3>公告</h3>
        <p>欢迎来到你的站点控制台。</p>
      </aside>
    </main>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  welcome: string
  modules: string[]
  loading: boolean
  errorText: string
}>()

defineEmits<{
  logout: []
}>()
</script>

<style scoped>
.home-shell {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.home-topbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: var(--space-16) var(--space-24);
  background: var(--color-surface);
  border-bottom: 1px solid var(--color-border);
  position: sticky;
  top: 0;
  z-index: 10;
  backdrop-filter: blur(12px);
}

.brand {
  font-size: var(--font-size-lg);
  font-weight: var(--font-weight-bold);
  letter-spacing: var(--letter-spacing-tight);
}

.user-slot {
  display: flex;
  align-items: center;
  gap: var(--space-12);
}

.welcome {
  color: var(--color-text-secondary);
  font-size: var(--font-size-sm);
}

.home-content {
  flex: 1;
  display: grid;
  grid-template-columns: 220px 1fr 220px;
  gap: var(--space-20);
  padding: var(--space-24);
  max-width: var(--container-max);
  margin: 0 auto;
  width: 100%;
}

.left-col,
.right-col,
.main-col {
  padding: var(--space-20);
  display: grid;
  gap: var(--space-12);
  align-content: start;
}

.left-col h3,
.right-col h3,
.main-col h3 {
  font-size: var(--font-size-md);
  font-weight: var(--font-weight-semibold);
}

.left-col p,
.right-col p {
  color: var(--color-text-secondary);
  font-size: var(--font-size-sm);
}

.feed-item {
  padding: var(--space-12) var(--space-16);
  border: 1px solid var(--color-border);
  border-radius: var(--radius-md);
  font-size: var(--font-size-sm);
  transition:
    border-color var(--transition-fast),
    background var(--transition-fast);
}

.feed-item:hover {
  border-color: var(--color-border-hover);
  background: var(--color-surface-hover);
}

@media (max-width: 980px) {
  .home-content {
    grid-template-columns: 1fr;
  }

  .left-col { order: -1; }
  .right-col { order: 1; }
}
</style>
