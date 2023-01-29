<template>
  <div class="box">
    <h2 class="title is-5">TimeSeries Chart</h2>
    <time-series-chart></time-series-chart>
  </div>
</template>

<script>
import TimeSeriesChart from './components/TimeSeriesChart';

export default {
  components: {
    TimeSeriesChart,
  },
};

</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
