package source

import (
	data "../data"
	source "../source"
	"gotest.tools/assert"
	"testing"
)

func TestParseLookupPath(t *testing.T) {
	var lookupPath LookupPath
	var data [][]string = [][]string{
		[]string{
			"v1.1.0/company/team_01/db/aws/postgres/project_01/subproject_01/stg/creds/readonly",
			"v1.0/company/team_02/kv/project_01/stg/environment",
		},
		LookupPath{
			data.FullVersion(1, 1, 0),
			[]string{"company", "team_01"},
			source.Db,
			[]string{"aws", "postgres", "project_01", "subproject_01"},
		},
	}
}
