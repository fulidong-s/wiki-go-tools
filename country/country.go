package country

import "strings"

type Country struct {
	CountryCode        string
	TwoCharCode        string
	ThreeCharCode      string
	EnglishCountryName string
	ChineseCountryName string
	AreaCode           string
}

type Area struct {
	CountryCode string
}

var (
	twoCharCodeToCountryCodeMapping = map[string]string{
		"AD": "020",
		"AE": "784",
		"AF": "004",
		"AG": "028",
		"AI": "660",
		"AL": "008",
		"AM": "051",
		"AO": "024",
		"AQ": "010",
		"AR": "032",
		"AS": "016",
		"AT": "040",
		"AU": "036",
		"AW": "533",
		"AX": "248",
		"AZ": "031",
		"BA": "070",
		"BB": "052",
		"BD": "050",
		"BE": "056",
		"BF": "854",
		"BG": "100",
		"BH": "048",
		"BI": "108",
		"BJ": "204",
		"BL": "652",
		"BM": "060",
		"BN": "096",
		"BO": "068",
		"BQ": "535",
		"BR": "076",
		"BS": "044",
		"BT": "064",
		"BV": "074",
		"BW": "072",
		"BY": "112",
		"BZ": "084",
		"CA": "124",
		"CC": "166",
		"CD": "180",
		"CF": "140",
		"CG": "178",
		"CH": "756",
		"CI": "384",
		"CK": "184",
		"CL": "152",
		"CM": "120",
		"CN": "156",
		"CO": "170",
		"CR": "188",
		"CU": "192",
		"CV": "132",
		"CX": "162",
		"CY": "196",
		"CZ": "203",
		"DE": "276",
		"DJ": "262",
		"DK": "208",
		"DM": "212",
		"DO": "214",
		"DZ": "012",
		"EC": "218",
		"EE": "233",
		"EG": "818",
		"EH": "732",
		"ER": "232",
		"ES": "724",
		"ET": "231",
		"FI": "246",
		"FJ": "242",
		"FK": "238",
		"FM": "583",
		"FO": "234",
		"FR": "250",
		"GA": "266",
		"GD": "308",
		"GE": "268",
		"GF": "254",
		"GG": "831",
		"GH": "288",
		"GI": "292",
		"GL": "304",
		"GM": "270",
		"GN": "324",
		"GP": "312",
		"GQ": "226",
		"GR": "300",
		"GS": "239",
		"GT": "320",
		"GU": "316",
		"GW": "624",
		"GY": "328",
		"HK": "344",
		"HM": "334",
		"HN": "340",
		"HR": "191",
		"HT": "332",
		"HU": "348",
		"ID": "360",
		"IE": "372",
		"IL": "376",
		"IM": "833",
		"IN": "356",
		"IO": "086",
		"IQ": "368",
		"IR": "364",
		"IS": "352",
		"IT": "380",
		"JE": "832",
		"JM": "388",
		"JO": "400",
		"JP": "392",
		"KE": "404",
		"KG": "417",
		"KH": "116",
		"KI": "296",
		"KM": "174",
		"KN": "659",
		"KP": "408",
		"KR": "410",
		"KW": "414",
		"KY": "136",
		"KZ": "398",
		"LA": "418",
		"LB": "422",
		"LC": "662",
		"LI": "438",
		"LK": "144",
		"LR": "430",
		"LS": "426",
		"LT": "440",
		"LU": "442",
		"LV": "428",
		"LY": "434",
		"MA": "504",
		"MC": "492",
		"MD": "498",
		"ME": "499",
		"MF": "663",
		"MG": "450",
		"MH": "584",
		"MK": "807",
		"ML": "466",
		"MM": "104",
		"MN": "496",
		"MO": "446",
		"MP": "580",
		"MQ": "474",
		"MR": "478",
		"MS": "500",
		"MT": "470",
		"MU": "480",
		"MV": "462",
		"MW": "454",
		"MX": "484",
		"MY": "458",
		"MZ": "508",
		"NA": "516",
		"NC": "540",
		"NE": "562",
		"NF": "574",
		"NG": "566",
		"NI": "558",
		"NL": "528",
		"NO": "578",
		"NP": "524",
		"NR": "520",
		"NU": "570",
		"NZ": "554",
		"OM": "512",
		"PA": "591",
		"PE": "604",
		"PF": "258",
		"PG": "598",
		"PH": "608",
		"PK": "586",
		"PL": "616",
		"PM": "666",
		"PN": "612",
		"PR": "630",
		"PS": "275",
		"PT": "620",
		"PW": "585",
		"PY": "600",
		"QA": "634",
		"QJ": "999",
		"RE": "638",
		"RO": "642",
		"RS": "688",
		"RU": "643",
		"RW": "646",
		"SA": "682",
		"SB": "090",
		"SC": "690",
		"SD": "729",
		"SE": "752",
		"SG": "702",
		"SH": "654",
		"SI": "705",
		"SJ": "744",
		"SK": "703",
		"SL": "694",
		"SM": "674",
		"SN": "686",
		"SO": "706",
		"SR": "740",
		"SS": "728",
		"ST": "678",
		"SV": "222",
		"SY": "760",
		"SZ": "748",
		"TC": "796",
		"TD": "148",
		"TF": "260",
		"TG": "768",
		"TH": "764",
		"TJ": "762",
		"TK": "772",
		"TL": "626",
		"TM": "795",
		"TN": "788",
		"TO": "776",
		"TR": "792",
		"TT": "780",
		"TV": "798",
		"TW": "158",
		"TZ": "834",
		"UA": "804",
		"UG": "800",
		"UK": "826",
		"UM": "581",
		"US": "840",
		"UY": "858",
		"UZ": "860",
		"VA": "336",
		"VC": "670",
		"VE": "862",
		"VG": "092",
		"VI": "850",
		"VN": "704",
		"VU": "548",
		"WF": "876",
		"WS": "882",
		"YE": "887",
		"YT": "175",
		"ZA": "710",
		"ZM": "894",
		"ZW": "716",
	}
	threeCharCodeToCountryCodeMapping = map[string]string{
		"AND": "020",
		"ARE": "784",
		"AFG": "004",
		"ATG": "028",
		"AIA": "660",
		"ALB": "008",
		"ARM": "051",
		"AGO": "024",
		"ATA": "010",
		"ARG": "032",
		"ASM": "016",
		"AUT": "040",
		"AUS": "036",
		"ABW": "533",
		"ALA": "248",
		"AZE": "031",
		"BIH": "070",
		"BRB": "052",
		"BGD": "050",
		"BEL": "056",
		"BFA": "854",
		"BGR": "100",
		"BHR": "048",
		"BDI": "108",
		"BEN": "204",
		"BLM": "652",
		"BMU": "060",
		"BRN": "096",
		"BOL": "068",
		"BES": "535",
		"BRA": "076",
		"BHS": "044",
		"BTN": "064",
		"BVT": "074",
		"BWA": "072",
		"BLR": "112",
		"BLZ": "084",
		"CAN": "124",
		"CCK": "166",
		"COD": "180",
		"CAF": "140",
		"COG": "178",
		"CHE": "756",
		"CIV": "384",
		"COK": "184",
		"CHL": "152",
		"CMR": "120",
		"CHN": "156",
		"COL": "170",
		"CRI": "188",
		"CUB": "192",
		"CPV": "132",
		"CXR": "162",
		"CYP": "196",
		"CZE": "203",
		"DEU": "276",
		"DJI": "262",
		"DNK": "208",
		"DMA": "212",
		"DOM": "214",
		"DZA": "012",
		"ECU": "218",
		"EST": "233",
		"EGY": "818",
		"ESH": "732",
		"ERI": "232",
		"ESP": "724",
		"ETH": "231",
		"FIN": "246",
		"FJI": "242",
		"FLK": "238",
		"FSM": "583",
		"FRO": "234",
		"FRA": "250",
		"GAB": "266",
		"GRD": "308",
		"GEO": "268",
		"GUF": "254",
		"GGY": "831",
		"GHA": "288",
		"GIB": "292",
		"GRL": "304",
		"GMB": "270",
		"GIN": "324",
		"GLP": "312",
		"GNQ": "226",
		"GRC": "300",
		"SGS": "239",
		"GTM": "320",
		"GUM": "316",
		"GNB": "624",
		"GUY": "328",
		"HKG": "344",
		"HMD": "334",
		"HND": "340",
		"HRV": "191",
		"HTI": "332",
		"HUN": "348",
		"IDN": "360",
		"IRL": "372",
		"ISR": "376",
		"IMN": "833",
		"IND": "356",
		"IOT": "086",
		"IRQ": "368",
		"IRN": "364",
		"ISL": "352",
		"ITA": "380",
		"JEY": "832",
		"JAM": "388",
		"JOR": "400",
		"JPN": "392",
		"KEN": "404",
		"KGZ": "417",
		"KHM": "116",
		"KIR": "296",
		"COM": "174",
		"KNA": "659",
		"PRK": "408",
		"KOR": "410",
		"KWT": "414",
		"CYM": "136",
		"KAZ": "398",
		"LAO": "418",
		"LBN": "422",
		"LCA": "662",
		"LIE": "438",
		"LKA": "144",
		"LBR": "430",
		"LSO": "426",
		"LTU": "440",
		"LUX": "442",
		"LVA": "428",
		"LBY": "434",
		"MAR": "504",
		"MCO": "492",
		"MDA": "498",
		"MNE": "499",
		"MAF": "663",
		"MDG": "450",
		"MHL": "584",
		"MKD": "807",
		"MLI": "466",
		"MMR": "104",
		"MNG": "496",
		"MAC": "446",
		"MNP": "580",
		"MTQ": "474",
		"MRT": "478",
		"MSR": "500",
		"MLT": "470",
		"MUS": "480",
		"MDV": "462",
		"MWI": "454",
		"MEX": "484",
		"MYS": "458",
		"MOZ": "508",
		"NAM": "516",
		"NCL": "540",
		"NER": "562",
		"NFK": "574",
		"NGA": "566",
		"NIC": "558",
		"NLD": "528",
		"NOR": "578",
		"NPL": "524",
		"NRU": "520",
		"NIU": "570",
		"NZL": "554",
		"OMN": "512",
		"PAN": "591",
		"PER": "604",
		"PYF": "258",
		"PNG": "598",
		"PHL": "608",
		"PAK": "586",
		"POL": "616",
		"SPM": "666",
		"PCN": "612",
		"PRI": "630",
		"PSE": "275",
		"PRT": "620",
		"PLW": "585",
		"PRY": "600",
		"QAT": "634",
		"QJ":  "999",
		"REU": "638",
		"ROU": "642",
		"SRB": "688",
		"RUS": "643",
		"RWA": "646",
		"SAU": "682",
		"SLB": "090",
		"SYC": "690",
		"SDN": "729",
		"SWE": "752",
		"SGP": "702",
		"SHN": "654",
		"SVN": "705",
		"SJM": "744",
		"SVK": "703",
		"SLE": "694",
		"SMR": "674",
		"SEN": "686",
		"SOM": "706",
		"SUR": "740",
		"SSD": "728",
		"STP": "678",
		"SLV": "222",
		"SYR": "760",
		"SWZ": "748",
		"TCA": "796",
		"TCD": "148",
		"ATF": "260",
		"TGO": "768",
		"THA": "764",
		"TJK": "762",
		"TKL": "772",
		"TLS": "626",
		"TKM": "795",
		"TUN": "788",
		"TON": "776",
		"TUR": "792",
		"TTO": "780",
		"TUV": "798",
		"TWN": "158",
		"TZA": "834",
		"UKR": "804",
		"UGA": "800",
		"GBR": "826",
		"UMI": "581",
		"USA": "840",
		"URY": "858",
		"UZB": "860",
		"VAT": "336",
		"VCT": "670",
		"VEN": "862",
		"VGB": "092",
		"VIR": "850",
		"VNM": "704",
		"VUT": "548",
		"WLF": "876",
		"WSM": "882",
		"YEM": "887",
		"MYT": "175",
		"ZAF": "710",
		"ZMB": "894",
		"ZWE": "716",
	}
	countryCodeMapping = map[string]*Country{
		"020": {"020", "AD", "AND", "Andorra", "安道尔", "00376"},
		"784": {"784", "AE", "ARE", "United Arab Emirates", "阿联酋", "00971"},
		"004": {"004", "AF", "AFG", "Afghanistan", "阿富汗	", "0093"},
		"028": {"028", "AG", "ATG", "Antigua & Barbuda", "安提瓜和巴布达", "001268"},
		"660": {"660", "AI", "AIA", "Anguilla", "安圭拉", "001264"},
		"008": {"008", "AL", "ALB", "Albania", "阿尔巴尼亚", "00355"},
		"051": {"051", "AM", "ARM", "Armenia", "亚美尼亚", "00374"},
		"024": {"024", "AO", "AGO", "Angola", "安哥拉", "00244"},
		"010": {"010", "AQ", "ATA", "Antarctica", "南极洲", "00672"},
		"032": {"032", "AR", "ARG", "Argentina", "阿根廷", "0054"},
		"016": {"016", "AS", "ASM", "American Samoa", "美属萨摩亚", "001684"},
		"040": {"040", "AT", "AUT", "Austria", "奥地利	", "0043"},
		"036": {"036", "AU", "AUS", "Australia", "澳大利亚", "0061"},
		"533": {"533", "AW", "ABW", "Aruba", "阿鲁巴", "00297"},
		"248": {"248", "AX", "ALA", "Aland Island", "奥兰群岛", "(null)"},
		"031": {"031", "AZ", "AZE", "Azerbaijan", "阿塞拜疆", "00994"},
		"070": {"070", "BA", "BIH", "Bosnia & Herzegovina", "波黑", "00387"},
		"052": {"052", "BB", "BRB", "Barbados", "巴巴多斯", "001246"},
		"050": {"050", "BD", "BGD", "Bangladesh", "孟加拉", "00880"},
		"056": {"056", "BE", "BEL", "Belgium", "比利时	", "0032"},
		"854": {"854", "BF", "BFA", "Burkina", "布基纳法索", "00226"},
		"100": {"100", "BG", "BGR", "Bulgaria", "保加利亚", "00359"},
		"048": {"048", "BH", "BHR", "Bahrain", "巴林", "00973"},
		"108": {"108", "BI", "BDI", "Burundi", "布隆迪	", "00257"},
		"204": {"204", "BJ", "BEN", "Benin", "贝宁", "00229"},
		"652": {"652", "BL", "BLM", "Saint Barthélemy", "圣巴泰勒米岛", "00590"},
		"060": {"060", "BM", "BMU", "Bermuda", "百慕大	", "001441"},
		"096": {"096", "BN", "BRN", "Brunei", "文莱", "00673"},
		"068": {"068", "BO", "BOL", "Bolivia", "玻利维亚", "00591"},
		"535": {"535", "BQ", "BES", "Caribbean Netherlands", "荷兰加勒比区", "(null)"},
		"076": {"076", "BR", "BRA", "Brazil", "巴西", "0055"},
		"044": {"044", "BS", "BHS", "The Bahamas", "巴哈马	", "001242"},
		"064": {"064", "BT", "BTN", "Bhutan", "不丹", "00975"},
		"074": {"074", "BV", "BVT", "Bouvet Island", "布韦岛", "(null)"},
		"072": {"072", "BW", "BWA", "Botswana", "博茨瓦纳", "00267"},
		"112": {"112", "BY", "BLR", "Belarus", "白俄罗斯", "00375"},
		"084": {"084", "BZ", "BLZ", "Belize", "伯利兹", "00501"},
		"124": {"124", "CA", "CAN", "Canada", "加拿大", "001"},
		"166": {"166", "CC", "CCK", "Cocos (Keeling) Islands", "科科斯群岛", "0061"},
		"180": {"180", "CD", "COD", "Democratic Republic of the Congo", "刚果（金）", "00243"},
		"140": {"140", "CF", "CAF", "Central African Republic", "中非", "00236"},
		"178": {"178", "CG", "COG", "Republic of the Congo", "刚果（布）", "00242"},
		"756": {"756", "CH", "CHE", "Switzerland", "瑞士", "0041"},
		"384": {"384", "CI", "CIV", "Cote d'ivoire", "科特迪瓦", "00225"},
		"184": {"184", "CK", "COK", "Cook Islands", "库克群岛", "00682"},
		"152": {"152", "CL", "CHL", "Chile", "智利", "0056"},
		"120": {"120", "CM", "CMR", "Cameroon", "喀麦隆", "00237"},
		"156": {"156", "CN", "CHN", "China", "中国", "0086"},
		"170": {"170", "CO", "COL", "Colombia", "哥伦比亚", "0057"},
		"188": {"188", "CR", "CRI", "Costa Rica", "哥斯达黎加", "00506"},
		"192": {"192", "CU", "CUB", "Cuba", "古巴", "0053"},
		"132": {"132", "CV", "CPV", "Cape Verde", "佛得角", "00238"},
		"162": {"162", "CX", "CXR", "Christmas Island", "圣诞岛", "0061"},
		"196": {"196", "CY", "CYP", "Cyprus", "塞浦路斯", "00357"},
		"203": {"203", "CZ", "CZE", "Czech Republic", "捷克", "00420"},
		"276": {"276", "DE", "DEU", "Germany", "德国", "0049"},
		"262": {"262", "DJ", "DJI", "Djibouti", "吉布提", "00253"},
		"208": {"208", "DK", "DNK", "Denmark", "丹麦", "0045"},
		"212": {"212", "DM", "DMA", "Dominica", "多米尼克", "001767"},
		"214": {"214", "DO", "DOM", "Dominican Republic", "多米尼加", "001849"},
		"012": {"012", "DZ", "DZA", "Algeria", "阿尔及利亚", "00213"},
		"218": {"218", "EC", "ECU", "Ecuador", "厄瓜多尔", "00593"},
		"233": {"233", "EE", "EST", "Estonia", "爱沙尼亚", "00372"},
		"818": {"818", "EG", "EGY", "Egypt", "埃及", "0020"},
		"732": {"732", "EH", "ESH", "Western Sahara", "西撒哈拉", "00212"},
		"232": {"232", "ER", "ERI", "Eritrea", "厄立特里亚", "00291"},
		"724": {"724", "ES", "ESP", "Spain", "西班牙", "0034"},
		"231": {"231", "ET", "ETH", "Ethiopia", "埃塞俄比亚", "00251"},
		"246": {"246", "FI", "FIN", "Finland", "芬兰", "00358"},
		"242": {"242", "FJ", "FJI", "Fiji", "斐济群岛", "00679"},
		"238": {"238", "FK", "FLK", "Falkland Islands", "马尔维纳斯群岛（ 福克兰）", "00500"},
		"583": {"583", "FM", "FSM", "Federated States of Micronesia", "密克罗尼西亚联邦", "00691"},
		"234": {"234", "FO", "FRO", "Faroe Islands", "法罗群岛", "00298"},
		"250": {"250", "FR", "FRA", "France", "法国", "0033"},
		"266": {"266", "GA", "GAB", "Gabon", "加蓬", "00241"},
		"308": {"308", "GD", "GRD", "Grenada", "格林纳达", "001473"},
		"268": {"268", "GE", "GEO", "Georgia", "格鲁吉亚", "00995"},
		"254": {"254", "GF", "GUF", "French Guiana", "法属圭亚那", "00594"},
		"831": {"831", "GG", "GGY", "Guernsey", "根西岛", "(null)"},
		"288": {"288", "GH", "GHA", "Ghana", "加纳", "00233"},
		"292": {"292", "GI", "GIB", "Gibraltar", "直布罗陀", "00350"},
		"304": {"304", "GL", "GRL", "Greenland", "格陵兰", "00299"},
		"270": {"270", "GM", "GMB", "Gambia", "冈比亚", "00220"},
		"324": {"324", "GN", "GIN", "Guinea", "几内亚", "00224"},
		"312": {"312", "GP", "GLP", "Guadeloupe", "瓜德罗普", "00590"},
		"226": {"226", "GQ", "GNQ", "Equatorial Guinea", "赤道几内亚", "00240"},
		"300": {"300", "GR", "GRC", "Greece", "希腊", "0030"},
		"239": {"239", "GS", "SGS", "South Georgia and the South Sandwich Islands", "南乔治亚岛和南桑威奇群岛", "(null)"},
		"320": {"320", "GT", "GTM", "Guatemala", "危地马拉", "00502"},
		"316": {"316", "GU", "GUM", "Guam", "关岛", "001671"},
		"624": {"624", "GW", "GNB", "Guinea-Bissau", "几内亚比绍", "00245"},
		"328": {"328", "GY", "GUY", "Guyana", "圭亚那", "00592"},
		"344": {"344", "HK", "HKG", "Hong Kong", "中国香港", "00852"},
		"334": {"334", "HM", "HMD", "Heard Island and McDonald Islands", "赫德岛和麦克唐纳群岛", "(null)"},
		"340": {"340", "HN", "HND", "Honduras", "洪都拉斯", "00504"},
		"191": {"191", "HR", "HRV", "Croatia", "克罗地亚", "00385"},
		"332": {"332", "HT", "HTI", "Haiti", "海地", "00509"},
		"348": {"348", "HU", "HUN", "Hungary", "匈牙利	", "0036"},
		"360": {"360", "ID", "IDN", "Indonesia", "印尼", "0062"},
		"372": {"372", "IE", "IRL", "Ireland", "爱尔兰	", "00353"},
		"376": {"376", "IL", "ISR", "Israel", "以色列", "00972"},
		"833": {"833", "IM", "IMN", "Isle of Man", "马恩岛	", "0044"},
		"356": {"356", "IN", "IND", "India", "印度", "0091"},
		"086": {"086", "IO", "IOT", "British Indian Ocean Territory", "印度洋领地", "00246"},
		"368": {"368", "IQ", "IRQ", "Iraq", "伊拉克", "00964"},
		"364": {"364", "IR", "IRN", "Iran", "伊朗", "0098"},
		"352": {"352", "IS", "ISL", "Iceland", "冰岛", "00354"},
		"380": {"380", "IT", "ITA", "Italy", "意大利", "0039"},
		"832": {"832", "JE", "JEY", "Jersey", "泽西岛", "0044"},
		"388": {"388", "JM", "JAM", "Jamaica", "牙买加	", "001876"},
		"400": {"400", "JO", "JOR", "Jordan", "约旦", "00962"},
		"392": {"392", "JP", "JPN", "Japan", "日本", "0081"},
		"404": {"404", "KE", "KEN", "Kenya", "肯尼亚", "00254"},
		"417": {"417", "KG", "KGZ", "Kyrgyzstan", "吉尔吉斯斯坦", "00996"},
		"116": {"116", "KH", "KHM", "Cambodia", "柬埔寨", "00855"},
		"296": {"296", "KI", "KIR", "Kiribati", "基里巴斯", "00686"},
		"174": {"174", "KM", "COM", "The Comoros", "科摩罗	", "00269"},
		"659": {"659", "KN", "KNA", "St. Kitts & Nevis", "圣基茨和尼维斯", "001869"},
		"408": {"408", "KP", "PRK", "North Korea", "朝鲜", "00850"},
		"410": {"410", "KR", "KOR", "South Korea", "韩国", "0082"},
		"414": {"414", "KW", "KWT", "Kuwait", "科威特", "00965"},
		"136": {"136", "KY", "CYM", "Cayman Islands", "开曼群岛", "001345"},
		"398": {"398", "KZ", "KAZ", "Kazakhstan", "哈萨克斯坦", "007"},
		"418": {"418", "LA", "LAO", "Laos", "老挝", "00856"},
		"422": {"422", "LB", "LBN", "Lebanon", "黎巴嫩	", "00961"},
		"662": {"662", "LC", "LCA", "St. Lucia", "圣卢西亚", "001758"},
		"438": {"438", "LI", "LIE", "Liechtenstein", "列支敦士登", "00423"},
		"144": {"144", "LK", "LKA", "Sri Lanka", "斯里兰卡", "0094"},
		"430": {"430", "LR", "LBR", "Liberia", "利比里亚", "00231"},
		"426": {"426", "LS", "LSO", "Lesotho", "莱索托	", "00266"},
		"440": {"440", "LT", "LTU", "Lithuania", "立陶宛", "00370"},
		"442": {"442", "LU", "LUX", "Luxembourg", "卢森堡", "00352"},
		"428": {"428", "LV", "LVA", "Latvia", "拉脱维亚", "00371"},
		"434": {"434", "LY", "LBY", "Libya", "利比亚", "00218"},
		"504": {"504", "MA", "MAR", "Morocco", "摩洛哥	", "0212"},
		"492": {"492", "MC", "MCO", "Monaco", "摩纳哥", "00377"},
		"498": {"498", "MD", "MDA", "Moldova", "摩尔多瓦", "00373"},
		"499": {"499", "ME", "MNE", "Montenegro", "黑山", "00382"},
		"663": {"663", "MF", "MAF", "Saint Martin (France)", "法属圣马丁", "(null)"},
		"450": {"450", "MG", "MDG", "Madagascar", "马达加斯加", "00261"},
		"584": {"584", "MH", "MHL", "Marshall islands", "马绍尔群岛", "00692"},
		"807": {"807", "MK", "MKD", "Republic of Macedonia (FYROM)", "马其顿", "00389"},
		"466": {"466", "ML", "MLI", "Mali", "马里", "00223"},
		"104": {"104", "MM", "MMR", "Myanmar (Burma)", "缅甸", "0095"},
		"496": {"496", "MN", "MNG", "Mongolia", "蒙古国 蒙古", "00976"},
		"446": {"446", "MO", "MAC", "Macao", "中国澳门", "00853"},
		"580": {"580", "MP", "MNP", "Northern Mariana Islands", "北马里亚纳群岛", "001670"},
		"474": {"474", "MQ", "MTQ", "Martinique", "马提尼克", "00596"},
		"478": {"478", "MR", "MRT", "Mauritania", "毛里塔尼亚", "00222"},
		"500": {"500", "MS", "MSR", "Montserrat", "蒙塞拉特岛", "(null)"},
		"470": {"470", "MT", "MLT", "Malta", "马耳他", "00356"},
		"480": {"480", "MU", "MUS", "Mauritius", "毛里求斯", "00230"},
		"462": {"462", "MV", "MDV", "Maldives", "马尔代夫", "00960"},
		"454": {"454", "MW", "MWI", "Malawi", "马拉维", "00265"},
		"484": {"484", "MX", "MEX", "Mexico", "墨西哥", "0052"},
		"458": {"458", "MY", "MYS", "Malaysia", "马来西亚", "0060"},
		"508": {"508", "MZ", "MOZ", "Mozambique", "莫桑比克", "00258"},
		"516": {"516", "NA", "NAM", "Namibia", "纳米比亚", "00264"},
		"540": {"540", "NC", "NCL", "New Caledonia", "新喀里多尼亚", "00687"},
		"562": {"562", "NE", "NER", "Niger", "尼日尔", "00227"},
		"574": {"574", "NF", "NFK", "Norfolk Island", "诺福克岛", "00672"},
		"566": {"566", "NG", "NGA", "Nigeria", "尼日利亚", "00234"},
		"558": {"558", "NI", "NIC", "Nicaragua", "尼加拉瓜", "00505"},
		"528": {"528", "NL", "NLD", "Netherlands", "荷兰", "0031"},
		"578": {"578", "NO", "NOR", "Norway", "挪威", "0047"},
		"524": {"524", "NP", "NPL", "Nepal", "尼泊尔", "00977"},
		"520": {"520", "NR", "NRU", "Nauru", "瑙鲁", "00674"},
		"570": {"570", "NU", "NIU", "Niue", "纽埃", "00683"},
		"554": {"554", "NZ", "NZL", "New Zealand", "新西兰	", "0064"},
		"512": {"512", "OM", "OMN", "Oman", "阿曼", "00968"},
		"591": {"591", "PA", "PAN", "Panama", "巴拿马", "00507"},
		"604": {"604", "PE", "PER", "Peru", "秘鲁", "0051"},
		"258": {"258", "PF", "PYF", "French polynesia", "法属波利尼西亚", "00689"},
		"598": {"598", "PG", "PNG", "Papua New Guinea", "巴布亚新几内亚", "00675"},
		"608": {"608", "PH", "PHL", "The Philippines", "菲律宾", "0063"},
		"586": {"586", "PK", "PAK", "Pakistan", "巴基斯坦", "0092"},
		"616": {"616", "PL", "POL", "Poland", "波兰", "0048"},
		"666": {"666", "PM", "SPM", "Saint-Pierre and Miquelon", "圣皮埃尔和密克隆", "00508"},
		"612": {"612", "PN", "PCN", "Pitcairn Islands", "皮特凯恩群岛", "00870"},
		"630": {"630", "PR", "PRI", "Puerto Rico", "波多黎各", "001939"},
		"275": {"275", "PS", "PSE", "Palestinian territories", "巴勒斯坦", "00970"},
		"620": {"620", "PT", "PRT", "Portugal", "葡萄牙", "00351"},
		"585": {"585", "PW", "PLW", "Palau", "帕劳", "00680"},
		"600": {"600", "PY", "PRY", "Paraguay", "巴拉圭", "00595"},
		"634": {"634", "QA", "QAT", "Qatar", "卡塔尔", "00974"},
		"999": {"999", "QJ", "QJ	", "quanjiu", "默认", "999"},
		"638": {"638", "RE", "REU", "Réunion", "留尼汪	", "00262"},
		"642": {"642", "RO", "ROU", "Romania", "罗马尼亚", "0040"},
		"688": {"688", "RS", "SRB", "Serbia", "塞尔维亚", "00381"},
		"643": {"643", "RU", "RUS", "Russian Federation", "俄罗斯", "007"},
		"646": {"646", "RW", "RWA", "Rwanda", "卢旺达", "00250"},
		"682": {"682", "SA", "SAU", "Saudi Arabia", "沙特阿拉伯", "00966"},
		"090": {"090", "SB", "SLB", "Solomon Islands", "所罗门群岛", "00677"},
		"690": {"690", "SC", "SYC", "Seychelles", "塞舌尔", "00248"},
		"729": {"729", "SD", "SDN", "Sudan", "苏丹", "00249"},
		"752": {"752", "SE", "SWE", "Sweden", "瑞典", "0046"},
		"702": {"702", "SG", "SGP", "Singapore", "新加坡", "0065"},
		"654": {"654", "SH", "SHN", "St. Helena & Dependencies", "圣赫勒拿", "00290"},
		"705": {"705", "SI", "SVN", "Slovenia", "斯洛文尼亚", "00386"},
		"744": {"744", "SJ", "SJM", "Template:Country data SJM Svalbard", "斯瓦尔巴群岛和 扬马延岛", "(null)"},
		"703": {"703", "SK", "SVK", "Slovakia", "斯洛伐克", "00421"},
		"694": {"694", "SL", "SLE", "Sierra Leone", "塞拉利昂", "00232"},
		"674": {"674", "SM", "SMR", "San Marino", "圣马力诺", "00378"},
		"686": {"686", "SN", "SEN", "Senegal", "塞内加尔", "00221"},
		"706": {"706", "SO", "SOM", "Somalia", "索马里	", "00252"},
		"740": {"740", "SR", "SUR", "Suriname", "苏里南", "00597"},
		"728": {"728", "SS", "SSD", "South Sudan", "南苏丹	", "00211"},
		"678": {"678", "ST", "STP", "Sao Tome & Principe", "圣多美和普林西比	", "00239"},
		"222": {"222", "SV", "SLV", "El Salvador", "萨尔瓦多", "00503"},
		"760": {"760", "SY", "SYR", "Syria", "叙利亚", "00963"},
		"748": {"748", "SZ", "SWZ", "Swaziland", "斯威士兰", "00268"},
		"796": {"796", "TC", "TCA", "Turks & Caicos Islands", "特克斯和凯科斯群岛", "(null)"},
		"148": {"148", "TD", "TCD", "Chad", "乍得", "00235"},
		"260": {"260", "TF", "ATF", "French Southern Territories", "法属南部领地", "(null)"},
		"768": {"768", "TG", "TGO", "Togo", "多哥", "00228"},
		"764": {"764", "TH", "THA", "Thailand", "泰国", "0066"},
		"762": {"762", "TJ", "TJK", "Tajikistan", "塔吉克斯坦", "00992"},
		"772": {"772", "TK", "TKL", "Tokelau", "托克劳	", "00690"},
		"626": {"626", "TL", "TLS", "Timor-Leste (East Timor)", "东帝汶", "00670"},
		"795": {"795", "TM", "TKM", "Turkmenistan", "土库曼斯坦", "00993"},
		"788": {"788", "TN", "TUN", "Tunisia", "突尼斯	", "00216"},
		"776": {"776", "TO", "TON", "Tonga", "汤加", "00676"},
		"792": {"792", "TR", "TUR", "Turkey", "土耳其", "0090"},
		"780": {"780", "TT", "TTO", "Trinidad & Tobago", "特立尼达和多巴哥", "001868"},
		"798": {"798", "TV", "TUV", "Tuvalu", "图瓦卢", "00688"},
		"158": {"158", "TW", "TWN", "Taiwan", "中国台湾", "00886"},
		"834": {"834", "TZ", "TZA", "Tanzania", "坦桑尼亚", "00255"},
		"804": {"804", "UA", "UKR", "Ukraine", "乌克兰	", "00380"},
		"800": {"800", "UG", "UGA", "Uganda", "乌干达", "00256"},
		"826": {"826", "UK", "GBR", "Great Britain (United Kingdom; England)", "英国", "0044"},
		"581": {"581", "UM", "UMI", "United States Minor Outlying Islands", "美国本土外小岛屿", "(null)"},
		"840": {"840", "US", "USA", "United States of America", "美国", "001"},
		"858": {"858", "UY", "URY", "Uruguay", "乌拉圭	", "00598"},
		"860": {"860", "UZ", "UZB", "Uzbekistan", "乌兹别克斯坦	", "00998"},
		"336": {"336", "VA", "VAT", "Vatican City (The Holy See)", "梵蒂冈	", "00379"},
		"670": {"670", "VC", "VCT", "St. Vincent & the Grenadines", "圣文森特和格林纳丁斯", "001784"},
		"862": {"862", "VE", "VEN", "Venezuela", "委内瑞拉", "0058"},
		"092": {"092", "VG", "VGB", "British Virgin Islands", "维尔京群岛", "001284"},
		"850": {"850", "VI", "VIR", "United States Virgin Islands", "美属维尔京群岛", "001340"},
		"704": {"704", "VN", "VNM", "Vietnam", "越南", "0084"},
		"548": {"548", "VU", "VUT", "Vanuatu", "瓦努阿图", "00678"},
		"876": {"876", "WF", "WLF", "Wallis and Futuna", "瓦利斯和富图纳", "00681"},
		"882": {"882", "WS", "WSM", "Samoa", "萨摩亚", "00685"},
		"887": {"887", "YE", "YEM", "Yemen", "也门", "00967"},
		"175": {"175", "YT", "MYT", "Mayotte", "马约特	", "00262"},
		"710": {"710", "ZA", "ZAF", "South Africa", "南非", "0027"},
		"894": {"894", "ZM", "ZMB", "Zambia", "赞比亚", "00260"},
		"716": {"716", "ZW", "ZWE", "Zimbabwe", "津巴布韦", "00263"},
	}
	countryAreaMapping = map[string]string{
		"158": "14481060",
		"156": "26731356",
		"458": "91405248",
		"704": "36431431",
		"360": "29406706",
		"764": "12398727",
		"608": "28942458",
		"096": "93609133",
		"104": "93609133",
		"418": "93609133",
		"446": "93609133",
		"175": "93609133",
		"242": "93609133",
		"474": "66006232",
		"728": "93609133",
		"258": "93609133",
		"162": "93609133",
		"166": "93609133",
		"184": "93609133",
		"234": "66006232",
		"248": "93609133",
		"798": "93609133",
		"584": "93609133",
		"570": "93609133",
		"520": "93609133",
		"663": "93609133",
		"744": "93609133",
		"074": "93609133",
		"239": "93609133",
		"580": "93609133",
		"583": "93609133",
		"776": "93609133",
		"612": "93609133",
		"772": "93609133",
		"598": "93609133",
		"585": "93609133",
		"882": "93609133",
		"574": "93609133",
		"548": "93609133",
		"540": "66006232",
		"500": "93609133",
		"334": "93609133",
		"292": "93609133",
		"296": "93609133",
		"090": "93609133",
		"016": "93609133",
		"028": "45030053",
		"010": "93609133",
		"032": "45030053",
		"044": "45030053",
		"052": "45030053",
		"060": "93609133",
		"076": "45030053",
		"084": "45030053",
		"068": "45030053",
		"086": "93609133",
		"092": "93609133",
		"124": "45030053",
		"136": "93609133",
		"152": "45030053",
		"170": "45030053",
		"188": "45030053",
		"192": "45030053",
		"212": "45030053",
		"214": "45030053",
		"218": "45030053",
		"222": "45030053",
		"238": "66006232",
		"260": "93609133",
		"304": "93609133",
		"308": "45030053",
		"316": "93609133",
		"320": "45030053",
		"328": "45030053",
		"332": "45030053",
		"340": "45030053",
		"388": "45030053",
		"484": "45030053",
		"533": "93609133",
		"535": "93609133",
		"558": "45030053",
		"581": "93609133",
		"591": "45030053",
		"600": "45030053",
		"630": "93609133",
		"604": "45030053",
		"659": "45030053",
		"662": "45030053",
		"666": "93609133",
		"670": "45030053",
		"740": "45030053",
		"780": "45030053",
		"796": "93609133",
		"840": "45030053",
		"850": "93609133",
		"858": "45030053",
		"862": "45030053",
		"008": "66006232",
		"020": "66006232",
		"040": "66006232",
		"100": "66006232",
		"070": "93609133",
		"112": "66006232",
		"056": "66006232",
		"191": "66006232",
		"203": "66006232",
		"208": "66006232",
		"233": "66006232",
		"250": "66006232",
		"246": "66006232",
		"276": "66006232",
		"300": "66006232",
		"254": "93609133",
		"312": "93609133",
		"336": "93609133",
		"348": "66006232",
		"352": "66006232",
		"428": "66006232",
		"438": "93609133",
		"372": "66006232",
		"380": "66006232",
		"440": "66006232",
		"470": "66006232",
		"442": "66006232",
		"498": "66006232",
		"492": "66006232",
		"499": "93609133",
		"528": "66006232",
		"578": "66006232",
		"616": "66006232",
		"620": "66006232",
		"642": "66006232",
		"643": "66006232",
		"652": "93609133",
		"674": "66006232",
		"660": "93609133",
		"703": "66006232",
		"688": "93609133",
		"724": "66006232",
		"752": "66006232",
		"705": "66006232",
		"804": "66006232",
		"807": "66006232",
		"756": "66006232",
		"831": "93609133",
		"826": "66006232",
		"833": "93609133",
		"832": "93609133",
		"876": "93609133",
		"004": "93609133",
		"398": "93609133",
		"408": "93609133",
		"417": "93609133",
		"496": "93609133",
		"626": "93609133",
		"762": "93609133",
		"795": "93609133",
		"860": "93609133",
		"392": "85830918",
		"024": "22923473",
		"854": "22923473",
		"108": "22923473",
		"204": "22923473",
		"072": "22923473",
		"180": "22923473",
		"140": "22923473",
		"178": "22923473",
		"120": "22923473",
		"384": "22923473",
		"132": "22923473",
		"262": "22923473",
		"818": "22923473",
		"012": "22923473",
		"732": "22923473",
		"231": "22923473",
		"266": "22923473",
		"232": "22923473",
		"270": "22923473",
		"288": "22923473",
		"324": "22923473",
		"226": "22923473",
		"624": "22923473",
		"404": "22923473",
		"174": "22923473",
		"430": "22923473",
		"426": "22923473",
		"434": "22923473",
		"504": "22923473",
		"450": "22923473",
		"466": "22923473",
		"478": "22923473",
		"480": "22923473",
		"454": "22923473",
		"508": "22923473",
		"516": "22923473",
		"562": "22923473",
		"638": "22923473",
		"729": "22923473",
		"646": "22923473",
		"690": "22923473",
		"654": "22923473",
		"694": "22923473",
		"686": "22923473",
		"706": "22923473",
		"678": "22923473",
		"748": "22923473",
		"768": "22923473",
		"148": "22923473",
		"788": "22923473",
		"834": "22923473",
		"800": "22923473",
		"710": "22923473",
		"894": "22923473",
		"716": "22923473",
		"566": "22923473",
		"344": "26731356",
		"702": "55196440",
		"196": "66006232",
		"116": "38153014",
		"356": "93011566",
		"784": "65998117",
		"050": "93011566",
		"064": "93011566",
		"462": "93011566",
		"144": "93011566",
		"524": "93011566",
		"586": "93011566",
		"051": "65998117",
		"031": "65998117",
		"048": "65998117",
		"268": "65998117",
		"376": "65998117",
		"368": "65998117",
		"364": "65998117",
		"400": "65998117",
		"414": "65998117",
		"422": "65998117",
		"512": "65998117",
		"275": "65998117",
		"634": "65998117",
		"682": "65998117",
		"760": "65998117",
		"792": "65998117",
		"887": "65998117",
		"410": "93609133",
		"036": "60414278",
		"554": "60414278",
	}
	areaNameMapping = map[string]string{
		"12398727": "泰国区",
		"13264857": "尼日利亚",
		"14481060": "台湾区",
		"22923473": "非洲区",
		"26731356": "香港区",
		"28942458": "菲律宾区",
		"29406706": "印尼区",
		"33873452": "大洋洲",
		"36431431": "越南区",
		"38153014": "柬埔寨",
		"45030053": "美洲",
		"47292932": "东南亚区",
		"54649693": "香港区（废弃）",
		"55196440": "新加坡区",
		"56780894": "韩国地区",
		"60414278": "澳新区",
		"65998117": "西亚区",
		"66006232": "欧洲",
		"85830918": "日本地区",
		"91405248": "马来区",
		"93011566": "南亚区",
		"93609133": "其他区",
		"999":      "全球",
	}
)

// GetAreaCodeByCode 同时兼容 二字码、三字码、CountryCode
func GetAreaCodeByCode(code string) string {
	// 先根据countryCode 获取
	out, ok := countryAreaMapping[code]
	if ok {
		return out
	}

	// 根据二字码获取country_code
	out, ok = countryAreaMapping[twoCharCodeToCountryCodeMapping[strings.ToUpper(code)]]
	if ok {
		return out
	}

	// 根据三字码获取
	out, ok = countryAreaMapping[threeCharCodeToCountryCodeMapping[strings.ToUpper(code)]]
	if ok {
		return out
	}
	return "93609133"
}

// GetCountryByCode 同时兼容 二字码、三字码、CountryCode
func GetCountryByCode(code string) (*Country, bool) {
	out, ok := countryCodeMapping[code]
	if ok {
		return out, true
	}

	out, ok = countryCodeMapping[twoCharCodeToCountryCodeMapping[strings.ToUpper(code)]]
	if ok {
		return out, true
	}

	out, ok = countryCodeMapping[threeCharCodeToCountryCodeMapping[strings.ToUpper(code)]]
	return out, ok
}

func GetAreaNameByAreaCode(code string) string {
	out, ok := areaNameMapping[code]
	if ok {
		return out
	}

	return "其他区"
}
