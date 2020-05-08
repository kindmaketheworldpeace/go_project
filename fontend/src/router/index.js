import Vue from 'vue'
import Router from 'vue-router'
import signin from '@/views/signin'
import blogbase from '@/views/blogbase'
import blog from '@/views/blog'
import blogs from '@/views/blogs'
import manage_blog_edit from '@/views/manage_blog_edit'
import manage_blogs from '@/views/manage_blogs'
import manage_comments from '@/views/manage_comments'
import manage_user from '@/views/manage_user'
import register from '@/views/register'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/signin',
      name: 'signin',
      component: signin
    }, {
      path: '/home',
      name: 'blogbase',
      component: blogbase,
      children: [{
        path: '/blog',
        name: 'blog',
        component: blog
      }, {
        path: '/blogs',
        name: 'blogs',
        component: blogs
      },{
        path: '/manage_blog_edit',
        name: 'manage_blog_edit',
        component: manage_blog_edit
      },{
        path: '/manage_blogs',
        name: 'manage_blogs',
        component: manage_blogs
      },{
        path: '/manage_comments',
        name: 'manage_comments',
        component: manage_comments
      },{
        path: '/manage_user',
        name: 'manage_user',
        component: manage_user
      },{
        path: '/register',
        name: 'register',
        component: register
      },]
    }
  ]
})
