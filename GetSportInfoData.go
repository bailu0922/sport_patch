package sport_patch

type GetSportInfoData struct {
	Policys        []Policy      `json:"policys"`
	SportArea      int64         `json:"sport_area"`
	Timeout        int64         `json:"timeout"`
	Dynamicid      bool          `json:"dynamicid"`
	Enable         bool          `json:"enable"`
	Lap            Lap           `json:"lap"`
	ID             string        `json:"id"`
	Value          string        `json:"value"`
	CD             int64         `json:"cd"`
	CategoryCode   string        `json:"category_code"`
	Custom         Custom        `json:"custom"`
	Illegal        []interface{} `json:"illegal"`
	ParentID       string        `json:"parentId"`
	ScoreType      string        `json:"score_type"`
	Mock           string        `json:"mock"`
	ScoreUnit      string        `json:"score_unit"`
	CheckinTimeout int64         `json:"checkin_timeout"`
	Transpose      int64         `json:"transpose"`
	Desc           string        `json:"desc"`
}

type Custom struct {
	DelayGo     int64    `json:"delay_go"`
	DelayPower  string   `json:"delay_power"`
	LimitSports []string `json:"limit_sports"`
}

type Lap struct {
	OrchestrationsShuttleDoneIndex1_Patchtime int64   `json:"orchestrations_shuttle_done_index_1_patchtime"`
	Freezeloop                                int64   `json:"freezeloop"`
	Size                                      int64   `json:"size"`
	Freezestop                                int64   `json:"freezestop"`
	OrchestrationsShuttleDoneIndex2_Patchtime int64   `json:"orchestrations_shuttle_done_index_2_patchtime"`
	Loop                                      []int64 `json:"loop"`
	ClassN                                    int64   `json:"class_n"`
	FacesearchLimit                           int64   `json:"facesearch_limit"`
	Freezestart                               int64   `json:"freezestart"`
}

type Policy struct {
	Code  string  `json:"code"`
	Level []int64 `json:"level"`
	Desc  string  `json:"desc"`
}
