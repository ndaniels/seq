package seq

var alignFasta = []string{
	"PPDHLWVHQEGIYRDEYQRTWVAVVEE--E--T--SF---------LR----------ARVQQIQVPLG---" +
		"----DAARPSHLLTS-----QL",
	"HPNRLWIWEKHVYLDEFRRSWLPVVIK--S--N--EK---------FQ----------VILRQEDVTLG---" +
		"----EAMSPSQLVPY-----EL",
	"PPRFLVCTRDDIYEDENGRQWVVAKVE--T--S--RSpygsrietcIT----------VHLQHMTTIPQ---" +
		"----EPTPQQPINNN-----SL",
	"HPDRLWAWEKFVYLDEKQHAWLPLTIEikD--R--LQ---------LR----------VLLRREDVVLG---" +
		"----RPMTPTQIGPS-----LL",
	"----------GIYEDEHHRVWIAVNVE--T--S--HS---------SHgnrietcvt-VHLQHMTTLPQ---" +
		"----EPTPQQPINNN-----SL",
	"LPVYLVSVRLGIYEDEHHRVWIVANVE--TshS--SH---------GN----------RRRTHVTVHLW---" +
		"----KLIPQQVIPFNplnydFL",
	"-PDRLWLWEKHVYLDEFRRSWLPIVIK--S--N--GK---------FQ----------VIMRQKDVILG---" +
		"----DSMTPSQLVPY-----EL",
	"-PHILTLRTHGIYEDEHHRLWVVLDLQ--A--ShlSF---------SN----------RLLIYLTVYLQqgv" +
		"afplESTPPSPMNLN-----GL",
	"PPCFLVCTRDDIYEDEHGRQWVAAKVE--T--S--SH---------SPycskietcvtVHLWQMTTLFQ---" +
		"----EPSPDSLKTFN-----FL",
	"---------PGFYEDEHHRLWMVAKLE--T--C--SH---------SPycnkietcvtVHLWQMTRYPQ---" +
		"----EPAPYNPMNYN-----FL",
	"---A--------------------------------------------------------------------" +
		"---------------------L",
}

var alignA2M = []string{
	"PPDHLWVHQEGIYRDEYQRTWVAVVEE..E..T..SF.........LR..........ARVQQIQVPLG..." +
		"....DAARPSHLLTS.....QL",
	"HPNRLWIWEKHVYLDEFRRSWLPVVIK..S..N..EK.........FQ..........VILRQEDVTLG..." +
		"....EAMSPSQLVPY.....EL",
	"PPRFLVCTRDDIYEDENGRQWVVAKVE..T..S..RSpygsrietcIT..........VHLQHMTTIPQ..." +
		"....EPTPQQPINNN.....SL",
	"HPDRLWAWEKFVYLDEKQHAWLPLTIEikD..R..LQ.........LR..........VLLRREDVVLG..." +
		"....RPMTPTQIGPS.....LL",
	"----------GIYEDEHHRVWIAVNVE..T..S..HS.........SHgnrietcvt.VHLQHMTTLPQ..." +
		"....EPTPQQPINNN.....SL",
	"LPVYLVSVRLGIYEDEHHRVWIVANVE..TshS..SH.........GN..........RRRTHVTVHLW..." +
		"....KLIPQQVIPFNplnydFL",
	"-PDRLWLWEKHVYLDEFRRSWLPIVIK..S..N..GK.........FQ..........VIMRQKDVILG..." +
		"....DSMTPSQLVPY.....EL",
	"-PHILTLRTHGIYEDEHHRLWVVLDLQ..A..ShlSF.........SN..........RLLIYLTVYLQqgv" +
		"afplESTPPSPMNLN.....GL",
	"PPCFLVCTRDDIYEDEHGRQWVAAKVE..T..S..SH.........SPycskietcvtVHLWQMTTLFQ..." +
		"....EPSPDSLKTFN.....FL",
	"---------PGFYEDEHHRLWMVAKLE..T..C..SH.........SPycnkietcvtVHLWQMTRYPQ..." +
		"....EPAPYNPMNYN.....FL",
	"---A-----------------------..-..-..--.........--..........-----------..." +
		"....-----------.....-L",
}

var alignA3M = []string{
	"PPDHLWVHQEGIYRDEYQRTWVAVVEEETSFLRARVQQIQVPLGDAARPSHLLTSQL",
	"HPNRLWIWEKHVYLDEFRRSWLPVVIKSNEKFQVILRQEDVTLGEAMSPSQLVPYEL",
	"PPRFLVCTRDDIYEDENGRQWVVAKVETSRSpygsrietcITVHLQHMTTIPQEPTPQQPINNNSL",
	"HPDRLWAWEKFVYLDEKQHAWLPLTIEikDRLQLRVLLRREDVVLGRPMTPTQIGPSLL",
	"----------GIYEDEHHRVWIAVNVETSHSSHgnrietcvtVHLQHMTTLPQEPTPQQPINNNSL",
	"LPVYLVSVRLGIYEDEHHRVWIVANVETshSSHGNRRRTHVTVHLWKLIPQQVIPFNplnydFL",
	"-PDRLWLWEKHVYLDEFRRSWLPIVIKSNGKFQVIMRQKDVILGDSMTPSQLVPYEL",
	"-PHILTLRTHGIYEDEHHRLWVVLDLQAShlSFSNRLLIYLTVYLQqgvafplESTPPSPMNLNGL",
	"PPCFLVCTRDDIYEDEHGRQWVAAKVETSSHSPycskietcvtVHLWQMTTLFQEPSPDSLKTFNFL",
	"---------PGFYEDEHHRLWMVAKLETCSHSPycnkietcvtVHLWQMTRYPQEPAPYNPMNYNFL",
	"---A----------------------------------------------------L",
}

var trickyA3M = []string{
	"GSGMKEFPCWLVEEFVVAEECSPCSNFRAKTTPECGPTGYVEKITCSSSKRNEFKSCRSALMEQR",
	"-----SLHCWQEEEFSILTECARCNPFQMKSWAPCARTGFIENINCAKSNKVEYKSCRSSRMDES",
	"--------CWLQEKYEVIEKCRPCKQFElts-NSVPVCSVTGYVERVSCKSSG-EVYRSCHNMRAEEN",
	"-----NSTCWRHESYTVVQECHPCSEFDivSRSLGVCIHTHYKEVLRC-KSGEIVTKSCDrVALIEQR",
	"-----TasnTTTNQCGEIIEVGECLPCSAFDkkSRAIRACAENGNKQLVKCRKSRQKFYKSCLmVPWIAER",
	"------ssEGC-wEREDYVVKQKCVPCTEFErvC-------aarlvapgq-TRWYEV-----------aan" +
		"kqikdvc---S--cene-NEAQ",
	"----IDLHCWQKEDFEILESCKLCSEFEkeASNINNCTEAsSYYDKVNCTNSG-VAFRPlcqili-cEN--" +
		"----",
	"----SSNKCWLREENTILKECHLCSD-----KLECINASYVETIKCKISG-LAYRNCrKP------",
	"------EECWKHEPFEVVTQCAPCKDFEIKAIkaGHCLKTGYFDRVNCSKSSTTVLRPCPSPK----",
	"------KTCWETENFTFIGPCDPCHLLASKvYTDSCSETGFRQLVQCTDSNQQVYKSCP-------",
}

var trickyA2M = []string{
	"GSGMKE...FPCW.LVEEFVVAEECSPCSNFR...AK.TT..PECG.........PT.GYVEKITCSSSKRN" +
		"EF..........KS......C.R.SA....LMEQR",
	"-----S...LHCW.QEEEFSILTECARCNPFQ...MK.SW..APCA.........RT.GFIENINCAKSNKV" +
		"EY..........KS......C.R.SS....RMDES",
	"------...--CW.LQEKYEVIEKCRPCKQFElts-N.SV..PVCS.........VT.GYVERVSCKSSG-E" +
		"VY..........RS......C.H.NM....RAEEN",
	"-----N...STCW.RHESYTVVQECHPCSEFDiv.SR.SL..GVCI.........HT.HYKEVLRC-KSGEI" +
		"VT..........KS......C.DrVA....LIEQR",
	"-----TasnTTTN.QCGEIIEVGECLPCSAFDkk.SR.AI..RACA.........EN.GNKQLVKCRKSRQK" +
		"FY..........KS......C.LmVP....WIAER",
	"------ss.EGC-wEREDYVVKQKCVPCTEFErv.C-.--..----aarlvapgq-T.RWYEV---------" +
		"--aankqikdvc--......-.S.--cene-NEAQ",
	"----ID...LHCW.QKEDFEILESCKLCSEFEke.AS.NI..NNCT.........EAsSYYDKVNCTNSG-V" +
		"AF..........RPlcqili-cE.N-....-----",
	"----SS...NKCW.LREENTILKECHLCSD--...--.-K..LECI.........NA.SYVETIKCKISG-L" +
		"AY..........RN......CrK.P-....-----",
	"------...EECW.KHEPFEVVTQCAPCKDFE...IK.AIkaGHCL.........KT.GYFDRVNCSKSSTT" +
		"VL..........RP......C.P.SP....K----",
	"------...KTCW.ETENFTFIGPCDPCHLLA...SKvYT..DSCS.........ET.GFRQLVQCTDSNQQ" +
		"VY..........KS......C.P.--....-----",
}

var trickyA3MShort = []string{
	"GSGMKEFPCWLVEEFVVAEECSPCSNFRAKTTPECGPTGYVEKITCSSSKRNEFKSCRSALMEQR",
	"-----SLHCWQEEEFSILTECARCNPFQMKSWAPCARTGFIENINCAKSNKVEYKSCRSSRMDES",
	"--------CWLQEKYEVIEKCRPCKQFElts-NSVPVCSVTGYVERVSCKSSG-EVYRSCHNMRAEEN",
	"-----NSTCWRHESYTVVQECHPCSEFDivSRSLGVCIHTHYKEVLRC-KSGEIVTKSCDrVALIEQR",
}

var trickyA2MShort = []string{
	"GSGMKEFPCWLVEEFVVAEECSPCSNFR...AKTTPECGPTGYVEKITCSSSKRNEFKSCR.SALMEQR",
	"-----SLHCWQEEEFSILTECARCNPFQ...MKSWAPCARTGFIENINCAKSNKVEYKSCR.SSRMDES",
	"--------CWLQEKYEVIEKCRPCKQFElts-NSVPVCSVTGYVERVSCKSSG-EVYRSCH.NMRAEEN",
	"-----NSTCWRHESYTVVQECHPCSEFDiv.SRSLGVCIHTHYKEVLRC-KSGEIVTKSCDrVALIEQR",
}
