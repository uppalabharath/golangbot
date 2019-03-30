1. This repo is a practise repo and follows examples from golangbot.com
2. The types in go are as follows: 
    * bool
        * zero / default value is false
    * Numeric Types
        * int8, int16, int32, int64, int -- Signed integers
        * uint8, uint16, uint32, uint64, uint -- Unsigned integers. 
            * Start with zero
            * 8 reprsnts the bits they occupy
            * int and uint occupy different sizes depending upon themachine the code runs
            * like for Eg., 32bit machines, int is a signed integer and occupies 32 bits
        * float32, float64
        * complex64, complex128
            * 64 means float32 real and imaginary parts
            * usually the real and imaginary part should be of same type ether float32 or float64
        * byte
            * alias to uint8
        * rune
            * alias to int32
            * Used in String manipulation tasks
    * string

3. https://golang.org/pkg/fmt/ -- Reference for ft package