import api from '@/api/api'

const state = {
}

const getters = {
}

const mutations = {
}

const actions = {

     login({commit, state}, param) {
        return api.login(param)
    },
  register({commit, state}, param) {
        return api.register(param)
    },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
