<template>
  <div>
    <div v-show="isRegister">
      <h1>注册</h1>
      <p>立即支付45元即可成为《百仕汇和》 会员</p>
      <button class="button">立即支付</button>
    </div>
    <div v-show="!isRegister">
      <p>恭喜你注册成功！</p>
      <router-link to="/">进入首页</router-link>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'index',
    data() {
      return {
        isRegister: false,
      };
    },
  };
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .title {
    text-align: left;
    font-size: 1.5rem;
    color: #607D8B;
  }
  .button {
    font-size: 1.06rem;
    height: 1.6rem;
    line-height: .6rem;
    padding: 0.5rem 1rem;
    border-radius: 1.1rem;
    display: inline-block;
    outline: 0 none;
    -webkit-appearance: none;
    box-sizing: border-box;
    text-align: center;
    color: #000;
    background-color: #fff;
    border: 1px solid #ddd;
    margin: 0 10px;
  }
</style>
