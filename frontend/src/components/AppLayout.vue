<template>
  <div class="app-layout">
    <!-- Mobile scrim -->
    <div v-if="mobileOpen" class="nav-scrim" @click="mobileOpen = false" />

    <!-- Navigation Drawer -->
    <aside :class="['nav-drawer', { collapsed: navCollapsed, 'mobile-open': mobileOpen }]">
      <div class="nav-logo">
        <div class="logo-icon">H</div>
        <span v-if="!navCollapsed" class="logo-text">HuangTao.dev</span>
      </div>

      <hr class="divider" />

      <nav class="nav-menu">
        <router-link v-if="isAdmin" to="/admin" :class="['nav-item', { active: isActive('/admin') }]">
          <span class="material-symbols-outlined">dashboard</span>
          <span v-if="!navCollapsed" class="nav-label">控制台</span>
        </router-link>

        <router-link to="/blog" :class="['nav-item', { active: isActive('/blog') }]">
          <span class="material-symbols-outlined">article</span>
          <span v-if="!navCollapsed" class="nav-label">博客</span>
        </router-link>

        <router-link v-if="isLoggedIn" to="/memo" :class="['nav-item', { active: isActive('/memo') }]">
          <span class="material-symbols-outlined">checklist</span>
          <span v-if="!navCollapsed" class="nav-label">备忘录</span>
        </router-link>
        <div v-if="!isLoggedIn" class="nav-item locked">
          <span class="material-symbols-outlined">lock</span>
          <span v-if="!navCollapsed" class="nav-label">备忘录</span>
          <span v-if="!navCollapsed" class="chip chip-outlined" style="height:24px;padding:0 8px;font-size:11px;margin-left:auto">登录</span>
        </div>
      </nav>

      <div class="nav-spacer" />

      <!-- User card at bottom -->
      <div v-if="isLoggedIn && !navCollapsed" class="nav-user-card">
        <div class="user-avatar-sm">{{ username.charAt(0).toUpperCase() }}</div>
        <div class="user-info-sm">
          <span class="user-name-sm">{{ username }}</span>
          <span class="user-role-sm">{{ role === 'admin' ? '管理员' : '用户' }}</span>
        </div>
      </div>
      <div v-if="isLoggedIn && navCollapsed" class="nav-user-card collapsed">
        <div class="user-avatar-sm">{{ username.charAt(0).toUpperCase() }}</div>
      </div>
    </aside>

    <!-- Main area -->
    <div class="main-area">
      <!-- Top App Bar -->
      <header class="top-bar">
        <div class="top-bar-left">
          <button class="icon-btn" @click="toggleNav">
            <span class="material-symbols-outlined">menu</span>
          </button>
          <h1 class="top-bar-title">{{ pageTitle }}</h1>
        </div>
        <div class="top-bar-right">
          <template v-if="isLoggedIn">
            <button class="icon-btn" @click="onLogout" title="退出登录">
              <span class="material-symbols-outlined">logout</span>
            </button>
          </template>
          <template v-else>
            <router-link to="/login" class="btn-filled" style="height:36px;padding:0 20px">
              <span class="material-symbols-outlined" style="font-size:18px">login</span>
              登录
            </router-link>
          </template>
        </div>
      </header>

      <!-- Content -->
      <main class="content-area">
        <slot />
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getToken, getRole, getUsername, clearAuth } from '../lib/auth'

const route = useRoute()
const router = useRouter()
const navCollapsed = ref(false)
const mobileOpen = ref(false)

const isMobile = () => window.innerWidth <= 768

const toggleNav = () => {
  if (isMobile()) {
    mobileOpen.value = !mobileOpen.value
  } else {
    navCollapsed.value = !navCollapsed.value
  }
}

const isLoggedIn = computed(() => !!getToken())
const role = computed(() => getRole())
const isAdmin = computed(() => role.value === 'admin')
const username = computed(() => getUsername() || 'U')

const isActive = (path: string) => route.path.startsWith(path)

const pageTitle = computed(() => {
  const p = route.path
  if (p.startsWith('/admin')) return '控制台'
  if (p.startsWith('/blog')) return '博客'
  if (p.startsWith('/memo')) return '备忘录'
  if (p === '/login') return '登录'
  if (p === '/register') return '注册'
  return ''
})

const onLogout = async () => {
  clearAuth()
  await router.push('/blog')
}

watch(() => route.path, () => {
  mobileOpen.value = false
})
</script>

<style scoped>
.app-layout {
  min-height: 100vh;
  display: flex;
}

/* === Navigation Drawer === */
.nav-drawer {
  width: var(--md-nav-width);
  background: var(--md-surface-container-low);
  display: flex;
  flex-direction: column;
  padding: 16px 12px;
  position: fixed;
  top: 0;
  left: 0;
  height: 100vh;
  z-index: 200;
  transition: width var(--md-duration-medium) var(--md-ease-standard);
  overflow: hidden;
}

.nav-drawer.collapsed {
  width: var(--md-nav-collapsed);
}

.nav-logo {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 8px 12px;
  margin-bottom: 4px;
}

.logo-icon {
  width: 40px;
  height: 40px;
  background: linear-gradient(135deg, var(--md-primary), var(--md-tertiary));
  color: #fff;
  border-radius: var(--md-radius-md);
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: 18px;
  flex-shrink: 0;
}

.logo-text {
  font: var(--md-title-large);
  color: var(--md-on-surface);
  white-space: nowrap;
}

.nav-menu {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 12px;
  height: 56px;
  padding: 0 16px;
  border-radius: var(--md-radius-full);
  font: var(--md-label-large);
  color: var(--md-on-surface-variant);
  text-decoration: none;
  transition: all var(--md-duration-short) var(--md-ease-standard);
  cursor: pointer;
  white-space: nowrap;
}

.nav-item:hover {
  background: rgba(28, 27, 31, 0.08);
  color: var(--md-on-surface);
}

.nav-item.active {
  background: var(--md-secondary-container);
  color: var(--md-on-secondary-container);
}

.nav-item.active .material-symbols-outlined {
  font-variation-settings: 'FILL' 1;
}

.nav-item.locked {
  opacity: 0.5;
  cursor: default;
}
.nav-item.locked:hover {
  background: transparent;
  color: var(--md-on-surface-variant);
}

.nav-label {
  flex: 1;
}

.nav-spacer {
  flex: 1;
}

.nav-user-card {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  border-radius: var(--md-radius-full);
  background: var(--md-surface-container);
}
.nav-user-card.collapsed {
  justify-content: center;
  padding: 12px;
}

.user-avatar-sm {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background: var(--md-primary-container);
  color: var(--md-on-primary-container);
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: 14px;
  flex-shrink: 0;
}

.user-info-sm {
  display: flex;
  flex-direction: column;
  min-width: 0;
}

.user-name-sm {
  font: var(--md-label-large);
  color: var(--md-on-surface);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.user-role-sm {
  font: var(--md-body-small);
  color: var(--md-on-surface-variant);
}

/* === Main Area === */
.main-area {
  flex: 1;
  margin-left: var(--md-nav-width);
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  transition: margin-left var(--md-duration-medium) var(--md-ease-standard);
}

.nav-drawer.collapsed ~ .main-area {
  margin-left: var(--md-nav-collapsed);
}

/* === Top App Bar === */
.top-bar {
  height: var(--md-topbar-height);
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 16px;
  background: var(--md-surface);
  position: sticky;
  top: 0;
  z-index: 100;
}

.top-bar-left {
  display: flex;
  align-items: center;
  gap: 8px;
}

.top-bar-title {
  font: var(--md-title-large);
  color: var(--md-on-surface);
}

.top-bar-right {
  display: flex;
  align-items: center;
  gap: 4px;
}

/* === Content === */
.content-area {
  flex: 1;
  padding: 24px 32px;
  max-width: 960px;
}

/* === Responsive === */
@media (max-width: 1024px) {
  .nav-drawer { width: var(--md-nav-collapsed); }
  .nav-drawer .nav-label,
  .nav-drawer .logo-text,
  .nav-drawer .nav-user-card:not(.collapsed) { display: none; }
  .main-area { margin-left: var(--md-nav-collapsed); }
}

@media (max-width: 768px) {
  .nav-drawer {
    width: var(--md-nav-width);
    transform: translateX(-100%);
    transition: transform var(--md-duration-medium) var(--md-ease-standard);
    box-shadow: none;
  }
  .nav-drawer.mobile-open {
    transform: translateX(0);
    box-shadow: var(--md-elevation-4);
  }
  .nav-drawer .nav-label,
  .nav-drawer .logo-text,
  .nav-drawer .nav-user-card:not(.collapsed) { display: flex; }
  .nav-drawer.mobile-open .nav-label,
  .nav-drawer.mobile-open .logo-text { display: inline; }
  .nav-drawer.mobile-open .nav-user-card:not(.collapsed) { display: flex; }
  .main-area { margin-left: 0; }
  .content-area { padding: 16px; }
}

.nav-scrim {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.32);
  z-index: 199;
}
</style>
