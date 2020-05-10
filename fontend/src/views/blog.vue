<template>
  <div style="width: 100%">
    <div style="width: 75%;display: inline-block">
      <div class="uk-width-medium-3-4">
        <article class="uk-article">
          <h2>{{ blog.Name }}</h2>
          <p class="uk-article-meta">发表于{{ blog.Created_at}}</p>
          <p>{{ blog.Content }}</p>
        </article>

        <hr class="uk-article-divider">

        <div v-if="login_user!==''">
          <h3>发表评论</h3>

          <article class="uk-comment">
            <header class="uk-comment-header">
              <img class="uk-comment-avatar uk-border-circle" width="50" height="50" :src="login_user.image">
              <h4 class="uk-comment-title">{{ login_user.name }}</h4>
            </header>
            <div class="uk-comment-body">
              <form id="form-comment" class="uk-form">
                <div class="uk-alert uk-alert-danger uk-hidden"></div>
                <div class="uk-form-row">
                  <textarea v-model="curcomment" rows="6" placeholder="说点什么吧"
                            style="width:100%;resize:none;"></textarea>
                </div>
                <div class="uk-form-row">
                  <button type="button" @click="create_comment()" class="uk-button uk-button-primary"><i
                    class="uk-icon-comment"></i> 发表评论
                  </button>
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
                    <p class="uk-comment-meta">{{ comment.Created_at }}</p>
                  </header>
                  <div class="uk-comment-body">
                    {{ comment.Content}}
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




    </div>
    <div style="width: 24%;float: right" class="uk-panel uk-panel-header">
       <div style="width: 100%;margin-bottom: 10px" class="uk-width-medium-1-4">
        <div class="uk-panel uk-panel-box">
          <div class="uk-text-center">
            <img class="uk-border-circle" width="120" height="120" :src="blog.User_image">
            <h3>{{ blog.User_name }}</h3>
          </div>
        </div>
      </div>
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
  import {mapGetters} from 'vuex'

  export default {
    name: "blog",
    data() {
      return {
        blog: {
          Name: "",
          Created_at: "",
          Content: "",
          User_image: '',
          User_name: ''

        },
        id: "",
        comments: [],
        curcomment: "",

      }
    },
    methods: {
      get_blog: function () {
        let params = {
          id: this.id,
        }
        this.$store.dispatch('blog/get_blog', params).then(res => {
            if (res.data.result) {
              this.blog = res.data.blog
              this.comments = res.data.comments
            } else {
              this.$message.error("失败！" + res.data.message)
            }
          }
        )
      },
      create_comment: function () {
        let params = {
          blog_id: this.id,
          user_id: this.login_user.user_id,
          user_name:this.login_user.user_name,
          user_image:this.login_user.user_image,
          content:this.curcomment
        }
        this.$store.dispatch('comment/create_comment', params).then(res => {
            if (res.data.result) {
              this.get_blog()
            } else {
              this.$message.error("失败！" + res.data.message)
            }
          }
        )

      }

    },
    created() {
      this.id = this.$route.params.id;
      this.get_blog()
    },
    computed: {
      ...mapGetters({'login_user': 'login/get_login_user'}),

    },
  }
</script>

<style scoped>
  @import '../assets/css/uikit.min.css';
  @import '../assets/css/uikit.gradient.min.css';
  @import '../assets/css/awesome.css';
</style>
