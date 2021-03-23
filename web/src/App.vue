<template>
  <div class="container">
    <img alt="Logo" src="./assets/logo.svg" />

    <div
      style="padding: 10px; color: #fff; border-radius: 4px"
      :class="{
        'bg-success': services.every((v) => !v.error),
        'bg-warning': services.some((v) => v.error),
        'bg-error': services.every((v) => v.error),
      }"
    >
      ALL PASS
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
        <tr v-for="v in services" :key="v.name">
          <td>
            <img
              v-if="v.error"
              style="widht: 30px; height: 30px"
              src="./assets/error.svg"
            />
            <img
              v-else
              style="widht: 30px; height: 30px"
              src="./assets/check.svg"
            />
          </td>
          <td>{{ v.name }}</td>
          <td>{{ formatDate(v.updated_at) }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { format } from "date-fns";

interface Message<T = unknown> {
  event: string;
  payload: T;
}

interface Service {
  name: string; // 服务名称
  error?: string; // 错误信息
  updated_at: string; // 更新日期
}

export default defineComponent({
  data: () => {
    const state: {
      services: Service[];
      ws?: WebSocket | null;
    } = {
      services: [],
      ws: null,
    };

    return state;
  },
  methods: {
    formatDate(val: string) {
      return format(new Date(val), "yyyy-MM-dd HH:mm:ss");
    },
    updateService(s: Service) {
      const service = this.services.find((v) => v.name === s.name);

      if (service) {
        for (const attr in s) {
          // @ts-expect-error ignore
          service[attr] = s[attr];
        }
      } else {
        this.services.push(s);
      }
    },
    connect() {
      if (this.ws) {
        this.ws?.close();
        this.ws = null;
      }
      const ws = new WebSocket("ws://localhost:9999/api/ws");

      this.ws = ws;

      ws.onopen = () => {
        console.log("Websocket 已连接");
      };

      ws.onclose = (event) => {
        console.log(
          "Socket is closed. Reconnect will be attempted in 1 second.",
          event.reason
        );
        setTimeout(() => {
          this.connect();
        }, 1000);
      };

      ws.onmessage = (event) => {
        const data = JSON.parse(event.data) as Message;

        switch (data.event) {
          case "init":
            {
              const payload = (data as Message<Service[]>).payload;
              for (const p of payload) {
                this.updateService(p);
              }
            }
            break;
          case "update":
            {
              const payload = (data as Message<Service>).payload;
              this.updateService(payload);
            }
            break;
        }
      };
    },
  },
  mounted() {
    this.connect();
  },
  unmounted() {},
});
</script>

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

.bg-success {
  background-color: #13ce66;
}

.bg-warning {
  background-color: #ebdd65;
}

.bg-error {
  background-color: #f44336;
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
