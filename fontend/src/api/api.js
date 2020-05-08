import {$axios} from './axios'

export default {
   login: params => {
        return $axios.post('api/blog_user/user/login/', params)
    },
   register: params => {
        return $axios.post('api/blog_user/user/register/', params)
    },
}
