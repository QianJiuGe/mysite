import { getToken } from './auth'

const API_BASE = import.meta.env.VITE_API_BASE ?? 'http://127.0.0.1:8080'

type ApiError = {
  code?: string
  message?: string
}

async function request<T>(path: string, init?: RequestInit, auth = false): Promise<T> {
  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
    ...(init?.headers as Record<string, string> | undefined),
  }
  if (auth) {
    const token = getToken()
    if (token) {
      headers.Authorization = `Bearer ${token}`
    }
  }

  const res = await fetch(`${API_BASE}${path}`, {
    ...init,
    headers,
  })

  const data = (await res.json().catch(() => ({}))) as ApiError & T
  if (!res.ok) {
    const err = new Error(data.message ?? 'request failed') as Error & { code?: string; status?: number }
    err.code = data.code
    err.status = res.status
    throw err
  }
  return data as T
}

// Auth
export async function register(payload: { username: string; password: string }): Promise<{ userId: number; status: string }> {
  return request('/v1/auth/register', { method: 'POST', body: JSON.stringify(payload) })
}

export async function login(payload: { username: string; password: string }): Promise<{ token: string; role: string }> {
  return request('/v1/auth/login', { method: 'POST', body: JSON.stringify(payload) })
}

// Admin
export async function getPendingUsers(): Promise<{ users: Array<{ userId: number; username: string; status: string }> }> {
  return request('/v1/admin/users/pending', { method: 'GET' }, true)
}

export async function approveUser(userId: number): Promise<{ userId: number; status: string }> {
  return request(`/v1/admin/users/${userId}/approve`, { method: 'POST' }, true)
}

// Blog
export type BlogPostSummary = {
  slug: string
  title: string
  summary: string
  tags: string[]
  date: string
}

export type BlogPostDetail = BlogPostSummary & {
  content: string
}

export async function listBlogPosts(): Promise<{ posts: BlogPostSummary[] }> {
  return request('/v1/blog/posts', { method: 'GET' })
}

export async function getBlogPost(slug: string): Promise<BlogPostDetail> {
  return request(`/v1/blog/posts/${slug}`, { method: 'GET' })
}

export async function createBlogPost(slug: string, content: string): Promise<{ slug: string }> {
  return request('/v1/admin/blog/posts', { method: 'POST', body: JSON.stringify({ slug, content }) }, true)
}

export async function updateBlogPost(slug: string, content: string): Promise<{ slug: string }> {
  return request(`/v1/admin/blog/posts/${slug}`, { method: 'PUT', body: JSON.stringify({ content }) }, true)
}

export async function deleteBlogPost(slug: string): Promise<{ slug: string }> {
  return request(`/v1/admin/blog/posts/${slug}`, { method: 'DELETE' }, true)
}

// Memo
export type Memo = {
  id: number
  userId: number
  username?: string
  content: string
  done: boolean
  priority: number
  createdAt: string
  updatedAt: string
}

export async function listMemos(): Promise<{ memos: Memo[] }> {
  return request('/v1/memos', { method: 'GET' }, true)
}

export async function createMemo(content: string, priority = 0): Promise<Memo> {
  return request('/v1/memos', { method: 'POST', body: JSON.stringify({ content, priority }) }, true)
}

export async function updateMemo(id: number, data: { content?: string; done?: boolean; priority?: number }): Promise<{ id: number }> {
  return request(`/v1/memos/${id}`, { method: 'PUT', body: JSON.stringify(data) }, true)
}

export async function deleteMemo(id: number): Promise<{ id: number }> {
  return request(`/v1/memos/${id}`, { method: 'DELETE' }, true)
}
