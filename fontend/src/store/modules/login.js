import api from '@/api/api'

const state = {
  login_user:{
     user_id:"1",
    user_image: "",
    user_name: "admin"
  }
}

const getters = {
get_login_user: state => state.login_user
}

const mutations = {
 update_user: (state, payload) => {
    state.login_user = {
      user_id: payload.user.User_id,
      user_image:payload.user.User_image,
      user_name:payload.user.User_name
    }
  },
}
var get_user=function (commit,data) {
  if (data.result) {
     commit('update_user', data)
  }

}
const actions = {

    async login({commit, state}, param) {
    const ret_data = await api.login(param)

     get_user(commit,ret_data.data)

    return ret_data

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
