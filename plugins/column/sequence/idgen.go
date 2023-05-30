package sequence

import "github.com/GUAIK-ORG/go-snowflake/snowflake"

func snowflakeID() int64 {
	s, err := snowflake.NewSnowflake(int64(0), int64(0))
	if err != nil {
		return 0
	}

	return s.NextVal()
}
