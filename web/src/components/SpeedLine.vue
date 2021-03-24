<template>
  <v-chart
    style="height: 80px; max-width: 300px; margin: 0 auto"
    :option="option"
  />
</template>

<script lang="ts">
import { LineChart } from "echarts/charts";
import { GridComponent } from "echarts/components";
import { use } from "echarts/core";
import { CanvasRenderer } from "echarts/renderers";
import { defineComponent, PropType } from "vue";
import VChart, { THEME_KEY } from "vue-echarts";
import type { Service } from "../App.vue";

use([GridComponent, LineChart, CanvasRenderer]);

export default defineComponent({
  props: {
    dataSource: {
      type: Array as PropType<Service[]>,
      required: true,
      default: () => [],
    },
  },
  components: { VChart },
  provide: {
    [THEME_KEY]: "dark",
  },
  data() {
    return {};
  },
  computed: {
    option(): any {
      const dateList = this.dataSource.map((v) => v.updated_at);
      const valueList = this.dataSource.map((v) => v.duration);

      return {
        backgroundColor: "transparent", //背景色
        xAxis: {
          data: dateList,
          show: false,
        },
        yAxis: { show: false },
        grid: {
          show: true,
          borderColor: "transparent",
          backgroundColor: "transparent",
        },
        series: [
          {
            type: "line",
            showSymbol: false,
            data: valueList,
          },
        ],
      };
    },
  },
});
</script>
