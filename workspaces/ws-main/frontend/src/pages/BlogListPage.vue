<template>
  <SiteShell :breadcrumbs="breadcrumbs" @logout="onLogout">
    <section class="page-heading">
      <h1>技术博客</h1>
      <p>记录工程实践、架构思考与交付经验，持续沉淀可复用的方法论。</p>
    </section>

    <section class="filter-bar card">
      <span class="filter-label">标签：</span>
      <button
        v-for="tag in tags"
        :key="tag"
        :class="['tag-filter', { active: selectedTag === tag }]"
        @click="selectedTag = selectedTag === tag ? '' : tag"
      >
        {{ tag }}
      </button>
    </section>

    <section class="list-stack">
      <article v-for="post in visiblePosts" :key="post.slug" class="card article-card">
        <p class="meta-row">{{ post.publishedAt }} · 更新于 {{ post.updatedAt }} · {{ post.readingMinutes }} min read</p>
        <h2>
          <router-link :to="`/blog/${post.slug}`">{{ post.title }}</router-link>
        </h2>
        <p>{{ post.summary }}</p>
        <div class="tag-row">
          <span v-for="tag in post.tags" :key="tag" class="tag">{{ tag }}</span>
        </div>
      </article>
    </section>
  </SiteShell>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router'

import SiteShell from '../components/SiteShell.vue'
import { clearAuth } from '../lib/auth'
import { blogPosts } from '../lib/content'

const router = useRouter()
const selectedTag = ref('')
const breadcrumbs = [
  { label: '首页', to: '/home' },
  { label: '博客' },
]

const tags = computed(() => {
  const set = new Set<string>()
  for (const post of blogPosts) {
    for (const tag of post.tags) {
      set.add(tag)
    }
  }
  return Array.from(set)
})

const visiblePosts = computed(() => {
  if (!selectedTag.value) {
    return blogPosts
  }
  return blogPosts.filter((post) => post.tags.includes(selectedTag.value))
})

const onLogout = async () => {
  clearAuth()
  await router.push('/login')
}
</script>
