<template>
  <div id="app"></div>
</template>

<script>
import axios from "axios";
import * as am4core from "@amcharts/amcharts4/core";
import * as am4charts from "@amcharts/amcharts4/charts";
import ru_RU from "@amcharts/amcharts4/lang/ru_RU";
import material from "@amcharts/amcharts4/themes/material";
am4core.useTheme(material);
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
  created() {
    // fetch data and create charts in methods
    axios.get("/UserVideos").then(res => {
      let appdiv = document.getElementById("app");
      res.data.forEach((data, index) => {
        let chartname = `chart${index}`;
        let div = document.createElement("div");
        div.innerText = "123";
        div.setAttribute("id", chartname);
        div.classList.add("chart");
        appdiv.appendChild(div);
        this.CrateChart(chartname, data);
      });
    });
  },
  methods: {
    CrateChart(elem, dataset) {
      // created charts for each vedio
      var chart = am4core.create(elem, am4charts.XYChart);
      chart.language.locale = ru_RU;
      chart.data = this.RebuildData(dataset.DateSlice, dataset.Views);
      var dateAxis = chart.xAxes.push(new am4charts.DateAxis());
      dateAxis.groupData = true;
      dateAxis.groupCount = 1;
      dateAxis.skipEmptyPeriods = true;
      var valueAxis = chart.yAxes.push(new am4charts.ValueAxis());
      valueAxis.tooltip.disabled = false;
      var series = chart.series.push(new am4charts.LineSeries());
      series.name = dataset.Title;
      series.tooltipText = "{valueY.value}";
      series.dataFields.dateX = `date`;
      series.dataFields.valueY = `value`;
      series.strokeWidth = 5;

      chart.cursor = new am4charts.XYCursor();
      chart.cursor.xAxis = dateAxis;
      chart.legend = new am4charts.Legend();
    },
    RebuildData(dates, values) {
      // rebuild data for amchart
      let newData = [];
      dates.forEach((date, i) => {
        newData.push({
          date: new Date(Date.parse(date)),
          value: values[i]
        });
      });
      return newData;
    }
  }
};
</script>

<style>
#app {
  margin-top: 50px;
}
.chart {
  margin: 0 50px;
  height: 33vh;
}
</style>