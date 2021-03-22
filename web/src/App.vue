<template>
  <!-- <HelloWorld msg="Hello Vue 3 + TypeScript + Vite" /> -->
  <div class="container">
    <img alt="Vue logo" src="./assets/logo.png" />

    <div
      style="
        background-color: #13ce66;
        padding: 10px;
        color: #fff;
        border-radius: 4px;
      "
    >
      ALL PASS {{ name }}
    </div>

    <table style="width: 100%; margin-top: 10px">
      <thead>
        <tr>
          <th>STATUS</th>
          <th>NAME</th>
          <th>LATEST CHECK</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>PASS</td>
          <td>TCP server</td>
          <td>2021-03-22 19:32:22</td>
        </tr>
        <tr>
          <td>PASS</td>
          <td>TCP server</td>
          <td>2021-03-22 19:32:22</td>
        </tr>
        <tr>
          <td>PASS</td>
          <td>TCP server</td>
          <td>2021-03-22 19:32:22</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts" setup></script>

<style>
.container {
  width: 960px;
  margin: 0 auto;
}

table {
  border-collapse: collapse;
}

tbody tr {
  border-bottom: 1px solid #e2e2e2;
}

tbody tr td {
  padding: 10px 0;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
