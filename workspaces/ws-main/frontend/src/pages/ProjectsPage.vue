<template>
  <SiteShell :breadcrumbs="breadcrumbs" @logout="onLogout">
    <section class="page-heading">
      <h1>项目展示</h1>
      <p>围绕“问题-方案-结果”记录项目实践，关注真实场景下的工程价值。</p>
    </section>

    <section class="content-grid">
      <article v-for="project in projectCases" :key="project.id" class="card project-card">
        <p class="meta-row">{{ project.status }} · {{ project.role }}</p>
        <h2>
          <router-link :to="`/projects/${project.id}`">{{ project.name }}</router-link>
        </h2>
        <p>{{ project.summary }}</p>
        <div class="tag-row">
          <span v-for="item in project.tech" :key="item" class="tag">{{ item }}</span>
        </div>
      </article>
    </section>
  </SiteShell>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

import SiteShell from '../components/SiteShell.vue'
import { clearAuth } from '../lib/auth'
import { projectCases } from '../lib/content'

const router = useRouter()
const breadcrumbs = [
  { label: '首页', to: '/home' },
  { label: '项目' },
]

const onLogout = async () => {
  clearAuth()
  await router.push('/login')
}
</script>
