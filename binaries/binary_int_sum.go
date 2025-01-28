package binaries

/*
Q 2.1-4: Dado dois vetores de bits A e B representando inteiros, contendo um arranjo de n elementos cada, realize a soma de ambos retornando um vetor C de tamanho n+1

A representação devera ser feita com o menor valor possível (bit/boolean)

Ex:
	A = [1, 1, 0, 1] = 13
	B = [1, 0, 1, 1] = 11
	C = [1, 1, 0, 0, 0] = 24

Por que o uso de booleanos para representação, simples. Binarios são valores de 1 ou 0 bit a bit, portanto não fazia sentido utilizar valores inteiros para cada posição pois seu uso seria mais pesadodo, tavez um int8? (2 bytes)

Nota: formato esperado BigEndian, a implementação não considera o MSB
Melhorias podem ser feitas outro dia, para fins de testes está ótimo
*/
func BinaryIntSum(a []bool, b []bool) []bool {

	r := make([]bool, len(a)+1)
	c := false

	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == b[i] && !c {
			r[i] = false
			c = a[i]
		} else if a[i] == b[i] && c {
			r[i] = a[i] || c
			c = a[i]
		} else if a[i] || a[i] && !c {
			r[i] = true
			c = false
		} else {
			r[i] = c
			c = false
		}
	}

	if c {
		r[0] = true
	}

	return r
}

/*
	Nota: formato esperado BigEndian, a implementação não considera o MSB
*/
func BitArrayToInt64(s []bool) uint64 {

	r := uint64(0)

	for i := 0; i < len(s); i++ {
		if s[i] {
			r = (r + 0x1)
		}
		r = r << 1
	}
	r = r >> 1

	return r
}

func Int64ToBitArray(n uint64) []bool {
	s := make([]bool, 64)

	for i := len(s) - 1; n != 0 && i >= 0; i-- {
		s[i] = n&0x1 == 0x1
		n >>= 1
	}

	return s
}
