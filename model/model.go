package model

type credentials struct {
	PgUser     string `json:"pgUser"`
	PgPassword string `json:"pgPassword"`
	PgHost     string `json:"pgHost"`
	PgPort     string `json:"pgPort"`
	PgDatabase string `json:"pgDatabase"`
}

type CheckPostgresqlLatestVersionModel struct {
	Current      bool   `json:"current"`
	EolDate      string `json:"eolDate"`
	FirstRelDate string `json:"firstRelDate"`
	LatestMinor  string `json:"latestMinor"`
	Major        string `json:"major"`
	RelDate      string `json:"relDate"`
	Supported    bool   `json:"supported"`
}
