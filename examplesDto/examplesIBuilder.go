package examplesDto

// TestIBuilderErrPref - Implements the IBuilderErrorPrefix
// interface declared in the 'errpref' software package
//    "github.com/MikeAustin71/errpref"
//
type TestIBuilderErrPref struct {
	errPrefInfo [][2]string
}

func (tIBuilder *TestIBuilderErrPref) GetEPrefStrings() [][2]string {

	if tIBuilder.errPrefInfo == nil {
		return nil
	}

	lenTwoDArray := len(tIBuilder.errPrefInfo)

	result := make([][2]string, lenTwoDArray)

	if lenTwoDArray == 0 {
		return result
	}

	copy(result, tIBuilder.errPrefInfo)

	return result
}

func (tIBuilder *TestIBuilderErrPref) SetEPrefStrings(twoDStrArray [][2]string) {

	lenTwoDArray := len(twoDStrArray)

	if lenTwoDArray == 0 {
		return
	}

	tIBuilder.errPrefInfo = make([][2]string, lenTwoDArray)

	copy(tIBuilder.errPrefInfo, twoDStrArray)

}

func (tIBuilder *TestIBuilderErrPref) String() string {

	str := ""

	lenTwoDArray := len(tIBuilder.errPrefInfo)

	if lenTwoDArray == 0 {
		return str
	}

	for i := 0; i < lenTwoDArray; i++ {

		if len(tIBuilder.errPrefInfo[i][0]) == 0 {
			continue
		}

		str += tIBuilder.errPrefInfo[i][0]

		if len(tIBuilder.errPrefInfo[i][1]) == 0 {
			str += "\n"
		} else {
			str += " : " + tIBuilder.errPrefInfo[i][1] +
				"\n"
		}
	}

	return str
}
