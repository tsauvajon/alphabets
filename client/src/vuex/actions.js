const toggleDrawer = ({ commit }) => {
  commit('TOGGLE_DRAWER')
}

const setFixtures = ({ commit }, { fixtures }) => {
  commit('SET_FIXTURES', { fixtures })
}

const setResults = ({ commit }, { results }) => {
  commit('SET_RESULTS', { results })
}

const addTeam = ({ commit }, { team }) => {
  commit('ADD_TEAM', { team })
}

const setStandings = ({ commit }, { standings }) => {
  commit('SET_STANDINGS', { standings })
}

const actions = {
  toggleDrawer,
  setFixtures,
  setResults,
  addTeam,
  setStandings
}

export default actions
