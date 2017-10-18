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

const actions = {
  toggleDrawer,
  setFixtures,
  setResults,
  addTeam
}

export default actions
