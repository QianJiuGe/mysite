import { ref } from 'vue'

const TOKEN_KEY = 'mysite_token'
const ROLE_KEY = 'mysite_role'
const USERNAME_KEY = 'mysite_username'

// Reactive revision counter — incremented on every auth change so Vue
// computed properties that call getToken/getRole/getUsername re-evaluate.
export const authRevision = ref(0)

export function getToken(): string {
  // access revision so Vue tracks the dependency
  void authRevision.value
  return localStorage.getItem(TOKEN_KEY) ?? ''
}

export function getRole(): string {
  void authRevision.value
  return localStorage.getItem(ROLE_KEY) ?? ''
}

export function getUsername(): string {
  void authRevision.value
  return localStorage.getItem(USERNAME_KEY) ?? ''
}

export function setAuth(token: string, role: string, username: string): void {
  localStorage.setItem(TOKEN_KEY, token)
  localStorage.setItem(ROLE_KEY, role)
  localStorage.setItem(USERNAME_KEY, username)
  authRevision.value++
}

export function clearAuth(): void {
  localStorage.removeItem(TOKEN_KEY)
  localStorage.removeItem(ROLE_KEY)
  localStorage.removeItem(USERNAME_KEY)
  authRevision.value++
}
