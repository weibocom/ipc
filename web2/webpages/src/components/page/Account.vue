<template>
  <div>
    <el-tabs v-model="activeTabName" type="border-card" @tab-click="handleTabClick">
      <el-tab-pane label="个人ID注册" name="registerTab">
        <div class="w-form">
          <el-form :model="regForm" label-width="90px" label-position="left" :rules="regFormRules" ref="regForm">
            <el-form-item label="发布平台" prop="company">
              <el-select class="form-item-content" v-model="regForm.company" placeholder="请选择发布平台">
                <el-option v-for="company in companys" :key="company.value" :label="company.name" :value="company.value">
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="用户ID" prop="uid">
              <el-input class="form-item-content" v-model="regForm.uid" placeholder="请输入用户ID"></el-input>
            </el-form-item>
            <el-form-item size="medium">
              <el-button class="form-item-content" type="primary" @click="submitRegister" :loading="regloading">账号注册</el-button>
            </el-form-item>
          </el-form>
        </div>
      </el-tab-pane>
      <el-tab-pane label="批量ID注册" name="batchRegisterTab">
        <div class="w-form">
          <el-form :model="batchRegForm" label-width="90px" label-position="left" :rules="batchRegFormRules" ref="batchRegForm">
            <el-form-item label="发布平台" prop="company">
              <el-select class="form-item-content" v-model="batchRegForm.company" placeholder="请选择发布平台">
                <el-option v-for="company in companys" :key="company.value" :label="company.name" :value="company.value">
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="导入文件" prop="filelist">
              <el-upload class="form-item-content" ref="upload" :limit="1" name="accounts_file" :data="batchRegData" :action="batchRegAction" :before-upload="beforeFileUpload" :on-exceed="handleFileExceed" :on-change="handleFileChange" :on-remove="handleFileRemove" :file-list="batchRegForm.filelist" :on-success="handleFileSuccess" :on-error="handleFileError" :auto-upload="false">
                <el-button slot="trigger" size="medium" type="primary">浏览</el-button>
                <div slot="tip" class="el-upload__tip">文件需为.csv格式，第一列为平台名，第二列为用户ID</div>
              </el-upload>
            </el-form-item>
            <el-form-item size="medium">
              <el-button class="form-item-content" type="primary" @click="submitUpload" :loading="batchloading">账号注册</el-button>
            </el-form-item>
          </el-form>
        </div>
      </el-tab-pane>
      <el-tab-pane label="联盟内用户查询" name="searchTab">
        <el-form :inline="true" :model="filters" style="margin-top: 10px;" @submit.native.prevent="handleUserSearch">
          <el-form-item size="medium" prop="company">
            <el-select class="form-item-content" style="margin-left: 0px;" v-model="filters.company" placeholder="请选择发布平台">
              <el-option v-for="company in companys" :key="company.value" :label="company.name" :value="company.value">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item size="medium">
            <el-input v-model="filters.id" prefix-icon="el-icon-search" placeholder="用户ID" @keyup.enter.native="handleSearch"></el-input>
          </el-form-item>
          <el-form-item size="medium">
            <el-button type="primary" icon="el-icon-search" v-on:click="handleUserSearch">查询</el-button>
          </el-form-item>
        </el-form>
        <!--user table-->
        <el-table :data="userdata" style="width: 100%" ref="userTable" v-loading="searchloading">
          <el-table-column prop="id" label="用户ID"></el-table-column>
          <el-table-column prop="created_at" label="入链时间"></el-table-column>
          <el-table-column align="center" label="详情" width="120" fixed="right">
            <template slot-scope="scope">
              <el-button size="mini" icon="el-icon-view" type="primary" @click="handleUserView(scope.$index, scope.row)">查看</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="pagination">
          <el-pagination :current-page="curpage" @current-change="handleCurrentChange" layout="prev, pager, next" :page-size="10" :total="total">
          </el-pagination>
        </div>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import API from '@/api/api.js'
import { FormMixin } from 'components/mixins'
import storage from '@/utils/storage'
import { STORAGE_KEY, COMPANYS } from '@/utils/constants'
import IBYTE from '@/utils/ibyte'
import validator from 'validator'
import jtfd from 'json-to-form-data'

export default {
  mixins: [FormMixin],
  data() {
    var validateUID = function(rules, value, callback) {
      if (value === '') {
        callback(new Error('请输入用户ID'))
      } else if (
        !validator.isInt(value, {
          allow_leading_zeroes: true,
          min: 0,
          max: 9999999999999
        })
      ) {
        callback(new Error('请输入正确的用户ID，数字类型，长度不超过13个字符'))
      } else {
        callback()
      }
    }
    return {
      activeTabName: 'registerTab',
      regloading: false,
      companys: COMPANYS,
      regForm: {
        uid: '',
        company: ''
      },
      regFormRules: {
        company: [
          { required: true, message: '请选择发布平台', trigger: 'change' }
        ],
        uid: [{ required: true, validator: validateUID, trigger: 'blur' }]
      },

      batchloading: false,
      batchRegAction: '/accounts?batch=true',
      batchRegData: {},
      batchRegForm: {
        company: '',
        filelist: []
      },
      batchRegFormRules: {
        company: [
          { required: true, message: '请选择发布平台', trigger: 'change' }
        ]
      },

      searchloading: false,
      filters: {
        company: '',
        id: ''
      },
      curpage: 1,
      total: 0,
      usertabledata: []
    }
  },
  created() {
    if (this.companys.length > 0) {
      this.regForm.company = this.companys[0].value
      this.batchRegForm.company = this.companys[0].value
      this.filters.company = this.companys[0].value
    }
  },
  computed: {
    userdata() {
      this.usertabledata.forEach(n => {
        n.created_at = new Date(n.created_at).Format('yyyy-MM-dd hh:mm:ss')
      })

      return this.usertabledata
    }
  },
  methods: {
    // 新增
    submitRegister() {
      this.validateForm('regForm')
        .then(valid => {
          let formData = Object.assign({}, this.regForm)
          formData = jtfd(formData)

          this.regloading = true
          API.createAccount(formData).then(
            res => {
              if (res.data.msg.toUpperCase() === 'OK') {
                this.$success('注册成功')
              } else {
                this.$error(res.data.msg)
              }
              this.regloading = false
            },
            _ => {
              this.regloading = false
            }
          )
        })
        .catch(valid => {})
    },
    submitUpload() {
      this.batchRegData.company = this.batchRegForm.company
      this.batchloading = true
      this.$refs.upload.submit()
    },
    beforeFileUpload(file) {
      const iscsv = file.type === 'text/csv'
      if (!iscsv) {
        this.$error('导入文件只能是 CSV 格式')
      }
      return iscsv
    },
    handleFileSuccess(res, file) {
      console.log(res)
      if (res.code === 200 && res.msg.toUpperCase() === 'OK') {
        this.$success('文件上传成功')
      } else {
        this.$error(res.msg)
      }
      this.batchloading = false
    },
    handleFileError(res, file) {
      this.$error('文件上传失败')
      this.batchloading = false
    },
    handleFileExceed(files, fileList) {
      this.$warning('当前限制选择 1 个文件')
    },
    handleFileRemove(file, filelist) {
      console.log(file)
    },
    handleFileChange(file, filelist) {
      console.log(file)
      if (file && file.raw && file.raw.type !== 'text/csv') {
        this.$error('文件：' + file.name + ' 不是csv文件')
      }
    },
    lookupAccount() {
      let params = {
        company: this.filters.company,
        page: this.curpage,
        pagesize: 10
      }
      if (this.filters.id) {
        params.uid = this.filters.id
      }

      this.searchloading = true
      API.lookupAccount(params)
        .then(
          res => {
            console.log(res)

            this.searchloading = false
            this.usertabledata = []
            this.total = 0
            let data = res.data

            if (data === undefined || data.data === undefined) {
              this.$error('未获取到用户数据')
              return
            }
            data = data.data

            if (data.count > 10) {
              this.usertabledata = data.users.slice(0, 10)
            } else {
              this.usertabledata = data.users.slice(0, data.count)
            }
            this.total = data.count
          },
          _ => {
            this.searchloading = false
          }
        )
        .catch(_ => {
          this.searchloading = false
        })
    },
    handleCurrentChange(value) {
      this.curpage = value
      this.lookupAccount()
    },
    handleUserSearch() {
      this.lookupAccount()
    },
    handleUserView(index, row) {
      this.$router.push({
        path: '/userdetail',
        query: { id: row.id, company: row.company }
      })
    },
    handleTabClick(tab, event) {
      if (tab.paneName === 'searchTab') {
        this.lookupAccount()
      }
    }
  }
}
</script>

<style scoped>
</style>

<style>
</style>
