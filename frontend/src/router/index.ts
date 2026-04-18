import { createRouter, createWebHistory } from 'vue-router'

import { getToken, getRole } from '../lib/auth'

const BlogListPage = () => import('../pages/BlogListPage.vue')
const BlogPostPage = () => import('../pages/BlogPostPage.vue')
const MemoPage = () => import('../pages/MemoPage.vue')
const LoginPage = () => import('../pages/LoginPage.vue')
const RegisterPage = () => import('../pages/RegisterPage.vue')
const PendingApprovalPage = () => import('../pages/PendingApprovalPage.vue')
const AdminPage = () => import('../pages/AdminPage.vue')

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', redirect: '/blog' },
    { path: '/blog', component: BlogListPage },
    { path: '/blog/:slug', component: BlogPostPage },
    { path: '/memo', component: MemoPage, meta: { requiresAuth: true } },
    { path: '/login', component: LoginPage },
    { path: '/register', component: RegisterPage },
    { path: '/pending-approval', component: PendingApprovalPage },
    { path: '/admin', component: AdminPage, meta: { requiresAuth: true, requiresAdmin: true } },
  ],
})

router.beforeEach((to) => {
  if (to.meta.requiresAuth && !getToken()) {
    return '/login'
  }
  if (to.meta.requiresAdmin && getRole() !== 'admin') {
    return '/blog'
  }
})

export default router
