package gcache

// 6                   42
// +------+------------------------------------------+
// | hint |          entity id                       |
// +------+------------------------------------------+

func parseMergeID(mergeID int64) (hintID, id int64) {
	hintID = mergeID >> 42
	// 63     = 0b111111
	// 63<<42 = 0b111111000000000000000000000000000000000000000000
	id = (^(int64(63) << 42)) & mergeID
	return
}

func generateMergeID(hintID, id int64) int64 {
	return (hintID << 42) | id
}
