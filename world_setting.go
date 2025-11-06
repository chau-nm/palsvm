package palsvm

import (
	"fmt"
	"strconv"
	"strings"
)

// WorldSetting The world settings configuration is defined in the official documentation at https://docs.palworldgame.com/settings-and-operation/configuration.
type WorldSetting struct {
	Difficulty                           string   `json:"Difficulty"`
	RandomizerType                       string   `json:"RandomizerType"`
	RandomizerSeed                       string   `json:"RandomizerSeed"`
	BIsRandomizerPalLevelRandom          bool     `json:"bIsRandomizerPalLevelRandom"`
	DayTimeSpeedRate                     float64  `json:"DayTimeSpeedRate"`
	NightTimeSpeedRate                   float64  `json:"NightTimeSpeedRate"`
	ExpRate                              float64  `json:"ExpRate"`
	PalCaptureRate                       float64  `json:"PalCaptureRate"`
	PalSpawnNumRate                      float64  `json:"PalSpawnNumRate"`
	PalDamageRateAttack                  float64  `json:"PalDamageRateAttack"`
	PalDamageRateDefense                 float64  `json:"PalDamageRateDefense"`
	PlayerDamageRateAttack               float64  `json:"PlayerDamageRateAttack"`
	PlayerDamageRateDefense              float64  `json:"PlayerDamageRateDefense"`
	PlayerStomachDecreaceRate            float64  `json:"PlayerStomachDecreaceRate"`
	PlayerStaminaDecreaceRate            float64  `json:"PlayerStaminaDecreaceRate"`
	PlayerAutoHPRegeneRate               float64  `json:"PlayerAutoHPRegeneRate"`
	PlayerAutoHpRegeneRateInSleep        float64  `json:"PlayerAutoHpRegeneRateInSleep"`
	PalStomachDecreaceRate               float64  `json:"PalStomachDecreaceRate"`
	PalStaminaDecreaceRate               float64  `json:"PalStaminaDecreaceRate"`
	PalAutoHPRegeneRate                  float64  `json:"PalAutoHPRegeneRate"`
	PalAutoHpRegeneRateInSleep           float64  `json:"PalAutoHpRegeneRateInSleep"`
	BuildObjectHpRate                    float64  `json:"BuildObjectHpRate"`
	BuildObjectDamageRate                float64  `json:"BuildObjectDamageRate"`
	BuildObjectDeteriorationDamageRate   float64  `json:"BuildObjectDeteriorationDamageRate"`
	CollectionDropRate                   float64  `json:"CollectionDropRate"`
	CollectionObjectHpRate               float64  `json:"CollectionObjectHpRate"`
	CollectionObjectRespawnSpeedRate     float64  `json:"CollectionObjectRespawnSpeedRate"`
	EnemyDropItemRate                    float64  `json:"EnemyDropItemRate"`
	DeathPenalty                         string   `json:"DeathPenalty"`
	BEnablePlayerToPlayerDamage          bool     `json:"bEnablePlayerToPlayerDamage"`
	BEnableFriendlyFire                  bool     `json:"bEnableFriendlyFire"`
	BEnableInvaderEnemy                  bool     `json:"bEnableInvaderEnemy"`
	EnablePredatorBossPal                bool     `json:"EnablePredatorBossPal"`
	BActiveUNKO                          bool     `json:"bActiveUNKO"`
	BEnableAimAssistPad                  bool     `json:"bEnableAimAssistPad"`
	BEnableAimAssistKeyboard             bool     `json:"bEnableAimAssistKeyboard"`
	DropItemMaxNum                       int      `json:"DropItemMaxNum"`
	DropItemMaxNumUNKO                   int      `json:"DropItemMaxNum_UNKO"`
	BaseCampMaxNum                       int      `json:"BaseCampMaxNum"`
	BaseCampMaxNumInGuild                int      `json:"BaseCampMaxNumInGuild"`
	BaseCampWorkerMaxNum                 int      `json:"BaseCampWorkerMaxNum"`
	DropItemAliveMaxHours                float64  `json:"DropItemAliveMaxHours"`
	BAutoResetGuildNoOnlinePlayers       bool     `json:"bAutoResetGuildNoOnlinePlayers"`
	AutoResetGuildTimeNoOnlinePlayers    float64  `json:"AutoResetGuildTimeNoOnlinePlayers"`
	GuildPlayerMaxNum                    int      `json:"GuildPlayerMaxNum"`
	PalEggDefaultHatchingTime            float64  `json:"PalEggDefaultHatchingTime"`
	WorkSpeedRate                        float64  `json:"WorkSpeedRate"`
	AutoSaveSpan                         float64  `json:"AutoSaveSpan"`
	CrossplayPlatforms                   []string `json:"CrossplayPlatforms"`
	LogFormatType                        string   `json:"LogFormatType"`
	BIsMultiplay                         bool     `json:"bIsMultiplay"`
	BIsPvP                               bool     `json:"bIsPvP"`
	BHardcore                            bool     `json:"bHardcore"`
	BPalLost                             bool     `json:"bPalLost"`
	BCharacterRecreateInHardcore         bool     `json:"bCharacterRecreateInHardcore"`
	BCanPickupOtherGuildDeathPenaltyDrop bool     `json:"bCanPickupOtherGuildDeathPenaltyDrop"`
	BEnableNonLoginPenalty               bool     `json:"bEnableNonLoginPenalty"`
	BEnableFastTravel                    bool     `json:"bEnableFastTravel"`
	BIsStartLocationSelectByMap          bool     `json:"bIsStartLocationSelectByMap"`
	BExistPlayerAfterLogout              bool     `json:"bExistPlayerAfterLogout"`
	BEnableDefenseOtherGuildPlayer       bool     `json:"bEnableDefenseOtherGuildPlayer"`
	BInvisibleOtherGuildBaseCampAreaFX   bool     `json:"bInvisibleOtherGuildBaseCampAreaFX"`
	BBuildAreaLimit                      bool     `json:"bBuildAreaLimit"`
	ItemWeightRate                       float64  `json:"ItemWeightRate"`
	BShowPlayerList                      bool     `json:"bShowPlayerList"`
	CoopPlayerMaxNum                     int      `json:"CoopPlayerMaxNum"`
	ServerPlayerMaxNum                   int      `json:"ServerPlayerMaxNum"`
	ServerName                           string   `json:"ServerName"`
	ServerDescription                    string   `json:"ServerDescription"`
	AdminPassword                        string   `json:"AdminPassword"`
	ServerPassword                       string   `json:"ServerPassword"`
	PublicPort                           int      `json:"PublicPort"`
	PublicIP                             string   `json:"PublicIP"`
	RCONEnabled                          bool     `json:"RCONEnabled"`
	RCONPort                             int      `json:"RCONPort"`
	RESTAPIEnabled                       bool     `json:"RESTAPIEnabled"`
	RESTAPIPort                          int      `json:"RESTAPIPort"`
	BIsUseBackupSaveData                 bool     `json:"bIsUseBackupSaveData"`
	Region                               string   `json:"Region"`
	BUseAuth                             bool     `json:"bUseAuth"`
	BanListURL                           string   `json:"BanListURL"`
	SupplyDropSpan                       int      `json:"SupplyDropSpan"`
	ChatPostLimitPerMinute               int      `json:"ChatPostLimitPerMinute"`
	MaxBuildingLimitNum                  int      `json:"MaxBuildingLimitNum"`
	ServerReplicatePawnCullDistance      float64  `json:"ServerReplicatePawnCullDistance"`
	BAllowGlobalPalboxExport             bool     `json:"bAllowGlobalPalboxExport"`
	BAllowGlobalPalboxImport             bool     `json:"bAllowGlobalPalboxImport"`
	EquipmentDurabilityDamageRate        float64  `json:"EquipmentDurabilityDamageRate"`
	ItemContainerForceMarkDirtyInterval  float64  `json:"ItemContainerForceMarkDirtyInterval"`
	ItemCorruptionMultiplier             float64  `json:"ItemCorruptionMultiplier"`
	//BIsShowJoinLeftMessage               bool     `json:"bIsShowJoinLeftMessage"`
}

// NewWorldSettingFromFileContentConfig parses a raw config string (comma-separated key=value)
// and returns a populated WorldSetting struct pointer.
func NewWorldSettingFromFileContentConfig(fileContent string) *WorldSetting {
	optionSettingString := extractOptionSettingsBlock(fileContent)

	ws := &WorldSetting{}

	pairs := parseKeyValuePairs(optionSettingString)

	for key, value := range pairs {
		applyConfigPair(key, value, ws)
	}
	return ws
}

// String returns the Palworld configuration string representation.
// It implements the fmt.Stringer interface.
func (ws *WorldSetting) String() string {
	return fmt.Sprintf("[/Script/Pal.PalGameWorldSettings]\nOptionSettings=(%s)", ws.buildOptionString())
}

// buildOptionString constructs the OptionSettings parameter string from WorldSetting fields.
// It formats all configuration values in the Palworld config format.
func (ws *WorldSetting) buildOptionString() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("Difficulty=%s,", ws.Difficulty))
	builder.WriteString(fmt.Sprintf("RandomizerType=%s,", ws.RandomizerType))
	builder.WriteString(fmt.Sprintf("RandomizerSeed=\"%s\",", ws.RandomizerSeed))
	builder.WriteString(fmt.Sprintf("bIsRandomizerPalLevelRandom=%s,", formatBool(ws.BIsRandomizerPalLevelRandom)))
	builder.WriteString(fmt.Sprintf("DayTimeSpeedRate=%f,", ws.DayTimeSpeedRate))
	builder.WriteString(fmt.Sprintf("NightTimeSpeedRate=%f,", ws.NightTimeSpeedRate))
	builder.WriteString(fmt.Sprintf("ExpRate=%f,", ws.ExpRate))
	builder.WriteString(fmt.Sprintf("PalCaptureRate=%f,", ws.PalCaptureRate))
	builder.WriteString(fmt.Sprintf("PalSpawnNumRate=%f,", ws.PalSpawnNumRate))
	builder.WriteString(fmt.Sprintf("PalDamageRateAttack=%f,", ws.PalDamageRateAttack))
	builder.WriteString(fmt.Sprintf("PalDamageRateDefense=%f,", ws.PalDamageRateDefense))
	builder.WriteString(fmt.Sprintf("PlayerDamageRateAttack=%f,", ws.PlayerDamageRateAttack))
	builder.WriteString(fmt.Sprintf("PlayerDamageRateDefense=%f,", ws.PlayerDamageRateDefense))
	builder.WriteString(fmt.Sprintf("PlayerStomachDecreaceRate=%f,", ws.PlayerStomachDecreaceRate))
	builder.WriteString(fmt.Sprintf("PlayerStaminaDecreaceRate=%f,", ws.PlayerStaminaDecreaceRate))
	builder.WriteString(fmt.Sprintf("PlayerAutoHPRegeneRate=%f,", ws.PlayerAutoHPRegeneRate))
	builder.WriteString(fmt.Sprintf("PlayerAutoHpRegeneRateInSleep=%f,", ws.PlayerAutoHpRegeneRateInSleep))
	builder.WriteString(fmt.Sprintf("PalStomachDecreaceRate=%f,", ws.PalStomachDecreaceRate))
	builder.WriteString(fmt.Sprintf("PalStaminaDecreaceRate=%f,", ws.PalStaminaDecreaceRate))
	builder.WriteString(fmt.Sprintf("PalAutoHPRegeneRate=%f,", ws.PalAutoHPRegeneRate))
	builder.WriteString(fmt.Sprintf("PalAutoHpRegeneRateInSleep=%f,", ws.PalAutoHpRegeneRateInSleep))
	builder.WriteString(fmt.Sprintf("BuildObjectHpRate=%f,", ws.BuildObjectHpRate))
	builder.WriteString(fmt.Sprintf("BuildObjectDamageRate=%f,", ws.BuildObjectDamageRate))
	builder.WriteString(fmt.Sprintf("BuildObjectDeteriorationDamageRate=%f,", ws.BuildObjectDeteriorationDamageRate))
	builder.WriteString(fmt.Sprintf("CollectionDropRate=%f,", ws.CollectionDropRate))
	builder.WriteString(fmt.Sprintf("CollectionObjectHpRate=%f,", ws.CollectionObjectHpRate))
	builder.WriteString(fmt.Sprintf("CollectionObjectRespawnSpeedRate=%f,", ws.CollectionObjectRespawnSpeedRate))
	builder.WriteString(fmt.Sprintf("EnemyDropItemRate=%f,", ws.EnemyDropItemRate))
	builder.WriteString(fmt.Sprintf("DeathPenalty=%s,", ws.DeathPenalty))
	builder.WriteString(fmt.Sprintf("bEnablePlayerToPlayerDamage=%s,", formatBool(ws.BEnablePlayerToPlayerDamage)))
	builder.WriteString(fmt.Sprintf("bEnableFriendlyFire=%s,", formatBool(ws.BEnableFriendlyFire)))
	builder.WriteString(fmt.Sprintf("bEnableInvaderEnemy=%s,", formatBool(ws.BEnableInvaderEnemy)))
	builder.WriteString(fmt.Sprintf("EnablePredatorBossPal=%s,", formatBool(ws.EnablePredatorBossPal)))
	builder.WriteString(fmt.Sprintf("bActiveUNKO=%s,", formatBool(ws.BActiveUNKO)))
	builder.WriteString(fmt.Sprintf("bEnableAimAssistPad=%s,", formatBool(ws.BEnableAimAssistPad)))
	builder.WriteString(fmt.Sprintf("bEnableAimAssistKeyboard=%s,", formatBool(ws.BEnableAimAssistKeyboard)))
	builder.WriteString(fmt.Sprintf("DropItemMaxNum=%d,", ws.DropItemMaxNum))
	builder.WriteString(fmt.Sprintf("DropItemMaxNum_UNKO=%d,", ws.DropItemMaxNumUNKO))
	builder.WriteString(fmt.Sprintf("BaseCampMaxNum=%d,", ws.BaseCampMaxNum))
	builder.WriteString(fmt.Sprintf("BaseCampMaxNumInGuild=%d,", ws.BaseCampMaxNumInGuild))
	builder.WriteString(fmt.Sprintf("BaseCampWorkerMaxNum=%d,", ws.BaseCampWorkerMaxNum))
	builder.WriteString(fmt.Sprintf("DropItemAliveMaxHours=%f,", ws.DropItemAliveMaxHours))
	builder.WriteString(fmt.Sprintf("bAutoResetGuildNoOnlinePlayers=%s,", formatBool(ws.BAutoResetGuildNoOnlinePlayers)))
	builder.WriteString(fmt.Sprintf("AutoResetGuildTimeNoOnlinePlayers=%f,", ws.AutoResetGuildTimeNoOnlinePlayers))
	builder.WriteString(fmt.Sprintf("GuildPlayerMaxNum=%d,", ws.GuildPlayerMaxNum))
	builder.WriteString(fmt.Sprintf("PalEggDefaultHatchingTime=%f,", ws.PalEggDefaultHatchingTime))
	builder.WriteString(fmt.Sprintf("WorkSpeedRate=%f,", ws.WorkSpeedRate))
	builder.WriteString(fmt.Sprintf("AutoSaveSpan=%f,", ws.AutoSaveSpan))
	builder.WriteString(fmt.Sprintf("CrossplayPlatforms=%s,", formatStringArray(ws.CrossplayPlatforms)))
	builder.WriteString(fmt.Sprintf("LogFormatType=%s,", ws.LogFormatType))
	builder.WriteString(fmt.Sprintf("bIsMultiplay=%s,", formatBool(ws.BIsMultiplay)))
	builder.WriteString(fmt.Sprintf("bIsPvP=%s,", formatBool(ws.BIsPvP)))
	builder.WriteString(fmt.Sprintf("bHardcore=%s,", formatBool(ws.BHardcore)))
	builder.WriteString(fmt.Sprintf("bPalLost=%s,", formatBool(ws.BPalLost)))
	builder.WriteString(fmt.Sprintf("bCharacterRecreateInHardcore=%s,", formatBool(ws.BCharacterRecreateInHardcore)))
	builder.WriteString(fmt.Sprintf("bCanPickupOtherGuildDeathPenaltyDrop=%s,", formatBool(ws.BCanPickupOtherGuildDeathPenaltyDrop)))
	builder.WriteString(fmt.Sprintf("bEnableNonLoginPenalty=%s,", formatBool(ws.BEnableNonLoginPenalty)))
	builder.WriteString(fmt.Sprintf("bEnableFastTravel=%s,", formatBool(ws.BEnableFastTravel)))
	builder.WriteString(fmt.Sprintf("bIsStartLocationSelectByMap=%s,", formatBool(ws.BIsStartLocationSelectByMap)))
	builder.WriteString(fmt.Sprintf("bExistPlayerAfterLogout=%s,", formatBool(ws.BExistPlayerAfterLogout)))
	builder.WriteString(fmt.Sprintf("bEnableDefenseOtherGuildPlayer=%s,", formatBool(ws.BEnableDefenseOtherGuildPlayer)))
	builder.WriteString(fmt.Sprintf("bInvisibleOtherGuildBaseCampAreaFX=%s,", formatBool(ws.BInvisibleOtherGuildBaseCampAreaFX)))
	builder.WriteString(fmt.Sprintf("bBuildAreaLimit=%s,", formatBool(ws.BBuildAreaLimit)))
	builder.WriteString(fmt.Sprintf("ItemWeightRate=%f,", ws.ItemWeightRate))
	builder.WriteString(fmt.Sprintf("bShowPlayerList=%s,", formatBool(ws.BShowPlayerList)))
	builder.WriteString(fmt.Sprintf("CoopPlayerMaxNum=%d,", ws.CoopPlayerMaxNum))
	builder.WriteString(fmt.Sprintf("ServerPlayerMaxNum=%d,", ws.ServerPlayerMaxNum))
	builder.WriteString(fmt.Sprintf("ServerName=\"%s\",", ws.ServerName))
	builder.WriteString(fmt.Sprintf("ServerDescription=\"%s\",", ws.ServerDescription))
	builder.WriteString(fmt.Sprintf("AdminPassword=\"%s\",", ws.AdminPassword))
	builder.WriteString(fmt.Sprintf("ServerPassword=\"%s\",", ws.ServerPassword))
	builder.WriteString(fmt.Sprintf("PublicPort=%d,", ws.PublicPort))
	builder.WriteString(fmt.Sprintf("PublicIP=\"%s\",", ws.PublicIP))
	builder.WriteString(fmt.Sprintf("RCONEnabled=%s,", formatBool(ws.RCONEnabled)))
	builder.WriteString(fmt.Sprintf("RCONPort=%d,", ws.RCONPort))
	builder.WriteString(fmt.Sprintf("RESTAPIEnabled=%s,", formatBool(ws.RESTAPIEnabled)))
	builder.WriteString(fmt.Sprintf("RESTAPIPort=%d,", ws.RESTAPIPort))
	builder.WriteString(fmt.Sprintf("bIsUseBackupSaveData=%s,", formatBool(ws.BIsUseBackupSaveData)))
	builder.WriteString(fmt.Sprintf("Region=\"%s\",", ws.Region))
	builder.WriteString(fmt.Sprintf("bUseAuth=%s,", formatBool(ws.BUseAuth)))
	builder.WriteString(fmt.Sprintf("BanListURL=\"%s\",", ws.BanListURL))
	builder.WriteString(fmt.Sprintf("SupplyDropSpan=%d,", ws.SupplyDropSpan))
	builder.WriteString(fmt.Sprintf("ChatPostLimitPerMinute=%d,", ws.ChatPostLimitPerMinute))
	builder.WriteString(fmt.Sprintf("MaxBuildingLimitNum=%d,", ws.MaxBuildingLimitNum))
	builder.WriteString(fmt.Sprintf("ServerReplicatePawnCullDistance=%f,", ws.ServerReplicatePawnCullDistance))
	builder.WriteString(fmt.Sprintf("bAllowGlobalPalboxExport=%s,", formatBool(ws.BAllowGlobalPalboxExport)))
	builder.WriteString(fmt.Sprintf("bAllowGlobalPalboxImport=%s,", formatBool(ws.BAllowGlobalPalboxImport)))
	builder.WriteString(fmt.Sprintf("EquipmentDurabilityDamageRate=%f,", ws.EquipmentDurabilityDamageRate))
	builder.WriteString(fmt.Sprintf("ItemContainerForceMarkDirtyInterval=%f,", ws.ItemContainerForceMarkDirtyInterval))
	builder.WriteString(fmt.Sprintf("ItemCorruptionMultiplier=%f", ws.ItemCorruptionMultiplier))

	return builder.String()
}

// extractOptionSettingsBlock extracts the content inside the parentheses
// of the "OptionSettings=(...)" section in a Palworld configuration file.
//
// Specifically:
//   - It searches for the substring "OptionSettings=(".
//   - It returns everything between the opening "(" and the last closing ")" in the file.
//   - If "OptionSettings=(" or the closing ")" is not found, it returns an empty string.
//
// Example:
// Input:
// [/Script/Pal.PalGameWorldSettings]
// OptionSettings=(Difficulty=None,ExpRate=1.0,ServerName="Default")
//
// Output:
// Difficulty=None,ExpRate=1.0,ServerName="Default"
func extractOptionSettingsBlock(content string) string {
	start := strings.Index(content, "OptionSettings=(")
	if start == -1 {
		return ""
	}
	start += len("OptionSettings=(")
	end := strings.LastIndex(content, ")")
	if end == -1 || end <= start {
		return ""
	}
	return content[start:end]
}

// parseKeyValuePairs parses a string containing key=value pairs separated by commas.
// It handles nested parentheses and quoted strings correctly.
func parseKeyValuePairs(str string) map[string]string {
	pairs := make(map[string]string)
	var key, value strings.Builder
	inQuotes := false
	inParens := 0
	readingKey := true

	for i := 0; i < len(str); i++ {
		char := str[i]

		switch char {
		case '"':
			inQuotes = !inQuotes
			if !readingKey {
				value.WriteByte(char)
			}
		case '(':
			if !inQuotes {
				inParens++
			}
			if !readingKey {
				value.WriteByte(char)
			}
		case ')':
			if !inQuotes {
				inParens--
			}
			if !readingKey && inParens >= 0 {
				value.WriteByte(char)
			}
		case '=':
			if !inQuotes && inParens == 0 && readingKey {
				readingKey = false
			} else if !readingKey {
				value.WriteByte(char)
			}
		case ',':
			if !inQuotes && inParens == 0 {
				if key.Len() > 0 {
					pairs[key.String()] = strings.Trim(value.String(), `"`)
				}
				key.Reset()
				value.Reset()
				readingKey = true
			} else if !readingKey {
				value.WriteByte(char)
			}
		default:
			if readingKey {
				key.WriteByte(char)
			} else {
				value.WriteByte(char)
			}
		}
	}

	if key.Len() > 0 {
		pairs[key.String()] = value.String()
	}

	return pairs
}

// applyConfigPair applies a single key-value configuration pair to the WorldSetting struct.
// It matches the key against known configuration fields and converts the value to the appropriate type.
func applyConfigPair(key, val string, ws *WorldSetting) {
	switch key {
	case "Difficulty":
		ws.Difficulty = val
	case "RandomizerType":
		ws.RandomizerType = val
	case "RandomizerSeed":
		ws.RandomizerSeed = val
	case "bIsRandomizerPalLevelRandom":
		ws.BIsRandomizerPalLevelRandom = parseBool(val)
	case "DayTimeSpeedRate":
		ws.DayTimeSpeedRate = parseFloat(val)
	case "NightTimeSpeedRate":
		ws.NightTimeSpeedRate = parseFloat(val)
	case "ExpRate":
		ws.ExpRate = parseFloat(val)
	case "PalCaptureRate":
		ws.PalCaptureRate = parseFloat(val)
	case "PalSpawnNumRate":
		ws.PalSpawnNumRate = parseFloat(val)
	case "PalDamageRateAttack":
		ws.PalDamageRateAttack = parseFloat(val)
	case "PalDamageRateDefense":
		ws.PalDamageRateDefense = parseFloat(val)
	case "PlayerDamageRateAttack":
		ws.PlayerDamageRateAttack = parseFloat(val)
	case "PlayerDamageRateDefense":
		ws.PlayerDamageRateDefense = parseFloat(val)
	case "PlayerStomachDecreaceRate":
		ws.PlayerStomachDecreaceRate = parseFloat(val)
	case "PlayerStaminaDecreaceRate":
		ws.PlayerStaminaDecreaceRate = parseFloat(val)
	case "PlayerAutoHPRegeneRate":
		ws.PlayerAutoHPRegeneRate = parseFloat(val)
	case "PlayerAutoHpRegeneRateInSleep":
		ws.PlayerAutoHpRegeneRateInSleep = parseFloat(val)
	case "PalStomachDecreaceRate":
		ws.PalStomachDecreaceRate = parseFloat(val)
	case "PalStaminaDecreaceRate":
		ws.PalStaminaDecreaceRate = parseFloat(val)
	case "PalAutoHPRegeneRate":
		ws.PalAutoHPRegeneRate = parseFloat(val)
	case "PalAutoHpRegeneRateInSleep":
		ws.PalAutoHpRegeneRateInSleep = parseFloat(val)
	case "BuildObjectHpRate":
		ws.BuildObjectHpRate = parseFloat(val)
	case "BuildObjectDamageRate":
		ws.BuildObjectDamageRate = parseFloat(val)
	case "BuildObjectDeteriorationDamageRate":
		ws.BuildObjectDeteriorationDamageRate = parseFloat(val)
	case "CollectionDropRate":
		ws.CollectionDropRate = parseFloat(val)
	case "CollectionObjectHpRate":
		ws.CollectionObjectHpRate = parseFloat(val)
	case "CollectionObjectRespawnSpeedRate":
		ws.CollectionObjectRespawnSpeedRate = parseFloat(val)
	case "EnemyDropItemRate":
		ws.EnemyDropItemRate = parseFloat(val)
	case "DeathPenalty":
		ws.DeathPenalty = val
	case "bEnablePlayerToPlayerDamage":
		ws.BEnablePlayerToPlayerDamage = parseBool(val)
	case "bEnableFriendlyFire":
		ws.BEnableFriendlyFire = parseBool(val)
	case "bEnableInvaderEnemy":
		ws.BEnableInvaderEnemy = parseBool(val)
	case "EnablePredatorBossPal":
		ws.EnablePredatorBossPal = parseBool(val)
	case "bActiveUNKO":
		ws.BActiveUNKO = parseBool(val)
	case "bEnableAimAssistPad":
		ws.BEnableAimAssistPad = parseBool(val)
	case "bEnableAimAssistKeyboard":
		ws.BEnableAimAssistKeyboard = parseBool(val)
	case "DropItemMaxNum":
		ws.DropItemMaxNum = parseInt(val)
	case "DropItemMaxNum_UNKO":
		ws.DropItemMaxNumUNKO = parseInt(val)
	case "BaseCampMaxNum":
		ws.BaseCampMaxNum = parseInt(val)
	case "BaseCampMaxNumInGuild":
		ws.BaseCampMaxNumInGuild = parseInt(val)
	case "BaseCampWorkerMaxNum":
		ws.BaseCampWorkerMaxNum = parseInt(val)
	case "DropItemAliveMaxHours":
		ws.DropItemAliveMaxHours = parseFloat(val)
	case "bAutoResetGuildNoOnlinePlayers":
		ws.BAutoResetGuildNoOnlinePlayers = parseBool(val)
	case "AutoResetGuildTimeNoOnlinePlayers":
		ws.AutoResetGuildTimeNoOnlinePlayers = parseFloat(val)
	case "GuildPlayerMaxNum":
		ws.GuildPlayerMaxNum = parseInt(val)
	case "PalEggDefaultHatchingTime":
		ws.PalEggDefaultHatchingTime = parseFloat(val)
	case "WorkSpeedRate":
		ws.WorkSpeedRate = parseFloat(val)
	case "AutoSaveSpan":
		ws.AutoSaveSpan = parseFloat(val)
	case "CrossplayPlatforms":
		ws.CrossplayPlatforms = parseList(val)
	case "LogFormatType":
		ws.LogFormatType = val
	case "bIsMultiplay":
		ws.BIsMultiplay = parseBool(val)
	case "bIsPvP":
		ws.BIsPvP = parseBool(val)
	case "bHardcore":
		ws.BHardcore = parseBool(val)
	case "bPalLost":
		ws.BPalLost = parseBool(val)
	case "bCharacterRecreateInHardcore":
		ws.BCharacterRecreateInHardcore = parseBool(val)
	case "bCanPickupOtherGuildDeathPenaltyDrop":
		ws.BCanPickupOtherGuildDeathPenaltyDrop = parseBool(val)
	case "bEnableNonLoginPenalty":
		ws.BEnableNonLoginPenalty = parseBool(val)
	case "bEnableFastTravel":
		ws.BEnableFastTravel = parseBool(val)
	case "bIsStartLocationSelectByMap":
		ws.BIsStartLocationSelectByMap = parseBool(val)
	case "bExistPlayerAfterLogout":
		ws.BExistPlayerAfterLogout = parseBool(val)
	case "bEnableDefenseOtherGuildPlayer":
		ws.BEnableDefenseOtherGuildPlayer = parseBool(val)
	case "bInvisibleOtherGuildBaseCampAreaFX":
		ws.BInvisibleOtherGuildBaseCampAreaFX = parseBool(val)
	case "bBuildAreaLimit":
		ws.BBuildAreaLimit = parseBool(val)
	case "ItemWeightRate":
		ws.ItemWeightRate = parseFloat(val)
	case "bShowPlayerList":
		ws.BShowPlayerList = parseBool(val)
	case "CoopPlayerMaxNum":
		ws.CoopPlayerMaxNum = parseInt(val)
	case "ServerPlayerMaxNum":
		ws.ServerPlayerMaxNum = parseInt(val)
	case "ServerName":
		ws.ServerName = val
	case "ServerDescription":
		ws.ServerDescription = val
	case "AdminPassword":
		ws.AdminPassword = val
	case "ServerPassword":
		ws.ServerPassword = val
	case "PublicPort":
		ws.PublicPort = parseInt(val)
	case "PublicIP":
		ws.PublicIP = val
	case "RCONEnabled":
		ws.RCONEnabled = parseBool(val)
	case "RCONPort":
		ws.RCONPort = parseInt(val)
	case "RESTAPIEnabled":
		ws.RESTAPIEnabled = parseBool(val)
	case "RESTAPIPort":
		ws.RESTAPIPort = parseInt(val)
	case "bIsUseBackupSaveData":
		ws.BIsUseBackupSaveData = parseBool(val)
	case "Region":
		ws.Region = val
	case "bUseAuth":
		ws.BUseAuth = parseBool(val)
	case "BanListURL":
		ws.BanListURL = val
	case "SupplyDropSpan":
		ws.SupplyDropSpan = parseInt(val)
	case "ChatPostLimitPerMinute":
		ws.ChatPostLimitPerMinute = parseInt(val)
	case "MaxBuildingLimitNum":
		ws.MaxBuildingLimitNum = parseInt(val)
	case "ServerReplicatePawnCullDistance":
		ws.ServerReplicatePawnCullDistance = parseFloat(val)
	case "bAllowGlobalPalboxExport":
		ws.BAllowGlobalPalboxExport = parseBool(val)
	case "bAllowGlobalPalboxImport":
		ws.BAllowGlobalPalboxImport = parseBool(val)
	case "EquipmentDurabilityDamageRate":
		ws.EquipmentDurabilityDamageRate = parseFloat(val)
	case "ItemContainerForceMarkDirtyInterval":
		ws.ItemContainerForceMarkDirtyInterval = parseFloat(val)
	case "ItemCorruptionMultiplier":
		ws.ItemCorruptionMultiplier = parseFloat(val)
	}
}

// parseBool safely converts string to bool (supports True/False, true/false)
func parseBool(s string) bool {
	s = strings.ToLower(s)
	return s == "true" || s == "1" || s == "yes"
}

// parseFloat safely converts string to float64
func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

// parseInt safely converts string to int
func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// parseList converts "(Steam,Xbox,PS5,Mac)" → []string{"Steam","Xbox","PS5","Mac"}
func parseList(s string) []string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "()")

	if s == "" {
		return []string{}
	}

	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			result = append(result, part)
		}
	}

	return result
}

// formatBool converts a boolean to Palworld config format (True/False)
func formatBool(b bool) string {
	if b {
		return "True"
	}
	return "False"
}

// formatStringArray converts a string slice to Palworld config array format
// Example: []string{"Steam", "Xbox"} -> "(Steam,Xbox)"
func formatStringArray(arr []string) string {
	if len(arr) == 0 {
		return "()"
	}
	return "(" + strings.Join(arr, ",") + ")"
}
