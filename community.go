package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Skill struct {
	Technology  string
	Level       int
	LastUpdated time.Time
}

type Member struct {
	DiscordID string
	Username  string

	Skills        []Skill
	LearningGoals []string

	CanMentor   bool
	WantsMentor bool

	JoinedAt   time.Time
	LastActive time.Time
}

type Community struct {
	mu      sync.RWMutex
	members map[string]*Member
}

func NewCommunity() *Community {
	return &Community{
		members: make(map[string]*Member),
	}
}

func (c *Community) AddOrUpdateMember(discordID, username string) *Member {
	c.mu.Lock()
	defer c.mu.Unlock()

	member, exists := c.members[discordID]

	if !exists {
		member = &Member{
			DiscordID:     discordID,
			Username:      username,
			Skills:        make([]Skill, 0, 3),
			LearningGoals: make([]string, 0, 2),
			JoinedAt:      time.Now(),
		}
		c.members[discordID] = member
	}
	member.LastActive = time.Now()
	return member
}

func (m *Member) AddSkill(technology string, level int) error {
	//no sabe nada, beginner, intermedio, avanzado
	if level < 1 || level > 4 {
		return fmt.Errorf("la habilidad debe estar entre 1 a 4, chau")
	}
	for i := range m.Skills {
		if strings.EqualFold(m.Skills[i].Technology, technology) {
			m.Skills[i].Level = level
			m.Skills[i].LastUpdated = time.Now()
		}
	}

	m.Skills = append(m.Skills, Skill{
		Technology:  technology,
		Level:       level,
		LastUpdated: time.Now(),
	})
	return nil
}

func (m *Member) AddLearningGoal(technology string) {
	for _, goal := range m.LearningGoals {
		if strings.EqualFold(goal, technology) {
			return
		}
		m.LearningGoals = append(m.LearningGoals, technology)
	}
}

func (m Member) GetSkillLevel(technology string) int {
	for _, skill := range m.Skills {
		if strings.EqualFold(skill.Technology, technology) {
			return skill.Level
		}
	}
	return 0
}

func (m Member) CanMentorF(technology string) bool {
	return m.CanMentor && m.GetSkillLevel(technology) >= 3
}

type MentorshipMatch struct {
	Mentor     *Member
	Mentee     *Member
	Technology string
}

//TODO: findmatches :)
