<template>
  <div>
    <el-card class="page-card" :body-style="{ padding: '20px 20px 0px 20px' }">
      <div class="w-form">
        <el-form :model="searchForm" label-width="90px" label-position="left" :rules="searchFormRules" ref="searchForm">
          <el-form-item label="发布平台" prop="company">
            <el-select class="form-item-content" v-model="searchForm.company" placeholder="请选择发布平台">
              <el-option v-for="company in companys" :key="company.value" :label="company.name" :value="company.value">
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="查询类型" prop="queryType">
            <el-select class="form-item-content" v-model="searchForm.queryType" placeholder="请选择查询类型">
              <el-option v-for="searchtype in searchtypes" :key="searchtype.value" :label="searchtype.label" :value="searchtype.value">
              </el-option>
            </el-select>
          </el-form-item>
          <div v-if="searchForm.queryType === 'user'">
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
          <div v-if="searchForm.queryType === 'dna'">
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
    <el-card style="margin-top: 10px;" :body-style="{ padding: '20px 20px 0px 20px' }">
      <el-table :data="resulttabledata" style="width: 100%" ref="resultTable" :loading="tableloading">
        <el-table-column prop="similarity" label="相似度"></el-table-column>
        <el-table-column prop="mid" label="MID">
            <template slot-scope="scope">
              <span class="span-button" @click="handleViewPost(scope.row.author, scope.row.mid)">{{scope.row.mid}}</span>
            </template>
        </el-table-column>
        <el-table-column prop="dna" label="哈希值"></el-table-column>
        <el-table-column prop="author" label="作者">
          <template slot-scope="scope">
              <span class="span-button" @click="handleViewUserHome(scope.row.author)">{{scope.row.author}}</span>
            </template>
        </el-table-column>
        <el-table-column prop="title" label="标题"></el-table-column>
        <el-table-column prop="show_content" label="内容">
          <template slot-scope="scope">
            <span class="content-span-button" @click="handleViewDetail(scope.row)">{{ scope.row.show_content }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间"></el-table-column>
      </el-table>
      <div class="pagination">
        <el-button size="mini" type="primary" :disabled="prevenable" @click="handlePrevPage">上一页</el-button>
        <span style="margin-left: 10px; margin-right: 10px; font-size: 14px;">{{ '第 ' + curpage + ' 页' }}</span>
        <el-button size="mini" type="primary" :disabled="nextenable" @click="handleNextPage">下一页</el-button>
      </div>
    </el-card>
    <el-dialog title="详情" width="680px" class="q-form-dialog" :visible.sync="contentDialogVisible">
      <el-form :model="contentForm" label-width="90px" label-position="left">
        <el-form-item label="内容" prop="content">
          <el-input class="form-item-content" :autosize="{ minRows: 12 }" type="textarea" v-model="contentForm.content" :readonly="true"></el-input>
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
      tableloading: false,
      companys: COMPANYS,
      searchtypes: [
        { label: 'URL查询', value: 'user' },
        { label: '哈希值查询', value: 'dna' }
      ],
      searchForm: {
        queryType: 'user',
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
      },

      curpage: 1,
      curparams: {},
      resultdata: [],
      contentDialogVisible: false,
      contentForm: {
        content: ''
      }
    }
  },
  created() {
    if (this.companys.length > 0) {
      this.searchForm.company = this.companys[0].value
    }
  },
  computed: {
    resulttabledata() {
      this.resultdata.forEach(n => {
        n.created_at = new Date(n.created_at).Format('yyyy-MM-dd hh:mm:ss')
        if (n.content.length > 25) {
          n.show_content = n.content.substring(0, 25) + '...'
        } else {
          n.show_content = n.content
        }
      })
      return this.resultdata
    },
    prevenable() {
      if (this.curpage === 1) {
        return true
      }
      return false
    },
    nextenable() {
      if (this.resultdata.length < 10) {
        return true
      }
      return false
    }
  },
  methods: {
    lookupSimilar(curpage) {
      this.curparams.page = curpage
      this.curparams.pagesize = 10
      this.loading = true
      API.lookupSimilar(this.curparams).then(
        res => {
          console.log(res)
          if (res.data.msg.toUpperCase() === 'OK') {
            this.resultdata = res.data.data.posts
            this.curpage = curpage
          } else {
            this.$error('未找到发布信息')
          }
          this.loading = false
        },
        _ => {
          this.loading = false
        }
      )
    },
    // 查询
    submitSearch() {
      this.validateForm('searchForm')
        .then(valid => {
          this.curparams = Object.assign({}, this.searchForm)
          this.lookupSimilar(1)
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
    },
    handleViewDetail(row) {
      this.contentForm.content = row.content
      this.contentDialogVisible = true
    },
    handlePrevPage() {
      this.lookupSimilar(this.curpage - 1)
    },
    handleNextPage() {
      this.lookupSimilar(this.curpage + 1)
    },
    handleViewPost(id, mid) {
      let obj = tool.encodeBase62(mid + '')
      window.open('https://weibo.com/' + id + '/' + obj)
    },
    handleViewUserHome(id) {
      window.open('https://weibo.com/' + id)
    }
  }
}
</script>

<style scoped>
.content-span-button {
  cursor: pointer;
  color: #409eff;
}
.span-button {
  cursor: pointer;
  color: #409eff;
}
</style>

<style>
</style>
