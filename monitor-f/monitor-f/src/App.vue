<template>
  <div id="app">
    <ReplicationStatusChart />
  </div>
</template>

<script>
import ReplicationStatusChart from "./components/ReplicationStatusChart.vue";

export default {
  name: "App",
  components: {
    ReplicationStatusChart,
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
