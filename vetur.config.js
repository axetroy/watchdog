// vscode 的 vetur 扩展 配置 https://vuejs.github.io/vetur/guide/setup.html#advanced
// 如果你不使用 vscode 作为开发工具，那么这个文件雨你无瓜

module.exports = {
  settings: {
    "vetur.useWorkspaceDependencies": true,
    "vetur.experimental.templateInterpolationService": true,
  },
  projects: [
    {
      root: "./web",
      package: "./package.json",
      tsconfig: "./tsconfig.json",
      globalComponents: ["./src/components/**/*.vue"],
    },
  ],
};
