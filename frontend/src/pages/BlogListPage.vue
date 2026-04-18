<template>
  <AppLayout>
    <div class="blog-list">
      <div class="page-header">
        <div>
          <h2 class="headline-medium">博客</h2>
          <p class="body-medium text-on-surface-variant">记录工程实践与技术思考</p>
        </div>
        <button v-if="isAdmin" class="fab-extended" @click="showNewPost = true">
          <span class="material-symbols-outlined">add</span>
          新建文章
        </button>
      </div>

      <!-- New post form -->
      <div v-if="showNewPost" class="card-outlined new-post-form">
        <h3 class="title-large" style="margin-bottom:16px">新建文章</h3>
        <label>
          Slug（URL 标识，英文、数字、横线）
          <input v-model.trim="newSlug" placeholder="my-first-post" />
        </label>
        <label style="margin-top:16px">
          内容（Markdown + frontmatter）
          <textarea v-model="newContent" rows="10" style="font-family:var(--md-font-mono)" :placeholder="newPostPlaceholder" />
        </label>
        <div class="form-actions">
          <button class="btn-filled" :disabled="!newSlug || !newContent || creating" @click="onCreate">
            <span class="material-symbols-outlined">publish</span>
            {{ creating ? '创建中...' : '创建' }}
          </button>
          <button class="btn-text" @click="showNewPost = false; newSlug = ''; newContent = ''">取消</button>
        </div>
        <p v-if="createError" class="text-error body-small" style="margin-top:8px">{{ createError }}</p>
      </div>

      <div v-if="loading" class="body-medium text-on-surface-variant" style="padding:40px 0">加载中...</div>
      <div v-else-if="errorText" class="text-error body-medium">{{ errorText }}</div>

      <div v-else class="posts-grid">
        <article v-for="post in posts" :key="post.slug" class="post-card card-elevated">
          <router-link :to="`/blog/${post.slug}`" class="post-link">
            <h3 class="title-large post-title">{{ post.title }}</h3>
            <p class="body-medium text-on-surface-variant post-summary">{{ post.summary }}</p>
            <div class="post-footer">
              <span class="label-medium text-on-surface-variant">{{ post.date }}</span>
              <div class="post-tags">
                <span v-for="tag in post.tags" :key="tag" class="chip chip-tonal" style="height:24px;padding:0 10px;font-size:11px">{{ tag }}</span>
              </div>
            </div>
          </router-link>
          <div v-if="isAdmin" class="post-actions">
            <router-link :to="`/blog/${post.slug}`" class="icon-btn" title="编辑">
              <span class="material-symbols-outlined" style="font-size:20px">edit</span>
            </router-link>
            <button class="icon-btn" @click="onDelete(post.slug, post.title)" :disabled="deletingSlug === post.slug" title="删除">
              <span class="material-symbols-outlined" style="font-size:20px;color:var(--md-error)">delete</span>
            </button>
          </div>
        </article>

        <div v-if="posts.length === 0" class="empty-state">
          <span class="material-symbols-outlined" style="font-size:48px;color:var(--md-outline)">draft</span>
          <p class="body-medium text-on-surface-variant">暂无文章</p>
        </div>
      </div>
    </div>
  </AppLayout>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import AppLayout from '../components/AppLayout.vue'
import { listBlogPosts, createBlogPost, deleteBlogPost, type BlogPostSummary } from '../lib/api'
import { getRole } from '../lib/auth'

const isAdmin = computed(() => getRole() === 'admin')
const posts = ref<BlogPostSummary[]>([])
const loading = ref(true)
const errorText = ref('')
const showNewPost = ref(false)
const newSlug = ref('')
const newContent = ref('')
const creating = ref(false)
const createError = ref('')
const deletingSlug = ref('')

const newPostPlaceholder = `---
title: 文章标题
summary: 一句话简介
tags: [技术]
date: ${new Date().toISOString().slice(0, 10)}
---

正文内容...`

const loadPosts = async () => {
  try {
    loading.value = true
    const res = await listBlogPosts()
    posts.value = res.posts
  } catch (e) {
    errorText.value = (e as Error).message ?? '加载失败'
  } finally {
    loading.value = false
  }
}

const onCreate = async () => {
  creating.value = true; createError.value = ''
  try {
    await createBlogPost(newSlug.value, newContent.value)
    showNewPost.value = false; newSlug.value = ''; newContent.value = ''
    await loadPosts()
  } catch (e) { createError.value = (e as Error).message ?? '创建失败' }
  finally { creating.value = false }
}

const onDelete = async (slug: string, title: string) => {
  if (!confirm(`确认删除「${title}」？此操作不可恢复。`)) return
  deletingSlug.value = slug
  try { await deleteBlogPost(slug); await loadPosts() }
  catch (e) { errorText.value = (e as Error).message ?? '删除失败' }
  finally { deletingSlug.value = '' }
}

onMounted(loadPosts)
</script>

<style scoped>
.blog-list { max-width: 800px; }

.page-header {
  display: flex; justify-content: space-between; align-items: flex-start;
  margin-bottom: 32px;
}

.new-post-form {
  margin-bottom: 24px;
}

.form-actions {
  display: flex; gap: 12px; margin-top: 16px;
}

.posts-grid {
  display: flex; flex-direction: column; gap: 12px;
}

.post-card {
  display: flex; align-items: center; padding: 20px 24px;
}

.post-link {
  flex: 1; min-width: 0;
  color: inherit; text-decoration: none;
}

.post-title {
  margin-bottom: 6px; color: var(--md-on-surface);
}

.post-summary {
  margin-bottom: 12px;
  display: -webkit-box; -webkit-line-clamp: 2; -webkit-box-orient: vertical; overflow: hidden;
}

.post-footer {
  display: flex; align-items: center; gap: 12px;
}

.post-tags { display: flex; gap: 6px; }

.post-actions {
  display: flex; gap: 4px; flex-shrink: 0; margin-left: 16px;
}

.empty-state {
  display: flex; flex-direction: column; align-items: center;
  gap: 12px; padding: 64px 0;
}
</style>
