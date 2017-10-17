const TOGGLE_DRAWER = (state) => {
  state.drawer = !state.drawer
}

const SET_FIXTURES = (state, { fixtures }) => {
  state.fixtures = fixtures

  state.fixtures.forEach(fixture => {
    ADD_TEAM(state, { team: fixture.LocalTeam })
    ADD_TEAM(state, { team: fixture.VisitorTeam })
  })
}

const ADD_TEAM = (state, { team }) => {
  const index = state.teams.findIndex(t => t.id === team.id)

  if (index !== -1) {
    state.teams[index] = team
  }

  state.teams = [
    ...state.teams,
    team
  ]
}

const mutations = {
  TOGGLE_DRAWER,
  SET_FIXTURES
}

export default mutations
