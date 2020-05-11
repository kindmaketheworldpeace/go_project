<template>
  <div style="width: 100%">
    <div style="display: inline-block" class="uk-width-medium-3-4">
      <div v-for=" (blog,index) in blogs">
        <article class="uk-article">
          <h2><a href="">{{ blog.Name }}</a></h2>
          <p class="uk-article-meta">发表于{{ blog.Created_at}}</p>d
          <p>{{ blog.Summary }}</p>
          <p><router-link :to="'/blog/'+blog.I">继续阅读 <i class="uk-icon-angle-double-right"></i></router-link></p>
        </article>
        <hr class="uk-article-divider">
      </div>
      {{ page }}

    </div>
 <div style="float: right;width: 24%" class="uk-width-medium-1-4">
      <div class="uk-panel uk-panel-header">
        <h3 class="uk-panel-title">友情链接</h3>
        <ul class="uk-list uk-list-line">
          <li><i class="uk-icon-thumbs-o-up"></i> <a target="_blank"
                                                     href="http://www.liaoxuefeng.com/category/0013738748415562fee26e070fa4664ad926c8e30146c67000">编程</a>
          </li>
          <li><i class="uk-icon-thumbs-o-up"></i> <a target="_blank"
                                                     href="http://www.liaoxuefeng.com/category/0013738748248885ddf38d8cd1b4803aa74bcda32f853fd000">读书</a>
          </li>
          <li><i class="uk-icon-thumbs-o-up"></i> <a target="_blank"
                                                     href="http://www.liaoxuefeng.com/wiki/001374738125095c955c1e6d8bb493182103fac9270762a000">Python教程</a>
          </li>
          <li><i class="uk-icon-thumbs-o-up"></i> <a target="_blank"
                                                     href="http://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000">Git教程</a>
          </li>
        </ul>
      </div>
    </div>

  </div>
</template>

<script>
  export default {
    name: "blogs",
    data() {
      return {
        blogs: [],
        page: 1

      }
    },
    methods: {
      get_blogs: function () {
        let params = {
        }
        this.$store.dispatch('blog/get_blogs', params).then(res => {
            if (res.data.result) {
                this.blogs = res.data.blogs
            } else {
              this.$message.error("失败！" + res.data.message)
            }
          }
        )
      }
    },
    created() {
      this.get_blogs()
    },
  }
</script>

<style scoped>
  @import '../assets/css/uikit.min.css';
  @import '../assets/css/uikit.gradient.min.css';
  @import '../assets/css/awesome.css';
</style>
