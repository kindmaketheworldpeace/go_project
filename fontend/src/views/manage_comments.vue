<template>
  <div style="width: 100%">
    <div class="uk-width-1-1 uk-margin-bottom">
      <div class="uk-panel uk-panel-box">
        <ul class="uk-breadcrumb">
          <li class="uk-active"><span>评论</span></li>
          <li>
            <router-link to="/manage_blogs">日志</router-link>
          </li>
          <li>
            <router-link to="/manage_users">用户</router-link>
          </li>
        </ul>
      </div>
    </div>

    <div id="error" class="uk-width-1-1">
    </div>

    <div id="loading" v-if="is_finish" class="uk-width-1-1 uk-text-center">
      <span><i class="uk-icon-spinner uk-icon-medium uk-icon-spin"></i> 正在加载...</span>
    </div>

    <div class="uk-width-1-1" >
      <table class="uk-table uk-table-hover">
        <thead>
        <tr>
          <th class="uk-width-2-10">作者</th>
          <th class="uk-width-5-10">内容</th>
          <th class="uk-width-2-10">创建时间</th>
          <th class="uk-width-1-10">操作</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for=" (comment) in comments">
          <td>
            <span>{{comment.User_name}}</span>
          </td>
          <td>
            <span>{{comment.Content}}</span>
          </td>
          <td>
            <span>{{comment.Created_at}}</span>
          </td>
          <td>
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
    name: "manage_comments",
    data() {
      return {
        comments: [],
        is_finish: false
      }
    },
    methods:{
       get_comment: function () {
        this.is_finish=!this.is_finish
        const login_user = this.login_user
        let params = {
          user_id: login_user.user_id,
        }
        this.$store.dispatch('comment/get_comment', params).then(res => {
            if (res.data.result) {
                this.comments = res.data.comments
               console.log(this.comments)

            } else {
              this.$message.error("失败！" + res.data.message)
            }
            this.is_finish=!this.is_finish
          }
        )
      }

    },
    created() {
      this.get_comment()
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
