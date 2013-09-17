package seq

type MatLookup func(Residue, Residue) int

func getBlosum62(a, b Residue) int {
	ai, ok := alpha62map[a]
	if !ok {
		ai = alpha62map['X']
	}
	bi, ok := alpha62map[b]
	if !ok {
		bi = alpha62map['X']
	}
	return blosum62[ai][bi]
}

func getDNA(a, b Residue) int {
	ai, ok := alphaDNAmap[a]
	if !ok {
		ai = alphaDNAmap['N']
	}
	bi, ok := alphaDNAmap[b]
	if !ok {
		bi = alphaDNAmap['N']
	}
	return identity[ai][bi]
}

func getRNA(a, b Residue) int {
	ai, ok := alphaRNAmap[a]
	if !ok {
		ai = alphaDNAmap['N']
	}
	bi, ok := alphaRNAmap[b]
	if !ok {
		bi = alphaRNAmap['N']
	}
	return identity[ai][bi]
}

var alpha62map = make(map[Residue]int)

var alphaDNAmap = make(map[Residue]int)

var alphaRNAmap = make(map[Residue]int)

func init() {
	for i, r := range AlphaBlosum62 {
		alpha62map[r] = i
	}
}

func init() {
	for i,r := range AlphaDNA {
		alphaDNAmap[r] = i
	}
}

func init() {
	for i,r := range AlphaRNA {
		alphaRNAmap[r] = i
	}
}

type substmatrix [][]int

var identity = substmatrix{
	{
		2, -1, -1, -1, 0, 0, 
	},
	{
		-1, 2, -1, -1, 0, 0, 
	},
	{
		-1, -1, 2, -1, 0, 0, 
	},
	{
		-1, -1, -1, 2, 0, 0, 
	},
	{
		0, 0, 0, 0, 0, 0,
	},
	{
		0, 0, 0, 0, 0, 0,
	},
}

// The indices correspond to alpha62, and can be retrieved using the
// alpha62map. The last element in each column/row is the gap penalty.
var blosum62 = substmatrix{
	{
		4.0, -2.0, 0.0, -2.0, -1.0, -2.0, 0.0, -2.0, -1.0, -1.0, -1.0, -1.0,
		-2.0, -1.0, -1.0, -1.0, 1.0, 0.0, 0.0, -3.0, 0.0, -2.0, -1.0, -4.0,
	},
	{
		-2.0, 4.0, -3.0, 4.0, 1.0, -3.0, -1.0, 0.0, -3.0, 0.0, -4.0, -3.0, 3.0,
		-2.0, 0.0, -1.0, 0.0, -1.0, -3.0, -4.0, -1.0, -3.0, 1.0, -4.0,
	},
	{
		0.0, -3.0, 9.0, -3.0, -4.0, -2.0, -3.0, -3.0, -1.0, -3.0, -1.0, -1.0,
		-3.0, -3.0, -3.0, -3.0, -1.0, -1.0, -1.0, -2.0, -2.0, -2.0, -3.0, -4.0,
	},
	{
		-2.0, 4.0, -3.0, 6.0, 2.0, -3.0, -1.0, -1.0, -3.0, -1.0, -4.0, -3.0,
		1.0, -1.0, 0.0, -2.0, 0.0, -1.0, -3.0, -4.0, -1.0, -3.0, 1.0, -4.0,
	},
	{
		-1.0, 1.0, -4.0, 2.0, 5.0, -3.0, -2.0, 0.0, -3.0, 1.0, -3.0, -2.0, 0.0,
		-1.0, 2.0, 0.0, 0.0, -1.0, -2.0, -3.0, -1.0, -2.0, 4.0, -4.0,
	},
	{
		-2.0, -3.0, -2.0, -3.0, -3.0, 6.0, -3.0, -1.0, 0.0, -3.0, 0.0, 0.0,
		-3.0, -4.0, -3.0, -3.0, -2.0, -2.0, -1.0, 1.0, -1.0, 3.0, -3.0, -4.0,
	},
	{
		0.0, -1.0, -3.0, -1.0, -2.0, -3.0, 6.0, -2.0, -4.0, -2.0, -4.0, -3.0,
		0.0, -2.0, -2.0, -2.0, 0.0, -2.0, -3.0, -2.0, -1.0, -3.0, -2.0, -4.0,
	},
	{
		-2.0, 0.0, -3.0, -1.0, 0.0, -1.0, -2.0, 8.0, -3.0, -1.0, -3.0, -2.0,
		1.0, -2.0, 0.0, 0.0, -1.0, -2.0, -3.0, -2.0, -1.0, 2.0, 0.0, -4.0,
	},
	{
		-1.0, -3.0, -1.0, -3.0, -3.0, 0.0, -4.0, -3.0, 4.0, -3.0, 2.0, 1.0,
		-3.0, -3.0, -3.0, -3.0, -2.0, -1.0, 3.0, -3.0, -1.0, -1.0, -3.0, -4.0,
	},
	{
		-1.0, 0.0, -3.0, -1.0, 1.0, -3.0, -2.0, -1.0, -3.0, 5.0, -2.0, -1.0,
		0.0, -1.0, 1.0, 2.0, 0.0, -1.0, -2.0, -3.0, -1.0, -2.0, 1.0, -4.0,
	},
	{
		-1.0, -4.0, -1.0, -4.0, -3.0, 0.0, -4.0, -3.0, 2.0, -2.0, 4.0, 2.0,
		-3.0, -3.0, -2.0, -2.0, -2.0, -1.0, 1.0, -2.0, -1.0, -1.0, -3.0, -4.0,
	},
	{
		-1.0, -3.0, -1.0, -3.0, -2.0, 0.0, -3.0, -2.0, 1.0, -1.0, 2.0, 5.0,
		-2.0, -2.0, 0.0, -1.0, -1.0, -1.0, 1.0, -1.0, -1.0, -1.0, -1.0, -4.0,
	},
	{
		-2.0, 3.0, -3.0, 1.0, 0.0, -3.0, 0.0, 1.0, -3.0, 0.0, -3.0, -2.0, 6.0,
		-2.0, 0.0, 0.0, 1.0, 0.0, -3.0, -4.0, -1.0, -2.0, 0.0, -4.0,
	},
	{
		-1.0, -2.0, -3.0, -1.0, -1.0, -4.0, -2.0, -2.0, -3.0, -1.0, -3.0, -2.0,
		-2.0, 7.0, -1.0, -2.0, -1.0, -1.0, -2.0, -4.0, -2.0, -3.0, -1.0, -4.0,
	},
	{
		-1.0, 0.0, -3.0, 0.0, 2.0, -3.0, -2.0, 0.0, -3.0, 1.0, -2.0, 0.0, 0.0,
		-1.0, 5.0, 1.0, 0.0, -1.0, -2.0, -2.0, -1.0, -1.0, 3.0, -4.0,
	},
	{
		-1.0, -1.0, -3.0, -2.0, 0.0, -3.0, -2.0, 0.0, -3.0, 2.0, -2.0, -1.0,
		0.0, -2.0, 1.0, 5.0, -1.0, -1.0, -3.0, -3.0, -1.0, -2.0, 0.0, -4.0,
	},
	{
		1.0, 0.0, -1.0, 0.0, 0.0, -2.0, 0.0, -1.0, -2.0, 0.0, -2.0, -1.0, 1.0,
		-1.0, 0.0, -1.0, 4.0, 1.0, -2.0, -3.0, 0.0, -2.0, 0.0, -4.0,
	},
	{
		0.0, -1.0, -1.0, -1.0, -1.0, -2.0, -2.0, -2.0, -1.0, -1.0, -1.0, -1.0,
		0.0, -1.0, -1.0, -1.0, 1.0, 5.0, 0.0, -2.0, 0.0, -2.0, -1.0, -4.0,
	},
	{
		0.0, -3.0, -1.0, -3.0, -2.0, -1.0, -3.0, -3.0, 3.0, -2.0, 1.0, 1.0,
		-3.0, -2.0, -2.0, -3.0, -2.0, 0.0, 4.0, -3.0, -1.0, -1.0, -2.0, -4.0,
	},
	{
		-3.0, -4.0, -2.0, -4.0, -3.0, 1.0, -2.0, -2.0, -3.0, -3.0, -2.0, -1.0,
		-4.0, -4.0, -2.0, -3.0, -3.0, -2.0, -3.0, 11.0, -2.0, 2.0, -3.0, -4.0,
	},
	{
		0.0, -1.0, -2.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0,
		-1.0, -2.0, -1.0, -1.0, 0.0, 0.0, -1.0, -2.0, -1.0, -1.0, -1.0, -4.0,
	},
	{
		-2.0, -3.0, -2.0, -3.0, -2.0, 3.0, -3.0, 2.0, -1.0, -2.0, -1.0, -1.0,
		-2.0, -3.0, -1.0, -2.0, -2.0, -2.0, -1.0, 2.0, -1.0, 7.0, -2.0, -4.0,
	},
	{
		-1.0, 1.0, -3.0, 1.0, 4.0, -3.0, -2.0, 0.0, -3.0, 1.0, -3.0, -1.0, 0.0,
		-1.0, 3.0, 0.0, 0.0, -1.0, -2.0, -3.0, -1.0, -2.0, 4.0, -4.0,
	},
	{
		-4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0,
		-4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, -4.0, 0.0,
	},
}
