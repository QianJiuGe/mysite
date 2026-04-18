<template>
  <AppLayout>
    <div class="memo-page">
      <div class="page-header">
        <div>
          <h2 class="headline-medium">备忘录</h2>
          <p class="body-medium text-on-surface-variant">{{ isAdmin ? '管理所有用户的备忘录' : '记录待办事项和笔记' }}</p>
        </div>
        <button v-if="!showInput" class="fab-extended" @click="showInput = true">
          <span class="material-symbols-outlined">add</span>
          添加
        </button>
      </div>

      <!-- Input panel (shown after clicking "添加") -->
      <div v-if="showInput" class="card-outlined memo-input">
        <textarea
          ref="inputRef"
          v-model="newContent"
          placeholder="输入备忘录内容..."
          rows="3"
          @keydown.meta.enter="addMemo"
          @keydown.ctrl.enter="addMemo"
        />
        <div class="input-actions">
          <select v-model="newPriority" class="priority-select">
            <option :value="0">普通</option>
            <option :value="1">重要</option>
            <option :value="2">紧急</option>
          </select>
          <div class="input-btns">
            <button class="btn-text" @click="cancelInput">取消</button>
            <button class="btn-filled" :disabled="!newContent.trim() || adding" @click="addMemo">
              {{ adding ? '添加中...' : '确认添加' }}
            </button>
          </div>
        </div>
      </div>

      <p v-if="errorText" class="text-error body-small" style="margin-bottom:16px">{{ errorText }}</p>
      <div v-if="loading" class="body-medium text-on-surface-variant" style="padding:40px 0;text-align:center">加载中...</div>

      <div v-else class="memo-list">
        <div v-if="memos.length === 0" class="empty-state">
          <span class="material-symbols-outlined" style="font-size:48px;color:var(--md-outline)">task_alt</span>
          <p class="body-medium text-on-surface-variant">暂无备忘录，点击右上角添加一条吧</p>
        </div>

        <div v-for="memo in memos" :key="memo.id" :class="['memo-item card-filled', { done: memo.done }]">
          <button class="icon-btn" @click="toggleDone(memo)" style="flex-shrink:0">
            <span class="material-symbols-outlined" :class="{ 'icon-filled': memo.done }" :style="{ color: memo.done ? 'var(--md-primary)' : 'var(--md-outline)' }">
              {{ memo.done ? 'check_circle' : 'radio_button_unchecked' }}
            </span>
          </button>

          <div class="memo-body">
            <div v-if="editingId === memo.id" class="memo-edit">
              <textarea v-model="editContent" rows="2" />
              <div class="edit-actions">
                <button class="btn-filled" style="height:32px;padding:0 16px;font-size:12px" @click="saveEdit(memo.id)">保存</button>
                <button class="btn-text" style="height:32px;font-size:12px" @click="editingId = null">取消</button>
              </div>
            </div>
            <div v-else>
              <p :class="['body-large', 'memo-content', { strikethrough: memo.done }]">{{ memo.content }}</p>
              <div class="memo-meta">
                <span v-if="isAdmin && memo.username" class="chip chip-tonal" style="height:20px;padding:0 8px;font-size:10px">{{ memo.username }}</span>
                <span :class="['chip', priorityClass(memo.priority)]" style="height:20px;padding:0 8px;font-size:10px">
                  {{ memo.priority === 2 ? '紧急' : memo.priority === 1 ? '重要' : '普通' }}
                </span>
                <span class="label-small text-on-surface-variant">{{ formatDate(memo.createdAt) }}</span>
              </div>
            </div>
          </div>

          <div class="memo-actions">
            <button class="icon-btn" @click="startEdit(memo)" title="编辑">
              <span class="material-symbols-outlined" style="font-size:20px">edit</span>
            </button>
            <button class="icon-btn" @click="removeMemo(memo.id)" title="删除">
              <span class="material-symbols-outlined" style="font-size:20px;color:var(--md-error)">delete</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </AppLayout>
</template>

<script setup lang="ts">
import { computed, nextTick, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import AppLayout from '../components/AppLayout.vue'
import { listMemos, createMemo, updateMemo, deleteMemo, type Memo } from '../lib/api'
import { clearAuth, getToken, getRole } from '../lib/auth'

const router = useRouter()
const isAdmin = computed(() => getRole() === 'admin')
const memos = ref<Memo[]>([])
const loading = ref(true)
const errorText = ref('')
const showInput = ref(false)
const newContent = ref('')
const newPriority = ref(0)
const adding = ref(false)
const editingId = ref<number | null>(null)
const editContent = ref('')
const inputRef = ref<HTMLTextAreaElement | null>(null)

const priorityClass = (p: number) => p === 2 ? 'chip-error' : p === 1 ? 'chip-warning' : 'chip-outlined'

const cancelInput = () => {
  showInput.value = false
  newContent.value = ''
  newPriority.value = 0
}

const loadMemos = async () => {
  if (!getToken()) { clearAuth(); await router.push('/login'); return }
  try {
    loading.value = true
    const res = await listMemos()
    memos.value = res.memos
  } catch (e) {
    const err = e as { status?: number; message?: string }
    if (err.status === 401) { clearAuth(); await router.push('/login'); return }
    errorText.value = err.message ?? '加载失败'
  } finally { loading.value = false }
}

const addMemo = async () => {
  if (!newContent.value.trim()) return
  adding.value = true; errorText.value = ''
  try {
    await createMemo(newContent.value.trim(), newPriority.value)
    newContent.value = ''; newPriority.value = 0
    showInput.value = false
    await loadMemos()
  } catch (e) { errorText.value = (e as Error).message ?? '添加失败' }
  finally { adding.value = false }
}

const toggleDone = async (memo: Memo) => {
  try { await updateMemo(memo.id, { done: !memo.done }); await loadMemos() }
  catch (e) { errorText.value = (e as Error).message ?? '更新失败' }
}

const startEdit = (memo: Memo) => { editingId.value = memo.id; editContent.value = memo.content }

const saveEdit = async (id: number) => {
  try { await updateMemo(id, { content: editContent.value }); editingId.value = null; await loadMemos() }
  catch (e) { errorText.value = (e as Error).message ?? '保存失败' }
}

const removeMemo = async (id: number) => {
  try { await deleteMemo(id); await loadMemos() }
  catch (e) { errorText.value = (e as Error).message ?? '删除失败' }
}

const formatDate = (d: string) =>
  new Date(d).toLocaleDateString('zh-CN', { month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' })

onMounted(loadMemos)
</script>

<style scoped>
.memo-page { max-width: 680px; }

.page-header {
  display: flex; justify-content: space-between; align-items: flex-start;
  margin-bottom: 24px;
}

.memo-input { margin-bottom: 24px; }
.memo-input textarea { margin-bottom: 12px; }

.input-actions {
  display: flex; justify-content: space-between; align-items: center;
}

.input-btns {
  display: flex; gap: 8px;
}

.priority-select {
  padding: 8px 12px; width: auto;
  border: 1px solid var(--md-outline); border-radius: var(--md-radius-xs);
  font: var(--md-body-medium); background: transparent; color: var(--md-on-surface);
}

.memo-list { display: flex; flex-direction: column; gap: 8px; }

.memo-item {
  display: flex; align-items: flex-start; gap: 8px;
  padding: 16px; border-radius: var(--md-radius-lg);
}

.memo-body { flex: 1; min-width: 0; }

.memo-content { word-break: break-word; }
.strikethrough { text-decoration: line-through; color: var(--md-outline); }

.memo-meta { display: flex; gap: 8px; margin-top: 8px; align-items: center; }

.memo-actions { display: flex; gap: 0; flex-shrink: 0; }

.memo-edit textarea { margin-bottom: 8px; }
.edit-actions { display: flex; gap: 8px; }

.icon-filled {
  font-variation-settings: 'FILL' 1;
}

.empty-state {
  display: flex; flex-direction: column; align-items: center;
  gap: 12px; padding: 64px 0;
}
</style>
