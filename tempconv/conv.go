package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }  // it's not a function, it is force type transformation
func FToC(f Fahrenheit) Celsius { return Celsius(f*9/5 + 32) }
