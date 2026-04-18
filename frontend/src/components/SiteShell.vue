<template>
  <div class="site-shell">
    <header class="site-header">
      <div class="site-header-inner">
        <router-link class="site-brand" to="/home">HuangTao.dev</router-link>
        <nav class="site-nav" aria-label="Main navigation">
          <router-link to="/home" class="site-nav-link">Home</router-link>
          <router-link to="/blog" class="site-nav-link">Blog</router-link>
          <router-link to="/projects" class="site-nav-link">Projects</router-link>
        </nav>
        <button class="btn-subtle" @click="$emit('logout')">退出</button>
      </div>
    </header>

    <main class="site-main container">
      <nav class="breadcrumb" aria-label="Breadcrumb">
        <span v-for="(item, index) in breadcrumbs" :key="item.label + index" class="crumb-item">
          <router-link v-if="item.to" :to="item.to">{{ item.label }}</router-link>
          <span v-else>{{ item.label }}</span>
          <span v-if="index < breadcrumbs.length - 1" class="crumb-sep">/</span>
        </span>
      </nav>
      <slot />
    </main>

    <footer class="site-footer">
      <div class="container">
        <p>&copy; {{ currentYear }} HuangTao.dev &middot; Built with Vue & Go</p>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  breadcrumbs: { label: string; to?: string }[]
}>()

defineEmits<{
  logout: []
}>()

const currentYear = new Date().getFullYear()
</script>
