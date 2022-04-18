package nutshttp

import (
	"github.com/gin-gonic/gin"
)

type BaseUri struct {
	Bucket string `uri:"bucket" binding:"required"`
	Key    string `uri:"key" binding:"required"`
}

func (s *NutsHTTPServer) SAdd(c *gin.Context) {

	type AddSetRequest struct {
		Value []string `json:"value" binding:"required"`
	}

	var (
		err           error
		baseUri       BaseUri
		addSetRequest AddSetRequest
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&addSetRequest); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = s.core.addSet(baseUri.Bucket, baseUri.Key, addSetRequest.Value...); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	WriteSucc(c, struct{}{})
}

func (s *NutsHTTPServer) SMembers(c *gin.Context) {

	type listSetResp struct {
		Items []string `json:"items"`
	}

	var (
		err     error
		baseUri BaseUri
		resp    listSetResp
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if resp.Items, err = s.core.listSet(baseUri.Bucket, baseUri.Key); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	WriteSucc(c, resp)
}

func (s *NutsHTTPServer) SAreMembers(c *gin.Context) {

	type (
		SAreMembersRequest struct {
			Value []string `json:"value" binding:"required"`
		}

		SAreMembersResp struct {
			IsExist bool `json:"is_exist"`
		}
	)

	var (
		ok               bool
		err              error
		baseUri          BaseUri
		sAreMembersquest SAreMembersRequest
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&sAreMembersquest); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if ok, err = s.core.sAreMembers(baseUri.Bucket, baseUri.Key, sAreMembersquest.Value...); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	WriteSucc(c, SAreMembersResp{
		IsExist: ok,
	})
}

func (s *NutsHTTPServer) SIsMember(c *gin.Context) {

	type (
		SIsMemberRequest struct {
			Value string `json:"value" binding:"required"`
		}

		SIsMemberResp struct {
			IsExist bool `json:"is_exist"`
		}
	)

	var (
		ok              bool
		err             error
		baseUri         BaseUri
		sIsMembersquest SIsMemberRequest
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})

		return
	}

	if err = c.ShouldBindJSON(&sIsMembersquest); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})

		return
	}

	if ok, err = s.core.sIsMember(baseUri.Bucket, baseUri.Key, sIsMembersquest.Value); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})

		return
	}

	WriteSucc(c, SIsMemberResp{
		IsExist: ok,
	})
}

func (s *NutsHTTPServer) SCard(c *gin.Context) {

	type ScardResp struct {
		Num int `json:"num"`
	}

	var (
		err     error
		baseUri BaseUri
		resp    ScardResp
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if resp.Num, err = s.core.sCard(baseUri.Bucket, baseUri.Key); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	WriteSucc(c, resp)
}

func (s *NutsHTTPServer) SHasKey(c *gin.Context) {

	type SHasKeyResp struct {
		IsExist bool `json:"is_exist"`
	}

	var (
		ok      bool
		err     error
		baseUri BaseUri
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})

		return
	}

	if ok, err = s.core.sHasKey(baseUri.Bucket, baseUri.Key); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})

		return
	}

	WriteSucc(c, SHasKeyResp{
		IsExist: ok,
	})
}

func (s *NutsHTTPServer) Spop(c *gin.Context) {

	type sPopResp struct {
		Item string `json:"item"`
	}

	var (
		err     error
		baseUri BaseUri
		resp    sPopResp
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})

		return
	}

	if resp.Item, err = s.core.sPop(baseUri.Bucket, baseUri.Key); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	WriteSucc(c, resp)
}

func (s *NutsHTTPServer) Srem(c *gin.Context) {

	type SRemRequest struct {
		Value []string `json:"value" binding:"required"`
	}

	var (
		err         error
		baseUri     BaseUri
		sRemRequest SRemRequest
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&sRemRequest); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = s.core.sRem(baseUri.Bucket, baseUri.Key, sRemRequest.Value); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	WriteSucc(c, struct{}{})
}

func (s *NutsHTTPServer) SDiffByOneBucket(c *gin.Context) {

	type (
		SDiffByOneBucketRequest struct {
			Key string `json:"key" binding:"required"`
		}

		SDiffByOneBucketResp struct {
			Items []string `json:"items"`
		}
	)

	var (
		err                     error
		baseUri                 BaseUri
		sDiffByOneBucketRequest SDiffByOneBucketRequest
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&sDiffByOneBucketRequest); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	items, err := s.core.sDiffByOneBucket(baseUri.Bucket, baseUri.Key, sDiffByOneBucketRequest.Key)
	if err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	WriteSucc(c, SDiffByOneBucketResp{
		Items: items,
	})
}

func (s *NutsHTTPServer) SDiffByTwoBuckets(c *gin.Context) {

	type (
		SDiffByTwoBucketsRequest struct {
			Bucket string `json:"bucket" binding:"required"`
			Key    string `json:"key" binding:"required"`
		}

		SDiffByTwoBucketsResp struct {
			Items []string `json:"items"`
		}
	)

	var (
		err                      error
		baseUri                  BaseUri
		sDiffByTwoBucketsRequest SDiffByTwoBucketsRequest
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&sDiffByTwoBucketsRequest); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	items, err := s.core.sDiffByTwoBuckets(baseUri.Bucket, sDiffByTwoBucketsRequest.Bucket, baseUri.Key, sDiffByTwoBucketsRequest.Key)
	if err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	WriteSucc(c, SDiffByTwoBucketsResp{
		Items: items,
	})
}

func (s *NutsHTTPServer) SMoveByOneBucket(c *gin.Context) {
	type (
		SMoveByOneBucketRequest struct {
			Key   string `json:"key" binding:"required"`
			Value string `json:"value" binding:"required"`
		}

		SMoveByOneBucketResp struct {
			Ok bool `json:"ok"`
		}
	)

	var (
		err                     error
		baseUri                 BaseUri
		sMoveByOneBucketRequest SMoveByOneBucketRequest
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&sMoveByOneBucketRequest); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	ok, err := s.core.sMoveByOneBucket(baseUri.Bucket, baseUri.Key, sMoveByOneBucketRequest.Key, sMoveByOneBucketRequest.Value)
	if err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	WriteSucc(c, SMoveByOneBucketResp{
		Ok: ok,
	})
}

func (s *NutsHTTPServer) SMoveByTwoBuckets(c *gin.Context) {
	type (
		SMoveByTwoBucketsRequest struct {
			Bucket string `json:"bucket" binding:"required"`
			Key    string `json:"key" binding:"required"`
			Value  string `json:"value" binding:"required"`
		}

		SMoveByTwoBucketsResp struct {
			Ok bool `json:"ok"`
		}
	)

	var (
		err                      error
		baseUri                  BaseUri
		sMoveByTwoBucketsRequest SMoveByTwoBucketsRequest
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&sMoveByTwoBucketsRequest); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	ok, err := s.core.sMoveByTwoBuckets(baseUri.Bucket, baseUri.Key, sMoveByTwoBucketsRequest.Bucket, sMoveByTwoBucketsRequest.Key, sMoveByTwoBucketsRequest.Value)
	if err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	WriteSucc(c, SMoveByTwoBucketsResp{
		Ok: ok,
	})
}

func (s *NutsHTTPServer) SUnionByOneBucket(c *gin.Context) {
	type (
		SUnionByOneBucketRequest struct {
			Key string `json:"key" binding:"required"`
		}

		SUnionByOneBucketResp struct {
			Items []string `json:"items"`
		}
	)

	var (
		err                      error
		baseUri                  BaseUri
		sUnionByOneBucketRequest SUnionByOneBucketRequest
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&sUnionByOneBucketRequest); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	items, err := s.core.sUnionByOneBucket(baseUri.Bucket, baseUri.Key, sUnionByOneBucketRequest.Key)
	if err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	WriteSucc(c, SUnionByOneBucketResp{
		Items: items,
	})
}

func (s *NutsHTTPServer) SUnionByTwoBuckets(c *gin.Context) {
	type (
		SUnionByTwoBucketsRequest struct {
			Bucket string `json:"bucket" binding:"required"`
			Key    string `json:"key" binding:"required"`
		}

		SUnionByTwoBucketsResp struct {
			Items []string `json:"items"`
		}
	)

	var (
		err                       error
		baseUri                   BaseUri
		sUnionByTwoBucketsRequest SUnionByTwoBucketsRequest
	)

	if err = c.ShouldBindUri(&baseUri); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&sUnionByTwoBucketsRequest); err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	items, err := s.core.sUnionByTwoBuckets(baseUri.Bucket, baseUri.Key, sUnionByTwoBucketsRequest.Bucket, sUnionByTwoBucketsRequest.Key)
	if err != nil {
		WriteError(c, APIMessage{
			Message: err.Error(),
		})
		return
	}

	WriteSucc(c, SUnionByTwoBucketsResp{
		Items: items,
	})
}
