<template>
  <div ref="btnSet" class="wrap">
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
import { articleAdopt } from '@/api/review/articleReview.js'
// import Cookies from "js-cookie";

export default {
  name: 'MarchAffirmation',
  mixins: [crud()],
  cruds() {
    return CRUD({ title: '组件', url: '' })
  },
  props: {
    data: {
      type: Object,
      required: true
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
    if (this.data.status === '3') {
      this.disabled = true
    } else {
      this.disabled = false
    }
  },
  created() {
    // console.log(this.data.status)
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
    judgeText() {
      this.$confirm('请再次确定是否驳回', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        console.log(this.data.userId)
        const id = parseInt(this.data.userId)
        const statuss = parseInt(this.data.status)
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
                title: '成功',
                message: '成功驳回',
                type: 'success'
              })
              this.data.status = '3'
            })
            .catch(() => {
              this.$message({
                type: 'info',
                message: '已取消驳回'
              })
              this.data.status = '2'
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
        const id = parseInt(this.data.userId)
        const statuss = parseInt(this.data.status)
        const isShow = statuss == '2'
        let status = 1
        const doAdoptData = {
          status,
          id
        }
        // console.log(doAdoptData)
        // console.log(isShow)
        if (isShow) {
          // console.log(isShow)
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
