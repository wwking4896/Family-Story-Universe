CREATE TABLE IF NOT EXISTS users (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  email VARCHAR(255) NOT NULL UNIQUE,
  password_hash VARCHAR(255) NOT NULL,
  display_name VARCHAR(100) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME NULL
);

CREATE TABLE IF NOT EXISTS families (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  owner_user_id BIGINT NOT NULL,
  name VARCHAR(100) NOT NULL,
  plan_type VARCHAR(50) NOT NULL DEFAULT 'free',
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME NULL,
  INDEX idx_families_owner_user_id (owner_user_id)
);

CREATE TABLE IF NOT EXISTS family_members (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  family_id BIGINT NOT NULL,
  user_id BIGINT NOT NULL,
  role VARCHAR(50) NOT NULL,
  status VARCHAR(50) NOT NULL,
  joined_at DATETIME NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY uk_family_members_family_user (family_id, user_id),
  INDEX idx_family_members_user_id (user_id)
);

CREATE TABLE IF NOT EXISTS children (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  family_id BIGINT NOT NULL,
  name VARCHAR(100) NOT NULL,
  nickname VARCHAR(100) NOT NULL,
  birth_date DATE NULL,
  gender_optional VARCHAR(50) NULL,
  avatar_url VARCHAR(500) NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME NULL,
  INDEX idx_children_family_deleted (family_id, deleted_at)
);

CREATE TABLE IF NOT EXISTS characters (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  family_id BIGINT NOT NULL,
  child_id BIGINT NOT NULL,
  real_name VARCHAR(100) NOT NULL,
  story_name VARCHAR(100) NOT NULL,
  role_type VARCHAR(100) NOT NULL,
  personality_traits JSON NULL,
  likes JSON NULL,
  fears JSON NULL,
  magic_power VARCHAR(255) NOT NULL,
  avatar_url VARCHAR(500) NULL,
  level INT NOT NULL DEFAULT 1,
  exp INT NOT NULL DEFAULT 0,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME NULL,
  INDEX idx_characters_family_child_deleted (family_id, child_id, deleted_at)
);

CREATE TABLE IF NOT EXISTS regions (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL,
  description TEXT NOT NULL,
  theme VARCHAR(100) NOT NULL,
  unlock_level INT NOT NULL DEFAULT 1,
  sort_order INT NOT NULL,
  is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS stories (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  family_id BIGINT NOT NULL,
  child_id BIGINT NOT NULL,
  main_character_id BIGINT NOT NULL,
  region_id BIGINT NOT NULL,
  title VARCHAR(255) NOT NULL,
  summary TEXT NOT NULL,
  content LONGTEXT NOT NULL,
  theme VARCHAR(100) NOT NULL,
  story_length VARCHAR(50) NOT NULL,
  real_life_event TEXT NULL,
  tone VARCHAR(50) NOT NULL,
  language VARCHAR(20) NOT NULL DEFAULT 'zh-TW',
  ai_provider VARCHAR(100) NOT NULL,
  ai_model VARCHAR(100) NOT NULL,
  status VARCHAR(50) NOT NULL,
  safety_check JSON NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME NULL,
  INDEX idx_stories_family_created (family_id, created_at),
  INDEX idx_stories_child_created (child_id, created_at),
  INDEX idx_stories_theme (theme)
);

CREATE TABLE IF NOT EXISTS story_memories (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  family_id BIGINT NOT NULL,
  child_id BIGINT NOT NULL,
  story_id BIGINT NOT NULL,
  tag VARCHAR(100) NOT NULL,
  memory_type VARCHAR(50) NOT NULL,
  importance_score INT NOT NULL DEFAULT 5,
  created_at DATETIME NOT NULL,
  INDEX idx_story_memories_family_child_created (family_id, child_id, created_at),
  INDEX idx_story_memories_tag (tag)
);

CREATE TABLE IF NOT EXISTS story_generation_jobs (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  family_id BIGINT NOT NULL,
  requested_by_user_id BIGINT NOT NULL,
  story_id BIGINT NULL,
  status VARCHAR(50) NOT NULL,
  request_payload JSON NULL,
  ai_response_payload JSON NULL,
  error_message TEXT NULL,
  started_at DATETIME NOT NULL,
  finished_at DATETIME NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  INDEX idx_story_generation_jobs_family_status_created (family_id, status, created_at)
);

INSERT IGNORE INTO schema_migrations (version) VALUES ('000002_core_schema');
