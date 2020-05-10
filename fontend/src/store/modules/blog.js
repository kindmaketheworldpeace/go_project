import api from '@/api/api'

const state = {
}

const getters = {
}

const mutations = {
}

const actions = {
  create_blog({commit, state}, param) {
        return api.create_blog(param)
    },
  get_blogs({commit, state}, param) {
        return api.get_blogs(param)
    },
  get_blog({commit, state}, param) {
        return api.get_blog(param)
    },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
