import api from '@/api/api'

const state = {
}

const getters = {
}

const mutations = {
}

const actions = {
  get_user({commit, state}, param) {
        return api.get_user(param)
    },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
