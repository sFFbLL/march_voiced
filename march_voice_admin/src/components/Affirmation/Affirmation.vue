<template>
  <!-- 1审核通过 2待审核 -->
  <div ref="btnSet" class="wrap">
    <el-button
      v-if="data.status == '2' || data.status == '1'"
      size="mini"
      class="filter-item"
      type="primary"
      @click="toDetails()"
    >详情</el-button>
    <el-button
      v-if="data.status == '2'"
      :disabled="disabled"
      size="mini"
      type="success"
      :class="{ active: disabled }"
      class="filter-item"
      @click="doAdopt()"
    >通过</el-button>
    <el-button
      v-if="data.status == '1' && data.recommend == 1"
      :class="{ active: disabled }"
      size="mini"
      type="warning"
      class="filter-item"
      @click="recommendText()"
    >推荐</el-button>
    <el-button
      v-if="data.status == '1' && data.recommend == 0"
      size="mini"
      type="warning"
      :class="{ active: disabled }"
      class="filter-item"
      @click="recommendfail()"
    >取消推荐</el-button>
    <el-button
      v-if="data.status == '2'"
      :disabled="disabled"
      class="filter-item"
      type="danger"
      size="mini"
      :class="{ active: disabled }"
      @click="judgeText()"
    >驳回</el-button>
  </div>
</template>

<script>
// import crudMenu from "@/api/system/menu";
import CRUD, { crud } from '@crud/crud'
import { recommendAdopt, articleAdopt } from '@/api/review/articleReview.js'
// import Cookies from "js-cookie";
export default {
  name: 'Affirmation',
  mixins: [crud()],
  cruds() {
    return CRUD({ title: '组件', url: '' })
  },
  props: {
    data: {
      type: Object,
      required: true
    },
    content: {
      type: Object,
      required: true
    },
    recommend: {
      type: Number
      // required: true
    }
  },
  data() {
    return {
      jio: false,
      disabled: false,
      superadmin: false,
      roleStatus: null
    }
  },
  updated() {
    if (this.data.status === '3') {
      this.disabled = true
    } else {
      this.disabled = false
    }
  },
  created() {
    if (this.data.status === '3') {
      this.disabled = true
    } else {
      this.disabled = false
    }
  },
  mounted() {
    this.$emit('btnSetWidthEvent', this.$refs.btnSet.clientWidth)
  },

  methods: {
    // 推荐
    recommendText() {
      this.$confirm('请再次确定是否推荐', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        const id = this.data.id
        const isShow = this.data.status == '1'
        if (isShow) {
          recommendAdopt(id)
            .then((rep) => { })
            .then(() => {
              this.$message({
                title: '成功',
                message: '成功推荐',
                type: 'success'
              })
              this.data.recommend = 0
            })
            .catch(() => {
              this.$message({
                type: 'info',
                message: '推荐失败'
              })
            })
        }
      })
    },

    // 详情
    toDetails() {
      const id = this.data.id
      // const id = 1;
      const recommend = this.data.recommend
      const status = this.data.status
      const path = this.content.toDetialPath
      this.$router.push({
        path: `${path}`,
        query: {
          id: id,
          recommend: recommend,
          status: status
        }
      })
    },
    // 取消推荐
    recommendfail() {
      // var that = this;
      this.$confirm('请再次确定是否取消推荐', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        const id = this.data.id
        const isShow = this.data.status == '1'
        // console.log(tid)
        if (isShow) {
          recommendAdopt(id)
            .then((rep) => { })
            .then(() => {
              this.$message({
                title: '成功',
                message: '推荐取消成功',
                type: 'success'
              })
              this.data.recommend = 1
            })
            .catch(() => {
              this.$message({
                type: 'info',
                message: '推荐取消失败'
              })
            })
        }
      })
    },

    judgeText() {
      this.$confirm('请再次确定是否驳回', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        const id = parseInt(this.data.id)
        const statuss = parseInt(this.data.status)
        const isShow = statuss == '2'
        let status = 1
        const doAdoptData = {
          status,
          id
        }
        if (isShow) {
          console.log(doAdoptData)
          articleAdopt(doAdoptData)
            .then((rep) => { })
            .then(() => {
              this.$message({
                title: '成功',
                message: '成功驳回',
                type: 'success'
              })
              status = '3'
            })
            .catch(() => {
              this.$message({
                type: 'info',
                message: '已取消驳回'
              })
              status = '2'
            })
        }
      })
    },

    // 通过方法
    doAdopt() {
      this.$confirm('请再次确定是否通过', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        const id = parseInt(this.data.id)
        const statuss = parseInt(this.data.status)
        const isShow = statuss == '2'
        let status = 1
        const doAdoptData = {
          status,
          id
        }
        console.log(doAdoptData)
        console.log(isShow)
        if (isShow) {
          console.log(isShow)
          articleAdopt(doAdoptData)
            .then((rep) => { })
            .then(() => {
              this.$message({
                type: 'success',
                message: '通过成功！'
              })
              status = '1'
            })
            .catch(() => {
              this.$message({
                type: 'info',
                message: '通过失败'
              })
            })
        }
      })
    }
  }
}
</script>

<style scoped>
.wrap {
  float: left;
  display: flex;
  flex-flow: row nowrap;
  justify-content: center;
  align-items: center;
}
.filter-item {
  margin: auto 3px;
  padding: 5px 5px;
}
.active {
  background-color: #909399 !important;
  border-color: #909399 !important;
}
</style>
