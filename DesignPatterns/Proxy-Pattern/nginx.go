package proxypattern

type nginx struct {
	application       *Application
	maxallowedrequest int
	rateLimiter       map[string]int
}

func newNginx() *nginx {
	return &nginx{
		application:       &Application{},
		maxallowedrequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *nginx) handlerequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.application.handleRequest(url, method)
}

func (n *nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxallowedrequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}
