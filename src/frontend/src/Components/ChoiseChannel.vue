<template>
  <div class="uk-container uk-margin-top">
    <div class="uk-margin">
      <div class="uk-form-controls">
        <select class="uk-select" v-model="choise">
          <option v-for="(val, i) in options" :key="i" :value="val.id">{{val.name}}</option>
        </select>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "choisechan",
  props: ["options"],
  data() {
    return {
      choise: null
    };
  },
  watch: {
    choise() {
      this.$emit("choise", this.choise);
    }
  }
};
</script>
