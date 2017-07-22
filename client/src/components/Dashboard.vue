<template>
  <div class="hello">
    <h1 class="title">总收益</h1>
    <h2>{{ total }}元<button class="button">提现</button></h2>
    <h1 class="title">下属人数</h1>
    <div class="tags">
      <span class="tag" :style="{ backgroundColor: '#f50' }">一级</span>122人
      <span class="tag" :style="{ backgroundColor: '#2db7f5' }">二级</span>123人
      <span class="tag" :style="{ backgroundColor: '#87d068' }">三级</span>124人
      <span class="tag" :style="{ backgroundColor: '#108ee9' }">四级</span>25人
    </div>
    <h1 class="title">统计数据</h1>
    <div class="statistics">
      <LineChart :data="data"></LineChart>
    </div>
  </div>
</template>

<script>
  import LineChart from './LineChart';

  export default {
    name: 'index',
    components: {
      LineChart,
    },
    data() {
      return {
        total: 11245,
        people: 112,
        data: {
          labels: ['1', '2', '3', '4', '5', '6', '7', '8', '9', '10', '11', '12'],
          datasets: [{
            label: '收益',
            backgroundColor: 'rgb(255, 99, 132)',
            borderColor: 'rgb(255, 99, 132)',
            data: [401, 210, 431, 112, 314, 412, 423, 213, 413, 114, 124, 112],
            fill: false,
          }, {
            label: '人数',
            backgroundColor: 'rgb(54, 162, 235)',
            borderColor: 'rgb(54, 162, 235)',
            data: [40, 20, 43, 12, 34, 42, 43, 23, 43, 4, 24, 12],
            fill: false,
          }],
        },
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
  .tags {
    font-size: 0.7rem;
  }
  .tag {
    display: inline-block;
    line-height: 20px;
    height: 22px;
    padding: 0 8px;
    border-radius: 4px;
    border: 1px solid #e9e9e9;
    background: #f3f3f3;
    font-size: 12px;
    -webkit-transition: all .3s cubic-bezier(.78,.14,.15,.86);
    transition: all .3s cubic-bezier(.78,.14,.15,.86);
    opacity: 1;
    margin-right: 8px;
    cursor: pointer;
    color: white;
    white-space: nowrap;
  }
</style>
