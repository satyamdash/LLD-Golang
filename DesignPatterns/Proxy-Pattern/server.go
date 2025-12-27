package proxypattern

type server interface {
	handlerequest(string, string) (int, string)
}
