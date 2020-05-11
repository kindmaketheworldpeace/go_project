<template>
  <div style="width: 100%">
    <div class="uk-width-1-1 uk-margin-bottom">
      <div class="uk-panel uk-panel-box">
        <ul class="uk-breadcrumb">
          <li><router-link to="/manage_comments">评论</router-link></li>
          <li><router-link to="/manage_blogs">日志</router-link></li>
          <li class="uk-active"><span>用户</span></li>
        </ul>
      </div>
    </div>

    <div id="error" class="uk-width-1-1">
    </div>

    <div id="loading" v-if="is_finish" class="uk-width-1-1 uk-text-center">
      <span><i class="uk-icon-spinner uk-icon-medium uk-icon-spin"></i> 正在加载...</span>
    </div>

    <div class="uk-width-1-1">
      <table class="uk-table uk-table-hover">
        <thead>
        <tr>
          <th class="uk-width-4-10">名字</th>
          <th class="uk-width-4-10">电子邮件</th>
          <th class="uk-width-2-10">注册时间</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="user in users">
          <td>
            <span v-text="user.Name"></span>
            <span v-if="user.admin" style="color:#d05"><i class="uk-icon-key"></i> 管理员</span>
          </td>
          <td>
            <a>{{user.Email}}</a>
          </td>
          <td>
            <span>{{user.Created_at}}</span>
          </td>
        </tr>
        </tbody>
      </table>
      <!--<pagination></pagination>-->
    </div>
  </div>
</template>

<script>
  export default {
    name: "manage_user",
    data() {
      return {
        is_finish: false,
        users: []
      }
    },
    methods: {
      get_user: function () {
        this.is_finish = !this.is_finish
        let params = {}
        this.$store.dispatch('user/get_user', params).then(res => {

            if (res.data.result) {
              this.users = res.data.users
            } else {
              this.$message.error("失败！" + res.data.message)
            }
             this.is_finish = !this.is_finish
          }
        )

      }
    },
    created() {
      this.get_user()
    }
  }
</script>

<style scoped>
  @import '../assets/css/uikit.min.css';
  @import '../assets/css/uikit.gradient.min.css';
  @import '../assets/css/awesome.css';

</style>
