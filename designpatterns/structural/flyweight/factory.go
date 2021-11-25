package flyweight

func getTeam(team int) *Team {
	switch team {
	case 1:
		return &Team{
			Id:   1,
			Name: "idToTeams 1",
		}
	case 2:
		return &Team{
			Id:   2,
			Name: "idToTeams 2",
		}
	default:
		return nil
	}
}

type Teams struct {
	idToTeams map[int]*Team
}

func (t *Teams) GetTeam(teamId int) *Team {
	if t.idToTeams[teamId] != nil {
		return t.idToTeams[teamId]
	}
	team := getTeam(teamId)
	t.idToTeams[teamId] = getTeam(teamId)
	return team
}

func (t *Teams) GetObjectCount() int {
	return len(t.idToTeams)
}
