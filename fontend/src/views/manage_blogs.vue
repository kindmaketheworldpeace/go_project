<template>
  <div style="width: 100%">
    <div class="uk-width-1-1 uk-margin-bottom">
      <div class="uk-panel uk-panel-box">
        <ul class="uk-breadcrumb">
          <li><router-link to="/manage_comments">评论</router-link></li>
          <li class="uk-active"><span>日志</span></li>
          <li><router-link to="/manage_users">用户</router-link></li>
        </ul>
      </div>
    </div>

    <div id="error" class="uk-width-1-1">
    </div>

    <div id="loading"  v-if="is_finish" class="uk-width-1-1 uk-text-center">
      <span><i class="uk-icon-spinner uk-icon-medium uk-icon-spin"></i> 正在加载...</span>
    </div>

    <div  class="uk-width-1-1">
      <router-link to="/manage_blog_edit" class="uk-button uk-button-primary"><i class="uk-icon-plus"></i> 新日志</router-link>

      <table class="uk-table uk-table-hover">
        <thead>
        <tr>
          <th class="uk-width-5-10">标题 / 摘要</th>
          <th class="uk-width-2-10">作者</th>
          <th class="uk-width-2-10">创建时间</th>
          <th class="uk-width-1-10">操作</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for=" (i) in blogs">
          <td>
            <a target="_blank">{{i.Name}}</a>
          </td>
          <td>
            <a target="_blank" href=''>{{i.User_name}}</a>
          </td>
          <td>
            <span>{{i.Created_at}}</span>
          </td>
          <td>
            <a href="#0" @click="" ><i class="uk-icon-edit"></i></a>
            <a href="#0" @click=""><i class="uk-icon-trash-o"></i></a>
          </td>
        </tr>
        </tbody>
      </table>

      <!--<pagination></pagination>-->
    </div>
  </div>
</template>

<script>
   import {mapGetters} from 'vuex'
  export default {
    name: "manage_blogs",
    data() {
      return {
        blogs: [],
        is_finish:false
      }
    },
    methods: {
      get_blogs: function () {
        this.is_finish=!this.is_finish
        const login_user = this.login_user
        let params = {
          user_id: login_user.user_id,
        }
        this.$store.dispatch('blog/get_blogs', params).then(res => {
            if (res.data.result) {
                this.blogs = res.data.blogs
              console.log(this.blogs)

            } else {
              this.$message.error("失败！" + res.data.message)
            }
            this.is_finish=!this.is_finish
          }
        )
      }
    },
    created() {
      this.get_blogs()
    },
    computed: {
      ...mapGetters({'login_user':'login/get_login_user'}),

    },
  }
</script>

<style scoped>
   @import '../assets/css/uikit.min.css';
  @import '../assets/css/uikit.gradient.min.css';
  @import '../assets/css/awesome.css';

</style>
