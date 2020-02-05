package pollbot

import (
	"github.com/keybase/managed-bots/base"
)

type Vote struct {
	ID     string
	Choice int
}

type voteToEncode struct {
	ID     string `codec:"d"`
	Choice int    `codec:"i"`
}

func NewVote(id string, choice int) Vote {
	return Vote{
		ID:     id,
		Choice: choice,
	}
}

func NewVoteFromEncoded(sdat string) Vote {
	var ve voteToEncode
	dat, _ := base.URLEncoder().DecodeString(sdat)
	_ = base.MsgpackDecode(&ve, dat)
	return Vote{
		ID:     ve.ID,
		Choice: ve.Choice,
	}
}

func (v Vote) Encode() string {
	mdat, _ := base.MsgpackEncode(voteToEncode{
		ID:     v.ID,
		Choice: v.Choice,
	})
	return base.URLEncoder().EncodeToString(mdat)
}
