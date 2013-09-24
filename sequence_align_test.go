package seq

import (
	"fmt"
	"strings"
	"testing"
)

func TestNeedlemanWunsch(t *testing.T) {
	type test struct {
		seq1, seq2 string
		out1, out2 string
		subst      MatLookup
	}

	tests := []test{
		{
			"ABCD",
			"ABCD",
			"ABCD",
			"ABCD",
			MatBlosum62,
		},
		{
			"GHIKLMNPQR",
			"GAAAHIKLMN",
			"---GHIKLMNPQR",
			"GAAAHIKLMN---",
			MatBlosum62,
		},
		{
			"GHIKLMNPQRSTVW",
			"GAAAHIKLMNPQRSTVW",
			"---GHIKLMNPQRSTVW",
			"GAAAHIKLMNPQRSTVW",
			MatBlosum62,
		},
		{
			"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			MatBlosum62,
		},
		{
			"NNNNNNNN",
			"NNNNNNNN",
			"NNNNNNNN",
			"NNNNNNNN",
			MatBlosum62,
		},
		{
			"NNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN",
			"NNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN",
			"NNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN",
			"NNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN",
			MatBlosum62,
		},
		{
			"ABCDEFGWXYZ",
			"ABCDEFMNPQRSTZABEGWXYZ",
			"ABCDEF-----------GWXYZ",
			"ABCDEFMNPQRSTZABEGWXYZ",
			MatBlosum62,
		},
		{
			"ASAECVSNENVEIEAPKTNIWTSLAKEEVQEVLDLLHSTYNITEVTKADFFSNYVLWIETLKPN" +
				"KTEALTYLDEDGDLPPRNARTVVYFGEGEEGYFEELKVGPLPVSDETTIEPLSFYNTNGK" +
				"SKLPFEVGHLDRIKSAAKSSFLNKNLNTTIMRDVLEGLIGVPYEDMGCHSAAPQLHDPAT" +
				"GATVDYGTCNINTENDAENLVPTGFFFKFDMTGRDVSQWKMLEYIYNNKVYTSAEELYEA" +
				"MQKDDFVTLPKIDVDNLDWTVIQRNDSAPVRHLDDRKSPRLVEPEGRRWAYDGDEEYFSW" +
				"MDWGFYTSWSRDTGISFYDITFKGERIVYELSLQELIAEYGSDDPFNQHTFYSDISYGVG" +
				"NRFSLVPGYDCPSTAGYFTTDTFEYDEFYNRTLSYCVFENQEDYSLLRHTGASYSAITQN" +
				"PTLNVRFISTIGNYDYNFLYKFFLDGTLEVSVRAAGYIQAGYWNPETSAPYGLKIHDVLS" +
				"GSFHDHVLNYKVDLDVGGTKNRASQYVMKDVDVEYPWAPGTVYNTKQIAREVFENEDFNG" +
				"INWPENGQGILLIESAEETNSFGNPRAYNIMPGGGGVHRIVKNSRSGPETQNWARSNLFL" +
				"TKHKDTELRSSTALNTNALYDPPVNFNAFLDDESLDGEDIVAWVNLGLHHLPNSNDLPNT" +
				"IFSTAHASFMLTPFNYFDSENSRDTTQQVFYTYDDETEESNWEFYGNDWSSCGVEVAEPN" +
				"FEDYTYGRGTRINKKMTNSDEVY",
			"AECVSNENVEIEAPKTNIWTSLAKEEVQEVLDLLHSTYNITEVTKADFFSNYVLWIETLKPNKT" +
				"EALTYLDEDGDLPPRNARTVVYFGEGEEGYFEELKVGPLPVSDETTIEPLSFYNTNGKSK" +
				"LPFEVGHLDRIKSAAKSSFLNKNLNTTIMRDVLEGLIGVPYEDMGCHSAAPQLHDPATGA" +
				"TVDYGTCNINTENDAENLVPTGFFFKFDMTGRDVSQWKMLEYIYNNKVYTSAEELYEAMQ" +
				"KDDFVTLPKIDVDNLDWTVIQRNDSAPVRHLDDRKSPRLVEPEGRRWAYDGDEEYFSWMD" +
				"WGFYTSWSRDTGISFYDITFKGERIVYELSLQELIAEYGSDDPFNQHTFYSDISYGVGNR" +
				"FSLVPGYDCPSTAGYFTTDTFEYDEFYNRTLSYCVFENQEDYSLLRHTGASYSAITQNPT" +
				"LNVRFISTIGNDYNFLYKFFLDGTLEVSVRAAGYIQAGYWNPETSAPYGLKIHDVLSGSF" +
				"HDHVLNYKVDLDVGGTKNRASQYVMKDVDVEYPWAPGTVYNTKQIAREVFENEDFNGINW" +
				"PENGQGILLIESAEETNSFGNPRAYNIMPGGGGVHRIVKNSRSGPETQNWARSNLFLTKH" +
				"KDTELRSSTALNTNALYDPPVNFNAFLDDESLDGEDIVAWVNLGLHHLPNSNDLPNTIFS" +
				"TAHASFMLTPFNYFDSENSRDTTQQVFYTYDDETEESNWEFYGNDWSSCGVEVAEPNFED" +
				"YTYGRGTRINKK",
			"ASAECVSNENVEIEAPKTNIWTSLAKEEVQEVLDLLHSTYNITEVTKADFFSNYVLWIETLKPN" +
				"KTEALTYLDEDGDLPPRNARTVVYFGEGEEGYFEELKVGPLPVSDETTIEPLSFYNTNGK" +
				"SKLPFEVGHLDRIKSAAKSSFLNKNLNTTIMRDVLEGLIGVPYEDMGCHSAAPQLHDPAT" +
				"GATVDYGTCNINTENDAENLVPTGFFFKFDMTGRDVSQWKMLEYIYNNKVYTSAEELYEA" +
				"MQKDDFVTLPKIDVDNLDWTVIQRNDSAPVRHLDDRKSPRLVEPEGRRWAYDGDEEYFSW" +
				"MDWGFYTSWSRDTGISFYDITFKGERIVYELSLQELIAEYGSDDPFNQHTFYSDISYGVG" +
				"NRFSLVPGYDCPSTAGYFTTDTFEYDEFYNRTLSYCVFENQEDYSLLRHTGASYSAITQN" +
				"PTLNVRFISTIGNYDYNFLYKFFLDGTLEVSVRAAGYIQAGYWNPETSAPYGLKIHDVLS" +
				"GSFHDHVLNYKVDLDVGGTKNRASQYVMKDVDVEYPWAPGTVYNTKQIAREVFENEDFNG" +
				"INWPENGQGILLIESAEETNSFGNPRAYNIMPGGGGVHRIVKNSRSGPETQNWARSNLFL" +
				"TKHKDTELRSSTALNTNALYDPPVNFNAFLDDESLDGEDIVAWVNLGLHHLPNSNDLPNT" +
				"IFSTAHASFMLTPFNYFDSENSRDTTQQVFYTYDDETEESNWEFYGNDWSSCGVEVAEPN" +
				"FEDYTYGRGTRINKKMTNSDEVY",
			"--AECVSNENVEIEAPKTNIWTSLAKEEVQEVLDLLHSTYNITEVTKADFFSNYVLWIETLKPN" +
				"KTEALTYLDEDGDLPPRNARTVVYFGEGEEGYFEELKVGPLPVSDETTIEPLSFYNTNGK" +
				"SKLPFEVGHLDRIKSAAKSSFLNKNLNTTIMRDVLEGLIGVPYEDMGCHSAAPQLHDPAT" +
				"GATVDYGTCNINTENDAENLVPTGFFFKFDMTGRDVSQWKMLEYIYNNKVYTSAEELYEA" +
				"MQKDDFVTLPKIDVDNLDWTVIQRNDSAPVRHLDDRKSPRLVEPEGRRWAYDGDEEYFSW" +
				"MDWGFYTSWSRDTGISFYDITFKGERIVYELSLQELIAEYGSDDPFNQHTFYSDISYGVG" +
				"NRFSLVPGYDCPSTAGYFTTDTFEYDEFYNRTLSYCVFENQEDYSLLRHTGASYSAITQN" +
				"PTLNVRFISTIGN-DYNFLYKFFLDGTLEVSVRAAGYIQAGYWNPETSAPYGLKIHDVLS" +
				"GSFHDHVLNYKVDLDVGGTKNRASQYVMKDVDVEYPWAPGTVYNTKQIAREVFENEDFNG" +
				"INWPENGQGILLIESAEETNSFGNPRAYNIMPGGGGVHRIVKNSRSGPETQNWARSNLFL" +
				"TKHKDTELRSSTALNTNALYDPPVNFNAFLDDESLDGEDIVAWVNLGLHHLPNSNDLPNT" +
				"IFSTAHASFMLTPFNYFDSENSRDTTQQVFYTYDDETEESNWEFYGNDWSSCGVEVAEPN" +
				"FEDYTYGRGTRINKK--------",
			MatBlosum62,
		},
		{
			"ACTG",
			"ACTG",
			"ACTG",
			"ACTG",
			MatDNA,
		},
	}
	sep := strings.Repeat("-", 45)
	for _, test := range tests {
		s1, s2 := stringToSeq(test.seq1), stringToSeq(test.seq2)
		substitution := test.subst
		aligned := NeedlemanWunsch(s1, s2, substitution)
		sout1 := fmt.Sprintf("%s", aligned.A)
		sout2 := fmt.Sprintf("%s", aligned.B)

		if sout1 != test.out1 || sout2 != test.out2 {
			t.Fatalf(
				`Alignment for:
%s
%s
%s
%s
resulted in
%s
%s
%s
%s
but should have been
%s
%s
%s
%s`,
				sep, test.seq1, test.seq2, sep,
				sep, sout1, sout2, sep,
				sep, test.out1, test.out2, sep)
		}
	}
}

func stringToSeq(s string) []Residue {
	residues := make([]Residue, len(s))
	for i := range s {
		residues[i] = Residue(s[i])
	}
	return residues
}
