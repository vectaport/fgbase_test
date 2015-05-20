all:: tbpromote tbadd tbadd_ring tbsub tbsteerc tbsteerv tbarbit tbrdy tbconst tbconst_local tbiterator tbfft tbffti tbsrcdst tbfanout tbgcd tbmul tbdiv tblsh tbrsh tbqsort tbread tbwrite tbdisplay tbcapture

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
	go run tbsrc.go -nodeid=1000 &
	go run tbdst.go -nodeid=2000

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

