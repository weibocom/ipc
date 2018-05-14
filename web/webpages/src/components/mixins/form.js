export default {
  methods: {
    validateForm(formName) {
      return new Promise((resolve, reject) => {
        this.$refs[formName].validate(valid => {
          if (valid) {
            resolve(valid)
          } else {
            reject(valid)
          }
        })
      })
    }
  }
}
