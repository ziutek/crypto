package xxtea

const delta uint32 = 0x9e3779b9

func mx(e, p int, y, z, sum uint32, key [4]uint32) uint32 {
	return ((z>>5 ^ y<<2) + (y>>3 ^ z<<4)) ^ ((sum ^ y) + (key[(p&3)^e] ^ z))
}

func Encrypt(v []uint32, key [4]uint32) {
	var y, sum uint32
	z := v[len(v)-1]
	for rounds := 6 + 52/len(v); rounds != 0; rounds-- {
		sum += delta
		e := int((sum >> 2) & 3)
		for p := range v[:len(v)-1] {
			y = v[p+1]
			v[p] += mx(e, p, y, z, sum, key)
			z = v[p]
			p++
		}
		y = v[0]
		v[len(v)-1] += mx(e,len(v)-1, y, z, sum, key)
		z = v[len(v)-1]
	}
}

func Decrypt(v []uint32, key [4]uint32) {
	var z uint32
	rounds := 6 + 52/len(v)
	y := v[0]
	for sum := uint32(rounds) * delta; sum != 0; sum -= delta {
		e := int((sum >> 2) & 3)
		for p := len(v) - 1; p > 0; p-- {
			z = v[p-1]
			v[p] -= mx(e, p, y, z, sum, key)
			y = v[p]
		}
		z = v[len(v)-1]
		v[0] -= mx(e, 0, y, z, sum, key)
		y = v[0]
	}

}
