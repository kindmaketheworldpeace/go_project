<template>
  <div style="width: 100%">
  <div style="width: 75%;display: inline-block">
    <div class="uk-width-medium-3-4">
      <article class="uk-article">
        <h2>{{ blog.name }}</h2>
        <p class="uk-article-meta">发表于{{ blog.created_at}}</p>
        <p>{{ blog.html_content }}</p>
      </article>

      <hr class="uk-article-divider">

      <div v-if="user!==''">
        <h3>发表评论</h3>

        <article class="uk-comment">
          <header class="uk-comment-header">
            <img class="uk-comment-avatar uk-border-circle" width="50" height="50" :src="user.image">
            <h4 class="uk-comment-title">{{ user.name }}</h4>
          </header>
          <div class="uk-comment-body">
            <form id="form-comment" class="uk-form">
              <div class="uk-alert uk-alert-danger uk-hidden"></div>
              <div class="uk-form-row">
                <textarea rows="6" placeholder="说点什么吧" style="width:100%;resize:none;"></textarea>
              </div>
              <div class="uk-form-row">
                <button type="submit" class="uk-button uk-button-primary"><i class="uk-icon-comment"></i> 发表评论</button>
              </div>
            </form>
          </div>
        </article>

        <hr class="uk-article-divider">
      </div>

      <h3>最新评论</h3>

      <ul class="uk-comment-list">
        <div v-if="comments.length!==0">
          <div v-for="(comment,index) in comments">
            <li>
              <article class="uk-comment">
                <header class="uk-comment-header">
                  <img class="uk-comment-avatar uk-border-circle" width="50" height="50"
                       :src="comment.user_image">
                  <h4 class="uk-comment-title">{{ comment.user_name }} <p v-if="comment.user_id==blog.user_id">作者</p>
                  </h4>
                  <p class="uk-comment-meta">{{ comment.created_at }}</p>
                </header>
                <div class="uk-comment-body">
                  {{ comment.html_content}}
                </div>
              </article>
            </li>
          </div>
        </div>

        <div v-else>
          <p>还没有人评论...</p>
        </div>
      </ul>

    </div>

    <div class="uk-width-medium-1-4">
      <div class="uk-panel uk-panel-box">
        <div class="uk-text-center">
          <img class="uk-border-circle" width="120" height="120" :src="blog.user_image">
          <h3>{{ blog.user_name }}</h3>
        </div>
      </div>
    </div>


  </div>
   <div style="width: 24%;float: right" class="uk-panel uk-panel-header">
      <h3 class="uk-panel-title">友情链接</h3>
      <ul class="uk-list uk-list-line">
        <li><i class="uk-icon-link"></i> <a href="#">编程</a></li>
        <li><i class="uk-icon-link"></i> <a href="#">思考</a></li>
        <li><i class="uk-icon-link"></i> <a href="#">读书</a></li>
      </ul>
    </div>
  </div>
</template>

<script>
  export default {
    name: "blog",
    data() {
      return {
        blog: {
          name: "",
          created_at: "",
          html_content: "",
          user_image: '',
          user_name: ''

        },
        user: {
          image: '',
          name: ''
        },
        comments: []
      }
    }
  }
</script>

<style scoped>
  @import '../assets/css/uikit.min.css';
  @import '../assets/css/uikit.gradient.min.css';
  @import '../assets/css/awesome.css';
</style>
