package nutshttp

func (s *NutsHTTPServer) initSetRouter() {
	sr := s.r.Group("/set")

	sr.POST("/sadd/:bucket/:key", s.SAdd)

	sr.POST("/saremembers/:bucket/:key", s.SAreMembers)

	sr.POST("/sismember/:bucket/:key", s.SIsMember)

	sr.POST("/srem/:bucket/:key", s.Srem)

	sr.GET("/smembers/:bucket/:key", s.SMembers)

	sr.GET("/shaskey/:bucket/:key", s.SHasKey)

	sr.GET("/spop/:bucket/:key", s.Spop)

	sr.GET("/scard/:bucket/:key", s.SCard)

	sr.POST("/sdiffbyonebucket/:bucket/:key", s.SDiffByOneBucket)

	sr.POST("/sdiffbytwobuckets/:bucket/:key", s.SDiffByTwoBuckets)

	sr.POST("/smovebyonebucket/:bucket/:key", s.SMoveByOneBucket)

	sr.POST("/smovebytwobuckets/:bucket/:key", s.SMoveByTwoBuckets)

	sr.POST("/sunionbyonebucket/:bucket/:key", s.SUnionByOneBucket)

	sr.POST("/sunionbytwobuckets/:bucket/:key", s.SUnionByTwoBuckets)
}
