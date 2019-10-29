<template>
  <div class="uk-container uk-margin-top">
    <div class="uk-text-center" uk-grid>
      <div class="uk-width-2-3">
        <input class="uk-input" type="text" placeholder="url on new video" v-model="url" />
      </div>
      <div class="uk-width-1-3">
        <button class="uk-button uk-button-primary" @click="AddUrl()">Button</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "addNew",
  data() {
    return {
      url: null
    };
  },
  methods: {
    AddUrl() {
      axios.post("/AddVideo", [{ url: this.url }]).then(res => {
        if (res.data.status === "add videos") {
          this.url = null;
        }
      });
    }
  }
};
</script>
