import Vue from 'vue'
import Vuex from 'vuex'
import login from './modules/login'
import user from './modules/user'
import blog from './modules/blog'
import comment from './modules/comment'
Vue.use(Vuex)

export default new Vuex.Store({
    modules: {
      login,
      user,
      blog,
      comment
    }
})
