<template>
  <div class="admin-shell">
    <aside class="admin-nav card-solid">
      <h2>管理端</h2>
      <button :class="['nav-item', { active: section === 'users' } ]" @click="section = 'users'">用户审批</button>
      <button :class="['nav-item', { active: section === 'content' } ]" @click="section = 'content'">内容管理</button>
      <button :class="['nav-item', { active: section === 'site' } ]" @click="section = 'site'">站点设置</button>
      <button class="nav-item danger" @click="onLogout">退出登录</button>
    </aside>

    <main class="admin-main">
      <section class="card-solid" v-if="section === 'users'">
        <div class="row between">
          <h3>待审批用户</h3>
          <button class="btn-primary" @click="loadPending" :disabled="loading">刷新</button>
        </div>
        <p v-if="errorText" class="error" style="margin-top: 12px">{{ errorText }}</p>
        <p v-if="loading" class="muted loading-pulse" style="margin-top: 12px">加载中...</p>
        <p v-else-if="pendingUsers.length === 0" class="muted">当前没有待审批用户。</p>
        <div class="pending-grid" v-else>
          <article class="pending-card" v-for="u in pendingUsers" :key="u.userId">
            <h4>{{ u.username }}</h4>
            <p>ID: {{ u.userId }}</p>
            <p>状态: {{ u.status }}</p>
            <button class="btn-primary" @click="onApprove(u.userId)" :disabled="approvingId === u.userId">
              {{ approvingId === u.userId ? '审批中...' : '同意注册' }}
            </button>
          </article>
        </div>
      </section>

      <section class="card-solid" v-if="section === 'content'">
        <h3>内容管理（预留）</h3>
        <p class="muted">后续用于管理首页模块、文章发布与推荐位。</p>
      </section>

      <section class="card-solid" v-if="section === 'site'">
        <h3>站点设置（预留）</h3>
        <p class="muted">后续用于导航、公告、主题和全局配置管理。</p>
      </section>
    </main>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

import { approveUser, getPendingUsers } from '../lib/api'
import { clearAuth } from '../lib/auth'

type PendingUser = {
  userId: number
  username: string
  status: string
}

const router = useRouter()
const section = ref<'users' | 'content' | 'site'>('users')
const pendingUsers = ref<PendingUser[]>([])
const loading = ref(false)
const approvingId = ref<number | null>(null)
const errorText = ref('')

const loadPending = async () => {
  errorText.value = ''
  try {
    loading.value = true
    const res = await getPendingUsers()
    pendingUsers.value = res.users
  } catch (e) {
    const err = e as { status?: number; message?: string }
    if (err.status === 401 || err.status === 403) {
      clearAuth()
      await router.push('/login')
      return
    }
    errorText.value = err.message ?? '加载待审批用户失败'
  } finally {
    loading.value = false
  }
}

const onApprove = async (userId: number) => {
  errorText.value = ''
  try {
    approvingId.value = userId
    await approveUser(userId)
    await loadPending()
  } catch (e) {
    const err = e as { message?: string }
    errorText.value = err.message ?? '审批失败'
  } finally {
    approvingId.value = null
  }
}

const onLogout = async () => {
  clearAuth()
  await router.push('/login')
}

onMounted(async () => {
  await loadPending()
})
</script>
