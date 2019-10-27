<template>
  <div id="app">
    <div>hello</div>
    <div id="chartdiv"></div>
  </div>
</template>

<script>
import axios from "axios";
import * as am4core from "@amcharts/amcharts4/core";
import * as am4charts from "@amcharts/amcharts4/charts";
import am4themes_material from "@amcharts/amcharts4/themes/animated";
am4core.useTheme(am4themes_material); //am4themes_animated
am4core.options.onlyShowOnViewport = true;
export default {
  name: "myvideos",
  data() {
    return {
      RawData: []
    };
  },
  created() {
    axios.get("/UserVideos").then(res => {
      this.RawData = res.data;
      var chart = am4core.create("chartdiv", am4charts.XYChart);
      chart.height = am4core.percent(100);
      chart.width = am4core.percent(100);

      var dateAxis = chart.xAxes.push(new am4charts.DateAxis());
      dateAxis.baseInterval = {
        timeUnit: "hour",
        count: 1
      };
      dateAxis.skipEmptyPeriods = true;
      // dateAxis.renderer.minGridDistance = 100000;
      dateAxis.renderer.labels.template.location = 0.5;
      dateAxis.renderer.grid.template.location = 0;
      var valueAxis = chart.yAxes.push(new am4charts.ValueAxis());
      valueAxis.tooltip.disabled = false;
      valueAxis.renderer.minWidth = 100;
      dateAxis.renderer.grid.template.strokeOpacity = 0.07;
      valueAxis.renderer.grid.template.strokeOpacity = 0.07;

      chart.data = [];
      var scrollbarX = new am4charts.XYChartScrollbar();
      chart.scrollbarX = scrollbarX;
      this.RebuildData.forEach((data, index) => {
        chart.data.push(...data.data);
        var series = chart.series.push(new am4charts.LineSeries());
        series.strokeWidth = 5;
        series.name = data.title;
        series.dataFields.dateX = `date${index}`;
        series.dataFields.valueY = `value${index}`;
        series.tooltipText = "{name}: {valueY.value}";
        scrollbarX.series.push(series);
      });
      chart.legend = new am4charts.Legend();
      chart.cursor = new am4charts.XYCursor();
    });
  },
  mounted() {},
  computed: {
    RebuildData() {
      let newData = [];
      this.RawData.forEach((element, index) => {
        let statObject = {};
        statObject.title = element.Title;
        statObject.data = [];
        element.Views.forEach((v, i) => {
          let vname = `value${index}`;
          let o = {};
          o[vname] = v;
          statObject.data.push(o);
        });
        element.DateSlice.forEach((v, i) => {
          let vdate = `date${index}`;
          statObject.data[i][vdate] = new Date(Date.parse(v));
        });
        newData.push(statObject);
        return;
      });
      return newData;
    }
  },
  methods: {}
};
</script>

<style>
#chartdiv {
  margin: 0 50px;
  height: 65vh;
}
</style>