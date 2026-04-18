<template>
  <AppLayout>
    <div v-if="loading" class="body-medium text-on-surface-variant" style="padding:40px 0">加载中...</div>
    <div v-else-if="errorText" class="text-error body-medium">{{ errorText }}</div>
    <div v-else-if="post" class="blog-post">
      <div class="post-header">
        <router-link to="/blog" class="btn-text" style="margin-left:-12px;margin-bottom:8px">
          <span class="material-symbols-outlined" style="font-size:18px">arrow_back</span>
          返回列表
        </router-link>
        <div class="header-row">
          <h1 class="display-small">{{ post.title }}</h1>
          <button v-if="isAdmin && !editing" class="btn-tonal" @click="startEdit">
            <span class="material-symbols-outlined" style="font-size:18px">edit</span>
            编辑
          </button>
        </div>
        <div class="post-meta">
          <span class="label-medium text-on-surface-variant">{{ post.date }}</span>
          <span v-for="tag in post.tags" :key="tag" class="chip chip-tonal" style="height:24px;padding:0 10px;font-size:11px">{{ tag }}</span>
        </div>
      </div>

      <div v-if="editing" class="edit-area card-outlined">
        <textarea v-model="editContent" class="edit-textarea" />
        <div class="edit-actions">
          <button class="btn-filled" :disabled="saving" @click="saveEdit">
            <span class="material-symbols-outlined" style="font-size:18px">save</span>
            {{ saving ? '保存中...' : '保存' }}
          </button>
          <button class="btn-text" @click="cancelEdit">取消</button>
        </div>
        <p v-if="saveError" class="text-error body-small" style="margin-top:8px">{{ saveError }}</p>
      </div>

      <div v-else class="markdown-body" v-html="renderedContent" />
    </div>
    <div v-else class="empty-state">
      <span class="material-symbols-outlined" style="font-size:48px;color:var(--md-outline)">search_off</span>
      <h2 class="title-large">文章不存在</h2>
      <router-link to="/blog" class="btn-text">返回博客列表</router-link>
    </div>
  </AppLayout>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { marked } from 'marked'
import AppLayout from '../components/AppLayout.vue'
import { getBlogPost, updateBlogPost, type BlogPostDetail } from '../lib/api'
import { getRole } from '../lib/auth'

const route = useRoute()
const post = ref<BlogPostDetail | null>(null)
const loading = ref(true)
const errorText = ref('')

const isAdmin = computed(() => getRole() === 'admin')
const editing = ref(false)
const editContent = ref('')
const saving = ref(false)
const saveError = ref('')

const renderedContent = computed(() => {
  if (!post.value) return ''
  return marked(post.value.content)
})

const loadPost = async (slug: string) => {
  loading.value = true; errorText.value = ''; post.value = null
  try { post.value = await getBlogPost(slug) }
  catch (e) { errorText.value = (e as Error).message ?? '加载失败' }
  finally { loading.value = false }
}

const startEdit = () => {
  if (!post.value) return
  const fm = `---\ntitle: ${post.value.title}\nsummary: ${post.value.summary}\ntags: [${post.value.tags.join(', ')}]\ndate: ${post.value.date}\n---\n\n`
  editContent.value = fm + post.value.content
  editing.value = true
}

const cancelEdit = () => { editing.value = false; editContent.value = ''; saveError.value = '' }

const saveEdit = async () => {
  if (!post.value) return
  saving.value = true; saveError.value = ''
  try {
    await updateBlogPost(post.value.slug, editContent.value)
    await loadPost(post.value.slug)
    editing.value = false; editContent.value = ''
  } catch (e) { saveError.value = (e as Error).message ?? '保存失败' }
  finally { saving.value = false }
}

onMounted(() => loadPost(String(route.params.slug)))
watch(() => route.params.slug, (slug) => { if (slug) loadPost(String(slug)) })
</script>

<style scoped>
.blog-post { max-width: 760px; }

.post-header { margin-bottom: 32px; }

.header-row {
  display: flex; justify-content: space-between; align-items: flex-start;
  gap: 16px; margin-bottom: 12px;
}

.post-meta {
  display: flex; align-items: center; gap: 8px;
}

.edit-area { margin-bottom: 32px; }

.edit-textarea {
  width: 100%; min-height: 400px;
  font-family: var(--md-font-mono);
  font-size: 14px; line-height: 1.7;
  padding: 16px; margin-bottom: 12px;
}

.edit-actions { display: flex; gap: 12px; }

.empty-state {
  display: flex; flex-direction: column; align-items: center;
  gap: 12px; padding: 80px 0;
  color: var(--md-on-surface-variant);
}
</style>
