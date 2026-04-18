const TOKEN_KEY = 'mysite_token'
const ROLE_KEY = 'mysite_role'

export function getToken(): string {
  return localStorage.getItem(TOKEN_KEY) ?? ''
}

export function getRole(): string {
  return localStorage.getItem(ROLE_KEY) ?? ''
}

export function setAuth(token: string, role: string): void {
  localStorage.setItem(TOKEN_KEY, token)
  localStorage.setItem(ROLE_KEY, role)
}

export function clearAuth(): void {
  localStorage.removeItem(TOKEN_KEY)
  localStorage.removeItem(ROLE_KEY)
}
