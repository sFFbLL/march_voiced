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
      :data="tableData"
      row-key="id"
      style="width: 100%"
    >
      <el-table-column prop="articleTitle" label="文章题目" />
      <el-table-column prop="type" label="所属类别" />
      <el-table-column prop="authorNickname" label="作者昵称" />
      <el-table-column prop="time" label="时间" width="85" />
      <!-- <el-table-column prop="objectState" label="状态" /> -->
      <el-table-column
        prop="statee"
        class="colum-name"
        fixed="right"
        label="状态"
        width="130"
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
      />
    </el-table>
    <!--分页组件-->
    <pagination />
  </div>
</template>

<script>
import eHeader from './module/header'
// import { validatePoint, validateSelectOptionId } from "@/utils/validate.js";
import CRUD, { presenter } from '@crud/crud'
import crudOperation from '@/components/AffirmationSearch/AffirmationSearch.vue'
import pagination from '@crud/Pagination'
import State from './module/State.vue'
// import Affirmation from "@/components/Affirmation/Affirmation.vue";

export default {
  name: 'ArticleReview',
  components: { eHeader, pagination, State, crudOperation },
  cruds() {
    return CRUD({
      title: '文章审核',
      // url: "api/expertLectures/condition-query-sel",
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
      tableData: [
        {
          articleTitle: '2016-05-02',
          type: '王小虎',
          authorNickname: '上海市普陀区金沙江路 1518 弄',
          time: '上海市普陀区金沙江路 1518 弄',
          objectState: 1
        },
        {
          articleTitle: '2016-05-04',
          type: '王小虎',
          authorNickname: '上海市普陀区金沙江路 1518 弄',
          time: '上海市普陀区金沙江路 1518 弄',
          objectState: 2
        },
        {
          articleTitle: '2016-05-01',
          type: '王小虎',
          authorNickname: '上海市普陀区金沙江路 1518 弄',
          time: '上海市普陀区金沙江路 1518 弄',
          objectState: 2
        },
        {
          articleTitle: '2016-05-03',
          type: '王小虎',
          authorNickname: '上海市普陀区金沙江路 1518 弄',
          time: '上海市普陀区金沙江路 1518 弄',
          objectState: 3
        }
      ],
      content: {
        // identity: "expert_lectures",
        toDetialPath: 'details'
      },
      permission: {
        add: ['admin', 'job:add'],
        edit: ['admin', 'job:edit'],
        del: ['admin', 'job:del']
      },
      // 操作列宽度
      operateWidth: 150
      // exportUrl: "api/expertLectures/export-excel",
      // identity: "expert_lectures",
    }
  },
  methods: {
    // validate(scope, identity) {
    //   validatePoint(scope, identity, this);
    // },
    getBtnSetWidth(width) {
      width += 15
      this.operateWidth = width > this.operateWidth ? width : this.operateWidth
    }
    // validateSelectOptionId,
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
