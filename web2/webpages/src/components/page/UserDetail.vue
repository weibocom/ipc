<template>
  <div>
    <el-card style="margin-bottom: 10px;" :body-style="{ padding: '20px 20px 0px 20px' }">
      <div class="w-form" style="margin-left: 0px; width: 650px;">
        <el-form :model="userForm" label-width="80px">
          <el-form-item size="medium" label="用户ID" prop="uid">
            <el-input class="form-item-content" v-model="userForm.uid" :readonly="true"></el-input>
          </el-form-item>
          <el-form-item size="medium" label="用户平台" prop="company">
            <el-input class="form-item-content" v-model="userForm.company" :readonly="true"></el-input>
          </el-form-item>
          <el-form-item size="medium" label="创建时间" prop="created_at">
            <el-input class="form-item-content" v-model="userForm.created_at" :readonly="true"></el-input>
          </el-form-item>
          <el-form-item size="medium" label="发布内容" prop="post_count">
            <el-input class="form-item-content" v-model="userForm.post_count" :readonly="true"></el-input>
          </el-form-item>
          <el-form-item size="medium" label="私钥" prop="private_key">
            <el-input class="form-item-content" v-model="userForm.private_key" :readonly="true"></el-input>
          </el-form-item>
          <el-form-item size="medium" label="公钥" prop="public_key">
            <el-input class="form-item-content" v-model="userForm.public_key" :readonly="true"></el-input>
          </el-form-item>
        </el-form>
      </div>
    </el-card>
    <el-card :body-style="{ padding: '20px 20px 0px 20px' }">
      <!--user table-->
      <el-table :data="posts" style="width: 100%" ref="multipleTable" v-loading="loading">
        <el-table-column prop="mid" label="编号" width="180"></el-table-column>
        <el-table-column prop="created_at" label="上链时间" width="180"></el-table-column>
        <el-table-column prop="dna" label="哈希值"></el-table-column>
        <el-table-column prop="title" label="标题"></el-table-column>
        <el-table-column prop="show_content" label="内容">
          <template slot-scope="scope">
            <span class="content-span-button" @click="handleViewDetail(scope.row)">{{ scope.row.show_content }}</span>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination :current-page="curpage" @current-change="handleCurrentChange" layout="prev, pager, next" :page-size="10" :total="total">
        </el-pagination>
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
import { STORAGE_KEY, COMPANYS } from '@/utils/constants'

export default {
  mixins: [FormMixin],
  data() {
    return {
      loading: false,
      total: 0,
      curpage: 1,
      table_data: [],
      userid: '',
      company: '',

      userForm: {
        uid: '',
        company: '',
        post_count: 0,
        created_at: '',
        private_key: '',
        public_key: ''
      },

      contentDialogVisible: false,
      contentForm: {
        content: ''
      }
    }
  },
  created() {
    this.userid = this.$route.query.id
    this.company = this.$route.query.company
    this.init()
  },
  computed: {
    posts() {
      this.table_data.forEach(n => {
        n.created_at = new Date(n.created_at).Format('yyyy-MM-dd hh:mm:ss')

        if (n.content.length > 25) {
          n.show_content = n.content.substring(0, 25) + '...'
        } else {
          n.show_content = n.content
        }
      })
      return this.table_data
    }
  },
  methods: {
    init() {
      this.lookupUserDetail()
      this.lookupUserContent()
    },
    lookupUserDetail() {
      let params = {
        company: this.company
      }
      API.detailAccount(this.userid, params).then(
        res => {
          console.log(res)
          if (res.data === undefined || res.data.data === undefined) {
            this.$error('未获取到用户详情')
            return
          }
          let rdd = res.data.data
          this.userForm.uid = rdd.user.id
          this.userForm.created_at = new Date(rdd.user.created_at).Format(
            'yyyy-MM-dd hh:mm:ss'
          )
          COMPANYS.forEach(n => {
            if (rdd.user.company === n.value) {
              this.userForm.company = n.name
            }
          })
          this.userForm.post_count = rdd.user.post_count
          this.userForm.private_key = rdd.user.private_key
          this.userForm.public_key = rdd.user.public_key
        },
        _ => {}
      )
    },
    lookupUserContent() {
      let params = {
        company: this.company,
        uid: this.userid,
        page: this.curpage,
        pagesize: 10
      }

      this.loading = true
      API.lookupUserContent(params)
        .then(
          res => {
            console.log(res)

            this.loading = false
            this.table_data = []
            this.total = 0
            let data = res.data

            if (data === undefined || data.data === undefined) {
              this.$error('未获取到用户发布数据')
              return
            }
            data = data.data

            if (data.post_count > 10) {
              this.table_data = data.posts.slice(0, 10)
            } else {
              this.table_data = data.posts.slice(0, data.post_count)
            }
            this.total = data.post_count
          },
          _ => {
            this.loading = false
          }
        )
        .catch(_ => {
          this.loading = false
        })
    },
    handleCurrentChange(val) {
      this.curpage = val
      this.lookupUserContent()
    },
    handleViewDetail(row) {
      this.contentForm.content = row.content
      this.contentDialogVisible = true
    }
  }
}
</script>

<style scoped>
.content-span-button {
  cursor: pointer;
  color: #409eff;
}
</style>

<style>
</style>
