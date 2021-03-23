<template>
  <!-- 1审核通过 2待审核 -->
  <div ref="btnSet" class="wrap">
    <el-button
      v-if="content.status == '2'"
      :disabled="disabled"
      size="mini"
      type="success"
      :class="{ active: disabled }"
      class="filter-item"
      @click="doAdopt()"
    >通过</el-button>
    <el-button
      v-if="content.status == '1' && content.recommend == 0"
      :class="{ active: disabled }"
      size="mini"
      type="warning"
      class="filter-item"
      @click="recommendText()"
    >推荐</el-button>
    <el-button
      v-if="content.status == '1' && content.recommend == 1"
      size="mini"
      type="warning"
      :class="{ active: disabled }"
      class="filter-item"
      @click="recommendfail()"
    >取消推荐</el-button>
    <el-button
      v-if="content.status == '2'"
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
import { articleAdopt, recommendAdopt } from '@/api/review/articleReview.js'
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
      disabled: false,
      superadmin: false,
      roleStatus: null
    }
  },
  updated() {
    if (this.content.status === '3') {
      this.disabled = true
    } else {
      this.disabled = false
    }
  },
  created() {
    // console.log(this.content.status, this.content.recommend, this.content.recommend)
    if (this.content.status === '3') {
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
        const id = this.content.id
        const isShow = this.content.status == '1'
        if (isShow) {
          recommendAdopt(id)
            .then((rep) => { })
            .then(() => {
              this.$message({
                title: '成功',
                message: '成功推荐',
                type: 'success'
              })
              this.content.recommend = 1
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

    // 取消推荐
    recommendfail() {
      this.$confirm('请再次确定是否取消推荐', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        const id = this.content.id
        const isShow = this.content.status == '1'
        if (isShow) {
          recommendAdopt(id)
            .then((rep) => { })
            .then(() => {
              this.$message({
                title: '成功',
                message: '推荐取消成功',
                type: 'success'
              })
              this.content.recommend = 0
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
        const id = parseInt(this.content.id)
        const statuss = parseInt(this.content.status)
        const isShow = statuss == '2'
        const status = 0
        const doAdoptData = {
          status,
          id
        }
        if (isShow) {
          // console.log(doAdoptData)
          articleAdopt(doAdoptData)
            .then((rep) => { })
            .then(() => {
              this.$message({
                title: '成功',
                message: '成功驳回',
                type: 'success'
              })
              this.content.status = '3'
              this.$emit('childEvent', this.content.status)
            })
            .catch(() => {
              this.$message({
                type: 'info',
                message: '已取消驳回'
              })
              this.content.status = '2'
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
        const id = parseInt(this.content.id)
        const statuss = parseInt(this.content.status)
        const isShow = statuss == '2'
        const status = 1
        const doAdoptData = {
          status,
          id
        }
        if (isShow) {
          articleAdopt(doAdoptData)
            .then((rep) => { })
            .then(() => {
              this.$message({
                type: 'success',
                message: '通过成功！'
              })
              this.content.status = '1'
              this.$emit('childEvent', this.content.status)
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
