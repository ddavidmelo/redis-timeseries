<template>

    <v-card variant="outlined" class="mx-auto" max-width="700">
        <v-toolbar color="deep-purple accent-4" dense>

            <v-text-field v-model="redisKey" style="width:200px" class="pl-2" label="Redis TimeSeries Key" required
                hide-details="auto" variant="solo"></v-text-field>
            <v-btn @click="run">
                RUN
            </v-btn>
            <v-spacer></v-spacer>
            <v-container shrink style="width:150px">
                <v-select v-model="updateIntervalText" label="Refresh" variant="solo" hide-details="auto"
                    :items="updateIntervalOptions">
                </v-select>
            </v-container>


            <v-btn @click="toggleStartPause">
                {{ btnPauseText }}
            </v-btn>

        </v-toolbar>

        <VueApexCharts ref="realtimeChart" width="100%" :options="chartOptions" :series="series" />
    </v-card>
</template>
  
<script>
import VueApexCharts from "vue3-apexcharts";
import { TimeSeriesRequest } from "@/timeseries_pb";
import { TimeSeriesServiceClient } from "@/timeseries_grpc_web_pb";

export default {
    components: {
        VueApexCharts,
    },
    data() {
        return {
            btnPauseText: "Pause",
            redisKey: "system:CPUPercentage:0",
            updateIntervalF: null,
            updateInterval: 1000,
            updateIntervalText: "1 s",
            updateIntervalOptions: ["500 ms", "1 s", "30 s", "1 min", "5 min", "15 min"],
            chartOptions: {
                chart: {
                    id: 'area-datetime',
                    type: 'line',
                    toolbar: {
                        show: false
                    },
                    zoom: {
                        enabled: false
                    },
                    animations: {
                        enabled: true,
                        easing: 'linear',
                        dynamicAnimation: {
                            speed: 500
                        }
                    }
                },
                stroke: {
                    curve: 'smooth'
                },
                dataLabels: {
                    enabled: false
                },
                markers: {
                    size: 0,
                    style: 'hollow',
                },
                xaxis: {
                    type: 'datetime',
                    range: 60000,
                },
            },
            series: [],
        }
    },
    watch: {
        updateIntervalText() {
            const timeVar = this.updateIntervalText.split(' ');
            switch (timeVar[1]) {
                case 'ms':
                    this.updateInterval = timeVar[0] * 1;
                    break;
                case 's':
                    this.updateInterval = timeVar[0] * 1000;
                    break;
                case 'min':
                    this.updateInterval = timeVar[0] * 1000 * 60;
                    break;
            }
            clearInterval(this.updateIntervalF);
            this.keepUpdatingData();
        },
    },
    created() {
        this.series = [{
            name: this.redisKey,
            data: []
        }]
    },
    methods: {
        keepUpdatingData() {
            this.updateIntervalF = setInterval(() => {
                this.grpcRequest();
                let seriesLength = this.series[0].data.length
                if (seriesLength > 100) {
                    this.series[0].data.splice(0, (seriesLength - 100));
                    this.updateSeriesLine();
                }
            }, this.updateInterval);
        },
        updateSeriesLine() {
            if (this.series[0].data != null) {
                this.$refs.realtimeChart.updateSeries([{
                    data: this.series[0].data,
                }], true);
            }
        },
        toggleStartPause() {
            if (this.btnPauseText == "Pause") {
                this.btnPauseText = "Start";
                clearInterval(this.updateIntervalF);
            } else {
                this.btnPauseText = "Pause";
                this.keepUpdatingData();
            }
        },
        run() {
            clearInterval(this.updateIntervalF);
            this.series = [{
                name: this.redisKey,
                data: []
            }];
            this.keepUpdatingData();
            this.grpcRequest();
        },
        grpcRequest() {
            var fromTimestamp = 0
            if (this.series[0].data.length > 1) {
                fromTimestamp = this.series[0].data[this.series[0].data.length - 1][0] + 1;
            }

            var host = window.location.protocol + "//" + window.location.host;
            const client = new TimeSeriesServiceClient(host, null, null);

            let request = new TimeSeriesRequest();
            request.setFromTimestamp(fromTimestamp);
            request.setToTimestamp(Date.now());
            request.setKey(this.redisKey);

            client.getTimeSeries(request, {}, (err, response) => {
                // handle the response
                if (err) {
                    console.log(err);
                } else {
                    response.getDataList().forEach((n) => {
                        this.series[0].data.push(n.array);
                        // console.log(n.array)
                    });
                }
            });
        }
    },
};

</script>
<style>

</style>