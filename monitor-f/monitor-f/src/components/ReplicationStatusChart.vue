<template>
    <div>
        <h1>Replication Status</h1>
        <canvas id="replicationStatusChart" ref="chart"></canvas>
    </div>
</template>
  
<script>
// import axios from "axios";
import { Bar } from "vue-chartjs";

export default {
    extends: Bar,
    data() {
        return {
            chartData: null,
        };
    },
    watch: {
        chartData: {
            deep: true,
            handler() {
                this.fetchDataAndUpdateChart();
            },
        },
    },

    async mounted() {
        this.fetchDataAndUpdateChart();

        // // Continuously fetch data and update the chart every 5 seconds
        // setInterval(() => {
        //     this.fetchDataAndUpdateChart();
        // }, 5000);
    },
    methods: {
        fetchDataAndUpdateChart() {
            const labels = this.chartData.map((record) => record.application_name);
            const replicationLags = this.chartData.map(
                (record) => parseInt(record.replication_lag)
            );

            this.renderChart(
                {
                    labels: labels,
                    datasets: [
                        {
                            label: "Replication Lag",
                            data: replicationLags,
                            backgroundColor: "rgba(75, 192, 192, 0.2)",
                            borderColor: "rgba(75, 192, 192, 1)",
                            borderWidth: 1,
                        },
                    ],
                },
                {
                    responsive: true,
                    maintainAspectRatio: false,
                    scales: {
                        y: {
                            beginAtZero: true,
                        },
                    },
                }
            );
        },
    },


};
</script>

<style scoped>
canvas {
    max-width: 800px;
    max-height: 400px;
}
</style>