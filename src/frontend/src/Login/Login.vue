<template>
  <div id="app">
    <div class="md-layout md-gutter md-alignment-center-center">
      <form class="md-layout-item-100" @submit.prevent="SendLogin()">
        <md-card class>
          <!-- <md-card-header>
            <div class="md-title">Users</div>
          </md-card-header>-->
          <md-card-content>
            <div class="md-layout md-gutter">
              <div class="md-layout-item md-small-size-100">
                <md-field :class="validate ? 'md-invalid' :''">
                  <label for="username">Login</label>
                  <md-input v-model="username" name="username" id="username" />
                  <!-- <span class="md-error" v-if="validate">incorrect login</span> -->
                </md-field>
                <md-field :class="validate ? 'md-invalid' :''">
                  <label for="password">Password</label>
                  <md-input v-model="password" name="password" id="password" type="password" />
                  <!-- <span class="md-error" v-if="validate">incorrect password</span> -->
                </md-field>
              </div>
            </div>
          </md-card-content>

          <md-card-actions>
            <md-button type="submit" class="md-primary">Create user</md-button>
          </md-card-actions>
        </md-card>
      </form>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "login",
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