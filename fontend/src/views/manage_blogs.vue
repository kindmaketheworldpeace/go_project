<template>
  <div style="width: 100%">
    <div class="uk-width-1-1 uk-margin-bottom">
      <div class="uk-panel uk-panel-box">
        <ul class="uk-breadcrumb">
          <li><a href="/manage/comments">评论</a></li>
          <li class="uk-active"><span>日志</span></li>
          <li><a href="/manage/users">用户</a></li>
        </ul>
      </div>
    </div>

    <div id="error" class="uk-width-1-1">
    </div>

    <div id="loading" class="uk-width-1-1 uk-text-center">
      <span><i class="uk-icon-spinner uk-icon-medium uk-icon-spin"></i> 正在加载...</span>
    </div>

    <div id="vm" class="uk-width-1-1">
      <a href="/manage/blogs/create" class="uk-button uk-button-primary"><i class="uk-icon-plus"></i> 新日志</a>

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
        <tr v-for="blog in blogs">
          <td>
            <a target="_blank">{{blog.name}}</a>
          </td>
          <td>
            <a target="_blank" href=''>{{blog.user_name}}</a>
          </td>
          <td>
            <span>{{blog.created_at.toDateTime()}}</span>
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
  export default {
    name: "manage_blogs",
    data() {
      return {
        blogs: []
      }
    }
  }
</script>

<style scoped>
   @import '../assets/css/uikit.min.css';
  @import '../assets/css/uikit.gradient.min.css';
  @import '../assets/css/awesome.css';

</style>
