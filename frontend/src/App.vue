<template>
  <div id="app">
    <div class="header">
      <div class="tomain"><span class="top">您好! 欢迎光临 </span>
        <font-awesome-icon icon="user-circle"  /><a class="top"> 登录</a><b class="side">|</b> <font-awesome-icon icon="check-square"  /><a  class="top">注册</a>
      </div>
      <div class="buycar"></div>
      <div class="menu">
        <div class="navcenter" >
          <div class="item">
            <a class="menu-item">
              <i style="font-style: normal">首页</i></a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'App'
  }
</script>

<style>
  .menu-item {
   color: white;
    line-height: 60px;
    font-family: sans-serif;
    font-style:normal;
    display: block;
    background-color: brown;
  }
  .item a :hover {
    background-color: red;
  }
  .item {
    float: left;
    height: 100%;
    width: 100px;
    text-align: center;
  }

  .navcenter {
    width: 1100px;
    margin: 0 auto;
    text-align: left;
  }
  .menu {
    height: 60px;
    background-color: black;
  }
  .buycar {
    height: 120px;
    background-color: #ffffff;
  }
  .side {
    margin-left: 10px;
    margin-right: 10px;
  }
 .top {
    font-size: 14px;
    height: 28px;
    line-height: 28px;
    font-family: sans-serif;
   padding: 0 0 0 2px;
 }
  .tomain {
    width: 1100px;
    height: 100%;
    margin: 0 auto;
    text-align: right;
  }

  .tomain span {

    float: left;
  }

  body {
    height: 100vh;
    margin: 0 0;
  }

  .header {
    height: 3%;
    width: 100%;
    background-color: green;
    position: fixed;
  }

  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    width: 100%;
    height: 100%
  }
</style>
