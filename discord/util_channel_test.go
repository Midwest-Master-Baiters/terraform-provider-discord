package discord

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestGetTextChannelType(t *testing.T) {
	params := []struct {
		id     uint
		chType string
		isHit  bool
	}{
		// success values
		{id: 0, chType: "text", isHit: true},
		{id: 2, chType: "voice", isHit: true},
		{id: 4, chType: "category", isHit: true},
		{id: 5, chType: "news", isHit: true},
		{id: 6, chType: "store", isHit: true},
		{id: 13, chType: "stage", isHit: true},
		{id: 15, chType: "forum", isHit: true},
		// failure values
		{id: 10, chType: "text", isHit: false},
		{id: 100, chType: "text", isHit: false},
	}

	//	case discordgo.ChannelTypeGuildStageVoice:
	//	return "stage", true
	//	case discordgo.ChannelTypeGuildForum:
	//	return "forum", true

	for _, p := range params {
		chType := discordgo.ChannelType(p.id)
		resChType, resIsHit := getTextChannelType(chType)
		if p.chType != resChType {
			t.Errorf("id: %v - chType Error: ex: %v, ac: %v", p.id, p.chType, resChType)
		}
		if p.isHit != resIsHit {
			t.Errorf("id: %v - isHit Error: ex: %v, ac: %v", p.id, p.isHit, resIsHit)
		}
	}
}

func TestGetDiscordChannelType(t *testing.T) {
	params := []struct {
		name   string
		chType uint
		isHit  bool
	}{
		// success values
		{chType: 0, name: "text", isHit: true},
		{chType: 2, name: "voice", isHit: true},
		{chType: 4, name: "category", isHit: true},
		{chType: 5, name: "news", isHit: true},
		{chType: 6, name: "store", isHit: true},
		{chType: 13, name: "stage", isHit: true},
		{chType: 15, name: "forum", isHit: true},
		// failure values
		{chType: 0, name: "lorem", isHit: false},
		{chType: 0, name: "pesudo", isHit: false},
	}

	for _, p := range params {
		resChType, resIsHit := getDiscordChannelType(p.name)
		if p.chType != uint(resChType) {
			t.Errorf("id: %v - chType Error: ex: %v, ac: %v", p.name, p.chType, resChType)
		}
		if p.isHit != resIsHit {
			t.Errorf("id: %v - isHit Error: ex: %v, ac: %v", p.name, p.isHit, resIsHit)
		}
	}
}
