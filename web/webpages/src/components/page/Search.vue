<template>
  <div>
    <el-card class="page-card">
      <div class="w-form">
        <el-form :model="searchForm" label-width="90px" label-position="left" :rules="searchFormRules" ref="searchForm">
          <el-form-item label="发布平台" prop="company">
            <el-select class="form-item-content" v-model="searchForm.company" placeholder="请选择发布平台">
              <el-option v-for="company in companys" :key="company" :label="company" :value="company">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="内容URL" prop="url">
            <el-input class="form-item-content" v-model="searchForm.url" placeholder="请输入URL" :disabled="searchForm.company!=='weibo'" @blur="handleURL(searchForm.url)"></el-input>
          </el-form-item>
          <el-form-item label="用户ID" prop="uid">
            <el-input class="form-item-content" v-model="searchForm.uid" placeholder="请输入用户ID"></el-input>
          </el-form-item>
          <el-form-item label="消息ID" prop="mid">
            <el-input class="form-item-content" v-model="searchForm.mid" placeholder="请输入消息ID"></el-input>
          </el-form-item>
          <el-form-item size="medium">
            <el-button class="form-item-content" type="primary" @click="submitSearch" :loading="loading">查询</el-button>
          </el-form-item>
        </el-form>
      </div>
    </el-card>
    <el-dialog title="查询结果" width="680px" class="q-form-dialog" :visible.sync="resultDialogVisible">
      <el-form :model="resultForm" label-width="90px" label-position="left">
        <el-form-item label="发布平台" prop="company">
          <el-input class="form-item-content" v-model="resultForm.company" :readonly="true"></el-input>
        </el-form-item>
        <el-form-item label="上链时间" prop="blockcreated">
          <el-input class="form-item-content" v-model="resultForm.blockcreated" :readonly="true"></el-input>
        </el-form-item>
        <el-form-item label="消息ID" prop="mid">
          <el-input class="form-item-content" v-model="resultForm.mid" :readonly="true"></el-input>
        </el-form-item>
        <el-form-item label="用户ID" prop="author">
          <el-input class="form-item-content" v-model="resultForm.author" :readonly="true"></el-input>
        </el-form-item>
        <el-form-item label="内容" prop="content">
          <el-input class="form-item-content" :autosize="{ minRows: 6 }" type="textarea" v-model="resultForm.content" :readonly="true"></el-input>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
import API from '@/api/api.js'
import { FormMixin } from 'components/mixins'
import storage from '@/utils/storage'
import tool from '@/utils/function'
import { STORAGE_KEY } from '@/utils/constants'
import IBYTE from '@/utils/ibyte'
import validator from 'validator'

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
    var validateMID = function(rules, value, callback) {
      if (value === '') {
        callback(new Error('请输入消息ID'))
      } else if (
        !validator.isInt(value, {
          allow_leading_zeroes: true,
          min: 0,
          max: 9999999999999999
        })
      ) {
        callback(new Error('请输入正确的消息ID，数字类型，长度不超过16个字符'))
      } else {
        callback()
      }
    }
    return {
      loading: false,
      companys: ['weibo', 'zhihu', 'facebook'],
      searchForm: {
        url: '',
        uid: '',
        mid: '',
        company: ''
      },
      searchFormRules: {
        company: [
          { required: true, message: '请选择发布平台', trigger: 'change' }
        ],
        uid: [{ required: true, validator: validateUID, trigger: 'blur' }],
        mid: [{ required: true, validator: validateMID, trigger: 'blur' }]
      },
      resultDialogVisible: false,
      resultForm: {
        company: '',
        blockcreated: '',
        mid: '',
        author: '',
        meta: '',
        content: ''
      }
    }
  },
  created() {
    if (this.companys.length > 0) {
      this.searchForm.company = this.companys[0]
    }
  },
  computed: {},
  methods: {
    // 查询
    submitSearch() {
      this.validateForm('searchForm')
        .then(valid => {
          let formData = Object.assign({}, this.searchForm)
          formData.action = 'search'
          this.loading = true
          API.search(formData).then(
            res => {
              if (res.data.msg.toUpperCase() === 'OK') {
                if (
                  res.data.data !== undefined &&
                  res.data.data.blockInfo !== undefined
                ) {
                  this.resultForm.company = formData.company
                  this.resultForm.blockcreated = tool.checkData(
                    res.data.data.blockInfo.blockCreateTime
                  )
                  this.resultForm.mid = formData.mid
                  this.resultForm.author = tool.checkData(
                    res.data.data.blockInfo.authorName
                  )
                  this.resultForm.author = tool.rmUIDPrefix(
                    this.resultForm.author
                  )
                  this.resultForm.content = tool.checkData(
                    res.data.data.blockInfo.content
                  )
                }
                this.resultDialogVisible = true
              } else {
                this.$error(res.data.msg)
              }
              this.loading = false
            },
            _ => {
              this.loading = false
            }
          )
        })
        .catch(valid => {})
    },
    handleURL(url) {
      if (url === '') {
        return
      }
      if (
        !validator.isURL(url, {
          protocols: ['http', 'https'],
          require_protocol: true,
          require_host: true,
          require_valid_protocol: true
        })
      ) {
        this.$warning('请输入合法的URL')
        return
      }
      // parse uid and mid from url, fill UID and MID input
      let obj = tool.uidmidFromURL(url)
      if (obj !== undefined && obj.uid !== undefined && obj.mid !== undefined) {
        this.searchForm.uid = obj.uid
        this.searchForm.mid = obj.mid
      }
    }
  }
}
</script>

<style scoped>

</style>

<style>
.page-card .el-card__header {
  background-color: rgb(238, 238, 238);
  color: rgba(0, 0, 0, 0.4);
  font-family: 'Roboto', 'sans-serif';
}
.page-card .el-table__header th {
  background-color: #fff;
}
</style>
