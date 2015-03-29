all:: tbpromote tbadd tbadd_ring tbsub tbstrcnd tbstrval tbarbit tbrdy tbconst tbconst_local tbiterator 

tbpromote:
	go run tbpromote.go

tbadd:
	go run tbadd.go

tbadd_ring:
	go run tbadd_ring.go

tbsub:
	go run tbsub.go

tbstrcnd:
	go run tbstrcnd.go

tbstrval:
	go run tbstrval.go

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

