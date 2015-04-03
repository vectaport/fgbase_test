all:: tbpromote tbadd tbadd_ring tbsub tbsteerc tbsteerv tbarbit tbrdy tbconst tbconst_local tbiterator 

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

