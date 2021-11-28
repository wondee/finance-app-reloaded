<template>
<layout :user="user">
  <router-view v-if="user"/>
  <v-container>
    <v-row>
      <v-col>
        <h1>Loading User...</h1> 
      </v-col>
    </v-row>
  </v-container>
</layout>
</template>
<script>
import VueRouter from "vue-router";
import Layout from "./Layout";

const Overview = () => import("./components/overview/Overview.vue");
const FixedCosts = () => import("./components/FixedCosts.vue");
const SpecialCosts = () => import("./components/SpecialCosts.vue");

const routes = [
  { path: "/", component: Overview },
  { path: "/fixedcosts", component: FixedCosts },
  { path: "/specialcosts", component: SpecialCosts }
];

const router = new VueRouter({
  routes
});

export default {
  router,
  components: {
    Layout
  },
  data() {
    return {
      user: null
    }
  },
  created: async function() {
    const response = await fetch("/api/user");

    if (response.status === 404) {
      const response = await fetch("/api/user", {
        method: "PUT"
      })

      if (!response.ok) {
        throw new Error(response)
      }
    }

    const userDetailsResponse = await fetch("/.auth/me");

    if (!userDetailsResponse.ok) {
      throw new Error(userDetailsResponse) 
    }

    const authData = await userDetailsResponse.json()

    const name = authData[0].user_claims.find(
      elem => elem.typ === "name"
    )

    this.user = {
      id: authData.user_id,
      name: name.val
    }

  }
};
</script>
