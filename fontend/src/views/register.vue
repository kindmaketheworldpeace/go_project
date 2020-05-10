<template>
  <div class="uk-width-2-3">
    <h1>欢迎注册！</h1>
    <form class="uk-form uk-form-stacked">
      <div class="uk-alert uk-alert-danger uk-hidden"></div>
      <div class="uk-form-row">
        <label class="uk-form-label">名字:</label>
        <div class="uk-form-controls">
          <input v-model="name" type="text" maxlength="50" placeholder="名字" class="uk-width-1-1">
        </div>
      </div>
      <div class="uk-form-row">
        <label class="uk-form-label">电子邮件:</label>
        <div class="uk-form-controls">
          <input v-model="email" type="text" maxlength="50" placeholder="your-name@example.com" class="uk-width-1-1">
        </div>
      </div>
      <div class="uk-form-row">
        <label class="uk-form-label">输入口令:</label>
        <div class="uk-form-controls">
          <input v-model="password1" type="password" maxlength="50" placeholder="输入口令" class="uk-width-1-1">
        </div>
      </div>
      <div class="uk-form-row">
        <label class="uk-form-label">重复口令:</label>
        <div class="uk-form-controls">
          <input v-model="password2" type="password" maxlength="50" placeholder="重复口令" class="uk-width-1-1">
        </div>
      </div>
      <div class="uk-form-row">
        <button @click="register()" type="submit" class="uk-button uk-button-primary"><i class="uk-icon-user"></i> 注册
        </button>
      </div>
    </form>
  </div>
</template>

<script>
  import md5 from 'js-md5'
  export default {
    name: "register",
    data() {
      return {
        name: '',
        email: '',
        password1: '',
        password2: ''
      }
    },
    methods: {
      register: function () {
        if (this.password1!==this.password2) {
             this.$message.error("密码不一致！")
             return
        }
        let params = {
          email: this.email,
          name: this.name,
          password1: md5(this.password1),
          password2: md5(this.password1)
        }
        this.$store.dispatch('login/register', params).then(res => {

          if(res.data.result){
            this.$message.success("注册成功！")
             this.$router.push({path:'/signin'})
          }else {
            this.$message.error("注册失败！" +res.data.message)
          }
          }
        )

      }

    }
  }
</script>

<style scoped>
  @import '../assets/css/uikit.min.css';
  @import '../assets/css/uikit.gradient.min.css';
  @import '../assets/css/awesome.css';

</style>
