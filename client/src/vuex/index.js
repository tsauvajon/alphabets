import Vue from 'vue'
import Vuex from 'vuex'

import actions from './actions'
import mutations from './mutations'

Vue.use(Vuex)

const initialState = {
  drawer: false,
  fixtures: null,
  teams: []
}

const getters = {
  drawer: ({ drawer }) => drawer,
  fixtures: ({ fixtures }) => fixtures,
  teams: ({ teams }) => teams
}

export default new Vuex.Store({
  state: initialState,
  mutations,
  actions,
  getters
})
