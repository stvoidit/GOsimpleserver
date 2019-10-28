<template>
  <div id="app">
    <div id="chartdiv"></div>
  </div>
</template>

<script>
import axios from "axios";
import * as am4core from "@amcharts/amcharts4/core";
import * as am4charts from "@amcharts/amcharts4/charts";
import ru_RU from "@amcharts/amcharts4/lang/ru_RU";
import am4themes_material from "@amcharts/amcharts4/themes/animated";
am4core.useTheme(am4themes_material);
am4core.options.onlyShowOnViewport = true;
export default {
  name: "myvideos",
  metaInfo: {
    title: "Контроль Обращений",
    link: [
      {
        rel: "shortcut icon",
        href: "/static/favicon.ico"
      },
      {
        rel: "shortcut icon",
        href: "/static/favicon-16x16.png"
      },
      {
        rel: "shortcut icon",
        href: "/static/favicon-32x32.png"
      }
    ]
  },
  data() {
    return {
      RawData: []
    };
  },
  created() {
    axios.get("/UserVideos").then(res => {
      this.RawData = res.data;
      var chart = am4core.create("chartdiv", am4charts.XYChart);
      chart.language.locale = ru_RU;
      chart.height = am4core.percent(100);
      chart.width = am4core.percent(100);

      var dateAxis = chart.xAxes.push(new am4charts.DateAxis());
      dateAxis.groupData = true;

      dateAxis.groupIntervals.setAll([
        { timeUnit: "minute", count: 1 },
        { timeUnit: "minute", count: 20 },
        { timeUnit: "hour", count: 1 },
        { timeUnit: "hour", count: 24 },
        { timeUnit: "day", count: 1 },
        { timeUnit: "month", count: 1 },
        { timeUnit: "year", count: 1 }
      ]);
      dateAxis.groupCount = 1;
      dateAxis.skipEmptyPeriods = true;

      var valueAxis = chart.yAxes.push(new am4charts.ValueAxis());
      valueAxis.tooltip.disabled = false;

      chart.data = [];
      this.RebuildData.forEach((data, index) => {
        var series = chart.series.push(new am4charts.LineSeries());
        chart.data.push(...data.data);
        series.name = data.title;
        series.strokeWidth = 3;
        series.tooltipText = "{valueY.value}";
        series.dataFields.dateX = `date${index}`;
        series.dataFields.valueY = `value${index}`;
        // chart.scrollbarX.series.push(series);
      });

      chart.legend = new am4charts.Legend();
      chart.cursor = new am4charts.XYCursor();
      chart.cursor.xAxis = dateAxis;
      chart.scrollbarX = new am4charts.XYChartScrollbar();
    });
  },
  computed: {
    RebuildData() {
      let newData = [];
      this.RawData.forEach((element, index) => {
        let statObject = {};
        statObject.title = element.Title;
        statObject.data = [];
        element.Views.forEach((v, i) => {
          let obj = {};
          obj[`value${index}`] = v;
          statObject.data.push(obj);
        });
        element.DateSlice.forEach((v, i) => {
          statObject.data[i][`date${index}`] = new Date(Date.parse(v));
        });
        newData.push(statObject);
        return;
      });
      return newData;
    }
  }
};
</script>

<style>
#chartdiv {
  margin: 0 50px;
  height: 65vh;
}
</style>