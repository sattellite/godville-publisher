package godville

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Godville struct {
	godname string
	token   string
}

type apiInfo struct {
	Name              string    `json:"name"`
	Godname           string    `json:"godname"`
	Gender            string    `json:"gender"`
	Level             int       `json:"level"`
	MaxHealth         int       `json:"max_health"`
	InventoryMaxNum   int       `json:"inventory_max_num"`
	Motto             string    `json:"motto"`
	Clan              string    `json:"clan"`
	ClanPosition      string    `json:"clan_position"`
	Alignment         string    `json:"alignment"`
	BricksCnt         int       `json:"bricks_cnt"`
	WoodCnt           int       `json:"wood_cnt"`
	TempleCompletedAt time.Time `json:"temple_completed_at"`
	Pet               struct {
		PetName  string `json:"pet_name"`
		PetClass string `json:"pet_class"`
		PetLevel int    `json:"pet_level"`
	} `json:"pet"`
	ArkCompletedAt time.Time `json:"ark_completed_at"`
	ArkF           int       `json:"ark_f"`
	ArkM           int       `json:"ark_m"`
	ArenaWon       int       `json:"arena_won"`
	ArenaLost      int       `json:"arena_lost"`
	Savings        string    `json:"savings"`

	// private fields
	Health        int           `json:"health"`
	QuestProgress int           `json:"quest_progress"`
	ExpProgress   int           `json:"exp_progress"`
	Godpower      int           `json:"godpower"`
	GoldApprox    string        `json:"gold_approx"`
	DiaryLast     string        `json:"diary_last"`
	TownName      string        `json:"town_name"`
	Distance      int           `json:"distance"`
	ArenaFight    bool          `json:"arena_fight"`
	InventoryNum  int           `json:"inventory_num"`
	Quest         string        `json:"quest"`
	Aura          string        `json:"aura"`
	Activatables  []interface{} `json:"activatables"`
}

type Info struct {
	Name      string
	Godname   string
	Alignment string
	Clan      string
	DiaryLast string
	TownName  string

	Level     int
	Distance  int
	Health    int
	MaxHealth int

	GoldApprox string
	Quest      string
}

func New(godname, token string) *Godville {
	return &Godville{
		godname: godname,
		token:   token,
	}
}

func (g *Godville) Info(ctx context.Context) (*Info, error) {
	req, reqErr := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("https://godville.net/gods/api/%s/%s", g.godname, g.token),
		http.NoBody,
	)
	if reqErr != nil {
		return nil, fmt.Errorf("failed to create request: %w", reqErr)
	}

	resp, getErr := http.DefaultClient.Do(req)
	if getErr != nil {
		return nil, getErr
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	if resp == nil || resp.Body == nil {
		return nil, fmt.Errorf("unexpected response: %v", resp)
	}
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			fmt.Printf("failed to close response body: %v\n", closeErr)
		}
	}()

	// print response body
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return nil, fmt.Errorf("failed to read response body: %w", readErr)
	}

	var info apiInfo
	if err := json.Unmarshal(body, &info); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &Info{
		Name:      info.Name,
		Godname:   info.Godname,
		Alignment: info.Alignment,
		Clan:      info.Clan,
		DiaryLast: info.DiaryLast,
		TownName:  info.TownName,

		Level:     info.Level,
		Distance:  info.Distance,
		Health:    info.Health,
		MaxHealth: info.MaxHealth,

		GoldApprox: info.GoldApprox,
		Quest:      info.Quest,
	}, nil
}
