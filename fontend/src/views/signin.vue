<template>
  <div class="uk-vertical-align uk-text-center uk-height-1-1">
    <div class="uk-vertical-align-middle" style="width: 320px">
      <p><a href="/" class="uk-icon-button"><i class="uk-icon-html5"></i></a> <a href="/">Awesome Python Webapp</a></p>
      <form  class="uk-panel uk-panel-box uk-form">
        <div class="uk-alert uk-alert-danger uk-hidden"></div>
        <div class="uk-form-row">
          <div class="uk-form-icon uk-width-1-1">
            <i class="uk-icon-envelope-o"></i>
            <input v-model="email" name="email" type="text" placeholder="电子邮件" maxlength="50"
                   class="uk-width-1-1 uk-form-large">
          </div>
        </div>
        <div class="uk-form-row">
          <div class="uk-form-icon uk-width-1-1">
            <i class="uk-icon-lock"></i>
            <input v-model="passwd" name="passwd" type="password" placeholder="口令" maxlength="50"
                   class="uk-width-1-1 uk-form-large">
          </div>
        </div>
        <div class="uk-form-row">
          <button type="submit" class="uk-width-1-1 uk-button uk-button-primary uk-button-large"><i
            class="uk-icon-sign-in"></i> 登录
          </button>
        </div>
      </form>
    </div>
  </div>

</template>

<script>
  export default {
    name: "signin",
    data () {
      return {
        email:'',
        passwd:''
      }
    }
  }
</script>

<style scoped>
  @import '../assets/css/uikit.min.css';
  @import '../assets/css/uikit.gradient.min.css';
  body {

    height: 100%


  }
</style>
