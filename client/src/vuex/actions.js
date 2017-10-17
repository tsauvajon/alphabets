const toggleDrawer = ({ commit }) => {
  commit('TOGGLE_DRAWER')
}

const setFixtures = ({ commit }, { fixtures }) => {
  commit('SET_FIXTURES', { fixtures })
}

const addTeam = ({ commit }, { team }) => {
  commit('ADD_TEAM', { team })
}

const actions = {
  toggleDrawer,
  setFixtures,
  addTeam
}

export default actions
