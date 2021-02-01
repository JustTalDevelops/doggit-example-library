package main

import (
	"github.com/JustTalDevelops/doggit"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/player/title"
)

type DoggitHandler struct {
	doggit.NopHandler
}

func (d DoggitHandler) HandleJoin(p *player.Player) {
	p.SendTitle(title.New("Example Library").WithSubtitle("Doggit Builder"))
}
