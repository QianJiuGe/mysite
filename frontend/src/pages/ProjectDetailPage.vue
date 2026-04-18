<template>
  <SiteShell :breadcrumbs="breadcrumbs" @logout="onLogout">
    <article v-if="project" class="detail-stack">
      <section class="card detail-head">
        <p class="meta-row">{{ project.status }} · {{ project.role }}</p>
        <h1>{{ project.name }}</h1>
        <p>{{ project.summary }}</p>
        <div class="tag-row">
          <span v-for="item in project.tech" :key="item" class="tag">{{ item }}</span>
        </div>
      </section>

      <section class="card detail-section">
        <h2>背景问题</h2>
        <p>{{ project.challenge }}</p>
      </section>

      <section class="card detail-section">
        <h2>技术方案</h2>
        <p>{{ project.solution }}</p>
      </section>

      <section class="card detail-section">
        <h2>结果与复盘</h2>
        <p>{{ project.impact }}</p>
      </section>

      <section class="card detail-section">
        <h2>相关链接</h2>
        <div class="link-row">
          <a v-for="link in project.links" :key="link.label" :href="link.href">{{ link.label }}</a>
        </div>
      </section>
    </article>

    <section v-else class="card empty-state">
      <h1>项目不存在</h1>
      <p class="muted">请返回项目列表重新选择。</p>
      <router-link to="/projects">返回项目列表</router-link>
    </section>
  </SiteShell>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'

import SiteShell from '../components/SiteShell.vue'
import { clearAuth } from '../lib/auth'
import { getProjectById } from '../lib/content'

const route = useRoute()
const router = useRouter()

const project = computed(() => getProjectById(String(route.params.id ?? '')))
const breadcrumbs = computed(() => {
  if (project.value) {
    return [
      { label: '首页', to: '/home' },
      { label: '项目', to: '/projects' },
      { label: project.value.name },
    ]
  }
  return [
    { label: '首页', to: '/home' },
    { label: '项目', to: '/projects' },
    { label: '项目不存在' },
  ]
})

const onLogout = async () => {
  clearAuth()
  await router.push('/login')
}
</script>
