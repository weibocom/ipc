<template>
  <div class="header">
    <div class="logo">微博版权追溯系统</div>
    <div class="user-info">
      <span class="msg-time"><el-tag type="success">当前入链时间：{{msgtime}}</el-tag></span>
      <span class="user-count"><el-tag type="success">联盟内用户：{{usercount}}</el-tag></span>
      <el-dropdown trigger="click" @command="handleCommand">
        <span class="el-dropdown-link">
          <icon class="user-logo" name="user"></icon> {{username}}
        </span>
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item command="loginout">退出</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
  </div>
</template>

<script>
import API from '@/api/api'
import storage from '@/utils/storage'
import { STORAGE_KEY } from '@/utils/constants'
import { FormMixin } from 'components/mixins'

export default {
  mixins: [FormMixin],
  data() {
    return {
      name: '微博版权追溯',
      usercount: 0,
      msgtime: '',
      timerid: ''
    }
  },
  created() {
    this.lookupUserCount()
    this.lookupMsgTs()

    var self = this
    this.timerid = setInterval(refreshUserCount, 10000)
    function refreshUserCount() {
      self.lookupUserCount()
      self.lookupMsgTs()
    }
  },
  beforeDestroy() {
    clearInterval(this.timerid)
  },
  computed: {
    username() {
      let username = storage.get(STORAGE_KEY.USERNAME)
      return username ? username : this.name
    }
  },
  methods: {
    lookupUserCount() {
      API.lookupUserCount()
        .then(
          res => {
            if (
              res.data.msg.toUpperCase() === 'OK' &&
              res.data.data !== undefined &&
              res.data.data.count !== undefined
            ) {
              this.usercount = res.data.data.count
            }
          },
          _ => {}
        )
        .catch(_ => {})
    },
    lookupMsgTs() {
      API.lookupMsgTs()
        .then(
          res => {
            if (
              res.data.msg.toUpperCase() === 'OK' &&
              res.data.data !== undefined &&
              res.data.data.timestamp !== undefined
            ) {
              this.msgtime = new Date(res.data.data.timestamp).Format('yyyy-MM-dd hh:mm:ss')
            }
          },
          _ => {}
        )
        .catch(_ => {})
    },
    handleCommand(command) {
      if (command == 'loginout') {
        storage.set(STORAGE_KEY.AUTHED, false)
        storage.set(STORAGE_KEY.AUTHTOKEN, '')
        storage.set(STORAGE_KEY.USERNAME, '')
        this.$router.push('/login')
      }
    }
  }
}
</script>
<style scoped>
.header {
  position: relative;
  box-sizing: border-box;
  width: 100%;
  height: 50px;
  font-size: 16px;
  line-height: 50px;
  color: #fff;
}
.header .logo {
  float: left;
  width: 180px;
  text-align: center;
}
.header .user-info {
  float: right;
  padding-right: 50px;
  font-size: 16px;
  color: #fff;
}
.header .user-info .user-count {
  /* margin-right: 30px; */
  font-size: 14px;
  color: #fff;
}
.header .user-info .msg-time {
  /* margin-right: 30px; */
  font-size: 14px;
  color: #fff;
}
.header .user-info .el-dropdown-link {
  position: relative;
  display: inline-block;
  padding-left: 50px;
  color: #fff;
  cursor: pointer;
  vertical-align: middle;
}
.header .user-info .user-logo {
  position: absolute;
  left: 15px;
  top: 15px;
  width: 20px;
  height: 20px;
  font-size: 20px;
}
.el-dropdown-menu__item {
  text-align: center;
}
</style>
