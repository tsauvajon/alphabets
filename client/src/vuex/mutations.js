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

const SET_RESULTS = (state, { results }) => {
  state.results = results

  state.results.forEach(result => {
    ADD_TEAM(state, { team: result.LocalTeam })
    ADD_TEAM(state, { team: result.VisitorTeam })
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

const SET_STANDINGS = (state, { standings }) => {
  state.standings = standings
}

const mutations = {
  TOGGLE_DRAWER,
  SET_FIXTURES,
  SET_RESULTS,
  ADD_TEAM,
  SET_STANDINGS
}

export default mutations
