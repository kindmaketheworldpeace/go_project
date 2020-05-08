<template>
  <div >
    <nav class="uk-navbar uk-navbar-attached uk-margin-bottom">
      <div class="uk-container uk-container-center">
        <a href="/" class="uk-navbar-brand">Awesome</a>
        <ul class="uk-navbar-nav">
          <li data-url="blogs"><a href="/"><i class="uk-icon-home"></i> 日志</a></li>
          <li><a target="_blank"
                 href="http://www.liaoxuefeng.com/wiki/0014316089557264a6b348958f449949df42a6d3a2e542c000"><i
            class="uk-icon-book"></i> 教程</a></li>
          <li><a target="_blank" href="https://github.com/michaelliao/awesome-python3-webapp"><i
            class="uk-icon-code"></i> 源码</a></li>
        </ul>
        <div class="uk-navbar-flip">
          <ul class="uk-navbar-nav">
            <div v-if="user!==''">
              <li class="uk-parent" data-uk-dropdown>
                <a href="#0"><i class="uk-icon-user"></i> </a>
                <div class="uk-dropdown uk-dropdown-navbar">
                  <ul class="uk-nav uk-nav-navbar">
                    <li><a href="/signout"><i class="uk-icon-sign-out"></i> 登出</a></li>
                  </ul>
                </div>
              </li>
            </div>
            <div v-else>
              <li><a href="/signin"><i class="uk-icon-sign-in"></i> 登陆</a></li>
              <li><a href="/register"><i class="uk-icon-edit"></i> 注册</a></li>
            </div>
          </ul>
        </div>
      </div>
    </nav>

    <div class="uk-container uk-container-center">
      <div class="uk-grid">
                <router-view/>

      </div>
    </div>

    <div class="uk-margin-large-top" style="background-color:#eee; border-top:1px solid #ccc;">
      <div class="uk-container uk-container-center uk-text-center">
        <div class="uk-panel uk-margin-top uk-margin-bottom">
          <p>
            <a target="_blank" href="http://weibo.com/liaoxuefeng" class="uk-icon-button uk-icon-weibo"></a>
            <a target="_blank" href="https://github.com/michaelliao" class="uk-icon-button uk-icon-github"></a>
            <a target="_blank" href="http://www.linkedin.com/in/liaoxuefeng"
               class="uk-icon-button uk-icon-linkedin-square"></a>
            <a target="_blank" href="https://twitter.com/liaoxuefeng" class="uk-icon-button uk-icon-twitter"></a>
          </p>
          <p>Powered by <a href="http://awesome.liaoxuefeng.com">Awesome Python Webapp</a>. Copyright &copy; 2014. [<a
            href="/manage/" target="_blank">Manage</a>]</p>
          <p><a href="http://www.liaoxuefeng.com/" target="_blank">www.liaoxuefeng.com</a>. All rights reserved.</p>
          <a target="_blank" href="http://www.w3.org/TR/html5/"><i class="uk-icon-html5"
                                                                   style="font-size:64px; color: #444;"></i></a>
        </div>

      </div>
    </div>
  </div>

</template>

<script>
  export default {
    name: 'blogbase',
    data() {
      return {
        user: ""
      }
    }
  }
</script>

<style scoped>
     @import '../assets/css/uikit.min.css';
     @import '../assets/css/uikit.gradient.min.css';
     @import '../assets/css/awesome.css';
</style>
