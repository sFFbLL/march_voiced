// import { getSelectOption, getDept } from '@/api/commonData'

const selectOptionArrays = {
  state: {
    articleTypeArray: [
      {
        key: '1',
        value: '待审核'
      },
      {
        key: '2',
        value: '审核通过'
      },
      {
        key: '3',
        value: '驳回'
      }
    ],
    cognizanceObjectStateArray: [
      {
        key: '1',
        value: '待审核'
      },
      {
        key: '2',
        value: '审核通过'
      },
      {
        key: '3',
        value: '驳回'
      }
    ]
  },

  mutations: {
    SET_SELECT_OPTION_ARRAYS: (state, obj) => {
      state[obj.option] = obj.res
    }
  },
  actions: {
    getAllSelectOptionArrays({ commit, state }) {
      //   const column = [
      //     'articleTypeArray'
      //   ]
      //   for (let index = 0; index < column.length; index++) {
      //     const element = column[index]
      //     new Promise((resolve, reject) => {
      //       getSelectOption(element).then(res => {
      //         const obj = {
      //           option: element,
      //           res: res
      //         }
      //         commit('SET_SELECT_OPTION_ARRAYS', obj)
      //         // commit('SET_SELECT_OPTION_ARRAYS', res)
      //         resolve()
      //       }).catch(error => {
      //         reject(error)
      //       })
      //     })
    }
    // new Promise((resolve, reject) => {
    //     getDept().then(res => {
    //         const obj = {
    //             option: 'dept',
    //             res: res
    //         }
    //         commit('SET_SELECT_OPTION_ARRAYS', obj)
    //         resolve()
    //     }).catch(error => {
    //         reject(error)
    //     })
    // })
  }
}
// }

export default selectOptionArrays
