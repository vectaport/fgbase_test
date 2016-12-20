all:: tbpromote tbadd tbadd_ring tbsub tbsteerc tbsteerv tbarbit tbrdy tbconst tbconst_local tbiterator tbsrcdst tbfanout tbgcd tbmul tbdiv tblsh tbrsh tbqsort tbread tbwrite tbsrcdst2 tbmap tbreduce tbcollect tbcrypt

imglab:: tbfft tbffti tbdisplay tbcapture

weblab:: tbserver

regexp:: tbmatch tbbar tbstar tbrepeat tbdot

tbpromote:
	go run tbpromote.go

tbadd:
	go run tbadd.go

tbadd_ring:
	go run tbadd_ring.go

tbsub:
	go run tbsub.go

tbsteerc:
	go run tbsteerc.go

tbsteerv:
	go run tbsteerv.go

tbarbit:
	go run tbarbit.go

tbrdy:
	go run tbrdy.go

tbconst:
	go run tbconst.go

tbconst_local:
	go run tbconst_local.go

tbiterator:
	go run tbiterator.go

tbfft:
	go run tbfft.go

tbffti:
	go run tbffti.go

tbsrcdst:
	go run tbsrc.go -nodeid=2000  &
	go run tbdst.go -nodeid=1000 

tbfanout:
	go run tbfanout.go

tbgcd:
	go run tbgcd.go

tbmul:
	go run tbmul.go

tbdiv:
	go run tbdiv.go

tblsh:
	go run tblsh.go

tbrsh:
	go run tbrsh.go

tbqsort:
	go run tbqsort.go

tbread:
	go run tbread.go Makefile

tbwrite:
	go run tbwrite.go testout

tbdisplay:
	go run tbdisplay.go -test

tbcapture:
	go run tbcapture.go -test

tbserver:
	go run tbserver.go -test &
	go run tbclient.go -test

tbsrcdst2:
	go run tbsrc2.go -nodeid=2000 -chansz 16 -trace V &
	go run tbdst2.go -nodeid=1000 -chansz 16 -trace V 

tbmap: 
	go run tbmap.go -chansz 32768 -nmap 6  -nreduce 26 -sec 4 -ncore 3 -trace Q

tbreduce:
	go run tbreduce.go

tbcollect:
	go run tbcollect.go

tbcrypt:
	go run tbcrypt.go

tbmatch:
	go run tbmatch.go

tbbar:
	go run tbbar.go

tbstar:
	go run tbstar.go

tbrepeat:
	go run tbrepeat.go

tbdot:
	go run tbdot.go
