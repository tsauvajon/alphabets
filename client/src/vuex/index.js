import Vue from 'vue'
import Vuex from 'vuex'

import actions from './actions'
import mutations from './mutations'

Vue.use(Vuex)

const initialState = {
  drawer: false
}

const getters = {
  drawer: ({ drawer }) => drawer
}

export default new Vuex.Store({
  state: initialState,
  mutations,
  actions,
  getters
})
