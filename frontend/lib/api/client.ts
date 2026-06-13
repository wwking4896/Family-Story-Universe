import type { AuthResult, Character, Child, Family, Region, Story } from './types';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL ?? 'http://localhost:8080/api/v1';

type ListResponse<T> = { items: T[] };

async function request<T>(path: string, init: RequestInit = {}): Promise<T> {
  const response = await fetch(`${API_BASE_URL}${path}`, {
    ...init,
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
      ...init.headers,
    },
  });

  if (!response.ok) {
    throw new Error(`API request failed: ${response.status}`);
  }

  return response.json() as Promise<T>;
}

export const fairyCastleApi = {
  register: (payload: { email: string; password: string; display_name: string }) =>
    request<AuthResult>('/auth/register', { method: 'POST', body: JSON.stringify(payload) }),
  login: (payload: { email: string; password: string }) =>
    request<AuthResult>('/auth/login', { method: 'POST', body: JSON.stringify(payload) }),
  createFamily: (token: string, payload: { name: string }) =>
    request<Family>('/families', { method: 'POST', headers: authHeader(token), body: JSON.stringify(payload) }),
  listChildren: (token: string, familyId: number) =>
    request<ListResponse<Child>>(`/children?family_id=${familyId}`, { headers: authHeader(token) }),
  createChild: (token: string, payload: Partial<Child>) =>
    request<Child>('/children', { method: 'POST', headers: authHeader(token), body: JSON.stringify(payload) }),
  listCharacters: (token: string, familyId: number) =>
    request<ListResponse<Character>>(`/characters?family_id=${familyId}`, { headers: authHeader(token) }),
  regions: () => request<ListResponse<Region>>('/regions'),
  stories: (token: string, familyId: number) =>
    request<ListResponse<Story>>(`/stories?family_id=${familyId}`, { headers: authHeader(token) }),
};

function authHeader(token: string) {
  return { Authorization: `Bearer ${token}` };
}
