package services

import (
	"testing"

	"github.com/fairy-castle/family-story-universe/backend/internal/domain"
)

func TestMVPFlow(t *testing.T) {
	store := NewMemoryStore("test-secret")
	auth, err := store.Register("parent@example.com", "password123", "小雨爸爸")
	if err != nil {
		t.Fatalf("register: %v", err)
	}
	user, err := store.UserFromToken(auth.AccessToken)
	if err != nil {
		t.Fatalf("token: %v", err)
	}
	family, err := store.CreateFamily(user.ID, "小雨的童話城堡")
	if err != nil {
		t.Fatalf("family: %v", err)
	}
	child, err := store.CreateChild(user.ID, domain.Child{FamilyID: family.ID, Name: "小雨", Nickname: "雨雨", BirthDate: "2022-05-01"})
	if err != nil {
		t.Fatalf("child: %v", err)
	}
	character, err := store.CreateCharacter(user.ID, domain.Character{FamilyID: family.ID, ChildID: child.ID, RealName: "小雨", StoryName: "星光小魔女", RoleType: "月光魔法學徒", PersonalityTraits: []string{"好奇", "善良"}, Likes: []string{"兔子"}, Fears: []string{"打雷"}, MagicPower: "讓星星發出溫柔的光"})
	if err != nil {
		t.Fatalf("character: %v", err)
	}
	result, err := store.GenerateStory(user.ID, StoryGenerateInput{FamilyID: family.ID, ChildID: child.ID, MainCharacterID: character.ID, RegionID: 2, Theme: "勇氣", StoryLength: "5_min", Tone: "睡前安撫", Language: "zh-TW"})
	if err != nil {
		t.Fatalf("story: %v", err)
	}
	if result.Story.Status != "published" || len(result.Story.MemoryTags) == 0 {
		t.Fatalf("unexpected story result: %+v", result.Story)
	}
	timebook, err := store.Timebook(user.ID, family.ID)
	if err != nil {
		t.Fatalf("timebook: %v", err)
	}
	if len(timebook.Years) != 1 || len(timebook.Years[0].Months) != 1 {
		t.Fatalf("unexpected timebook: %+v", timebook)
	}
}

func TestFamilyIsolation(t *testing.T) {
	store := NewMemoryStore("test-secret")
	a, err := store.Register("a@example.com", "password123", "A")
	if err != nil {
		t.Fatal(err)
	}
	b, err := store.Register("b@example.com", "password123", "B")
	if err != nil {
		t.Fatal(err)
	}
	family, err := store.CreateFamily(a.User.ID, "A family")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := store.ListChildren(b.User.ID, family.ID); err != ErrForbidden {
		t.Fatalf("expected forbidden, got %v", err)
	}
}
