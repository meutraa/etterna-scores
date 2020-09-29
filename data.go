package main

import "time"

type Scores []Score

type BestScore struct {
	WifeScore      float32
	SSRNormPercent float32
	EtternaValid   int
	Pack           string
	Date           time.Time
	Steps          string
	Grade          string
	Song           string
	Rate           float32
	Overall        float32
	Stream         float32
	Jumpstream     float32
	Handstream     float32
	Stamina        float32
	JackSpeed      float32
	Chordjack      float32
	Technical      float32
}

type BestScores []BestScore
type BestScoresByTime []BestScore

type Score struct {
	// Text            string  `xml:",chardata"`
	Key string `xml:"Key,attr"`
	// SSRCalcVersion  string  `xml:"SSRCalcVersion"`
	Grade     string  `xml:"Grade"`
	WifeScore float32 `xml:"WifeScore"`
	// WifePoints      string  `xml:"WifePoints"`
	SSRNormPercent float32 `xml:"SSRNormPercent"`
	// JudgeScale      string  `xml:"JudgeScale"`
	// NoChordCohesion string  `xml:"NoChordCohesion"`
	EtternaValid int `xml:"EtternaValid"`
	// SurviveSeconds  float64 `xml:"SurviveSeconds"`
	// MaxCombo        int     `xml:"MaxCombo"`
	// Modifiers       string  `xml:"Modifiers"`
	// MachineGuid     string  `xml:"MachineGuid"`
	DateTime string `xml:"DateTime"`
	// TopScore        string  `xml:"TopScore"`
	// Servs           struct {
	// Text   string `xml:",chardata"`
	// Server string `xml:"server"`
	// } `xml:"Servs"`
	// TapNoteScores struct {
	// Text      string `xml:",chardata"`
	// HitMine   string `xml:"HitMine"`
	// AvoidMine string `xml:"AvoidMine"`
	// Miss      string `xml:"Miss"`
	// W5        string `xml:"W5"`
	// 	W4        string `xml:"W4"`
	// 	W3        string `xml:"W3"`
	// 	W2        string `xml:"W2"`
	// 	W1        string `xml:"W1"`
	// } `xml:"TapNoteScores"`
	// HoldNoteScores struct {
	// 	Text       string `xml:",chardata"`
	// 	LetGo      string `xml:"LetGo"`
	// 	Held       string `xml:"Held"`
	// 	MissedHold string `xml:"MissedHold"`
	// } `xml:"HoldNoteScores"`
	SkillsetSSRs struct {
		// Text       string  `xml:",chardata"`
		Overall    float32 `xml:"Overall"`
		Stream     float32 `xml:"Stream"`
		Jumpstream float32 `xml:"Jumpstream"`
		Handstream float32 `xml:"Handstream"`
		Stamina    float32 `xml:"Stamina"`
		JackSpeed  float32 `xml:"JackSpeed"`
		Chordjack  float32 `xml:"Chordjack"`
		Technical  float32 `xml:"Technical"`
	} `xml:"SkillsetSSRs"`
	// ValidationKeys struct {
	// 	Text    string `xml:",chardata"`
	// 	Brittle string `xml:"Brittle"`
	// 	Weak    string `xml:"Weak"`
	// } `xml:"ValidationKeys"`
}

type Stats struct {
	// XMLName     xml.Name `xml:"Stats"`
	// Text        string   `xml:",chardata"`
	// GeneralData struct {
	// 	Text           string `xml:",chardata"`
	// 	DisplayName    string `xml:"DisplayName"`
	// 	CharacterID    string `xml:"CharacterID"`
	// 	Guid           string `xml:"Guid"`
	// 	SortOrder      string `xml:"SortOrder"`
	// 	LastDifficulty string `xml:"LastDifficulty"`
	// 	LastStepsType  string `xml:"LastStepsType"`
	// 	Song           struct {
	// 		Text string `xml:",chardata"`
	// 		Dir  string `xml:"Dir,attr"`
	// 	} `xml:"Song"`
	// 	TotalSessions         string `xml:"TotalSessions"`
	// 	TotalSessionSeconds   string `xml:"TotalSessionSeconds"`
	// 	TotalGameplaySeconds  string `xml:"TotalGameplaySeconds"`
	// 	LastPlayedMachineGuid string `xml:"LastPlayedMachineGuid"`
	// 	LastPlayedDate        string `xml:"LastPlayedDate"`
	// 	TotalDancePoints      string `xml:"TotalDancePoints"`
	// 	NumToasties           string `xml:"NumToasties"`
	// 	TotalTapsAndHolds     string `xml:"TotalTapsAndHolds"`
	// 	TotalJumps            string `xml:"TotalJumps"`
	// 	TotalHolds            string `xml:"TotalHolds"`
	// 	TotalRolls            string `xml:"TotalRolls"`
	// 	TotalMines            string `xml:"TotalMines"`
	// 	TotalHands            string `xml:"TotalHands"`
	// 	TotalLifts            string `xml:"TotalLifts"`
	// 	PlayerRating          string `xml:"PlayerRating"`
	// 	DefaultModifiers      struct {
	// 		Text  string `xml:",chardata"`
	// 		Dance string `xml:"dance"`
	// 	} `xml:"DefaultModifiers"`
	// 	PlayerSkillsets struct {
	// 		Text       string `xml:",chardata"`
	// 		Overall    string `xml:"Overall"`
	// 		Stream     string `xml:"Stream"`
	// 		Jumpstream string `xml:"Jumpstream"`
	// 		Handstream string `xml:"Handstream"`
	// 		Stamina    string `xml:"Stamina"`
	// 		JackSpeed  string `xml:"JackSpeed"`
	// 		Chordjack  string `xml:"Chordjack"`
	// 		Technical  string `xml:"Technical"`
	// 	} `xml:"PlayerSkillsets"`
	// 	NumTotalSongsPlayed       string `xml:"NumTotalSongsPlayed"`
	// 	NumStagesPassedByPlayMode string `xml:"NumStagesPassedByPlayMode"`
	// 	UserTable                 string `xml:"UserTable"`
	// } `xml:"GeneralData"`
	// Playlists struct {
	// 	Text     string `xml:",chardata"`
	// 	Playlist struct {
	// 		Text      string `xml:",chardata"`
	// 		Name      string `xml:"Name,attr"`
	// 		Chartlist struct {
	// 			Text  string `xml:",chardata"`
	// 			Chart []struct {
	// 				Text  string `xml:",chardata"`
	// 				Key   string `xml:"Key,attr"`
	// 				Pack  string `xml:"Pack,attr"`
	// 				Rate  string `xml:"Rate,attr"`
	// 				Song  string `xml:"Song,attr"`
	// 				Steps string `xml:"Steps,attr"`
	// 			} `xml:"Chart"`
	// 		} `xml:"Chartlist"`
	// 	} `xml:"Playlist"`
	// } `xml:"Playlists"`
	// ScoreGoals struct {
	// 	Text          string `xml:",chardata"`
	// 	GoalsForChart []struct {
	// 		Text      string `xml:",chardata"`
	// 		Key       string `xml:"Key,attr"`
	// 		ScoreGoal struct {
	// 			Text         string `xml:",chardata"`
	// 			Rate         string `xml:"Rate"`
	// 			Percent      string `xml:"Percent"`
	// 			Priority     string `xml:"Priority"`
	// 			Achieved     string `xml:"Achieved"`
	// 			TimeAssigned string `xml:"TimeAssigned"`
	// 			TimeAchieved string `xml:"TimeAchieved"`
	// 			ScoreKey     string `xml:"ScoreKey"`
	// 			Comment      string `xml:"Comment"`
	// 		} `xml:"ScoreGoal"`
	// 	} `xml:"GoalsForChart"`
	// } `xml:"ScoreGoals"`
	PlayerScores struct {
		// Text  string `xml:",chardata"`
		Chart []struct {
			// Text     string `xml:",chardata"`
			// Key      string `xml:"Key,attr"`
			Pack     string `xml:"Pack,attr"`
			Song     string `xml:"Song,attr"`
			Steps    string `xml:"Steps,attr"`
			ScoresAt []struct {
				// Text      string  `xml:",chardata"`
				// BestGrade string  `xml:"BestGrade,attr"`
				PBKey string  `xml:"PBKey,attr"`
				Rate  float32 `xml:"Rate,attr"`
				// NoccPBKey string  `xml:"noccPBKey,attr"`
				Scores []Score `xml:"Score"`
			} `xml:"ScoresAt"`
		} `xml:"Chart"`
	} `xml:"PlayerScores"`
}
