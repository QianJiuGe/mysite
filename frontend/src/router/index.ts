import { createRouter, createWebHistory } from 'vue-router'

import AdminPage from '../pages/AdminPage.vue'
import BlogListPage from '../pages/BlogListPage.vue'
import BlogPostPage from '../pages/BlogPostPage.vue'
import HomePage from '../pages/HomePage.vue'
import LoginPage from '../pages/LoginPage.vue'
import PendingApprovalPage from '../pages/PendingApprovalPage.vue'
import ProjectDetailPage from '../pages/ProjectDetailPage.vue'
import ProjectsPage from '../pages/ProjectsPage.vue'
import RegisterPage from '../pages/RegisterPage.vue'
import { getRole, getToken } from '../lib/auth'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', redirect: '/login' },
    { path: '/login', component: LoginPage },
    { path: '/register', component: RegisterPage },
    { path: '/pending-approval', component: PendingApprovalPage },
    { path: '/home', component: HomePage, meta: { requiresAuth: true } },
    { path: '/blog', component: BlogListPage, meta: { requiresAuth: true } },
    { path: '/blog/:slug', component: BlogPostPage, meta: { requiresAuth: true } },
    { path: '/projects', component: ProjectsPage, meta: { requiresAuth: true } },
    { path: '/projects/:id', component: ProjectDetailPage, meta: { requiresAuth: true } },
    { path: '/admin', component: AdminPage, meta: { requiresAuth: true, adminOnly: true } },
  ],
})

router.beforeEach((to) => {
  const token = getToken()
  const role = getRole()
  if (to.meta.requiresAuth && !token) {
    return '/login'
  }
  if (to.meta.adminOnly && role !== 'admin') {
    return '/home'
  }
  return true
})

export default router
