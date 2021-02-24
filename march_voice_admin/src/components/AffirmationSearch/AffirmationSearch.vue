<template>
  <div class="crud-opts">
    <span class="crud-opts-left">
      <!--左侧插槽-->
      <slot name="left" />
      <el-button
        class="filter-item"
        size="mini"
        type="success"
        icon="el-icon-search"
        @click="crud.toQuery"
      >搜索</el-button>
      <el-button
        v-if="crud.optShow.reset"
        class="filter-item"
        size="mini"
        type="warning"
        icon="el-icon-refresh-left"
        @click="crud.resetQuery()"
      >重置</el-button>
      <!--右侧-->
      <slot name="right" />
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
          @change="handleCheckAllChange"
        >全选</el-checkbox>
        <el-checkbox
          v-for="item of tableColumns"
          :key="item.property"
          v-model="item.visible"
          @change="handleCheckedTableColumnsChange(item)"
        >
          {{ item.label }}</el-checkbox>
      </el-popover>
    </el-button-group>
  </div>
</template>
<script>
import CRUD, { crud } from '@crud/crud'

function sortWithRef(src, ref) {
  const result = Object.assign([], ref)
  let cursor = -1
  src.forEach(e => {
    const idx = result.indexOf(e)
    if (idx === -1) {
      cursor += 1
      result.splice(cursor, 0, e)
    } else {
      cursor = idx
    }
  })
  return result
}

export default {
  mixins: [crud()],
  props: {
    permission: {
      type: Object,
      default: () => {
        return {}
      }
    },
    ignoreColumns: {
      type: Array,
      default: () => { return [] }
    },
    hiddenColumns: {
      type: Array,
      default: () => { return [] }
    },
    url: {
      type: String,
      default: () => {
        return ''
      }
    }
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
      allColumnsSelectedIndeterminate: false,
      tableUnwatcher: null,
      // 忽略下次表格列变动
      ignoreNextTableColumnsChange: false
    }
  },
  watch: {
    'crud.props.table'() {
      this.updateTableColumns()
      this.tableColumns.forEach(column => {
        if (this.hiddenColumns.indexOf(column.property) !== -1) {
          column.visible = false
          this.updateColumnVisible(column)
        }
      })
    },
    'crud.props.table.store.states.columns'() {
      this.updateTableColumns()
    }
  },
  created() {
    this.crud.updateProp('searchToggle', true)
  },
  methods: {
    updateTableColumns() {
      const table = this.crud.getTable()
      if (!table) {
        this.tableColumns = []
        return
      }
      let cols = null
      const columnFilter = e => e && e.type === 'default' && e.property && this.ignoreColumns.indexOf(e.property) === -1
      const refCols = table.columns.filter(columnFilter)
      if (this.ignoreNextTableColumnsChange) {
        this.ignoreNextTableColumnsChange = false
        return
      }
      this.ignoreNextTableColumnsChange = false
      const columns = []
      const fullTableColumns = table.$children.map(e => e.columnConfig).filter(columnFilter)
      cols = sortWithRef(fullTableColumns, refCols)
      cols.forEach(config => {
        const column = {
          property: config.property,
          label: config.label,
          visible: refCols.indexOf(config) !== -1
        }
        columns.push(column)
      })
      this.tableColumns = columns
    },
    toQuery() {
      console.log(111)
      this.crud.toQuery()
    },
    toggleSearch() {
      this.crud.props.searchToggle = !this.crud.props.searchToggle
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
.crud-opts .crud-opts-right {
  margin-left: auto;
}
</style>
