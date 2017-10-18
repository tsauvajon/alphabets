const toggleDrawer = ({ commit }) => {
  commit('TOGGLE_DRAWER')
}

const setFixtures = ({ commit }, { fixtures }) => {
  commit('SET_FIXTURES', { fixtures })
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
  addTeam,
  setStandings
}

export default actions
