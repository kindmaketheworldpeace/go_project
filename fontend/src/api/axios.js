import Vue from 'vue'
import axios from 'axios'

axios.defaults.baseURL = 'http://127.0.0.1:8080/';
axios.defaults.withCredentials = false;

axios.interceptors.response.use(response => {
  if (response.status >= 400) {
      return {
          code: response.status,
          message: '请求异常，请刷新重试',
          result: false
      }
  }
  return response
}, error => {

  // 未认证
  if (error.response.status == 401) {
    window.location.href = "http://127.0.0.1:8000/#/login"
  }
 return error.response
});
Vue.prototype.$http = axios;
export const $axios = axios
