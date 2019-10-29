<template>
  <div id="app" class="uk-container">
    <div class="uk-flex uk-flex-center">
      <form class="uk-flex uk-flex-column" @submit.prevent="SendLogin()">
        <div class="uk-margin">
          <label class="uk-form-label" for="username">Username</label>
          <div class="uk-form-controls">
            <input class="uk-input" type="text" name="username" id="username" v-model="username" />
          </div>
        </div>
        <div class="uk-inline">
          <label class="uk-form-label" for="password">Password</label>
          <div class="uk-form-controls">
            <input
              class="uk-input"
              type="password"
              name="password"
              id="password"
              v-model="password"
            />
          </div>
        </div>
        <div class="uk-margin uk-flex uk-flex-center">
          <button class="uk-button uk-button-primary uk-button-small" type="submit">login</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "login",
  metaInfo: {
    title: "Login",
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
      username: "",
      password: "",
      validate: false
    };
  },
  methods: {
    SendLogin() {
      axios
        .post("login", { username: this.username, password: this.password })
        .then(res => {
          if (res.data.status !== "ok") {
            this.validate = true;
          } else {
            this.validate = false;
            window.location.href = res.data.goto;
          }
        })
        .catch(() => {
          this.validate = true;
        });
    }
  }
};
</script>

<style>
#app {
  margin-top: 100px;
}
</style>