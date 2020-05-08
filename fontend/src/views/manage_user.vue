<template>
  <div style="width: 100%">
    <div class="uk-width-1-1 uk-margin-bottom">
      <div class="uk-panel uk-panel-box">
        <ul class="uk-breadcrumb">
          <li><a href="/manage/comments">评论</a></li>
          <li><a href="/manage/blogs">日志</a></li>
          <li class="uk-active"><span>用户</span></li>
        </ul>
      </div>
    </div>

    <div id="error" class="uk-width-1-1">
    </div>

    <div id="loading" class="uk-width-1-1 uk-text-center">
      <span><i class="uk-icon-spinner uk-icon-medium uk-icon-spin"></i> 正在加载...</span>
    </div>

    <div  class="uk-width-1-1">
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
            <span v-text="user.name"></span>
            <span v-if="user.admin" style="color:#d05"><i class="uk-icon-key"></i> 管理员</span>
          </td>
          <td>
            <a>{{user.email}}</a>
          </td>
          <td>
            <span>{{user.created_at.toDateTime()}}</span>
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
        users: []
      }
    }
  }
</script>

<style scoped>
   @import '../assets/css/uikit.min.css';
  @import '../assets/css/uikit.gradient.min.css';
  @import '../assets/css/awesome.css';

</style>
