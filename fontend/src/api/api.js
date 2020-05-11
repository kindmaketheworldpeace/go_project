import {$axios} from './axios'

export default {
   login: params => {
        return $axios.post('api/login/', params)
    },
   register: params => {
        return $axios.post('api/register/', params)
    },
  get_user: params => {
        return $axios.get('api/get_user/', params)
    },
  create_blog: params => {
        return $axios.post('api/create_blog/', params)
    },
   get_blogs: params => {
        return $axios.get('api/get_blogs/', {params:params})
    },
   get_blog: params => {
        return $axios.get('api/get_blog/', {params:params})
    },
   create_comment: params => {
        return $axios.post('api/create_comment/', params)
    },
  get_comment: params => {
        return $axios.get('api/get_comments/', {params:params})
    },
}
