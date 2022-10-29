package currency

// Code is represent currency code.
type Code int16

//go:generate stringer -output=stringer.Code.go -type=Code
const (
	_ Code = iota
	AED
	AFN
	ALL
	AMD
	ANG
	AOA
	ARS
	AUD
	AWG
	AZN
	BAM
	BBD
	BDT
	BGN
	BHD
	BIF
	BMD
	BND
	BOB
	BRL
	BSD
	BTN
	BWP
	BYR
	BZD
	CAD
	CDF
	CHF
	CLP
	CNY
	COP
	CRC
	CUC
	CVE
	CZK
	DJF
	DKK
	DOP
	DZD
	EEK
	EGP
	ERN
	ETB
	EUR
	FJD
	FKP
	GBP
	GEL
	GHS
	GIP
	GMD
	GNF
	GQE
	GTQ
	GYD
	HKD
	HNL
	HRK
	HTG
	HUF
	IDR
	ILS
	INR
	IQD
	IRR
	ISK
	JMD
	JOD
	JPY
	KES
	KGS
	KHR
	KMF
	KPW
	KRW
	KWD
	KYD
	KZT
	LAK
	LBP
	LKR
	LRD
	LSL
	LTL
	LVL
	LYD
	MAD
	MDL
	MGA
	MKD
	MMK
	MNT
	MOP
	MRO
	MUR
	MVR
	MWK
	MXN
	MYR
	MZM
	NAD
	NGN
	NIO
	NOK
	NPR
	NZD
	OMR
	PAB
	PEN
	PGK
	PHP
	PKR
	PLN
	PYG
	QAR
	RON
	RSD
	RUB
	SAR
	SBD
	SCR
	SDG
	SEK
	SGD
	SHP
	SLL
	SOS
	SRD
	SYP
	SZL
	THB
	TJS
	TMT
	TND
	TRY
	TTD
	TWD
	TZS
	UAH
	UGX
	USD
	UYU
	UZS
	VEB
	VND
	VUV
	WST
	XAF
	XCD
	XDR
	XOF
	XPF
	YER
	ZAR
	ZMK
	ZWR
)

var currencies = map[string]Code{
	"AED": AED,
	"AFN": AFN,
	"ALL": ALL,
	"AMD": AMD,
	"ANG": ANG,
	"AOA": AOA,
	"ARS": ARS,
	"AUD": AUD,
	"AWG": AWG,
	"AZN": AZN,
	"BAM": BAM,
	"BBD": BBD,
	"BDT": BDT,
	"BGN": BGN,
	"BHD": BHD,
	"BIF": BIF,
	"BMD": BMD,
	"BND": BND,
	"BOB": BOB,
	"BRL": BRL,
	"BSD": BSD,
	"BTN": BTN,
	"BWP": BWP,
	"BYR": BYR,
	"BZD": BZD,
	"CAD": CAD,
	"CDF": CDF,
	"CHF": CHF,
	"CLP": CLP,
	"CNY": CNY,
	"COP": COP,
	"CRC": CRC,
	"CUC": CUC,
	"CVE": CVE,
	"CZK": CZK,
	"DJF": DJF,
	"DKK": DKK,
	"DOP": DOP,
	"DZD": DZD,
	"EEK": EEK,
	"EGP": EGP,
	"ERN": ERN,
	"ETB": ETB,
	"EUR": EUR,
	"FJD": FJD,
	"FKP": FKP,
	"GBP": GBP,
	"GEL": GEL,
	"GHS": GHS,
	"GIP": GIP,
	"GMD": GMD,
	"GNF": GNF,
	"GQE": GQE,
	"GTQ": GTQ,
	"GYD": GYD,
	"HKD": HKD,
	"HNL": HNL,
	"HRK": HRK,
	"HTG": HTG,
	"HUF": HUF,
	"IDR": IDR,
	"ILS": ILS,
	"INR": INR,
	"IQD": IQD,
	"IRR": IRR,
	"ISK": ISK,
	"JMD": JMD,
	"JOD": JOD,
	"JPY": JPY,
	"KES": KES,
	"KGS": KGS,
	"KHR": KHR,
	"KMF": KMF,
	"KPW": KPW,
	"KRW": KRW,
	"KWD": KWD,
	"KYD": KYD,
	"KZT": KZT,
	"LAK": LAK,
	"LBP": LBP,
	"LKR": LKR,
	"LRD": LRD,
	"LSL": LSL,
	"LTL": LTL,
	"LVL": LVL,
	"LYD": LYD,
	"MAD": MAD,
	"MDL": MDL,
	"MGA": MGA,
	"MKD": MKD,
	"MMK": MMK,
	"MNT": MNT,
	"MOP": MOP,
	"MRO": MRO,
	"MUR": MUR,
	"MVR": MVR,
	"MWK": MWK,
	"MXN": MXN,
	"MYR": MYR,
	"MZM": MZM,
	"NAD": NAD,
	"NGN": NGN,
	"NIO": NIO,
	"NOK": NOK,
	"NPR": NPR,
	"NZD": NZD,
	"OMR": OMR,
	"PAB": PAB,
	"PEN": PEN,
	"PGK": PGK,
	"PHP": PHP,
	"PKR": PKR,
	"PLN": PLN,
	"PYG": PYG,
	"QAR": QAR,
	"RON": RON,
	"RSD": RSD,
	"RUB": RUB,
	"SAR": SAR,
	"SBD": SBD,
	"SCR": SCR,
	"SDG": SDG,
	"SEK": SEK,
	"SGD": SGD,
	"SHP": SHP,
	"SLL": SLL,
	"SOS": SOS,
	"SRD": SRD,
	"SYP": SYP,
	"SZL": SZL,
	"THB": THB,
	"TJS": TJS,
	"TMT": TMT,
	"TND": TND,
	"TRY": TRY,
	"TTD": TTD,
	"TWD": TWD,
	"TZS": TZS,
	"UAH": UAH,
	"UGX": UGX,
	"USD": USD,
	"UYU": UYU,
	"UZS": UZS,
	"VEB": VEB,
	"VND": VND,
	"VUV": VUV,
	"WST": WST,
	"XAF": XAF,
	"XCD": XCD,
	"XDR": XDR,
	"XOF": XOF,
	"XPF": XPF,
	"YER": YER,
	"ZAR": ZAR,
	"ZMK": ZMK,
	"ZWR": ZWR,
}

// Parse code from string.
func Parse(txt string) (Code, bool) {
	code, ok := currencies[txt]
	return code, ok
}
