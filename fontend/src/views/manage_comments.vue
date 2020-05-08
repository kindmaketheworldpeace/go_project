<template>
<div style="width: 100%">
   <div class="uk-width-1-1 uk-margin-bottom">
        <div class="uk-panel uk-panel-box">
            <ul class="uk-breadcrumb">
                <li class="uk-active"><span>评论</span></li>
                <li><a href="/manage/blogs">日志</a></li>
                <li><a href="/manage/users">用户</a></li>
            </ul>
        </div>
    </div>

    <div id="error" class="uk-width-1-1">
    </div>

    <div id="loading" class="uk-width-1-1 uk-text-center">
        <span><i class="uk-icon-spinner uk-icon-medium uk-icon-spin"></i> 正在加载...</span>
    </div>

    <div class="uk-width-1-1" style="display:none">
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
                <tr v-for="comment in comments" >
                    <td>
                        <span >{{comment.user_name}}</span>
                    </td>
                    <td>
                        <span >{{comment.content}}</span>
                    </td>
                    <td>
                        <span >{{comment.created_at.toDateTime()}}</span>
                    </td>
                    <td>
                      <a href="#0" @click="" ><i class="uk-icon-trash-o"></i></a>
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
        name: "manage_comments",
      data () {
          return {
            comments:[]
          }
      }
    }
</script>

<style scoped>
   @import '../assets/css/uikit.min.css';
  @import '../assets/css/uikit.gradient.min.css';
  @import '../assets/css/awesome.css';

</style>
