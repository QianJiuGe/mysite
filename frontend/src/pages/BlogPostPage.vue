<template>
  <SiteShell :breadcrumbs="breadcrumbs" @logout="onLogout">
    <article v-if="post" class="post-layout">
      <div class="post-main card">
        <header class="post-header">
          <h1>{{ post.title }}</h1>
          <p class="meta-row">
            发布于 {{ post.publishedAt }} · 更新于 {{ post.updatedAt }} · {{ post.readingMinutes }} min read
          </p>
          <div class="tag-row">
            <span v-for="tag in post.tags" :key="tag" class="tag">{{ tag }}</span>
          </div>
        </header>

        <section v-for="section in post.sections" :key="section.heading" class="post-section">
          <h2>{{ section.heading }}</h2>
          <p>{{ section.body }}</p>
        </section>

        <footer class="post-foot-nav">
          <router-link v-if="prevPost" :to="`/blog/${prevPost.slug}`">← {{ prevPost.title }}</router-link>
          <span v-else></span>
          <router-link v-if="nextPost" :to="`/blog/${nextPost.slug}`">{{ nextPost.title }} →</router-link>
        </footer>
      </div>

      <aside class="post-toc card">
        <h3>目录</h3>
        <ul>
          <li v-for="section in post.sections" :key="section.heading">{{ section.heading }}</li>
        </ul>
      </aside>
    </article>

    <section v-else class="card empty-state">
      <h1>文章不存在</h1>
      <p class="muted">请返回博客列表重新选择。</p>
      <router-link to="/blog">返回博客列表</router-link>
    </section>
  </SiteShell>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'

import SiteShell from '../components/SiteShell.vue'
import { clearAuth } from '../lib/auth'
import { blogPosts, getBlogPostBySlug } from '../lib/content'

const route = useRoute()
const router = useRouter()

const post = computed(() => getBlogPostBySlug(String(route.params.slug ?? '')))
const currentIndex = computed(() => blogPosts.findIndex((item) => item.slug === post.value?.slug))
const prevPost = computed(() => {
  if (currentIndex.value <= 0) {
    return undefined
  }
  return blogPosts[currentIndex.value - 1]
})
const nextPost = computed(() => {
  if (currentIndex.value < 0 || currentIndex.value >= blogPosts.length - 1) {
    return undefined
  }
  return blogPosts[currentIndex.value + 1]
})

const breadcrumbs = computed(() => {
  if (post.value) {
    return [
      { label: '首页', to: '/home' },
      { label: '博客', to: '/blog' },
      { label: post.value.title },
    ]
  }
  return [
    { label: '首页', to: '/home' },
    { label: '博客', to: '/blog' },
    { label: '文章不存在' },
  ]
})

const onLogout = async () => {
  clearAuth()
  await router.push('/login')
}
</script>
