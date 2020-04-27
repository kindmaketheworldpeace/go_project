<template>
  <div id="app">
    <div class="header">
      <div class="tomain"><span class="top">您好, 欢迎光临!   </span>
        <font-awesome-icon icon="user-circle"/>
        <a class="top"> 登录</a><b class="side">|</b>
        <font-awesome-icon icon="check-square"/>
        <a class="top">注册</a>
      </div>
      <div style="    background-color: white;">
        <div class="buycar">
          <img src="../static/image/logo.png"></img>
        </div>
      </div>
      <div class="menu">
        <div class="navcenter">
          <div class="item">
            <a class="menu-item">
              <i style="font-style: normal">首页</i></a>
          </div>
          <div class="item">
            <a class="menu-item">
              <i style="font-style: normal">办公用品</i></a>
          </div>
          <div class="item">
            <a class="menu-item">
              <i style="font-style: normal">笔记本</i></a>
          </div>
          <div class="item">
            <a class="menu-item">
              <i style="font-style: normal">平板电脑</i></a>
          </div>
          <div class="item">
            <a class="menu-item">
              <i style="font-style: normal">台式机</i></a>
          </div>
        </div>
      </div>
      <div style="background-image: url('../static/image/pad.jpg');height: 380px">
      </div>
    </div>
    <div style="margin: 0 auto">
      <div style="margin: 0 auto;width: 1600px;">
        <div
          style="display: inline-block;width: 350px;background-image: url('../static/image/tpad (1).jpg');height: 400px;margin: 30px 30px">

        </div>
        <div
          style="display: inline-block;width: 350px;background-image: url('../static/image/tpad (2).jpg');height: 400px;margin: 30px 30px">

        </div>
        <div
          style="display: inline-block;width: 350px;background-image: url('../static/image/tpad (3).jpg');height: 400px;margin: 30px 30px">

        </div>
      </div>
      <div style="margin: 0 auto;width: 1180px;">
        <div style="border-bottom-style: solid; border-bottom-width:2px; text-align: left;border-bottom-color: black">
          <span style="color: #ff1639;font-size: 25px;">热门产品</span>

        </div>
        <div>
          <div class="goods-item">
            <img width="250" height="250" src="/static/image/computer1.jpg">
            <div><span class="price">￥6000</span></div>
            <div><span class="goods-name">惠普战神</span></div>

          </div>
          <div class="goods-item">
            <img width="250" height="250" src="/static/image/computer2.jpg">
            <div><span class="price">￥6000</span></div>
            <div><span class="goods-name">惠普战神</span></div>

          </div>
          <div class="goods-item">
            <img width="250" height="250" src="/static/image/computer3.jpg">
            <div><span class="price">￥6000</span></div>
            <div><span class="goods-name">惠普战神</span></div>
          </div>
          <div class="goods-item">
            <img width="250" height="250" src="/static/image/computer4.jpg">
            <div><span class="price">￥6000</span></div>
            <div><span class="goods-name">惠普战神</span></div>
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
  .goods-item {
    display: inline-block;
    margin: 20px 15px;
    border-color: #d4d1d1;
    border-style: solid;
    border-width: 1px;
    padding: 5px 3px;

  }

  .goods-item .price {
    color: #ff1639;
  }

  .goods-item .goods-name {
    color: dimgray;
    font-family: sans-serif;
  }

  .menu-item {
    color: white;
    line-height: 60px;
    font-family: sans-serif;
    font-style: normal;
    display: block;
  }

  .item :hover {
    background-color: #ff1639;
  }

  .item {
    float: left;
    height: 100%;
    width: 200px;
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
    width: 1100px;
    height: 50px;
    background-color: #ffffff;
    text-align: left;
    margin: 0 auto;
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
    -webkit-animation: top 10s linear infinite;
    animation: top 10s linear infinite;
  }

  @keyframes top {
    0% {
      transform: translateX(-100);
    }
    100% {
      transform: translateX(800px);
    }
  }

  .tomain {
    width: 1100px;
    height: 3%;
    margin: 0 auto;
    text-align: right;
  }

  .tomain span {

    float: left;
  }

  body {
    margin: 0 0;
  }

  .header {

    width: 100%;
    background-color: #c8c8c8;
  }

  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    width: 100%;
  }
</style>
