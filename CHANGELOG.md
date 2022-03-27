## v0.2.1 (2022-03-28)

### New feature:

- add --daemon flag([`ab6c4b0`](https://github.com/axetroy/watchdog/commit/ab6c4b056013a5d497e32c3fff33ce484b907dc4)) (by Axetroy)
- upgrade to go 1.18([`6b3b0bd`](https://github.com/axetroy/watchdog/commit/6b3b0bd2b746ed974f7029ba0be281864db24bb0)) (by Axetroy)

### Bugs fixed:

- **deps**: update golang.org/x/net commit hash to 3ad01bb (#155)([`0c6b9b0`](https://github.com/axetroy/watchdog/commit/0c6b9b02fdceb50158001b4aedf241c5fc577cc0)) (by renovate[bot])
- **deps**: update golang.org/x/net commit hash to cf34111 (#154)([`1106057`](https://github.com/axetroy/watchdog/commit/1106057432edab1670a686179ffd502fe59d4d2f)) (by renovate[bot])
- **deps**: update golang.org/x/crypto commit hash to 089bfa5 (#147)([`4010cdc`](https://github.com/axetroy/watchdog/commit/4010cdc1c457021941632dd809843ce60b2ea945)) (by renovate[bot])
- **deps**: update module github.com/fatih/color to v1.13.0 (#149)([`8d37dca`](https://github.com/axetroy/watchdog/commit/8d37dca22c9f884fca623c10c37c55859ba84418)) (by renovate[bot])
- **deps**: update module github.com/go-playground/validator/v10 to v10.9.0 (#146)([`57f3bf2`](https://github.com/axetroy/watchdog/commit/57f3bf21d1f6b52aafb3a67ba7e7f7509246b45b)) (by renovate[bot])
- **deps**: update golang.org/x/crypto commit hash to 84f3576 (#145)([`acddc04`](https://github.com/axetroy/watchdog/commit/acddc042590c22ef58fcf77230a112ad822b7924)) (by renovate[bot])

### Revert:

- revert [`ab6c4b0`](https://github.com/axetroy/watchdog/commit/ab6c4b056013a5d497e32c3fff33ce484b907dc4), feat: add --daemon flag([`0cef6ef`](https://github.com/axetroy/watchdog/commit/0cef6ef5b947cb2125a048c9712e58a8db76d50e))

## v0.2.0 (2021-09-21)

### New feature:

- 支持 yml 配置文件([`ea2170d`](https://github.com/axetroy/watchdog/commit/ea2170d9f4b2216237773badfc425f5504c0641f)) (by axetroy)
- 支持休息时段内，不发送通知([`874402c`](https://github.com/axetroy/watchdog/commit/874402c273901860ae666defd0ecd6676c952911)) (by axetroy)
- 支持每小时最多发送的通知数量，防止频率过高而打扰([`08c1d28`](https://github.com/axetroy/watchdog/commit/08c1d28790a2bfacaf319e29e72b8d593f7d07a6)) (by axetroy)

### Bugs fixed:

- **deps**: update module github.com/mhale/smtpd to v0.8.0 (#101)([`0463511`](https://github.com/axetroy/watchdog/commit/046351101766ad6f15e85741b03540b9c5372676)) (by renovate[bot])
- **deps**: update dependency echarts to v5.2.1 (#143)([`f7b0fb2`](https://github.com/axetroy/watchdog/commit/f7b0fb26e509f1b2e608e92b16f44bab36d74c43)) (by renovate[bot])
- http test([`ae4e462`](https://github.com/axetroy/watchdog/commit/ae4e4629fcb9c27a2d228401908573a6e856cf14)) (by Axetroy)
- lint([`e3e5367`](https://github.com/axetroy/watchdog/commit/e3e53672c42e0fe071c8be5ce5511461a31c2ebf)) (by Axetroy)
- **deps**: update dependency date-fns to v2.24.0 (#140)([`de70198`](https://github.com/axetroy/watchdog/commit/de70198cef2e6b5f495a8117bf89fae955edb910)) (by renovate[bot])
- **deps**: update module github.com/mitchellh/mapstructure to v1.4.2 (#137)([`bb638ae`](https://github.com/axetroy/watchdog/commit/bb638ae62343caca2f8349538e59a772a96052a8)) (by renovate[bot])
- **deps**: update dependency echarts to v5.2.0 (#131)([`05e88b9`](https://github.com/axetroy/watchdog/commit/05e88b98dc169afa28e6b19d8b572169fe679e24)) (by renovate[bot])
- **deps**: update golang.org/x/crypto commit hash to 32db794 (#124)([`da8ac21`](https://github.com/axetroy/watchdog/commit/da8ac21789043a83cb2c8f9f66f79d4ee36590fd)) (by renovate[bot])
- **deps**: update golang.org/x/crypto commit hash to 0a44fdf (#118)([`60a707c`](https://github.com/axetroy/watchdog/commit/60a707ced4d74781a803d0092b445bab44083efa)) (by renovate[bot])
- **deps**: update golang.org/x/crypto commit hash to 0ba0e8f (#117)([`1ad0a01`](https://github.com/axetroy/watchdog/commit/1ad0a01b80973644c25b68b34e5f91870a31386a)) (by renovate[bot])
- **deps**: update module github.com/go-playground/universal-translator to v0.18.0 (#114)([`14ae07a`](https://github.com/axetroy/watchdog/commit/14ae07a4e8840818320a9d75545af75e384689fa)) (by renovate[bot])
- **deps**: update dependency vue-echarts to v6.0.0 (#110)([`0c41e5d`](https://github.com/axetroy/watchdog/commit/0c41e5d326df0d1cf1c83919cd094d4192d7a44d)) (by renovate[bot])
- **deps**: update dependency date-fns to v2.23.0 (#106)([`7713756`](https://github.com/axetroy/watchdog/commit/7713756f130c6704aa68f95c2fa3800ec302d116)) (by renovate[bot])
- **deps**: update dependency vue to v3.1.5 (#103)([`b20f24c`](https://github.com/axetroy/watchdog/commit/b20f24cefa8d201870c29ac1d484a47cf6292997)) (by renovate[bot])
- **deps**: update module github.com/google/uuid to v1.3.0 (#98)([`ca39470`](https://github.com/axetroy/watchdog/commit/ca394704d1dbfdf29de0afe1220b45ae006e564c)) (by renovate[bot])
- **deps**: update module github.com/go-playground/validator/v10 to v10.7.0 (#60)([`069bd25`](https://github.com/axetroy/watchdog/commit/069bd257380ba345b59937de4d21a60a979abf42)) (by renovate[bot])
- **deps**: update golang.org/x/crypto commit hash to a769d52 (#95)([`d0b9182`](https://github.com/axetroy/watchdog/commit/d0b918286ef5acfc2a50f5c8c91a5cc3ce35f852)) (by renovate[bot])
- **deps**: update dependency vue-echarts to v6.0.0-rc.6 (#83)([`d582725`](https://github.com/axetroy/watchdog/commit/d5827252c6e99474ae8ab9d453c3bdb092cc298a)) (by renovate[bot])
- **deps**: update module github.com/gliderlabs/ssh to v0.3.3 (#91)([`8ca7e0d`](https://github.com/axetroy/watchdog/commit/8ca7e0d620a64a9d3cba95480e15003571b19ac7)) (by renovate[bot])
- **deps**: update golang.org/x/crypto commit hash to 5ff15b2 (#85)([`bf1ab26`](https://github.com/axetroy/watchdog/commit/bf1ab2697a7d0eaafaf65499fa8c3213cf0a6640)) (by renovate[bot])
- **deps**: update dependency echarts to v5.1.2 (#82)([`92ccf40`](https://github.com/axetroy/watchdog/commit/92ccf40a24bc6f048743ddf55ff6e8d6211e8a0a)) (by renovate[bot])
- **deps**: update golang.org/x/crypto commit hash to c07d793 (#69)([`a0fcfbd`](https://github.com/axetroy/watchdog/commit/a0fcfbdf0e960815a531a62cee0c82b7a8674b96)) (by renovate[bot])
- **deps**: update module github.com/fatih/color to v1.12.0 (#73)([`d31d806`](https://github.com/axetroy/watchdog/commit/d31d806fa8d36a15aba327f829b902aec510af21)) (by renovate[bot])
- **deps**: update dependency date-fns to v2.22.1 (#76)([`906b815`](https://github.com/axetroy/watchdog/commit/906b815f00d288e0cdea59b6608d75dd3d17b876)) (by renovate[bot])
- **deps**: update golang.org/x/crypto commit hash to cd7d49e (#58)([`3ae3b2f`](https://github.com/axetroy/watchdog/commit/3ae3b2fb668b885f9b165565b264ab431e6ba0af)) (by renovate[bot])
- **deps**: update module github.com/fatih/color to v1.11.0 (#68)([`75d297e`](https://github.com/axetroy/watchdog/commit/75d297eb78ee18e057840a39f8560fb4dd4deb48)) (by renovate[bot])
- **deps**: update dependency date-fns to v2.21.3 (#64)([`223b333`](https://github.com/axetroy/watchdog/commit/223b3330cdf2eb24b5fcf09745d4d2c12ff8717d)) (by renovate[bot])
- **deps**: update dependency echarts to v5.1.1 (#53)([`2d6d3d9`](https://github.com/axetroy/watchdog/commit/2d6d3d93b1982fc4fbcaa4a121d96fef045ea3c3)) (by renovate[bot])
- **deps**: update dependency date-fns to v2.21.2 (#59)([`0f56353`](https://github.com/axetroy/watchdog/commit/0f563535f4d825a8b7ca7523433fede199aded24)) (by renovate[bot])
- **deps**: update dependency vue-echarts to v6.0.0-rc.5 (#54)([`39e1751`](https://github.com/axetroy/watchdog/commit/39e1751fdd9af781b4cf5336ad9f6d9f69422808)) (by renovate[bot])
- **deps**: update golang.org/x/crypto commit hash to 83a5a9b (#49)([`45bc416`](https://github.com/axetroy/watchdog/commit/45bc416b4d04d491cfa4024d04f1acba56990b0e)) (by renovate[bot])
- **deps**: update golang.org/x/crypto commit hash to 5bf0f12 (#48)([`0bc6aaa`](https://github.com/axetroy/watchdog/commit/0bc6aaabf54a8e632132920666daa775d7484b7b)) (by renovate[bot])
- **deps**: update dependency echarts to v5.1.0 (#45)([`fb139b3`](https://github.com/axetroy/watchdog/commit/fb139b323837154b5dd6873b7d47021670a0a94d)) (by renovate[bot])
- **deps**: update dependency date-fns to v2.21.1 (#40)([`22b30e9`](https://github.com/axetroy/watchdog/commit/22b30e9d63f19749db4f57a07cd65ad0fe77396e)) (by renovate[bot])
- **deps**: update golang.org/x/crypto commit hash to 4f45737 (#44)([`8648b20`](https://github.com/axetroy/watchdog/commit/8648b20ee91922c8aa97257bf1e2841558ad26ea)) (by renovate[bot])
- **deps**: update module gopkg.in/yaml.v2 to v2.4.0 (#38)([`ea036b0`](https://github.com/axetroy/watchdog/commit/ea036b0e6c91d60b86830b1b1e71fb60cb78dad2)) (by renovate[bot])
- **deps**: update golang.org/x/net commit hash to afb366f (#36)([`e4d7989`](https://github.com/axetroy/watchdog/commit/e4d79892c407887772edb462d7c57712a4d8f688)) (by renovate[bot])
- **deps**: update module github.com/gookit/color to v1.4.2 (#34)([`71f9978`](https://github.com/axetroy/watchdog/commit/71f9978a6127dfb4d8c6aeff6704f238ad15d6a4)) (by renovate[bot])
- **deps**: update module github.com/go-playground/validator/v10 to v10.5.0 (#33)([`36c1fab`](https://github.com/axetroy/watchdog/commit/36c1fabef6a23f7f8013d6c924b486fab9e3bfc1)) (by renovate[bot])
- **deps**: update module github.com/go-playground/validator/v10 to v10.4.2 (#28)([`eebdbb7`](https://github.com/axetroy/watchdog/commit/eebdbb775229f4f97ade3ea57544c249c60b20a6)) (by renovate[bot])
- **deps**: update module github.com/gookit/color to v1.4.1 (#30)([`d8481a0`](https://github.com/axetroy/watchdog/commit/d8481a05360fb77856c481216fd18b98676a33e1)) (by renovate[bot])
- **deps**: update dependency date-fns to v2.20.1 (#32)([`af91d5e`](https://github.com/axetroy/watchdog/commit/af91d5e1060f7d90ae937cfc724e7a97c49b3cac)) (by renovate[bot])
- **deps**: update module github.com/gookit/color to v1.4.0([`0d08289`](https://github.com/axetroy/watchdog/commit/0d0828977d723d4293644dd940f4471021d7cd4b)) (by Renovate Bot)
- **deps**: update golang.org/x/net commit hash to a5a99cb([`bde60ca`](https://github.com/axetroy/watchdog/commit/bde60cacdd888d91157a5b43ae01022e10bfb603)) (by Renovate Bot)
- **deps**: update golang.org/x/net commit hash to 0fccb6f([`d621a2b`](https://github.com/axetroy/watchdog/commit/d621a2b2ae54ab3572e8be07d0932ea600c26975)) (by Renovate Bot)
- **deps**: update golang.org/x/net commit hash to 22f4162([`826276c`](https://github.com/axetroy/watchdog/commit/826276cb8f037fdf891f562a3046fa535e92cd45)) (by Renovate Bot)
- **deps**: update golang.org/x/net commit hash to df645c7([`9901635`](https://github.com/axetroy/watchdog/commit/990163527000664b1317c1d6fa5daaaa309503d2)) (by Renovate Bot)
- **deps**: update dependency vue-echarts to v6.0.0-rc.4([`14d12cd`](https://github.com/axetroy/watchdog/commit/14d12cde5fca087a61985d14854aa6797a2c5e23)) (by Renovate Bot)

### Revert:

- revert [`534085d`](https://github.com/axetroy/watchdog/commit/534085de4e9901b8f6e96cff2d310d15d714795c), test: fix([`1208e61`](https://github.com/axetroy/watchdog/commit/1208e6163a55deeaebb835b6b8ef7545268d7332))

## v0.1.1 (2021-03-28)

### Bugs fixed:

- **deps**: update module github.com/sirupsen/logrus to v1.8.1([`7669672`](https://github.com/axetroy/watchdog/commit/7669672b3f411947d1640b092c6bdc6d5dcdac6e)) (by Renovate Bot)

## v0.1.0 (2021-03-28)

### New feature:

- 输出日志到可切割的文件中([`4a67f5d`](https://github.com/axetroy/watchdog/commit/4a67f5d168157babf2cb8bedcefaab2b02ab1628)) (by axetroy)
- read PORT for env([`0591180`](https://github.com/axetroy/watchdog/commit/05911808a9d7b496db4b349e00a0693789e29486)) (by axetroy)
- 支持 smtp 的通知协议([`334e0c7`](https://github.com/axetroy/watchdog/commit/334e0c7d87fe3a68af1a93eacb0ef4f6a92c1ef1)) (by axetroy)
- 支持 SMTP 协议([`274be47`](https://github.com/axetroy/watchdog/commit/274be47589e284f9a04830b86ca830cf00a3b066)) (by axetroy)
- TCP 协议支持超时([`c47dc58`](https://github.com/axetroy/watchdog/commit/c47dc583797df22ae728a747696583711fe78fd9)) (by axetroy)
- HTTP 协议添加超时取消([`f754503`](https://github.com/axetroy/watchdog/commit/f7545038e4bcd223aa61bf94b4205f1775f0b1a9)) (by axetroy)
- 添加超时 context([`825fb9d`](https://github.com/axetroy/watchdog/commit/825fb9d10f2c30422a09fb756d9e7a257da174bf)) (by axetroy)
- SSH 协议用户名和密码不必填，可选项([`e7397ab`](https://github.com/axetroy/watchdog/commit/e7397abc01fe19d107969dfb991b3993e708eb67)) (by axetroy)
- 支持 SSH 协议([`846c093`](https://github.com/axetroy/watchdog/commit/846c093c0a12c39f48162215cb4c81e5aeebb706)) (by axetroy)
- 支持 ftp/sftp 协议([`a6c4126`](https://github.com/axetroy/watchdog/commit/a6c41263615895bf1a99f2cb253b02137472e6ca)) (by axetroy)
- 添加服务的速度曲线([`a5dddc5`](https://github.com/axetroy/watchdog/commit/a5dddc521f2919ff232074b795f12dccb0f81305)) (by axetroy)
- 添加持续时间([`1764373`](https://github.com/axetroy/watchdog/commit/17643739ad94b03e1badc3f83304c230a2e7de49)) (by axetroy)
- support webhook([`4e2a7e4`](https://github.com/axetroy/watchdog/commit/4e2a7e482e2749293f292ea8f90b7192b456bfc3)) (by axetroy)
- socket 推送([`ad7bc9e`](https://github.com/axetroy/watchdog/commit/ad7bc9e3a084a5da0a8d1bab98c516f1a9cc4107)) (by axetroy)
- 添加 Websocket([`5e8a3c9`](https://github.com/axetroy/watchdog/commit/5e8a3c9da373e7c51b9447c359be06d9e1c9da14)) (by axetroy)

### Bugs fixed:

- 完善配置文件的校验([`b54b862`](https://github.com/axetroy/watchdog/commit/b54b862a74d05ff70a12d9529b70eb0e0d5818f2)) (by axetroy)
- 完善配置文件的校验([`981c4ee`](https://github.com/axetroy/watchdog/commit/981c4ee96252139624454ed2d2945abe387aa044)) (by axetroy)
- 完善配置文件的校验([`e26ac11`](https://github.com/axetroy/watchdog/commit/e26ac11e1e6d433ffd6eefaf61539df93f811a39)) (by axetroy)
- **deps**: update golang.org/x/net commit hash to 61e0566([`c71b559`](https://github.com/axetroy/watchdog/commit/c71b559ddd3e497cb66ed74f66a035fe473a887c)) (by Renovate Bot)
- webhook test([`c6d96cf`](https://github.com/axetroy/watchdog/commit/c6d96cf8cc8117b89fef149d20a1808bceb5d99e)) (by axetroy)
- web([`46f7946`](https://github.com/axetroy/watchdog/commit/46f7946c47053d1f71e5b668bae90ed3724ef600)) (by axetroy)
- **deps**: update golang.org/x/crypto commit hash to 0c34fe9([`34446b2`](https://github.com/axetroy/watchdog/commit/34446b2ed9b67ad914dfc97b621bcd99bb0a309b)) (by Renovate Bot)
- test([`b8b0b0e`](https://github.com/axetroy/watchdog/commit/b8b0b0ee0bc7d7e31534853cf8a0bd8f9aafcecf)) (by axetroy)
- 修复 SSH 协议解析认证信息不正确([`07e1ab4`](https://github.com/axetroy/watchdog/commit/07e1ab4f590ae34e8616b59171ad877c0165b947)) (by axetroy)
- test([`efd8cdd`](https://github.com/axetroy/watchdog/commit/efd8cddbf2a1fc050b2975264f70962164758070)) (by axetroy)
- build for web([`03f0dca`](https://github.com/axetroy/watchdog/commit/03f0dcae72903cc4365b58d8a2d8e132e2d1a579)) (by axetroy)
- build([`da9339c`](https://github.com/axetroy/watchdog/commit/da9339c822862b8ee3b5004379885f05075c895e)) (by axetroy)
