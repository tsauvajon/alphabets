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

const SET_BETS = (state, { bets }) => {
  state.bets = bets
}

const ADD_MONEY = (state, { amount }) => {
  state.wallet = parseInt(state.wallet) + amount
}

const CONNECT = (state, { userName, userId, wallet }) => {
  state.connected = true
  state.username = userName
  state.id = userId
  state.wallet = wallet
}

const DISCONNECT = (state) => {
  state.connected = false
  state.username = null
  state.id = null
  state.wallet = null
}

const mutations = {
  TOGGLE_DRAWER,
  SET_FIXTURES,
  SET_RESULTS,
  SET_BETS,
  ADD_MONEY,
  ADD_TEAM,
  SET_STANDINGS,
  CONNECT,
  DISCONNECT
}

export default mutations
