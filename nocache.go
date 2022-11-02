package dnsr

type NoCache struct{}

func (*NoCache) addNX(qname string)      {}
func (*NoCache) add(qname string, rr RR) {}
func (*NoCache) get(qname string) RRs    { return nil }
