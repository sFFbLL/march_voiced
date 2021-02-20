<template>
  <div class="crud-opts">
    <span class="crud-opts-left">
      <span>
        <el-button
          class="filter-item"
          size="mini"
          type="success"
          icon="el-icon-search"
          @click="toQuery"
          >搜索</el-button
        >
        <el-button
          v-if="crud.optShow.reset"
          class="filter-item"
          size="mini"
          type="warning"
          icon="el-icon-refresh-left"
          @click="crud.resetQuery(true, role, deptId)"
          >重置</el-button
        >
      </span>
    </span>
    <el-button-group class="crud-opts-right">
      <!-- 查询图标 -->
      <el-button
        size="mini"
        plain
        type="info"
        icon="el-icon-search"
        @click="toggleSearch()"
      />
      <!-- 刷新图标 -->
      <el-button size="mini" icon="el-icon-refresh" @click="toQuery" />

      <el-popover placement="bottom-end" width="150" trigger="click">
        <el-button slot="reference" size="mini" icon="el-icon-s-grid">
          <i class="fa fa-caret-down" aria-hidden="true" />
        </el-button>
        <el-checkbox
          v-model="allColumnsSelected"
          :indeterminate="allColumnsSelectedIndeterminate"
          >全选</el-checkbox
        >
        <el-checkbox
          v-for="item of tableColumns"
          :key="item"
          v-model="item.visible"
          :label="item"
          >{{ item }}</el-checkbox
        >
      </el-popover>
    </el-button-group>
  </div>
</template>
<script>
import CRUD, { crud } from "@crud/crud";
export default {
  mixins: [crud()],
  props: {
    permission: {
      type: Object,
      default: () => {
        return {};
      },
    },
    url: {
      type: String,
      default: () => {
        return "";
      },
    },
    // identity: {
    //   type: String,
    //   default: () => {
    //     return "cognizanceObjectStateArray";
    //   },
    // },
  },
  isMyExport: {
    type: Boolean,
    default: () => {
      return false;
    },
  },
  data() {
    return {
      tableColumns: ["上海", "北京", "广州", "深圳"],
      allColumnsSelected: true,
      allColumnsSelectedIndeterminate: false,
    };
  },
  methods: {
    toQuery() {
      this.crud.toQuery();
    },
    toggleSearch() {
      this.crud.props.searchToggle = !this.crud.props.searchToggle;
    },
  },
};
</script>

<style>
.crud-opts {
  padding: 4px 0;
  display: -webkit-flex;
  display: flex;
  align-items: center;
  margin-left: auto;
}
</style>
