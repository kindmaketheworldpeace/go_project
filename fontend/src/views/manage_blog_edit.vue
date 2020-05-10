<template>
  <div style="width: 100%">
    <div class="uk-width-1-1 uk-margin-bottom">
      <div class="uk-panel uk-panel-box">
        <ul class="uk-breadcrumb">
          <li><router-link to="/manage_comments">评论</router-link></li>
          <li><router-link to="/manage_blogs">日志</router-link></li>
          <li><router-link to="/manage_users">用户</router-link></li>
        </ul>
      </div>
    </div>

    <div id="error" class="uk-width-1-1">
    </div>

    <div id="loading" v-if="is_finish" class="uk-width-1-1 uk-text-center">
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

          <button @click="create_blog()"   type="button" class="uk-button uk-button-primary">   <i class="uk-icon-save"></i> 保存</button>
          <router-link to="/manage_blogs" class="uk-button"><i class="uk-icon-times"></i> 取消</router-link>
        </div>
      </form>
    </div>
  </div>

</template>

<script>
  import {mapGetters} from 'vuex'
  export default {
    name: "manage_blog_edit",
    data() {
      return {
        name:'',
        summary: "",
        content: '',
        is_finish:false
      }
    },
    methods:{
      create_blog:function () {
        this.is_finish = !this.is_finish
        const login_user=this.login_user
        let params={
          user_id:login_user.user_id,
          user_name:login_user.user_name,
          user_image:login_user.user_image,
          name:this.name,
          summary:this.summary,
          content:this.content
        }
         this.$store.dispatch('blog/create_blog', params).then(res => {
            if (res.data.result) {
               this.$message.success("成功！")
                this.$router.push({path:'/manage_blogs'})
            } else {
              this.$message.error("失败！" + res.data.message)
            }
             this.is_finish = !this.is_finish
          }
        )

      }
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
