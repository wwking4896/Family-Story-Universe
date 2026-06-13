export type AuthResult = {
  user: { id: number; email: string; display_name: string };
  access_token: string;
  expires_in: number;
};

export type Family = {
  id: number;
  owner_user_id: number;
  name: string;
  plan_type: string;
};

export type Child = {
  id: number;
  family_id: number;
  name: string;
  nickname: string;
  birth_date: string;
  gender_optional?: string | null;
  avatar_url?: string | null;
};

export type Character = {
  id: number;
  family_id: number;
  child_id: number;
  real_name: string;
  story_name: string;
  role_type: string;
  personality_traits: string[];
  likes: string[];
  fears: string[];
  magic_power: string;
  level: number;
  exp: number;
};

export type Region = {
  id: number;
  name: string;
  description: string;
  theme: string;
  unlock_level: number;
  sort_order: number;
  is_active: boolean;
};

export type Story = {
  id: number;
  family_id: number;
  child_id: number;
  main_character_id: number;
  region_id: number;
  title: string;
  summary: string;
  content: string;
  theme: string;
  story_length: string;
  tone: string;
  language: string;
  status: string;
  memory_tags: string[];
  created_at: string;
};
