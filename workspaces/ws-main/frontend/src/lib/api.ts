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

export async function register(payload: { username: string; password: string }): Promise<{ userId: number; status: string }> {
  return request('/v1/auth/register', {
    method: 'POST',
    body: JSON.stringify(payload),
  })
}

export async function login(payload: { username: string; password: string }): Promise<{ token: string; role: string }> {
  return request('/v1/auth/login', {
    method: 'POST',
    body: JSON.stringify(payload),
  })
}

export async function getHome(): Promise<{ welcome: string; modules: string[] }> {
  return request('/v1/home', { method: 'GET' }, true)
}

export async function getPendingUsers(): Promise<{ users: Array<{ userId: number; username: string; status: string }> }> {
  return request('/v1/admin/users/pending', { method: 'GET' }, true)
}

export async function approveUser(userId: number): Promise<{ userId: number; status: string }> {
  return request(`/v1/admin/users/${userId}/approve`, { method: 'POST' }, true)
}
