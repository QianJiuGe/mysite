<template>
  <SiteShell :breadcrumbs="breadcrumbs" @logout="onLogout">
    <section class="hero card">
      <p class="eyebrow">PERSONAL SITE</p>
      <h1>构建可靠、可演进的 Web 工程体验</h1>
      <p class="hero-lead">
        我关注前后端协作、交付效率和技术写作，把工程实践沉淀成可复用的知识资产。
      </p>
      <div class="hero-actions">
        <router-link to="/blog" class="btn-primary">阅读博客</router-link>
        <router-link to="/projects" class="btn-secondary">查看项目</router-link>
      </div>
    </section>

    <section class="section-block">
      <div class="section-title-row">
        <h2>Latest Posts</h2>
        <router-link to="/blog">查看全部</router-link>
      </div>
      <div class="content-grid">
        <article v-for="post in latestPosts" :key="post.slug" class="card post-card">
          <p class="meta-row">{{ post.publishedAt }} · {{ post.readingMinutes }} min read</p>
          <h3>
            <router-link :to="`/blog/${post.slug}`">{{ post.title }}</router-link>
          </h3>
          <p>{{ post.summary }}</p>
          <div class="tag-row">
            <span v-for="tag in post.tags" :key="tag" class="tag">{{ tag }}</span>
          </div>
        </article>
      </div>
    </section>

    <section class="section-block">
      <div class="section-title-row">
        <h2>Featured Projects</h2>
        <router-link to="/projects">项目归档</router-link>
      </div>
      <div class="content-grid">
        <article v-for="project in featuredProjects" :key="project.id" class="card project-card">
          <p class="meta-row">{{ project.status }} · {{ project.role }}</p>
          <h3>
            <router-link :to="`/projects/${project.id}`">{{ project.name }}</router-link>
          </h3>
          <p>{{ project.summary }}</p>
          <div class="tag-row">
            <span v-for="item in project.tech" :key="item" class="tag">{{ item }}</span>
          </div>
        </article>
      </div>
    </section>

    <section class="section-block card stats-card">
      <h2>Tech Snapshot</h2>
      <p v-if="loading" class="muted loading-pulse">同步中...</p>
      <p v-else-if="errorText" class="error">{{ errorText }}</p>
      <div v-else class="snapshot-grid">
        <div>
          <p class="snapshot-label">欢迎语</p>
          <p class="snapshot-value">{{ welcome || '欢迎回来' }}</p>
        </div>
        <div>
          <p class="snapshot-label">当前模块</p>
          <p class="snapshot-value">{{ modules.length > 0 ? modules.join(' / ') : 'Home · Blog · Projects' }}</p>
        </div>
      </div>
    </section>
  </SiteShell>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

import SiteShell from '../components/SiteShell.vue'
import { clearAuth } from '../lib/auth'
import { getHome } from '../lib/api'
import { blogPosts, projectCases } from '../lib/content'

const router = useRouter()
const welcome = ref('')
const modules = ref<string[]>([])
const loading = ref(true)
const errorText = ref('')

const breadcrumbs = [{ label: '首页' }]
const latestPosts = computed(() => blogPosts.slice(0, 3))
const featuredProjects = computed(() => projectCases.slice(0, 3))

const onLogout = async () => {
  clearAuth()
  await router.push('/login')
}

onMounted(async () => {
  try {
    const res = await getHome()
    welcome.value = res.welcome
    modules.value = res.modules
  } catch (e) {
    const err = e as { status?: number; message?: string }
    if (err.status === 401) {
      clearAuth()
      await router.push('/login')
      return
    }
    errorText.value = err.message ?? '加载首页数据失败'
  } finally {
    loading.value = false
  }
})
</script>
