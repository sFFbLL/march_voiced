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
        >搜索</el-button>
        <el-button
          v-if="crud.optShow.reset"
          class="filter-item"
          size="mini"
          type="warning"
          icon="el-icon-refresh-left"
          @click="crud.resetQuery(true, role, deptId)"
        >重置</el-button>
      </span>
    </span>
    <el-button-group class="crud-opts-right">
      <el-button
        size="mini"
        plain
        type="info"
        icon="el-icon-search"
        @click="toggleSearch()"
      />

      <el-button size="mini" icon="el-icon-refresh" @click="toQuery" />
      <el-popover placement="bottom-end" width="150" trigger="click">
        <el-button slot="reference" size="mini" icon="el-icon-s-grid">
          <i class="fa fa-caret-down" aria-hidden="true" />
        </el-button>
        <el-checkbox
          v-model="allColumnsSelected"
          :indeterminate="allColumnsSelectedIndeterminate"
          @change="handleCheckAllChange"
        >全选</el-checkbox>
        <el-checkbox
          v-for="item in tableColumns"
          :key="item.property"
          v-model="item.visible"
          @change="handleCheckedTableColumnsChange(item)"
        >{{ item.label }}</el-checkbox>
      </el-popover>
    </el-button-group>
  </div>
</template>
<script>
import CRUD, { crud } from '@crud/crud'

export default {
  mixins: [crud()],
  props: {
    permission: {
      type: Object,
      default: () => {
        return {}
      }
    },
    url: {
      type: String,
      default: () => {
        return ''
      }
    }
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
      return false
    }
  },
  data() {
    return {
      tableColumns: [],
      allColumnsSelected: true,
      allColumnsSelectedIndeterminate: false
    }
  },
  methods: {
    toQuery() {
      this.crud.toQuery()
    },
    toggleSearch() {
      this.crud.props.searchToggle = !this.crud.props.searchToggle
    }
  },
  handleCheckAllChange(val) {
    if (val === false) {
      this.allColumnsSelected = true
      return
    }
    this.tableColumns.forEach((column) => {
      if (!column.visible) {
        column.visible = true
        this.updateColumnVisible(column)
      }
    })
    this.allColumnsSelected = val
    this.allColumnsSelectedIndeterminate = false
  },
  handleCheckedTableColumnsChange(item) {
    let totalCount = 0
    let selectedCount = 0
    this.tableColumns.forEach((column) => {
      ++totalCount
      selectedCount += column.visible ? 1 : 0
    })
    if (selectedCount === 0) {
      this.crud.notify('请至少选择一列', CRUD.NOTIFICATION_TYPE.WARNING)
      this.$nextTick(function() {
        item.visible = true
      })
      return
    }
    this.allColumnsSelected = selectedCount === totalCount
    this.allColumnsSelectedIndeterminate =
      selectedCount !== totalCount && selectedCount !== 0
    this.updateColumnVisible(item)
  },
  updateColumnVisible(item) {
    const table = this.crud.props.table
    const vm = table.$children.find((e) => e.prop === item.property)
    const columnConfig = vm.columnConfig
    if (item.visible) {
      // 找出合适的插入点
      const columnIndex = this.tableColumns.indexOf(item)
      vm.owner.store.commit(
        'insertColumn',
        columnConfig,
        columnIndex + 1,
        null
      )
    } else {
      vm.owner.store.commit('removeColumn', columnConfig, null)
    }
    this.ignoreNextTableColumnsChange = true
  }
}
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
