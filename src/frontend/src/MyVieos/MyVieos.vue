<template>
  <div id="app">
    <Menu />
    <AddNew />
    <ChoiseChannel :options="channels" @choise="ChoiseChannel" />
    <div id="charts" class="uk-container">
      <div v-if="data" class="uk-text-center">
        <h4>Choise channel</h4>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import * as am4core from "@amcharts/amcharts4/core";
import * as am4charts from "@amcharts/amcharts4/charts";
import ru_RU from "@amcharts/amcharts4/lang/ru_RU";
import material from "@amcharts/amcharts4/themes/material";
am4core.useTheme(material);

import Menu from "@/Components/Navbar.vue";
import AddNew from "@/Components/AddNew.vue";
import ChoiseChannel from "@/Components/ChoiseChannel.vue";
export default {
  name: "myvideos",
  components: { Menu, AddNew, ChoiseChannel },
  metaInfo: {
    title: "Videos charts",
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
      data: [],
      channels: [],
      charts: []
    };
  },
  created() {
    axios.get("/UserChannels").then(res => (this.channels = res.data));
    this.FetchData();
  },
  methods: {
    ChoiseChannel(chan) {
      this.charts.forEach(e => e.remove());
      this.charts = [];
      let payload = { chanId: chan };
      this.FetchData(payload);
    },
    FetchData(payload) {
      // fetch data and create charts in methods
      axios.get("/UserVideos", { params: payload }).then(res => {
        this.data = res.data ? res.data : [];
        this.CreateElems();
      });
    },
    CreateElems() {
      this.data.forEach((dataset, index) => {
        let chartname = `chart${index}`;
        let app = document.getElementById("charts");
        let container = document.createElement("div");
        container.classList.add("uk-flex-wrap-around");
        container.classList.add("uk-flex");
        this.charts.push(container);

        let textdiv = document.createElement("div");
        textdiv.classList.add("uk-width-1-3");
        textdiv.classList.add("uk-card");
        textdiv.classList.add("uk-card-body");
        let titleDiv = document.createElement("div");
        textdiv.appendChild(titleDiv);
        let videoHref = document.createElement("a");
        videoHref.setAttribute("href", dataset.URL);
        videoHref.setAttribute("target", "_blank");
        videoHref.innerHTML = dataset.Title;
        titleDiv.appendChild(videoHref);

        let chartdiv = document.createElement("div");
        chartdiv.classList.add("uk-width-1-1");
        chartdiv.classList.add("uk-card");
        chartdiv.classList.add("uk-card-body");
        chartdiv.classList.add("chart");
        chartdiv.setAttribute("id", chartname);

        app.appendChild(container);
        container.appendChild(textdiv);
        container.appendChild(chartdiv);
        this.CrateChart(chartname, dataset);
      });
    },
    CrateChart(elem, dataset) {
      // created charts for each vedio
      var chart = am4core.create(elem, am4charts.XYChart);
      chart.language.locale = ru_RU;
      chart.data = this.RebuildData(
        dataset.DateSlice,
        dataset.Views,
        dataset.Likes,
        dataset.Dislikes
      );
      var dateAxis = chart.xAxes.push(new am4charts.DateAxis());
      dateAxis.groupData = true;
      dateAxis.groupCount = 1;
      dateAxis.skipEmptyPeriods = true;
      var valueAxis = chart.yAxes.push(new am4charts.ValueAxis());
      valueAxis.tooltip.disabled = false;
      var series = chart.series.push(new am4charts.LineSeries());
      series.name = "Views";
      series.tooltipText = "{valueY.value}";
      series.dataFields.dateX = `date`;
      series.dataFields.valueY = `value`;
      series.strokeWidth = 5;

      var likes = chart.series.push(new am4charts.LineSeries());
      likes.name = "Likes";
      likes.tooltipText = "{valueY.value}";
      likes.dataFields.dateX = `date`;
      likes.dataFields.valueY = `likes`;
      likes.strokeWidth = 1;

      var dislikes = chart.series.push(new am4charts.LineSeries());
      dislikes.name = "Dislikes";
      dislikes.tooltipText = "{valueY.value}";
      dislikes.dataFields.dateX = `date`;
      dislikes.dataFields.valueY = `dislikes`;
      dislikes.strokeWidth = 1;

      chart.legend = new am4charts.Legend();
      chart.cursor = new am4charts.XYCursor();
      chart.cursor.xAxis = dateAxis;
      chart.exporting.menu = new am4core.ExportMenu();
      chart.exporting.menu.items = [
        {
          label: "...",
          menu: [
            // { type: "png", label: "PNG" },
            { type: "xlsx", label: "XLSX" }
          ]
        }
      ];
    },
    RebuildData(dates, values, valueLikes, valueDislikes) {
      // rebuild data for amchart
      let newData = [];
      dates.forEach((date, i) => {
        newData.push({
          date: new Date(Date.parse(date)),
          value: values[i],
          likes: valueLikes[i],
          dislikes: valueDislikes[i]
        });
      });
      return newData;
    }
  }
};
</script>

<style>
.chart {
  height: 30vh;
}
</style>