package palsvm

import (
	"reflect"
	"testing"
)

func TestNewWorldSettingFromFileContentConfig(t *testing.T) {
	fileContent := "[/Script/Pal.PalGameWorldSettings]\nOptionSettings=(Difficulty=None,RandomizerType=None,RandomizerSeed=\"\",bIsRandomizerPalLevelRandom=False,DayTimeSpeedRate=1.000000,NightTimeSpeedRate=1.000000,ExpRate=1.000000,PalCaptureRate=1.000000,PalSpawnNumRate=1.000000,PalDamageRateAttack=1.000000,PalDamageRateDefense=1.000000,PlayerDamageRateAttack=1.000000,PlayerDamageRateDefense=1.000000,PlayerStomachDecreaceRate=1.000000,PlayerStaminaDecreaceRate=1.000000,PlayerAutoHPRegeneRate=1.000000,PlayerAutoHpRegeneRateInSleep=1.000000,PalStomachDecreaceRate=1.000000,PalStaminaDecreaceRate=1.000000,PalAutoHPRegeneRate=1.000000,PalAutoHpRegeneRateInSleep=1.000000,BuildObjectHpRate=1.000000,BuildObjectDamageRate=1.000000,BuildObjectDeteriorationDamageRate=1.000000,CollectionDropRate=1.000000,CollectionObjectHpRate=1.000000,CollectionObjectRespawnSpeedRate=1.000000,EnemyDropItemRate=1.000000,DeathPenalty=All,bEnablePlayerToPlayerDamage=False,bEnableFriendlyFire=False,bEnableInvaderEnemy=True,EnablePredatorBossPal=True,bActiveUNKO=False,bEnableAimAssistPad=True,bEnableAimAssistKeyboard=False,DropItemMaxNum=3000,DropItemMaxNum_UNKO=100,BaseCampMaxNum=128,BaseCampMaxNumInGuild=3,BaseCampWorkerMaxNum=15,DropItemAliveMaxHours=1.000000,bAutoResetGuildNoOnlinePlayers=False,AutoResetGuildTimeNoOnlinePlayers=72.000000,GuildPlayerMaxNum=20,PalEggDefaultHatchingTime=72.000000,WorkSpeedRate=1.000000,AutoSaveSpan=30.000000,CrossplayPlatforms=(Steam,Xbox,PS5,Mac),LogFormatType=Text,bIsMultiplay=False,bIsPvP=False,bHardcore=False,bPalLost=False,bCharacterRecreateInHardcore=False,bCanPickupOtherGuildDeathPenaltyDrop=False,bEnableNonLoginPenalty=True,bEnableFastTravel=True,bIsStartLocationSelectByMap=True,bExistPlayerAfterLogout=False,bEnableDefenseOtherGuildPlayer=False,bInvisibleOtherGuildBaseCampAreaFX=False,bBuildAreaLimit=False,ItemWeightRate=1.000000,bShowPlayerList=False,CoopPlayerMaxNum=4,ServerPlayerMaxNum=32,ServerName=\"Default Palworld Server\",ServerDescription=\"\",AdminPassword=\"\",ServerPassword=\"\",PublicPort=8211,PublicIP=\"\",RCONEnabled=False,RCONPort=25575,RESTAPIEnabled=False,RESTAPIPort=8212,bIsUseBackupSaveData=True,Region=\"\",bUseAuth=True,BanListURL=\"https://api.palworldgame.com/api/banlist.txt\",SupplyDropSpan=180,ChatPostLimitPerMinute=10,MaxBuildingLimitNum=0,ServerReplicatePawnCullDistance=15000.000000,bAllowGlobalPalboxExport=True,bAllowGlobalPalboxImport=False,EquipmentDurabilityDamageRate=1.000000,ItemContainerForceMarkDirtyInterval=1.000000,ItemCorruptionMultiplier=1.000000)"
	ws := NewWorldSettingFromFileContentConfig(fileContent)

	if ws == nil {
		t.Fatalf("expected non-nil WorldSetting, got nil")
	}

	tests := []struct {
		name string
		got  any
		want any
	}{
		// String field
		{"Difficulty", ws.Difficulty, "None"},
		{"RandomizerType", ws.RandomizerType, "None"},
		{"RandomizerSeed", ws.RandomizerSeed, ""},
		{"DeathPenalty", ws.DeathPenalty, "All"},
		{"LogFormatType", ws.LogFormatType, "Text"},
		{"ServerName", ws.ServerName, "Default Palworld Server"},
		{"ServerDescription", ws.ServerDescription, ""},
		{"AdminPassword", ws.AdminPassword, ""},
		{"ServerPassword", ws.ServerPassword, ""},
		{"PublicIP", ws.PublicIP, ""},
		{"Region", ws.Region, ""},
		{"BanListURL", ws.BanListURL, "https://api.palworldgame.com/api/banlist.txt"},

		// Bool fields
		{"BIsRandomizerPalLevelRandom", ws.BIsRandomizerPalLevelRandom, false},
		{"BEnablePlayerToPlayerDamage", ws.BEnablePlayerToPlayerDamage, false},
		{"BEnableFriendlyFire", ws.BEnableFriendlyFire, false},
		{"BEnableInvaderEnemy", ws.BEnableInvaderEnemy, true},
		{"EnablePredatorBossPal", ws.EnablePredatorBossPal, true},
		{"BActiveUNKO", ws.BActiveUNKO, false},
		{"BEnableAimAssistPad", ws.BEnableAimAssistPad, true},
		{"BEnableAimAssistKeyboard", ws.BEnableAimAssistKeyboard, false},
		{"BAutoResetGuildNoOnlinePlayers", ws.BAutoResetGuildNoOnlinePlayers, false},
		{"BIsMultiplay", ws.BIsMultiplay, false},
		{"BIsPvP", ws.BIsPvP, false},
		{"BHardcore", ws.BHardcore, false},
		{"BPalLost", ws.BPalLost, false},
		{"BCharacterRecreateInHardcore", ws.BCharacterRecreateInHardcore, false},
		{"BCanPickupOtherGuildDeathPenaltyDrop", ws.BCanPickupOtherGuildDeathPenaltyDrop, false},
		{"BEnableNonLoginPenalty", ws.BEnableNonLoginPenalty, true},
		{"BEnableFastTravel", ws.BEnableFastTravel, true},
		{"BIsStartLocationSelectByMap", ws.BIsStartLocationSelectByMap, true},
		{"BExistPlayerAfterLogout", ws.BExistPlayerAfterLogout, false},
		{"BEnableDefenseOtherGuildPlayer", ws.BEnableDefenseOtherGuildPlayer, false},
		{"BInvisibleOtherGuildBaseCampAreaFX", ws.BInvisibleOtherGuildBaseCampAreaFX, false},
		{"BBuildAreaLimit", ws.BBuildAreaLimit, false},
		{"BShowPlayerList", ws.BShowPlayerList, false},
		{"RCONEnabled", ws.RCONEnabled, false},
		{"RESTAPIEnabled", ws.RESTAPIEnabled, false},
		{"BIsUseBackupSaveData", ws.BIsUseBackupSaveData, true},
		{"BUseAuth", ws.BUseAuth, true},
		{"BAllowGlobalPalboxExport", ws.BAllowGlobalPalboxExport, true},
		{"BAllowGlobalPalboxImport", ws.BAllowGlobalPalboxImport, false},

		// Int fields
		{"DropItemMaxNum", ws.DropItemMaxNum, 3000},
		{"DropItemMaxNumUNKO", ws.DropItemMaxNumUNKO, 100},
		{"BaseCampMaxNum", ws.BaseCampMaxNum, 128},
		{"BaseCampMaxNumInGuild", ws.BaseCampMaxNumInGuild, 3},
		{"BaseCampWorkerMaxNum", ws.BaseCampWorkerMaxNum, 15},
		{"GuildPlayerMaxNum", ws.GuildPlayerMaxNum, 20},
		{"CoopPlayerMaxNum", ws.CoopPlayerMaxNum, 4},
		{"ServerPlayerMaxNum", ws.ServerPlayerMaxNum, 32},
		{"PublicPort", ws.PublicPort, 8211},
		{"RCONPort", ws.RCONPort, 25575},
		{"RESTAPIPort", ws.RESTAPIPort, 8212},
		{"SupplyDropSpan", ws.SupplyDropSpan, 180},
		{"ChatPostLimitPerMinute", ws.ChatPostLimitPerMinute, 10},
		{"MaxBuildingLimitNum", ws.MaxBuildingLimitNum, 0},

		// Float64 fields
		{"DayTimeSpeedRate", ws.DayTimeSpeedRate, 1.0},
		{"NightTimeSpeedRate", ws.NightTimeSpeedRate, 1.0},
		{"ExpRate", ws.ExpRate, 1.0},
		{"PalCaptureRate", ws.PalCaptureRate, 1.0},
		{"PalSpawnNumRate", ws.PalSpawnNumRate, 1.0},
		{"PalDamageRateAttack", ws.PalDamageRateAttack, 1.0},
		{"PalDamageRateDefense", ws.PalDamageRateDefense, 1.0},
		{"PlayerDamageRateAttack", ws.PlayerDamageRateAttack, 1.0},
		{"PlayerDamageRateDefense", ws.PlayerDamageRateDefense, 1.0},
		{"PlayerStomachDecreaceRate", ws.PlayerStomachDecreaceRate, 1.0},
		{"PlayerStaminaDecreaceRate", ws.PlayerStaminaDecreaceRate, 1.0},
		{"PlayerAutoHPRegeneRate", ws.PlayerAutoHPRegeneRate, 1.0},
		{"PlayerAutoHpRegeneRateInSleep", ws.PlayerAutoHpRegeneRateInSleep, 1.0},
		{"PalStomachDecreaceRate", ws.PalStomachDecreaceRate, 1.0},
		{"PalStaminaDecreaceRate", ws.PalStaminaDecreaceRate, 1.0},
		{"PalAutoHPRegeneRate", ws.PalAutoHPRegeneRate, 1.0},
		{"PalAutoHpRegeneRateInSleep", ws.PalAutoHpRegeneRateInSleep, 1.0},
		{"BuildObjectHpRate", ws.BuildObjectHpRate, 1.0},
		{"BuildObjectDamageRate", ws.BuildObjectDamageRate, 1.0},
		{"BuildObjectDeteriorationDamageRate", ws.BuildObjectDeteriorationDamageRate, 1.0},
		{"CollectionDropRate", ws.CollectionDropRate, 1.0},
		{"CollectionObjectHpRate", ws.CollectionObjectHpRate, 1.0},
		{"CollectionObjectRespawnSpeedRate", ws.CollectionObjectRespawnSpeedRate, 1.0},
		{"EnemyDropItemRate", ws.EnemyDropItemRate, 1.0},
		{"DropItemAliveMaxHours", ws.DropItemAliveMaxHours, 1.0},
		{"AutoResetGuildTimeNoOnlinePlayers", ws.AutoResetGuildTimeNoOnlinePlayers, 72.0},
		{"PalEggDefaultHatchingTime", ws.PalEggDefaultHatchingTime, 72.0},
		{"WorkSpeedRate", ws.WorkSpeedRate, 1.0},
		{"AutoSaveSpan", ws.AutoSaveSpan, 30.0},
		{"ItemWeightRate", ws.ItemWeightRate, 1.0},
		{"ServerReplicatePawnCullDistance", ws.ServerReplicatePawnCullDistance, 15000.0},
		{"EquipmentDurabilityDamageRate", ws.EquipmentDurabilityDamageRate, 1.0},
		{"ItemContainerForceMarkDirtyInterval", ws.ItemContainerForceMarkDirtyInterval, 1.0},
		{"ItemCorruptionMultiplier", ws.ItemCorruptionMultiplier, 1.0},

		// Array fields
		{"CrossplayPlatforms", ws.CrossplayPlatforms, []string{"Steam", "Xbox", "PS5", "Mac"}},
	}

	for _, tt := range tests {
		if !reflect.DeepEqual(tt.got, tt.want) {
			t.Errorf("%s: got %v, want %v", tt.name, tt.got, tt.want)
		}
	}
}

func TestNewWorldSettingString(t *testing.T) {
	fileContent := "[/Script/Pal.PalGameWorldSettings]\nOptionSettings=(Difficulty=None,RandomizerType=None,RandomizerSeed=\"\",bIsRandomizerPalLevelRandom=False,DayTimeSpeedRate=1.000000,NightTimeSpeedRate=1.000000,ExpRate=1.000000,PalCaptureRate=1.000000,PalSpawnNumRate=1.000000,PalDamageRateAttack=1.000000,PalDamageRateDefense=1.000000,PlayerDamageRateAttack=1.000000,PlayerDamageRateDefense=1.000000,PlayerStomachDecreaceRate=1.000000,PlayerStaminaDecreaceRate=1.000000,PlayerAutoHPRegeneRate=1.000000,PlayerAutoHpRegeneRateInSleep=1.000000,PalStomachDecreaceRate=1.000000,PalStaminaDecreaceRate=1.000000,PalAutoHPRegeneRate=1.000000,PalAutoHpRegeneRateInSleep=1.000000,BuildObjectHpRate=1.000000,BuildObjectDamageRate=1.000000,BuildObjectDeteriorationDamageRate=1.000000,CollectionDropRate=1.000000,CollectionObjectHpRate=1.000000,CollectionObjectRespawnSpeedRate=1.000000,EnemyDropItemRate=1.000000,DeathPenalty=All,bEnablePlayerToPlayerDamage=False,bEnableFriendlyFire=False,bEnableInvaderEnemy=True,EnablePredatorBossPal=True,bActiveUNKO=False,bEnableAimAssistPad=True,bEnableAimAssistKeyboard=False,DropItemMaxNum=3000,DropItemMaxNum_UNKO=100,BaseCampMaxNum=128,BaseCampMaxNumInGuild=3,BaseCampWorkerMaxNum=15,DropItemAliveMaxHours=1.000000,bAutoResetGuildNoOnlinePlayers=False,AutoResetGuildTimeNoOnlinePlayers=72.000000,GuildPlayerMaxNum=20,PalEggDefaultHatchingTime=72.000000,WorkSpeedRate=1.000000,AutoSaveSpan=30.000000,CrossplayPlatforms=(Steam,Xbox,PS5,Mac),LogFormatType=Text,bIsMultiplay=False,bIsPvP=False,bHardcore=False,bPalLost=False,bCharacterRecreateInHardcore=False,bCanPickupOtherGuildDeathPenaltyDrop=False,bEnableNonLoginPenalty=True,bEnableFastTravel=True,bIsStartLocationSelectByMap=True,bExistPlayerAfterLogout=False,bEnableDefenseOtherGuildPlayer=False,bInvisibleOtherGuildBaseCampAreaFX=False,bBuildAreaLimit=False,ItemWeightRate=1.000000,bShowPlayerList=False,CoopPlayerMaxNum=4,ServerPlayerMaxNum=32,ServerName=\"Default Palworld Server\",ServerDescription=\"\",AdminPassword=\"\",ServerPassword=\"\",PublicPort=8211,PublicIP=\"\",RCONEnabled=False,RCONPort=25575,RESTAPIEnabled=False,RESTAPIPort=8212,bIsUseBackupSaveData=True,Region=\"\",bUseAuth=True,BanListURL=\"https://api.palworldgame.com/api/banlist.txt\",SupplyDropSpan=180,ChatPostLimitPerMinute=10,MaxBuildingLimitNum=0,ServerReplicatePawnCullDistance=15000.000000,bAllowGlobalPalboxExport=True,bAllowGlobalPalboxImport=False,EquipmentDurabilityDamageRate=1.000000,ItemContainerForceMarkDirtyInterval=1.000000,ItemCorruptionMultiplier=1.000000)"
	ws := NewWorldSettingFromFileContentConfig(fileContent)
	wsString := ws.String()
	if fileContent != wsString {
		t.Errorf("\nExpected: %s\nGot: %s", fileContent, wsString)
	}
}
