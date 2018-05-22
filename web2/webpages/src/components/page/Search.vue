<template>
  <div>
    <el-card class="page-card">
      <div class="w-form">
        <el-form :model="searchForm" label-width="90px" label-position="left" :rules="searchFormRules" ref="searchForm">
          <el-form-item label="发布平台" prop="company">
            <el-select class="form-item-content" v-model="searchForm.company" placeholder="请选择发布平台">
              <el-option v-for="company in companys" :key="company.value" :label="company.name" :value="company.value">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="查询类型" prop="type">
            <el-select class="form-item-content" v-model="searchForm.type" placeholder="请选择发布平台">
              <el-option v-for="searchtype in searchtypes" :key="searchtype.value" :label="searchtype.label" :value="searchtype.value">
              </el-option>
            </el-select>
          </el-form-item>
          <div v-if="searchForm.type === 'url'">
            <el-form-item label="内容URL" prop="url">
              <el-input class="form-item-content" v-model="searchForm.url" placeholder="请输入URL" @blur="handleURL(searchForm.url)"></el-input>
            </el-form-item>
            <el-form-item label="用户ID" prop="uid">
              <el-input class="form-item-content" v-model="searchForm.uid" placeholder="请输入用户ID"></el-input>
            </el-form-item>
            <el-form-item label="消息ID" prop="mid">
              <el-input class="form-item-content" v-model="searchForm.mid" placeholder="请输入消息ID"></el-input>
            </el-form-item>
          </div>
          <div v-if="searchForm.type === 'dna'">
            <el-form-item label="哈希值" prop="dna">
              <el-input class="form-item-content" v-model="searchForm.dna" placeholder="请输入哈希值"></el-input>
            </el-form-item>
          </div>
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
        <el-form-item label="上链时间" prop="created_at">
          <el-input class="form-item-content" v-model="resultForm.created_at" :readonly="true"></el-input>
        </el-form-item>
        <el-form-item label="哈希值" prop="dna">
          <el-input class="form-item-content" v-model="resultForm.dna" :readonly="true"></el-input>
        </el-form-item>
        <el-form-item label="作者" prop="author">
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
import { STORAGE_KEY, COMPANYS } from '@/utils/constants'
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
      companys: COMPANYS,
      searchtypes: [
        { label: 'URL查询', value: 'url' },
        { label: '哈希值查询', value: 'dna' }
      ],
      searchForm: {
        type: 'url',
        url: '',
        uid: '',
        mid: '',
        dna: '',
        company: ''
      },
      searchFormRules: {
        company: [
          { required: true, message: '请选择发布平台', trigger: 'change' }
        ],
        uid: [{ required: true, validator: validateUID, trigger: 'blur' }],
        mid: [{ required: true, validator: validateMID, trigger: 'blur' }],
        dna: [{ required: true, message: '请输入哈希值', trigger: 'blur' }]
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
      this.searchForm.company = this.companys[0].value
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
          API.lookupContent(formData).then(
            res => {
              console.log(res)

              if (res.data.msg.toUpperCase() === 'OK') {
                let rdd = res.data.data
                if (rdd !== undefined && rdd.post !== undefined) {
                  this.resultForm.company = formData.company
                  this.resultForm.created_at = tool.checkData(
                    rdd.post.created_at
                  )
                  this.resultForm.created_at = new Date(
                    this.resultForm.created_at
                  ).Format('yyyy-MM-dd hh:mm:ss')
                  this.resultForm.author = tool.checkData(rdd.post.author)
                  this.resultForm.dna = tool.checkData(rdd.post.dna)
                  this.resultForm.content = tool.checkData(rdd.post.content)
                }
                this.resultDialogVisible = true
              } else {
                this.$error('未找到发布信息')
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
</style>
