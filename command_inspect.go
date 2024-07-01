package main

import (
	"errors"
	"fmt"
)

func commandInspect(s *Status) error {
	if s.extraArgs == nil || len(s.extraArgs) > 1 {
		return errors.New("incorrect number of arg: expect 1")
	}
	name := s.extraArgs[0]
	pkm, ok := s.pkmCaught[name]
	if !ok {
		return errors.New("you have not caught " + name)
	}
	fmt.Println("Name:", pkm.Name)
	fmt.Println("Height:", pkm.Height)
	fmt.Println("Weight:", pkm.Weight)
	fmt.Println("Stat:")
	stats := [6]string{"HP:", "Attack:", "Defense:", "Special Attack:", "Speical Defense:", "Speed:"}
	for i := 0; i < 6; i++ {
		fmt.Println(" -", stats[i], pkm.Stats[i])
	}
	fmt.Println("Types:")
	fmt.Println(" -", pkm.Types[0])
	if pkm.Types[1] != "" {
		fmt.Println(" -", pkm.Types[1])
	}
	return nil
}
