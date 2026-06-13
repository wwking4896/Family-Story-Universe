package domain

import "time"

type User struct {
	ID           int64     `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	DisplayName  string    `json:"display_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Family struct {
	ID          int64     `json:"id"`
	OwnerUserID int64     `json:"owner_user_id"`
	Name        string    `json:"name"`
	PlanType    string    `json:"plan_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type FamilyMember struct {
	ID       int64     `json:"id"`
	FamilyID int64     `json:"family_id"`
	UserID   int64     `json:"user_id"`
	Role     string    `json:"role"`
	Status   string    `json:"status"`
	JoinedAt time.Time `json:"joined_at"`
}

type Child struct {
	ID             int64      `json:"id"`
	FamilyID       int64      `json:"family_id"`
	Name           string     `json:"name"`
	Nickname       string     `json:"nickname"`
	BirthDate      string     `json:"birth_date"`
	GenderOptional *string    `json:"gender_optional"`
	AvatarURL      *string    `json:"avatar_url"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"-"`
}

type Character struct {
	ID                int64      `json:"id"`
	FamilyID          int64      `json:"family_id"`
	ChildID           int64      `json:"child_id"`
	RealName          string     `json:"real_name"`
	StoryName         string     `json:"story_name"`
	RoleType          string     `json:"role_type"`
	PersonalityTraits []string   `json:"personality_traits"`
	Likes             []string   `json:"likes"`
	Fears             []string   `json:"fears"`
	MagicPower        string     `json:"magic_power"`
	AvatarURL         *string    `json:"avatar_url"`
	Level             int        `json:"level"`
	Exp               int        `json:"exp"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"-"`
}

type Region struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Theme       string `json:"theme"`
	UnlockLevel int    `json:"unlock_level"`
	SortOrder   int    `json:"sort_order"`
	IsActive    bool   `json:"is_active"`
}

type Story struct {
	ID              int64             `json:"id"`
	FamilyID        int64             `json:"family_id"`
	ChildID         int64             `json:"child_id"`
	MainCharacterID int64             `json:"main_character_id"`
	RegionID        int64             `json:"region_id"`
	Title           string            `json:"title"`
	Summary         string            `json:"summary"`
	Content         string            `json:"content"`
	Theme           string            `json:"theme"`
	StoryLength     string            `json:"story_length"`
	RealLifeEvent   string            `json:"real_life_event"`
	Tone            string            `json:"tone"`
	Language        string            `json:"language"`
	AIProvider      string            `json:"ai_provider"`
	AIModel         string            `json:"ai_model"`
	Status          string            `json:"status"`
	SafetyCheck     map[string]bool   `json:"safety_check"`
	MemoryTags      []string          `json:"memory_tags"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	DeletedAt       *time.Time        `json:"-"`
	Metadata        map[string]string `json:"metadata,omitempty"`
}

type StoryMemory struct {
	ID              int64     `json:"id"`
	FamilyID        int64     `json:"family_id"`
	ChildID         int64     `json:"child_id"`
	StoryID         int64     `json:"story_id"`
	Tag             string    `json:"tag"`
	MemoryType      string    `json:"memory_type"`
	ImportanceScore int       `json:"importance_score"`
	CreatedAt       time.Time `json:"created_at"`
}

type StoryGenerationJob struct {
	ID                int64      `json:"id"`
	FamilyID          int64      `json:"family_id"`
	RequestedByUserID int64      `json:"requested_by_user_id"`
	StoryID           *int64     `json:"story_id"`
	Status            string     `json:"status"`
	ErrorMessage      string     `json:"error_message,omitempty"`
	StartedAt         time.Time  `json:"started_at"`
	FinishedAt        *time.Time `json:"finished_at"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}
