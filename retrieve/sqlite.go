package retrieve

import (
	"job-funnel/load"
	"job-funnel/utils"
)

func RetrieveDataFromSqliteAll() []utils.JobPost {
	var retrievedData []utils.JobPost
	load.DB.Table(*load.TableName).Find(&retrievedData)
	return retrievedData
}
