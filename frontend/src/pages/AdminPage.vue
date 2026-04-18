<template>
  <AppLayout>
    <div class="admin-page">
      <div class="page-header">
        <h2 class="headline-medium">控制台</h2>
        <p class="body-medium text-on-surface-variant">管理用户、内容和站点配置</p>
      </div>

      <!-- Tab bar -->
      <div class="tab-bar">
        <button v-for="tab in tabs" :key="tab.key" :class="['tab-item', { active: section === tab.key }]" @click="section = tab.key">
          <span class="material-symbols-outlined" style="font-size:20px">{{ tab.icon }}</span>
          {{ tab.label }}
        </button>
      </div>

      <!-- Users section -->
      <div v-if="section === 'users'" class="section">
        <div class="card-elevated">
          <div class="section-top">
            <div>
              <h3 class="title-large">待审批用户</h3>
              <p class="body-small text-on-surface-variant" style="margin-top:4px">审核新注册用户的访问权限</p>
            </div>
            <button class="btn-tonal" @click="loadPending" :disabled="loading" style="height:36px;padding:0 16px">
              <span class="material-symbols-outlined" style="font-size:18px">refresh</span>
              刷新
            </button>
          </div>

          <p v-if="errorText" class="text-error body-small" style="margin-top:16px">{{ errorText }}</p>

          <div v-if="loading" style="padding:32px 0;text-align:center">
            <span class="body-medium text-on-surface-variant">加载中...</span>
          </div>

          <div v-else-if="pendingUsers.length === 0" class="empty-card">
            <span class="material-symbols-outlined" style="font-size:40px;color:var(--md-outline)">verified</span>
            <p class="body-medium text-on-surface-variant">当前没有待审批用户</p>
          </div>

          <div v-else class="user-list">
            <div v-for="u in pendingUsers" :key="u.userId" class="user-row">
              <div class="user-avatar">{{ u.username.charAt(0).toUpperCase() }}</div>
              <div class="user-info">
                <span class="label-large">{{ u.username }}</span>
                <span class="body-small text-on-surface-variant font-mono">ID: {{ u.userId }}</span>
              </div>
              <span class="chip chip-warning" style="height:24px;padding:0 10px;font-size:11px">{{ u.status }}</span>
              <button class="btn-filled" style="height:32px;padding:0 16px;font-size:12px" @click="onApprove(u.userId)" :disabled="approvingId === u.userId">
                {{ approvingId === u.userId ? '处理中' : '通过' }}
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Site settings section -->
      <div v-if="section === 'site'" class="section">
        <div class="card-elevated">
          <div class="section-top">
            <div>
              <h3 class="title-large">站点设置</h3>
              <p class="body-small text-on-surface-variant" style="margin-top:4px">管理导航、公告和全局配置</p>
            </div>
          </div>
          <div class="empty-card">
            <span class="material-symbols-outlined" style="font-size:40px;color:var(--md-outline)">construction</span>
            <p class="body-medium text-on-surface-variant">即将推出</p>
          </div>
        </div>
      </div>
    </div>
  </AppLayout>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import AppLayout from '../components/AppLayout.vue'
import { approveUser, getPendingUsers } from '../lib/api'
import { clearAuth } from '../lib/auth'

const tabs = [
  { key: 'users', label: '用户审批', icon: 'group' },
  { key: 'site', label: '站点设置', icon: 'settings' },
]

type PendingUser = { userId: number; username: string; status: string }

const router = useRouter()
const section = ref('users')
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
    if (err.status === 401 || err.status === 403) { clearAuth(); await router.push('/login'); return }
    errorText.value = err.message ?? '加载待审批用户失败'
  } finally { loading.value = false }
}

const onApprove = async (userId: number) => {
  errorText.value = ''
  try { approvingId.value = userId; await approveUser(userId); await loadPending() }
  catch (e) { errorText.value = (e as { message?: string }).message ?? '审批失败' }
  finally { approvingId.value = null }
}

onMounted(loadPending)
</script>

<style scoped>
.admin-page { max-width: 760px; }
.page-header { margin-bottom: 24px; }

.tab-bar {
  display: flex; gap: 8px;
  margin-bottom: 24px;
}

.tab-item {
  display: flex; align-items: center; gap: 8px;
  height: 40px; padding: 0 20px;
  border: 1px solid var(--md-outline-variant); border-radius: var(--md-radius-full);
  background: transparent;
  font: var(--md-label-large); color: var(--md-on-surface-variant);
  cursor: pointer;
  transition: all var(--md-duration-short) var(--md-ease-standard);
}
.tab-item:hover { background: rgba(28,27,31,0.08); }
.tab-item.active {
  background: var(--md-secondary-container); border-color: transparent;
  color: var(--md-on-secondary-container);
}
.tab-item.active .material-symbols-outlined { font-variation-settings: 'FILL' 1; }

.section-top {
  display: flex; justify-content: space-between; align-items: flex-start;
  margin-bottom: 20px;
}

.empty-card {
  display: flex; flex-direction: column; align-items: center;
  gap: 12px; padding: 40px 0;
}

.user-list { display: flex; flex-direction: column; gap: 8px; }

.user-row {
  display: flex; align-items: center; gap: 12px;
  padding: 12px 16px;
  background: var(--md-surface-container); border-radius: var(--md-radius-md);
  transition: background var(--md-duration-short) var(--md-ease-standard);
}
.user-row:hover { background: var(--md-surface-container-high); }

.user-avatar {
  width: 40px; height: 40px; border-radius: 50%;
  background: var(--md-primary-container); color: var(--md-on-primary-container);
  display: flex; align-items: center; justify-content: center;
  font-weight: 600; font-size: 16px; flex-shrink: 0;
}

.user-info {
  flex: 1; display: flex; flex-direction: column; min-width: 0;
}
</style>
