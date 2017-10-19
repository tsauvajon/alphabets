import api from '@/helpers/url'

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

const connect = async ({ commit }, { username, password }) => {
  try {
    const resultAsync = api.blockchain.connect(username, password)

    const result = await resultAsync

    if (result.status !== 200) {
      return false
    }

    console.log(result.data)
    commit('CONNECT', { ...result.data })
    return true
  } catch (e) {
    console.error(e)
    return false
  }
}

const disconnect = ({ commit }) => commit('DISCONNECT')

const actions = {
  toggleDrawer,
  setFixtures,
  setResults,
  addTeam,
  setStandings,
  connect,
  disconnect
}

export default actions
