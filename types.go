package generic

type orderable interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

type numeric interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}
