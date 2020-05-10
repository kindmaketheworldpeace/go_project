import api from '@/api/api'

const state = {
}

const getters = {
}

const mutations = {
}

const actions = {
  create_comment({commit, state}, param) {
        return api.create_comment(param)
    },
  get_comment({commit, state}, param) {
        return api.get_comment(param)
    },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
