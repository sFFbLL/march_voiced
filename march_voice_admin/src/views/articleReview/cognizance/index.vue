<template>
  <div class="app-container">
    <div class="head-container">
      <eHeader :permission="permission" />
      <crudOperation :permission="permission" :is-my-export="true" />
    </div>
    <!--表格渲染-->
    <el-table
      ref="table"
      v-loading="crud.loading"
      stripe
      lazy
      :data="unsigns"
      row-key="id"
      style="width: 100%"
    >
      <el-table-column prop="title" label="文章题目" />
      <el-table-column prop="tag" label="所属类别" />
      <el-table-column prop="nickname" label="作者昵称" />
      <el-table-column prop="createTime" label="时间" width="285">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.createTime) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        prop="status"
        class="colum-name"
        label="状态"
        width="180"
      >
        <template slot-scope="scope">
          <State :data="scope.row" />
        </template>
      </el-table-column>
      <!--  操作   -->
      <el-table-column
        label="操作"
        align="center"
        fixed="right"
        :width="operateWidth"
      >
        <template slot-scope="scope">
          <Affirmation
            :data="scope.row"
            :content="content"
            @btnSetWidthEvent="getBtnSetWidth"
          />
        </template>
      </el-table-column>
    </el-table>
    <!--分页组件-->
    <pagination />
  </div>
</template>

<script>
import eHeader from './module/header'
import CRUD, { presenter } from '@crud/crud'
import crudOperation from '@/components/AffirmationSearch/AffirmationSearch.vue'
import pagination from '@crud/Pagination'
import State from '@/components/State/State.vue'
import Affirmation from '@/components/Affirmation/Affirmation.vue'
export default {
  name: 'ArticleReview',
  components: { Affirmation, eHeader, pagination, State, crudOperation },
  cruds() {
    return CRUD({
      title: '文章审核',
      url: '/api/article/admin',
      sort: [{ column: 'level', asc: 'true' }],
      // 主页操作栏显示哪些按钮
      optShow: {
        add: false,
        edit: false,
        del: false,
        download: true,
        reset: true
      }
    })
  },
  mixins: [presenter()],
  data() {
    return {
      recommend: null,
      content: {
        recommend: '',
        toDetialPath: 'Detail',
        arrnew: [],
        arr: []
      },
      permission: {
        add: ['admin', 'job:add'],
        edit: ['admin', 'job:edit'],
        del: ['admin', 'job:del']
      },
      // 操作列宽度
      operateWidth: 150
    }
  },
  computed: {
    unsigns() {
      return this.crud.data.filter((item) => item.status === 1 || item.status === 2)
    }
  },

  methods: {
    getBtnSetWidth(width) {
      width += 15
      this.operateWidth = width > this.operateWidth ? width : this.operateWidth
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
// @import "@/assets/styles/point.scss";
::v-deep .el-input-number .el-input__inner {
  text-align: left;
}
::v-deep.el-table .cell {
  white-space: nowrap;
  text-overflow: ellipsis;
}
</style>
