import types from '../types'
import storage from '../../utils/storage'
import { USER_TYPE, STORAGE_KEY } from '../../utils/constants'
import API from '../../api/api.js'

const state = {
  profile: {}
}

const getters = {
  user: state => state.profile.user || {},
  profile: state => state.profile,
  isAuthenticated: (state, getters) => {
    return !!getters.user.id
  },
  isAdminRole: (state, getters) => {
    return (
      getters.user.admin_type === USER_TYPE.ADMIN ||
      getters.user.admin_type === USER_TYPE.SUPER_ADMIN
    )
  },
  isSuperAdmin: (state, getters) => {
    return getters.user.admin_type === USER_TYPE.SUPER_ADMIN
  }
}

const mutations = {
  [types.CHANGE_PROFILE](state, { profile }) {
    state.profile = profile
    storage.set(STORAGE_KEY.AUTHED, !!profile.user)
    storage.set(STORAGE_KEY.USERNAME, profile.user)
  }
}

const actions = {
  getProfile({ commit }, username) {
    commit(types.CHANGE_PROFILE, {
      profile: { user: username }
    })
  },
  clearProfile({ commit }) {
    commit(types.CHANGE_PROFILE, {
      profile: {}
    })
    storage.clear()
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
