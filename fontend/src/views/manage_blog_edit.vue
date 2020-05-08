<template>
  <div style="width: 100%">
    <div class="uk-width-1-1 uk-margin-bottom">
      <div class="uk-panel uk-panel-box">
        <ul class="uk-breadcrumb">
          <li><a href="/manage/comments">评论</a></li>
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

    <div  class="uk-width-2-3">
      <form class="uk-form uk-form-stacked">
        <div class="uk-alert uk-alert-danger uk-hidden"></div>
        <div class="uk-form-row">
          <label class="uk-form-label">标题:</label>
          <div class="uk-form-controls">
            <input v-model="name" name="name" type="text" placeholder="标题" class="uk-width-1-1">
          </div>
        </div>
        <div class="uk-form-row">
          <label class="uk-form-label">摘要:</label>
          <div class="uk-form-controls">
            <textarea v-model="summary" rows="4" name="summary" placeholder="摘要" class="uk-width-1-1"
                      style="resize:none;"></textarea>
          </div>
        </div>
        <div class="uk-form-row">
          <label class="uk-form-label">内容:</label>
          <div class="uk-form-controls">
            <textarea v-model="content" rows="16" name="content" placeholder="内容" class="uk-width-1-1"
                      style="resize:none;"></textarea>
          </div>
        </div>
        <div class="uk-form-row">
          <button type="submit" class="uk-button uk-button-primary"><i class="uk-icon-save"></i> 保存</button>
          <a href="/manage/blogs" class="uk-button"><i class="uk-icon-times"></i> 取消</a>
        </div>
      </form>
    </div>
  </div>

</template>

<script>
  export default {
    name: "manage_blog_edit",
    data() {
      return {
        name:'',
        summary: "",
        content: ''
      }
    }

  }
</script>

<style scoped>
  @import '../assets/css/uikit.min.css';
  @import '../assets/css/uikit.gradient.min.css';
  @import '../assets/css/awesome.css';

</style>
