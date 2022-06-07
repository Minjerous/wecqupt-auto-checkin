package tool

var DayCode = []string{
	"s9ZS",
	"jQkB",
	"RuQM",
	"O0_L",
	"Buxf",
	"LepV",
	"Ec6w",
	"zPLD",
	"eZry",
	"QjBF",
	"XPB0",
	"zlTr",
	"YDr2",
	"Mfdu",
	"HSoi",
	"frhT",
	"GOdB",
	"AEN0",
	"zX0T",
	"wJg1",
	"fCmn",
	"SM3z",
	"2U5I",
	"LI3u",
	"3rAY",
	"aoa4",
	"Jf9u",
	"M69T",
	"XCea",
	"63gc",
	"6_Kf",
}

var hourCode = []string{
	"89KC",
	"pzTS",
	"wgte",
	"29_3",
	"GpdG",
	"FDYl",
	"vsE9",
	"SPJk",
	"_buC",
	"GPHN",
	"OKax",
	"_Kk4",
	"hYxa",
	"1BC5",
	"oBk_",
	"JgUW",
	"0CPR",
	"jlEh",
	"gBGg",
	"frS6",
	"4ads",
	"Iwfk",
	"TCgR",
	"wbjP",
}

//生成 mrdkkey
//https://github.com/leishui/WCY 感谢学长的文档

func GetMirKey(day, hour int) string {
	return DayCode[day] + hourCode[hour]
}

