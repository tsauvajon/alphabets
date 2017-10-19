import Vue from 'vue'
import Vuex from 'vuex'

import actions from './actions'
import mutations from './mutations'

Vue.use(Vuex)

const initialState = {
  drawer: false,
  fixtures: null,
  results: null,
  teams: [],
  standings: null,
  id: null,
  username: null,
  wallet: null,
  connected: false
}

const getters = {
  drawer: ({ drawer }) => drawer,
  fixtures: ({ fixtures }) => fixtures,
  results: ({ results }) => results,
  teams: ({ teams }) => teams,
  standings: ({ standings }) => standings,
  connected: ({ connected }) => connected,
  username: ({ username }) => username,
  id: ({ id }) => id,
  wallet: ({ wallet }) => wallet
}

export default new Vuex.Store({
  state: initialState,
  mutations,
  actions,
  getters
})
