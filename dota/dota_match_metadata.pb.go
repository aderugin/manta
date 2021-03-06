// Code generated by protoc-gen-go.
// source: dota_match_metadata.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CDOTAMatchMetadataFile struct {
	Version          *int32              `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	MatchId          *uint64             `protobuf:"varint,2,req,name=match_id,json=matchId" json:"match_id,omitempty"`
	Metadata         *CDOTAMatchMetadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
	PrivateMetadata  []byte              `protobuf:"bytes,5,opt,name=private_metadata,json=privateMetadata" json:"private_metadata,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *CDOTAMatchMetadataFile) Reset()                    { *m = CDOTAMatchMetadataFile{} }
func (m *CDOTAMatchMetadataFile) String() string            { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadataFile) ProtoMessage()               {}
func (*CDOTAMatchMetadataFile) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{0} }

func (m *CDOTAMatchMetadataFile) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *CDOTAMatchMetadataFile) GetMatchId() uint64 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *CDOTAMatchMetadataFile) GetMetadata() *CDOTAMatchMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CDOTAMatchMetadataFile) GetPrivateMetadata() []byte {
	if m != nil {
		return m.PrivateMetadata
	}
	return nil
}

type CDOTAMatchMetadata struct {
	Teams                []*CDOTAMatchMetadata_Team  `protobuf:"bytes,1,rep,name=teams" json:"teams,omitempty"`
	ItemRewards          []*CLobbyTimedRewardDetails `protobuf:"bytes,2,rep,name=item_rewards,json=itemRewards" json:"item_rewards,omitempty"`
	LobbyId              *uint64                     `protobuf:"fixed64,3,opt,name=lobby_id,json=lobbyId" json:"lobby_id,omitempty"`
	ReportUntilTime      *uint64                     `protobuf:"fixed64,4,opt,name=report_until_time,json=reportUntilTime" json:"report_until_time,omitempty"`
	EventGameCustomTable []byte                      `protobuf:"bytes,5,opt,name=event_game_custom_table,json=eventGameCustomTable" json:"event_game_custom_table,omitempty"`
	XXX_unrecognized     []byte                      `json:"-"`
}

func (m *CDOTAMatchMetadata) Reset()                    { *m = CDOTAMatchMetadata{} }
func (m *CDOTAMatchMetadata) String() string            { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata) ProtoMessage()               {}
func (*CDOTAMatchMetadata) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{1} }

func (m *CDOTAMatchMetadata) GetTeams() []*CDOTAMatchMetadata_Team {
	if m != nil {
		return m.Teams
	}
	return nil
}

func (m *CDOTAMatchMetadata) GetItemRewards() []*CLobbyTimedRewardDetails {
	if m != nil {
		return m.ItemRewards
	}
	return nil
}

func (m *CDOTAMatchMetadata) GetLobbyId() uint64 {
	if m != nil && m.LobbyId != nil {
		return *m.LobbyId
	}
	return 0
}

func (m *CDOTAMatchMetadata) GetReportUntilTime() uint64 {
	if m != nil && m.ReportUntilTime != nil {
		return *m.ReportUntilTime
	}
	return 0
}

func (m *CDOTAMatchMetadata) GetEventGameCustomTable() []byte {
	if m != nil {
		return m.EventGameCustomTable
	}
	return nil
}

type CDOTAMatchMetadata_Team struct {
	DotaTeam          *uint32                           `protobuf:"varint,1,opt,name=dota_team,json=dotaTeam" json:"dota_team,omitempty"`
	Players           []*CDOTAMatchMetadata_Team_Player `protobuf:"bytes,2,rep,name=players" json:"players,omitempty"`
	GraphExperience   []float32                         `protobuf:"fixed32,3,rep,name=graph_experience,json=graphExperience" json:"graph_experience,omitempty"`
	GraphGoldEarned   []float32                         `protobuf:"fixed32,4,rep,name=graph_gold_earned,json=graphGoldEarned" json:"graph_gold_earned,omitempty"`
	GraphNetWorth     []float32                         `protobuf:"fixed32,5,rep,name=graph_net_worth,json=graphNetWorth" json:"graph_net_worth,omitempty"`
	CmFirstPick       *bool                             `protobuf:"varint,6,opt,name=cm_first_pick,json=cmFirstPick" json:"cm_first_pick,omitempty"`
	CmCaptainPlayerId *uint32                           `protobuf:"varint,7,opt,name=cm_captain_player_id,json=cmCaptainPlayerId" json:"cm_captain_player_id,omitempty"`
	CmBans            []uint32                          `protobuf:"varint,8,rep,name=cm_bans,json=cmBans" json:"cm_bans,omitempty"`
	CmPicks           []uint32                          `protobuf:"varint,9,rep,name=cm_picks,json=cmPicks" json:"cm_picks,omitempty"`
	CmPenalty         *uint32                           `protobuf:"varint,10,opt,name=cm_penalty,json=cmPenalty" json:"cm_penalty,omitempty"`
	XXX_unrecognized  []byte                            `json:"-"`
}

func (m *CDOTAMatchMetadata_Team) Reset()                    { *m = CDOTAMatchMetadata_Team{} }
func (m *CDOTAMatchMetadata_Team) String() string            { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata_Team) ProtoMessage()               {}
func (*CDOTAMatchMetadata_Team) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{1, 0} }

func (m *CDOTAMatchMetadata_Team) GetDotaTeam() uint32 {
	if m != nil && m.DotaTeam != nil {
		return *m.DotaTeam
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team) GetPlayers() []*CDOTAMatchMetadata_Team_Player {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetGraphExperience() []float32 {
	if m != nil {
		return m.GraphExperience
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetGraphGoldEarned() []float32 {
	if m != nil {
		return m.GraphGoldEarned
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetGraphNetWorth() []float32 {
	if m != nil {
		return m.GraphNetWorth
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetCmFirstPick() bool {
	if m != nil && m.CmFirstPick != nil {
		return *m.CmFirstPick
	}
	return false
}

func (m *CDOTAMatchMetadata_Team) GetCmCaptainPlayerId() uint32 {
	if m != nil && m.CmCaptainPlayerId != nil {
		return *m.CmCaptainPlayerId
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team) GetCmBans() []uint32 {
	if m != nil {
		return m.CmBans
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetCmPicks() []uint32 {
	if m != nil {
		return m.CmPicks
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team) GetCmPenalty() uint32 {
	if m != nil && m.CmPenalty != nil {
		return *m.CmPenalty
	}
	return 0
}

type CDOTAMatchMetadata_Team_PlayerKill struct {
	VictimSlot       *uint32 `protobuf:"varint,1,opt,name=victim_slot,json=victimSlot" json:"victim_slot,omitempty"`
	Count            *uint32 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAMatchMetadata_Team_PlayerKill) Reset()         { *m = CDOTAMatchMetadata_Team_PlayerKill{} }
func (m *CDOTAMatchMetadata_Team_PlayerKill) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata_Team_PlayerKill) ProtoMessage()    {}
func (*CDOTAMatchMetadata_Team_PlayerKill) Descriptor() ([]byte, []int) {
	return fileDescriptor22, []int{1, 0, 0}
}

func (m *CDOTAMatchMetadata_Team_PlayerKill) GetVictimSlot() uint32 {
	if m != nil && m.VictimSlot != nil {
		return *m.VictimSlot
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_PlayerKill) GetCount() uint32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

type CDOTAMatchMetadata_Team_ItemPurchase struct {
	ItemId           *uint32 `protobuf:"varint,1,opt,name=item_id,json=itemId" json:"item_id,omitempty"`
	PurchaseTime     *int32  `protobuf:"varint,2,opt,name=purchase_time,json=purchaseTime" json:"purchase_time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAMatchMetadata_Team_ItemPurchase) Reset()         { *m = CDOTAMatchMetadata_Team_ItemPurchase{} }
func (m *CDOTAMatchMetadata_Team_ItemPurchase) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata_Team_ItemPurchase) ProtoMessage()    {}
func (*CDOTAMatchMetadata_Team_ItemPurchase) Descriptor() ([]byte, []int) {
	return fileDescriptor22, []int{1, 0, 1}
}

func (m *CDOTAMatchMetadata_Team_ItemPurchase) GetItemId() uint32 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_ItemPurchase) GetPurchaseTime() int32 {
	if m != nil && m.PurchaseTime != nil {
		return *m.PurchaseTime
	}
	return 0
}

type CDOTAMatchMetadata_Team_InventorySnapshot struct {
	ItemId           []uint32 `protobuf:"varint,1,rep,name=item_id,json=itemId" json:"item_id,omitempty"`
	GameTime         *int32   `protobuf:"varint,2,opt,name=game_time,json=gameTime" json:"game_time,omitempty"`
	Kills            *uint32  `protobuf:"varint,3,opt,name=kills" json:"kills,omitempty"`
	Deaths           *uint32  `protobuf:"varint,4,opt,name=deaths" json:"deaths,omitempty"`
	Assists          *uint32  `protobuf:"varint,5,opt,name=assists" json:"assists,omitempty"`
	Level            *uint32  `protobuf:"varint,6,opt,name=level" json:"level,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CDOTAMatchMetadata_Team_InventorySnapshot) Reset() {
	*m = CDOTAMatchMetadata_Team_InventorySnapshot{}
}
func (m *CDOTAMatchMetadata_Team_InventorySnapshot) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata_Team_InventorySnapshot) ProtoMessage()    {}
func (*CDOTAMatchMetadata_Team_InventorySnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor22, []int{1, 0, 2}
}

func (m *CDOTAMatchMetadata_Team_InventorySnapshot) GetItemId() []uint32 {
	if m != nil {
		return m.ItemId
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_InventorySnapshot) GetGameTime() int32 {
	if m != nil && m.GameTime != nil {
		return *m.GameTime
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_InventorySnapshot) GetKills() uint32 {
	if m != nil && m.Kills != nil {
		return *m.Kills
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_InventorySnapshot) GetDeaths() uint32 {
	if m != nil && m.Deaths != nil {
		return *m.Deaths
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_InventorySnapshot) GetAssists() uint32 {
	if m != nil && m.Assists != nil {
		return *m.Assists
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_InventorySnapshot) GetLevel() uint32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

type CDOTAMatchMetadata_Team_AutoStyleCriteria struct {
	NameToken        *uint32  `protobuf:"varint,1,opt,name=name_token,json=nameToken" json:"name_token,omitempty"`
	Value            *float32 `protobuf:"fixed32,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CDOTAMatchMetadata_Team_AutoStyleCriteria) Reset() {
	*m = CDOTAMatchMetadata_Team_AutoStyleCriteria{}
}
func (m *CDOTAMatchMetadata_Team_AutoStyleCriteria) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata_Team_AutoStyleCriteria) ProtoMessage()    {}
func (*CDOTAMatchMetadata_Team_AutoStyleCriteria) Descriptor() ([]byte, []int) {
	return fileDescriptor22, []int{1, 0, 3}
}

func (m *CDOTAMatchMetadata_Team_AutoStyleCriteria) GetNameToken() uint32 {
	if m != nil && m.NameToken != nil {
		return *m.NameToken
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_AutoStyleCriteria) GetValue() float32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type CDOTAMatchMetadata_Team_Player struct {
	AccountId          *uint32                                      `protobuf:"varint,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	AbilityUpgrades    []uint32                                     `protobuf:"varint,2,rep,name=ability_upgrades,json=abilityUpgrades" json:"ability_upgrades,omitempty"`
	PlayerSlot         *uint32                                      `protobuf:"varint,3,opt,name=player_slot,json=playerSlot" json:"player_slot,omitempty"`
	EquippedEconItems  []*CSOEconItem                               `protobuf:"bytes,4,rep,name=equipped_econ_items,json=equippedEconItems" json:"equipped_econ_items,omitempty"`
	Kills              []*CDOTAMatchMetadata_Team_PlayerKill        `protobuf:"bytes,5,rep,name=kills" json:"kills,omitempty"`
	Items              []*CDOTAMatchMetadata_Team_ItemPurchase      `protobuf:"bytes,6,rep,name=items" json:"items,omitempty"`
	AvgKillsX16        *uint32                                      `protobuf:"varint,7,opt,name=avg_kills_x16,json=avgKillsX16" json:"avg_kills_x16,omitempty"`
	AvgDeathsX16       *uint32                                      `protobuf:"varint,8,opt,name=avg_deaths_x16,json=avgDeathsX16" json:"avg_deaths_x16,omitempty"`
	AvgAssistsX16      *uint32                                      `protobuf:"varint,9,opt,name=avg_assists_x16,json=avgAssistsX16" json:"avg_assists_x16,omitempty"`
	AvgGpmX16          *uint32                                      `protobuf:"varint,10,opt,name=avg_gpm_x16,json=avgGpmX16" json:"avg_gpm_x16,omitempty"`
	AvgXpmX16          *uint32                                      `protobuf:"varint,11,opt,name=avg_xpm_x16,json=avgXpmX16" json:"avg_xpm_x16,omitempty"`
	BestKillsX16       *uint32                                      `protobuf:"varint,12,opt,name=best_kills_x16,json=bestKillsX16" json:"best_kills_x16,omitempty"`
	BestAssistsX16     *uint32                                      `protobuf:"varint,13,opt,name=best_assists_x16,json=bestAssistsX16" json:"best_assists_x16,omitempty"`
	BestGpmX16         *uint32                                      `protobuf:"varint,14,opt,name=best_gpm_x16,json=bestGpmX16" json:"best_gpm_x16,omitempty"`
	BestXpmX16         *uint32                                      `protobuf:"varint,15,opt,name=best_xpm_x16,json=bestXpmX16" json:"best_xpm_x16,omitempty"`
	WinStreak          *uint32                                      `protobuf:"varint,16,opt,name=win_streak,json=winStreak" json:"win_streak,omitempty"`
	BestWinStreak      *uint32                                      `protobuf:"varint,17,opt,name=best_win_streak,json=bestWinStreak" json:"best_win_streak,omitempty"`
	FightScore         *float32                                     `protobuf:"fixed32,18,opt,name=fight_score,json=fightScore" json:"fight_score,omitempty"`
	FarmScore          *float32                                     `protobuf:"fixed32,19,opt,name=farm_score,json=farmScore" json:"farm_score,omitempty"`
	SupportScore       *float32                                     `protobuf:"fixed32,20,opt,name=support_score,json=supportScore" json:"support_score,omitempty"`
	PushScore          *float32                                     `protobuf:"fixed32,21,opt,name=push_score,json=pushScore" json:"push_score,omitempty"`
	LevelUpTimes       []uint32                                     `protobuf:"varint,22,rep,name=level_up_times,json=levelUpTimes" json:"level_up_times,omitempty"`
	GraphNetWorth      []float32                                    `protobuf:"fixed32,23,rep,name=graph_net_worth,json=graphNetWorth" json:"graph_net_worth,omitempty"`
	InventorySnapshot  []*CDOTAMatchMetadata_Team_InventorySnapshot `protobuf:"bytes,24,rep,name=inventory_snapshot,json=inventorySnapshot" json:"inventory_snapshot,omitempty"`
	AvgStatsCalibrated *bool                                        `protobuf:"varint,25,opt,name=avg_stats_calibrated,json=avgStatsCalibrated" json:"avg_stats_calibrated,omitempty"`
	AutoStyleCriteria  []*CDOTAMatchMetadata_Team_AutoStyleCriteria `protobuf:"bytes,26,rep,name=auto_style_criteria,json=autoStyleCriteria" json:"auto_style_criteria,omitempty"`
	EventId            *uint32                                      `protobuf:"varint,27,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	EventPoints        *uint32                                      `protobuf:"varint,28,opt,name=event_points,json=eventPoints" json:"event_points,omitempty"`
	XXX_unrecognized   []byte                                       `json:"-"`
}

func (m *CDOTAMatchMetadata_Team_Player) Reset()         { *m = CDOTAMatchMetadata_Team_Player{} }
func (m *CDOTAMatchMetadata_Team_Player) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchMetadata_Team_Player) ProtoMessage()    {}
func (*CDOTAMatchMetadata_Team_Player) Descriptor() ([]byte, []int) {
	return fileDescriptor22, []int{1, 0, 4}
}

func (m *CDOTAMatchMetadata_Team_Player) GetAccountId() uint32 {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetAbilityUpgrades() []uint32 {
	if m != nil {
		return m.AbilityUpgrades
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_Player) GetPlayerSlot() uint32 {
	if m != nil && m.PlayerSlot != nil {
		return *m.PlayerSlot
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetEquippedEconItems() []*CSOEconItem {
	if m != nil {
		return m.EquippedEconItems
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_Player) GetKills() []*CDOTAMatchMetadata_Team_PlayerKill {
	if m != nil {
		return m.Kills
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_Player) GetItems() []*CDOTAMatchMetadata_Team_ItemPurchase {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_Player) GetAvgKillsX16() uint32 {
	if m != nil && m.AvgKillsX16 != nil {
		return *m.AvgKillsX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetAvgDeathsX16() uint32 {
	if m != nil && m.AvgDeathsX16 != nil {
		return *m.AvgDeathsX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetAvgAssistsX16() uint32 {
	if m != nil && m.AvgAssistsX16 != nil {
		return *m.AvgAssistsX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetAvgGpmX16() uint32 {
	if m != nil && m.AvgGpmX16 != nil {
		return *m.AvgGpmX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetAvgXpmX16() uint32 {
	if m != nil && m.AvgXpmX16 != nil {
		return *m.AvgXpmX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetBestKillsX16() uint32 {
	if m != nil && m.BestKillsX16 != nil {
		return *m.BestKillsX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetBestAssistsX16() uint32 {
	if m != nil && m.BestAssistsX16 != nil {
		return *m.BestAssistsX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetBestGpmX16() uint32 {
	if m != nil && m.BestGpmX16 != nil {
		return *m.BestGpmX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetBestXpmX16() uint32 {
	if m != nil && m.BestXpmX16 != nil {
		return *m.BestXpmX16
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetWinStreak() uint32 {
	if m != nil && m.WinStreak != nil {
		return *m.WinStreak
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetBestWinStreak() uint32 {
	if m != nil && m.BestWinStreak != nil {
		return *m.BestWinStreak
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetFightScore() float32 {
	if m != nil && m.FightScore != nil {
		return *m.FightScore
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetFarmScore() float32 {
	if m != nil && m.FarmScore != nil {
		return *m.FarmScore
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetSupportScore() float32 {
	if m != nil && m.SupportScore != nil {
		return *m.SupportScore
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetPushScore() float32 {
	if m != nil && m.PushScore != nil {
		return *m.PushScore
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetLevelUpTimes() []uint32 {
	if m != nil {
		return m.LevelUpTimes
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_Player) GetGraphNetWorth() []float32 {
	if m != nil {
		return m.GraphNetWorth
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_Player) GetInventorySnapshot() []*CDOTAMatchMetadata_Team_InventorySnapshot {
	if m != nil {
		return m.InventorySnapshot
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_Player) GetAvgStatsCalibrated() bool {
	if m != nil && m.AvgStatsCalibrated != nil {
		return *m.AvgStatsCalibrated
	}
	return false
}

func (m *CDOTAMatchMetadata_Team_Player) GetAutoStyleCriteria() []*CDOTAMatchMetadata_Team_AutoStyleCriteria {
	if m != nil {
		return m.AutoStyleCriteria
	}
	return nil
}

func (m *CDOTAMatchMetadata_Team_Player) GetEventId() uint32 {
	if m != nil && m.EventId != nil {
		return *m.EventId
	}
	return 0
}

func (m *CDOTAMatchMetadata_Team_Player) GetEventPoints() uint32 {
	if m != nil && m.EventPoints != nil {
		return *m.EventPoints
	}
	return 0
}

type CDOTAMatchPrivateMetadata struct {
	Teams            []*CDOTAMatchPrivateMetadata_Team `protobuf:"bytes,1,rep,name=teams" json:"teams,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *CDOTAMatchPrivateMetadata) Reset()                    { *m = CDOTAMatchPrivateMetadata{} }
func (m *CDOTAMatchPrivateMetadata) String() string            { return proto.CompactTextString(m) }
func (*CDOTAMatchPrivateMetadata) ProtoMessage()               {}
func (*CDOTAMatchPrivateMetadata) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{2} }

func (m *CDOTAMatchPrivateMetadata) GetTeams() []*CDOTAMatchPrivateMetadata_Team {
	if m != nil {
		return m.Teams
	}
	return nil
}

type CDOTAMatchPrivateMetadata_Team struct {
	DotaTeam         *uint32                                    `protobuf:"varint,1,opt,name=dota_team,json=dotaTeam" json:"dota_team,omitempty"`
	Players          []*CDOTAMatchPrivateMetadata_Team_Player   `protobuf:"bytes,2,rep,name=players" json:"players,omitempty"`
	Buildings        []*CDOTAMatchPrivateMetadata_Team_Building `protobuf:"bytes,3,rep,name=buildings" json:"buildings,omitempty"`
	XXX_unrecognized []byte                                     `json:"-"`
}

func (m *CDOTAMatchPrivateMetadata_Team) Reset()         { *m = CDOTAMatchPrivateMetadata_Team{} }
func (m *CDOTAMatchPrivateMetadata_Team) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchPrivateMetadata_Team) ProtoMessage()    {}
func (*CDOTAMatchPrivateMetadata_Team) Descriptor() ([]byte, []int) {
	return fileDescriptor22, []int{2, 0}
}

func (m *CDOTAMatchPrivateMetadata_Team) GetDotaTeam() uint32 {
	if m != nil && m.DotaTeam != nil {
		return *m.DotaTeam
	}
	return 0
}

func (m *CDOTAMatchPrivateMetadata_Team) GetPlayers() []*CDOTAMatchPrivateMetadata_Team_Player {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *CDOTAMatchPrivateMetadata_Team) GetBuildings() []*CDOTAMatchPrivateMetadata_Team_Building {
	if m != nil {
		return m.Buildings
	}
	return nil
}

type CDOTAMatchPrivateMetadata_Team_Player struct {
	AccountId        *uint32 `protobuf:"varint,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	PlayerSlot       *uint32 `protobuf:"varint,2,opt,name=player_slot,json=playerSlot" json:"player_slot,omitempty"`
	PositionStream   []byte  `protobuf:"bytes,3,opt,name=position_stream,json=positionStream" json:"position_stream,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDOTAMatchPrivateMetadata_Team_Player) Reset()         { *m = CDOTAMatchPrivateMetadata_Team_Player{} }
func (m *CDOTAMatchPrivateMetadata_Team_Player) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchPrivateMetadata_Team_Player) ProtoMessage()    {}
func (*CDOTAMatchPrivateMetadata_Team_Player) Descriptor() ([]byte, []int) {
	return fileDescriptor22, []int{2, 0, 0}
}

func (m *CDOTAMatchPrivateMetadata_Team_Player) GetAccountId() uint32 {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return 0
}

func (m *CDOTAMatchPrivateMetadata_Team_Player) GetPlayerSlot() uint32 {
	if m != nil && m.PlayerSlot != nil {
		return *m.PlayerSlot
	}
	return 0
}

func (m *CDOTAMatchPrivateMetadata_Team_Player) GetPositionStream() []byte {
	if m != nil {
		return m.PositionStream
	}
	return nil
}

type CDOTAMatchPrivateMetadata_Team_Building struct {
	UnitName         *string  `protobuf:"bytes,1,opt,name=unit_name,json=unitName" json:"unit_name,omitempty"`
	PositionQuantX   *uint32  `protobuf:"varint,2,opt,name=position_quant_x,json=positionQuantX" json:"position_quant_x,omitempty"`
	PositionQuantY   *uint32  `protobuf:"varint,3,opt,name=position_quant_y,json=positionQuantY" json:"position_quant_y,omitempty"`
	DeathTime        *float32 `protobuf:"fixed32,4,opt,name=death_time,json=deathTime" json:"death_time,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CDOTAMatchPrivateMetadata_Team_Building) Reset() {
	*m = CDOTAMatchPrivateMetadata_Team_Building{}
}
func (m *CDOTAMatchPrivateMetadata_Team_Building) String() string { return proto.CompactTextString(m) }
func (*CDOTAMatchPrivateMetadata_Team_Building) ProtoMessage()    {}
func (*CDOTAMatchPrivateMetadata_Team_Building) Descriptor() ([]byte, []int) {
	return fileDescriptor22, []int{2, 0, 1}
}

func (m *CDOTAMatchPrivateMetadata_Team_Building) GetUnitName() string {
	if m != nil && m.UnitName != nil {
		return *m.UnitName
	}
	return ""
}

func (m *CDOTAMatchPrivateMetadata_Team_Building) GetPositionQuantX() uint32 {
	if m != nil && m.PositionQuantX != nil {
		return *m.PositionQuantX
	}
	return 0
}

func (m *CDOTAMatchPrivateMetadata_Team_Building) GetPositionQuantY() uint32 {
	if m != nil && m.PositionQuantY != nil {
		return *m.PositionQuantY
	}
	return 0
}

func (m *CDOTAMatchPrivateMetadata_Team_Building) GetDeathTime() float32 {
	if m != nil && m.DeathTime != nil {
		return *m.DeathTime
	}
	return 0
}

func init() {
	proto.RegisterType((*CDOTAMatchMetadataFile)(nil), "dota.CDOTAMatchMetadataFile")
	proto.RegisterType((*CDOTAMatchMetadata)(nil), "dota.CDOTAMatchMetadata")
	proto.RegisterType((*CDOTAMatchMetadata_Team)(nil), "dota.CDOTAMatchMetadata.Team")
	proto.RegisterType((*CDOTAMatchMetadata_Team_PlayerKill)(nil), "dota.CDOTAMatchMetadata.Team.PlayerKill")
	proto.RegisterType((*CDOTAMatchMetadata_Team_ItemPurchase)(nil), "dota.CDOTAMatchMetadata.Team.ItemPurchase")
	proto.RegisterType((*CDOTAMatchMetadata_Team_InventorySnapshot)(nil), "dota.CDOTAMatchMetadata.Team.InventorySnapshot")
	proto.RegisterType((*CDOTAMatchMetadata_Team_AutoStyleCriteria)(nil), "dota.CDOTAMatchMetadata.Team.AutoStyleCriteria")
	proto.RegisterType((*CDOTAMatchMetadata_Team_Player)(nil), "dota.CDOTAMatchMetadata.Team.Player")
	proto.RegisterType((*CDOTAMatchPrivateMetadata)(nil), "dota.CDOTAMatchPrivateMetadata")
	proto.RegisterType((*CDOTAMatchPrivateMetadata_Team)(nil), "dota.CDOTAMatchPrivateMetadata.Team")
	proto.RegisterType((*CDOTAMatchPrivateMetadata_Team_Player)(nil), "dota.CDOTAMatchPrivateMetadata.Team.Player")
	proto.RegisterType((*CDOTAMatchPrivateMetadata_Team_Building)(nil), "dota.CDOTAMatchPrivateMetadata.Team.Building")
}

func init() { proto.RegisterFile("dota_match_metadata.proto", fileDescriptor22) }

var fileDescriptor22 = []byte{
	// 1412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xcd, 0x76, 0x13, 0x47,
	0x16, 0x1e, 0x49, 0x96, 0x2d, 0x5d, 0x49, 0x16, 0x6a, 0x0c, 0xb4, 0xc5, 0x00, 0x1a, 0xe3, 0xc3,
	0x68, 0x98, 0x83, 0x19, 0x3c, 0x33, 0x2c, 0x66, 0xc1, 0x19, 0x63, 0x0c, 0xd1, 0xe1, 0xcf, 0x69,
	0x9b, 0x03, 0xd9, 0xa4, 0x4e, 0xa9, 0xbb, 0x68, 0xd5, 0x71, 0x57, 0x77, 0xd3, 0x55, 0x2d, 0xdb,
	0xbb, 0x3c, 0x47, 0x76, 0x59, 0x65, 0x97, 0xd7, 0x48, 0x1e, 0x22, 0x2f, 0x90, 0xb7, 0xc8, 0xb9,
	0xb7, 0xba, 0x65, 0xd9, 0x22, 0xc4, 0xcb, 0xfa, 0xee, 0x77, 0x7f, 0xea, 0xfe, 0x55, 0xc1, 0x7a,
	0x90, 0x18, 0xce, 0x14, 0x37, 0xfe, 0x84, 0x29, 0x61, 0x78, 0xc0, 0x0d, 0xdf, 0x4a, 0xb3, 0xc4,
	0x24, 0xce, 0x12, 0x8a, 0xfa, 0xd7, 0xc6, 0x5c, 0x0b, 0x16, 0xfa, 0x4a, 0x68, 0xcd, 0x43, 0xa1,
	0xad, 0xb0, 0xff, 0x80, 0xf4, 0xce, 0x60, 0xe6, 0x27, 0x4a, 0x25, 0x71, 0x69, 0x89, 0xc7, 0x3c,
	0x14, 0x4a, 0xc4, 0xc6, 0xd2, 0x37, 0x7e, 0xaa, 0xc0, 0xf5, 0xdd, 0x67, 0x6f, 0x0f, 0x77, 0x5e,
	0xa3, 0xfc, 0x75, 0xe1, 0xe8, 0xb9, 0x8c, 0x84, 0xe3, 0xc2, 0xca, 0x54, 0x64, 0x5a, 0x26, 0xb1,
	0x5b, 0x19, 0x54, 0x87, 0x75, 0xaf, 0x3c, 0x3a, 0xeb, 0xd0, 0xb0, 0xe6, 0x64, 0xe0, 0x56, 0x07,
	0xd5, 0xe1, 0x92, 0xb7, 0x42, 0xe7, 0x51, 0xe0, 0xfc, 0x07, 0x1a, 0x65, 0xb4, 0x6e, 0x6d, 0x50,
	0x19, 0xb6, 0xb6, 0xdd, 0x2d, 0x8c, 0x68, 0x6b, 0xd1, 0x89, 0x37, 0x63, 0x3a, 0xff, 0x80, 0x2b,
	0x69, 0x26, 0xa7, 0xdc, 0x88, 0xd9, 0x5d, 0xdd, 0xfa, 0xa0, 0x32, 0x6c, 0x7b, 0xdd, 0x02, 0x2f,
	0x95, 0x36, 0x7e, 0xeb, 0x81, 0xb3, 0x68, 0xcb, 0xf9, 0x37, 0xd4, 0x8d, 0xe0, 0x4a, 0xbb, 0x95,
	0x41, 0x6d, 0xd8, 0xda, 0xbe, 0xf5, 0x47, 0x4e, 0xb7, 0x0e, 0x05, 0x57, 0x9e, 0xe5, 0x3a, 0x3b,
	0xd0, 0x96, 0x46, 0x28, 0x96, 0x89, 0x63, 0x9e, 0x05, 0xda, 0xad, 0x92, 0xee, 0xed, 0x42, 0xf7,
	0x55, 0x32, 0x1e, 0x9f, 0x1e, 0x4a, 0x25, 0x02, 0x8f, 0xe4, 0xcf, 0x84, 0xe1, 0x32, 0xd2, 0x5e,
	0x0b, 0x75, 0x2c, 0xa4, 0x31, 0x15, 0x11, 0xf2, 0x30, 0x15, 0x78, 0xdf, 0x65, 0x6f, 0x85, 0xce,
	0xa3, 0xc0, 0xb9, 0x0f, 0xbd, 0x4c, 0xa4, 0x49, 0x66, 0x58, 0x1e, 0x1b, 0x19, 0x31, 0x23, 0x95,
	0x70, 0x97, 0x88, 0xd3, 0xb5, 0x82, 0x77, 0x88, 0xa3, 0x07, 0xe7, 0xbf, 0x70, 0x43, 0x4c, 0x45,
	0x6c, 0x58, 0xc8, 0x95, 0x60, 0x7e, 0xae, 0x4d, 0xa2, 0x98, 0xe1, 0xe3, 0x48, 0x14, 0x79, 0x58,
	0x23, 0xf1, 0x0b, 0xae, 0xc4, 0x2e, 0x09, 0x0f, 0x51, 0xd6, 0xff, 0xb5, 0x0b, 0x4b, 0x78, 0x21,
	0xe7, 0x26, 0x34, 0xa9, 0xee, 0x78, 0x2f, 0xb7, 0x32, 0xa8, 0x0c, 0x3b, 0x5e, 0x03, 0x01, 0x12,
	0x3e, 0x81, 0x95, 0x34, 0xe2, 0xa7, 0x22, 0x2b, 0x6f, 0xb8, 0xf9, 0xc5, 0xec, 0x6c, 0xed, 0x13,
	0xd9, 0x2b, 0x95, 0xb0, 0x3a, 0x61, 0xc6, 0xd3, 0x09, 0x13, 0x27, 0xa9, 0xc8, 0xa4, 0x88, 0x7d,
	0xe1, 0xd6, 0x06, 0xb5, 0x61, 0xd5, 0xeb, 0x12, 0xbe, 0x37, 0x83, 0xf1, 0xce, 0x96, 0x1a, 0x26,
	0x51, 0xc0, 0x04, 0xcf, 0x62, 0x11, 0xb8, 0x4b, 0x73, 0xdc, 0x17, 0x49, 0x14, 0xec, 0x11, 0xec,
	0xdc, 0x03, 0x0b, 0xb1, 0x58, 0x18, 0x76, 0x9c, 0x64, 0x66, 0xe2, 0xd6, 0x89, 0xd9, 0x21, 0xf8,
	0x8d, 0x30, 0xef, 0x11, 0x74, 0x36, 0xa0, 0xe3, 0x2b, 0xf6, 0x51, 0x66, 0xda, 0xb0, 0x54, 0xfa,
	0x47, 0xee, 0xf2, 0xa0, 0x32, 0x6c, 0x78, 0x2d, 0x5f, 0x3d, 0x47, 0x6c, 0x5f, 0xfa, 0x47, 0xce,
	0x43, 0x58, 0xf3, 0x15, 0xf3, 0x79, 0x6a, 0xb8, 0x8c, 0x99, 0x0d, 0x1c, 0x4b, 0xb2, 0x42, 0xa9,
	0xe8, 0xf9, 0x6a, 0xd7, 0x8a, 0xec, 0xd5, 0x46, 0x81, 0x73, 0x03, 0x56, 0x7c, 0xc5, 0xc6, 0x3c,
	0xd6, 0x6e, 0x63, 0x50, 0x1b, 0x76, 0xbc, 0x65, 0x5f, 0x3d, 0xe5, 0x31, 0x15, 0xd4, 0x57, 0xe4,
	0x47, 0xbb, 0x4d, 0x92, 0xac, 0xf8, 0x0a, 0x7d, 0x68, 0xe7, 0x16, 0x00, 0x8a, 0x44, 0xcc, 0x23,
	0x73, 0xea, 0x02, 0x99, 0x6e, 0xfa, 0x6a, 0xdf, 0x02, 0xfd, 0x5d, 0x00, 0x6b, 0xfe, 0xa5, 0x8c,
	0x22, 0xe7, 0x0e, 0xb4, 0xa6, 0xd2, 0x37, 0x52, 0x31, 0x1d, 0x25, 0xa6, 0xa8, 0x09, 0x58, 0xe8,
	0x20, 0x4a, 0x8c, 0xb3, 0x06, 0x75, 0x3f, 0xc9, 0x63, 0xe3, 0x56, 0x49, 0x64, 0x0f, 0xfd, 0x57,
	0xd0, 0x1e, 0x19, 0xa1, 0xf6, 0xf3, 0xcc, 0x9f, 0x70, 0x2d, 0x30, 0x4e, 0x6a, 0x51, 0x19, 0x14,
	0x26, 0x96, 0xf1, 0x38, 0x0a, 0x9c, 0xbb, 0xd0, 0x49, 0x0b, 0x92, 0xed, 0x2c, 0x34, 0x53, 0xf7,
	0xda, 0x25, 0x88, 0x6d, 0xd5, 0xff, 0xb1, 0x02, 0xbd, 0x51, 0x8c, 0x9d, 0x93, 0x64, 0xa7, 0x07,
	0x31, 0x4f, 0xf5, 0x24, 0x31, 0xe7, 0x6d, 0xd6, 0xe6, 0x6c, 0xde, 0x84, 0x26, 0xf5, 0xdf, 0x9c,
	0xbd, 0x06, 0x02, 0xd4, 0xa2, 0x6b, 0x50, 0x3f, 0x92, 0x51, 0xa4, 0xa9, 0xcd, 0x3b, 0x9e, 0x3d,
	0x38, 0xd7, 0x61, 0x39, 0x10, 0xdc, 0x4c, 0x34, 0x75, 0x76, 0xc7, 0x2b, 0x4e, 0xb8, 0x3c, 0xb8,
	0xd6, 0x52, 0x1b, 0x4d, 0x0d, 0xdc, 0xf1, 0xca, 0x23, 0xda, 0x89, 0xc4, 0x54, 0x44, 0x54, 0xc6,
	0x8e, 0x67, 0x0f, 0xfd, 0xaf, 0xa0, 0xb7, 0x93, 0x9b, 0xe4, 0xc0, 0x9c, 0x46, 0x62, 0x37, 0x93,
	0x46, 0x64, 0x92, 0x63, 0xc2, 0x63, 0x8a, 0x27, 0x39, 0x12, 0x71, 0x71, 0xff, 0x26, 0x22, 0x87,
	0x08, 0xa0, 0xa5, 0x29, 0x8f, 0x72, 0x1b, 0x6a, 0xd5, 0xb3, 0x87, 0xfe, 0x2f, 0x4d, 0x58, 0xb6,
	0x75, 0x40, 0x7d, 0xee, 0x53, 0x5e, 0xcf, 0xf2, 0xd7, 0x2c, 0x90, 0x51, 0x80, 0x7d, 0xcd, 0xc7,
	0x32, 0x92, 0xe6, 0x94, 0xe5, 0x69, 0x98, 0xf1, 0x40, 0xd8, 0x01, 0xe9, 0x78, 0xdd, 0x02, 0x7f,
	0x57, 0xc0, 0x58, 0xcd, 0xa2, 0xa9, 0xa8, 0x9a, 0x36, 0x05, 0x60, 0x21, 0xaa, 0xe6, 0x0e, 0x5c,
	0x15, 0x9f, 0x72, 0x99, 0xa6, 0x22, 0x60, 0xc2, 0x4f, 0x62, 0x86, 0x29, 0xd5, 0xd4, 0xfa, 0xad,
	0xed, 0x5e, 0x31, 0x6f, 0x07, 0x6f, 0xf7, 0xfc, 0x24, 0xc6, 0xfa, 0x7a, 0xbd, 0x92, 0x5d, 0x22,
	0xda, 0x79, 0x52, 0x26, 0xb8, 0x4e, 0x4a, 0xc3, 0xcb, 0x0c, 0x29, 0xb6, 0x5a, 0x59, 0x8a, 0xff,
	0x43, 0xdd, 0x3a, 0x5d, 0x26, 0xfd, 0xfb, 0x5f, 0xd6, 0x9f, 0xef, 0x32, 0xcf, 0x2a, 0xe2, 0xa4,
	0xf1, 0x69, 0xc8, 0xc8, 0x1c, 0x3b, 0x79, 0xf4, 0xb8, 0x18, 0x9f, 0x16, 0x9f, 0x86, 0xe8, 0x48,
	0x7f, 0x78, 0xf4, 0xd8, 0xd9, 0x84, 0x55, 0xe4, 0xd8, 0x32, 0x13, 0xa9, 0x41, 0xa4, 0x36, 0x9f,
	0x86, 0xcf, 0x08, 0x44, 0xd6, 0x3d, 0xe8, 0x22, 0xab, 0xa8, 0x39, 0xd1, 0x9a, 0x44, 0x43, 0x07,
	0x3b, 0x16, 0x45, 0xde, 0x6d, 0x40, 0xe3, 0x2c, 0x4c, 0x15, 0x71, 0x8a, 0x99, 0xe2, 0xd3, 0xf0,
	0x45, 0xaa, 0xe6, 0xe4, 0x27, 0x85, 0xbc, 0x35, 0x93, 0x7f, 0xb0, 0xf2, 0x4d, 0x58, 0x1d, 0x0b,
	0x6d, 0xe6, 0x42, 0x6e, 0xdb, 0x68, 0x10, 0x9d, 0xc5, 0x3c, 0x84, 0x2b, 0xc4, 0x9a, 0x0f, 0xa7,
	0x43, 0x3c, 0xd2, 0x9e, 0x8b, 0x67, 0x00, 0xa4, 0x39, 0x0b, 0x68, 0xd5, 0x16, 0x1a, 0xb1, 0x22,
	0xa2, 0x92, 0x51, 0x86, 0xd4, 0x3d, 0x63, 0x14, 0x31, 0xdd, 0x02, 0x38, 0x96, 0x31, 0xd3, 0x26,
	0x13, 0xfc, 0xc8, 0xbd, 0x62, 0x43, 0x3e, 0x96, 0xf1, 0x01, 0x01, 0x98, 0x1a, 0x32, 0x30, 0xc7,
	0xe9, 0xd9, 0xd4, 0x20, 0xfc, 0x7e, 0xc6, 0xbb, 0x03, 0xad, 0x8f, 0x32, 0x9c, 0x18, 0xa6, 0xfd,
	0x24, 0x13, 0xae, 0x43, 0x3d, 0x0e, 0x04, 0x1d, 0x20, 0x82, 0x7e, 0x3e, 0xf2, 0x4c, 0x15, 0xf2,
	0xab, 0x24, 0x6f, 0x22, 0x62, 0xc5, 0x77, 0xa1, 0xa3, 0xf3, 0x94, 0xde, 0x1f, 0xcb, 0x58, 0x23,
	0x46, 0xbb, 0x00, 0x67, 0x36, 0xd2, 0x5c, 0x4f, 0x0a, 0xc6, 0x35, 0x6b, 0x03, 0x11, 0x2b, 0xde,
	0x84, 0x55, 0x1a, 0x4f, 0x96, 0xa7, 0xb4, 0x14, 0xb4, 0x7b, 0x9d, 0xe6, 0xa3, 0x4d, 0xe8, 0xbb,
	0x14, 0x17, 0x83, 0xfe, 0xdc, 0x22, 0xbf, 0xf1, 0xb9, 0x45, 0xfe, 0x2d, 0x38, 0xb2, 0x5c, 0x46,
	0x4c, 0x17, 0xdb, 0xc8, 0x75, 0xa9, 0x5b, 0x1f, 0xfe, 0x49, 0xb7, 0x5e, 0x5c, 0x62, 0x5e, 0x4f,
	0x2e, 0xec, 0xb5, 0x7f, 0xc1, 0x1a, 0x36, 0x8b, 0x36, 0xdc, 0x68, 0xe6, 0xf3, 0x48, 0x8e, 0x33,
	0x6e, 0x44, 0xe0, 0xae, 0xd3, 0x7b, 0xe1, 0xf0, 0x69, 0x78, 0x80, 0xa2, 0xdd, 0x99, 0xc4, 0x61,
	0x70, 0x95, 0xe7, 0x26, 0x61, 0x1a, 0xd7, 0x0e, 0xf3, 0x8b, 0xbd, 0xe3, 0xf6, 0x2f, 0x13, 0xd2,
	0xc2, 0xba, 0xf2, 0x7a, 0x7c, 0x61, 0x83, 0xad, 0x43, 0xc3, 0xbe, 0xeb, 0x32, 0x70, 0x6f, 0xda,
	0x3d, 0x48, 0xe7, 0x51, 0xe0, 0xfc, 0x0d, 0xda, 0x56, 0x94, 0x26, 0x32, 0x36, 0xda, 0xfd, 0xab,
	0x9d, 0x35, 0xc2, 0xf6, 0x09, 0xda, 0xf8, 0x61, 0x09, 0xd6, 0xcf, 0xdc, 0xef, 0x9f, 0xff, 0x09,
	0x39, 0xff, 0x3b, 0xff, 0xe5, 0x59, 0x78, 0xd4, 0x2f, 0xf0, 0xe7, 0x7f, 0x3e, 0xfd, 0x9f, 0x6b,
	0x97, 0xf9, 0x38, 0xec, 0x5d, 0xfc, 0x38, 0xfc, 0xf3, 0x32, 0x3e, 0x16, 0xfe, 0x0f, 0x2f, 0xa1,
	0x39, 0xce, 0x65, 0x14, 0xc8, 0x38, 0xd4, 0xf4, 0x71, 0x68, 0x6d, 0x3f, 0xb8, 0x94, 0xa1, 0xa7,
	0x85, 0x96, 0x77, 0xa6, 0xdf, 0xff, 0x74, 0xd9, 0xed, 0x7e, 0x61, 0x65, 0x57, 0x17, 0x56, 0xf6,
	0xdf, 0xa1, 0x9b, 0x26, 0x5a, 0x1a, 0x99, 0x14, 0x83, 0xa8, 0x68, 0xaf, 0xb7, 0xbd, 0xd5, 0x12,
	0xa6, 0x49, 0x54, 0xfd, 0xef, 0x2b, 0xd0, 0x28, 0x43, 0xc1, 0x84, 0xe5, 0xb1, 0x34, 0x0c, 0x9f,
	0x21, 0x72, 0xda, 0xf4, 0x1a, 0x08, 0xbc, 0xe1, 0x4a, 0xe0, 0xa2, 0x99, 0x99, 0xfc, 0x94, 0xf3,
	0xd8, 0xb0, 0x93, 0xc2, 0xf1, 0xcc, 0xe6, 0xd7, 0x08, 0x7f, 0xf8, 0x0c, 0xf3, 0xb4, 0x78, 0x55,
	0xce, 0x33, 0xbf, 0xc1, 0x6b, 0xd2, 0xb2, 0x3d, 0xfb, 0x3f, 0x56, 0xbd, 0x26, 0x21, 0x38, 0x7d,
	0x4f, 0x6b, 0xdf, 0x55, 0xfe, 0xf2, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x8d, 0x5c, 0xfa,
	0x2d, 0x0c, 0x00, 0x00,
}
