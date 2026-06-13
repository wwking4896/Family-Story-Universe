package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/fairy-castle/family-story-universe/backend/internal/domain"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrNotFound     = errors.New("not found")
	ErrValidation   = errors.New("validation error")
	ErrConflict     = errors.New("conflict")
)

type MemoryStore struct {
	mu sync.RWMutex

	jwtSecret string

	nextUserID      int64
	nextFamilyID    int64
	nextMemberID    int64
	nextChildID     int64
	nextCharacterID int64
	nextStoryID     int64
	nextMemoryID    int64
	nextJobID       int64

	usersByID    map[int64]domain.User
	usersByEmail map[string]int64
	families     map[int64]domain.Family
	members      map[int64]domain.FamilyMember
	children     map[int64]domain.Child
	characters   map[int64]domain.Character
	regions      map[int64]domain.Region
	stories      map[int64]domain.Story
	memories     map[int64]domain.StoryMemory
	jobs         map[int64]domain.StoryGenerationJob
}

type AuthResult struct {
	User        domain.User `json:"user"`
	AccessToken string      `json:"access_token"`
	ExpiresIn   int64       `json:"expires_in"`
}

type TimebookYear struct {
	Year   int             `json:"year"`
	Months []TimebookMonth `json:"months"`
}

type TimebookMonth struct {
	Month   int            `json:"month"`
	Stories []domain.Story `json:"stories"`
}

type TimebookResponse struct {
	Years []TimebookYear `json:"years"`
}

type StoryGenerateInput struct {
	FamilyID              int64  `json:"family_id"`
	ChildID               int64  `json:"child_id"`
	MainCharacterID       int64  `json:"main_character_id"`
	RegionID              int64  `json:"region_id"`
	Theme                 string `json:"theme"`
	StoryLength           string `json:"story_length"`
	RealLifeEventOptional string `json:"real_life_event_optional"`
	Tone                  string `json:"tone"`
	Language              string `json:"language"`
}

type StoryGenerateResult struct {
	JobID  int64        `json:"job_id"`
	Status string       `json:"status"`
	Story  domain.Story `json:"story"`
}

func NewMemoryStore(jwtSecret string) *MemoryStore {
	store := &MemoryStore{
		jwtSecret:       jwtSecret,
		nextUserID:      1,
		nextFamilyID:    1,
		nextMemberID:    1,
		nextChildID:     1,
		nextCharacterID: 1,
		nextStoryID:     1,
		nextMemoryID:    1,
		nextJobID:       1,
		usersByID:       map[int64]domain.User{},
		usersByEmail:    map[string]int64{},
		families:        map[int64]domain.Family{},
		members:         map[int64]domain.FamilyMember{},
		children:        map[int64]domain.Child{},
		characters:      map[int64]domain.Character{},
		regions:         map[int64]domain.Region{},
		stories:         map[int64]domain.Story{},
		memories:        map[int64]domain.StoryMemory{},
		jobs:            map[int64]domain.StoryGenerationJob{},
	}
	store.seedRegions()
	return store
}

func (s *MemoryStore) Register(email, password, displayName string) (AuthResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" || password == "" || displayName == "" {
		return AuthResult{}, ErrValidation
	}
	if _, exists := s.usersByEmail[email]; exists {
		return AuthResult{}, ErrConflict
	}
	now := time.Now().UTC()
	user := domain.User{ID: s.nextUserID, Email: email, PasswordHash: hashPassword(password), DisplayName: displayName, CreatedAt: now, UpdatedAt: now}
	s.nextUserID++
	s.usersByID[user.ID] = user
	s.usersByEmail[email] = user.ID
	return s.authResult(user)
}

func (s *MemoryStore) Login(email, password string) (AuthResult, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	id, exists := s.usersByEmail[strings.ToLower(strings.TrimSpace(email))]
	if !exists {
		return AuthResult{}, ErrUnauthorized
	}
	user := s.usersByID[id]
	if user.PasswordHash != hashPassword(password) {
		return AuthResult{}, ErrUnauthorized
	}
	return s.authResult(user)
}

func (s *MemoryStore) UserFromToken(token string) (domain.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	claims, err := s.parseToken(token)
	if err != nil {
		return domain.User{}, ErrUnauthorized
	}
	user, exists := s.usersByID[claims.UserID]
	if !exists {
		return domain.User{}, ErrUnauthorized
	}
	return user, nil
}

func (s *MemoryStore) CreateFamily(userID int64, name string) (domain.Family, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if strings.TrimSpace(name) == "" {
		return domain.Family{}, ErrValidation
	}
	now := time.Now().UTC()
	family := domain.Family{ID: s.nextFamilyID, OwnerUserID: userID, Name: name, PlanType: "free", CreatedAt: now, UpdatedAt: now}
	s.nextFamilyID++
	s.families[family.ID] = family
	member := domain.FamilyMember{ID: s.nextMemberID, FamilyID: family.ID, UserID: userID, Role: "owner", Status: "active", JoinedAt: now}
	s.nextMemberID++
	s.members[member.ID] = member
	return family, nil
}

func (s *MemoryStore) MyFamilies(userID int64) []domain.Family {
	s.mu.RLock()
	defer s.mu.RUnlock()
	items := []domain.Family{}
	for _, member := range s.members {
		if member.UserID == userID && member.Status == "active" {
			items = append(items, s.families[member.FamilyID])
		}
	}
	return items
}

func (s *MemoryStore) CreateChild(userID int64, child domain.Child) (domain.Child, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.isMemberLocked(userID, child.FamilyID) || strings.TrimSpace(child.Name) == "" {
		return domain.Child{}, ErrForbidden
	}
	now := time.Now().UTC()
	child.ID = s.nextChildID
	child.CreatedAt = now
	child.UpdatedAt = now
	s.nextChildID++
	s.children[child.ID] = child
	return child, nil
}

func (s *MemoryStore) ListChildren(userID, familyID int64) ([]domain.Child, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if !s.isMemberLocked(userID, familyID) {
		return nil, ErrForbidden
	}
	items := []domain.Child{}
	for _, child := range s.children {
		if child.FamilyID == familyID && child.DeletedAt == nil {
			items = append(items, child)
		}
	}
	return items, nil
}

func (s *MemoryStore) CreateCharacter(userID int64, character domain.Character) (domain.Character, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	child, ok := s.children[character.ChildID]
	if !ok || child.FamilyID != character.FamilyID || !s.isMemberLocked(userID, character.FamilyID) || strings.TrimSpace(character.StoryName) == "" {
		return domain.Character{}, ErrForbidden
	}
	now := time.Now().UTC()
	character.ID = s.nextCharacterID
	character.Level = 1
	character.Exp = 0
	character.CreatedAt = now
	character.UpdatedAt = now
	s.nextCharacterID++
	s.characters[character.ID] = character
	return character, nil
}

func (s *MemoryStore) ListCharacters(userID, familyID int64) ([]domain.Character, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if !s.isMemberLocked(userID, familyID) {
		return nil, ErrForbidden
	}
	items := []domain.Character{}
	for _, character := range s.characters {
		if character.FamilyID == familyID && character.DeletedAt == nil {
			items = append(items, character)
		}
	}
	return items, nil
}

func (s *MemoryStore) Regions() []domain.Region {
	s.mu.RLock()
	defer s.mu.RUnlock()
	items := make([]domain.Region, 0, len(s.regions))
	for _, region := range s.regions {
		items = append(items, region)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].SortOrder < items[j].SortOrder })
	return items
}

func (s *MemoryStore) GenerateStory(userID int64, input StoryGenerateInput) (StoryGenerateResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.isMemberLocked(userID, input.FamilyID) {
		return StoryGenerateResult{}, ErrForbidden
	}
	child, childOK := s.children[input.ChildID]
	character, characterOK := s.characters[input.MainCharacterID]
	region, regionOK := s.regions[input.RegionID]
	if !childOK || !characterOK || !regionOK || child.FamilyID != input.FamilyID || character.FamilyID != input.FamilyID {
		return StoryGenerateResult{}, ErrForbidden
	}
	now := time.Now().UTC()
	job := domain.StoryGenerationJob{ID: s.nextJobID, FamilyID: input.FamilyID, RequestedByUserID: userID, Status: "processing", StartedAt: now, CreatedAt: now, UpdatedAt: now}
	s.nextJobID++
	s.jobs[job.ID] = job

	memories := s.recentMemoryTagsLocked(input.FamilyID, input.ChildID, 8)
	content, tags := generateMockStory(child, character, region, input, memories)
	story := domain.Story{ID: s.nextStoryID, FamilyID: input.FamilyID, ChildID: input.ChildID, MainCharacterID: input.MainCharacterID, RegionID: input.RegionID, Title: fmt.Sprintf("%s的%s冒險", character.StoryName, input.Theme), Summary: fmt.Sprintf("%s在%s學會了%s。", character.StoryName, region.Name, input.Theme), Content: content, Theme: input.Theme, StoryLength: input.StoryLength, RealLifeEvent: input.RealLifeEventOptional, Tone: input.Tone, Language: defaultString(input.Language, "zh-TW"), AIProvider: "mock", AIModel: "fairy-castle-mock-v1", Status: "published", SafetyCheck: safeCheck(), MemoryTags: tags, CreatedAt: now, UpdatedAt: now}
	s.nextStoryID++
	s.stories[story.ID] = story
	for _, tag := range tags {
		memory := domain.StoryMemory{ID: s.nextMemoryID, FamilyID: story.FamilyID, ChildID: story.ChildID, StoryID: story.ID, Tag: tag, MemoryType: "story", ImportanceScore: 5, CreatedAt: now}
		s.nextMemoryID++
		s.memories[memory.ID] = memory
	}
	finished := time.Now().UTC()
	job.Status = "completed"
	job.StoryID = &story.ID
	job.FinishedAt = &finished
	job.UpdatedAt = finished
	s.jobs[job.ID] = job
	return StoryGenerateResult{JobID: job.ID, Status: job.Status, Story: story}, nil
}

func (s *MemoryStore) ListStories(userID, familyID int64) ([]domain.Story, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if !s.isMemberLocked(userID, familyID) {
		return nil, ErrForbidden
	}
	items := []domain.Story{}
	for _, story := range s.stories {
		if story.FamilyID == familyID && story.DeletedAt == nil {
			items = append(items, story)
		}
	}
	sort.Slice(items, func(i, j int) bool { return items[i].CreatedAt.After(items[j].CreatedAt) })
	return items, nil
}

func (s *MemoryStore) GetStory(userID, storyID int64) (domain.Story, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	story, ok := s.stories[storyID]
	if !ok || story.DeletedAt != nil {
		return domain.Story{}, ErrNotFound
	}
	if !s.isMemberLocked(userID, story.FamilyID) {
		return domain.Story{}, ErrForbidden
	}
	return story, nil
}

func (s *MemoryStore) Timebook(userID, familyID int64) (TimebookResponse, error) {
	stories, err := s.ListStories(userID, familyID)
	if err != nil {
		return TimebookResponse{}, err
	}
	byYearMonth := map[int]map[int][]domain.Story{}
	for _, story := range stories {
		year, month, _ := story.CreatedAt.Date()
		if _, ok := byYearMonth[year]; !ok {
			byYearMonth[year] = map[int][]domain.Story{}
		}
		byYearMonth[year][int(month)] = append(byYearMonth[year][int(month)], story)
	}
	years := make([]int, 0, len(byYearMonth))
	for year := range byYearMonth {
		years = append(years, year)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(years)))
	response := TimebookResponse{Years: []TimebookYear{}}
	for _, year := range years {
		months := make([]int, 0, len(byYearMonth[year]))
		for month := range byYearMonth[year] {
			months = append(months, month)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(months)))
		y := TimebookYear{Year: year, Months: []TimebookMonth{}}
		for _, month := range months {
			y.Months = append(y.Months, TimebookMonth{Month: month, Stories: byYearMonth[year][month]})
		}
		response.Years = append(response.Years, y)
	}
	return response, nil
}

func (s *MemoryStore) isMemberLocked(userID, familyID int64) bool {
	for _, member := range s.members {
		if member.UserID == userID && member.FamilyID == familyID && member.Status == "active" {
			return true
		}
	}
	return false
}

func (s *MemoryStore) seedRegions() {
	regions := []domain.Region{
		{ID: 1, Name: "童話城堡", Description: "每個家庭故事開始的地方。", Theme: "home", UnlockLevel: 1, SortOrder: 1, IsActive: true},
		{ID: 2, Name: "魔法森林", Description: "充滿溫柔魔法與小任務的森林。", Theme: "forest", UnlockLevel: 1, SortOrder: 2, IsActive: true},
		{ID: 3, Name: "糖果村", Description: "學習分享與禮貌的甜甜村莊。", Theme: "sharing", UnlockLevel: 1, SortOrder: 3, IsActive: true},
		{ID: 4, Name: "星光湖", Description: "適合睡前安撫與親子陪伴的湖畔。", Theme: "bedtime", UnlockLevel: 1, SortOrder: 4, IsActive: true},
		{ID: 5, Name: "彩虹山谷", Description: "練習情緒管理與同理心的山谷。", Theme: "emotion", UnlockLevel: 1, SortOrder: 5, IsActive: true},
		{ID: 6, Name: "龍之谷", Description: "學習勇氣但不過度刺激的冒險區域。", Theme: "courage", UnlockLevel: 1, SortOrder: 6, IsActive: true},
		{ID: 7, Name: "夢境花園", Description: "保存想像力與美好夢境的花園。", Theme: "dream", UnlockLevel: 1, SortOrder: 7, IsActive: true},
		{ID: 8, Name: "時光塔", Description: "回顧家庭故事與成長記憶的高塔。", Theme: "memory", UnlockLevel: 1, SortOrder: 8, IsActive: true},
	}
	for _, region := range regions {
		s.regions[region.ID] = region
	}
}

type tokenClaims struct {
	UserID int64 `json:"user_id"`
	Exp    int64 `json:"exp"`
}

func (s *MemoryStore) authResult(user domain.User) (AuthResult, error) {
	token, err := s.signToken(user.ID)
	if err != nil {
		return AuthResult{}, err
	}
	return AuthResult{User: user, AccessToken: token, ExpiresIn: 3600}, nil
}

func (s *MemoryStore) signToken(userID int64) (string, error) {
	claims := tokenClaims{UserID: userID, Exp: time.Now().Add(time.Hour).Unix()}
	payload, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	encoded := base64.RawURLEncoding.EncodeToString(payload)
	sig := s.signature(encoded)
	return encoded + "." + sig, nil
}

func (s *MemoryStore) parseToken(token string) (tokenClaims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 2 || !hmac.Equal([]byte(parts[1]), []byte(s.signature(parts[0]))) {
		return tokenClaims{}, ErrUnauthorized
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return tokenClaims{}, ErrUnauthorized
	}
	var claims tokenClaims
	if err := json.Unmarshal(payload, &claims); err != nil || time.Now().Unix() > claims.Exp {
		return tokenClaims{}, ErrUnauthorized
	}
	return claims, nil
}

func (s *MemoryStore) signature(payload string) string {
	mac := hmac.New(sha256.New, []byte(s.jwtSecret))
	mac.Write([]byte(payload))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

func hashPassword(password string) string {
	sum := sha256.Sum256([]byte("fairy-castle:" + password))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

func (s *MemoryStore) recentMemoryTagsLocked(familyID, childID int64, limit int) []string {
	items := []domain.StoryMemory{}
	for _, memory := range s.memories {
		if memory.FamilyID == familyID && memory.ChildID == childID {
			items = append(items, memory)
		}
	}
	sort.Slice(items, func(i, j int) bool { return items[i].CreatedAt.After(items[j].CreatedAt) })
	tags := []string{}
	for i, memory := range items {
		if i >= limit {
			break
		}
		tags = append(tags, memory.Tag)
	}
	return tags
}

func generateMockStory(child domain.Child, character domain.Character, region domain.Region, input StoryGenerateInput, memories []string) (string, []string) {
	theme := defaultString(input.Theme, "勇氣")
	tone := defaultString(input.Tone, "溫柔")
	event := strings.TrimSpace(input.RealLifeEventOptional)
	if event == "" {
		event = "今天沒有特別的大事，但心裡有一顆小小的星星想發光。"
	}
	memoryText := ""
	if len(memories) > 0 {
		memoryText = "她想起以前的故事記憶：「" + strings.Join(memories, "、") + "」，心裡變得更安定。"
	}
	content := fmt.Sprintf("在%s裡，%s帶著%s的心情出發了。她記得%s最喜歡的事物，也記得自己有一個特別的魔法：%s。\n\n今天的任務來自一件真實的小事情：%s 於是，%s把它變成一個溫柔的童話任務，一步一步完成。%s\n\n最後，%s發現，%s不是一下子變得很厲害，而是在有人陪伴、願意嘗試的時候，心裡慢慢長出亮晶晶的力量。夜晚安靜下來，%s也帶著微笑，準備做一個甜甜的夢。", region.Name, character.StoryName, tone, child.Name, character.MagicPower, event, character.StoryName, memoryText, character.StoryName, theme, child.Nickname)
	tags := []string{theme, region.Name, character.StoryName}
	if input.RealLifeEventOptional != "" {
		tags = append(tags, "真實事件")
	}
	return content, tags
}

func safeCheck() map[string]bool {
	return map[string]bool{"violence": false, "death": false, "adult_content": false, "scary_content": false, "discrimination": false, "unsafe_behavior": false}
}

func defaultString(value, fallback string) string {
	if strings.TrimSpace(value) == "" {
		return fallback
	}
	return value
}
