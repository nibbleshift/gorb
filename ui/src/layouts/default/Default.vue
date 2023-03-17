<template>
  <v-app >
    <v-layout>
      <v-app-bar elevation="2" title="Go Regression and Benchmarks" ></v-app-bar>
      <v-navigation-drawer
        class="primary"
        theme="dark"
        permanent
      >
        <v-list nav>
          <v-list-item to="dashboard" prepend-icon="mdi-view-dashboard" title="Dashboard" value="dashboard"></v-list-item>
          <v-list-item to="tests" prepend-icon="mdi-list-status" title="Tests" value="tests"></v-list-item>
          <v-list-item to="benchmarks" prepend-icon="mdi-speedometer" title="Benchmarks" value="benchmarks"></v-list-item>
        </v-list>

        <template v-slot:append>
          <div class="pa-2">
            <v-icon icon="mdi-github" size="large" outlined color="white" flat></v-icon>
          </div>
        </template>
      </v-navigation-drawer>
      <v-main style="min-height: 300px;"><router-view /></v-main>
    </v-layout>
  </v-app>
</template>

<script lang="ts" setup>
</script>
