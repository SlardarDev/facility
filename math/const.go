package math

type Word uintptr

const (
	// Compute the size _S of a Word in bytes.
	_m    = ^Word(0)
	_logS = _m>>8&1 + _m>>16&1 + _m>>32&1
	_S    = 1 << _logS

	_W = _S << 3 // word size in bits
	_B = 1 << _W // digit base
	_M = _B - 1  // digit mask

	_W2 = _W / 2   // half word size in bits
	_B2 = 1 << _W2 // half digit base
	_M2 = _B2 - 1  // half digit mask
)

const (
	LimitsMaxInt64  int64  = 1<<63 - 1
	LimitsMaxInt32  int32  = 1<<31 - 1
	LimitsMaxInt8   int8   = 1<<7 - 1
	LimitsMaxInt16  int16  = 1<<15 - 1
	LimitsMaxInt    int    = 1<<(_W-1) - 1
	LimitsMaxUint8  uint8  = 1<<8 - 1
	LimitsMaxUint32 uint32 = 1<<32 - 1
	LimitsMaxUint64 uint64 = 1<<64 - 1
	LimitsMaxUint16 uint16 = 1<<16 - 1
	LimitsMinInt64  int64  = -1 << 63
	LimitsMinInt32  int32  = -1 << 31
	LimitsMinInt8   int8   = -1 << 7
	LimitsMinInt16  int16  = -1 << 15
	LimitsMinInt    int    = -1 << (_W - 1)

	LimitsMinUint8  uint8  = 0
	LimitsMinUint32 uint32 = 0
	LimitsMinUint64 uint64 = 0
	LimitsMinUint16 uint16 = 0

	LimitsMaxFloat32             float32 = 3.40282346638528859811704183484516925440e+38  // 2**127 * (2**24 - 1) / 2**23
	LimitsMinFloat32             float32 = -3.40282346638528859811704183484516925440e+38 // 2**127 * (2**24 - 1) / 2**23
	LimitsSmallestNonzeroFloat32 float32 = 1.401298464324817070923729583289916131280e-45 // 1 / 2**(127 - 1 + 23)

	LimitsMaxFloat64             float64 = 1.797693134862315708145274237317043567981e+308  // 2**1023 * (2**53 - 1) / 2**52
	LimitsMinFloat64             float64 = -1.797693134862315708145274237317043567981e+308 // 2**1023 * (2**53 - 1) / 2**52
	LimitsSmallestNonzeroFloat64 float64 = 4.940656458412465441765687928682213723651e-324  // 1 / 2**(1023 - 1 + 52)
)
